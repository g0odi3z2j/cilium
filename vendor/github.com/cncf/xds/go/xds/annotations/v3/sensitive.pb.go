// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: xds/annotations/v3/sensitive.proto

package v3

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
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

var file_xds_annotations_v3_sensitive_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         61008053,
		Name:          "xds.annotations.v3.sensitive",
		Tag:           "varint,61008053,opt,name=sensitive",
		Filename:      "xds/annotations/v3/sensitive.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool sensitive = 61008053;
	E_Sensitive = &file_xds_annotations_v3_sensitive_proto_extTypes[0]
)

var File_xds_annotations_v3_sensitive_proto protoreflect.FileDescriptor

var file_xds_annotations_v3_sensitive_proto_rawDesc = []byte{
	0x0a, 0x22, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x78, 0x64, 0x73, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3e, 0x0a, 0x09, 0x73, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb5, 0xd1, 0x8b, 0x1d, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e, 0x63, 0x66, 0x2f, 0x78, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_xds_annotations_v3_sensitive_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_xds_annotations_v3_sensitive_proto_depIdxs = []int32{
	0, // 0: xds.annotations.v3.sensitive:extendee -> google.protobuf.FieldOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_xds_annotations_v3_sensitive_proto_init() }
func file_xds_annotations_v3_sensitive_proto_init() {
	if File_xds_annotations_v3_sensitive_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_xds_annotations_v3_sensitive_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_xds_annotations_v3_sensitive_proto_goTypes,
		DependencyIndexes: file_xds_annotations_v3_sensitive_proto_depIdxs,
		ExtensionInfos:    file_xds_annotations_v3_sensitive_proto_extTypes,
	}.Build()
	File_xds_annotations_v3_sensitive_proto = out.File
	file_xds_annotations_v3_sensitive_proto_rawDesc = nil
	file_xds_annotations_v3_sensitive_proto_goTypes = nil
	file_xds_annotations_v3_sensitive_proto_depIdxs = nil
}
