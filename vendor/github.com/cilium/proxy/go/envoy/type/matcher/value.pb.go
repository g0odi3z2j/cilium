// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/value.proto

package envoy_type_matcher

import (
	fmt "fmt"
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

// Specifies the way to match a ProtobufWkt::Value. Primitive values and ListValue are supported.
// StructValue is not supported and is always not matched.
// [#next-free-field: 7]
type ValueMatcher struct {
	// Specifies how to match a value.
	//
	// Types that are valid to be assigned to MatchPattern:
	//	*ValueMatcher_NullMatch_
	//	*ValueMatcher_DoubleMatch
	//	*ValueMatcher_StringMatch
	//	*ValueMatcher_BoolMatch
	//	*ValueMatcher_PresentMatch
	//	*ValueMatcher_ListMatch
	MatchPattern         isValueMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ValueMatcher) Reset()         { *m = ValueMatcher{} }
func (m *ValueMatcher) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher) ProtoMessage()    {}
func (*ValueMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{0}
}

func (m *ValueMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueMatcher.Unmarshal(m, b)
}
func (m *ValueMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueMatcher.Marshal(b, m, deterministic)
}
func (m *ValueMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher.Merge(m, src)
}
func (m *ValueMatcher) XXX_Size() int {
	return xxx_messageInfo_ValueMatcher.Size(m)
}
func (m *ValueMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher proto.InternalMessageInfo

type isValueMatcher_MatchPattern interface {
	isValueMatcher_MatchPattern()
}

type ValueMatcher_NullMatch_ struct {
	NullMatch *ValueMatcher_NullMatch `protobuf:"bytes,1,opt,name=null_match,json=nullMatch,proto3,oneof"`
}

type ValueMatcher_DoubleMatch struct {
	DoubleMatch *DoubleMatcher `protobuf:"bytes,2,opt,name=double_match,json=doubleMatch,proto3,oneof"`
}

type ValueMatcher_StringMatch struct {
	StringMatch *StringMatcher `protobuf:"bytes,3,opt,name=string_match,json=stringMatch,proto3,oneof"`
}

type ValueMatcher_BoolMatch struct {
	BoolMatch bool `protobuf:"varint,4,opt,name=bool_match,json=boolMatch,proto3,oneof"`
}

type ValueMatcher_PresentMatch struct {
	PresentMatch bool `protobuf:"varint,5,opt,name=present_match,json=presentMatch,proto3,oneof"`
}

type ValueMatcher_ListMatch struct {
	ListMatch *ListMatcher `protobuf:"bytes,6,opt,name=list_match,json=listMatch,proto3,oneof"`
}

func (*ValueMatcher_NullMatch_) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_DoubleMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_StringMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_BoolMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_PresentMatch) isValueMatcher_MatchPattern() {}

func (*ValueMatcher_ListMatch) isValueMatcher_MatchPattern() {}

func (m *ValueMatcher) GetMatchPattern() isValueMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ValueMatcher) GetNullMatch() *ValueMatcher_NullMatch {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_NullMatch_); ok {
		return x.NullMatch
	}
	return nil
}

func (m *ValueMatcher) GetDoubleMatch() *DoubleMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_DoubleMatch); ok {
		return x.DoubleMatch
	}
	return nil
}

func (m *ValueMatcher) GetStringMatch() *StringMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_StringMatch); ok {
		return x.StringMatch
	}
	return nil
}

func (m *ValueMatcher) GetBoolMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_BoolMatch); ok {
		return x.BoolMatch
	}
	return false
}

func (m *ValueMatcher) GetPresentMatch() bool {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_PresentMatch); ok {
		return x.PresentMatch
	}
	return false
}

func (m *ValueMatcher) GetListMatch() *ListMatcher {
	if x, ok := m.GetMatchPattern().(*ValueMatcher_ListMatch); ok {
		return x.ListMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ValueMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ValueMatcher_NullMatch_)(nil),
		(*ValueMatcher_DoubleMatch)(nil),
		(*ValueMatcher_StringMatch)(nil),
		(*ValueMatcher_BoolMatch)(nil),
		(*ValueMatcher_PresentMatch)(nil),
		(*ValueMatcher_ListMatch)(nil),
	}
}

// NullMatch is an empty message to specify a null value.
type ValueMatcher_NullMatch struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueMatcher_NullMatch) Reset()         { *m = ValueMatcher_NullMatch{} }
func (m *ValueMatcher_NullMatch) String() string { return proto.CompactTextString(m) }
func (*ValueMatcher_NullMatch) ProtoMessage()    {}
func (*ValueMatcher_NullMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{0, 0}
}

func (m *ValueMatcher_NullMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueMatcher_NullMatch.Unmarshal(m, b)
}
func (m *ValueMatcher_NullMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueMatcher_NullMatch.Marshal(b, m, deterministic)
}
func (m *ValueMatcher_NullMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueMatcher_NullMatch.Merge(m, src)
}
func (m *ValueMatcher_NullMatch) XXX_Size() int {
	return xxx_messageInfo_ValueMatcher_NullMatch.Size(m)
}
func (m *ValueMatcher_NullMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueMatcher_NullMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ValueMatcher_NullMatch proto.InternalMessageInfo

// Specifies the way to match a list value.
type ListMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*ListMatcher_OneOf
	MatchPattern         isListMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ListMatcher) Reset()         { *m = ListMatcher{} }
func (m *ListMatcher) String() string { return proto.CompactTextString(m) }
func (*ListMatcher) ProtoMessage()    {}
func (*ListMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_145b36501d266253, []int{1}
}

func (m *ListMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMatcher.Unmarshal(m, b)
}
func (m *ListMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMatcher.Marshal(b, m, deterministic)
}
func (m *ListMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMatcher.Merge(m, src)
}
func (m *ListMatcher) XXX_Size() int {
	return xxx_messageInfo_ListMatcher.Size(m)
}
func (m *ListMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_ListMatcher proto.InternalMessageInfo

type isListMatcher_MatchPattern interface {
	isListMatcher_MatchPattern()
}

type ListMatcher_OneOf struct {
	OneOf *ValueMatcher `protobuf:"bytes,1,opt,name=one_of,json=oneOf,proto3,oneof"`
}

func (*ListMatcher_OneOf) isListMatcher_MatchPattern() {}

func (m *ListMatcher) GetMatchPattern() isListMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *ListMatcher) GetOneOf() *ValueMatcher {
	if x, ok := m.GetMatchPattern().(*ListMatcher_OneOf); ok {
		return x.OneOf
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListMatcher_OneOf)(nil),
	}
}

func init() {
	proto.RegisterType((*ValueMatcher)(nil), "envoy.type.matcher.ValueMatcher")
	proto.RegisterType((*ValueMatcher_NullMatch)(nil), "envoy.type.matcher.ValueMatcher.NullMatch")
	proto.RegisterType((*ListMatcher)(nil), "envoy.type.matcher.ListMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/value.proto", fileDescriptor_145b36501d266253) }

var fileDescriptor_145b36501d266253 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x57, 0xe7, 0x86, 0x7b, 0xdd, 0x2e, 0x41, 0x50, 0x76, 0x70, 0x75, 0x20, 0x0c, 0x0f,
	0xa9, 0xe8, 0xc9, 0x9b, 0x14, 0x91, 0x81, 0xbf, 0xc6, 0x04, 0x8f, 0x8e, 0xd6, 0xbd, 0x69, 0x21,
	0x4b, 0x4a, 0x9a, 0x0e, 0xf7, 0x9f, 0xf8, 0xa7, 0x7a, 0x94, 0xfc, 0xa8, 0x1b, 0x18, 0xf1, 0x96,
	0xbc, 0xf7, 0xf9, 0x7e, 0xfa, 0xd2, 0x04, 0x8e, 0x90, 0xaf, 0xc4, 0x3a, 0x56, 0xeb, 0x02, 0xe3,
	0x65, 0xaa, 0x5e, 0xdf, 0x51, 0xc6, 0xab, 0x94, 0x55, 0x48, 0x0b, 0x29, 0x94, 0x20, 0xc4, 0xf4,
	0xa9, 0xee, 0x53, 0xd7, 0xef, 0x0f, 0x3c, 0x19, 0x5e, 0x2d, 0x33, 0x94, 0x36, 0xe4, 0x05, 0x4a,
	0x25, 0x73, 0xfe, 0xe6, 0x80, 0x83, 0x55, 0xca, 0xf2, 0x79, 0xaa, 0x30, 0xae, 0x17, 0xb6, 0x31,
	0xfc, 0x6c, 0x42, 0xf7, 0x59, 0x7f, 0xfe, 0xde, 0xc6, 0xc8, 0x2d, 0x00, 0xaf, 0x18, 0x9b, 0x19,
	0xcd, 0x61, 0x10, 0x05, 0xa3, 0xf0, 0xfc, 0x94, 0xfe, 0x1e, 0x8a, 0x6e, 0xa7, 0xe8, 0x43, 0xc5,
	0x98, 0x59, 0x8f, 0x1b, 0xd3, 0x0e, 0xaf, 0x37, 0xe4, 0x06, 0xba, 0x73, 0x51, 0x65, 0x0c, 0x9d,
	0x6e, 0xc7, 0xe8, 0x8e, 0x7d, 0xba, 0x6b, 0xc3, 0x39, 0xdf, 0xb8, 0x31, 0x0d, 0xe7, 0x9b, 0x82,
	0xf6, 0xd8, 0xe3, 0x38, 0x4f, 0xf3, 0x6f, 0xcf, 0x93, 0xe1, 0xb6, 0x3c, 0xe5, 0xa6, 0x40, 0x06,
	0x00, 0x99, 0x10, 0xf5, 0xe1, 0x76, 0xa3, 0x60, 0xb4, 0xa7, 0x07, 0xd6, 0x35, 0x0b, 0x9c, 0x40,
	0xaf, 0x90, 0x58, 0x22, 0x57, 0x8e, 0x69, 0x39, 0xa6, 0xeb, 0xca, 0x16, 0xbb, 0x02, 0x60, 0x79,
	0x59, 0x33, 0x6d, 0x33, 0xcd, 0xc0, 0x37, 0xcd, 0x5d, 0x5e, 0xaa, 0xcd, 0x2c, 0x1d, 0x56, 0x6f,
	0xfb, 0x21, 0x74, 0x7e, 0xfe, 0x59, 0xb2, 0x0f, 0x3d, 0x13, 0x98, 0x15, 0xa9, 0x52, 0x28, 0x39,
	0x69, 0x7e, 0x25, 0xc1, 0xf0, 0x05, 0xc2, 0xad, 0x38, 0xb9, 0x84, 0xb6, 0xe0, 0x38, 0x13, 0x0b,
	0x77, 0x29, 0xd1, 0x7f, 0x97, 0x32, 0x6e, 0x4c, 0x5b, 0x82, 0xe3, 0xe3, 0xc2, 0xef, 0x4f, 0xce,
	0x20, 0xca, 0x85, 0x95, 0x14, 0x52, 0x7c, 0xac, 0x3d, 0xbe, 0x04, 0x8c, 0x70, 0xa2, 0x9f, 0xca,
	0x24, 0xc8, 0xda, 0xe6, 0xcd, 0x5c, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0xec, 0x38, 0x3f, 0x72,
	0xc4, 0x02, 0x00, 0x00,
}
