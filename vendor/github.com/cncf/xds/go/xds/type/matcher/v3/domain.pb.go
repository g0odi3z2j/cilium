// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: xds/type/matcher/v3/domain.proto

package v3

import (
	_ "github.com/cncf/xds/go/xds/annotations/v3"
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

type ServerNameMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainMatchers []*ServerNameMatcher_DomainMatcher `protobuf:"bytes,1,rep,name=domain_matchers,json=domainMatchers,proto3" json:"domain_matchers,omitempty"`
}

func (x *ServerNameMatcher) Reset() {
	*x = ServerNameMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerNameMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerNameMatcher) ProtoMessage() {}

func (x *ServerNameMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerNameMatcher.ProtoReflect.Descriptor instead.
func (*ServerNameMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_domain_proto_rawDescGZIP(), []int{0}
}

func (x *ServerNameMatcher) GetDomainMatchers() []*ServerNameMatcher_DomainMatcher {
	if x != nil {
		return x.DomainMatchers
	}
	return nil
}

type ServerNameMatcher_DomainMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domains []string         `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty"`
	OnMatch *Matcher_OnMatch `protobuf:"bytes,2,opt,name=on_match,json=onMatch,proto3" json:"on_match,omitempty"`
}

func (x *ServerNameMatcher_DomainMatcher) Reset() {
	*x = ServerNameMatcher_DomainMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_xds_type_matcher_v3_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerNameMatcher_DomainMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerNameMatcher_DomainMatcher) ProtoMessage() {}

func (x *ServerNameMatcher_DomainMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_xds_type_matcher_v3_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerNameMatcher_DomainMatcher.ProtoReflect.Descriptor instead.
func (*ServerNameMatcher_DomainMatcher) Descriptor() ([]byte, []int) {
	return file_xds_type_matcher_v3_domain_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ServerNameMatcher_DomainMatcher) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *ServerNameMatcher_DomainMatcher) GetOnMatch() *Matcher_OnMatch {
	if x != nil {
		return x.OnMatch
	}
	return nil
}

var File_xds_type_matcher_v3_domain_proto protoreflect.FileDescriptor

var file_xds_type_matcher_v3_domain_proto_rawDesc = []byte{
	0x0a, 0x20, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x5d, 0x0a, 0x0f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x1a, 0x74, 0x0a, 0x0d, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x07, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x12, 0x3f,
	0x0a, 0x08, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x78, 0x64, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x4f,
	0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x42,
	0x6e, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x78, 0x64,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x42, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_xds_type_matcher_v3_domain_proto_rawDescOnce sync.Once
	file_xds_type_matcher_v3_domain_proto_rawDescData = file_xds_type_matcher_v3_domain_proto_rawDesc
)

func file_xds_type_matcher_v3_domain_proto_rawDescGZIP() []byte {
	file_xds_type_matcher_v3_domain_proto_rawDescOnce.Do(func() {
		file_xds_type_matcher_v3_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_xds_type_matcher_v3_domain_proto_rawDescData)
	})
	return file_xds_type_matcher_v3_domain_proto_rawDescData
}

var file_xds_type_matcher_v3_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_xds_type_matcher_v3_domain_proto_goTypes = []interface{}{
	(*ServerNameMatcher)(nil),               // 0: xds.type.matcher.v3.ServerNameMatcher
	(*ServerNameMatcher_DomainMatcher)(nil), // 1: xds.type.matcher.v3.ServerNameMatcher.DomainMatcher
	(*Matcher_OnMatch)(nil),                 // 2: xds.type.matcher.v3.Matcher.OnMatch
}
var file_xds_type_matcher_v3_domain_proto_depIdxs = []int32{
	1, // 0: xds.type.matcher.v3.ServerNameMatcher.domain_matchers:type_name -> xds.type.matcher.v3.ServerNameMatcher.DomainMatcher
	2, // 1: xds.type.matcher.v3.ServerNameMatcher.DomainMatcher.on_match:type_name -> xds.type.matcher.v3.Matcher.OnMatch
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_xds_type_matcher_v3_domain_proto_init() }
func file_xds_type_matcher_v3_domain_proto_init() {
	if File_xds_type_matcher_v3_domain_proto != nil {
		return
	}
	file_xds_type_matcher_v3_matcher_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_xds_type_matcher_v3_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerNameMatcher); i {
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
		file_xds_type_matcher_v3_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerNameMatcher_DomainMatcher); i {
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
			RawDescriptor: file_xds_type_matcher_v3_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_xds_type_matcher_v3_domain_proto_goTypes,
		DependencyIndexes: file_xds_type_matcher_v3_domain_proto_depIdxs,
		MessageInfos:      file_xds_type_matcher_v3_domain_proto_msgTypes,
	}.Build()
	File_xds_type_matcher_v3_domain_proto = out.File
	file_xds_type_matcher_v3_domain_proto_rawDesc = nil
	file_xds_type_matcher_v3_domain_proto_goTypes = nil
	file_xds_type_matcher_v3_domain_proto_depIdxs = nil
}
