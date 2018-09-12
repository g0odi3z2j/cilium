// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/cilium_l7policy.proto

package cilium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type L7Policy struct {
	// Path to the unix domain socket for the cilium access log.
	AccessLogPath string `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath,proto3" json:"access_log_path,omitempty"`
	// Cilium endpoint security policy to enforce.
	PolicyName string `protobuf:"bytes,2,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// HTTP response body message for 403 status code.
	// If empty, "Access denied" will be used.
	Denied_403Body string `protobuf:"bytes,3,opt,name=denied_403_body,json=denied403Body,proto3" json:"denied_403_body,omitempty"`
	// 'true' if the filter is on ingress listener, 'false' for egress listener.
	// Value from the listener filter will be used if not specified here.
	IsIngress            *wrappers.BoolValue `protobuf:"bytes,4,opt,name=is_ingress,json=isIngress,proto3" json:"is_ingress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *L7Policy) Reset()         { *m = L7Policy{} }
func (m *L7Policy) String() string { return proto.CompactTextString(m) }
func (*L7Policy) ProtoMessage()    {}
func (*L7Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b4ffb6fa1d76306, []int{0}
}
func (m *L7Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L7Policy.Unmarshal(m, b)
}
func (m *L7Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L7Policy.Marshal(b, m, deterministic)
}
func (dst *L7Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L7Policy.Merge(dst, src)
}
func (m *L7Policy) XXX_Size() int {
	return xxx_messageInfo_L7Policy.Size(m)
}
func (m *L7Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_L7Policy.DiscardUnknown(m)
}

var xxx_messageInfo_L7Policy proto.InternalMessageInfo

func (m *L7Policy) GetAccessLogPath() string {
	if m != nil {
		return m.AccessLogPath
	}
	return ""
}

func (m *L7Policy) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *L7Policy) GetDenied_403Body() string {
	if m != nil {
		return m.Denied_403Body
	}
	return ""
}

func (m *L7Policy) GetIsIngress() *wrappers.BoolValue {
	if m != nil {
		return m.IsIngress
	}
	return nil
}

func init() {
	proto.RegisterType((*L7Policy)(nil), "cilium.L7Policy")
}

func init() { proto.RegisterFile("cilium/cilium_l7policy.proto", fileDescriptor_8b4ffb6fa1d76306) }

var fileDescriptor_8b4ffb6fa1d76306 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x95, 0x62, 0x53, 0x44, 0xc8, 0x29, 0x14, 0xd1, 0xe2, 0x41, 0x7a, 0xca, 0x16,
	0xb7, 0x50, 0xbc, 0xf6, 0x26, 0x14, 0x29, 0x3d, 0x78, 0x0d, 0xd9, 0x64, 0xcc, 0x06, 0xb2, 0x3b,
	0x21, 0xd9, 0x45, 0xf6, 0xc5, 0x7c, 0x3e, 0x31, 0x83, 0xa7, 0x81, 0x6f, 0xbe, 0xf9, 0xe7, 0x67,
	0x0f, 0xc6, 0x07, 0x3f, 0xf5, 0x35, 0x0d, 0x15, 0x0e, 0x11, 0x83, 0x37, 0xb3, 0x8c, 0x09, 0x47,
	0xe4, 0x0b, 0xc2, 0xeb, 0x47, 0x87, 0xe8, 0x02, 0xd4, 0x85, 0xb6, 0xd3, 0x57, 0xfd, 0x9d, 0x74,
	0x8c, 0x90, 0x32, 0x79, 0xcf, 0x3f, 0x15, 0xbb, 0x3d, 0x1d, 0xce, 0xe5, 0x94, 0xbf, 0xb0, 0x7b,
	0x6d, 0x0c, 0xe4, 0xac, 0x02, 0x3a, 0x15, 0xf5, 0xd8, 0x89, 0x6a, 0x53, 0x6d, 0x97, 0x97, 0x3b,
	0xc2, 0x27, 0x74, 0x67, 0x3d, 0x76, 0xfc, 0x89, 0xad, 0xe8, 0x99, 0x1a, 0x74, 0x0f, 0xe2, 0xaa,
	0x38, 0x8c, 0xd0, 0x87, 0xee, 0xe1, 0x2f, 0xc8, 0xc2, 0xe0, 0xc1, 0xaa, 0xfd, 0xae, 0x51, 0x2d,
	0xda, 0x59, 0x5c, 0x53, 0x10, 0xe1, 0xfd, 0xae, 0x39, 0xa2, 0x9d, 0xf9, 0x1b, 0x63, 0x3e, 0x2b,
	0x3f, 0xb8, 0x04, 0x39, 0x8b, 0x9b, 0x4d, 0xb5, 0x5d, 0xbd, 0xae, 0x25, 0x55, 0x96, 0xff, 0x95,
	0xe5, 0x11, 0x31, 0x7c, 0xea, 0x30, 0xc1, 0x65, 0xe9, 0xf3, 0x3b, 0xc9, 0xed, 0xa2, 0xac, 0x9b,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xee, 0x32, 0x36, 0x07, 0x01, 0x00, 0x00,
}
