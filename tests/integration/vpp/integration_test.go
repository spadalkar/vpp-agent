//  Copyright (c) 2018 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
	"time"

	"git.fd.io/govpp.git/adapter/socketclient"
	"git.fd.io/govpp.git/adapter/statsclient"
	govppapi "git.fd.io/govpp.git/api"
	govppcore "git.fd.io/govpp.git/core"
	"github.com/mitchellh/go-ps"
	. "github.com/onsi/gomega"

	"go.ligato.io/vpp-agent/v3/plugins/govppmux/vppcalls"
	"go.ligato.io/vpp-agent/v3/plugins/vpp"
)

const (
	vppConnectRetryDelay = time.Millisecond * 500
	vppBootDelay         = time.Millisecond * 200
	vppTermDelay         = time.Millisecond * 50
	vppExitTimeout       = time.Second * 1

	defaultVPPConfig = `
		unix {
			nodaemon
			cli-listen /run/vpp/cli.sock
			cli-no-pager
			log /tmp/vpp.log
			full-coredump
		}
		api-trace {
			on
			save-api-table /run/vpp/api-table
		}
		socksvr {
			socket-name /run/vpp/api.sock
		}
		statseg {
			socket-name /run/vpp/stats.sock
			per-node-counters on
		}
		plugins {
			plugin dpdk_plugin.so { disable }
		}`
)

type TestCtx struct {
	t              *testing.T
	Ctx            context.Context
	vppCmd         *exec.Cmd
	stderr, stdout *bytes.Buffer
	Conn           *govppcore.Connection
	StatsConn      *govppcore.StatsConnection
	vppBinapi      govppapi.Channel
	vppStats       govppapi.StatsProvider
	vpp            vppcalls.VppCoreAPI
	versionInfo    *vppcalls.VersionInfo
	vppClient      *vppClient
}

func startVPP(t *testing.T, stdout, stderr io.Writer) *exec.Cmd {
	// check if VPP process is not running already
	processes, err := ps.Processes()
	if err != nil {
		t.Fatalf("listing processes failed: %v", err)
	}
	for _, process := range processes {
		proc := process.Executable()
		if strings.Contains(proc, "vpp") && process.Pid() != os.Getpid() {
			t.Logf(" - found process: %+v", process)
		}
		switch proc {
		case *vppPath, "vpp", "vpp_main":
			t.Fatalf("VPP is already running (PID: %v)", process.Pid())
		}
	}

	// remove binapi files from previous run
	var removeFile = func(path string) {
		if err := os.Remove(path); err == nil {
			t.Logf("removed file %q", path)
		} else if !os.IsNotExist(err) {
			t.Fatalf("removing file %q failed: %v", path, err)
		}
	}
	removeFile(*vppSockAddr)

	// ensure VPP runtime directory exists
	if err := os.Mkdir("/run/vpp", 0755); err != nil && !os.IsExist(err) {
		t.Logf("mkdir failed: %v", err)
	}

	// setup VPP process
	vppCmd := exec.Command(*vppPath)
	if *vppConfig != "" {
		vppCmd.Args = append(vppCmd.Args, "-c", *vppConfig)
	} else {
		vppCmd.Args = append(vppCmd.Args, defaultVPPConfig)
	}
	vppCmd.Stderr = stderr
	vppCmd.Stdout = stdout

	// ensure that process is killed when current process exits
	vppCmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGKILL}

	if err := vppCmd.Start(); err != nil {
		t.Fatalf("starting VPP failed: %v", err)
	}

	t.Logf("VPP start OK (PID: %v)", vppCmd.Process.Pid)
	return vppCmd
}

func setupVPP(t *testing.T) *TestCtx {
	if os.Getenv("TRAVIS") != "" {
		t.Skip("skipping test for Travis")
	}
	RegisterTestingT(t)

	start := time.Now()

	ctx := context.TODO()

	// start VPP process
	var stderr, stdout bytes.Buffer
	vppCmd := startVPP(t, &stdout, &stderr)
	vppPID := uint32(vppCmd.Process.Pid)

	// if setupVPP fails we need stop the VPP process
	defer func() {
		if t.Failed() {
			stopVPP(t, vppCmd)
		}
	}()

	// wait until the socket is ready
	adapter := socketclient.NewVppClient(*vppSockAddr)
	if err := adapter.WaitReady(); err != nil {
		t.Logf("WaitReady error: %v", err)
	}
	time.Sleep(vppBootDelay)

	connectRetry := func(retries int) (conn *govppcore.Connection, err error) {
		for i := 1; i <= retries; i++ {
			conn, err = govppcore.Connect(adapter)
			if err != nil {
				t.Logf("attempt #%d failed: %v, retrying in %v", i, err, vppConnectRetryDelay)
				time.Sleep(vppConnectRetryDelay)
				continue
			}
			return
		}
		return nil, fmt.Errorf("failed to connect after %d retries", retries)
	}

	// connect to binapi
	conn, err := connectRetry(int(*vppRetry))
	if err != nil {
		t.Errorf("connecting to VPP failed: %v", err)
		if err := vppCmd.Process.Kill(); err != nil {
			t.Fatalf("killing VPP failed: %v", err)
		}
		if state, err := vppCmd.Process.Wait(); err != nil {
			t.Logf("VPP wait failed: %v", err)
		} else {
			t.Logf("VPP wait OK: %v", state)
		}
		t.FailNow()
	}

	apiChannel, err := conn.NewAPIChannel()
	if err != nil {
		t.Fatalf("creating channel failed: %v", err)
	}

	vppClient := &vppClient{
		t:               t,
		ChannelProvider: conn,
		ch:              apiChannel,
	}

	vpeHandler := vppcalls.CompatibleHandler(vppClient)

	// retrieve VPP version
	versionInfo, err := vpeHandler.GetVersion(ctx)
	if err != nil {
		t.Fatalf("getting version info failed: %v", err)
	}
	t.Logf("VPP version: %v", versionInfo.Version)
	if versionInfo.Version == "" {
		t.Fatal("expected VPP version to not be empty")
	}
	// verify connected session
	vpeInfo, err := vpeHandler.GetSession(ctx)
	if err != nil {
		t.Fatalf("getting vpp info failed: %v", err)
	}
	if vpeInfo.PID != vppPID {
		t.Fatalf("expected VPP PID to be %v, got %v", vppPID, vpeInfo.PID)
	}

	vppClient.vpp = vpeHandler

	// connect to stats
	statsClient := statsclient.NewStatsClient("")
	statsConn, err := govppcore.ConnectStats(statsClient)
	if err != nil {
		t.Logf("connecting to VPP stats API failed: %v", err)
	} else {
		vppClient.stats = statsConn
	}

	t.Logf("-> VPP ready (took %v)", time.Since(start).Seconds())

	return &TestCtx{
		t:           t,
		Ctx:         ctx,
		versionInfo: versionInfo,
		vpp:         vpeHandler,
		vppCmd:      vppCmd,
		stderr:      &stderr,
		stdout:      &stdout,
		Conn:        conn,
		vppBinapi:   apiChannel,
		vppStats:    statsConn,
		vppClient:   vppClient,
	}
}

func (ctx *TestCtx) teardownVPP() {
	// disconnect sometimes hangs
	done := make(chan struct{})
	go func() {
		ctx.StatsConn.Disconnect()
		ctx.vppBinapi.Close()
		ctx.Conn.Disconnect()
		close(done)
	}()
	select {
	case <-done:
		time.Sleep(vppTermDelay)
	case <-time.After(vppExitTimeout):
		ctx.t.Logf("VPP disconnect timeout")
	}

	stopVPP(ctx.t, ctx.vppCmd)
}

func stopVPP(t *testing.T, vppCmd *exec.Cmd) {
	if err := vppCmd.Process.Signal(syscall.SIGTERM); err != nil {
		t.Fatalf("sending SIGTERM to VPP failed: %v", err)
	}
	// wait until VPP exits
	exit := make(chan struct{})
	go func() {
		if err := vppCmd.Wait(); err != nil {
			var exiterr *exec.ExitError
			if errors.As(err, &exiterr) && strings.Contains(exiterr.Error(), "core dumped") {
				t.Logf("VPP process CRASHED: %s", exiterr.Error())
			} else {
				t.Logf("VPP process wait failed: %v", err)
			}
		} else {
			t.Logf("VPP exit OK")
		}
		close(exit)
	}()
	select {
	case <-exit:
		// exited
	case <-time.After(vppExitTimeout):
		t.Logf("VPP exit timeout")
		t.Logf("sending SIGKILL to VPP..")
		if err := vppCmd.Process.Signal(syscall.SIGKILL); err != nil {
			t.Fatalf("sending SIGKILL to VPP failed: %v", err)
		}
	}
}

type vppClient struct {
	t *testing.T
	govppapi.ChannelProvider
	ch      govppapi.Channel
	stats   govppapi.StatsProvider
	vpp     vppcalls.VppCoreAPI
	version vpp.Version
}

func (v *vppClient) Version() vpp.Version {
	return v.version
}

func (v *vppClient) BinapiVersion() vpp.Version {
	return ""
}

func (v *vppClient) CheckCompatiblity(msgs ...govppapi.Message) error {
	return v.ch.CheckCompatiblity(msgs...)
}

func (v *vppClient) Stats() govppapi.StatsProvider {
	return v.stats
}

func (v *vppClient) IsPluginLoaded(plugin string) bool {
	ctx := context.Background()
	plugins, err := v.vpp.GetPlugins(ctx)
	if err != nil {
		v.t.Fatalf("GetPlugins failed: %v", plugins)
	}
	for _, p := range plugins {
		if p.Name == plugin {
			return true
		}
	}
	return false
}

func (v *vppClient) OnReconnect(h func()) {
	// no-op
}
