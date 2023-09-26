// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.1
// source: envoy/extensions/request_id/uuid/v3/uuid.proto

package uuidv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// Configuration for the default UUID request ID extension which has the following behavior:
//
//  1. Request ID is propagated using the :ref:`x-request-id
//     <config_http_conn_man_headers_x-request-id>` header.
//
//  2. Request ID is a universally unique identifier `(UUID4)
//     <https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)>`_.
//
//  3. Tracing decision (sampled, forced, etc) is set in 14th nibble of the UUID. By default this will
//     overwrite existing UUIDs received in the “x-request-id“ header if the trace sampling decision
//     is changed. The 14th nibble of the UUID4 has been chosen because it is fixed to '4' by the
//     standard. Thus, '4' indicates a default UUID and no trace status. This nibble is swapped to:
//
//     a. '9': Sampled.
//     b. 'a': Force traced due to server-side override.
//     c. 'b': Force traced due to client-side request ID joining.
//
//     See the :ref:`x-request-id <config_http_conn_man_headers_x-request-id>` documentation for
//     more information.
type UuidRequestIdConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the implementation alters the UUID to contain the trace sampling decision as per the
	// “UuidRequestIdConfig“ message documentation. This defaults to true. If disabled no
	// modification to the UUID will be performed. It is important to note that if disabled,
	// stable sampling of traces, access logs, etc. will no longer work and only random sampling will
	// be possible.
	PackTraceReason *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=pack_trace_reason,json=packTraceReason,proto3" json:"pack_trace_reason,omitempty"`
	// Set whether to use :ref:`x-request-id<config_http_conn_man_headers_x-request-id>` for sampling or not.
	// This defaults to true. See the :ref:`context propagation <arch_overview_tracing_context_propagation>`
	// overview for more information.
	UseRequestIdForTraceSampling *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=use_request_id_for_trace_sampling,json=useRequestIdForTraceSampling,proto3" json:"use_request_id_for_trace_sampling,omitempty"`
}

func (x *UuidRequestIdConfig) Reset() {
	*x = UuidRequestIdConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_request_id_uuid_v3_uuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UuidRequestIdConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UuidRequestIdConfig) ProtoMessage() {}

func (x *UuidRequestIdConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_request_id_uuid_v3_uuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UuidRequestIdConfig.ProtoReflect.Descriptor instead.
func (*UuidRequestIdConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDescGZIP(), []int{0}
}

func (x *UuidRequestIdConfig) GetPackTraceReason() *wrapperspb.BoolValue {
	if x != nil {
		return x.PackTraceReason
	}
	return nil
}

func (x *UuidRequestIdConfig) GetUseRequestIdForTraceSampling() *wrapperspb.BoolValue {
	if x != nil {
		return x.UseRequestIdForTraceSampling
	}
	return nil
}

var File_envoy_extensions_request_id_uuid_v3_uuid_proto protoreflect.FileDescriptor

var file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x2f, 0x75, 0x75,
	0x69, 0x64, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x2e, 0x75, 0x75,
	0x69, 0x64, 0x2e, 0x76, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x13, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x46, 0x0a, 0x11,
	0x70, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x63, 0x0a, 0x21, 0x75, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x1c, 0x75, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x9b, 0x01, 0x0a, 0x31, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x76, 0x33, 0x42,
	0x09, 0x55, 0x75, 0x69, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x2f, 0x75, 0x75, 0x69, 0x64, 0x2f, 0x76, 0x33, 0x3b, 0x75, 0x75, 0x69, 0x64, 0x76, 0x33, 0xba,
	0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDescOnce sync.Once
	file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDescData = file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDesc
)

func file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDescGZIP() []byte {
	file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDescData)
	})
	return file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDescData
}

var file_envoy_extensions_request_id_uuid_v3_uuid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_request_id_uuid_v3_uuid_proto_goTypes = []interface{}{
	(*UuidRequestIdConfig)(nil),  // 0: envoy.extensions.request_id.uuid.v3.UuidRequestIdConfig
	(*wrapperspb.BoolValue)(nil), // 1: google.protobuf.BoolValue
}
var file_envoy_extensions_request_id_uuid_v3_uuid_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.request_id.uuid.v3.UuidRequestIdConfig.pack_trace_reason:type_name -> google.protobuf.BoolValue
	1, // 1: envoy.extensions.request_id.uuid.v3.UuidRequestIdConfig.use_request_id_for_trace_sampling:type_name -> google.protobuf.BoolValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_request_id_uuid_v3_uuid_proto_init() }
func file_envoy_extensions_request_id_uuid_v3_uuid_proto_init() {
	if File_envoy_extensions_request_id_uuid_v3_uuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_request_id_uuid_v3_uuid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UuidRequestIdConfig); i {
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
			RawDescriptor: file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_request_id_uuid_v3_uuid_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_request_id_uuid_v3_uuid_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_request_id_uuid_v3_uuid_proto_msgTypes,
	}.Build()
	File_envoy_extensions_request_id_uuid_v3_uuid_proto = out.File
	file_envoy_extensions_request_id_uuid_v3_uuid_proto_rawDesc = nil
	file_envoy_extensions_request_id_uuid_v3_uuid_proto_goTypes = nil
	file_envoy_extensions_request_id_uuid_v3_uuid_proto_depIdxs = nil
}
