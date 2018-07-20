// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/accesslog.proto

/*
Package cilium is a generated protocol buffer package.

It is generated from these files:
	cilium/accesslog.proto
	cilium/cilium_bpf_metadata.proto
	cilium/cilium_l7policy.proto
	cilium/npds.proto
	cilium/nphds.proto

It has these top-level messages:
	KeyValue
	HttpLogEntry
	BpfMetadata
	L7Policy
	NetworkPolicy
	PortNetworkPolicy
	PortNetworkPolicyRule
	HttpNetworkPolicyRules
	HttpNetworkPolicyRule
	KafkaNetworkPolicyRules
	KafkaNetworkPolicyRule
	NetworkPolicyHosts
*/
package cilium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Protocol int32

const (
	Protocol_HTTP10 Protocol = 0
	Protocol_HTTP11 Protocol = 1
	Protocol_HTTP2  Protocol = 2
)

var Protocol_name = map[int32]string{
	0: "HTTP10",
	1: "HTTP11",
	2: "HTTP2",
}
var Protocol_value = map[string]int32{
	"HTTP10": 0,
	"HTTP11": 1,
	"HTTP2":  2,
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type EntryType int32

const (
	EntryType_Request  EntryType = 0
	EntryType_Response EntryType = 1
	EntryType_Denied   EntryType = 2
)

var EntryType_name = map[int32]string{
	0: "Request",
	1: "Response",
	2: "Denied",
}
var EntryType_value = map[string]int32{
	"Request":  0,
	"Response": 1,
	"Denied":   2,
}

func (x EntryType) String() string {
	return proto.EnumName(EntryType_name, int32(x))
}
func (EntryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HttpLogEntry struct {
	// The time that Cilium filter captured this log entry,
	// in, nanoseconds since 1/1/1970.
	Timestamp    uint64    `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	HttpProtocol Protocol  `protobuf:"varint,2,opt,name=http_protocol,json=httpProtocol,enum=cilium.Protocol" json:"http_protocol,omitempty"`
	EntryType    EntryType `protobuf:"varint,3,opt,name=entry_type,json=entryType,enum=cilium.EntryType" json:"entry_type,omitempty"`
	// Cilium network policy resource name
	PolicyName string `protobuf:"bytes,4,opt,name=policy_name,json=policyName" json:"policy_name,omitempty"`
	// Cilium rule reference
	CiliumRuleRef string `protobuf:"bytes,5,opt,name=cilium_rule_ref,json=ciliumRuleRef" json:"cilium_rule_ref,omitempty"`
	// Cilium security ID of the source
	SourceSecurityId uint32 `protobuf:"varint,6,opt,name=source_security_id,json=sourceSecurityId" json:"source_security_id,omitempty"`
	// These fields record the original source and destination addresses,
	// stored in ipv4:port or [ipv6]:port format.
	SourceAddress      string `protobuf:"bytes,7,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	DestinationAddress string `protobuf:"bytes,8,opt,name=destination_address,json=destinationAddress" json:"destination_address,omitempty"`
	// Request info that is also retained for the response
	Scheme string `protobuf:"bytes,9,opt,name=scheme" json:"scheme,omitempty"`
	Host   string `protobuf:"bytes,10,opt,name=host" json:"host,omitempty"`
	Path   string `protobuf:"bytes,11,opt,name=path" json:"path,omitempty"`
	Method string `protobuf:"bytes,12,opt,name=method" json:"method,omitempty"`
	// Response info
	Status uint32 `protobuf:"varint,13,opt,name=status" json:"status,omitempty"`
	// Request headers not included above
	Headers []*KeyValue `protobuf:"bytes,14,rep,name=headers" json:"headers,omitempty"`
	// 'true' if the request was received by an ingress listener,
	// 'false' if received by an egress listener
	IsIngress bool `protobuf:"varint,15,opt,name=is_ingress,json=isIngress" json:"is_ingress,omitempty"`
}

func (m *HttpLogEntry) Reset()                    { *m = HttpLogEntry{} }
func (m *HttpLogEntry) String() string            { return proto.CompactTextString(m) }
func (*HttpLogEntry) ProtoMessage()               {}
func (*HttpLogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HttpLogEntry) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *HttpLogEntry) GetHttpProtocol() Protocol {
	if m != nil {
		return m.HttpProtocol
	}
	return Protocol_HTTP10
}

func (m *HttpLogEntry) GetEntryType() EntryType {
	if m != nil {
		return m.EntryType
	}
	return EntryType_Request
}

func (m *HttpLogEntry) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *HttpLogEntry) GetCiliumRuleRef() string {
	if m != nil {
		return m.CiliumRuleRef
	}
	return ""
}

func (m *HttpLogEntry) GetSourceSecurityId() uint32 {
	if m != nil {
		return m.SourceSecurityId
	}
	return 0
}

func (m *HttpLogEntry) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *HttpLogEntry) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *HttpLogEntry) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *HttpLogEntry) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HttpLogEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HttpLogEntry) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpLogEntry) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *HttpLogEntry) GetHeaders() []*KeyValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpLogEntry) GetIsIngress() bool {
	if m != nil {
		return m.IsIngress
	}
	return false
}

func init() {
	proto.RegisterType((*KeyValue)(nil), "cilium.KeyValue")
	proto.RegisterType((*HttpLogEntry)(nil), "cilium.HttpLogEntry")
	proto.RegisterEnum("cilium.Protocol", Protocol_name, Protocol_value)
	proto.RegisterEnum("cilium.EntryType", EntryType_name, EntryType_value)
}

func init() { proto.RegisterFile("cilium/accesslog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x6f, 0x6b, 0xd4, 0x40,
	0x10, 0xc6, 0x9b, 0xfb, 0x93, 0x4b, 0xe6, 0xfe, 0x34, 0x8e, 0x52, 0xf6, 0x85, 0xe2, 0x51, 0x50,
	0x8e, 0x43, 0xaf, 0xed, 0x89, 0x1f, 0x40, 0x50, 0x68, 0x51, 0xa4, 0xac, 0x87, 0x6f, 0xc3, 0x9a,
	0x4c, 0x2f, 0x8b, 0x49, 0x36, 0x66, 0x37, 0x42, 0x3e, 0x8d, 0x5f, 0x55, 0xb2, 0x9b, 0x5c, 0xfb,
	0xee, 0x99, 0xdf, 0x3c, 0xcf, 0xec, 0x2e, 0xb3, 0x70, 0x91, 0xc8, 0x5c, 0x36, 0xc5, 0x95, 0x48,
	0x12, 0xd2, 0x3a, 0x57, 0xc7, 0x5d, 0x55, 0x2b, 0xa3, 0xd0, 0x77, 0xfc, 0x72, 0x0f, 0xc1, 0x57,
	0x6a, 0x7f, 0x8a, 0xbc, 0x21, 0x8c, 0x60, 0xfc, 0x9b, 0x5a, 0xe6, 0xad, 0xbd, 0x4d, 0xc8, 0x3b,
	0x89, 0x2f, 0x60, 0xfa, 0xb7, 0x6b, 0xb1, 0x91, 0x65, 0xae, 0xb8, 0xfc, 0x37, 0x81, 0xc5, 0xad,
	0x31, 0xd5, 0x37, 0x75, 0xfc, 0x52, 0x9a, 0xba, 0xc5, 0x97, 0x10, 0x1a, 0x59, 0x90, 0x36, 0xa2,
	0xa8, 0x6c, 0x7c, 0xc2, 0x1f, 0x01, 0x7e, 0x84, 0x65, 0x66, 0x4c, 0x15, 0xdb, 0x83, 0x13, 0x95,
	0xdb, 0x61, 0xab, 0x7d, 0xb4, 0x73, 0x57, 0xd8, 0xdd, 0xf7, 0x9c, 0x2f, 0x3a, 0xdb, 0x50, 0xe1,
	0x35, 0x00, 0x75, 0xd3, 0x63, 0xd3, 0x56, 0xc4, 0xc6, 0x36, 0xf3, 0x6c, 0xc8, 0xd8, 0x73, 0x0f,
	0x6d, 0x45, 0x3c, 0xa4, 0x41, 0xe2, 0x6b, 0x98, 0x57, 0x2a, 0x97, 0x49, 0x1b, 0x97, 0xa2, 0x20,
	0x36, 0xb1, 0x77, 0x06, 0x87, 0xbe, 0x8b, 0x82, 0xf0, 0x2d, 0x9c, 0xbb, 0x7c, 0x5c, 0x37, 0x39,
	0xc5, 0x35, 0x3d, 0xb0, 0xa9, 0x35, 0x2d, 0x1d, 0xe6, 0x4d, 0x4e, 0x9c, 0x1e, 0xf0, 0x1d, 0xa0,
	0x56, 0x4d, 0x9d, 0x50, 0xac, 0x29, 0x69, 0x6a, 0x69, 0xda, 0x58, 0xa6, 0xcc, 0x5f, 0x7b, 0x9b,
	0x25, 0x8f, 0x5c, 0xe7, 0x47, 0xdf, 0xb8, 0x4b, 0xf1, 0x0d, 0xac, 0x7a, 0xb7, 0x48, 0xd3, 0x9a,
	0xb4, 0x66, 0x33, 0x37, 0xd4, 0xd1, 0x4f, 0x0e, 0xe2, 0x15, 0x3c, 0x4f, 0x49, 0x1b, 0x59, 0x0a,
	0x23, 0x55, 0x79, 0xf2, 0x06, 0xd6, 0x8b, 0x4f, 0x5a, 0x43, 0xe0, 0x02, 0x7c, 0x9d, 0x64, 0x54,
	0x10, 0x0b, 0xad, 0xa7, 0xaf, 0x10, 0x61, 0x92, 0x29, 0x6d, 0x18, 0x58, 0x6a, 0x75, 0xc7, 0x2a,
	0x61, 0x32, 0x36, 0x77, 0xac, 0xd3, 0x5d, 0xbe, 0x20, 0x93, 0xa9, 0x94, 0x2d, 0x5c, 0xde, 0x55,
	0x76, 0xae, 0x11, 0xa6, 0xd1, 0x6c, 0x69, 0x5f, 0xd4, 0x57, 0xb8, 0x85, 0x59, 0x46, 0x22, 0xa5,
	0x5a, 0xb3, 0xd5, 0x7a, 0xbc, 0x99, 0x3f, 0x6e, 0x68, 0xf8, 0x21, 0x7c, 0x30, 0xe0, 0x2b, 0x00,
	0xa9, 0x63, 0x59, 0x1e, 0xed, 0x1b, 0xce, 0xd7, 0xde, 0x26, 0xe0, 0xa1, 0xd4, 0x77, 0x0e, 0x6c,
	0xdf, 0x43, 0x70, 0xda, 0x23, 0x80, 0x7f, 0x7b, 0x38, 0xdc, 0xdf, 0x5c, 0x47, 0x67, 0x27, 0x7d,
	0x13, 0x79, 0x18, 0xc2, 0xb4, 0xd3, 0xfb, 0x68, 0xb4, 0xdd, 0x43, 0x78, 0x5a, 0x28, 0xce, 0x61,
	0xc6, 0xe9, 0x4f, 0x43, 0xda, 0x44, 0x67, 0xb8, 0x80, 0x80, 0x93, 0xae, 0x54, 0xa9, 0x29, 0xf2,
	0xba, 0xf8, 0x67, 0x2a, 0x25, 0xa5, 0xd1, 0xe8, 0x97, 0x6f, 0xbf, 0xd3, 0x87, 0xff, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xcc, 0xc0, 0x49, 0x3a, 0xe1, 0x02, 0x00, 0x00,
}
