// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ext_authz/v2/ext_authz.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/filter/http/ext_authz/v2/ext_authz.proto

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
// gRPC Authorization API defined by
// :ref:`external_auth <envoy_api_msg_auth.CheckRequest>`.
// A failed check will cause this filter to return 403 Forbidden.
type ExtAuthz struct {
	// The external authorization gRPC service configuration.
	GrpcService *envoy_api_v2_core1.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService" json:"grpc_service,omitempty"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. If set to true then in case of failure to get a
	// response back from the authorization service or getting a response that
	// is NOT denied then traffic will be permitted.
	// Defaults to false.
	FailureModeAllow bool `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow" json:"failure_mode_allow,omitempty"`
}

func (m *ExtAuthz) Reset()                    { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string            { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()               {}
func (*ExtAuthz) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.http.ext_authz.v2.ExtAuthz")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ext_authz/v2/ext_authz.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0x22, 0x52, 0x52, 0x0f, 0x92, 0x8b, 0xa5, 0x07, 0x29, 0xa2, 0xd0, 0x83, 0x4c,
	0x20, 0xe2, 0x03, 0xac, 0x20, 0x9e, 0xbc, 0xd4, 0x9b, 0x97, 0x10, 0xb3, 0xb3, 0xdb, 0x40, 0xec,
	0x84, 0x38, 0x8d, 0x55, 0x5f, 0x5e, 0xd2, 0x28, 0xdb, 0x5b, 0x26, 0xf3, 0x7f, 0xdf, 0xfc, 0xe2,
	0x1e, 0xb7, 0x99, 0xbe, 0x94, 0xa3, 0x6d, 0xef, 0x07, 0xd5, 0xfb, 0xc0, 0x98, 0xd4, 0x86, 0x39,
	0x2a, 0xdc, 0xb3, 0xb1, 0x3b, 0xde, 0x7c, 0xab, 0xac, 0xc7, 0x01, 0x62, 0x22, 0x26, 0x79, 0x73,
	0xc0, 0xa0, 0x62, 0x50, 0x31, 0x28, 0x18, 0x8c, 0xc9, 0xac, 0x17, 0xd7, 0xd5, 0x6e, 0xa3, 0x2f,
	0x12, 0x47, 0x09, 0xd5, 0x90, 0xa2, 0x33, 0x1f, 0x98, 0xb2, 0x77, 0x58, 0x65, 0x8b, 0x8b, 0x6c,
	0x83, 0xef, 0x2c, 0xa3, 0xfa, 0x7f, 0xd4, 0xc5, 0xd5, 0x8f, 0x98, 0x3e, 0xee, 0xb9, 0x2d, 0x36,
	0xd9, 0x8a, 0xb3, 0x63, 0x74, 0xde, 0x2c, 0x9b, 0xd5, 0x4c, 0x5f, 0x42, 0x2d, 0x62, 0xa3, 0x87,
	0xac, 0xa1, 0x5c, 0x80, 0xa7, 0x14, 0xdd, 0x4b, 0x4d, 0xad, 0x67, 0xc3, 0x38, 0xc8, 0x5b, 0x21,
	0x7b, 0xeb, 0xc3, 0x2e, 0xa1, 0x79, 0xa7, 0x0e, 0x8d, 0x0d, 0x81, 0x3e, 0xe7, 0x93, 0x65, 0xb3,
	0x9a, 0xae, 0xcf, 0xff, 0x36, 0xcf, 0xd4, 0x61, 0x5b, 0xfe, 0x1f, 0x4e, 0x5e, 0x27, 0x59, 0xbf,
	0x9d, 0x1e, 0x9a, 0xdc, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x48, 0xf0, 0x98, 0xd6, 0x28, 0x01,
	0x00, 0x00,
}
