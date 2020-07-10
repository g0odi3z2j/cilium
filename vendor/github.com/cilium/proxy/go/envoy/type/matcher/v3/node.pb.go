// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/node.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

// Specifies the way to match a Node.
// The match follows AND semantics.
type NodeMatcher struct {
	// Specifies match criteria on the node id.
	NodeId *StringMatcher `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Specifies match criteria on the node metadata.
	NodeMetadatas        []*StructMatcher `protobuf:"bytes,2,rep,name=node_metadatas,json=nodeMetadatas,proto3" json:"node_metadatas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeMatcher) Reset()         { *m = NodeMatcher{} }
func (m *NodeMatcher) String() string { return proto.CompactTextString(m) }
func (*NodeMatcher) ProtoMessage()    {}
func (*NodeMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccc0c3eef80eb4fc, []int{0}
}

func (m *NodeMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeMatcher.Unmarshal(m, b)
}
func (m *NodeMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeMatcher.Marshal(b, m, deterministic)
}
func (m *NodeMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeMatcher.Merge(m, src)
}
func (m *NodeMatcher) XXX_Size() int {
	return xxx_messageInfo_NodeMatcher.Size(m)
}
func (m *NodeMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_NodeMatcher proto.InternalMessageInfo

func (m *NodeMatcher) GetNodeId() *StringMatcher {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *NodeMatcher) GetNodeMetadatas() []*StructMatcher {
	if m != nil {
		return m.NodeMetadatas
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeMatcher)(nil), "envoy.type.matcher.v3.NodeMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/node.proto", fileDescriptor_ccc0c3eef80eb4fc) }

var fileDescriptor_ccc0c3eef80eb4fc = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2f,
	0x33, 0xd6, 0xcf, 0xcb, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x05, 0xab,
	0xd0, 0x03, 0xa9, 0xd0, 0x83, 0xaa, 0xd0, 0x2b, 0x33, 0x96, 0x52, 0xc2, 0xae, 0xb1, 0xb8, 0xa4,
	0x28, 0x33, 0x2f, 0x1d, 0xa2, 0x15, 0x8f, 0x9a, 0xd2, 0xe4, 0x12, 0xa8, 0x1a, 0xd9, 0xd2, 0x94,
	0x82, 0x44, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0xfd, 0xe2,
	0x92, 0xc4, 0x92, 0xd2, 0x62, 0xa8, 0xb4, 0x22, 0x86, 0x74, 0x59, 0x6a, 0x51, 0x71, 0x66, 0x7e,
	0x1e, 0xdc, 0x16, 0xa5, 0x03, 0x8c, 0x5c, 0xdc, 0x7e, 0xf9, 0x29, 0xa9, 0xbe, 0x10, 0x1b, 0x84,
	0x6c, 0xb9, 0xd8, 0x41, 0xce, 0x8f, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52,
	0xd1, 0xc3, 0xea, 0x05, 0xbd, 0x60, 0xb0, 0x5b, 0xa1, 0xda, 0x82, 0xd8, 0x40, 0x9a, 0x3c, 0x53,
	0x84, 0xbc, 0xb9, 0xf8, 0xc0, 0xda, 0x73, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x8b, 0x25,
	0x98, 0x14, 0x98, 0xf1, 0x9b, 0x52, 0x9a, 0x5c, 0x02, 0x33, 0x85, 0x17, 0xa4, 0xd7, 0x17, 0xa6,
	0xd5, 0x4a, 0x75, 0xd6, 0xd1, 0x0e, 0x39, 0x05, 0x2e, 0x39, 0x2c, 0x5a, 0x91, 0x9c, 0xec, 0x64,
	0xb5, 0xab, 0xe1, 0xc4, 0x45, 0x36, 0x26, 0x01, 0x26, 0x2e, 0xe5, 0xcc, 0x7c, 0x88, 0x3d, 0x05,
	0x45, 0xf9, 0x15, 0x95, 0xd8, 0xad, 0x74, 0xe2, 0x04, 0xe9, 0x0d, 0x00, 0x79, 0x3e, 0x80, 0x31,
	0x89, 0x0d, 0x1c, 0x0a, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x5e, 0x40, 0xe7, 0xca,
	0x01, 0x00, 0x00,
}
