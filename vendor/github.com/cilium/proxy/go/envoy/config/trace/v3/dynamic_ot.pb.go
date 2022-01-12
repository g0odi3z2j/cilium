// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.1
// source: envoy/config/trace/v3/dynamic_ot.proto

package tracev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

// DynamicOtConfig is used to dynamically load a tracer from a shared library
// that implements the `OpenTracing dynamic loading API
// <https://github.com/opentracing/opentracing-cpp>`_.
// [#extension: envoy.tracers.dynamic_ot]
type DynamicOtConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dynamic library implementing the `OpenTracing API
	// <https://github.com/opentracing/opentracing-cpp>`_.
	Library string `protobuf:"bytes,1,opt,name=library,proto3" json:"library,omitempty"`
	// The configuration to use when creating a tracer from the given dynamic
	// library.
	Config *structpb.Struct `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *DynamicOtConfig) Reset() {
	*x = DynamicOtConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_trace_v3_dynamic_ot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicOtConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicOtConfig) ProtoMessage() {}

func (x *DynamicOtConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_trace_v3_dynamic_ot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicOtConfig.ProtoReflect.Descriptor instead.
func (*DynamicOtConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_trace_v3_dynamic_ot_proto_rawDescGZIP(), []int{0}
}

func (x *DynamicOtConfig) GetLibrary() string {
	if x != nil {
		return x.Library
	}
	return ""
}

func (x *DynamicOtConfig) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_envoy_config_trace_v3_dynamic_ot_proto protoreflect.FileDescriptor

var file_envoy_config_trace_v3_dynamic_ot_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x4f, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x07,
	0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x12,
	0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x3a, 0x2c, 0x9a, 0xc5, 0x88, 0x1e, 0x27, 0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4f, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0xb8,
	0x01, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4f,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x76, 0x33, 0xf2, 0x98,
	0xfe, 0x8f, 0x05, 0x2d, 0x12, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x73, 0x2e, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6f, 0x74, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_config_trace_v3_dynamic_ot_proto_rawDescOnce sync.Once
	file_envoy_config_trace_v3_dynamic_ot_proto_rawDescData = file_envoy_config_trace_v3_dynamic_ot_proto_rawDesc
)

func file_envoy_config_trace_v3_dynamic_ot_proto_rawDescGZIP() []byte {
	file_envoy_config_trace_v3_dynamic_ot_proto_rawDescOnce.Do(func() {
		file_envoy_config_trace_v3_dynamic_ot_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_trace_v3_dynamic_ot_proto_rawDescData)
	})
	return file_envoy_config_trace_v3_dynamic_ot_proto_rawDescData
}

var file_envoy_config_trace_v3_dynamic_ot_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_trace_v3_dynamic_ot_proto_goTypes = []interface{}{
	(*DynamicOtConfig)(nil), // 0: envoy.config.trace.v3.DynamicOtConfig
	(*structpb.Struct)(nil), // 1: google.protobuf.Struct
}
var file_envoy_config_trace_v3_dynamic_ot_proto_depIdxs = []int32{
	1, // 0: envoy.config.trace.v3.DynamicOtConfig.config:type_name -> google.protobuf.Struct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_trace_v3_dynamic_ot_proto_init() }
func file_envoy_config_trace_v3_dynamic_ot_proto_init() {
	if File_envoy_config_trace_v3_dynamic_ot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_trace_v3_dynamic_ot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicOtConfig); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_trace_v3_dynamic_ot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_trace_v3_dynamic_ot_proto_goTypes,
		DependencyIndexes: file_envoy_config_trace_v3_dynamic_ot_proto_depIdxs,
		MessageInfos:      file_envoy_config_trace_v3_dynamic_ot_proto_msgTypes,
	}.Build()
	File_envoy_config_trace_v3_dynamic_ot_proto = out.File
	file_envoy_config_trace_v3_dynamic_ot_proto_rawDesc = nil
	file_envoy_config_trace_v3_dynamic_ot_proto_goTypes = nil
	file_envoy_config_trace_v3_dynamic_ot_proto_depIdxs = nil
}
