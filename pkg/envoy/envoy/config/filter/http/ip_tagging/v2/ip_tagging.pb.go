// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _ "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of requests the filter should apply to. The supported types
// are internal, external or both. The
// :ref:`x-forwarded-for<config_http_conn_man_headers_x-forwarded-for_internal_origin>` header is
// used to determine if a request is internal and will result in
// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>`
// being set. The filter defaults to both, and it will apply to all request types.
type IPTagging_RequestType int32

const (
	// Both external and internal requests will be tagged. This is the default value.
	IPTagging_BOTH IPTagging_RequestType = 0
	// Only internal requests will be tagged.
	IPTagging_INTERNAL IPTagging_RequestType = 1
	// Only external requests will be tagged.
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

var IPTagging_RequestType_name = map[int32]string{
	0: "BOTH",
	1: "INTERNAL",
	2: "EXTERNAL",
}
var IPTagging_RequestType_value = map[string]int32{
	"BOTH":     0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
}

func (x IPTagging_RequestType) String() string {
	return proto.EnumName(IPTagging_RequestType_name, int32(x))
}
func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ip_tagging_a13e9561239cbddc, []int{0, 0}
}

type IPTagging struct {
	// The type of request the filter should apply to.
	RequestType IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType" json:"request_type,omitempty"`
	// [#comment:TODO(ccaraman): Extend functionality to load IP tags from file system.
	// Tracked by issue https://github.com/envoyproxy/envoy/issues/2695]
	// The set of IP tags for the filter.
	IpTags               []*IPTagging_IPTag `protobuf:"bytes,4,rep,name=ip_tags,json=ipTags,proto3" json:"ip_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IPTagging) Reset()         { *m = IPTagging{} }
func (m *IPTagging) String() string { return proto.CompactTextString(m) }
func (*IPTagging) ProtoMessage()    {}
func (*IPTagging) Descriptor() ([]byte, []int) {
	return fileDescriptor_ip_tagging_a13e9561239cbddc, []int{0}
}
func (m *IPTagging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging.Unmarshal(m, b)
}
func (m *IPTagging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging.Marshal(b, m, deterministic)
}
func (dst *IPTagging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging.Merge(dst, src)
}
func (m *IPTagging) XXX_Size() int {
	return xxx_messageInfo_IPTagging.Size(m)
}
func (m *IPTagging) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging proto.InternalMessageInfo

func (m *IPTagging) GetRequestType() IPTagging_RequestType {
	if m != nil {
		return m.RequestType
	}
	return IPTagging_BOTH
}

func (m *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if m != nil {
		return m.IpTags
	}
	return nil
}

// Supplies the IP tag name and the IP address subnets.
type IPTagging_IPTag struct {
	// Specifies the IP tag name to apply.
	IpTagName string `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName,proto3" json:"ip_tag_name,omitempty"`
	// A list of IP address subnets that will be tagged with
	// ip_tag_name. Both IPv4 and IPv6 are supported.
	IpList               []*core.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList,proto3" json:"ip_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IPTagging_IPTag) Reset()         { *m = IPTagging_IPTag{} }
func (m *IPTagging_IPTag) String() string { return proto.CompactTextString(m) }
func (*IPTagging_IPTag) ProtoMessage()    {}
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ip_tagging_a13e9561239cbddc, []int{0, 0}
}
func (m *IPTagging_IPTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging_IPTag.Unmarshal(m, b)
}
func (m *IPTagging_IPTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging_IPTag.Marshal(b, m, deterministic)
}
func (dst *IPTagging_IPTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging_IPTag.Merge(dst, src)
}
func (m *IPTagging_IPTag) XXX_Size() int {
	return xxx_messageInfo_IPTagging_IPTag.Size(m)
}
func (m *IPTagging_IPTag) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging_IPTag.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging_IPTag proto.InternalMessageInfo

func (m *IPTagging_IPTag) GetIpTagName() string {
	if m != nil {
		return m.IpTagName
	}
	return ""
}

func (m *IPTagging_IPTag) GetIpList() []*core.CidrRange {
	if m != nil {
		return m.IpList
	}
	return nil
}

func init() {
	proto.RegisterType((*IPTagging)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging")
	proto.RegisterType((*IPTagging_IPTag)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging.IPTag")
	proto.RegisterEnum("envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType", IPTagging_RequestType_name, IPTagging_RequestType_value)
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto", fileDescriptor_ip_tagging_a13e9561239cbddc)
}

var fileDescriptor_ip_tagging_a13e9561239cbddc = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xcd, 0x4a, 0x2b, 0x31,
	0x14, 0xbe, 0x33, 0xfd, 0xb9, 0x6d, 0xa6, 0x5c, 0x4a, 0x36, 0xb7, 0x14, 0xa9, 0xa5, 0x0b, 0xe9,
	0x2a, 0x81, 0x29, 0xd2, 0x95, 0x0b, 0x2b, 0x05, 0x0b, 0xa5, 0xca, 0x30, 0x0b, 0x11, 0xb1, 0xa4,
	0x9d, 0x34, 0x06, 0xa6, 0x93, 0x98, 0xa4, 0x23, 0xdd, 0xfa, 0x08, 0x3e, 0x8e, 0x2b, 0x5f, 0x47,
	0xf0, 0x21, 0x64, 0x92, 0x56, 0xbb, 0xd4, 0xdd, 0xf9, 0xf8, 0xce, 0xf7, 0x73, 0x0e, 0x18, 0xd2,
	0x2c, 0x17, 0x5b, 0xbc, 0x14, 0xd9, 0x8a, 0x33, 0xbc, 0xe2, 0xa9, 0xa1, 0x0a, 0x3f, 0x18, 0x23,
	0x31, 0x97, 0x73, 0x43, 0x18, 0xe3, 0x19, 0xc3, 0x79, 0x78, 0x80, 0x90, 0x54, 0xc2, 0x08, 0x78,
	0x62, 0x85, 0xc8, 0x09, 0x91, 0x13, 0xa2, 0x42, 0x88, 0x0e, 0x56, 0xf3, 0xb0, 0x7d, 0xec, 0x02,
	0x88, 0xe4, 0x85, 0xcd, 0x52, 0x28, 0x8a, 0x49, 0x92, 0x28, 0xaa, 0xb5, 0x33, 0x6a, 0x77, 0x98,
	0x10, 0x2c, 0xa5, 0xd8, 0xa2, 0xc5, 0x66, 0x85, 0x9f, 0x14, 0x91, 0x92, 0xaa, 0x3d, 0xff, 0x3f,
	0x27, 0x29, 0x4f, 0x88, 0xa1, 0x78, 0x3f, 0x38, 0xa2, 0xf7, 0xe1, 0x83, 0xfa, 0xe4, 0x3a, 0x76,
	0x51, 0x30, 0x05, 0x0d, 0x45, 0x1f, 0x37, 0x54, 0x9b, 0xb9, 0xd9, 0x4a, 0xda, 0xf2, 0xba, 0x5e,
	0xff, 0x5f, 0x78, 0x86, 0x7e, 0x56, 0x13, 0x7d, 0x19, 0xa1, 0xc8, 0xb9, 0xc4, 0x5b, 0x49, 0x47,
	0xe0, 0xf5, 0xfd, 0xad, 0x54, 0x79, 0xf6, 0xfc, 0xa6, 0x17, 0x05, 0xea, 0x9b, 0x80, 0x77, 0xe0,
	0xaf, 0xd3, 0xeb, 0x56, 0xb9, 0x5b, 0xea, 0x07, 0xe1, 0xf0, 0xf7, 0x41, 0x76, 0xda, 0x45, 0xbc,
	0x78, 0x7e, 0xcd, 0x8b, 0xaa, 0x5c, 0xc6, 0x84, 0xe9, 0xf6, 0x3d, 0xa8, 0x58, 0x12, 0x76, 0x40,
	0xe0, 0xd4, 0xf3, 0x8c, 0xac, 0xdd, 0x4d, 0xf5, 0xa8, 0x6e, 0xb7, 0x66, 0x64, 0x4d, 0xe1, 0xa9,
	0xad, 0x91, 0x72, 0x6d, 0x5a, 0xbe, 0xad, 0x71, 0xb4, 0xab, 0x41, 0x24, 0x2f, 0xc2, 0x8a, 0x77,
	0xa3, 0x0b, 0x9e, 0xa8, 0x88, 0x64, 0x8c, 0x16, 0xfe, 0x53, 0xae, 0x4d, 0x6f, 0x00, 0x82, 0x83,
	0x2b, 0x61, 0x0d, 0x94, 0x47, 0x57, 0xf1, 0x65, 0xf3, 0x0f, 0x6c, 0x80, 0xda, 0x64, 0x16, 0x8f,
	0xa3, 0xd9, 0xf9, 0xb4, 0xe9, 0x15, 0x68, 0x7c, 0xb3, 0x43, 0xfe, 0xa8, 0x7c, 0xeb, 0xe7, 0xe1,
	0xa2, 0x6a, 0x7f, 0x3f, 0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x57, 0x2f, 0x1a, 0x38, 0x02,
	0x00, 0x00,
}
