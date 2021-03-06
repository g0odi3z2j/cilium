// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/extensions/upstreams/http/v3/http_protocol_options.proto

package envoy_extensions_upstreams_http_v3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// HttpProtocolOptions specifies Http upstream protocol options. This object
// is used in
// :ref:`typed_extension_protocol_options<envoy_api_field_config.cluster.v3.Cluster.typed_extension_protocol_options>`,
// keyed by the name `envoy.extensions.upstreams.http.v3.HttpProtocolOptions`.
//
// This controls what protocol(s) should be used for upstream and how said protocol(s) are configured.
//
// This replaces the prior pattern of explicit protocol configuration directly
// in the cluster. So a configuration like this, explicitly configuring the use of HTTP/2 upstream:
//
// .. code::
//
//   clusters:
//     - name: some_service
//       connect_timeout: 5s
//       upstream_http_protocol_options:
//         auto_sni: true
//       common_http_protocol_options:
//         idle_timeout: 1s
//       http2_protocol_options:
//         max_concurrent_streams: 100
//        .... [further cluster config]
//
// Would now look like this:
//
// .. code::
//
//   clusters:
//     - name: some_service
//       connect_timeout: 5s
//       typed_extension_protocol_options:
//         envoy.extensions.upstreams.http.v3.HttpProtocolOptions:
//           "@type": type.googleapis.com/envoy.extensions.upstreams.http.v3.HttpProtocolOptions
//           upstream_http_protocol_options:
//             auto_sni: true
//           common_http_protocol_options:
//             idle_timeout: 1s
//           explicit_http_config:
//             http2_protocol_options:
//               max_concurrent_streams: 100
//        .... [further cluster config]
// [#next-free-field: 6]
type HttpProtocolOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This contains options common across HTTP/1 and HTTP/2
	CommonHttpProtocolOptions *v3.HttpProtocolOptions `protobuf:"bytes,1,opt,name=common_http_protocol_options,json=commonHttpProtocolOptions,proto3" json:"common_http_protocol_options,omitempty"`
	// This contains common protocol options which are only applied upstream.
	UpstreamHttpProtocolOptions *v3.UpstreamHttpProtocolOptions `protobuf:"bytes,2,opt,name=upstream_http_protocol_options,json=upstreamHttpProtocolOptions,proto3" json:"upstream_http_protocol_options,omitempty"`
	// This controls the actual protocol to be used upstream.
	//
	// Types that are assignable to UpstreamProtocolOptions:
	//	*HttpProtocolOptions_ExplicitHttpConfig_
	//	*HttpProtocolOptions_UseDownstreamProtocolConfig
	//	*HttpProtocolOptions_AutoConfig
	UpstreamProtocolOptions isHttpProtocolOptions_UpstreamProtocolOptions `protobuf_oneof:"upstream_protocol_options"`
}

func (x *HttpProtocolOptions) Reset() {
	*x = HttpProtocolOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpProtocolOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpProtocolOptions) ProtoMessage() {}

func (x *HttpProtocolOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpProtocolOptions.ProtoReflect.Descriptor instead.
func (*HttpProtocolOptions) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescGZIP(), []int{0}
}

func (x *HttpProtocolOptions) GetCommonHttpProtocolOptions() *v3.HttpProtocolOptions {
	if x != nil {
		return x.CommonHttpProtocolOptions
	}
	return nil
}

func (x *HttpProtocolOptions) GetUpstreamHttpProtocolOptions() *v3.UpstreamHttpProtocolOptions {
	if x != nil {
		return x.UpstreamHttpProtocolOptions
	}
	return nil
}

func (m *HttpProtocolOptions) GetUpstreamProtocolOptions() isHttpProtocolOptions_UpstreamProtocolOptions {
	if m != nil {
		return m.UpstreamProtocolOptions
	}
	return nil
}

func (x *HttpProtocolOptions) GetExplicitHttpConfig() *HttpProtocolOptions_ExplicitHttpConfig {
	if x, ok := x.GetUpstreamProtocolOptions().(*HttpProtocolOptions_ExplicitHttpConfig_); ok {
		return x.ExplicitHttpConfig
	}
	return nil
}

func (x *HttpProtocolOptions) GetUseDownstreamProtocolConfig() *HttpProtocolOptions_UseDownstreamHttpConfig {
	if x, ok := x.GetUpstreamProtocolOptions().(*HttpProtocolOptions_UseDownstreamProtocolConfig); ok {
		return x.UseDownstreamProtocolConfig
	}
	return nil
}

func (x *HttpProtocolOptions) GetAutoConfig() *HttpProtocolOptions_AutoHttpConfig {
	if x, ok := x.GetUpstreamProtocolOptions().(*HttpProtocolOptions_AutoConfig); ok {
		return x.AutoConfig
	}
	return nil
}

type isHttpProtocolOptions_UpstreamProtocolOptions interface {
	isHttpProtocolOptions_UpstreamProtocolOptions()
}

type HttpProtocolOptions_ExplicitHttpConfig_ struct {
	// To explicitly configure either HTTP/1 or HTTP/2 (but not both!) use *explicit_http_config*.
	// If the *explicit_http_config* is empty, HTTP/1.1 is used.
	ExplicitHttpConfig *HttpProtocolOptions_ExplicitHttpConfig `protobuf:"bytes,3,opt,name=explicit_http_config,json=explicitHttpConfig,proto3,oneof"`
}

type HttpProtocolOptions_UseDownstreamProtocolConfig struct {
	// This allows switching on protocol based on what protocol the downstream
	// connection used.
	UseDownstreamProtocolConfig *HttpProtocolOptions_UseDownstreamHttpConfig `protobuf:"bytes,4,opt,name=use_downstream_protocol_config,json=useDownstreamProtocolConfig,proto3,oneof"`
}

type HttpProtocolOptions_AutoConfig struct {
	// This allows switching on protocol based on ALPN
	AutoConfig *HttpProtocolOptions_AutoHttpConfig `protobuf:"bytes,5,opt,name=auto_config,json=autoConfig,proto3,oneof"`
}

func (*HttpProtocolOptions_ExplicitHttpConfig_) isHttpProtocolOptions_UpstreamProtocolOptions() {}

func (*HttpProtocolOptions_UseDownstreamProtocolConfig) isHttpProtocolOptions_UpstreamProtocolOptions() {
}

func (*HttpProtocolOptions_AutoConfig) isHttpProtocolOptions_UpstreamProtocolOptions() {}

// If this is used, the cluster will only operate on one of the possible upstream protocols (HTTP/1.1, HTTP/2).
// Note that HTTP/2 should generally be used for upstream clusters doing gRPC.
type HttpProtocolOptions_ExplicitHttpConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ProtocolConfig:
	//	*HttpProtocolOptions_ExplicitHttpConfig_HttpProtocolOptions
	//	*HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions
	ProtocolConfig isHttpProtocolOptions_ExplicitHttpConfig_ProtocolConfig `protobuf_oneof:"protocol_config"`
}

func (x *HttpProtocolOptions_ExplicitHttpConfig) Reset() {
	*x = HttpProtocolOptions_ExplicitHttpConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpProtocolOptions_ExplicitHttpConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpProtocolOptions_ExplicitHttpConfig) ProtoMessage() {}

func (x *HttpProtocolOptions_ExplicitHttpConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpProtocolOptions_ExplicitHttpConfig.ProtoReflect.Descriptor instead.
func (*HttpProtocolOptions_ExplicitHttpConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescGZIP(), []int{0, 0}
}

func (m *HttpProtocolOptions_ExplicitHttpConfig) GetProtocolConfig() isHttpProtocolOptions_ExplicitHttpConfig_ProtocolConfig {
	if m != nil {
		return m.ProtocolConfig
	}
	return nil
}

func (x *HttpProtocolOptions_ExplicitHttpConfig) GetHttpProtocolOptions() *v3.Http1ProtocolOptions {
	if x, ok := x.GetProtocolConfig().(*HttpProtocolOptions_ExplicitHttpConfig_HttpProtocolOptions); ok {
		return x.HttpProtocolOptions
	}
	return nil
}

func (x *HttpProtocolOptions_ExplicitHttpConfig) GetHttp2ProtocolOptions() *v3.Http2ProtocolOptions {
	if x, ok := x.GetProtocolConfig().(*HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions); ok {
		return x.Http2ProtocolOptions
	}
	return nil
}

type isHttpProtocolOptions_ExplicitHttpConfig_ProtocolConfig interface {
	isHttpProtocolOptions_ExplicitHttpConfig_ProtocolConfig()
}

type HttpProtocolOptions_ExplicitHttpConfig_HttpProtocolOptions struct {
	HttpProtocolOptions *v3.Http1ProtocolOptions `protobuf:"bytes,1,opt,name=http_protocol_options,json=httpProtocolOptions,proto3,oneof"`
}

type HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions struct {
	Http2ProtocolOptions *v3.Http2ProtocolOptions `protobuf:"bytes,2,opt,name=http2_protocol_options,json=http2ProtocolOptions,proto3,oneof"`
}

func (*HttpProtocolOptions_ExplicitHttpConfig_HttpProtocolOptions) isHttpProtocolOptions_ExplicitHttpConfig_ProtocolConfig() {
}

func (*HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions) isHttpProtocolOptions_ExplicitHttpConfig_ProtocolConfig() {
}

// If this is used, the cluster can use either of the configured protocols, and
// will use whichever protocol was used by the downstream connection.
type HttpProtocolOptions_UseDownstreamHttpConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpProtocolOptions  *v3.Http1ProtocolOptions `protobuf:"bytes,1,opt,name=http_protocol_options,json=httpProtocolOptions,proto3" json:"http_protocol_options,omitempty"`
	Http2ProtocolOptions *v3.Http2ProtocolOptions `protobuf:"bytes,2,opt,name=http2_protocol_options,json=http2ProtocolOptions,proto3" json:"http2_protocol_options,omitempty"`
}

func (x *HttpProtocolOptions_UseDownstreamHttpConfig) Reset() {
	*x = HttpProtocolOptions_UseDownstreamHttpConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpProtocolOptions_UseDownstreamHttpConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpProtocolOptions_UseDownstreamHttpConfig) ProtoMessage() {}

func (x *HttpProtocolOptions_UseDownstreamHttpConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpProtocolOptions_UseDownstreamHttpConfig.ProtoReflect.Descriptor instead.
func (*HttpProtocolOptions_UseDownstreamHttpConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescGZIP(), []int{0, 1}
}

func (x *HttpProtocolOptions_UseDownstreamHttpConfig) GetHttpProtocolOptions() *v3.Http1ProtocolOptions {
	if x != nil {
		return x.HttpProtocolOptions
	}
	return nil
}

func (x *HttpProtocolOptions_UseDownstreamHttpConfig) GetHttp2ProtocolOptions() *v3.Http2ProtocolOptions {
	if x != nil {
		return x.Http2ProtocolOptions
	}
	return nil
}

// If this is used, the cluster can use either HTTP/1 or HTTP/2, and will use whichever
// protocol is negotiated by ALPN with the upstream.
// Clusters configured with *AutoHttpConfig* will use the highest available
// protocol; HTTP/2 if supported, otherwise HTTP/1.
// If the upstream does not support ALPN, *AutoHttpConfig* will fail over to HTTP/1.
// This can only be used with transport sockets which support ALPN. Using a
// transport socket which does not support ALPN will result in configuration
// failure. The transport layer may be configured with custom ALPN, but the default ALPN
// for the cluster (or if custom ALPN fails) will be "h2,http/1.1".
type HttpProtocolOptions_AutoHttpConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpProtocolOptions  *v3.Http1ProtocolOptions `protobuf:"bytes,1,opt,name=http_protocol_options,json=httpProtocolOptions,proto3" json:"http_protocol_options,omitempty"`
	Http2ProtocolOptions *v3.Http2ProtocolOptions `protobuf:"bytes,2,opt,name=http2_protocol_options,json=http2ProtocolOptions,proto3" json:"http2_protocol_options,omitempty"`
}

func (x *HttpProtocolOptions_AutoHttpConfig) Reset() {
	*x = HttpProtocolOptions_AutoHttpConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpProtocolOptions_AutoHttpConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpProtocolOptions_AutoHttpConfig) ProtoMessage() {}

func (x *HttpProtocolOptions_AutoHttpConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpProtocolOptions_AutoHttpConfig.ProtoReflect.Descriptor instead.
func (*HttpProtocolOptions_AutoHttpConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescGZIP(), []int{0, 2}
}

func (x *HttpProtocolOptions_AutoHttpConfig) GetHttpProtocolOptions() *v3.Http1ProtocolOptions {
	if x != nil {
		return x.HttpProtocolOptions
	}
	return nil
}

func (x *HttpProtocolOptions_AutoHttpConfig) GetHttp2ProtocolOptions() *v3.Http2ProtocolOptions {
	if x != nil {
		return x.Http2ProtocolOptions
	}
	return nil
}

var File_envoy_extensions_upstreams_http_v3_http_protocol_options_proto protoreflect.FileDescriptor

var file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x76, 0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x76, 0x33, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc7, 0x0a, 0x0a, 0x13, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6a, 0x0a, 0x1c, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x19, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x76, 0x0a, 0x1e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x74, 0x74,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x1b, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x7e, 0x0a,
	0x14, 0x65, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x33,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x48, 0x74, 0x74,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x12, 0x65, 0x78, 0x70, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x96, 0x01,
	0x0a, 0x1e, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x74, 0x74,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x1b, 0x75, 0x73, 0x65, 0x44, 0x6f,
	0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x69, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x46, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x75,
	0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x33,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0xf2, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x48, 0x74,
	0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x60, 0x0a, 0x15, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x31, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x13, 0x68, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x62, 0x0a, 0x16, 0x68, 0x74,
	0x74, 0x70, 0x32, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x14, 0x68, 0x74, 0x74, 0x70, 0x32, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x16,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x1a, 0xdb, 0x01, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x44, 0x6f,
	0x77, 0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x5e, 0x0a, 0x15, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x31, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x68,
	0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x60, 0x0a, 0x16, 0x68, 0x74, 0x74, 0x70, 0x32, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x32, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x14,
	0x68, 0x74, 0x74, 0x70, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xd2, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x6f, 0x48, 0x74, 0x74,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5e, 0x0a, 0x15, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x31, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x13, 0x68, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x60, 0x0a, 0x16, 0x68, 0x74, 0x74, 0x70, 0x32,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x14, 0x68, 0x74, 0x74, 0x70, 0x32, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x20, 0x0a, 0x19, 0x75, 0x70, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x56, 0x0a, 0x30, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x76, 0x33, 0x42,
	0x18, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescOnce sync.Once
	file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescData = file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDesc
)

func file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescGZIP() []byte {
	file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescData)
	})
	return file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDescData
}

var file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_goTypes = []interface{}{
	(*HttpProtocolOptions)(nil),                         // 0: envoy.extensions.upstreams.http.v3.HttpProtocolOptions
	(*HttpProtocolOptions_ExplicitHttpConfig)(nil),      // 1: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.ExplicitHttpConfig
	(*HttpProtocolOptions_UseDownstreamHttpConfig)(nil), // 2: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.UseDownstreamHttpConfig
	(*HttpProtocolOptions_AutoHttpConfig)(nil),          // 3: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.AutoHttpConfig
	(*v3.HttpProtocolOptions)(nil),                      // 4: envoy.config.core.v3.HttpProtocolOptions
	(*v3.UpstreamHttpProtocolOptions)(nil),              // 5: envoy.config.core.v3.UpstreamHttpProtocolOptions
	(*v3.Http1ProtocolOptions)(nil),                     // 6: envoy.config.core.v3.Http1ProtocolOptions
	(*v3.Http2ProtocolOptions)(nil),                     // 7: envoy.config.core.v3.Http2ProtocolOptions
}
var file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_depIdxs = []int32{
	4,  // 0: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.common_http_protocol_options:type_name -> envoy.config.core.v3.HttpProtocolOptions
	5,  // 1: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.upstream_http_protocol_options:type_name -> envoy.config.core.v3.UpstreamHttpProtocolOptions
	1,  // 2: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.explicit_http_config:type_name -> envoy.extensions.upstreams.http.v3.HttpProtocolOptions.ExplicitHttpConfig
	2,  // 3: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.use_downstream_protocol_config:type_name -> envoy.extensions.upstreams.http.v3.HttpProtocolOptions.UseDownstreamHttpConfig
	3,  // 4: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.auto_config:type_name -> envoy.extensions.upstreams.http.v3.HttpProtocolOptions.AutoHttpConfig
	6,  // 5: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.ExplicitHttpConfig.http_protocol_options:type_name -> envoy.config.core.v3.Http1ProtocolOptions
	7,  // 6: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.ExplicitHttpConfig.http2_protocol_options:type_name -> envoy.config.core.v3.Http2ProtocolOptions
	6,  // 7: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.UseDownstreamHttpConfig.http_protocol_options:type_name -> envoy.config.core.v3.Http1ProtocolOptions
	7,  // 8: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.UseDownstreamHttpConfig.http2_protocol_options:type_name -> envoy.config.core.v3.Http2ProtocolOptions
	6,  // 9: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.AutoHttpConfig.http_protocol_options:type_name -> envoy.config.core.v3.Http1ProtocolOptions
	7,  // 10: envoy.extensions.upstreams.http.v3.HttpProtocolOptions.AutoHttpConfig.http2_protocol_options:type_name -> envoy.config.core.v3.Http2ProtocolOptions
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_init() }
func file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_init() {
	if File_envoy_extensions_upstreams_http_v3_http_protocol_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpProtocolOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpProtocolOptions_ExplicitHttpConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpProtocolOptions_UseDownstreamHttpConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpProtocolOptions_AutoHttpConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HttpProtocolOptions_ExplicitHttpConfig_)(nil),
		(*HttpProtocolOptions_UseDownstreamProtocolConfig)(nil),
		(*HttpProtocolOptions_AutoConfig)(nil),
	}
	file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*HttpProtocolOptions_ExplicitHttpConfig_HttpProtocolOptions)(nil),
		(*HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_msgTypes,
	}.Build()
	File_envoy_extensions_upstreams_http_v3_http_protocol_options_proto = out.File
	file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_rawDesc = nil
	file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_goTypes = nil
	file_envoy_extensions_upstreams_http_v3_http_protocol_options_proto_depIdxs = nil
}
