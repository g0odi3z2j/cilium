// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/tracing/v2/custom_tag.proto

package envoy_type_tracing_v2

import (
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/type/metadata/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Describes custom tags for the active span.
// [#next-free-field: 6]
type CustomTag struct {
	// Used to populate the tag name.
	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	// Used to specify what kind of custom tag.
	//
	// Types that are valid to be assigned to Type:
	//	*CustomTag_Literal_
	//	*CustomTag_Environment_
	//	*CustomTag_RequestHeader
	//	*CustomTag_Metadata_
	Type                 isCustomTag_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CustomTag) Reset()         { *m = CustomTag{} }
func (m *CustomTag) String() string { return proto.CompactTextString(m) }
func (*CustomTag) ProtoMessage()    {}
func (*CustomTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0}
}

func (m *CustomTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag.Unmarshal(m, b)
}
func (m *CustomTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag.Marshal(b, m, deterministic)
}
func (m *CustomTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag.Merge(m, src)
}
func (m *CustomTag) XXX_Size() int {
	return xxx_messageInfo_CustomTag.Size(m)
}
func (m *CustomTag) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag proto.InternalMessageInfo

func (m *CustomTag) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type isCustomTag_Type interface {
	isCustomTag_Type()
}

type CustomTag_Literal_ struct {
	Literal *CustomTag_Literal `protobuf:"bytes,2,opt,name=literal,proto3,oneof"`
}

type CustomTag_Environment_ struct {
	Environment *CustomTag_Environment `protobuf:"bytes,3,opt,name=environment,proto3,oneof"`
}

type CustomTag_RequestHeader struct {
	RequestHeader *CustomTag_Header `protobuf:"bytes,4,opt,name=request_header,json=requestHeader,proto3,oneof"`
}

type CustomTag_Metadata_ struct {
	Metadata *CustomTag_Metadata `protobuf:"bytes,5,opt,name=metadata,proto3,oneof"`
}

func (*CustomTag_Literal_) isCustomTag_Type() {}

func (*CustomTag_Environment_) isCustomTag_Type() {}

func (*CustomTag_RequestHeader) isCustomTag_Type() {}

func (*CustomTag_Metadata_) isCustomTag_Type() {}

func (m *CustomTag) GetType() isCustomTag_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CustomTag) GetLiteral() *CustomTag_Literal {
	if x, ok := m.GetType().(*CustomTag_Literal_); ok {
		return x.Literal
	}
	return nil
}

func (m *CustomTag) GetEnvironment() *CustomTag_Environment {
	if x, ok := m.GetType().(*CustomTag_Environment_); ok {
		return x.Environment
	}
	return nil
}

func (m *CustomTag) GetRequestHeader() *CustomTag_Header {
	if x, ok := m.GetType().(*CustomTag_RequestHeader); ok {
		return x.RequestHeader
	}
	return nil
}

func (m *CustomTag) GetMetadata() *CustomTag_Metadata {
	if x, ok := m.GetType().(*CustomTag_Metadata_); ok {
		return x.Metadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomTag) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomTag_Literal_)(nil),
		(*CustomTag_Environment_)(nil),
		(*CustomTag_RequestHeader)(nil),
		(*CustomTag_Metadata_)(nil),
	}
}

// Literal type custom tag with static value for the tag value.
type CustomTag_Literal struct {
	// Static literal value to populate the tag value.
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Literal) Reset()         { *m = CustomTag_Literal{} }
func (m *CustomTag_Literal) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Literal) ProtoMessage()    {}
func (*CustomTag_Literal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0, 0}
}

func (m *CustomTag_Literal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Literal.Unmarshal(m, b)
}
func (m *CustomTag_Literal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Literal.Marshal(b, m, deterministic)
}
func (m *CustomTag_Literal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Literal.Merge(m, src)
}
func (m *CustomTag_Literal) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Literal.Size(m)
}
func (m *CustomTag_Literal) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Literal.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Literal proto.InternalMessageInfo

func (m *CustomTag_Literal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Environment type custom tag with environment name and default value.
type CustomTag_Environment struct {
	// Environment variable name to obtain the value to populate the tag value.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When the environment variable is not found,
	// the tag value will be populated with this default value if specified,
	// otherwise no tag will be populated.
	DefaultValue         string   `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Environment) Reset()         { *m = CustomTag_Environment{} }
func (m *CustomTag_Environment) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Environment) ProtoMessage()    {}
func (*CustomTag_Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0, 1}
}

func (m *CustomTag_Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Environment.Unmarshal(m, b)
}
func (m *CustomTag_Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Environment.Marshal(b, m, deterministic)
}
func (m *CustomTag_Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Environment.Merge(m, src)
}
func (m *CustomTag_Environment) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Environment.Size(m)
}
func (m *CustomTag_Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Environment proto.InternalMessageInfo

func (m *CustomTag_Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomTag_Environment) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

// Header type custom tag with header name and default value.
type CustomTag_Header struct {
	// Header name to obtain the value to populate the tag value.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// When the header does not exist,
	// the tag value will be populated with this default value if specified,
	// otherwise no tag will be populated.
	DefaultValue         string   `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Header) Reset()         { *m = CustomTag_Header{} }
func (m *CustomTag_Header) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Header) ProtoMessage()    {}
func (*CustomTag_Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0, 2}
}

func (m *CustomTag_Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Header.Unmarshal(m, b)
}
func (m *CustomTag_Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Header.Marshal(b, m, deterministic)
}
func (m *CustomTag_Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Header.Merge(m, src)
}
func (m *CustomTag_Header) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Header.Size(m)
}
func (m *CustomTag_Header) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Header.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Header proto.InternalMessageInfo

func (m *CustomTag_Header) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CustomTag_Header) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

// Metadata type custom tag using
// :ref:`MetadataKey <envoy_api_msg_type.metadata.v2.MetadataKey>` to retrieve the protobuf value
// from :ref:`Metadata <envoy_api_msg_core.Metadata>`, and populate the tag value with
// `the canonical JSON <https://developers.google.com/protocol-buffers/docs/proto3#json>`_
// representation of it.
type CustomTag_Metadata struct {
	// Specify what kind of metadata to obtain tag value from.
	Kind *v2.MetadataKind `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// Metadata key to define the path to retrieve the tag value.
	MetadataKey *v2.MetadataKey `protobuf:"bytes,2,opt,name=metadata_key,json=metadataKey,proto3" json:"metadata_key,omitempty"`
	// When no valid metadata is found,
	// the tag value would be populated with this default value if specified,
	// otherwise no tag would be populated.
	DefaultValue         string   `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomTag_Metadata) Reset()         { *m = CustomTag_Metadata{} }
func (m *CustomTag_Metadata) String() string { return proto.CompactTextString(m) }
func (*CustomTag_Metadata) ProtoMessage()    {}
func (*CustomTag_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6b7ba7857eb6848, []int{0, 3}
}

func (m *CustomTag_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomTag_Metadata.Unmarshal(m, b)
}
func (m *CustomTag_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomTag_Metadata.Marshal(b, m, deterministic)
}
func (m *CustomTag_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomTag_Metadata.Merge(m, src)
}
func (m *CustomTag_Metadata) XXX_Size() int {
	return xxx_messageInfo_CustomTag_Metadata.Size(m)
}
func (m *CustomTag_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomTag_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_CustomTag_Metadata proto.InternalMessageInfo

func (m *CustomTag_Metadata) GetKind() *v2.MetadataKind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *CustomTag_Metadata) GetMetadataKey() *v2.MetadataKey {
	if m != nil {
		return m.MetadataKey
	}
	return nil
}

func (m *CustomTag_Metadata) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomTag)(nil), "envoy.type.tracing.v2.CustomTag")
	proto.RegisterType((*CustomTag_Literal)(nil), "envoy.type.tracing.v2.CustomTag.Literal")
	proto.RegisterType((*CustomTag_Environment)(nil), "envoy.type.tracing.v2.CustomTag.Environment")
	proto.RegisterType((*CustomTag_Header)(nil), "envoy.type.tracing.v2.CustomTag.Header")
	proto.RegisterType((*CustomTag_Metadata)(nil), "envoy.type.tracing.v2.CustomTag.Metadata")
}

func init() {
	proto.RegisterFile("envoy/type/tracing/v2/custom_tag.proto", fileDescriptor_d6b7ba7857eb6848)
}

var fileDescriptor_d6b7ba7857eb6848 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4b, 0xaf, 0x93, 0x40,
	0x14, 0xc7, 0xe1, 0x42, 0x6f, 0x6f, 0x0f, 0xb7, 0x5d, 0x4c, 0x62, 0x44, 0x8c, 0x49, 0x63, 0x7d,
	0x60, 0x62, 0x20, 0xc1, 0x85, 0xae, 0xf1, 0x45, 0x7c, 0x44, 0x32, 0x31, 0x6e, 0xc9, 0x58, 0x46,
	0x24, 0x85, 0xa1, 0x4e, 0x07, 0x22, 0x1f, 0xcb, 0x8f, 0xe4, 0xb7, 0x30, 0xae, 0x0c, 0xc3, 0xa3,
	0xc4, 0x36, 0x61, 0x71, 0x77, 0xc3, 0xe1, 0xff, 0xff, 0xcd, 0x9c, 0xff, 0x99, 0x81, 0x47, 0x94,
	0x55, 0x45, 0xed, 0x8a, 0x7a, 0x4f, 0x5d, 0xc1, 0xc9, 0x36, 0x65, 0x89, 0x5b, 0x79, 0xee, 0xb6,
	0x3c, 0x88, 0x22, 0x8f, 0x04, 0x49, 0x9c, 0x3d, 0x2f, 0x44, 0x81, 0x6e, 0x49, 0x9d, 0xd3, 0xe8,
	0x9c, 0x4e, 0xe7, 0x54, 0x9e, 0xf5, 0x70, 0x64, 0xcf, 0xa9, 0x20, 0x31, 0x11, 0xa4, 0xf1, 0xf7,
	0xeb, 0xd6, 0x6d, 0xdd, 0xae, 0x48, 0x96, 0xc6, 0x44, 0x50, 0xb7, 0x5f, 0xb4, 0x3f, 0xee, 0xff,
	0x9e, 0xc1, 0xe2, 0xa5, 0xdc, 0xeb, 0x33, 0x49, 0xd0, 0x1d, 0xd0, 0x04, 0x49, 0x4c, 0x75, 0xad,
	0xda, 0x0b, 0x7f, 0xfe, 0xd7, 0xd7, 0xf9, 0xc5, 0x5a, 0xc5, 0x4d, 0x0d, 0xbd, 0x82, 0x79, 0x96,
	0x0a, 0xca, 0x49, 0x66, 0x5e, 0xac, 0x55, 0xdb, 0xf0, 0x6c, 0xe7, 0xec, 0x89, 0x9c, 0x81, 0xe6,
	0x7c, 0x68, 0xf5, 0x81, 0x82, 0x7b, 0x2b, 0x0a, 0xc1, 0xa0, 0xac, 0x4a, 0x79, 0xc1, 0x72, 0xca,
	0x84, 0xa9, 0x49, 0xd2, 0xd3, 0x49, 0xd2, 0xeb, 0xa3, 0x27, 0x50, 0xf0, 0x18, 0x81, 0x42, 0x58,
	0x71, 0xfa, 0xa3, 0xa4, 0x07, 0x11, 0x7d, 0xa7, 0x24, 0xa6, 0xdc, 0xd4, 0x25, 0xf4, 0xf1, 0x24,
	0x34, 0x90, 0xf2, 0x40, 0xc1, 0xcb, 0x0e, 0xd0, 0x16, 0xd0, 0x5b, 0xb8, 0xea, 0xd3, 0x33, 0x67,
	0x92, 0xf5, 0x64, 0x92, 0xf5, 0xb1, 0x33, 0x04, 0x0a, 0x1e, 0xcc, 0x96, 0x0d, 0xf3, 0x2e, 0x02,
	0x74, 0x0f, 0x66, 0x15, 0xc9, 0x4a, 0xfa, 0x7f, 0xb4, 0x6d, 0xd5, 0xfa, 0x04, 0xc6, 0xa8, 0x45,
	0x74, 0x17, 0x74, 0x46, 0xf2, 0x13, 0xb1, 0x2c, 0xa2, 0x0d, 0x2c, 0x63, 0xfa, 0x8d, 0x94, 0x99,
	0x88, 0x5a, 0x64, 0x33, 0x8e, 0x05, 0xbe, 0xee, 0x8a, 0x5f, 0x24, 0xf0, 0x1d, 0x5c, 0x76, 0xdd,
	0xdc, 0x9c, 0xf5, 0x4b, 0x85, 0xab, 0xbe, 0x3f, 0xf4, 0x02, 0xf4, 0x5d, 0xca, 0x62, 0x89, 0x33,
	0xbc, 0x07, 0xe3, 0x60, 0x86, 0x2b, 0x57, 0x79, 0x43, 0x1e, 0xef, 0x53, 0x16, 0x63, 0xe9, 0x40,
	0x6f, 0xe0, 0xba, 0x57, 0x44, 0x3b, 0x5a, 0x77, 0xb7, 0x68, 0x33, 0x49, 0xa0, 0x35, 0x36, 0xf2,
	0xe3, 0xc7, 0xe9, 0x99, 0xb5, 0xd3, 0x33, 0xfb, 0x06, 0xe8, 0x0d, 0x11, 0x69, 0x7f, 0x7c, 0xd5,
	0x7f, 0x0e, 0x9b, 0xb4, 0x68, 0xf7, 0xd9, 0xf3, 0xe2, 0x67, 0x7d, 0x7e, 0x9a, 0xfe, 0x6a, 0x18,
	0x67, 0xd8, 0x3c, 0x8d, 0x50, 0xfd, 0x7a, 0x29, 0xdf, 0xc8, 0xb3, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x39, 0x6d, 0xdc, 0x0c, 0xa4, 0x03, 0x00, 0x00,
}
