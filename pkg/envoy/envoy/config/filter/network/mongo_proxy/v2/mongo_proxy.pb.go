// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto

It has these top-level messages:
	MongoProxy
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_config_filter_fault_v2 "github.com/cilium/cilium/pkg/envoy/envoy/config/filter/fault/v2"
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

type MongoProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mongo_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix" json:"stat_prefix,omitempty"`
	// The optional path to use for writing Mongo access logs. If not access log
	// path is specified no access logs will be written. Note that access log is
	// also gated :ref:`runtime <config_network_filters_mongo_proxy_runtime>`.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog" json:"access_log,omitempty"`
	// Inject a fixed delay before proxying a Mongo operation. Delays are
	// applied to the following MongoDB operations: Query, Insert, GetMore,
	// and KillCursors. Once an active delay is in progress, all incoming
	// data up until the timer event fires will be a part of the delay.
	Delay *envoy_config_filter_fault_v2.FaultDelay `protobuf:"bytes,3,opt,name=delay" json:"delay,omitempty"`
}

func (m *MongoProxy) Reset()                    { *m = MongoProxy{} }
func (m *MongoProxy) String() string            { return proto.CompactTextString(m) }
func (*MongoProxy) ProtoMessage()               {}
func (*MongoProxy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MongoProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MongoProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *MongoProxy) GetDelay() *envoy_config_filter_fault_v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func init() {
	proto.RegisterType((*MongoProxy)(nil), "envoy.config.filter.network.mongo_proxy.v2.MongoProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/mongo_proxy/v2/mongo_proxy.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0xcf, 0xcd, 0xcf, 0x4b, 0xcf, 0x8f, 0x2f, 0x28,
	0xca, 0xaf, 0xa8, 0xd4, 0x2f, 0x33, 0x42, 0xe6, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x69,
	0x81, 0x75, 0xeb, 0x41, 0x74, 0xeb, 0x41, 0x74, 0xeb, 0x41, 0x75, 0xeb, 0x21, 0x2b, 0x2f, 0x33,
	0x92, 0xd2, 0xc0, 0x66, 0x53, 0x5a, 0x62, 0x69, 0x4e, 0x09, 0xc8, 0x6c, 0x30, 0x03, 0x62, 0xaa,
	0x94, 0x78, 0x59, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0xaa, 0x3e, 0x8c, 0x01, 0x91, 0x50, 0x9a,
	0xce, 0xc8, 0xc5, 0xe5, 0x0b, 0x32, 0x35, 0x00, 0x64, 0xa8, 0x90, 0x16, 0x17, 0x77, 0x71, 0x49,
	0x62, 0x49, 0x7c, 0x41, 0x51, 0x6a, 0x5a, 0x66, 0x85, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xa7, 0x13,
	0xe7, 0xae, 0x97, 0x07, 0x98, 0x59, 0x8a, 0x98, 0x14, 0x18, 0x83, 0xb8, 0x40, 0xb2, 0x01, 0x60,
	0x49, 0x21, 0x59, 0x2e, 0xae, 0xc4, 0xe4, 0xe4, 0xd4, 0xe2, 0xe2, 0xf8, 0x9c, 0xfc, 0x74, 0x09,
	0x26, 0x90, 0xd2, 0x20, 0x4e, 0x88, 0x88, 0x4f, 0x7e, 0xba, 0x90, 0x1d, 0x17, 0x6b, 0x4a, 0x6a,
	0x4e, 0x62, 0xa5, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x86, 0x1e, 0x36, 0x8f, 0x41, 0xdc,
	0x58, 0x66, 0xa4, 0xe7, 0x06, 0x62, 0xb8, 0x80, 0xd4, 0x07, 0x41, 0xb4, 0x39, 0xb1, 0x44, 0x31,
	0x95, 0x19, 0x25, 0xb1, 0x81, 0x9d, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x0d, 0x43,
	0x42, 0x55, 0x01, 0x00, 0x00,
}
