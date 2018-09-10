// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/cilium_bpf_metadata.proto

package cilium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BpfMetadata struct {
	// File system root for bpf. Defaults to "/sys/fs/bpf" if left empty.
	BpfRoot string `protobuf:"bytes,1,opt,name=bpf_root,json=bpfRoot" json:"bpf_root,omitempty"`
	// 'true' if the filter is on ingress listener, 'false' for egress listener.
	IsIngress bool `protobuf:"varint,2,opt,name=is_ingress,json=isIngress" json:"is_ingress,omitempty"`
}

func (m *BpfMetadata) Reset()                    { *m = BpfMetadata{} }
func (m *BpfMetadata) String() string            { return proto.CompactTextString(m) }
func (*BpfMetadata) ProtoMessage()               {}
func (*BpfMetadata) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *BpfMetadata) GetBpfRoot() string {
	if m != nil {
		return m.BpfRoot
	}
	return ""
}

func (m *BpfMetadata) GetIsIngress() bool {
	if m != nil {
		return m.IsIngress
	}
	return false
}

func init() {
	proto.RegisterType((*BpfMetadata)(nil), "cilium.BpfMetadata")
}

func init() { proto.RegisterFile("cilium/cilium_bpf_metadata.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0xcc, 0xc9,
	0x2c, 0xcd, 0xd5, 0x87, 0x50, 0xf1, 0x49, 0x05, 0x69, 0xf1, 0xb9, 0xa9, 0x25, 0x89, 0x29, 0x89,
	0x25, 0x89, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x29, 0x25, 0x77, 0x2e, 0x6e,
	0xa7, 0x82, 0x34, 0x5f, 0xa8, 0xa4, 0x90, 0x24, 0x17, 0x07, 0x48, 0x71, 0x51, 0x7e, 0x7e, 0x89,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x7b, 0x52, 0x41, 0x5a, 0x50, 0x7e, 0x7e, 0x89, 0x90,
	0x2c, 0x17, 0x57, 0x66, 0x71, 0x7c, 0x66, 0x5e, 0x7a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x47, 0x10, 0x67, 0x66, 0xb1, 0x27, 0x44, 0x20, 0x89, 0x0d, 0x6c, 0xae, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0xb0, 0xea, 0x7e, 0xbe, 0x7b, 0x00, 0x00, 0x00,
}
