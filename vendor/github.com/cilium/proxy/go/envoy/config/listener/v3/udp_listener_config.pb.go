// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/config/listener/v3/udp_listener_config.proto

package envoy_config_listener_v3

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/structpb"
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

type UdpListenerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Used to look up UDP listener factory, matches "raw_udp_listener" or
	// "quic_listener" to create a specific udp listener.
	// If not specified, treat as "raw_udp_listener".
	UdpListenerName string `protobuf:"bytes,1,opt,name=udp_listener_name,json=udpListenerName,proto3" json:"udp_listener_name,omitempty"`
	// Used to create a specific listener factory. To some factory, e.g.
	// "raw_udp_listener", config is not needed.
	//
	// Types that are assignable to ConfigType:
	//	*UdpListenerConfig_TypedConfig
	ConfigType isUdpListenerConfig_ConfigType `protobuf_oneof:"config_type"`
}

func (x *UdpListenerConfig) Reset() {
	*x = UdpListenerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UdpListenerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UdpListenerConfig) ProtoMessage() {}

func (x *UdpListenerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UdpListenerConfig.ProtoReflect.Descriptor instead.
func (*UdpListenerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP(), []int{0}
}

func (x *UdpListenerConfig) GetUdpListenerName() string {
	if x != nil {
		return x.UdpListenerName
	}
	return ""
}

func (m *UdpListenerConfig) GetConfigType() isUdpListenerConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (x *UdpListenerConfig) GetTypedConfig() *anypb.Any {
	if x, ok := x.GetConfigType().(*UdpListenerConfig_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

type isUdpListenerConfig_ConfigType interface {
	isUdpListenerConfig_ConfigType()
}

type UdpListenerConfig_TypedConfig struct {
	TypedConfig *anypb.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*UdpListenerConfig_TypedConfig) isUdpListenerConfig_ConfigType() {}

type ActiveRawUdpListenerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActiveRawUdpListenerConfig) Reset() {
	*x = ActiveRawUdpListenerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRawUdpListenerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRawUdpListenerConfig) ProtoMessage() {}

func (x *ActiveRawUdpListenerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRawUdpListenerConfig.ProtoReflect.Descriptor instead.
func (*ActiveRawUdpListenerConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP(), []int{1}
}

var File_envoy_config_listener_v3_udp_listener_config_proto protoreflect.FileDescriptor

var file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc = []byte{
	0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x11, 0x55, 0x64,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2a, 0x0a, 0x11, 0x75, 0x64, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x64, 0x70, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x2e, 0x9a, 0xc5, 0x88, 0x1e, 0x29, 0x0a, 0x27, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2e, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x22, 0x55, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x77,
	0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x77, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x4a, 0x0a, 0x26, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x2e, 0x76, 0x33, 0x42, 0x16, 0x55, 0x64, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescOnce sync.Once
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData = file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc
)

func file_envoy_config_listener_v3_udp_listener_config_proto_rawDescGZIP() []byte {
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDescOnce.Do(func() {
		file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData)
	})
	return file_envoy_config_listener_v3_udp_listener_config_proto_rawDescData
}

var file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_config_listener_v3_udp_listener_config_proto_goTypes = []interface{}{
	(*UdpListenerConfig)(nil),          // 0: envoy.config.listener.v3.UdpListenerConfig
	(*ActiveRawUdpListenerConfig)(nil), // 1: envoy.config.listener.v3.ActiveRawUdpListenerConfig
	(*anypb.Any)(nil),                  // 2: google.protobuf.Any
}
var file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs = []int32{
	2, // 0: envoy.config.listener.v3.UdpListenerConfig.typed_config:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_config_listener_v3_udp_listener_config_proto_init() }
func file_envoy_config_listener_v3_udp_listener_config_proto_init() {
	if File_envoy_config_listener_v3_udp_listener_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UdpListenerConfig); i {
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
		file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRawUdpListenerConfig); i {
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
	file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UdpListenerConfig_TypedConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_listener_v3_udp_listener_config_proto_goTypes,
		DependencyIndexes: file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs,
		MessageInfos:      file_envoy_config_listener_v3_udp_listener_config_proto_msgTypes,
	}.Build()
	File_envoy_config_listener_v3_udp_listener_config_proto = out.File
	file_envoy_config_listener_v3_udp_listener_config_proto_rawDesc = nil
	file_envoy_config_listener_v3_udp_listener_config_proto_goTypes = nil
	file_envoy_config_listener_v3_udp_listener_config_proto_depIdxs = nil
}
