// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/cilium_l7policy.proto

package cilium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core2 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type L7Policy struct {
	// Path to the unix domain socket for the cilium access log.
	AccessLogPath string `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath" json:"access_log_path,omitempty"`
	// Cilium identifier for the listener.
	ListenerId string `protobuf:"bytes,2,opt,name=listener_id,json=listenerId" json:"listener_id,omitempty"`
	// Cilium endpoint security policy to enforce
	PolicyName string `protobuf:"bytes,3,opt,name=policy_name,json=policyName" json:"policy_name,omitempty"`
	// gRPC API config source for network policy
	ApiConfigSource *envoy_api_v2_core2.ApiConfigSource `protobuf:"bytes,4,opt,name=api_config_source,json=apiConfigSource" json:"api_config_source,omitempty"`
}

func (m *L7Policy) Reset()                    { *m = L7Policy{} }
func (m *L7Policy) String() string            { return proto.CompactTextString(m) }
func (*L7Policy) ProtoMessage()               {}
func (*L7Policy) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *L7Policy) GetAccessLogPath() string {
	if m != nil {
		return m.AccessLogPath
	}
	return ""
}

func (m *L7Policy) GetListenerId() string {
	if m != nil {
		return m.ListenerId
	}
	return ""
}

func (m *L7Policy) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *L7Policy) GetApiConfigSource() *envoy_api_v2_core2.ApiConfigSource {
	if m != nil {
		return m.ApiConfigSource
	}
	return nil
}

func init() {
	proto.RegisterType((*L7Policy)(nil), "cilium.L7Policy")
}

func init() { proto.RegisterFile("cilium/cilium_l7policy.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x46, 0x59, 0x95, 0x52, 0x53, 0xa4, 0xb8, 0xa7, 0x45, 0x04, 0x4b, 0x41, 0xe9, 0x29, 0x81,
	0xf5, 0xd0, 0xb3, 0x7a, 0x12, 0x4a, 0x29, 0xf5, 0xe6, 0x25, 0x8c, 0xe9, 0xb8, 0x1d, 0xc8, 0x66,
	0x42, 0x92, 0x2e, 0xf4, 0xff, 0xf9, 0xc3, 0xa4, 0x89, 0x1e, 0x7a, 0x1a, 0x78, 0x3c, 0x86, 0xf7,
	0x89, 0x7b, 0x43, 0x96, 0x0e, 0xbd, 0x2a, 0x47, 0xdb, 0xa5, 0x67, 0x4b, 0xe6, 0x28, 0x7d, 0xe0,
	0xc4, 0xf5, 0xa8, 0xe0, 0xbb, 0x47, 0x74, 0x03, 0x1f, 0x15, 0x78, 0x52, 0x43, 0xab, 0x0c, 0x07,
	0x54, 0x86, 0xdd, 0x37, 0x75, 0x3a, 0xf2, 0x21, 0x18, 0x2c, 0xfa, 0xfc, 0xa7, 0x12, 0xe3, 0xd5,
	0x72, 0x93, 0x3f, 0xd4, 0x4f, 0x62, 0x0a, 0xc6, 0x60, 0x8c, 0xda, 0x72, 0xa7, 0x3d, 0xa4, 0x7d,
	0x53, 0xcd, 0xaa, 0xc5, 0xf5, 0xf6, 0xa6, 0xe0, 0x15, 0x77, 0x1b, 0x48, 0xfb, 0xfa, 0x41, 0x4c,
	0x2c, 0xc5, 0x84, 0x0e, 0x83, 0xa6, 0x5d, 0x73, 0x91, 0x1d, 0xf1, 0x8f, 0xde, 0x77, 0x27, 0xa1,
	0x44, 0x69, 0x07, 0x3d, 0x36, 0x97, 0x45, 0x28, 0x68, 0x0d, 0x3d, 0xd6, 0x6b, 0x71, 0x0b, 0x9e,
	0xf4, 0x59, 0x51, 0x73, 0x35, 0xab, 0x16, 0x93, 0x76, 0x2e, 0x73, 0xb9, 0x04, 0x4f, 0x72, 0x68,
	0xe5, 0xa9, 0x5c, 0xbe, 0x78, 0x7a, 0xcb, 0xea, 0x47, 0x36, 0xb7, 0x53, 0x38, 0x07, 0xaf, 0xe3,
	0xcf, 0xbf, 0xdd, 0x5f, 0xa3, 0xbc, 0xeb, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xff, 0xa2, 0x3a,
	0xa7, 0x26, 0x01, 0x00, 0x00,
}
