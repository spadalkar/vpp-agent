// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vppcalls

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"

	"github.com/ligato/cn-infra/utils/addrs"
	ipsec_api "github.com/ligato/vpp-agent/plugins/vpp/binapi/ipsec"
	"github.com/ligato/vpp-agent/plugins/vpp/model/ipsec"
)

func (handler *ipSecVppHandler) tunnelIfAddDel(tunnel *ipsec.TunnelInterfaces_Tunnel, isAdd bool) (uint32, error) {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsec_api.IpsecTunnelIfAddDel{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	localCryptoKey, err := hex.DecodeString(tunnel.LocalCryptoKey)
	if err != nil {
		return 0, err
	}
	remoteCryptoKey, err := hex.DecodeString(tunnel.RemoteCryptoKey)
	if err != nil {
		return 0, err
	}
	localIntegKey, err := hex.DecodeString(tunnel.LocalIntegKey)
	if err != nil {
		return 0, err
	}
	remoteIntegKey, err := hex.DecodeString(tunnel.RemoteIntegKey)
	if err != nil {
		return 0, err
	}

	req := &ipsec_api.IpsecTunnelIfAddDel{
		IsAdd:              boolToUint(isAdd),
		Esn:                boolToUint(tunnel.Esn),
		AntiReplay:         boolToUint(tunnel.AntiReplay),
		LocalIP:            net.ParseIP(tunnel.LocalIp).To4(),
		RemoteIP:           net.ParseIP(tunnel.RemoteIp).To4(),
		LocalSpi:           tunnel.LocalSpi,
		RemoteSpi:          tunnel.RemoteSpi,
		CryptoAlg:          uint8(tunnel.CryptoAlg),
		LocalCryptoKey:     localCryptoKey,
		LocalCryptoKeyLen:  uint8(len(localCryptoKey)),
		RemoteCryptoKey:    remoteCryptoKey,
		RemoteCryptoKeyLen: uint8(len(remoteCryptoKey)),
		IntegAlg:           uint8(tunnel.IntegAlg),
		LocalIntegKey:      localIntegKey,
		LocalIntegKeyLen:   uint8(len(localIntegKey)),
		RemoteIntegKey:     remoteIntegKey,
		RemoteIntegKeyLen:  uint8(len(remoteIntegKey)),
	}

	reply := &ipsec_api.IpsecTunnelIfAddDelReply{}
	if err := handler.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return 0, err
	}
	if reply.Retval != 0 {
		return 0, fmt.Errorf("%s returned %d", reply.GetMessageName(), reply.Retval)
	}

	return reply.SwIfIndex, nil
}

func (handler *ipSecVppHandler) AddTunnelInterface(tunnel *ipsec.TunnelInterfaces_Tunnel) (uint32, error) {
	return handler.tunnelIfAddDel(tunnel, true)
}

func (handler *ipSecVppHandler) DelTunnelInterface(ifIdx uint32, tunnel *ipsec.TunnelInterfaces_Tunnel) error {
	// Note: ifIdx is not used now, tunnel shiould be matched based on paramters
	_, err := handler.tunnelIfAddDel(tunnel, false)
	return err
}

func (handler *ipSecVppHandler) spdAddDel(spdID uint32, isAdd bool) error {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsec_api.IpsecSpdAddDel{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	req := &ipsec_api.IpsecSpdAddDel{
		IsAdd: boolToUint(isAdd),
		SpdID: spdID,
	}

	reply := &ipsec_api.IpsecSpdAddDelReply{}
	if err := handler.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("%s returned %d", reply.GetMessageName(), reply.Retval)
	}

	return nil
}

func (handler *ipSecVppHandler) AddSPD(spdID uint32) error {
	return handler.spdAddDel(spdID, true)
}

func (handler *ipSecVppHandler) DelSPD(spdID uint32) error {
	return handler.spdAddDel(spdID, false)
}

func (handler *ipSecVppHandler) interfaceAddDelSpd(spdID, swIfIdx uint32, isAdd bool) error {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsec_api.IpsecInterfaceAddDelSpd{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	req := &ipsec_api.IpsecInterfaceAddDelSpd{
		IsAdd:     boolToUint(isAdd),
		SwIfIndex: swIfIdx,
		SpdID:     spdID,
	}

	reply := &ipsec_api.IpsecInterfaceAddDelSpdReply{}
	if err := handler.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("%s returned %d", reply.GetMessageName(), reply.Retval)
	}

	return nil
}

func (handler *ipSecVppHandler) InterfaceAddSPD(spdID, swIfIdx uint32) error {
	return handler.interfaceAddDelSpd(spdID, swIfIdx, true)
}

func (handler *ipSecVppHandler) InterfaceDelSPD(spdID, swIfIdx uint32) error {
	return handler.interfaceAddDelSpd(spdID, swIfIdx, false)
}

func (handler *ipSecVppHandler) spdAddDelEntry(spdID, saID uint32, spd *ipsec.SecurityPolicyDatabases_SPD_PolicyEntry, isAdd bool) error {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsec_api.IpsecSpdAddDelEntry{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	req := &ipsec_api.IpsecSpdAddDelEntry{
		IsAdd:           boolToUint(isAdd),
		SpdID:           spdID,
		Priority:        spd.Priority,
		IsOutbound:      boolToUint(spd.IsOutbound),
		Protocol:        uint8(spd.Protocol),
		RemotePortStart: uint16(spd.RemotePortStart),
		RemotePortStop:  uint16(spd.RemotePortStop),
		LocalPortStart:  uint16(spd.LocalPortStart),
		LocalPortStop:   uint16(spd.LocalPortStop),
		Policy:          uint8(spd.Action),
		SaID:            saID,
	}
	if req.RemotePortStop == 0 {
		req.RemotePortStop = ^req.RemotePortStop
	}
	if req.LocalPortStop == 0 {
		req.LocalPortStop = ^req.LocalPortStop
	}
	if spd.RemoteAddrStart != "" {
		isIPv6, err := addrs.IsIPv6(spd.RemoteAddrStart)
		if err != nil {
			return err
		}
		if isIPv6 {
			req.IsIPv6 = 1
			req.RemoteAddressStart = net.ParseIP(spd.RemoteAddrStart).To16()
			req.RemoteAddressStop = net.ParseIP(spd.RemoteAddrStop).To16()
			req.LocalAddressStart = net.ParseIP(spd.LocalAddrStart).To16()
			req.LocalAddressStop = net.ParseIP(spd.LocalAddrStop).To16()
		} else {
			req.IsIPv6 = 0
			req.RemoteAddressStart = net.ParseIP(spd.RemoteAddrStart).To4()
			req.RemoteAddressStop = net.ParseIP(spd.RemoteAddrStop).To4()
			req.LocalAddressStart = net.ParseIP(spd.LocalAddrStart).To4()
			req.LocalAddressStop = net.ParseIP(spd.LocalAddrStop).To4()
		}
	} else {
		req.RemoteAddressStart = net.ParseIP("0.0.0.0").To4()
		req.RemoteAddressStop = net.ParseIP("255.255.255.255").To4()
		req.LocalAddressStart = net.ParseIP("0.0.0.0").To4()
		req.LocalAddressStop = net.ParseIP("255.255.255.255").To4()
	}

	reply := &ipsec_api.IpsecSpdAddDelEntryReply{}
	if err := handler.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("%s returned %d", reply.GetMessageName(), reply.Retval)
	}

	return nil
}

func (handler *ipSecVppHandler) AddSPDEntry(spdID, saID uint32, spd *ipsec.SecurityPolicyDatabases_SPD_PolicyEntry) error {
	return handler.spdAddDelEntry(spdID, saID, spd, true)
}

func (handler *ipSecVppHandler) DelSPDEntry(spdID, saID uint32, spd *ipsec.SecurityPolicyDatabases_SPD_PolicyEntry) error {
	return handler.spdAddDelEntry(spdID, saID, spd, false)
}

func (handler *ipSecVppHandler) sadAddDelEntry(saID uint32, sa *ipsec.SecurityAssociations_SA, isAdd bool) error {
	defer func(t time.Time) {
		handler.stopwatch.TimeLog(ipsec_api.IpsecSadAddDelEntry{}).LogTimeEntry(time.Since(t))
	}(time.Now())

	cryptoKey, err := hex.DecodeString(sa.CryptoKey)
	if err != nil {
		return err
	}
	integKey, err := hex.DecodeString(sa.IntegKey)
	if err != nil {
		return err
	}

	req := &ipsec_api.IpsecSadAddDelEntry{
		IsAdd:                     boolToUint(isAdd),
		SadID:                     saID,
		Spi:                       sa.Spi,
		Protocol:                  uint8(sa.Protocol),
		CryptoAlgorithm:           uint8(sa.CryptoAlg),
		CryptoKey:                 cryptoKey,
		CryptoKeyLength:           uint8(len(cryptoKey)),
		IntegrityAlgorithm:        uint8(sa.IntegAlg),
		IntegrityKey:              integKey,
		IntegrityKeyLength:        uint8(len(integKey)),
		UseExtendedSequenceNumber: boolToUint(sa.UseEsn),
		UseAntiReplay:             boolToUint(sa.UseAntiReplay),
		UDPEncap:                  boolToUint(sa.EnableUdpEncap),
	}
	if sa.TunnelSrcAddr != "" {
		req.IsTunnel = 1
		isIPv6, err := addrs.IsIPv6(sa.TunnelSrcAddr)
		if err != nil {
			return err
		}
		if isIPv6 {
			req.IsTunnelIPv6 = 1
			req.TunnelSrcAddress = net.ParseIP(sa.TunnelSrcAddr).To16()
			req.TunnelDstAddress = net.ParseIP(sa.TunnelDstAddr).To16()
		} else {
			req.IsTunnelIPv6 = 0
			req.TunnelSrcAddress = net.ParseIP(sa.TunnelSrcAddr).To4()
			req.TunnelDstAddress = net.ParseIP(sa.TunnelDstAddr).To4()
		}
	}

	reply := &ipsec_api.IpsecSadAddDelEntryReply{}
	if err := handler.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}
	if reply.Retval != 0 {
		return fmt.Errorf("%s returned %d", reply.GetMessageName(), reply.Retval)
	}

	return nil
}

func (handler *ipSecVppHandler) AddSAEntry(saID uint32, sa *ipsec.SecurityAssociations_SA) error {
	return handler.sadAddDelEntry(saID, sa, true)
}

func (handler *ipSecVppHandler) DelSAEntry(saID uint32, sa *ipsec.SecurityAssociations_SA) error {
	return handler.sadAddDelEntry(saID, sa, false)
}

func boolToUint(value bool) uint8 {
	if value {
		return 1
	}
	return 0
}
