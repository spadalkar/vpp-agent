// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// source: /usr/share/vpp/api/plugins/stn.api.json

/*
Package stn is a generated VPP binary API for 'stn' module.

It consists of:
	 10 enums
	  6 aliases
	  6 types
	  1 union
	  4 messages
	  2 services
*/
package stn

import (
	"bytes"
	"context"
	"io"
	"strconv"

	api "git.fd.io/govpp.git/api"
	struc "github.com/lunixbochs/struc"

	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/interface_types"
	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2001/ip_types"
)

const (
	// ModuleName is the name of this module.
	ModuleName = "stn"
	// APIVersion is the API version of this module.
	APIVersion = "2.0.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x619d8f3
)

type AddressFamily = ip_types.AddressFamily

type IfStatusFlags = interface_types.IfStatusFlags

type IfType = interface_types.IfType

type IPDscp = ip_types.IPDscp

type IPEcn = ip_types.IPEcn

type IPProto = ip_types.IPProto

type LinkDuplex = interface_types.LinkDuplex

type MtuProto = interface_types.MtuProto

type RxMode = interface_types.RxMode

type SubIfFlags = interface_types.SubIfFlags

type AddressWithPrefix = ip_types.AddressWithPrefix

type InterfaceIndex = interface_types.InterfaceIndex

type IP4Address = ip_types.IP4Address

type IP4AddressWithPrefix = ip_types.IP4AddressWithPrefix

type IP6Address = ip_types.IP6Address

type IP6AddressWithPrefix = ip_types.IP6AddressWithPrefix

type Address = ip_types.Address

type IP4Prefix = ip_types.IP4Prefix

type IP6Prefix = ip_types.IP6Prefix

type Mprefix = ip_types.Mprefix

type Prefix = ip_types.Prefix

type PrefixMatcher = ip_types.PrefixMatcher

type AddressUnion = ip_types.AddressUnion

// StnAddDelRule represents VPP binary API message 'stn_add_del_rule'.
type StnAddDelRule struct {
	IPAddress Address
	SwIfIndex InterfaceIndex
	IsAdd     bool
}

func (m *StnAddDelRule) Reset()                        { *m = StnAddDelRule{} }
func (*StnAddDelRule) GetMessageName() string          { return "stn_add_del_rule" }
func (*StnAddDelRule) GetCrcString() string            { return "53f751e6" }
func (*StnAddDelRule) GetMessageType() api.MessageType { return api.RequestMessage }

// StnAddDelRuleReply represents VPP binary API message 'stn_add_del_rule_reply'.
type StnAddDelRuleReply struct {
	Retval int32
}

func (m *StnAddDelRuleReply) Reset()                        { *m = StnAddDelRuleReply{} }
func (*StnAddDelRuleReply) GetMessageName() string          { return "stn_add_del_rule_reply" }
func (*StnAddDelRuleReply) GetCrcString() string            { return "e8d4e804" }
func (*StnAddDelRuleReply) GetMessageType() api.MessageType { return api.ReplyMessage }

// StnRulesDetails represents VPP binary API message 'stn_rules_details'.
type StnRulesDetails struct {
	IPAddress Address
	SwIfIndex InterfaceIndex
}

func (m *StnRulesDetails) Reset()                        { *m = StnRulesDetails{} }
func (*StnRulesDetails) GetMessageName() string          { return "stn_rules_details" }
func (*StnRulesDetails) GetCrcString() string            { return "b0f6606c" }
func (*StnRulesDetails) GetMessageType() api.MessageType { return api.ReplyMessage }

// StnRulesDump represents VPP binary API message 'stn_rules_dump'.
type StnRulesDump struct{}

func (m *StnRulesDump) Reset()                        { *m = StnRulesDump{} }
func (*StnRulesDump) GetMessageName() string          { return "stn_rules_dump" }
func (*StnRulesDump) GetCrcString() string            { return "51077d14" }
func (*StnRulesDump) GetMessageType() api.MessageType { return api.RequestMessage }

func init() {
	api.RegisterMessage((*StnAddDelRule)(nil), "stn.StnAddDelRule")
	api.RegisterMessage((*StnAddDelRuleReply)(nil), "stn.StnAddDelRuleReply")
	api.RegisterMessage((*StnRulesDetails)(nil), "stn.StnRulesDetails")
	api.RegisterMessage((*StnRulesDump)(nil), "stn.StnRulesDump")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*StnAddDelRule)(nil),
		(*StnAddDelRuleReply)(nil),
		(*StnRulesDetails)(nil),
		(*StnRulesDump)(nil),
	}
}

// RPCService represents RPC service API for stn module.
type RPCService interface {
	DumpStnRules(ctx context.Context, in *StnRulesDump) (RPCService_DumpStnRulesClient, error)
	StnAddDelRule(ctx context.Context, in *StnAddDelRule) (*StnAddDelRuleReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpStnRules(ctx context.Context, in *StnRulesDump) (RPCService_DumpStnRulesClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpStnRulesClient{stream}
	return x, nil
}

type RPCService_DumpStnRulesClient interface {
	Recv() (*StnRulesDetails, error)
}

type serviceClient_DumpStnRulesClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpStnRulesClient) Recv() (*StnRulesDetails, error) {
	m := new(StnRulesDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) StnAddDelRule(ctx context.Context, in *StnAddDelRule) (*StnAddDelRuleReply, error) {
	out := new(StnAddDelRuleReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion1 // please upgrade the GoVPP api package

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
