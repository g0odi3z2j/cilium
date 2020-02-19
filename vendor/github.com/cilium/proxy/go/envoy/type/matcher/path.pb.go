// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/path.proto

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

// Specifies the way to match a path on HTTP request.
type PathMatcher struct {
	// Types that are valid to be assigned to Rule:
	//	*PathMatcher_Path
	Rule                 isPathMatcher_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PathMatcher) Reset()         { *m = PathMatcher{} }
func (m *PathMatcher) String() string { return proto.CompactTextString(m) }
func (*PathMatcher) ProtoMessage()    {}
func (*PathMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_bec7ed88adc90b4e, []int{0}
}

func (m *PathMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathMatcher.Unmarshal(m, b)
}
func (m *PathMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathMatcher.Marshal(b, m, deterministic)
}
func (m *PathMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathMatcher.Merge(m, src)
}
func (m *PathMatcher) XXX_Size() int {
	return xxx_messageInfo_PathMatcher.Size(m)
}
func (m *PathMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_PathMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_PathMatcher proto.InternalMessageInfo

type isPathMatcher_Rule interface {
	isPathMatcher_Rule()
}

type PathMatcher_Path struct {
	Path *StringMatcher `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

func (*PathMatcher_Path) isPathMatcher_Rule() {}

func (m *PathMatcher) GetRule() isPathMatcher_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *PathMatcher) GetPath() *StringMatcher {
	if x, ok := m.GetRule().(*PathMatcher_Path); ok {
		return x.Path
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PathMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PathMatcher_Path)(nil),
	}
}

func init() {
	proto.RegisterType((*PathMatcher)(nil), "envoy.type.matcher.PathMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/path.proto", fileDescriptor_bec7ed88adc90b4e) }

var fileDescriptor_bec7ed88adc90b4e = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2f,
	0x48, 0x2c, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0x4b, 0xeb, 0x81, 0xa4,
	0xf5, 0xa0, 0xd2, 0x52, 0xf2, 0x58, 0xb4, 0x14, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x43, 0x34, 0x49,
	0x89, 0x97, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xa4, 0xea, 0xc3, 0x18, 0x10, 0x09, 0xa5, 0x58,
	0x2e, 0xee, 0x80, 0xc4, 0x92, 0x0c, 0x5f, 0x88, 0x26, 0x21, 0x47, 0x2e, 0x16, 0x90, 0x55, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x8a, 0x7a, 0x98, 0x76, 0xe9, 0x05, 0x83, 0xcd, 0x85, 0x6a,
	0x70, 0xe2, 0xf8, 0xe5, 0xc4, 0xda, 0xc5, 0xc8, 0x24, 0xc0, 0xe8, 0xc1, 0x10, 0x04, 0xd6, 0xea,
	0xc4, 0xcd, 0xc5, 0x52, 0x54, 0x9a, 0x93, 0x2a, 0xc4, 0xfc, 0xc3, 0x89, 0xd1, 0x49, 0x9f, 0x4b,
	0x21, 0x33, 0x1f, 0x62, 0x4a, 0x41, 0x51, 0x7e, 0x45, 0x25, 0x16, 0x03, 0x9d, 0x38, 0x41, 0x0e,
	0x08, 0x00, 0xb9, 0x26, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0x1b, 0xd1, 0xbf, 0x05, 0x01, 0x00, 0x00,
}
