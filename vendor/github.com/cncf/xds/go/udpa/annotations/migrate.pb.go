// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: udpa/annotations/migrate.proto

package annotations

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

type MigrateAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rename string `protobuf:"bytes,1,opt,name=rename,proto3" json:"rename,omitempty"`
}

func (x *MigrateAnnotation) Reset() {
	*x = MigrateAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_udpa_annotations_migrate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrateAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrateAnnotation) ProtoMessage() {}

func (x *MigrateAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_udpa_annotations_migrate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrateAnnotation.ProtoReflect.Descriptor instead.
func (*MigrateAnnotation) Descriptor() ([]byte, []int) {
	return file_udpa_annotations_migrate_proto_rawDescGZIP(), []int{0}
}

func (x *MigrateAnnotation) GetRename() string {
	if x != nil {
		return x.Rename
	}
	return ""
}

type FieldMigrateAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rename         string `protobuf:"bytes,1,opt,name=rename,proto3" json:"rename,omitempty"`
	OneofPromotion string `protobuf:"bytes,2,opt,name=oneof_promotion,json=oneofPromotion,proto3" json:"oneof_promotion,omitempty"`
}

func (x *FieldMigrateAnnotation) Reset() {
	*x = FieldMigrateAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_udpa_annotations_migrate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMigrateAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMigrateAnnotation) ProtoMessage() {}

func (x *FieldMigrateAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_udpa_annotations_migrate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMigrateAnnotation.ProtoReflect.Descriptor instead.
func (*FieldMigrateAnnotation) Descriptor() ([]byte, []int) {
	return file_udpa_annotations_migrate_proto_rawDescGZIP(), []int{1}
}

func (x *FieldMigrateAnnotation) GetRename() string {
	if x != nil {
		return x.Rename
	}
	return ""
}

func (x *FieldMigrateAnnotation) GetOneofPromotion() string {
	if x != nil {
		return x.OneofPromotion
	}
	return ""
}

type FileMigrateAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MoveToPackage string `protobuf:"bytes,2,opt,name=move_to_package,json=moveToPackage,proto3" json:"move_to_package,omitempty"`
}

func (x *FileMigrateAnnotation) Reset() {
	*x = FileMigrateAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_udpa_annotations_migrate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMigrateAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMigrateAnnotation) ProtoMessage() {}

func (x *FileMigrateAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_udpa_annotations_migrate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMigrateAnnotation.ProtoReflect.Descriptor instead.
func (*FileMigrateAnnotation) Descriptor() ([]byte, []int) {
	return file_udpa_annotations_migrate_proto_rawDescGZIP(), []int{2}
}

func (x *FileMigrateAnnotation) GetMoveToPackage() string {
	if x != nil {
		return x.MoveToPackage
	}
	return ""
}

var file_udpa_annotations_migrate_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MigrateAnnotation)(nil),
		Field:         171962766,
		Name:          "udpa.annotations.message_migrate",
		Tag:           "bytes,171962766,opt,name=message_migrate",
		Filename:      "udpa/annotations/migrate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldMigrateAnnotation)(nil),
		Field:         171962766,
		Name:          "udpa.annotations.field_migrate",
		Tag:           "bytes,171962766,opt,name=field_migrate",
		Filename:      "udpa/annotations/migrate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*MigrateAnnotation)(nil),
		Field:         171962766,
		Name:          "udpa.annotations.enum_migrate",
		Tag:           "bytes,171962766,opt,name=enum_migrate",
		Filename:      "udpa/annotations/migrate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*MigrateAnnotation)(nil),
		Field:         171962766,
		Name:          "udpa.annotations.enum_value_migrate",
		Tag:           "bytes,171962766,opt,name=enum_value_migrate",
		Filename:      "udpa/annotations/migrate.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileMigrateAnnotation)(nil),
		Field:         171962766,
		Name:          "udpa.annotations.file_migrate",
		Tag:           "bytes,171962766,opt,name=file_migrate",
		Filename:      "udpa/annotations/migrate.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional udpa.annotations.MigrateAnnotation message_migrate = 171962766;
	E_MessageMigrate = &file_udpa_annotations_migrate_proto_extTypes[0]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional udpa.annotations.FieldMigrateAnnotation field_migrate = 171962766;
	E_FieldMigrate = &file_udpa_annotations_migrate_proto_extTypes[1]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional udpa.annotations.MigrateAnnotation enum_migrate = 171962766;
	E_EnumMigrate = &file_udpa_annotations_migrate_proto_extTypes[2]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional udpa.annotations.MigrateAnnotation enum_value_migrate = 171962766;
	E_EnumValueMigrate = &file_udpa_annotations_migrate_proto_extTypes[3]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional udpa.annotations.FileMigrateAnnotation file_migrate = 171962766;
	E_FileMigrate = &file_udpa_annotations_migrate_proto_extTypes[4]
)

var File_udpa_annotations_migrate_proto protoreflect.FileDescriptor

var file_udpa_annotations_migrate_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x11, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x59, 0x0a, 0x16, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x70, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x15,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x74, 0x6f,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6d, 0x6f, 0x76, 0x65, 0x54, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x3a, 0x70, 0x0a,
	0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x8e, 0xe3, 0xff, 0x51, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x64, 0x70,
	0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x3a,
	0x6f, 0x0a, 0x0d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x8e, 0xe3, 0xff, 0x51, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x3a, 0x67, 0x0a, 0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8e,
	0xe3, 0xff, 0x51, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x6e,
	0x75, 0x6d, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x3a, 0x77, 0x0a, 0x12, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x8e, 0xe3, 0xff, 0x51, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x75, 0x64,
	0x70, 0x61, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x10, 0x65, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x3a, 0x6b, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x8e, 0xe3, 0xff, 0x51, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x75, 0x64, 0x70, 0x61,
	0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6e,
	0x63, 0x66, 0x2f, 0x78, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_udpa_annotations_migrate_proto_rawDescOnce sync.Once
	file_udpa_annotations_migrate_proto_rawDescData = file_udpa_annotations_migrate_proto_rawDesc
)

func file_udpa_annotations_migrate_proto_rawDescGZIP() []byte {
	file_udpa_annotations_migrate_proto_rawDescOnce.Do(func() {
		file_udpa_annotations_migrate_proto_rawDescData = protoimpl.X.CompressGZIP(file_udpa_annotations_migrate_proto_rawDescData)
	})
	return file_udpa_annotations_migrate_proto_rawDescData
}

var file_udpa_annotations_migrate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_udpa_annotations_migrate_proto_goTypes = []interface{}{
	(*MigrateAnnotation)(nil),             // 0: udpa.annotations.MigrateAnnotation
	(*FieldMigrateAnnotation)(nil),        // 1: udpa.annotations.FieldMigrateAnnotation
	(*FileMigrateAnnotation)(nil),         // 2: udpa.annotations.FileMigrateAnnotation
	(*descriptorpb.MessageOptions)(nil),   // 3: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),     // 4: google.protobuf.FieldOptions
	(*descriptorpb.EnumOptions)(nil),      // 5: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 6: google.protobuf.EnumValueOptions
	(*descriptorpb.FileOptions)(nil),      // 7: google.protobuf.FileOptions
}
var file_udpa_annotations_migrate_proto_depIdxs = []int32{
	3,  // 0: udpa.annotations.message_migrate:extendee -> google.protobuf.MessageOptions
	4,  // 1: udpa.annotations.field_migrate:extendee -> google.protobuf.FieldOptions
	5,  // 2: udpa.annotations.enum_migrate:extendee -> google.protobuf.EnumOptions
	6,  // 3: udpa.annotations.enum_value_migrate:extendee -> google.protobuf.EnumValueOptions
	7,  // 4: udpa.annotations.file_migrate:extendee -> google.protobuf.FileOptions
	0,  // 5: udpa.annotations.message_migrate:type_name -> udpa.annotations.MigrateAnnotation
	1,  // 6: udpa.annotations.field_migrate:type_name -> udpa.annotations.FieldMigrateAnnotation
	0,  // 7: udpa.annotations.enum_migrate:type_name -> udpa.annotations.MigrateAnnotation
	0,  // 8: udpa.annotations.enum_value_migrate:type_name -> udpa.annotations.MigrateAnnotation
	2,  // 9: udpa.annotations.file_migrate:type_name -> udpa.annotations.FileMigrateAnnotation
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	5,  // [5:10] is the sub-list for extension type_name
	0,  // [0:5] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_udpa_annotations_migrate_proto_init() }
func file_udpa_annotations_migrate_proto_init() {
	if File_udpa_annotations_migrate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_udpa_annotations_migrate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrateAnnotation); i {
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
		file_udpa_annotations_migrate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMigrateAnnotation); i {
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
		file_udpa_annotations_migrate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMigrateAnnotation); i {
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
			RawDescriptor: file_udpa_annotations_migrate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_udpa_annotations_migrate_proto_goTypes,
		DependencyIndexes: file_udpa_annotations_migrate_proto_depIdxs,
		MessageInfos:      file_udpa_annotations_migrate_proto_msgTypes,
		ExtensionInfos:    file_udpa_annotations_migrate_proto_extTypes,
	}.Build()
	File_udpa_annotations_migrate_proto = out.File
	file_udpa_annotations_migrate_proto_rawDesc = nil
	file_udpa_annotations_migrate_proto_goTypes = nil
	file_udpa_annotations_migrate_proto_depIdxs = nil
}
