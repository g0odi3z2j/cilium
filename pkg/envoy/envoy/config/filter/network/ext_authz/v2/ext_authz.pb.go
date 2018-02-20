// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/ext_authz/v2/ext_authz.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/filter/network/ext_authz/v2/ext_authz.proto

It has these top-level messages:
	ExtAuthz
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core1 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [#not-implemented-hide:]
// External Authorization filter calls out to an external service over the
// gRPC Authorization API defined by :ref:`external_auth <envoy_api_msg_auth.CheckRequest>`.
// A failed check will cause this filter to close the TCP connection.
type ExtAuthz struct {
	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix" json:"stat_prefix,omitempty"`
	// The external authorization gRPC service configuration.
	GrpcService *envoy_api_v2_core1.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService" json:"grpc_service,omitempty"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. If set to true then in case of failure to get a
	// response back from the authorization service allow the traffic.
	// Defaults to false.
	// If set to true and the response from the authorization service is NOT
	// Denied then the traffic will be permitted.
	FailureModeAllow bool `protobuf:"varint,3,opt,name=failure_mode_allow,json=failureModeAllow" json:"failure_mode_allow,omitempty"`
}

func (m *ExtAuthz) Reset()                    { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string            { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()               {}
func (*ExtAuthz) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExtAuthz) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ExtAuthz) GetGrpcService() *envoy_api_v2_core1.GrpcService {
	if m != nil {
		return m.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.network.ext_authz.v2.ExtAuthz")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/ext_authz/v2/ext_authz.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xc9, 0xee, 0x22, 0xbb, 0xa9, 0x07, 0xe9, 0xc5, 0xb2, 0x07, 0x29, 0xe2, 0xa1, 0x88,
	0x24, 0x50, 0x2f, 0x5e, 0xbb, 0x20, 0x9e, 0x04, 0xa9, 0x37, 0x2f, 0x25, 0xb6, 0xd3, 0x1a, 0x8c,
	0x4d, 0x98, 0xcd, 0x66, 0xab, 0x2f, 0xe4, 0x3b, 0x78, 0xf2, 0x75, 0x7c, 0x0b, 0x49, 0xb3, 0x52,
	0x6f, 0x99, 0x7c, 0xf3, 0xfd, 0xf3, 0xd3, 0x1b, 0xe8, 0x9d, 0x7e, 0xe7, 0xb5, 0xee, 0x5b, 0xd9,
	0xf1, 0x56, 0x2a, 0x0b, 0xc8, 0x7b, 0xb0, 0x7b, 0x8d, 0xaf, 0x1c, 0x06, 0x5b, 0x89, 0x9d, 0x7d,
	0xf9, 0xe0, 0x2e, 0x9f, 0x06, 0x66, 0x50, 0x5b, 0x1d, 0x67, 0xa3, 0xc9, 0x82, 0xc9, 0x82, 0xc9,
	0x0e, 0x26, 0x9b, 0x96, 0x5d, 0xbe, 0xbe, 0x08, 0x37, 0x84, 0x91, 0x3e, 0xa7, 0xd6, 0x08, 0xbc,
	0x43, 0x53, 0x57, 0x5b, 0x40, 0x27, 0x6b, 0x08, 0x79, 0xeb, 0x53, 0x27, 0x94, 0x6c, 0x84, 0x05,
	0xfe, 0xf7, 0x08, 0xe0, 0xfc, 0x93, 0xd0, 0xe5, 0xed, 0x60, 0x0b, 0x1f, 0x17, 0x5f, 0xd2, 0x68,
	0x6b, 0x85, 0xad, 0x0c, 0x42, 0x2b, 0x87, 0x84, 0xa4, 0x24, 0x5b, 0x6d, 0x56, 0x5f, 0x3f, 0xdf,
	0xf3, 0x05, 0xce, 0x52, 0x52, 0x52, 0x4f, 0x1f, 0x46, 0x18, 0x17, 0xf4, 0xf8, 0xff, 0x9d, 0x64,
	0x96, 0x92, 0x2c, 0xca, 0xcf, 0x58, 0x28, 0x2e, 0x8c, 0x64, 0x2e, 0x67, 0xbe, 0x0e, 0xbb, 0x43,
	0x53, 0x3f, 0x86, 0xad, 0x32, 0xea, 0xa6, 0x21, 0xbe, 0xa2, 0x71, 0x2b, 0xa4, 0xda, 0x21, 0x54,
	0x6f, 0xba, 0x81, 0x4a, 0x28, 0xa5, 0xf7, 0xc9, 0x3c, 0x25, 0xd9, 0xb2, 0x3c, 0x39, 0x90, 0x7b,
	0xdd, 0x40, 0xe1, 0xff, 0x37, 0x8b, 0xa7, 0x99, 0xcb, 0x9f, 0x8f, 0xc6, 0xda, 0xd7, 0xbf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x91, 0x6a, 0xdd, 0x47, 0x5b, 0x01, 0x00, 0x00,
}
