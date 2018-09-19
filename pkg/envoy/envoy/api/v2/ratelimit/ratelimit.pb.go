// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/ratelimit/ratelimit.proto

package ratelimit

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A RateLimitDescriptor is a list of hierarchical entries that are used by the service to
// determine the final rate limit key and overall allowed limit. Here are some examples of how
// they might be used for the domain "envoy".
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits all unauthenticated traffic for the IP address 10.0.0.1. The
// configuration supplies a default limit for the *remote_address* key. If there is a desire to
// raise the limit for 10.0.0.1 or block it entirely it can be specified directly in the
// configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"]
//
// What it does: Limits all unauthenticated traffic globally for a specific path (or prefix if
// configured that way in the service).
//
// .. code-block:: cpp
//
//   ["authenticated": "false"], ["path": "/foo/bar"], ["remote_address": "10.0.0.1"]
//
// What it does: Limits unauthenticated traffic to a specific path for a specific IP address.
// Like (1) we can raise/block specific IP addresses if we want with an override configuration.
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"]
//
// What it does: Limits all traffic for an authenticated client "foo"
//
// .. code-block:: cpp
//
//   ["authenticated": "true"], ["client_id": "foo"], ["path": "/foo/bar"]
//
// What it does: Limits traffic to a specific path for an authenticated client "foo"
//
// The idea behind the API is that (1)/(2)/(3) and (4)/(5) can be sent in 1 request if desired.
// This enables building complex application scenarios with a generic backend.
type RateLimitDescriptor struct {
	// Descriptor entries.
	Entries              []*RateLimitDescriptor_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitDescriptor) Reset()         { *m = RateLimitDescriptor{} }
func (m *RateLimitDescriptor) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor) ProtoMessage()    {}
func (*RateLimitDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_5684844e04543b8d, []int{0}
}

func (m *RateLimitDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor.Unmarshal(m, b)
}
func (m *RateLimitDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor.Merge(m, src)
}
func (m *RateLimitDescriptor) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor.Size(m)
}
func (m *RateLimitDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor proto.InternalMessageInfo

func (m *RateLimitDescriptor) GetEntries() []*RateLimitDescriptor_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type RateLimitDescriptor_Entry struct {
	// Descriptor key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Descriptor value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimitDescriptor_Entry) Reset()         { *m = RateLimitDescriptor_Entry{} }
func (m *RateLimitDescriptor_Entry) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptor_Entry) ProtoMessage()    {}
func (*RateLimitDescriptor_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_5684844e04543b8d, []int{0, 0}
}

func (m *RateLimitDescriptor_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Unmarshal(m, b)
}
func (m *RateLimitDescriptor_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptor_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptor_Entry.Merge(m, src)
}
func (m *RateLimitDescriptor_Entry) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptor_Entry.Size(m)
}
func (m *RateLimitDescriptor_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptor_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptor_Entry proto.InternalMessageInfo

func (m *RateLimitDescriptor_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptor_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*RateLimitDescriptor)(nil), "envoy.api.v2.ratelimit.RateLimitDescriptor")
	proto.RegisterType((*RateLimitDescriptor_Entry)(nil), "envoy.api.v2.ratelimit.RateLimitDescriptor.Entry")
}

func init() {
	proto.RegisterFile("envoy/api/v2/ratelimit/ratelimit.proto", fileDescriptor_5684844e04543b8d)
}

var fileDescriptor_5684844e04543b8d = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x33, 0xd2, 0x2f, 0x4a, 0x2c, 0x49, 0xcd, 0xc9, 0xcc,
	0xcd, 0x2c, 0x41, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xc0, 0xea, 0xf4, 0x12,
	0x0b, 0x32, 0xf5, 0xca, 0x8c, 0xf4, 0xe0, 0xb2, 0x52, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89,
	0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x83, 0xd2, 0x56, 0x46, 0x2e, 0xe1, 0xa0, 0xc4, 0x92, 0x54,
	0x1f, 0x90, 0x32, 0x97, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0x82, 0x92, 0xfc, 0x22, 0xa1, 0x70, 0x2e,
	0xf6, 0xd4, 0xbc, 0x92, 0xa2, 0xcc, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x43,
	0x3d, 0xec, 0x46, 0xeb, 0x61, 0xd1, 0xad, 0xe7, 0x9a, 0x57, 0x52, 0x54, 0xe9, 0xc4, 0xb5, 0xeb,
	0xe5, 0x01, 0x66, 0xd6, 0x49, 0x8c, 0x4c, 0x1c, 0x8c, 0x41, 0x30, 0xd3, 0xa4, 0x5c, 0xb9, 0x58,
	0xc1, 0xb2, 0x42, 0xd2, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e,
	0x9c, 0x20, 0xa5, 0x2c, 0x45, 0x4c, 0x0a, 0x8c, 0x41, 0x20, 0x51, 0x21, 0x79, 0x2e, 0xd6, 0xb2,
	0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x74, 0x69, 0x88, 0xb8, 0x13, 0x77, 0x14, 0x27, 0xdc, 0x09,
	0x49, 0x6c, 0x60, 0xbf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x86, 0x2d, 0x05, 0x19, 0x26,
	0x01, 0x00, 0x00,
}
