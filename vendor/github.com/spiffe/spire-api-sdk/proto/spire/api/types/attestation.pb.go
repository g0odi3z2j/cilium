// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/api/types/attestation.proto

package types

import (
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

type AttestationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of attestation data. This is typically the name of the plugin
	// that produced that data.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The attestation data payload.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *AttestationData) Reset() {
	*x = AttestationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_types_attestation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestationData) ProtoMessage() {}

func (x *AttestationData) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_types_attestation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestationData.ProtoReflect.Descriptor instead.
func (*AttestationData) Descriptor() ([]byte, []int) {
	return file_spire_api_types_attestation_proto_rawDescGZIP(), []int{0}
}

func (x *AttestationData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttestationData) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_spire_api_types_attestation_proto protoreflect.FileDescriptor

var file_spire_api_types_attestation_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_types_attestation_proto_rawDescOnce sync.Once
	file_spire_api_types_attestation_proto_rawDescData = file_spire_api_types_attestation_proto_rawDesc
)

func file_spire_api_types_attestation_proto_rawDescGZIP() []byte {
	file_spire_api_types_attestation_proto_rawDescOnce.Do(func() {
		file_spire_api_types_attestation_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_types_attestation_proto_rawDescData)
	})
	return file_spire_api_types_attestation_proto_rawDescData
}

var file_spire_api_types_attestation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spire_api_types_attestation_proto_goTypes = []interface{}{
	(*AttestationData)(nil), // 0: spire.api.types.AttestationData
}
var file_spire_api_types_attestation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spire_api_types_attestation_proto_init() }
func file_spire_api_types_attestation_proto_init() {
	if File_spire_api_types_attestation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_api_types_attestation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestationData); i {
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
			RawDescriptor: file_spire_api_types_attestation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spire_api_types_attestation_proto_goTypes,
		DependencyIndexes: file_spire_api_types_attestation_proto_depIdxs,
		MessageInfos:      file_spire_api_types_attestation_proto_msgTypes,
	}.Build()
	File_spire_api_types_attestation_proto = out.File
	file_spire_api_types_attestation_proto_rawDesc = nil
	file_spire_api_types_attestation_proto_goTypes = nil
	file_spire_api_types_attestation_proto_depIdxs = nil
}
