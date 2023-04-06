// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.12
// source: envoy/config/core/v3/grpc_method_list.proto

package corev3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// A list of gRPC methods which can be used as an allowlist, for example.
type GrpcMethodList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*GrpcMethodList_Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *GrpcMethodList) Reset() {
	*x = GrpcMethodList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_grpc_method_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcMethodList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcMethodList) ProtoMessage() {}

func (x *GrpcMethodList) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_grpc_method_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcMethodList.ProtoReflect.Descriptor instead.
func (*GrpcMethodList) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_grpc_method_list_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcMethodList) GetServices() []*GrpcMethodList_Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type GrpcMethodList_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the gRPC service.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The names of the gRPC methods in this service.
	MethodNames []string `protobuf:"bytes,2,rep,name=method_names,json=methodNames,proto3" json:"method_names,omitempty"`
}

func (x *GrpcMethodList_Service) Reset() {
	*x = GrpcMethodList_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_grpc_method_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcMethodList_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcMethodList_Service) ProtoMessage() {}

func (x *GrpcMethodList_Service) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_grpc_method_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcMethodList_Service.ProtoReflect.Descriptor instead.
func (*GrpcMethodList_Service) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_grpc_method_list_proto_rawDescGZIP(), []int{0, 0}
}

func (x *GrpcMethodList_Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GrpcMethodList_Service) GetMethodNames() []string {
	if x != nil {
		return x.MethodNames
	}
	return nil
}

var File_envoy_config_core_v3_grpc_method_list_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_grpc_method_list_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x02, 0x0a, 0x0e, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x84, 0x01, 0x0a, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x3a, 0x2f, 0x9a, 0xc5, 0x88, 0x1e, 0x2a, 0x0a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x72, 0x70, 0x63,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x3a, 0x27, 0x9a, 0xc5, 0x88, 0x1e, 0x22, 0x0a, 0x20, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x87, 0x01, 0x0a, 0x22,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x42, 0x13, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6f, 0x72, 0x65, 0x76, 0x33, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_grpc_method_list_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_grpc_method_list_proto_rawDescData = file_envoy_config_core_v3_grpc_method_list_proto_rawDesc
)

func file_envoy_config_core_v3_grpc_method_list_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_grpc_method_list_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_grpc_method_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_grpc_method_list_proto_rawDescData)
	})
	return file_envoy_config_core_v3_grpc_method_list_proto_rawDescData
}

var file_envoy_config_core_v3_grpc_method_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_core_v3_grpc_method_list_proto_goTypes = []interface{}{
	(*GrpcMethodList)(nil),         // 0: envoy.config.core.v3.GrpcMethodList
	(*GrpcMethodList_Service)(nil), // 1: envoy.config.core.v3.GrpcMethodList.Service
}
var file_envoy_config_core_v3_grpc_method_list_proto_depIdxs = []int32{
	1, // 0: envoy.config.core.v3.GrpcMethodList.services:type_name -> envoy.config.core.v3.GrpcMethodList.Service
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_grpc_method_list_proto_init() }
func file_envoy_config_core_v3_grpc_method_list_proto_init() {
	if File_envoy_config_core_v3_grpc_method_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_grpc_method_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcMethodList); i {
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
		file_envoy_config_core_v3_grpc_method_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcMethodList_Service); i {
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
			RawDescriptor: file_envoy_config_core_v3_grpc_method_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_grpc_method_list_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_grpc_method_list_proto_depIdxs,
		MessageInfos:      file_envoy_config_core_v3_grpc_method_list_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_grpc_method_list_proto = out.File
	file_envoy_config_core_v3_grpc_method_list_proto_rawDesc = nil
	file_envoy_config_core_v3_grpc_method_list_proto_goTypes = nil
	file_envoy_config_core_v3_grpc_method_list_proto_depIdxs = nil
}
