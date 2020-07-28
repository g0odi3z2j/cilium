// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/api/network_filter.proto

package cilium

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NetworkFilter struct {
	// Path to the proxylib to be opened
	Proxylib string `protobuf:"bytes,1,opt,name=proxylib,proto3" json:"proxylib,omitempty"`
	// Transparent set of parameters provided for proxylib initialization
	ProxylibParams map[string]string `protobuf:"bytes,2,rep,name=proxylib_params,json=proxylibParams,proto3" json:"proxylib_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// L7 protocol identifier
	L7Proto string `protobuf:"bytes,3,opt,name=l7_proto,json=l7Proto,proto3" json:"l7_proto,omitempty"`
	// Name of the policy to be enforced
	PolicyName string `protobuf:"bytes,4,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// Path to the unix domain socket for the cilium access log.
	AccessLogPath        string   `protobuf:"bytes,5,opt,name=access_log_path,json=accessLogPath,proto3" json:"access_log_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkFilter) Reset()         { *m = NetworkFilter{} }
func (m *NetworkFilter) String() string { return proto.CompactTextString(m) }
func (*NetworkFilter) ProtoMessage()    {}
func (*NetworkFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd4806b6c243641, []int{0}
}

func (m *NetworkFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkFilter.Unmarshal(m, b)
}
func (m *NetworkFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkFilter.Marshal(b, m, deterministic)
}
func (m *NetworkFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkFilter.Merge(m, src)
}
func (m *NetworkFilter) XXX_Size() int {
	return xxx_messageInfo_NetworkFilter.Size(m)
}
func (m *NetworkFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkFilter.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkFilter proto.InternalMessageInfo

func (m *NetworkFilter) GetProxylib() string {
	if m != nil {
		return m.Proxylib
	}
	return ""
}

func (m *NetworkFilter) GetProxylibParams() map[string]string {
	if m != nil {
		return m.ProxylibParams
	}
	return nil
}

func (m *NetworkFilter) GetL7Proto() string {
	if m != nil {
		return m.L7Proto
	}
	return ""
}

func (m *NetworkFilter) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *NetworkFilter) GetAccessLogPath() string {
	if m != nil {
		return m.AccessLogPath
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkFilter)(nil), "cilium.NetworkFilter")
	proto.RegisterMapType((map[string]string)(nil), "cilium.NetworkFilter.ProxylibParamsEntry")
}

func init() { proto.RegisterFile("cilium/api/network_filter.proto", fileDescriptor_0bd4806b6c243641) }

var fileDescriptor_0bd4806b6c243641 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0x5d, 0x5b, 0xeb, 0x94, 0x5a, 0x89, 0x1e, 0x62, 0x2f, 0x2d, 0x1e, 0xa4, 0x5e,
	0xb6, 0xa0, 0x87, 0x8a, 0x37, 0x0f, 0x7a, 0x92, 0xb2, 0xec, 0x0b, 0x84, 0xe9, 0x12, 0xdb, 0xd0,
	0xec, 0x26, 0x64, 0x53, 0x35, 0x2f, 0xe4, 0x73, 0x8a, 0x93, 0x2a, 0x2c, 0xf4, 0x36, 0xf3, 0xcf,
	0xc7, 0x9f, 0x2f, 0x30, 0xad, 0x94, 0x56, 0xfb, 0x7a, 0x81, 0x56, 0x2d, 0x1a, 0xe9, 0x3f, 0x8d,
	0xdb, 0x89, 0x77, 0xa5, 0xbd, 0x74, 0xb9, 0x75, 0xc6, 0x1b, 0xd6, 0x8f, 0xc0, 0xcd, 0x77, 0x0a,
	0xa3, 0x55, 0x04, 0x5e, 0xe9, 0xce, 0x26, 0x30, 0xb0, 0xce, 0x7c, 0x05, 0xad, 0xd6, 0x3c, 0x99,
	0x25, 0xf3, 0xb3, 0xf2, 0x7f, 0x67, 0x25, 0x8c, 0xff, 0x66, 0x61, 0xd1, 0x61, 0xdd, 0xf2, 0x74,
	0x96, 0xcd, 0x87, 0xf7, 0x77, 0x79, 0xec, 0xcb, 0x3b, 0x5d, 0x79, 0x71, 0x80, 0x0b, 0x62, 0x5f,
	0x1a, 0xef, 0x42, 0x79, 0x6e, 0x3b, 0x21, 0xbb, 0x86, 0x81, 0x5e, 0x0a, 0xb2, 0xe2, 0x19, 0xbd,
	0x77, 0xaa, 0x97, 0x05, 0x49, 0x4e, 0x61, 0x68, 0x8d, 0x56, 0x55, 0x10, 0x0d, 0xd6, 0x92, 0x9f,
	0xd0, 0x15, 0x62, 0xb4, 0xc2, 0x5a, 0xb2, 0x5b, 0x18, 0x63, 0x55, 0xc9, 0xb6, 0x15, 0xda, 0x6c,
	0x84, 0x45, 0xbf, 0xe5, 0x3d, 0x82, 0x46, 0x31, 0x7e, 0x33, 0x9b, 0x02, 0xfd, 0x76, 0xf2, 0x0c,
	0x97, 0x47, 0x54, 0xd8, 0x05, 0x64, 0x3b, 0x19, 0x0e, 0xbf, 0xfc, 0x1d, 0xd9, 0x15, 0xf4, 0x3e,
	0x50, 0xef, 0x25, 0x4f, 0x29, 0x8b, 0xcb, 0x53, 0xfa, 0x98, 0xac, 0xfb, 0x64, 0xf8, 0xf0, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0x4d, 0x7d, 0x30, 0x5a, 0x01, 0x00, 0x00,
}
