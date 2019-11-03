// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/metadata.proto

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

// MetadataMatcher provides a general interface to check if a given value is matched in
// :ref:`Metadata <envoy_api_msg_core.Metadata>`. It uses `filter` and `path` to retrieve the value
// from the Metadata and then check if it's matched to the specified value.
//
// For example, for the following Metadata:
//
// .. code-block:: yaml
//
//    filter_metadata:
//      envoy.filters.http.rbac:
//        fields:
//          a:
//            struct_value:
//              fields:
//                b:
//                  struct_value:
//                    fields:
//                      c:
//                        string_value: pro
//                t:
//                  list_value:
//                    values:
//                      - string_value: m
//                      - string_value: n
//
// The following MetadataMatcher is matched as the path [a, b, c] will retrieve a string value "pro"
// from the Metadata which is matched to the specified prefix match.
//
// .. code-block:: yaml
//
//    filter: envoy.filters.http.rbac
//    path:
//    - key: a
//    - key: b
//    - key: c
//    value:
//      string_match:
//        prefix: pr
//
// The following MetadataMatcher is matched as the code will match one of the string values in the
// list at the path [a, t].
//
// .. code-block:: yaml
//
//    filter: envoy.filters.http.rbac
//    path:
//    - key: a
//    - key: t
//    value:
//      list_match:
//        one_of:
//          string_match:
//            exact: m
//
// An example use of MetadataMatcher is specifying additional metadata in envoy.filters.http.rbac to
// enforce access control based on dynamic metadata in a request. See :ref:`Permission
// <envoy_api_msg_config.rbac.v2.Permission>` and :ref:`Principal
// <envoy_api_msg_config.rbac.v2.Principal>`.
type MetadataMatcher struct {
	// The filter name to retrieve the Struct from the Metadata.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The path to retrieve the Value from the Struct.
	Path []*MetadataMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	// The MetadataMatcher is matched if the value retrieved by path is matched to this value.
	Value                *ValueMatcher `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MetadataMatcher) Reset()         { *m = MetadataMatcher{} }
func (m *MetadataMatcher) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher) ProtoMessage()    {}
func (*MetadataMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_865eaf6a1e9e266d, []int{0}
}

func (m *MetadataMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataMatcher.Unmarshal(m, b)
}
func (m *MetadataMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataMatcher.Marshal(b, m, deterministic)
}
func (m *MetadataMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher.Merge(m, src)
}
func (m *MetadataMatcher) XXX_Size() int {
	return xxx_messageInfo_MetadataMatcher.Size(m)
}
func (m *MetadataMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher proto.InternalMessageInfo

func (m *MetadataMatcher) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *MetadataMatcher) GetPath() []*MetadataMatcher_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *MetadataMatcher) GetValue() *ValueMatcher {
	if m != nil {
		return m.Value
	}
	return nil
}

// Specifies the segment in a path to retrieve value from Metadata.
// Note: Currently it's not supported to retrieve a value from a list in Metadata. This means that
// if the segment key refers to a list, it has to be the last segment in a path.
type MetadataMatcher_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*MetadataMatcher_PathSegment_Key
	Segment              isMetadataMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *MetadataMatcher_PathSegment) Reset()         { *m = MetadataMatcher_PathSegment{} }
func (m *MetadataMatcher_PathSegment) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher_PathSegment) ProtoMessage()    {}
func (*MetadataMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_865eaf6a1e9e266d, []int{0, 0}
}

func (m *MetadataMatcher_PathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Unmarshal(m, b)
}
func (m *MetadataMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Marshal(b, m, deterministic)
}
func (m *MetadataMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher_PathSegment.Merge(m, src)
}
func (m *MetadataMatcher_PathSegment) XXX_Size() int {
	return xxx_messageInfo_MetadataMatcher_PathSegment.Size(m)
}
func (m *MetadataMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher_PathSegment proto.InternalMessageInfo

type isMetadataMatcher_PathSegment_Segment interface {
	isMetadataMatcher_PathSegment_Segment()
}

type MetadataMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*MetadataMatcher_PathSegment_Key) isMetadataMatcher_PathSegment_Segment() {}

func (m *MetadataMatcher_PathSegment) GetSegment() isMetadataMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *MetadataMatcher_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*MetadataMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetadataMatcher_PathSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetadataMatcher_PathSegment_Key)(nil),
	}
}

func init() {
	proto.RegisterType((*MetadataMatcher)(nil), "envoy.type.matcher.MetadataMatcher")
	proto.RegisterType((*MetadataMatcher_PathSegment)(nil), "envoy.type.matcher.MetadataMatcher.PathSegment")
}

func init() { proto.RegisterFile("envoy/type/matcher/metadata.proto", fileDescriptor_865eaf6a1e9e266d) }

var fileDescriptor_865eaf6a1e9e266d = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4f, 0x84, 0x40,
	0x10, 0xc5, 0x1d, 0xb8, 0x7f, 0x2e, 0xf1, 0x4f, 0xb6, 0x91, 0x60, 0xa2, 0xab, 0x15, 0xd5, 0x92,
	0xdc, 0x75, 0x56, 0x66, 0x2b, 0x1b, 0x12, 0x82, 0x89, 0xfd, 0xea, 0x8d, 0x42, 0x84, 0x83, 0xe0,
	0x48, 0xe4, 0x2b, 0x58, 0xfa, 0x69, 0xcd, 0x35, 0x1a, 0xd8, 0xbd, 0xc4, 0x9c, 0x74, 0x9b, 0xd9,
	0xf7, 0x7e, 0xef, 0xcd, 0xb0, 0x2b, 0xdc, 0xb4, 0x55, 0x17, 0x51, 0x57, 0x63, 0x54, 0x6a, 0x7a,
	0xca, 0xb0, 0x89, 0x4a, 0x24, 0xbd, 0xd6, 0xa4, 0x65, 0xdd, 0x54, 0x54, 0x71, 0x3e, 0x48, 0x64,
	0x2f, 0x91, 0x56, 0x12, 0x5c, 0x8c, 0xd8, 0x5a, 0x5d, 0xbc, 0xa3, 0xf1, 0x04, 0x67, 0xad, 0x2e,
	0xf2, 0xb5, 0x26, 0x8c, 0x76, 0x0f, 0xf3, 0x71, 0xfd, 0x03, 0xec, 0x24, 0xb6, 0xfc, 0xd8, 0x18,
	0xf9, 0x25, 0x9b, 0x3d, 0xe7, 0x05, 0x61, 0xe3, 0x83, 0x80, 0xf0, 0x50, 0xcd, 0xb7, 0x6a, 0xd2,
	0x38, 0x02, 0x52, 0x3b, 0xe6, 0x31, 0x9b, 0xd4, 0x9a, 0x32, 0xdf, 0x11, 0x6e, 0xe8, 0x2d, 0x23,
	0xf9, 0xbf, 0x90, 0xdc, 0x63, 0xca, 0x44, 0x53, 0x76, 0x8f, 0x2f, 0x25, 0x6e, 0x48, 0x2d, 0xb6,
	0x6a, 0xfa, 0x05, 0xce, 0x02, 0xd2, 0x01, 0xc3, 0x6f, 0xd9, 0x74, 0xe8, 0xea, 0xbb, 0x02, 0x42,
	0x6f, 0x29, 0xc6, 0x78, 0x0f, 0xbd, 0xc0, 0xc2, 0x06, 0xc0, 0x27, 0x38, 0xa7, 0x90, 0x1a, 0x63,
	0x70, 0xc3, 0xbc, 0x3f, 0x01, 0xfc, 0x9c, 0xb9, 0xaf, 0xd8, 0xed, 0xb5, 0xbf, 0x3b, 0x48, 0xfb,
	0xa9, 0x3a, 0x66, 0xf3, 0x37, 0xab, 0x73, 0xbf, 0x15, 0xa8, 0x15, 0x13, 0x79, 0x65, 0x22, 0xeb,
	0xa6, 0xfa, 0xe8, 0x46, 0xd2, 0xd5, 0xd1, 0x6e, 0x9d, 0xa4, 0x3f, 0x5a, 0x02, 0x8f, 0xb3, 0xe1,
	0x7a, 0xab, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x42, 0x4b, 0x24, 0xaf, 0x01, 0x00, 0x00,
}
