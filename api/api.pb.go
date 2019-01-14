// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package api // import "github.com/ligato/vpp-agent/api"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Module struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Specs                []*ModelSpec `protobuf:"bytes,2,rep,name=specs,proto3" json:"specs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Module) Reset()         { *m = Module{} }
func (m *Module) String() string { return proto.CompactTextString(m) }
func (*Module) ProtoMessage()    {}
func (*Module) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{0}
}
func (m *Module) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Module.Unmarshal(m, b)
}
func (m *Module) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Module.Marshal(b, m, deterministic)
}
func (dst *Module) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Module.Merge(dst, src)
}
func (m *Module) XXX_Size() int {
	return xxx_messageInfo_Module.Size(m)
}
func (m *Module) XXX_DiscardUnknown() {
	xxx_messageInfo_Module.DiscardUnknown(m)
}

var xxx_messageInfo_Module proto.InternalMessageInfo

func (m *Module) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Module) GetSpecs() []*ModelSpec {
	if m != nil {
		return m.Specs
	}
	return nil
}

func (*Module) XXX_MessageName() string {
	return "api.Module"
}

type ModelSpec struct {
	Version              string            `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Class                string            `protobuf:"bytes,2,opt,name=class,proto3" json:"class,omitempty"`
	Module               string            `protobuf:"bytes,3,opt,name=module,proto3" json:"module,omitempty"`
	Type                 string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,5,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ModelSpec) Reset()         { *m = ModelSpec{} }
func (m *ModelSpec) String() string { return proto.CompactTextString(m) }
func (*ModelSpec) ProtoMessage()    {}
func (*ModelSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{1}
}
func (m *ModelSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelSpec.Unmarshal(m, b)
}
func (m *ModelSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelSpec.Marshal(b, m, deterministic)
}
func (dst *ModelSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelSpec.Merge(dst, src)
}
func (m *ModelSpec) XXX_Size() int {
	return xxx_messageInfo_ModelSpec.Size(m)
}
func (m *ModelSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ModelSpec proto.InternalMessageInfo

func (m *ModelSpec) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ModelSpec) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *ModelSpec) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *ModelSpec) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ModelSpec) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (*ModelSpec) XXX_MessageName() string {
	return "api.ModelSpec"
}

type Model struct {
	Any                  *types.Any        `protobuf:"bytes,1,opt,name=any,proto3" json:"any,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{2}
}
func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (dst *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(dst, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

func (m *Model) GetAny() *types.Any {
	if m != nil {
		return m.Any
	}
	return nil
}

func (m *Model) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (*Model) XXX_MessageName() string {
	return "api.Model"
}

type ListModulesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListModulesRequest) Reset()         { *m = ListModulesRequest{} }
func (m *ListModulesRequest) String() string { return proto.CompactTextString(m) }
func (*ListModulesRequest) ProtoMessage()    {}
func (*ListModulesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{3}
}
func (m *ListModulesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListModulesRequest.Unmarshal(m, b)
}
func (m *ListModulesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListModulesRequest.Marshal(b, m, deterministic)
}
func (dst *ListModulesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListModulesRequest.Merge(dst, src)
}
func (m *ListModulesRequest) XXX_Size() int {
	return xxx_messageInfo_ListModulesRequest.Size(m)
}
func (m *ListModulesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListModulesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListModulesRequest proto.InternalMessageInfo

func (*ListModulesRequest) XXX_MessageName() string {
	return "api.ListModulesRequest"
}

type ListModulesResponse struct {
	Modules              []*Module `protobuf:"bytes,1,rep,name=modules,proto3" json:"modules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListModulesResponse) Reset()         { *m = ListModulesResponse{} }
func (m *ListModulesResponse) String() string { return proto.CompactTextString(m) }
func (*ListModulesResponse) ProtoMessage()    {}
func (*ListModulesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{4}
}
func (m *ListModulesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListModulesResponse.Unmarshal(m, b)
}
func (m *ListModulesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListModulesResponse.Marshal(b, m, deterministic)
}
func (dst *ListModulesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListModulesResponse.Merge(dst, src)
}
func (m *ListModulesResponse) XXX_Size() int {
	return xxx_messageInfo_ListModulesResponse.Size(m)
}
func (m *ListModulesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListModulesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListModulesResponse proto.InternalMessageInfo

func (m *ListModulesResponse) GetModules() []*Module {
	if m != nil {
		return m.Modules
	}
	return nil
}

func (*ListModulesResponse) XXX_MessageName() string {
	return "api.ListModulesResponse"
}

type SyncItem struct {
	Model                *Model   `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Delete               bool     `protobuf:"varint,2,opt,name=delete,proto3" json:"delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncItem) Reset()         { *m = SyncItem{} }
func (m *SyncItem) String() string { return proto.CompactTextString(m) }
func (*SyncItem) ProtoMessage()    {}
func (*SyncItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{5}
}
func (m *SyncItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncItem.Unmarshal(m, b)
}
func (m *SyncItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncItem.Marshal(b, m, deterministic)
}
func (dst *SyncItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncItem.Merge(dst, src)
}
func (m *SyncItem) XXX_Size() int {
	return xxx_messageInfo_SyncItem.Size(m)
}
func (m *SyncItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncItem.DiscardUnknown(m)
}

var xxx_messageInfo_SyncItem proto.InternalMessageInfo

func (m *SyncItem) GetModel() *Model {
	if m != nil {
		return m.Model
	}
	return nil
}

func (m *SyncItem) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

func (*SyncItem) XXX_MessageName() string {
	return "api.SyncItem"
}

type SyncOptions struct {
	Resync               bool     `protobuf:"varint,1,opt,name=resync,proto3" json:"resync,omitempty"`
	RetryCount           uint32   `protobuf:"varint,2,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncOptions) Reset()         { *m = SyncOptions{} }
func (m *SyncOptions) String() string { return proto.CompactTextString(m) }
func (*SyncOptions) ProtoMessage()    {}
func (*SyncOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{6}
}
func (m *SyncOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncOptions.Unmarshal(m, b)
}
func (m *SyncOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncOptions.Marshal(b, m, deterministic)
}
func (dst *SyncOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncOptions.Merge(dst, src)
}
func (m *SyncOptions) XXX_Size() int {
	return xxx_messageInfo_SyncOptions.Size(m)
}
func (m *SyncOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncOptions.DiscardUnknown(m)
}

var xxx_messageInfo_SyncOptions proto.InternalMessageInfo

func (m *SyncOptions) GetResync() bool {
	if m != nil {
		return m.Resync
	}
	return false
}

func (m *SyncOptions) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (*SyncOptions) XXX_MessageName() string {
	return "api.SyncOptions"
}

type SyncRequest struct {
	Items                []*SyncItem  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Options              *SyncOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{7}
}
func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRequest.Unmarshal(m, b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
}
func (dst *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(dst, src)
}
func (m *SyncRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRequest.Size(m)
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetItems() []*SyncItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SyncRequest) GetOptions() *SyncOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (*SyncRequest) XXX_MessageName() string {
	return "api.SyncRequest"
}

type SyncResponse struct {
	Results              []*SyncItemResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SyncResponse) Reset()         { *m = SyncResponse{} }
func (m *SyncResponse) String() string { return proto.CompactTextString(m) }
func (*SyncResponse) ProtoMessage()    {}
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{8}
}
func (m *SyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncResponse.Unmarshal(m, b)
}
func (m *SyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncResponse.Marshal(b, m, deterministic)
}
func (dst *SyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncResponse.Merge(dst, src)
}
func (m *SyncResponse) XXX_Size() int {
	return xxx_messageInfo_SyncResponse.Size(m)
}
func (m *SyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncResponse proto.InternalMessageInfo

func (m *SyncResponse) GetResults() []*SyncItemResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func (*SyncResponse) XXX_MessageName() string {
	return "api.SyncResponse"
}

type SyncItemResult struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncItemResult) Reset()         { *m = SyncItemResult{} }
func (m *SyncItemResult) String() string { return proto.CompactTextString(m) }
func (*SyncItemResult) ProtoMessage()    {}
func (*SyncItemResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{9}
}
func (m *SyncItemResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncItemResult.Unmarshal(m, b)
}
func (m *SyncItemResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncItemResult.Marshal(b, m, deterministic)
}
func (dst *SyncItemResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncItemResult.Merge(dst, src)
}
func (m *SyncItemResult) XXX_Size() int {
	return xxx_messageInfo_SyncItemResult.Size(m)
}
func (m *SyncItemResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncItemResult.DiscardUnknown(m)
}

var xxx_messageInfo_SyncItemResult proto.InternalMessageInfo

func (m *SyncItemResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SyncItemResult) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SyncItemResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (*SyncItemResult) XXX_MessageName() string {
	return "api.SyncItemResult"
}

type ObtainOptions struct {
	ObtainAll            bool     `protobuf:"varint,1,opt,name=obtain_all,json=obtainAll,proto3" json:"obtain_all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObtainOptions) Reset()         { *m = ObtainOptions{} }
func (m *ObtainOptions) String() string { return proto.CompactTextString(m) }
func (*ObtainOptions) ProtoMessage()    {}
func (*ObtainOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{10}
}
func (m *ObtainOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObtainOptions.Unmarshal(m, b)
}
func (m *ObtainOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObtainOptions.Marshal(b, m, deterministic)
}
func (dst *ObtainOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObtainOptions.Merge(dst, src)
}
func (m *ObtainOptions) XXX_Size() int {
	return xxx_messageInfo_ObtainOptions.Size(m)
}
func (m *ObtainOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ObtainOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ObtainOptions proto.InternalMessageInfo

func (m *ObtainOptions) GetObtainAll() bool {
	if m != nil {
		return m.ObtainAll
	}
	return false
}

func (*ObtainOptions) XXX_MessageName() string {
	return "api.ObtainOptions"
}

type ObtainRequest struct {
	Ids                  []string       `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	Specs                []string       `protobuf:"bytes,2,rep,name=specs,proto3" json:"specs,omitempty"`
	Options              *ObtainOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ObtainRequest) Reset()         { *m = ObtainRequest{} }
func (m *ObtainRequest) String() string { return proto.CompactTextString(m) }
func (*ObtainRequest) ProtoMessage()    {}
func (*ObtainRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{11}
}
func (m *ObtainRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObtainRequest.Unmarshal(m, b)
}
func (m *ObtainRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObtainRequest.Marshal(b, m, deterministic)
}
func (dst *ObtainRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObtainRequest.Merge(dst, src)
}
func (m *ObtainRequest) XXX_Size() int {
	return xxx_messageInfo_ObtainRequest.Size(m)
}
func (m *ObtainRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObtainRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObtainRequest proto.InternalMessageInfo

func (m *ObtainRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ObtainRequest) GetSpecs() []string {
	if m != nil {
		return m.Specs
	}
	return nil
}

func (m *ObtainRequest) GetOptions() *ObtainOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (*ObtainRequest) XXX_MessageName() string {
	return "api.ObtainRequest"
}

type ObtainResponse struct {
	Models               []*ModelInfo `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ObtainResponse) Reset()         { *m = ObtainResponse{} }
func (m *ObtainResponse) String() string { return proto.CompactTextString(m) }
func (*ObtainResponse) ProtoMessage()    {}
func (*ObtainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{12}
}
func (m *ObtainResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObtainResponse.Unmarshal(m, b)
}
func (m *ObtainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObtainResponse.Marshal(b, m, deterministic)
}
func (dst *ObtainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObtainResponse.Merge(dst, src)
}
func (m *ObtainResponse) XXX_Size() int {
	return xxx_messageInfo_ObtainResponse.Size(m)
}
func (m *ObtainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ObtainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ObtainResponse proto.InternalMessageInfo

func (m *ObtainResponse) GetModels() []*ModelInfo {
	if m != nil {
		return m.Models
	}
	return nil
}

func (*ObtainResponse) XXX_MessageName() string {
	return "api.ObtainResponse"
}

type ModelInfo struct {
	Spec                 *ModelSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Model                *Model     `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Status               string     `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Error                string     `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ModelInfo) Reset()         { *m = ModelInfo{} }
func (m *ModelInfo) String() string { return proto.CompactTextString(m) }
func (*ModelInfo) ProtoMessage()    {}
func (*ModelInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_5fb666e15fab7b09, []int{13}
}
func (m *ModelInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelInfo.Unmarshal(m, b)
}
func (m *ModelInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelInfo.Marshal(b, m, deterministic)
}
func (dst *ModelInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelInfo.Merge(dst, src)
}
func (m *ModelInfo) XXX_Size() int {
	return xxx_messageInfo_ModelInfo.Size(m)
}
func (m *ModelInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ModelInfo proto.InternalMessageInfo

func (m *ModelInfo) GetSpec() *ModelSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *ModelInfo) GetModel() *Model {
	if m != nil {
		return m.Model
	}
	return nil
}

func (m *ModelInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ModelInfo) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (*ModelInfo) XXX_MessageName() string {
	return "api.ModelInfo"
}
func init() {
	proto.RegisterType((*Module)(nil), "api.Module")
	proto.RegisterType((*ModelSpec)(nil), "api.ModelSpec")
	proto.RegisterMapType((map[string]string)(nil), "api.ModelSpec.MetaEntry")
	proto.RegisterType((*Model)(nil), "api.Model")
	proto.RegisterMapType((map[string]string)(nil), "api.Model.MetadataEntry")
	proto.RegisterType((*ListModulesRequest)(nil), "api.ListModulesRequest")
	proto.RegisterType((*ListModulesResponse)(nil), "api.ListModulesResponse")
	proto.RegisterType((*SyncItem)(nil), "api.SyncItem")
	proto.RegisterType((*SyncOptions)(nil), "api.SyncOptions")
	proto.RegisterType((*SyncRequest)(nil), "api.SyncRequest")
	proto.RegisterType((*SyncResponse)(nil), "api.SyncResponse")
	proto.RegisterType((*SyncItemResult)(nil), "api.SyncItemResult")
	proto.RegisterType((*ObtainOptions)(nil), "api.ObtainOptions")
	proto.RegisterType((*ObtainRequest)(nil), "api.ObtainRequest")
	proto.RegisterType((*ObtainResponse)(nil), "api.ObtainResponse")
	proto.RegisterType((*ModelInfo)(nil), "api.ModelInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncServiceClient interface {
	ListModules(ctx context.Context, in *ListModulesRequest, opts ...grpc.CallOption) (*ListModulesResponse, error)
	Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error)
	Obtain(ctx context.Context, in *ObtainRequest, opts ...grpc.CallOption) (*ObtainResponse, error)
}

type syncServiceClient struct {
	cc *grpc.ClientConn
}

func NewSyncServiceClient(cc *grpc.ClientConn) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) ListModules(ctx context.Context, in *ListModulesRequest, opts ...grpc.CallOption) (*ListModulesResponse, error) {
	out := new(ListModulesResponse)
	err := c.cc.Invoke(ctx, "/api.SyncService/ListModules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) Sync(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (*SyncResponse, error) {
	out := new(SyncResponse)
	err := c.cc.Invoke(ctx, "/api.SyncService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) Obtain(ctx context.Context, in *ObtainRequest, opts ...grpc.CallOption) (*ObtainResponse, error) {
	out := new(ObtainResponse)
	err := c.cc.Invoke(ctx, "/api.SyncService/Obtain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServiceServer is the server API for SyncService service.
type SyncServiceServer interface {
	ListModules(context.Context, *ListModulesRequest) (*ListModulesResponse, error)
	Sync(context.Context, *SyncRequest) (*SyncResponse, error)
	Obtain(context.Context, *ObtainRequest) (*ObtainResponse, error)
}

func RegisterSyncServiceServer(s *grpc.Server, srv SyncServiceServer) {
	s.RegisterService(&_SyncService_serviceDesc, srv)
}

func _SyncService_ListModules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListModulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).ListModules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncService/ListModules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).ListModules(ctx, req.(*ListModulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Sync(ctx, req.(*SyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_Obtain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObtainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).Obtain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SyncService/Obtain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).Obtain(ctx, req.(*ObtainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListModules",
			Handler:    _SyncService_ListModules_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _SyncService_Sync_Handler,
		},
		{
			MethodName: "Obtain",
			Handler:    _SyncService_Obtain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_5fb666e15fab7b09) }

var fileDescriptor_api_5fb666e15fab7b09 = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x4a,
	0x14, 0x3d, 0x8e, 0x73, 0xdd, 0x39, 0x89, 0x7a, 0xa6, 0xd5, 0x39, 0x3e, 0x91, 0xa0, 0x95, 0xa1,
	0x55, 0x85, 0x1a, 0x47, 0x4a, 0x91, 0xa8, 0xb8, 0x3c, 0xb4, 0x5c, 0xa4, 0x4a, 0x94, 0x4a, 0xd3,
	0x37, 0x1e, 0xa8, 0x26, 0xce, 0x6e, 0xb0, 0x70, 0x66, 0x8c, 0x3d, 0x8e, 0xe4, 0x27, 0xfe, 0x85,
	0x7f, 0xe0, 0x1f, 0x78, 0xe2, 0x27, 0xf8, 0x11, 0x34, 0xb7, 0x5c, 0x28, 0x42, 0xe2, 0x6d, 0xf6,
	0x5a, 0x7b, 0xf6, 0x5e, 0xb3, 0x66, 0xcf, 0x40, 0x87, 0x65, 0x49, 0x94, 0xe5, 0x42, 0x0a, 0xe2,
	0xb3, 0x2c, 0x19, 0x0c, 0x67, 0x89, 0x7c, 0x5f, 0x4e, 0xa2, 0x58, 0xcc, 0x47, 0x33, 0x31, 0x13,
	0x23, 0xcd, 0x4d, 0xca, 0x1b, 0x1d, 0xe9, 0x40, 0xaf, 0xcc, 0x9e, 0xc1, 0xff, 0x33, 0x21, 0x66,
	0x29, 0xae, 0xb2, 0x18, 0xaf, 0x0c, 0x15, 0x9e, 0x41, 0xf3, 0x42, 0x4c, 0xcb, 0x14, 0x09, 0x81,
	0x3a, 0x67, 0x73, 0x0c, 0xbc, 0x3d, 0xef, 0xb0, 0x43, 0xf5, 0x9a, 0xdc, 0x87, 0x46, 0x91, 0x61,
	0x5c, 0x04, 0xb5, 0x3d, 0xff, 0xb0, 0x3b, 0xee, 0x47, 0x4a, 0xc7, 0x85, 0x98, 0x62, 0x7a, 0x95,
	0x61, 0x4c, 0x0d, 0x19, 0x7e, 0xf3, 0xa0, 0xb3, 0x04, 0x49, 0x00, 0xad, 0x05, 0xe6, 0x45, 0x22,
	0xb8, 0x2d, 0xe5, 0x42, 0xb2, 0x03, 0x8d, 0x38, 0x65, 0x85, 0xaa, 0xa6, 0x70, 0x13, 0x90, 0x7f,
	0xa1, 0x39, 0xd7, 0x0a, 0x02, 0x5f, 0xc3, 0x36, 0x52, 0x7a, 0x64, 0x95, 0x61, 0x50, 0x37, 0x7a,
	0xd4, 0x9a, 0x1c, 0x41, 0x7d, 0x8e, 0x92, 0x05, 0x0d, 0x2d, 0x27, 0xd8, 0x94, 0x13, 0x5d, 0xa0,
	0x64, 0x2f, 0xb9, 0xcc, 0x2b, 0xaa, 0xb3, 0x06, 0x8f, 0xa0, 0xb3, 0x84, 0xc8, 0x16, 0xf8, 0x1f,
	0xb0, 0xb2, 0x92, 0xd4, 0x52, 0xc9, 0x59, 0xb0, 0xb4, 0x44, 0x27, 0x47, 0x07, 0x8f, 0x6b, 0x27,
	0x5e, 0xf8, 0xd9, 0x83, 0x86, 0x2e, 0x4b, 0x0e, 0xc0, 0x67, 0xdc, 0xec, 0xea, 0x8e, 0x77, 0x22,
	0xe3, 0x63, 0xe4, 0x7c, 0x8c, 0x4e, 0x79, 0x45, 0x55, 0x02, 0x79, 0x08, 0x6d, 0xd5, 0x72, 0xca,
	0x24, 0xb3, 0x5e, 0xad, 0x89, 0xd3, 0xc2, 0x14, 0x65, 0xc4, 0x2d, 0x33, 0x07, 0x4f, 0xa0, 0xb7,
	0x41, 0xfd, 0x91, 0xc8, 0x1d, 0x20, 0xaf, 0x93, 0x42, 0x9a, 0xdb, 0x2b, 0x28, 0x7e, 0x2c, 0xb1,
	0x90, 0xe1, 0x53, 0xd8, 0xde, 0x40, 0x8b, 0x4c, 0xf0, 0x02, 0xc9, 0x3e, 0xb4, 0x8c, 0xad, 0x45,
	0xe0, 0x69, 0x79, 0x5d, 0x27, 0xaf, 0x4c, 0x91, 0x3a, 0x2e, 0x7c, 0x01, 0xed, 0xab, 0x8a, 0xc7,
	0xe7, 0x12, 0xe7, 0x64, 0x0f, 0x1a, 0x73, 0xa5, 0xde, 0x1e, 0x1e, 0x56, 0xe7, 0xa1, 0x86, 0x50,
	0x37, 0x37, 0xc5, 0x14, 0xa5, 0x11, 0xd7, 0xa6, 0x36, 0x0a, 0x5f, 0x41, 0x57, 0x55, 0xb9, 0xcc,
	0x64, 0x22, 0xb8, 0xbe, 0xe0, 0x1c, 0x8b, 0x8a, 0xc7, 0xba, 0x52, 0x9b, 0xda, 0x88, 0xec, 0x42,
	0x37, 0x47, 0x99, 0x57, 0xd7, 0xb1, 0x28, 0xb9, 0xd4, 0x35, 0x7a, 0x14, 0x34, 0xf4, 0x5c, 0x21,
	0xe1, 0x3b, 0x53, 0xc7, 0x1e, 0x8d, 0xdc, 0x83, 0x46, 0x22, 0x71, 0xee, 0x4e, 0xd0, 0xd3, 0x82,
	0x9c, 0x5c, 0x6a, 0x38, 0xf2, 0x00, 0x5a, 0xc2, 0xf4, 0xd5, 0x05, 0xbb, 0xe3, 0xad, 0x65, 0x9a,
	0xd5, 0x43, 0x5d, 0x42, 0xf8, 0x0c, 0xfe, 0x36, 0xf5, 0xad, 0x49, 0x43, 0x68, 0xe5, 0x58, 0x94,
	0xa9, 0x74, 0x2d, 0xb6, 0x37, 0x5b, 0x68, 0x8e, 0xba, 0x9c, 0xf0, 0x0d, 0xf4, 0x37, 0x29, 0xd2,
	0x87, 0x5a, 0x32, 0xb5, 0xb7, 0x57, 0x4b, 0xa6, 0xea, 0xe4, 0x85, 0x64, 0xb2, 0x74, 0x13, 0x6f,
	0x23, 0x75, 0xa9, 0x98, 0xe7, 0x22, 0xb7, 0x13, 0x6f, 0x82, 0x30, 0x82, 0xde, 0xe5, 0x44, 0xb2,
	0x84, 0x3b, 0xe3, 0xee, 0x00, 0x08, 0x0d, 0x5c, 0xb3, 0x34, 0xb5, 0xe6, 0x75, 0x0c, 0x72, 0x9a,
	0xa6, 0x21, 0xba, 0x7c, 0x67, 0xd0, 0x16, 0xf8, 0xc9, 0xd4, 0x68, 0xef, 0x50, 0xb5, 0x54, 0x8d,
	0x56, 0xef, 0xb7, 0x63, 0xdf, 0x2b, 0x39, 0x5a, 0x79, 0xe4, 0x6b, 0x8f, 0x88, 0x3e, 0xe7, 0x46,
	0xf3, 0x95, 0x4b, 0x27, 0xd0, 0x77, 0x6d, 0xac, 0x4f, 0x07, 0xfa, 0xc5, 0x62, 0xea, 0x6c, 0x5a,
	0xfb, 0x16, 0xce, 0xf9, 0x8d, 0xa0, 0x96, 0x0d, 0x3f, 0xd9, 0x6f, 0x41, 0x81, 0x24, 0x84, 0xba,
	0xea, 0x6e, 0xa7, 0xe9, 0xe7, 0x9f, 0x44, 0x73, 0xab, 0x91, 0xab, 0xfd, 0x66, 0xe4, 0xac, 0xa3,
	0xfe, 0xaf, 0x1d, 0xad, 0xaf, 0x39, 0x3a, 0xfe, 0xe2, 0x99, 0x09, 0xba, 0xc2, 0x7c, 0x91, 0xc4,
	0x48, 0xce, 0xa0, 0xbb, 0xf6, 0x38, 0xc8, 0x7f, 0xba, 0xfe, 0xed, 0x47, 0x34, 0x08, 0x6e, 0x13,
	0xe6, 0xe8, 0xe1, 0x5f, 0x64, 0x08, 0x75, 0x55, 0x92, 0xac, 0xe6, 0xca, 0xed, 0xfa, 0x67, 0x0d,
	0x59, 0xa6, 0x1f, 0x43, 0xd3, 0xb8, 0x47, 0xd6, 0x4d, 0x76, 0x5b, 0xb6, 0x37, 0x30, 0xb7, 0xe9,
	0x6c, 0xff, 0xeb, 0xf7, 0xbb, 0xde, 0xdb, 0xdd, 0xb5, 0x4f, 0x3e, 0x4d, 0x66, 0x4c, 0x8a, 0xd1,
	0x22, 0xcb, 0x86, 0x6c, 0x86, 0x5c, 0x8e, 0x58, 0x96, 0x4c, 0x9a, 0xfa, 0x1f, 0x3a, 0xfe, 0x11,
	0x00, 0x00, 0xff, 0xff, 0x82, 0x74, 0xcf, 0xa2, 0x1e, 0x06, 0x00, 0x00,
}
