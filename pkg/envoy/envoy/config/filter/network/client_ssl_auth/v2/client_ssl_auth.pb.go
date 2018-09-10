// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/client_ssl_auth/v2/client_ssl_auth.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/filter/network/client_ssl_auth/v2/client_ssl_auth.proto

It has these top-level messages:
	ClientSSLAuth
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core1 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import google_protobuf3 "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientSSLAuth struct {
	// The :ref:`cluster manager <arch_overview_cluster_manager>` cluster that runs
	// the authentication service. The filter will connect to the service every 60s to fetch the list
	// of principals. The service must support the expected :ref:`REST API
	// <config_network_filters_client_ssl_auth_rest_api>`.
	AuthApiCluster string `protobuf:"bytes,1,opt,name=auth_api_cluster,json=authApiCluster" json:"auth_api_cluster,omitempty"`
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_client_ssl_auth_stats>`.
	StatPrefix string `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix" json:"stat_prefix,omitempty"`
	// Time in milliseconds between principal refreshes from the
	// authentication service. Default is 60000 (60s). The actual fetch time
	// will be this value plus a random jittered value between
	// 0-refresh_delay_ms milliseconds.
	RefreshDelay *google_protobuf3.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay" json:"refresh_delay,omitempty"`
	// An optional list of IP address and subnet masks that should be white
	// listed for access by the filter. If no list is provided, there is no
	// IP white list.
	IpWhiteList []*envoy_api_v2_core1.CidrRange `protobuf:"bytes,4,rep,name=ip_white_list,json=ipWhiteList" json:"ip_white_list,omitempty"`
}

func (m *ClientSSLAuth) Reset()                    { *m = ClientSSLAuth{} }
func (m *ClientSSLAuth) String() string            { return proto.CompactTextString(m) }
func (*ClientSSLAuth) ProtoMessage()               {}
func (*ClientSSLAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientSSLAuth) GetAuthApiCluster() string {
	if m != nil {
		return m.AuthApiCluster
	}
	return ""
}

func (m *ClientSSLAuth) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ClientSSLAuth) GetRefreshDelay() *google_protobuf3.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ClientSSLAuth) GetIpWhiteList() []*envoy_api_v2_core1.CidrRange {
	if m != nil {
		return m.IpWhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientSSLAuth)(nil), "envoy.config.filter.network.client_ssl_auth.v2.ClientSSLAuth")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/client_ssl_auth/v2/client_ssl_auth.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x6a, 0xeb, 0x30,
	0x10, 0xc6, 0x71, 0x62, 0x1e, 0xc4, 0x7e, 0x79, 0x3c, 0x4c, 0xa1, 0x6e, 0x28, 0x8d, 0xe9, 0x2a,
	0x74, 0x21, 0x81, 0x73, 0x81, 0x26, 0xf1, 0x32, 0x8b, 0xe2, 0x2c, 0x0a, 0xdd, 0x08, 0x25, 0x1e,
	0xdb, 0x43, 0x85, 0x25, 0x24, 0xd9, 0x69, 0x6e, 0xd2, 0xb3, 0x74, 0xd5, 0x9b, 0x74, 0xdd, 0x1b,
	0x74, 0x59, 0xfc, 0x27, 0x9b, 0xec, 0x46, 0xfa, 0xe6, 0xfb, 0xcd, 0x7c, 0xe3, 0x25, 0x50, 0x35,
	0xf2, 0x44, 0x0f, 0xb2, 0xca, 0xb1, 0xa0, 0x39, 0x0a, 0x0b, 0x9a, 0x56, 0x60, 0x8f, 0x52, 0xbf,
	0xd2, 0x83, 0x40, 0xa8, 0x2c, 0x33, 0x46, 0x30, 0x5e, 0xdb, 0x92, 0x36, 0xf1, 0xe5, 0x17, 0x51,
	0x5a, 0x5a, 0x19, 0x90, 0x8e, 0x42, 0x7a, 0x0a, 0xe9, 0x29, 0x64, 0xa0, 0x90, 0x4b, 0x4b, 0x13,
	0xcf, 0xe6, 0xfd, 0x54, 0xae, 0xb0, 0x63, 0x4a, 0x0d, 0x94, 0x67, 0x99, 0x06, 0x63, 0x7a, 0xe0,
	0xec, 0xae, 0x90, 0xb2, 0x10, 0x40, 0xbb, 0xd7, 0xbe, 0xce, 0x69, 0x56, 0x6b, 0x6e, 0x51, 0x56,
	0x83, 0x7e, 0xdd, 0x70, 0x81, 0x19, 0xb7, 0x40, 0xcf, 0xc5, 0x20, 0x5c, 0x15, 0xb2, 0x90, 0x5d,
	0x49, 0xdb, 0xaa, 0xff, 0xbd, 0xff, 0x71, 0xbc, 0xe9, 0xa6, 0x5b, 0x63, 0xb7, 0xdb, 0xae, 0x6a,
	0x5b, 0x06, 0x4b, 0xef, 0x7f, 0xbb, 0x0c, 0xe3, 0x0a, 0xd9, 0x41, 0xd4, 0xc6, 0x82, 0x0e, 0x9d,
	0xc8, 0x59, 0x4c, 0xd6, 0x93, 0x8f, 0xef, 0xcf, 0xb1, 0xab, 0x47, 0x91, 0x93, 0xfe, 0x6b, 0x5b,
	0x56, 0x0a, 0x37, 0x7d, 0x43, 0xf0, 0xe0, 0xf9, 0xc6, 0x72, 0xcb, 0x94, 0x86, 0x1c, 0xdf, 0xc2,
	0xd1, 0x65, 0xbf, 0xd7, 0xaa, 0x4f, 0x9d, 0x18, 0x24, 0xde, 0x54, 0x43, 0xae, 0xc1, 0x94, 0x2c,
	0x03, 0xc1, 0x4f, 0xe1, 0x38, 0x72, 0x16, 0x7e, 0x7c, 0x43, 0xfa, 0x64, 0xe4, 0x9c, 0x8c, 0x24,
	0x43, 0xb2, 0xb5, 0xfb, 0xfe, 0x35, 0x77, 0xd2, 0xbf, 0x83, 0x2b, 0x69, 0x4d, 0xc1, 0xa3, 0x37,
	0x45, 0xc5, 0x8e, 0x25, 0x5a, 0x60, 0x02, 0x8d, 0x0d, 0xdd, 0x68, 0xbc, 0xf0, 0xe3, 0xdb, 0xe1,
	0xe0, 0x5c, 0x21, 0x69, 0x62, 0xd2, 0x1e, 0x90, 0x6c, 0x30, 0xd3, 0x29, 0xaf, 0x0a, 0x48, 0x7d,
	0x54, 0xcf, 0xad, 0x63, 0x8b, 0xc6, 0xae, 0xdd, 0x97, 0x51, 0x13, 0xef, 0xff, 0x74, 0xe3, 0x96,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x31, 0xfe, 0x47, 0xef, 0x01, 0x00, 0x00,
}
