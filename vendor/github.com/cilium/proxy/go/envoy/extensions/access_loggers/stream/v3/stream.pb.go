// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.23.1
// source: envoy/extensions/access_loggers/stream/v3/stream.proto

package streamv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
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

// Custom configuration for an :ref:`AccessLog <envoy_v3_api_msg_config.accesslog.v3.AccessLog>`
// that writes log entries directly to the operating system's standard output.
// [#extension: envoy.access_loggers.stdout]
type StdoutAccessLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AccessLogFormat:
	//
	//	*StdoutAccessLog_LogFormat
	AccessLogFormat isStdoutAccessLog_AccessLogFormat `protobuf_oneof:"access_log_format"`
}

func (x *StdoutAccessLog) Reset() {
	*x = StdoutAccessLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StdoutAccessLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StdoutAccessLog) ProtoMessage() {}

func (x *StdoutAccessLog) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StdoutAccessLog.ProtoReflect.Descriptor instead.
func (*StdoutAccessLog) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescGZIP(), []int{0}
}

func (m *StdoutAccessLog) GetAccessLogFormat() isStdoutAccessLog_AccessLogFormat {
	if m != nil {
		return m.AccessLogFormat
	}
	return nil
}

func (x *StdoutAccessLog) GetLogFormat() *v3.SubstitutionFormatString {
	if x, ok := x.GetAccessLogFormat().(*StdoutAccessLog_LogFormat); ok {
		return x.LogFormat
	}
	return nil
}

type isStdoutAccessLog_AccessLogFormat interface {
	isStdoutAccessLog_AccessLogFormat()
}

type StdoutAccessLog_LogFormat struct {
	// Configuration to form access log data and format.
	// If not specified, use :ref:`default format <config_access_log_default_format>`.
	LogFormat *v3.SubstitutionFormatString `protobuf:"bytes,1,opt,name=log_format,json=logFormat,proto3,oneof"`
}

func (*StdoutAccessLog_LogFormat) isStdoutAccessLog_AccessLogFormat() {}

// Custom configuration for an :ref:`AccessLog <envoy_v3_api_msg_config.accesslog.v3.AccessLog>`
// that writes log entries directly to the operating system's standard error.
// [#extension: envoy.access_loggers.stderr]
type StderrAccessLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AccessLogFormat:
	//
	//	*StderrAccessLog_LogFormat
	AccessLogFormat isStderrAccessLog_AccessLogFormat `protobuf_oneof:"access_log_format"`
}

func (x *StderrAccessLog) Reset() {
	*x = StderrAccessLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StderrAccessLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StderrAccessLog) ProtoMessage() {}

func (x *StderrAccessLog) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StderrAccessLog.ProtoReflect.Descriptor instead.
func (*StderrAccessLog) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescGZIP(), []int{1}
}

func (m *StderrAccessLog) GetAccessLogFormat() isStderrAccessLog_AccessLogFormat {
	if m != nil {
		return m.AccessLogFormat
	}
	return nil
}

func (x *StderrAccessLog) GetLogFormat() *v3.SubstitutionFormatString {
	if x, ok := x.GetAccessLogFormat().(*StderrAccessLog_LogFormat); ok {
		return x.LogFormat
	}
	return nil
}

type isStderrAccessLog_AccessLogFormat interface {
	isStderrAccessLog_AccessLogFormat()
}

type StderrAccessLog_LogFormat struct {
	// Configuration to form access log data and format.
	// If not specified, use :ref:`default format <config_access_log_default_format>`.
	LogFormat *v3.SubstitutionFormatString `protobuf:"bytes,1,opt,name=log_format,json=logFormat,proto3,oneof"`
}

func (*StderrAccessLog_LogFormat) isStderrAccessLog_AccessLogFormat() {}

var File_envoy_extensions_access_loggers_stream_v3_stream_proto protoreflect.FileDescriptor

var file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDesc = []byte{
	0x0a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x33, 0x1a, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69,
	0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x59, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x33, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x42, 0x13, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x5f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x59, 0x0a, 0x0a, 0x6c, 0x6f,
	0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x6c, 0x6f, 0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0xab, 0x01, 0x0a, 0x37, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x76, 0x33, 0x42, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x76, 0x33,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescOnce sync.Once
	file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescData = file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDesc
)

func file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescGZIP() []byte {
	file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescData)
	})
	return file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDescData
}

var file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_access_loggers_stream_v3_stream_proto_goTypes = []interface{}{
	(*StdoutAccessLog)(nil),             // 0: envoy.extensions.access_loggers.stream.v3.StdoutAccessLog
	(*StderrAccessLog)(nil),             // 1: envoy.extensions.access_loggers.stream.v3.StderrAccessLog
	(*v3.SubstitutionFormatString)(nil), // 2: envoy.config.core.v3.SubstitutionFormatString
}
var file_envoy_extensions_access_loggers_stream_v3_stream_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.access_loggers.stream.v3.StdoutAccessLog.log_format:type_name -> envoy.config.core.v3.SubstitutionFormatString
	2, // 1: envoy.extensions.access_loggers.stream.v3.StderrAccessLog.log_format:type_name -> envoy.config.core.v3.SubstitutionFormatString
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_access_loggers_stream_v3_stream_proto_init() }
func file_envoy_extensions_access_loggers_stream_v3_stream_proto_init() {
	if File_envoy_extensions_access_loggers_stream_v3_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StdoutAccessLog); i {
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
		file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StderrAccessLog); i {
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
	file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StdoutAccessLog_LogFormat)(nil),
	}
	file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StderrAccessLog_LogFormat)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_access_loggers_stream_v3_stream_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_access_loggers_stream_v3_stream_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_access_loggers_stream_v3_stream_proto_msgTypes,
	}.Build()
	File_envoy_extensions_access_loggers_stream_v3_stream_proto = out.File
	file_envoy_extensions_access_loggers_stream_v3_stream_proto_rawDesc = nil
	file_envoy_extensions_access_loggers_stream_v3_stream_proto_goTypes = nil
	file_envoy_extensions_access_loggers_stream_v3_stream_proto_depIdxs = nil
}
