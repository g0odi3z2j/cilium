// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.1
// source: envoy/extensions/filters/network/mongo_proxy/v3/mongo_proxy.proto

package mongo_proxyv3

import (
	v3 "github.com/cilium/proxy/go/envoy/extensions/filters/common/fault/v3"
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

// [#next-free-field: 6]
type MongoProxy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mongo_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The optional path to use for writing Mongo access logs. If not access log
	// path is specified no access logs will be written. Note that access log is
	// also gated :ref:`runtime <config_network_filters_mongo_proxy_runtime>`.
	AccessLog string `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	// Inject a fixed delay before proxying a Mongo operation. Delays are
	// applied to the following MongoDB operations: Query, Insert, GetMore,
	// and KillCursors. Once an active delay is in progress, all incoming
	// data up until the timer event fires will be a part of the delay.
	Delay *v3.FaultDelay `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
	// Flag to specify whether :ref:`dynamic metadata
	// <config_network_filters_mongo_proxy_dynamic_metadata>` should be emitted. Defaults to false.
	EmitDynamicMetadata bool `protobuf:"varint,4,opt,name=emit_dynamic_metadata,json=emitDynamicMetadata,proto3" json:"emit_dynamic_metadata,omitempty"`
	// List of commands to emit metrics for. Defaults to "delete", "insert", and "update".
	// Note that metrics will not be emitted for "find" commands, since those are considered
	// queries, and metrics for those are emitted under a dedicated "query" namespace.
	Commands []string `protobuf:"bytes,5,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *MongoProxy) Reset() {
	*x = MongoProxy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MongoProxy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MongoProxy) ProtoMessage() {}

func (x *MongoProxy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MongoProxy.ProtoReflect.Descriptor instead.
func (*MongoProxy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *MongoProxy) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *MongoProxy) GetAccessLog() string {
	if x != nil {
		return x.AccessLog
	}
	return ""
}

func (x *MongoProxy) GetDelay() *v3.FaultDelay {
	if x != nil {
		return x.Delay
	}
	return nil
}

func (x *MongoProxy) GetEmitDynamicMetadata() bool {
	if x != nil {
		return x.EmitDynamicMetadata
	}
	return false
}

func (x *MongoProxy) GetCommands() []string {
	if x != nil {
		return x.Commands
	}
	return nil
}

var File_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDesc = []byte{
	0x0a, 0x41, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76,
	0x33, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x33, 0x1a, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x02, 0x0a, 0x0a, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x4a, 0x0a,
	0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x46, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c,
	0x61, 0x79, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x6d, 0x69,
	0x74, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x65, 0x6d, 0x69, 0x74, 0x44, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x3a, 0x3c, 0x9a, 0xc5, 0x88, 0x1e, 0x37,
	0x0a, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x6f,
	0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x4d, 0x6f, 0x6e,
	0x67, 0x6f, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x42, 0xc0, 0x01, 0x0a, 0x3d, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x0f, 0x4d, 0x6f, 0x6e, 0x67, 0x6f,
	0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x64, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDescData = file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_goTypes = []interface{}{
	(*MongoProxy)(nil),    // 0: envoy.extensions.filters.network.mongo_proxy.v3.MongoProxy
	(*v3.FaultDelay)(nil), // 1: envoy.extensions.filters.common.fault.v3.FaultDelay
}
var file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.network.mongo_proxy.v3.MongoProxy.delay:type_name -> envoy.extensions.filters.common.fault.v3.FaultDelay
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_init() }
func file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_init() {
	if File_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MongoProxy); i {
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
			RawDescriptor: file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto = out.File
	file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_network_mongo_proxy_v3_mongo_proxy_proto_depIdxs = nil
}
