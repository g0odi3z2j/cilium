// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.1
// source: envoy/config/listener/v3/api_listener.proto

package listenerv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// Describes a type of API listener, which is used in non-proxy clients. The type of API
// exposed to the non-proxy application depends on the type of API listener.
type ApiListener struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type in this field determines the type of API listener. At present, the following
	// types are supported:
	// envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager (HTTP)
	// envoy.extensions.filters.network.http_connection_manager.v3.EnvoyMobileHttpConnectionManager (HTTP)
	// [#next-major-version: In the v3 API, replace this Any field with a oneof containing the
	// specific config message for each type of API listener. We could not do this in v2 because
	// it would have caused circular dependencies for go protos: lds.proto depends on this file,
	// and http_connection_manager.proto depends on rds.proto, which is in the same directory as
	// lds.proto, so lds.proto cannot depend on this file.]
	ApiListener *anypb.Any `protobuf:"bytes,1,opt,name=api_listener,json=apiListener,proto3" json:"api_listener,omitempty"`
}

func (x *ApiListener) Reset() {
	*x = ApiListener{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v3_api_listener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiListener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiListener) ProtoMessage() {}

func (x *ApiListener) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v3_api_listener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiListener.ProtoReflect.Descriptor instead.
func (*ApiListener) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v3_api_listener_proto_rawDescGZIP(), []int{0}
}

func (x *ApiListener) GetApiListener() *anypb.Any {
	if x != nil {
		return x.ApiListener
	}
	return nil
}

var File_envoy_config_listener_v3_api_listener_proto protoreflect.FileDescriptor

var file_envoy_config_listener_v3_api_listener_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0b, 0x41, 0x70, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0b, 0x61, 0x70, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x3a, 0x2b, 0x9a, 0xc5,
	0x88, 0x1e, 0x26, 0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70,
	0x69, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x42, 0x90, 0x01, 0x0a, 0x26, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x42, 0x10, 0x41, 0x70, 0x69, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_listener_v3_api_listener_proto_rawDescOnce sync.Once
	file_envoy_config_listener_v3_api_listener_proto_rawDescData = file_envoy_config_listener_v3_api_listener_proto_rawDesc
)

func file_envoy_config_listener_v3_api_listener_proto_rawDescGZIP() []byte {
	file_envoy_config_listener_v3_api_listener_proto_rawDescOnce.Do(func() {
		file_envoy_config_listener_v3_api_listener_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_listener_v3_api_listener_proto_rawDescData)
	})
	return file_envoy_config_listener_v3_api_listener_proto_rawDescData
}

var file_envoy_config_listener_v3_api_listener_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_listener_v3_api_listener_proto_goTypes = []interface{}{
	(*ApiListener)(nil), // 0: envoy.config.listener.v3.ApiListener
	(*anypb.Any)(nil),   // 1: google.protobuf.Any
}
var file_envoy_config_listener_v3_api_listener_proto_depIdxs = []int32{
	1, // 0: envoy.config.listener.v3.ApiListener.api_listener:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_listener_v3_api_listener_proto_init() }
func file_envoy_config_listener_v3_api_listener_proto_init() {
	if File_envoy_config_listener_v3_api_listener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_listener_v3_api_listener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiListener); i {
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
			RawDescriptor: file_envoy_config_listener_v3_api_listener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_listener_v3_api_listener_proto_goTypes,
		DependencyIndexes: file_envoy_config_listener_v3_api_listener_proto_depIdxs,
		MessageInfos:      file_envoy_config_listener_v3_api_listener_proto_msgTypes,
	}.Build()
	File_envoy_config_listener_v3_api_listener_proto = out.File
	file_envoy_config_listener_v3_api_listener_proto_rawDesc = nil
	file_envoy_config_listener_v3_api_listener_proto_goTypes = nil
	file_envoy_config_listener_v3_api_listener_proto_depIdxs = nil
}
