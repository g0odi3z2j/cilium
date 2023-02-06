// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.4
// source: envoy/config/trace/v3/skywalking.proto

package tracev3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Configuration for the SkyWalking tracer. Please note that if SkyWalking tracer is used as the
// provider of http tracer, then
// :ref:`start_child_span <envoy_v3_api_field_extensions.filters.http.router.v3.Router.start_child_span>`
// in the router must be set to true to get the correct topology and tracing data. Moreover, SkyWalking
// Tracer does not support SkyWalking extension header (“sw8-x“) temporarily.
// [#extension: envoy.tracers.skywalking]
type SkyWalkingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SkyWalking collector service.
	GrpcService  *v3.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	ClientConfig *ClientConfig   `protobuf:"bytes,2,opt,name=client_config,json=clientConfig,proto3" json:"client_config,omitempty"`
}

func (x *SkyWalkingConfig) Reset() {
	*x = SkyWalkingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_trace_v3_skywalking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkyWalkingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkyWalkingConfig) ProtoMessage() {}

func (x *SkyWalkingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_trace_v3_skywalking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkyWalkingConfig.ProtoReflect.Descriptor instead.
func (*SkyWalkingConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_trace_v3_skywalking_proto_rawDescGZIP(), []int{0}
}

func (x *SkyWalkingConfig) GetGrpcService() *v3.GrpcService {
	if x != nil {
		return x.GrpcService
	}
	return nil
}

func (x *SkyWalkingConfig) GetClientConfig() *ClientConfig {
	if x != nil {
		return x.ClientConfig
	}
	return nil
}

// Client config for SkyWalking tracer.
type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Service name for SkyWalking tracer. If this field is empty, then local service cluster name
	// that configured by :ref:`Bootstrap node <envoy_v3_api_field_config.bootstrap.v3.Bootstrap.node>`
	// message's :ref:`cluster <envoy_v3_api_field_config.core.v3.Node.cluster>` field or command line
	// option :option:`--service-cluster` will be used. If both this field and local service cluster
	// name are empty, “EnvoyProxy“ is used as the service name by default.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Service instance name for SkyWalking tracer. If this field is empty, then local service node
	// that configured by :ref:`Bootstrap node <envoy_v3_api_field_config.bootstrap.v3.Bootstrap.node>`
	// message's :ref:`id <envoy_v3_api_field_config.core.v3.Node.id>` field or command line  option
	// :option:`--service-node` will be used. If both this field and local service node are empty,
	// “EnvoyProxy“ is used as the instance name by default.
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	// Authentication token config for SkyWalking. SkyWalking can use token authentication to secure
	// that monitoring application data can be trusted. In current version, Token is considered as a
	// simple string.
	// [#comment:TODO(wbpcode): Get backend token through the SDS API.]
	//
	// Types that are assignable to BackendTokenSpecifier:
	//
	//	*ClientConfig_BackendToken
	BackendTokenSpecifier isClientConfig_BackendTokenSpecifier `protobuf_oneof:"backend_token_specifier"`
	// Envoy caches the segment in memory when the SkyWalking backend service is temporarily unavailable.
	// This field specifies the maximum number of segments that can be cached. If not specified, the
	// default is 1024.
	MaxCacheSize *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=max_cache_size,json=maxCacheSize,proto3" json:"max_cache_size,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_trace_v3_skywalking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_trace_v3_skywalking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_trace_v3_skywalking_proto_rawDescGZIP(), []int{1}
}

func (x *ClientConfig) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ClientConfig) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (m *ClientConfig) GetBackendTokenSpecifier() isClientConfig_BackendTokenSpecifier {
	if m != nil {
		return m.BackendTokenSpecifier
	}
	return nil
}

func (x *ClientConfig) GetBackendToken() string {
	if x, ok := x.GetBackendTokenSpecifier().(*ClientConfig_BackendToken); ok {
		return x.BackendToken
	}
	return ""
}

func (x *ClientConfig) GetMaxCacheSize() *wrapperspb.UInt32Value {
	if x != nil {
		return x.MaxCacheSize
	}
	return nil
}

type isClientConfig_BackendTokenSpecifier interface {
	isClientConfig_BackendTokenSpecifier()
}

type ClientConfig_BackendToken struct {
	// Inline authentication token string.
	BackendToken string `protobuf:"bytes,3,opt,name=backend_token,json=backendToken,proto3,oneof"`
}

func (*ClientConfig_BackendToken) isClientConfig_BackendTokenSpecifier() {}

var File_envoy_config_trace_v3_skywalking_proto protoreflect.FileDescriptor

var file_envoy_config_trace_v3_skywalking_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x1a,
	0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x10, 0x53, 0x6b, 0x79, 0x57, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4e, 0x0a, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0xe4, 0x01, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xb8, 0xb7, 0x8b, 0xa4, 0x02, 0x01, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x42, 0x0a, 0x0e, 0x6d, 0x61, 0x78,
	0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0c, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x19, 0x0a,
	0x17, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0xb9, 0x01, 0x0a, 0x23, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33,
	0x42, 0x0f, 0x53, 0x6b, 0x79, 0x77, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76,
	0x33, 0x3b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x76, 0x33, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x2d, 0x12,
	0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x6b, 0x79, 0x77, 0x61, 0x6c,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_trace_v3_skywalking_proto_rawDescOnce sync.Once
	file_envoy_config_trace_v3_skywalking_proto_rawDescData = file_envoy_config_trace_v3_skywalking_proto_rawDesc
)

func file_envoy_config_trace_v3_skywalking_proto_rawDescGZIP() []byte {
	file_envoy_config_trace_v3_skywalking_proto_rawDescOnce.Do(func() {
		file_envoy_config_trace_v3_skywalking_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_trace_v3_skywalking_proto_rawDescData)
	})
	return file_envoy_config_trace_v3_skywalking_proto_rawDescData
}

var file_envoy_config_trace_v3_skywalking_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_trace_v3_skywalking_proto_goTypes = []interface{}{
	(*SkyWalkingConfig)(nil),       // 0: envoy.config.trace.v3.SkyWalkingConfig
	(*ClientConfig)(nil),           // 1: envoy.config.trace.v3.ClientConfig
	(*v3.GrpcService)(nil),         // 2: envoy.config.core.v3.GrpcService
	(*wrapperspb.UInt32Value)(nil), // 3: google.protobuf.UInt32Value
}
var file_envoy_config_trace_v3_skywalking_proto_depIdxs = []int32{
	2, // 0: envoy.config.trace.v3.SkyWalkingConfig.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	1, // 1: envoy.config.trace.v3.SkyWalkingConfig.client_config:type_name -> envoy.config.trace.v3.ClientConfig
	3, // 2: envoy.config.trace.v3.ClientConfig.max_cache_size:type_name -> google.protobuf.UInt32Value
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_config_trace_v3_skywalking_proto_init() }
func file_envoy_config_trace_v3_skywalking_proto_init() {
	if File_envoy_config_trace_v3_skywalking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_trace_v3_skywalking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkyWalkingConfig); i {
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
		file_envoy_config_trace_v3_skywalking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
	file_envoy_config_trace_v3_skywalking_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ClientConfig_BackendToken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_trace_v3_skywalking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_trace_v3_skywalking_proto_goTypes,
		DependencyIndexes: file_envoy_config_trace_v3_skywalking_proto_depIdxs,
		MessageInfos:      file_envoy_config_trace_v3_skywalking_proto_msgTypes,
	}.Build()
	File_envoy_config_trace_v3_skywalking_proto = out.File
	file_envoy_config_trace_v3_skywalking_proto_rawDesc = nil
	file_envoy_config_trace_v3_skywalking_proto_goTypes = nil
	file_envoy_config_trace_v3_skywalking_proto_depIdxs = nil
}
