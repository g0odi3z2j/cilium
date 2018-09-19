// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/accesslog/v2/als.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
	v2 "github.com/cilium/cilium/pkg/envoy/envoy/data/accesslog/v2"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Empty response for the StreamAccessLogs API. Will never be sent. See below.
type StreamAccessLogsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccessLogsResponse) Reset()         { *m = StreamAccessLogsResponse{} }
func (m *StreamAccessLogsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsResponse) ProtoMessage()    {}
func (*StreamAccessLogsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{0}
}

func (m *StreamAccessLogsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsResponse.Unmarshal(m, b)
}
func (m *StreamAccessLogsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsResponse.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsResponse.Merge(m, src)
}
func (m *StreamAccessLogsResponse) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsResponse.Size(m)
}
func (m *StreamAccessLogsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsResponse proto.InternalMessageInfo

// Stream message for the StreamAccessLogs API. Envoy will open a stream to the server and stream
// access logs without ever expecting a response.
type StreamAccessLogsMessage struct {
	// Identifier data that will only be sent in the first message on the stream. This is effectively
	// structured metadata and is a performance optimization.
	Identifier *StreamAccessLogsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Batches of log entries of a single type. Generally speaking, a given stream should only
	// ever include one type of log entry.
	//
	// Types that are valid to be assigned to LogEntries:
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	LogEntries           isStreamAccessLogsMessage_LogEntries `protobuf_oneof:"log_entries"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *StreamAccessLogsMessage) Reset()         { *m = StreamAccessLogsMessage{} }
func (m *StreamAccessLogsMessage) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage) ProtoMessage()    {}
func (*StreamAccessLogsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1}
}

func (m *StreamAccessLogsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage.Merge(m, src)
}
func (m *StreamAccessLogsMessage) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage.Size(m)
}
func (m *StreamAccessLogsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage proto.InternalMessageInfo

func (m *StreamAccessLogsMessage) GetIdentifier() *StreamAccessLogsMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

type isStreamAccessLogsMessage_LogEntries interface {
	isStreamAccessLogsMessage_LogEntries()
}

type StreamAccessLogsMessage_HttpLogs struct {
	HttpLogs *StreamAccessLogsMessage_HTTPAccessLogEntries `protobuf:"bytes,2,opt,name=http_logs,json=httpLogs,proto3,oneof"`
}

type StreamAccessLogsMessage_TcpLogs struct {
	TcpLogs *StreamAccessLogsMessage_TCPAccessLogEntries `protobuf:"bytes,3,opt,name=tcp_logs,json=tcpLogs,proto3,oneof"`
}

func (*StreamAccessLogsMessage_HttpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (*StreamAccessLogsMessage_TcpLogs) isStreamAccessLogsMessage_LogEntries() {}

func (m *StreamAccessLogsMessage) GetLogEntries() isStreamAccessLogsMessage_LogEntries {
	if m != nil {
		return m.LogEntries
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetHttpLogs() *StreamAccessLogsMessage_HTTPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_HttpLogs); ok {
		return x.HttpLogs
	}
	return nil
}

func (m *StreamAccessLogsMessage) GetTcpLogs() *StreamAccessLogsMessage_TCPAccessLogEntries {
	if x, ok := m.GetLogEntries().(*StreamAccessLogsMessage_TcpLogs); ok {
		return x.TcpLogs
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamAccessLogsMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamAccessLogsMessage_OneofMarshaler, _StreamAccessLogsMessage_OneofUnmarshaler, _StreamAccessLogsMessage_OneofSizer, []interface{}{
		(*StreamAccessLogsMessage_HttpLogs)(nil),
		(*StreamAccessLogsMessage_TcpLogs)(nil),
	}
}

func _StreamAccessLogsMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamAccessLogsMessage)
	// log_entries
	switch x := m.LogEntries.(type) {
	case *StreamAccessLogsMessage_HttpLogs:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpLogs); err != nil {
			return err
		}
	case *StreamAccessLogsMessage_TcpLogs:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TcpLogs); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamAccessLogsMessage.LogEntries has unexpected type %T", x)
	}
	return nil
}

func _StreamAccessLogsMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamAccessLogsMessage)
	switch tag {
	case 2: // log_entries.http_logs
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamAccessLogsMessage_HTTPAccessLogEntries)
		err := b.DecodeMessage(msg)
		m.LogEntries = &StreamAccessLogsMessage_HttpLogs{msg}
		return true, err
	case 3: // log_entries.tcp_logs
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamAccessLogsMessage_TCPAccessLogEntries)
		err := b.DecodeMessage(msg)
		m.LogEntries = &StreamAccessLogsMessage_TcpLogs{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamAccessLogsMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamAccessLogsMessage)
	// log_entries
	switch x := m.LogEntries.(type) {
	case *StreamAccessLogsMessage_HttpLogs:
		s := proto.Size(x.HttpLogs)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamAccessLogsMessage_TcpLogs:
		s := proto.Size(x.TcpLogs)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamAccessLogsMessage_Identifier struct {
	// The node sending the access log messages over the stream.
	Node *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// The friendly name of the log configured in :ref:`CommonGrpcAccessLogConfig
	// <envoy_api_msg_config.accesslog.v2.CommonGrpcAccessLogConfig>`.
	LogName              string   `protobuf:"bytes,2,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamAccessLogsMessage_Identifier) Reset()         { *m = StreamAccessLogsMessage_Identifier{} }
func (m *StreamAccessLogsMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamAccessLogsMessage_Identifier) ProtoMessage()    {}
func (*StreamAccessLogsMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 0}
}

func (m *StreamAccessLogsMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_Identifier.Merge(m, src)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_Identifier.Size(m)
}
func (m *StreamAccessLogsMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_Identifier proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *StreamAccessLogsMessage_Identifier) GetLogName() string {
	if m != nil {
		return m.LogName
	}
	return ""
}

// Wrapper for batches of HTTP access log entries.
type StreamAccessLogsMessage_HTTPAccessLogEntries struct {
	LogEntry             []*v2.HTTPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_HTTPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_HTTPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 1}
}

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Merge(m, src)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.Size(m)
}
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_HTTPAccessLogEntries proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) GetLogEntry() []*v2.HTTPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

// [#not-implemented-hide:]
// Wrapper for batches of TCP access log entries.
type StreamAccessLogsMessage_TCPAccessLogEntries struct {
	LogEntry             []*v2.TCPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Reset() {
	*m = StreamAccessLogsMessage_TCPAccessLogEntries{}
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) String() string {
	return proto.CompactTextString(m)
}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) ProtoMessage() {}
func (*StreamAccessLogsMessage_TCPAccessLogEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4f3a3a69261b513, []int{1, 2}
}

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Unmarshal(m, b)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Marshal(b, m, deterministic)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Merge(m, src)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_Size() int {
	return xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.Size(m)
}
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries.DiscardUnknown(m)
}

var xxx_messageInfo_StreamAccessLogsMessage_TCPAccessLogEntries proto.InternalMessageInfo

func (m *StreamAccessLogsMessage_TCPAccessLogEntries) GetLogEntry() []*v2.TCPAccessLogEntry {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamAccessLogsResponse)(nil), "envoy.service.accesslog.v2.StreamAccessLogsResponse")
	proto.RegisterType((*StreamAccessLogsMessage)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage")
	proto.RegisterType((*StreamAccessLogsMessage_Identifier)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.Identifier")
	proto.RegisterType((*StreamAccessLogsMessage_HTTPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.HTTPAccessLogEntries")
	proto.RegisterType((*StreamAccessLogsMessage_TCPAccessLogEntries)(nil), "envoy.service.accesslog.v2.StreamAccessLogsMessage.TCPAccessLogEntries")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccessLogServiceClient is the client API for AccessLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccessLogServiceClient interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error)
}

type accessLogServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccessLogServiceClient(cc *grpc.ClientConn) AccessLogServiceClient {
	return &accessLogServiceClient{cc}
}

func (c *accessLogServiceClient) StreamAccessLogs(ctx context.Context, opts ...grpc.CallOption) (AccessLogService_StreamAccessLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccessLogService_serviceDesc.Streams[0], "/envoy.service.accesslog.v2.AccessLogService/StreamAccessLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &accessLogServiceStreamAccessLogsClient{stream}
	return x, nil
}

type AccessLogService_StreamAccessLogsClient interface {
	Send(*StreamAccessLogsMessage) error
	CloseAndRecv() (*StreamAccessLogsResponse, error)
	grpc.ClientStream
}

type accessLogServiceStreamAccessLogsClient struct {
	grpc.ClientStream
}

func (x *accessLogServiceStreamAccessLogsClient) Send(m *StreamAccessLogsMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsClient) CloseAndRecv() (*StreamAccessLogsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamAccessLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccessLogServiceServer is the server API for AccessLogService service.
type AccessLogServiceServer interface {
	// Envoy will connect and send StreamAccessLogsMessage messages forever. It does not expect any
	// response to be sent as nothing would be done in the case of failure. The server should
	// disconnect if it expects Envoy to reconnect. In the future we may decide to add a different
	// API for "critical" access logs in which Envoy will buffer access logs for some period of time
	// until it gets an ACK so it could then retry. This API is designed for high throughput with the
	// expectation that it might be lossy.
	StreamAccessLogs(AccessLogService_StreamAccessLogsServer) error
}

func RegisterAccessLogServiceServer(s *grpc.Server, srv AccessLogServiceServer) {
	s.RegisterService(&_AccessLogService_serviceDesc, srv)
}

func _AccessLogService_StreamAccessLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AccessLogServiceServer).StreamAccessLogs(&accessLogServiceStreamAccessLogsServer{stream})
}

type AccessLogService_StreamAccessLogsServer interface {
	SendAndClose(*StreamAccessLogsResponse) error
	Recv() (*StreamAccessLogsMessage, error)
	grpc.ServerStream
}

type accessLogServiceStreamAccessLogsServer struct {
	grpc.ServerStream
}

func (x *accessLogServiceStreamAccessLogsServer) SendAndClose(m *StreamAccessLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *accessLogServiceStreamAccessLogsServer) Recv() (*StreamAccessLogsMessage, error) {
	m := new(StreamAccessLogsMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AccessLogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.accesslog.v2.AccessLogService",
	HandlerType: (*AccessLogServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAccessLogs",
			Handler:       _AccessLogService_StreamAccessLogs_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/accesslog/v2/als.proto",
}

func init() {
	proto.RegisterFile("envoy/service/accesslog/v2/als.proto", fileDescriptor_e4f3a3a69261b513)
}

var fileDescriptor_e4f3a3a69261b513 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x77, 0xd2, 0xad, 0x6d, 0x5f, 0x2f, 0x65, 0x5c, 0x68, 0x09, 0x1e, 0xca, 0xb2, 0x60,
	0x51, 0x98, 0x40, 0x56, 0xf0, 0x26, 0x58, 0x11, 0x2b, 0xe8, 0x22, 0x69, 0x4f, 0x1e, 0x5c, 0x66,
	0x93, 0x67, 0x1c, 0x4d, 0x33, 0x61, 0x66, 0x08, 0xf4, 0xe8, 0x4d, 0x3c, 0x7a, 0xf0, 0x1b, 0xf8,
	0x1d, 0xc4, 0xd3, 0x7e, 0x9d, 0xfd, 0x16, 0x32, 0x93, 0x6c, 0xd4, 0xb6, 0x39, 0xd8, 0x5b, 0x98,
	0x79, 0xff, 0xdf, 0x6f, 0xde, 0xbc, 0x0c, 0x9c, 0x61, 0x5e, 0xca, 0x4d, 0xa0, 0x51, 0x95, 0x22,
	0xc6, 0x80, 0xc7, 0x31, 0x6a, 0x9d, 0xc9, 0x34, 0x28, 0xc3, 0x80, 0x67, 0x9a, 0x15, 0x4a, 0x1a,
	0x49, 0x7d, 0x57, 0xc5, 0xea, 0x2a, 0xd6, 0x54, 0xb1, 0x32, 0xf4, 0xef, 0x55, 0x04, 0x5e, 0x08,
	0x9b, 0x89, 0xa5, 0xc2, 0xe0, 0x8a, 0x6b, 0xac, 0x92, 0xfe, 0xfd, 0x6a, 0x37, 0xe1, 0x86, 0x6f,
	0xc1, 0x1b, 0x46, 0x55, 0x38, 0x2e, 0x79, 0x26, 0x12, 0x6e, 0x30, 0xb8, 0xfd, 0xa8, 0x36, 0x4e,
	0x7d, 0x98, 0x2c, 0x8d, 0x42, 0xbe, 0x7e, 0xea, 0x12, 0xaf, 0x64, 0xaa, 0x23, 0xd4, 0x85, 0xcc,
	0x35, 0x9e, 0xfe, 0xe8, 0xc2, 0x78, 0x7b, 0xf3, 0x35, 0x6a, 0xcd, 0x53, 0xa4, 0xef, 0x00, 0x44,
	0x82, 0xb9, 0x11, 0xef, 0x05, 0xaa, 0x09, 0x99, 0x92, 0xd9, 0x30, 0x7c, 0xc2, 0xda, 0x1b, 0x61,
	0x2d, 0x20, 0xf6, 0xb2, 0xa1, 0x44, 0x7f, 0x11, 0x69, 0x0a, 0x83, 0x0f, 0xc6, 0x14, 0x97, 0x99,
	0x4c, 0xf5, 0xc4, 0x73, 0xf8, 0xc5, 0x21, 0xf8, 0xc5, 0x6a, 0xf5, 0xa6, 0x59, 0x7d, 0x9e, 0x1b,
	0x25, 0x50, 0x2f, 0x8e, 0xa2, 0xbe, 0x85, 0xdb, 0x3a, 0x9a, 0x40, 0xdf, 0xc4, 0xb5, 0xa7, 0xe3,
	0x3c, 0x2f, 0x0e, 0xf1, 0xac, 0x9e, 0xed, 0xd3, 0xf4, 0x4c, 0xec, 0x2c, 0xfe, 0x27, 0x80, 0x3f,
	0x8d, 0xd2, 0xc7, 0x70, 0x9c, 0xcb, 0x04, 0xeb, 0x6b, 0x1b, 0xd7, 0x3e, 0x5e, 0x08, 0x6b, 0xb0,
	0x33, 0x66, 0x17, 0x32, 0xc1, 0x39, 0xfc, 0xba, 0xb9, 0xee, 0x74, 0xbf, 0x12, 0x6f, 0x44, 0x22,
	0x17, 0xa0, 0x67, 0xd0, 0xcf, 0x64, 0x7a, 0x99, 0xf3, 0x35, 0xba, 0x4b, 0x19, 0xcc, 0x07, 0xb6,
	0xe6, 0x58, 0x79, 0x53, 0x12, 0xf5, 0x32, 0x99, 0x5e, 0xf0, 0x35, 0xfa, 0x19, 0x9c, 0xec, 0x6b,
	0x9b, 0xae, 0x60, 0x60, 0xd3, 0x98, 0x1b, 0xb5, 0x99, 0x90, 0x69, 0x67, 0x36, 0x0c, 0x1f, 0xd6,
	0x6e, 0xfb, 0x07, 0xfd, 0xdb, 0xe8, 0x0e, 0x61, 0x53, 0x9f, 0xe7, 0x1b, 0xf1, 0xfa, 0x24, 0xb2,
	0xe7, 0x70, 0xab, 0xfe, 0x47, 0xb8, 0xbb, 0xa7, 0x79, 0xba, 0xdc, 0x95, 0x3d, 0x68, 0x95, 0x6d,
	0x03, 0x5a, 0x5c, 0xf3, 0x13, 0x18, 0xde, 0x42, 0xad, 0xa3, 0xfb, 0xf3, 0xe6, 0xba, 0x43, 0xc2,
	0xef, 0x04, 0x46, 0x4d, 0x7c, 0x59, 0x4d, 0x8d, 0x7e, 0x26, 0x30, 0xda, 0x1e, 0x16, 0x3d, 0x3f,
	0x60, 0xb4, 0xfe, 0xa3, 0xff, 0x09, 0x35, 0x8f, 0xe7, 0x68, 0x46, 0xe6, 0xbd, 0xb7, 0x5e, 0x19,
	0x7e, 0x21, 0xe4, 0xea, 0x8e, 0x7b, 0x6c, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x00,
	0xb2, 0x16, 0x10, 0x04, 0x00, 0x00,
}
