// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/metadata.proto

package matcher

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
// <envoy_api_msg_config.rbac.v2alpha.Permission>` and :ref:`Principal
// <envoy_api_msg_config.rbac.v2alpha.Principal>`.
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

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetadataMatcher_PathSegment) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetadataMatcher_PathSegment_OneofMarshaler, _MetadataMatcher_PathSegment_OneofUnmarshaler, _MetadataMatcher_PathSegment_OneofSizer, []interface{}{
		(*MetadataMatcher_PathSegment_Key)(nil),
	}
}

func _MetadataMatcher_PathSegment_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetadataMatcher_PathSegment)
	// segment
	switch x := m.Segment.(type) {
	case *MetadataMatcher_PathSegment_Key:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Key)
	case nil:
	default:
		return fmt.Errorf("MetadataMatcher_PathSegment.Segment has unexpected type %T", x)
	}
	return nil
}

func _MetadataMatcher_PathSegment_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetadataMatcher_PathSegment)
	switch tag {
	case 1: // segment.key
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Segment = &MetadataMatcher_PathSegment_Key{x}
		return true, err
	default:
		return false, nil
	}
}

func _MetadataMatcher_PathSegment_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetadataMatcher_PathSegment)
	// segment
	switch x := m.Segment.(type) {
	case *MetadataMatcher_PathSegment_Key:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Key)))
		n += len(x.Key)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*MetadataMatcher)(nil), "envoy.type.matcher.MetadataMatcher")
	proto.RegisterType((*MetadataMatcher_PathSegment)(nil), "envoy.type.matcher.MetadataMatcher.PathSegment")
}

func init() { proto.RegisterFile("envoy/type/matcher/metadata.proto", fileDescriptor_865eaf6a1e9e266d) }

var fileDescriptor_865eaf6a1e9e266d = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0xcf,
	0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02,
	0x2b, 0xd1, 0x03, 0x29, 0xd1, 0x83, 0x2a, 0x91, 0x92, 0xc3, 0xa2, 0xad, 0x2c, 0x31, 0xa7, 0x34,
	0x15, 0xa2, 0x47, 0x4a, 0xbc, 0x2c, 0x31, 0x27, 0x33, 0x25, 0xb1, 0x24, 0x55, 0x1f, 0xc6, 0x80,
	0x48, 0x28, 0x75, 0x32, 0x71, 0xf1, 0xfb, 0x42, 0xcd, 0xf7, 0x85, 0x68, 0x14, 0x52, 0xe4, 0x62,
	0x4b, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0xe2, 0xdc, 0xf5,
	0xf2, 0x00, 0x33, 0x4b, 0x11, 0x93, 0x02, 0x63, 0x10, 0x54, 0x42, 0xc8, 0x9f, 0x8b, 0xa5, 0x20,
	0xb1, 0x24, 0x43, 0x82, 0x49, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x5f, 0x0f, 0xd3, 0x49, 0x7a, 0x68,
	0xa6, 0xea, 0x05, 0x24, 0x96, 0x64, 0x04, 0xa7, 0xa6, 0xe7, 0xa6, 0xe6, 0x95, 0x38, 0x71, 0x81,
	0x4c, 0x64, 0x9d, 0xc4, 0xc8, 0xc4, 0xc1, 0x18, 0x04, 0x36, 0x48, 0xc8, 0x89, 0x8b, 0x15, 0xec,
	0x5e, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x05, 0x6c, 0x26, 0x86, 0x81, 0x14, 0x40, 0x8d,
	0x83, 0x1a, 0xd1, 0xc5, 0xc8, 0x24, 0xc0, 0x18, 0x04, 0xd1, 0x2a, 0x65, 0xc7, 0xc5, 0x8d, 0x64,
	0x89, 0x90, 0x2c, 0x17, 0x73, 0x76, 0x6a, 0x25, 0x86, 0x1f, 0x3c, 0x18, 0x82, 0x40, 0xe2, 0x4e,
	0x02, 0x5c, 0xec, 0xc5, 0x50, 0x95, 0xac, 0x3b, 0x5e, 0x1e, 0x60, 0x66, 0x74, 0xe2, 0x8c, 0x62,
	0x87, 0x5a, 0x95, 0xc4, 0x06, 0x0e, 0x1d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x6e,
	0x90, 0x30, 0x8f, 0x01, 0x00, 0x00,
}
