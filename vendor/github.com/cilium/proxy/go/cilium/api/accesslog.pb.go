// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/api/accesslog.proto

package cilium

import (
	fmt "fmt"
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

type HttpProtocol int32

const (
	HttpProtocol_HTTP10 HttpProtocol = 0
	HttpProtocol_HTTP11 HttpProtocol = 1
	HttpProtocol_HTTP2  HttpProtocol = 2
)

var HttpProtocol_name = map[int32]string{
	0: "HTTP10",
	1: "HTTP11",
	2: "HTTP2",
}

var HttpProtocol_value = map[string]int32{
	"HTTP10": 0,
	"HTTP11": 1,
	"HTTP2":  2,
}

func (x HttpProtocol) String() string {
	return proto.EnumName(HttpProtocol_name, int32(x))
}

func (HttpProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_02c2856a892fe6b6, []int{0}
}

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

func (EntryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_02c2856a892fe6b6, []int{1}
}

type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_02c2856a892fe6b6, []int{0}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

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
	HttpProtocol HttpProtocol `protobuf:"varint,1,opt,name=http_protocol,json=httpProtocol,proto3,enum=cilium.HttpProtocol" json:"http_protocol,omitempty"`
	// Request info that is also retained for the response
	Scheme string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Host   string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Path   string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Method string `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	// Request headers not included above
	Headers []*KeyValue `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty"`
	// Response info
	Status uint32 `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	// missing_headers includes both headers that were added to the
	// request, and headers that were merely logged as missing
	MissingHeaders []*KeyValue `protobuf:"bytes,8,rep,name=missing_headers,json=missingHeaders,proto3" json:"missing_headers,omitempty"`
	// rejected_headers includes headers that were flagged as unallowed,
	// which may have been removed, or merely logged and the request still
	// allowed, or the request may have been dropped due to them.
	RejectedHeaders      []*KeyValue `protobuf:"bytes,9,rep,name=rejected_headers,json=rejectedHeaders,proto3" json:"rejected_headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HttpLogEntry) Reset()         { *m = HttpLogEntry{} }
func (m *HttpLogEntry) String() string { return proto.CompactTextString(m) }
func (*HttpLogEntry) ProtoMessage()    {}
func (*HttpLogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_02c2856a892fe6b6, []int{1}
}

func (m *HttpLogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpLogEntry.Unmarshal(m, b)
}
func (m *HttpLogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpLogEntry.Marshal(b, m, deterministic)
}
func (m *HttpLogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpLogEntry.Merge(m, src)
}
func (m *HttpLogEntry) XXX_Size() int {
	return xxx_messageInfo_HttpLogEntry.Size(m)
}
func (m *HttpLogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpLogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_HttpLogEntry proto.InternalMessageInfo

func (m *HttpLogEntry) GetHttpProtocol() HttpProtocol {
	if m != nil {
		return m.HttpProtocol
	}
	return HttpProtocol_HTTP10
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

func (m *HttpLogEntry) GetHeaders() []*KeyValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpLogEntry) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *HttpLogEntry) GetMissingHeaders() []*KeyValue {
	if m != nil {
		return m.MissingHeaders
	}
	return nil
}

func (m *HttpLogEntry) GetRejectedHeaders() []*KeyValue {
	if m != nil {
		return m.RejectedHeaders
	}
	return nil
}

type KafkaLogEntry struct {
	// correlation_id is a user-supplied integer value that will be passed
	// back with the response
	CorrelationId int32 `protobuf:"varint,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// error_code is the Kafka error code being returned
	// Ref. https://kafka.apache.org/protocol#protocol_error_codes
	ErrorCode int32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	// api_version of the Kafka api used
	// Ref. https://kafka.apache.org/protocol#protocol_compatibility
	ApiVersion int32 `protobuf:"varint,3,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// api_key for Kafka message
	// Reference: https://kafka.apache.org/protocol#protocol_api_keys
	ApiKey int32 `protobuf:"varint,4,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// Topics of the request
	// Optional, as not all messages have topics (ex. LeaveGroup, Heartbeat)
	Topics               []string `protobuf:"bytes,5,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaLogEntry) Reset()         { *m = KafkaLogEntry{} }
func (m *KafkaLogEntry) String() string { return proto.CompactTextString(m) }
func (*KafkaLogEntry) ProtoMessage()    {}
func (*KafkaLogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_02c2856a892fe6b6, []int{2}
}

func (m *KafkaLogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaLogEntry.Unmarshal(m, b)
}
func (m *KafkaLogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaLogEntry.Marshal(b, m, deterministic)
}
func (m *KafkaLogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaLogEntry.Merge(m, src)
}
func (m *KafkaLogEntry) XXX_Size() int {
	return xxx_messageInfo_KafkaLogEntry.Size(m)
}
func (m *KafkaLogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaLogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaLogEntry proto.InternalMessageInfo

func (m *KafkaLogEntry) GetCorrelationId() int32 {
	if m != nil {
		return m.CorrelationId
	}
	return 0
}

func (m *KafkaLogEntry) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *KafkaLogEntry) GetApiVersion() int32 {
	if m != nil {
		return m.ApiVersion
	}
	return 0
}

func (m *KafkaLogEntry) GetApiKey() int32 {
	if m != nil {
		return m.ApiKey
	}
	return 0
}

func (m *KafkaLogEntry) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type L7LogEntry struct {
	Proto                string            `protobuf:"bytes,1,opt,name=proto,proto3" json:"proto,omitempty"`
	Fields               map[string]string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L7LogEntry) Reset()         { *m = L7LogEntry{} }
func (m *L7LogEntry) String() string { return proto.CompactTextString(m) }
func (*L7LogEntry) ProtoMessage()    {}
func (*L7LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_02c2856a892fe6b6, []int{3}
}

func (m *L7LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L7LogEntry.Unmarshal(m, b)
}
func (m *L7LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L7LogEntry.Marshal(b, m, deterministic)
}
func (m *L7LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L7LogEntry.Merge(m, src)
}
func (m *L7LogEntry) XXX_Size() int {
	return xxx_messageInfo_L7LogEntry.Size(m)
}
func (m *L7LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_L7LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_L7LogEntry proto.InternalMessageInfo

func (m *L7LogEntry) GetProto() string {
	if m != nil {
		return m.Proto
	}
	return ""
}

func (m *L7LogEntry) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type LogEntry struct {
	// The time that Cilium filter captured this log entry,
	// in, nanoseconds since 1/1/1970.
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 'true' if the request was received by an ingress listener,
	// 'false' if received by an egress listener
	IsIngress bool      `protobuf:"varint,15,opt,name=is_ingress,json=isIngress,proto3" json:"is_ingress,omitempty"`
	EntryType EntryType `protobuf:"varint,3,opt,name=entry_type,json=entryType,proto3,enum=cilium.EntryType" json:"entry_type,omitempty"`
	// Cilium network policy resource name
	PolicyName string `protobuf:"bytes,4,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// Cilium rule reference
	CiliumRuleRef string `protobuf:"bytes,5,opt,name=cilium_rule_ref,json=ciliumRuleRef,proto3" json:"cilium_rule_ref,omitempty"`
	// Cilium security ID of the source and destination
	SourceSecurityId      uint32 `protobuf:"varint,6,opt,name=source_security_id,json=sourceSecurityId,proto3" json:"source_security_id,omitempty"`
	DestinationSecurityId uint32 `protobuf:"varint,16,opt,name=destination_security_id,json=destinationSecurityId,proto3" json:"destination_security_id,omitempty"`
	// These fields record the original source and destination addresses,
	// stored in ipv4:port or [ipv6]:port format.
	SourceAddress      string `protobuf:"bytes,7,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress string `protobuf:"bytes,8,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	// Types that are valid to be assigned to L7:
	//	*LogEntry_Http
	//	*LogEntry_Kafka
	//	*LogEntry_GenericL7
	L7 isLogEntry_L7 `protobuf_oneof:"l7"`
	//
	// Deprecated HTTP fields. Use the http field above instead.
	//
	HttpProtocol HttpProtocol `protobuf:"varint,2,opt,name=http_protocol,json=httpProtocol,proto3,enum=cilium.HttpProtocol" json:"http_protocol,omitempty"` // Deprecated: Do not use.
	// Request info that is also retained for the response
	Scheme string `protobuf:"bytes,9,opt,name=scheme,proto3" json:"scheme,omitempty"`   // Deprecated: Do not use.
	Host   string `protobuf:"bytes,10,opt,name=host,proto3" json:"host,omitempty"`      // Deprecated: Do not use.
	Path   string `protobuf:"bytes,11,opt,name=path,proto3" json:"path,omitempty"`      // Deprecated: Do not use.
	Method string `protobuf:"bytes,12,opt,name=method,proto3" json:"method,omitempty"`  // Deprecated: Do not use.
	Status uint32 `protobuf:"varint,13,opt,name=status,proto3" json:"status,omitempty"` // Deprecated: Do not use.
	// Request headers not included above
	Headers              []*KeyValue `protobuf:"bytes,14,rep,name=headers,proto3" json:"headers,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_02c2856a892fe6b6, []int{4}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *LogEntry) GetIsIngress() bool {
	if m != nil {
		return m.IsIngress
	}
	return false
}

func (m *LogEntry) GetEntryType() EntryType {
	if m != nil {
		return m.EntryType
	}
	return EntryType_Request
}

func (m *LogEntry) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *LogEntry) GetCiliumRuleRef() string {
	if m != nil {
		return m.CiliumRuleRef
	}
	return ""
}

func (m *LogEntry) GetSourceSecurityId() uint32 {
	if m != nil {
		return m.SourceSecurityId
	}
	return 0
}

func (m *LogEntry) GetDestinationSecurityId() uint32 {
	if m != nil {
		return m.DestinationSecurityId
	}
	return 0
}

func (m *LogEntry) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *LogEntry) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

type isLogEntry_L7 interface {
	isLogEntry_L7()
}

type LogEntry_Http struct {
	Http *HttpLogEntry `protobuf:"bytes,100,opt,name=http,proto3,oneof"`
}

type LogEntry_Kafka struct {
	Kafka *KafkaLogEntry `protobuf:"bytes,101,opt,name=kafka,proto3,oneof"`
}

type LogEntry_GenericL7 struct {
	GenericL7 *L7LogEntry `protobuf:"bytes,102,opt,name=generic_l7,json=genericL7,proto3,oneof"`
}

func (*LogEntry_Http) isLogEntry_L7() {}

func (*LogEntry_Kafka) isLogEntry_L7() {}

func (*LogEntry_GenericL7) isLogEntry_L7() {}

func (m *LogEntry) GetL7() isLogEntry_L7 {
	if m != nil {
		return m.L7
	}
	return nil
}

func (m *LogEntry) GetHttp() *HttpLogEntry {
	if x, ok := m.GetL7().(*LogEntry_Http); ok {
		return x.Http
	}
	return nil
}

func (m *LogEntry) GetKafka() *KafkaLogEntry {
	if x, ok := m.GetL7().(*LogEntry_Kafka); ok {
		return x.Kafka
	}
	return nil
}

func (m *LogEntry) GetGenericL7() *L7LogEntry {
	if x, ok := m.GetL7().(*LogEntry_GenericL7); ok {
		return x.GenericL7
	}
	return nil
}

// Deprecated: Do not use.
func (m *LogEntry) GetHttpProtocol() HttpProtocol {
	if m != nil {
		return m.HttpProtocol
	}
	return HttpProtocol_HTTP10
}

// Deprecated: Do not use.
func (m *LogEntry) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

// Deprecated: Do not use.
func (m *LogEntry) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

// Deprecated: Do not use.
func (m *LogEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// Deprecated: Do not use.
func (m *LogEntry) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

// Deprecated: Do not use.
func (m *LogEntry) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

// Deprecated: Do not use.
func (m *LogEntry) GetHeaders() []*KeyValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LogEntry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LogEntry_Http)(nil),
		(*LogEntry_Kafka)(nil),
		(*LogEntry_GenericL7)(nil),
	}
}

func init() {
	proto.RegisterEnum("cilium.HttpProtocol", HttpProtocol_name, HttpProtocol_value)
	proto.RegisterEnum("cilium.EntryType", EntryType_name, EntryType_value)
	proto.RegisterType((*KeyValue)(nil), "cilium.KeyValue")
	proto.RegisterType((*HttpLogEntry)(nil), "cilium.HttpLogEntry")
	proto.RegisterType((*KafkaLogEntry)(nil), "cilium.KafkaLogEntry")
	proto.RegisterType((*L7LogEntry)(nil), "cilium.L7LogEntry")
	proto.RegisterMapType((map[string]string)(nil), "cilium.L7LogEntry.FieldsEntry")
	proto.RegisterType((*LogEntry)(nil), "cilium.LogEntry")
}

func init() { proto.RegisterFile("cilium/api/accesslog.proto", fileDescriptor_02c2856a892fe6b6) }

var fileDescriptor_02c2856a892fe6b6 = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xed, 0x6e, 0xe3, 0x44,
	0x14, 0x8d, 0xdd, 0xd8, 0x89, 0x6f, 0x9a, 0xd4, 0x0c, 0xdd, 0xae, 0x15, 0xf1, 0x11, 0x45, 0x5a,
	0x14, 0x55, 0x90, 0x6e, 0xb3, 0xd2, 0x96, 0x82, 0xf8, 0xc1, 0xf2, 0xa1, 0x54, 0xad, 0xd0, 0x6a,
	0xa8, 0xf6, 0xaf, 0x65, 0xec, 0x9b, 0x64, 0xa8, 0xe3, 0x31, 0x9e, 0xc9, 0x4a, 0x79, 0x04, 0x1e,
	0x80, 0x67, 0xe0, 0xa9, 0x78, 0x17, 0x34, 0x1f, 0x76, 0x4c, 0xa1, 0x12, 0xff, 0xee, 0x3d, 0xf7,
	0x9c, 0x3b, 0x33, 0xd7, 0xd7, 0x07, 0xc6, 0x29, 0xcb, 0xd9, 0x6e, 0x7b, 0x91, 0x94, 0xec, 0x22,
	0x49, 0x53, 0x14, 0x22, 0xe7, 0xeb, 0x79, 0x59, 0x71, 0xc9, 0x89, 0x6f, 0x6a, 0xd3, 0x05, 0xf4,
	0x6f, 0x71, 0xff, 0x2e, 0xc9, 0x77, 0x48, 0x42, 0x38, 0x7a, 0xc0, 0x7d, 0xe4, 0x4c, 0x9c, 0x59,
	0x40, 0x55, 0x48, 0x4e, 0xc1, 0x7b, 0xaf, 0x4a, 0x91, 0xab, 0x31, 0x93, 0x4c, 0xff, 0x72, 0xe1,
	0x78, 0x29, 0x65, 0x79, 0xc7, 0xd7, 0x3f, 0x14, 0xb2, 0xda, 0x93, 0x6b, 0x18, 0x6e, 0xa4, 0x2c,
	0x63, 0xdd, 0x3a, 0xe5, 0xb9, 0x6e, 0x31, 0x5a, 0x9c, 0xce, 0xcd, 0x21, 0x73, 0x45, 0x7e, 0x6b,
	0x6b, 0xf4, 0x78, 0xd3, 0xca, 0xc8, 0x19, 0xf8, 0x22, 0xdd, 0xe0, 0xb6, 0x3e, 0xc2, 0x66, 0x84,
	0x40, 0x77, 0xc3, 0x85, 0x8c, 0x8e, 0x34, 0xaa, 0x63, 0x85, 0x95, 0x89, 0xdc, 0x44, 0x5d, 0x83,
	0xa9, 0x58, 0xe9, 0xb7, 0x28, 0x37, 0x3c, 0x8b, 0x3c, 0xa3, 0x37, 0x19, 0x39, 0x87, 0xde, 0x06,
	0x93, 0x0c, 0x2b, 0x11, 0xf9, 0x93, 0xa3, 0xd9, 0x60, 0x11, 0xd6, 0x97, 0xa9, 0x9f, 0x4b, 0x6b,
	0x82, 0xbe, 0x83, 0x4c, 0xe4, 0x4e, 0x44, 0xbd, 0x89, 0x33, 0x1b, 0x52, 0x9b, 0x91, 0x6b, 0x38,
	0xd9, 0x32, 0x21, 0x58, 0xb1, 0x8e, 0xeb, 0x5e, 0xfd, 0x27, 0x7a, 0x8d, 0x2c, 0x71, 0x69, 0x5b,
	0x7e, 0x0d, 0x61, 0x85, 0xbf, 0x62, 0x2a, 0x31, 0x6b, 0xb4, 0xc1, 0x13, 0xda, 0x93, 0x9a, 0x69,
	0xc5, 0xd3, 0x3f, 0x1d, 0x18, 0xde, 0x26, 0xab, 0x87, 0xa4, 0x19, 0xf0, 0x0b, 0x18, 0xa5, 0xbc,
	0xaa, 0x30, 0x4f, 0x24, 0xe3, 0x45, 0xcc, 0x32, 0x3d, 0x61, 0x8f, 0x0e, 0x5b, 0xe8, 0x4d, 0x46,
	0x3e, 0x06, 0xc0, 0xaa, 0xe2, 0x55, 0x9c, 0xf2, 0xcc, 0x0c, 0xd4, 0xa3, 0x81, 0x46, 0xbe, 0xe3,
	0x19, 0x92, 0x4f, 0x61, 0x90, 0x94, 0x2c, 0x7e, 0x8f, 0x95, 0x60, 0xbc, 0xd0, 0xa3, 0xf5, 0x28,
	0x24, 0x25, 0x7b, 0x67, 0x10, 0xf2, 0x1c, 0x7a, 0x8a, 0xa0, 0x96, 0xa0, 0xab, 0x8b, 0x7e, 0x52,
	0xb2, 0x5b, 0xdc, 0xab, 0x09, 0x49, 0x5e, 0xb2, 0x54, 0x44, 0xde, 0xe4, 0x48, 0x4d, 0xd9, 0x64,
	0xd3, 0x3f, 0x1c, 0x80, 0xbb, 0xab, 0xe6, 0x9a, 0xa7, 0xe0, 0xe9, 0x15, 0xb0, 0x2b, 0x64, 0x12,
	0xf2, 0x1a, 0xfc, 0x15, 0xc3, 0x3c, 0x13, 0x91, 0xab, 0x27, 0xf0, 0x49, 0x3d, 0x81, 0x83, 0x72,
	0xfe, 0xa3, 0x26, 0xe8, 0x98, 0x5a, 0xf6, 0xf8, 0x1a, 0x06, 0x2d, 0xf8, 0xff, 0x6e, 0xe7, 0x57,
	0xee, 0x97, 0xce, 0xf4, 0x77, 0x1f, 0xfa, 0xcd, 0xad, 0x3e, 0x82, 0x40, 0xb2, 0x2d, 0x0a, 0x99,
	0x6c, 0x4b, 0x2d, 0xef, 0xd2, 0x03, 0xa0, 0x66, 0xc6, 0x44, 0xcc, 0x8a, 0x75, 0x85, 0x42, 0x44,
	0x27, 0x13, 0x67, 0xd6, 0xa7, 0x01, 0x13, 0x37, 0x06, 0x20, 0x2f, 0x01, 0x50, 0x75, 0x89, 0xe5,
	0xbe, 0x44, 0x3d, 0xb2, 0xd1, 0xe2, 0x83, 0xfa, 0x01, 0xba, 0xff, 0xfd, 0xbe, 0x44, 0x1a, 0x60,
	0x1d, 0xaa, 0x29, 0x97, 0x3c, 0x67, 0xe9, 0x3e, 0x2e, 0x92, 0x2d, 0xda, 0x65, 0x05, 0x03, 0xfd,
	0x94, 0x6c, 0x91, 0x7c, 0x06, 0x27, 0x46, 0x1f, 0x57, 0xbb, 0x1c, 0xe3, 0x0a, 0x57, 0x76, 0x77,
	0x87, 0x06, 0xa6, 0xbb, 0x1c, 0x29, 0xae, 0xc8, 0xe7, 0x40, 0x04, 0xdf, 0x55, 0x29, 0xc6, 0x02,
	0xd3, 0x5d, 0xc5, 0xe4, 0x5e, 0x7d, 0x78, 0x5f, 0xaf, 0x68, 0x68, 0x2a, 0x3f, 0xdb, 0xc2, 0x4d,
	0x46, 0x5e, 0xc3, 0xf3, 0x0c, 0x85, 0x64, 0x85, 0x59, 0x91, 0xb6, 0x24, 0xd4, 0x92, 0x67, 0xad,
	0x72, 0x4b, 0xf7, 0x02, 0x46, 0xf6, 0x94, 0x24, 0xcb, 0xf4, 0x0c, 0x7a, 0xe6, 0x32, 0x06, 0xfd,
	0xd6, 0x80, 0xe4, 0x02, 0x3e, 0x6c, 0xb7, 0xaf, 0xb9, 0x7d, 0xcd, 0x25, 0xad, 0x52, 0x2d, 0x38,
	0x87, 0xae, 0xfa, 0xd1, 0xa3, 0x6c, 0xe2, 0xcc, 0x06, 0xff, 0xb4, 0x82, 0xfa, 0xcb, 0x2c, 0x3b,
	0x54, 0x73, 0xc8, 0x17, 0xe0, 0x3d, 0xa8, 0x7d, 0x8f, 0x50, 0x93, 0x9f, 0x35, 0xbf, 0x48, 0xfb,
	0x27, 0x58, 0x76, 0xa8, 0x61, 0x91, 0x57, 0x00, 0x6b, 0x2c, 0xb0, 0x62, 0x69, 0x9c, 0x5f, 0x45,
	0x2b, 0xad, 0x21, 0xff, 0x5e, 0xaa, 0x65, 0x87, 0x06, 0x96, 0x77, 0x77, 0x45, 0xbe, 0x79, 0xec,
	0x51, 0xee, 0xd3, 0x1e, 0xf5, 0xc6, 0x8d, 0x9c, 0x47, 0x3e, 0x35, 0x6e, 0x7c, 0x2a, 0x50, 0x4f,
	0xd6, 0x8c, 0xda, 0xab, 0xce, 0xac, 0x57, 0x41, 0x53, 0x31, 0x7e, 0x75, 0x66, 0xfd, 0x6a, 0x70,
	0xc0, 0xb5, 0x67, 0x8d, 0x1b, 0xcf, 0x3a, 0x3e, 0xf4, 0xb2, 0xbe, 0x35, 0x6e, 0xbc, 0x68, 0xa8,
	0xbe, 0x9a, 0x3d, 0xc7, 0xf8, 0xd1, 0xfc, 0xe0, 0x69, 0xa3, 0xff, 0xf6, 0x12, 0x4d, 0xaf, 0x49,
	0x6f, 0xba, 0xe0, 0xe6, 0x57, 0xe7, 0x97, 0xc6, 0xac, 0x9b, 0x97, 0x00, 0xf8, 0xcb, 0xfb, 0xfb,
	0xb7, 0x97, 0x2f, 0xc3, 0x4e, 0x13, 0x5f, 0x86, 0x0e, 0x09, 0xc0, 0x53, 0xf1, 0x22, 0x74, 0xcf,
	0x17, 0x10, 0x34, 0xab, 0x4d, 0x06, 0xd0, 0xa3, 0xf8, 0xdb, 0x0e, 0x85, 0x0c, 0x3b, 0xe4, 0x18,
	0xfa, 0x14, 0x45, 0xc9, 0x0b, 0x81, 0xa1, 0xa3, 0xe4, 0xdf, 0x63, 0xc1, 0x30, 0x0b, 0xdd, 0x5f,
	0x7c, 0x3d, 0xd8, 0x57, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x06, 0x6b, 0x47, 0x1f, 0x75, 0x06,
	0x00, 0x00,
}
