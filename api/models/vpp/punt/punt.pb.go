// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vpp/punt/punt.proto

package vpp_punt // import "github.com/ligato/vpp-agent/api/models/vpp/punt"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type L3Protocol int32

const (
	L3Protocol_UNDEFINED_L3 L3Protocol = 0
	L3Protocol_IPv4         L3Protocol = 4
	L3Protocol_IPv6         L3Protocol = 6
	L3Protocol_ALL          L3Protocol = 10
)

var L3Protocol_name = map[int32]string{
	0:  "UNDEFINED_L3",
	4:  "IPv4",
	6:  "IPv6",
	10: "ALL",
}
var L3Protocol_value = map[string]int32{
	"UNDEFINED_L3": 0,
	"IPv4":         4,
	"IPv6":         6,
	"ALL":          10,
}

func (x L3Protocol) String() string {
	return proto.EnumName(L3Protocol_name, int32(x))
}
func (L3Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_punt_bc7c8ac0f93efeb8, []int{0}
}

type L4Protocol int32

const (
	L4Protocol_UNDEFINED_L4 L4Protocol = 0
	L4Protocol_TCP          L4Protocol = 6
	L4Protocol_UDP          L4Protocol = 17
)

var L4Protocol_name = map[int32]string{
	0:  "UNDEFINED_L4",
	6:  "TCP",
	17: "UDP",
}
var L4Protocol_value = map[string]int32{
	"UNDEFINED_L4": 0,
	"TCP":          6,
	"UDP":          17,
}

func (x L4Protocol) String() string {
	return proto.EnumName(L4Protocol_name, int32(x))
}
func (L4Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_punt_bc7c8ac0f93efeb8, []int{1}
}

// IpRedirect allows otherwise dropped packet which destination IP address matching some of the VPP addresses
// to redirect to the defined next hop address via the TX interface
type IpRedirect struct {
	L3Protocol           L3Protocol `protobuf:"varint,1,opt,name=l3_protocol,json=l3Protocol,proto3,enum=vpp.punt.L3Protocol" json:"l3_protocol,omitempty"`
	RxInterface          string     `protobuf:"bytes,2,opt,name=rx_interface,json=rxInterface,proto3" json:"rx_interface,omitempty"`
	TxInterface          string     `protobuf:"bytes,3,opt,name=tx_interface,json=txInterface,proto3" json:"tx_interface,omitempty"`
	NextHop              string     `protobuf:"bytes,4,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IpRedirect) Reset()         { *m = IpRedirect{} }
func (m *IpRedirect) String() string { return proto.CompactTextString(m) }
func (*IpRedirect) ProtoMessage()    {}
func (*IpRedirect) Descriptor() ([]byte, []int) {
	return fileDescriptor_punt_bc7c8ac0f93efeb8, []int{0}
}
func (m *IpRedirect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpRedirect.Unmarshal(m, b)
}
func (m *IpRedirect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpRedirect.Marshal(b, m, deterministic)
}
func (dst *IpRedirect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpRedirect.Merge(dst, src)
}
func (m *IpRedirect) XXX_Size() int {
	return xxx_messageInfo_IpRedirect.Size(m)
}
func (m *IpRedirect) XXX_DiscardUnknown() {
	xxx_messageInfo_IpRedirect.DiscardUnknown(m)
}

var xxx_messageInfo_IpRedirect proto.InternalMessageInfo

func (m *IpRedirect) GetL3Protocol() L3Protocol {
	if m != nil {
		return m.L3Protocol
	}
	return L3Protocol_UNDEFINED_L3
}

func (m *IpRedirect) GetRxInterface() string {
	if m != nil {
		return m.RxInterface
	}
	return ""
}

func (m *IpRedirect) GetTxInterface() string {
	if m != nil {
		return m.TxInterface
	}
	return ""
}

func (m *IpRedirect) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (*IpRedirect) XXX_MessageName() string {
	return "vpp.punt.IpRedirect"
}

// allows otherwise dropped packet which destination IP address matching some of the VPP interface IP addresses to be
// punted to the host. L3 and L4 protocols can be used for filtering
type ToHost struct {
	L3Protocol           L3Protocol `protobuf:"varint,2,opt,name=l3_protocol,json=l3Protocol,proto3,enum=vpp.punt.L3Protocol" json:"l3_protocol,omitempty"`
	L4Protocol           L4Protocol `protobuf:"varint,3,opt,name=l4_protocol,json=l4Protocol,proto3,enum=vpp.punt.L4Protocol" json:"l4_protocol,omitempty"`
	Port                 uint32     `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	SocketPath           string     `protobuf:"bytes,5,opt,name=socket_path,json=socketPath,proto3" json:"socket_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ToHost) Reset()         { *m = ToHost{} }
func (m *ToHost) String() string { return proto.CompactTextString(m) }
func (*ToHost) ProtoMessage()    {}
func (*ToHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_punt_bc7c8ac0f93efeb8, []int{1}
}
func (m *ToHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToHost.Unmarshal(m, b)
}
func (m *ToHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToHost.Marshal(b, m, deterministic)
}
func (dst *ToHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToHost.Merge(dst, src)
}
func (m *ToHost) XXX_Size() int {
	return xxx_messageInfo_ToHost.Size(m)
}
func (m *ToHost) XXX_DiscardUnknown() {
	xxx_messageInfo_ToHost.DiscardUnknown(m)
}

var xxx_messageInfo_ToHost proto.InternalMessageInfo

func (m *ToHost) GetL3Protocol() L3Protocol {
	if m != nil {
		return m.L3Protocol
	}
	return L3Protocol_UNDEFINED_L3
}

func (m *ToHost) GetL4Protocol() L4Protocol {
	if m != nil {
		return m.L4Protocol
	}
	return L4Protocol_UNDEFINED_L4
}

func (m *ToHost) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ToHost) GetSocketPath() string {
	if m != nil {
		return m.SocketPath
	}
	return ""
}

func (*ToHost) XXX_MessageName() string {
	return "vpp.punt.ToHost"
}
func init() {
	proto.RegisterType((*IpRedirect)(nil), "vpp.punt.IpRedirect")
	proto.RegisterType((*ToHost)(nil), "vpp.punt.ToHost")
	proto.RegisterEnum("vpp.punt.L3Protocol", L3Protocol_name, L3Protocol_value)
	proto.RegisterEnum("vpp.punt.L4Protocol", L4Protocol_name, L4Protocol_value)
}

func init() { proto.RegisterFile("vpp/punt/punt.proto", fileDescriptor_punt_bc7c8ac0f93efeb8) }

var fileDescriptor_punt_bc7c8ac0f93efeb8 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x8b, 0xda, 0x40,
	0x18, 0xc6, 0x8d, 0xa6, 0x6a, 0x5f, 0x6d, 0x49, 0xa7, 0x3d, 0xa4, 0x3d, 0xb4, 0xd6, 0x93, 0x08,
	0x26, 0xa5, 0x49, 0x4b, 0x41, 0x28, 0xb4, 0xd5, 0x62, 0x20, 0x48, 0x08, 0x7a, 0xd9, 0x4b, 0x88,
	0x71, 0x4c, 0xc2, 0xc6, 0xcc, 0x4b, 0x1c, 0x83, 0x1f, 0x68, 0x6f, 0xfb, 0x45, 0xf6, 0x7b, 0xec,
	0x17, 0x59, 0x32, 0xf1, 0x4f, 0x58, 0x2f, 0x7b, 0x09, 0xcf, 0xf3, 0xbe, 0xcf, 0x2f, 0x3c, 0xcc,
	0x0c, 0xbc, 0xcf, 0x11, 0x75, 0xdc, 0xa7, 0x5c, 0x7c, 0x34, 0xcc, 0x18, 0x67, 0xa4, 0x9d, 0x23,
	0x6a, 0x85, 0xff, 0x34, 0x0a, 0x63, 0x1e, 0xed, 0x57, 0x5a, 0xc0, 0xb6, 0x7a, 0xc8, 0x42, 0xa6,
	0x8b, 0xc0, 0x6a, 0xbf, 0x11, 0x4e, 0x18, 0xa1, 0x4a, 0xb0, 0x7f, 0x27, 0x01, 0x58, 0xe8, 0xd2,
	0x75, 0x9c, 0xd1, 0x80, 0x93, 0x1f, 0xd0, 0x49, 0x0c, 0x4f, 0xac, 0x02, 0x96, 0xa8, 0x52, 0x4f,
	0x1a, 0xbc, 0xfd, 0xfe, 0x41, 0x3b, 0xfd, 0x5d, 0xb3, 0x0d, 0xe7, 0xb8, 0x73, 0x21, 0x39, 0x6b,
	0xf2, 0x15, 0xba, 0xd9, 0xc1, 0x8b, 0x53, 0x4e, 0xb3, 0x8d, 0x1f, 0x50, 0xb5, 0xde, 0x93, 0x06,
	0xaf, 0xdd, 0x4e, 0x76, 0xb0, 0x4e, 0xa3, 0x22, 0xc2, 0xab, 0x91, 0x46, 0x19, 0xe1, 0x95, 0xc8,
	0x47, 0x68, 0xa7, 0xf4, 0xc0, 0xbd, 0x88, 0xa1, 0x2a, 0x8b, 0x75, 0xab, 0xf0, 0x33, 0x86, 0xfd,
	0x7b, 0x09, 0x9a, 0x0b, 0x36, 0x63, 0xbb, 0xab, 0x8a, 0xf5, 0x17, 0x56, 0x2c, 0x30, 0xf3, 0x82,
	0x35, 0xae, 0x30, 0xb3, 0x82, 0x9d, 0x35, 0x21, 0x20, 0x23, 0xcb, 0xb8, 0xe8, 0xf3, 0xc6, 0x15,
	0x9a, 0x7c, 0x81, 0xce, 0x8e, 0x05, 0xb7, 0x94, 0x7b, 0xe8, 0xf3, 0x48, 0x7d, 0x25, 0xaa, 0x42,
	0x39, 0x72, 0x7c, 0x1e, 0x0d, 0xc7, 0x00, 0x97, 0x16, 0x44, 0x81, 0xee, 0x72, 0x3e, 0x99, 0xfe,
	0xb7, 0xe6, 0xd3, 0x89, 0x67, 0x1b, 0x4a, 0x8d, 0xb4, 0x41, 0xb6, 0x9c, 0xdc, 0x54, 0xe4, 0xa3,
	0xfa, 0xa9, 0x34, 0x49, 0x0b, 0x1a, 0x7f, 0x6c, 0x5b, 0x81, 0xe1, 0x37, 0x80, 0x4b, 0x97, 0x67,
	0xb0, 0xa9, 0xd4, 0x8a, 0xe0, 0xe2, 0x9f, 0x53, 0x12, 0xcb, 0x89, 0xa3, 0xbc, 0xfb, 0xfb, 0xfb,
	0xe1, 0xf1, 0xb3, 0x74, 0xf3, 0xab, 0x72, 0xf1, 0x49, 0x1c, 0xfa, 0x9c, 0xe9, 0x39, 0xe2, 0xc8,
	0x0f, 0x69, 0xca, 0x75, 0x1f, 0x63, 0x7d, 0xcb, 0xd6, 0x34, 0xd9, 0xe9, 0xa7, 0xe7, 0x33, 0xce,
	0x11, 0xbd, 0x42, 0xac, 0x9a, 0xe2, 0x54, 0x8c, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0xe9,
	0x37, 0x1a, 0x5a, 0x02, 0x00, 0x00,
}
