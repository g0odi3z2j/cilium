// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/health_check.proto

package core

import (
	fmt "fmt"
	_type "github.com/cilium/proxy/go/envoy/type"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Endpoint health status.
type HealthStatus int32

const (
	// The health status is not known. This is interpreted by Envoy as *HEALTHY*.
	HealthStatus_UNKNOWN HealthStatus = 0
	// Healthy.
	HealthStatus_HEALTHY HealthStatus = 1
	// Unhealthy.
	HealthStatus_UNHEALTHY HealthStatus = 2
	// Connection draining in progress. E.g.,
	// `<https://aws.amazon.com/blogs/aws/elb-connection-draining-remove-instances-from-service-with-care/>`_
	// or
	// `<https://cloud.google.com/compute/docs/load-balancing/enabling-connection-draining>`_.
	// This is interpreted by Envoy as *UNHEALTHY*.
	HealthStatus_DRAINING HealthStatus = 3
	// Health check timed out. This is part of HDS and is interpreted by Envoy as
	// *UNHEALTHY*.
	HealthStatus_TIMEOUT HealthStatus = 4
	// Degraded.
	HealthStatus_DEGRADED HealthStatus = 5
)

var HealthStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "HEALTHY",
	2: "UNHEALTHY",
	3: "DRAINING",
	4: "TIMEOUT",
	5: "DEGRADED",
}

var HealthStatus_value = map[string]int32{
	"UNKNOWN":   0,
	"HEALTHY":   1,
	"UNHEALTHY": 2,
	"DRAINING":  3,
	"TIMEOUT":   4,
	"DEGRADED":  5,
}

func (x HealthStatus) String() string {
	return proto.EnumName(HealthStatus_name, int32(x))
}

func (HealthStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0}
}

type HealthCheck struct {
	// The time to wait for a health check response. If the timeout is reached the
	// health check attempt will be considered a failure.
	Timeout *duration.Duration `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The interval between health checks.
	Interval *duration.Duration `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	// An optional jitter amount in milliseconds. If specified, Envoy will start health
	// checking after for a random time in ms between 0 and initial_jitter. This only
	// applies to the first health check.
	InitialJitter *duration.Duration `protobuf:"bytes,20,opt,name=initial_jitter,json=initialJitter,proto3" json:"initial_jitter,omitempty"`
	// An optional jitter amount in milliseconds. If specified, during every
	// interval Envoy will add interval_jitter to the wait time.
	IntervalJitter *duration.Duration `protobuf:"bytes,3,opt,name=interval_jitter,json=intervalJitter,proto3" json:"interval_jitter,omitempty"`
	// An optional jitter amount as a percentage of interval_ms. If specified,
	// during every interval Envoy will add interval_ms *
	// interval_jitter_percent / 100 to the wait time.
	//
	// If interval_jitter_ms and interval_jitter_percent are both set, both of
	// them will be used to increase the wait time.
	IntervalJitterPercent uint32 `protobuf:"varint,18,opt,name=interval_jitter_percent,json=intervalJitterPercent,proto3" json:"interval_jitter_percent,omitempty"`
	// The number of unhealthy health checks required before a host is marked
	// unhealthy. Note that for *http* health checking if a host responds with 503
	// this threshold is ignored and the host is considered unhealthy immediately.
	UnhealthyThreshold *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=unhealthy_threshold,json=unhealthyThreshold,proto3" json:"unhealthy_threshold,omitempty"`
	// The number of healthy health checks required before a host is marked
	// healthy. Note that during startup, only a single successful health check is
	// required to mark a host healthy.
	HealthyThreshold *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=healthy_threshold,json=healthyThreshold,proto3" json:"healthy_threshold,omitempty"`
	// [#not-implemented-hide:] Non-serving port for health checking.
	AltPort *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=alt_port,json=altPort,proto3" json:"alt_port,omitempty"`
	// Reuse health check connection between health checks. Default is true.
	ReuseConnection *wrappers.BoolValue `protobuf:"bytes,7,opt,name=reuse_connection,json=reuseConnection,proto3" json:"reuse_connection,omitempty"`
	// Types that are valid to be assigned to HealthChecker:
	//	*HealthCheck_HttpHealthCheck_
	//	*HealthCheck_TcpHealthCheck_
	//	*HealthCheck_GrpcHealthCheck_
	//	*HealthCheck_CustomHealthCheck_
	HealthChecker isHealthCheck_HealthChecker `protobuf_oneof:"health_checker"`
	// The "no traffic interval" is a special health check interval that is used when a cluster has
	// never had traffic routed to it. This lower interval allows cluster information to be kept up to
	// date, without sending a potentially large amount of active health checking traffic for no
	// reason. Once a cluster has been used for traffic routing, Envoy will shift back to using the
	// standard health check interval that is defined. Note that this interval takes precedence over
	// any other.
	//
	// The default value for "no traffic interval" is 60 seconds.
	NoTrafficInterval *duration.Duration `protobuf:"bytes,12,opt,name=no_traffic_interval,json=noTrafficInterval,proto3" json:"no_traffic_interval,omitempty"`
	// The "unhealthy interval" is a health check interval that is used for hosts that are marked as
	// unhealthy. As soon as the host is marked as healthy, Envoy will shift back to using the
	// standard health check interval that is defined.
	//
	// The default value for "unhealthy interval" is the same as "interval".
	UnhealthyInterval *duration.Duration `protobuf:"bytes,14,opt,name=unhealthy_interval,json=unhealthyInterval,proto3" json:"unhealthy_interval,omitempty"`
	// The "unhealthy edge interval" is a special health check interval that is used for the first
	// health check right after a host is marked as unhealthy. For subsequent health checks
	// Envoy will shift back to using either "unhealthy interval" if present or the standard health
	// check interval that is defined.
	//
	// The default value for "unhealthy edge interval" is the same as "unhealthy interval".
	UnhealthyEdgeInterval *duration.Duration `protobuf:"bytes,15,opt,name=unhealthy_edge_interval,json=unhealthyEdgeInterval,proto3" json:"unhealthy_edge_interval,omitempty"`
	// The "healthy edge interval" is a special health check interval that is used for the first
	// health check right after a host is marked as healthy. For subsequent health checks
	// Envoy will shift back to using the standard health check interval that is defined.
	//
	// The default value for "healthy edge interval" is the same as the default interval.
	HealthyEdgeInterval *duration.Duration `protobuf:"bytes,16,opt,name=healthy_edge_interval,json=healthyEdgeInterval,proto3" json:"healthy_edge_interval,omitempty"`
	// Specifies the path to the :ref:`health check event log <arch_overview_health_check_logging>`.
	// If empty, no event log will be written.
	EventLogPath string `protobuf:"bytes,17,opt,name=event_log_path,json=eventLogPath,proto3" json:"event_log_path,omitempty"`
	// If set to true, health check failure events will always be logged. If set to false, only the
	// initial health check failure event will be logged.
	// The default value is false.
	AlwaysLogHealthCheckFailures bool     `protobuf:"varint,19,opt,name=always_log_health_check_failures,json=alwaysLogHealthCheckFailures,proto3" json:"always_log_health_check_failures,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *HealthCheck) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *HealthCheck) GetInitialJitter() *duration.Duration {
	if m != nil {
		return m.InitialJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitter() *duration.Duration {
	if m != nil {
		return m.IntervalJitter
	}
	return nil
}

func (m *HealthCheck) GetIntervalJitterPercent() uint32 {
	if m != nil {
		return m.IntervalJitterPercent
	}
	return 0
}

func (m *HealthCheck) GetUnhealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.UnhealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetHealthyThreshold() *wrappers.UInt32Value {
	if m != nil {
		return m.HealthyThreshold
	}
	return nil
}

func (m *HealthCheck) GetAltPort() *wrappers.UInt32Value {
	if m != nil {
		return m.AltPort
	}
	return nil
}

func (m *HealthCheck) GetReuseConnection() *wrappers.BoolValue {
	if m != nil {
		return m.ReuseConnection
	}
	return nil
}

type isHealthCheck_HealthChecker interface {
	isHealthCheck_HealthChecker()
}

type HealthCheck_HttpHealthCheck_ struct {
	HttpHealthCheck *HealthCheck_HttpHealthCheck `protobuf:"bytes,8,opt,name=http_health_check,json=httpHealthCheck,proto3,oneof"`
}

type HealthCheck_TcpHealthCheck_ struct {
	TcpHealthCheck *HealthCheck_TcpHealthCheck `protobuf:"bytes,9,opt,name=tcp_health_check,json=tcpHealthCheck,proto3,oneof"`
}

type HealthCheck_GrpcHealthCheck_ struct {
	GrpcHealthCheck *HealthCheck_GrpcHealthCheck `protobuf:"bytes,11,opt,name=grpc_health_check,json=grpcHealthCheck,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_ struct {
	CustomHealthCheck *HealthCheck_CustomHealthCheck `protobuf:"bytes,13,opt,name=custom_health_check,json=customHealthCheck,proto3,oneof"`
}

func (*HealthCheck_HttpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_TcpHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_GrpcHealthCheck_) isHealthCheck_HealthChecker() {}

func (*HealthCheck_CustomHealthCheck_) isHealthCheck_HealthChecker() {}

func (m *HealthCheck) GetHealthChecker() isHealthCheck_HealthChecker {
	if m != nil {
		return m.HealthChecker
	}
	return nil
}

func (m *HealthCheck) GetHttpHealthCheck() *HealthCheck_HttpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_HttpHealthCheck_); ok {
		return x.HttpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetTcpHealthCheck() *HealthCheck_TcpHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_TcpHealthCheck_); ok {
		return x.TcpHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetGrpcHealthCheck() *HealthCheck_GrpcHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_GrpcHealthCheck_); ok {
		return x.GrpcHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetCustomHealthCheck() *HealthCheck_CustomHealthCheck {
	if x, ok := m.GetHealthChecker().(*HealthCheck_CustomHealthCheck_); ok {
		return x.CustomHealthCheck
	}
	return nil
}

func (m *HealthCheck) GetNoTrafficInterval() *duration.Duration {
	if m != nil {
		return m.NoTrafficInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyInterval
	}
	return nil
}

func (m *HealthCheck) GetUnhealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.UnhealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetHealthyEdgeInterval() *duration.Duration {
	if m != nil {
		return m.HealthyEdgeInterval
	}
	return nil
}

func (m *HealthCheck) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

func (m *HealthCheck) GetAlwaysLogHealthCheckFailures() bool {
	if m != nil {
		return m.AlwaysLogHealthCheckFailures
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_HttpHealthCheck_)(nil),
		(*HealthCheck_TcpHealthCheck_)(nil),
		(*HealthCheck_GrpcHealthCheck_)(nil),
		(*HealthCheck_CustomHealthCheck_)(nil),
	}
}

// Describes the encoding of the payload bytes in the payload.
type HealthCheck_Payload struct {
	// Types that are valid to be assigned to Payload:
	//	*HealthCheck_Payload_Text
	//	*HealthCheck_Payload_Binary
	Payload              isHealthCheck_Payload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *HealthCheck_Payload) Reset()         { *m = HealthCheck_Payload{} }
func (m *HealthCheck_Payload) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_Payload) ProtoMessage()    {}
func (*HealthCheck_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 0}
}

func (m *HealthCheck_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_Payload.Unmarshal(m, b)
}
func (m *HealthCheck_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_Payload.Marshal(b, m, deterministic)
}
func (m *HealthCheck_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_Payload.Merge(m, src)
}
func (m *HealthCheck_Payload) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_Payload.Size(m)
}
func (m *HealthCheck_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_Payload proto.InternalMessageInfo

type isHealthCheck_Payload_Payload interface {
	isHealthCheck_Payload_Payload()
}

type HealthCheck_Payload_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type HealthCheck_Payload_Binary struct {
	Binary []byte `protobuf:"bytes,2,opt,name=binary,proto3,oneof"`
}

func (*HealthCheck_Payload_Text) isHealthCheck_Payload_Payload() {}

func (*HealthCheck_Payload_Binary) isHealthCheck_Payload_Payload() {}

func (m *HealthCheck_Payload) GetPayload() isHealthCheck_Payload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *HealthCheck_Payload) GetText() string {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Text); ok {
		return x.Text
	}
	return ""
}

func (m *HealthCheck_Payload) GetBinary() []byte {
	if x, ok := m.GetPayload().(*HealthCheck_Payload_Binary); ok {
		return x.Binary
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_Payload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_Payload_Text)(nil),
		(*HealthCheck_Payload_Binary)(nil),
	}
}

// [#comment:next free field: 10]
type HealthCheck_HttpHealthCheck struct {
	// The value of the host header in the HTTP health check request. If
	// left empty (default value), the name of the cluster this health check is associated
	// with will be used.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Specifies the HTTP path that will be requested during health checking. For example
	// */healthcheck*.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// [#not-implemented-hide:] HTTP specific payload.
	Send *HealthCheck_Payload `protobuf:"bytes,3,opt,name=send,proto3" json:"send,omitempty"`
	// [#not-implemented-hide:] HTTP specific response.
	Receive *HealthCheck_Payload `protobuf:"bytes,4,opt,name=receive,proto3" json:"receive,omitempty"`
	// An optional service name parameter which is used to validate the identity of
	// the health checked cluster. See the :ref:`architecture overview
	// <arch_overview_health_checking_identity>` for more information.
	ServiceName string `protobuf:"bytes,5,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request that is sent to the
	// health checked cluster. For more information, including details on header value syntax, see
	// the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	RequestHeadersToAdd []*HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request that is sent to the
	// health checked cluster.
	RequestHeadersToRemove []string `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	// If set, health checks will be made using http/2.
	UseHttp2 bool `protobuf:"varint,7,opt,name=use_http2,json=useHttp2,proto3" json:"use_http2,omitempty"`
	// Specifies a list of HTTP response statuses considered healthy. If provided, replaces default
	// 200-only policy - 200 must be included explicitly as needed. Ranges follow half-open
	// semantics of :ref:`Int64Range <envoy_api_msg_type.Int64Range>`.
	ExpectedStatuses     []*_type.Int64Range `protobuf:"bytes,9,rep,name=expected_statuses,json=expectedStatuses,proto3" json:"expected_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HealthCheck_HttpHealthCheck) Reset()         { *m = HealthCheck_HttpHealthCheck{} }
func (m *HealthCheck_HttpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_HttpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_HttpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 1}
}

func (m *HealthCheck_HttpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_HttpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_HttpHealthCheck.Size(m)
}
func (m *HealthCheck_HttpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_HttpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_HttpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_HttpHealthCheck) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetReceive() *HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToAdd() []*HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *HealthCheck_HttpHealthCheck) GetUseHttp2() bool {
	if m != nil {
		return m.UseHttp2
	}
	return false
}

func (m *HealthCheck_HttpHealthCheck) GetExpectedStatuses() []*_type.Int64Range {
	if m != nil {
		return m.ExpectedStatuses
	}
	return nil
}

type HealthCheck_TcpHealthCheck struct {
	// Empty payloads imply a connect-only health check.
	Send *HealthCheck_Payload `protobuf:"bytes,1,opt,name=send,proto3" json:"send,omitempty"`
	// When checking the response, “fuzzy” matching is performed such that each
	// binary block must be found, and in the order specified, but not
	// necessarily contiguous.
	Receive              []*HealthCheck_Payload `protobuf:"bytes,2,rep,name=receive,proto3" json:"receive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HealthCheck_TcpHealthCheck) Reset()         { *m = HealthCheck_TcpHealthCheck{} }
func (m *HealthCheck_TcpHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_TcpHealthCheck) ProtoMessage()    {}
func (*HealthCheck_TcpHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 2}
}

func (m *HealthCheck_TcpHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.Merge(m, src)
}
func (m *HealthCheck_TcpHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_TcpHealthCheck.Size(m)
}
func (m *HealthCheck_TcpHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_TcpHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_TcpHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_TcpHealthCheck) GetSend() *HealthCheck_Payload {
	if m != nil {
		return m.Send
	}
	return nil
}

func (m *HealthCheck_TcpHealthCheck) GetReceive() []*HealthCheck_Payload {
	if m != nil {
		return m.Receive
	}
	return nil
}

type HealthCheck_RedisHealthCheck struct {
	// If set, optionally perform ``EXISTS <key>`` instead of ``PING``. A return value
	// from Redis of 0 (does not exist) is considered a passing healthcheck. A return value other
	// than 0 is considered a failure. This allows the user to mark a Redis instance for maintenance
	// by setting the specified key to any value and waiting for traffic to drain.
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_RedisHealthCheck) Reset()         { *m = HealthCheck_RedisHealthCheck{} }
func (m *HealthCheck_RedisHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_RedisHealthCheck) ProtoMessage()    {}
func (*HealthCheck_RedisHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 3}
}

func (m *HealthCheck_RedisHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.Merge(m, src)
}
func (m *HealthCheck_RedisHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_RedisHealthCheck.Size(m)
}
func (m *HealthCheck_RedisHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_RedisHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_RedisHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_RedisHealthCheck) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// `grpc.health.v1.Health
// <https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto>`_-based
// healthcheck. See `gRPC doc <https://github.com/grpc/grpc/blob/master/doc/health-checking.md>`_
// for details.
type HealthCheck_GrpcHealthCheck struct {
	// An optional service name parameter which will be sent to gRPC service in
	// `grpc.health.v1.HealthCheckRequest
	// <https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto#L20>`_.
	// message. See `gRPC health-checking overview
	// <https://github.com/grpc/grpc/blob/master/doc/health-checking.md>`_ for more information.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The value of the :authority header in the gRPC health check request. If
	// left empty (default value), the name of the cluster this health check is associated
	// with will be used.
	Authority            string   `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck_GrpcHealthCheck) Reset()         { *m = HealthCheck_GrpcHealthCheck{} }
func (m *HealthCheck_GrpcHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_GrpcHealthCheck) ProtoMessage()    {}
func (*HealthCheck_GrpcHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 4}
}

func (m *HealthCheck_GrpcHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.Merge(m, src)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_GrpcHealthCheck.Size(m)
}
func (m *HealthCheck_GrpcHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_GrpcHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_GrpcHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_GrpcHealthCheck) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *HealthCheck_GrpcHealthCheck) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// Custom health check.
type HealthCheck_CustomHealthCheck struct {
	// The registered name of the custom health checker.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A custom health checker specific configuration which depends on the custom health checker
	// being instantiated. See :api:`envoy/config/health_checker` for reference.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*HealthCheck_CustomHealthCheck_Config
	//	*HealthCheck_CustomHealthCheck_TypedConfig
	ConfigType           isHealthCheck_CustomHealthCheck_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *HealthCheck_CustomHealthCheck) Reset()         { *m = HealthCheck_CustomHealthCheck{} }
func (m *HealthCheck_CustomHealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck_CustomHealthCheck) ProtoMessage()    {}
func (*HealthCheck_CustomHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6ca44dd529b90bd, []int{0, 5}
}

func (m *HealthCheck_CustomHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.Merge(m, src)
}
func (m *HealthCheck_CustomHealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck_CustomHealthCheck.Size(m)
}
func (m *HealthCheck_CustomHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck_CustomHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck_CustomHealthCheck proto.InternalMessageInfo

func (m *HealthCheck_CustomHealthCheck) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isHealthCheck_CustomHealthCheck_ConfigType interface {
	isHealthCheck_CustomHealthCheck_ConfigType()
}

type HealthCheck_CustomHealthCheck_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type HealthCheck_CustomHealthCheck_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*HealthCheck_CustomHealthCheck_Config) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (*HealthCheck_CustomHealthCheck_TypedConfig) isHealthCheck_CustomHealthCheck_ConfigType() {}

func (m *HealthCheck_CustomHealthCheck) GetConfigType() isHealthCheck_CustomHealthCheck_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_Config); ok {
		return x.Config
	}
	return nil
}

func (m *HealthCheck_CustomHealthCheck) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*HealthCheck_CustomHealthCheck_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheck_CustomHealthCheck) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheck_CustomHealthCheck_Config)(nil),
		(*HealthCheck_CustomHealthCheck_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.api.v2.core.HealthStatus", HealthStatus_name, HealthStatus_value)
	proto.RegisterType((*HealthCheck)(nil), "envoy.api.v2.core.HealthCheck")
	proto.RegisterType((*HealthCheck_Payload)(nil), "envoy.api.v2.core.HealthCheck.Payload")
	proto.RegisterType((*HealthCheck_HttpHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.HttpHealthCheck")
	proto.RegisterType((*HealthCheck_TcpHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.TcpHealthCheck")
	proto.RegisterType((*HealthCheck_RedisHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.RedisHealthCheck")
	proto.RegisterType((*HealthCheck_GrpcHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.GrpcHealthCheck")
	proto.RegisterType((*HealthCheck_CustomHealthCheck)(nil), "envoy.api.v2.core.HealthCheck.CustomHealthCheck")
}

func init() {
	proto.RegisterFile("envoy/api/v2/core/health_check.proto", fileDescriptor_b6ca44dd529b90bd)
}

var fileDescriptor_b6ca44dd529b90bd = []byte{
	// 1209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xcb, 0x6e, 0xdb, 0x46,
	0x17, 0xc7, 0x45, 0x4b, 0xb6, 0xa4, 0x23, 0x59, 0xa6, 0x46, 0x89, 0xcd, 0xe8, 0xf3, 0x97, 0xa8,
	0x85, 0x51, 0x18, 0x01, 0x2a, 0xb5, 0x4a, 0x91, 0x22, 0x5d, 0xc5, 0xb2, 0x15, 0x4b, 0x69, 0xa2,
	0x18, 0x13, 0x39, 0x45, 0x80, 0x16, 0xec, 0x98, 0x1c, 0x53, 0x6c, 0x68, 0x0e, 0x3b, 0x1c, 0x2a,
	0xd1, 0x4b, 0x74, 0xd1, 0x55, 0x1f, 0xa1, 0xc8, 0xa6, 0xdb, 0xa2, 0xab, 0xbc, 0x49, 0x81, 0xae,
	0x9a, 0x3e, 0x42, 0x57, 0xc5, 0x0c, 0x29, 0x59, 0x12, 0x5d, 0xd8, 0x46, 0x77, 0x9c, 0x73, 0xce,
	0xff, 0x37, 0x97, 0x73, 0x01, 0x61, 0x87, 0xfa, 0x63, 0x36, 0x69, 0x91, 0xc0, 0x6d, 0x8d, 0xdb,
	0x2d, 0x8b, 0x71, 0xda, 0x1a, 0x51, 0xe2, 0x89, 0x91, 0x69, 0x8d, 0xa8, 0xf5, 0xaa, 0x19, 0x70,
	0x26, 0x18, 0xaa, 0xaa, 0xa8, 0x26, 0x09, 0xdc, 0xe6, 0xb8, 0xdd, 0x94, 0x51, 0xf5, 0xed, 0xb4,
	0xf0, 0x84, 0x84, 0x34, 0x16, 0xd4, 0x37, 0x63, 0xaf, 0x98, 0x04, 0xb4, 0xc5, 0x89, 0xef, 0x4c,
	0xed, 0xb7, 0x1c, 0xc6, 0x1c, 0x8f, 0xb6, 0xd4, 0xea, 0x24, 0x3a, 0x6d, 0x11, 0x7f, 0x92, 0xb8,
	0x6e, 0x2f, 0xbb, 0xec, 0x88, 0x13, 0xe1, 0x32, 0x3f, 0xf1, 0x6f, 0x2f, 0xfb, 0x43, 0xc1, 0x23,
	0x4b, 0xfc, 0x9b, 0xfa, 0x35, 0x27, 0x41, 0x40, 0x79, 0x98, 0xf8, 0xb7, 0xc6, 0xc4, 0x73, 0x6d,
	0x22, 0x68, 0x6b, 0xfa, 0x91, 0x38, 0x6e, 0x38, 0xcc, 0x61, 0xea, 0xb3, 0x25, 0xbf, 0x62, 0xeb,
	0x87, 0x7f, 0xd7, 0xa0, 0xd4, 0x53, 0xef, 0xb0, 0x2f, 0x9f, 0x01, 0x75, 0x21, 0x2f, 0xdc, 0x33,
	0xca, 0x22, 0x61, 0x68, 0x0d, 0x6d, 0xb7, 0xd4, 0xbe, 0xd5, 0x8c, 0x37, 0x6c, 0x4e, 0x37, 0x6c,
	0x1e, 0x24, 0xc7, 0xed, 0xe8, 0x3f, 0xfd, 0x7e, 0x47, 0xfb, 0xed, 0xfd, 0xbb, 0x6c, 0xfe, 0xad,
	0x96, 0x2b, 0x68, 0x77, 0x33, 0x78, 0xaa, 0x45, 0x87, 0x50, 0x70, 0x7d, 0x41, 0xf9, 0x98, 0x78,
	0xc6, 0xca, 0xf5, 0x39, 0x33, 0x31, 0x7a, 0x08, 0x15, 0xd7, 0x77, 0x85, 0x4b, 0x3c, 0xf3, 0x3b,
	0x57, 0x08, 0xca, 0x8d, 0x1b, 0x97, 0xe0, 0xf0, 0x7a, 0x22, 0x78, 0xac, 0xe2, 0x51, 0x07, 0x36,
	0xa6, 0xb4, 0x29, 0x22, 0x7b, 0x19, 0xa2, 0x32, 0x55, 0x24, 0x8c, 0xfb, 0xb0, 0xb5, 0xc4, 0x30,
	0x03, 0xca, 0x2d, 0xea, 0x0b, 0x03, 0x35, 0xb4, 0xdd, 0x75, 0x7c, 0x73, 0x51, 0x70, 0x14, 0x3b,
	0xd1, 0x53, 0xa8, 0x45, 0x7e, 0x5c, 0x66, 0x13, 0x53, 0x8c, 0x38, 0x0d, 0x47, 0xcc, 0xb3, 0x8d,
	0x9c, 0xda, 0x7f, 0x3b, 0xb5, 0xff, 0x71, 0xdf, 0x17, 0xf7, 0xda, 0x2f, 0x88, 0x17, 0x51, 0x8c,
	0x66, 0xc2, 0xe1, 0x54, 0x87, 0xfa, 0x50, 0x4d, 0xc3, 0x56, 0xaf, 0x00, 0xd3, 0x53, 0xa8, 0xcf,
	0xa1, 0x40, 0x3c, 0x61, 0x06, 0x8c, 0x0b, 0x63, 0xed, 0x0a, 0x84, 0x3c, 0xf1, 0xc4, 0x11, 0xe3,
	0x02, 0x75, 0x41, 0xe7, 0x34, 0x0a, 0xa9, 0x69, 0x31, 0xdf, 0xa7, 0x96, 0x7c, 0x2e, 0x23, 0xaf,
	0x00, 0xf5, 0x14, 0xa0, 0xc3, 0x98, 0x17, 0xcb, 0x37, 0x94, 0x66, 0x7f, 0x26, 0x41, 0x5f, 0x43,
	0x75, 0x24, 0x44, 0x60, 0xce, 0xf7, 0xa0, 0x51, 0x50, 0x9c, 0x66, 0x33, 0xd5, 0x84, 0xcd, 0xb9,
	0x12, 0x6d, 0xf6, 0x84, 0x08, 0xe6, 0xd6, 0xbd, 0x0c, 0xde, 0x18, 0x2d, 0x9a, 0xd0, 0x4b, 0xd0,
	0x85, 0xb5, 0x04, 0x2f, 0x2a, 0xf8, 0xc7, 0x97, 0xc0, 0x87, 0xd6, 0x12, 0xbb, 0x22, 0x16, 0x2c,
	0xf2, 0xe0, 0x0e, 0x0f, 0xac, 0x45, 0x76, 0xe9, 0x4a, 0x07, 0x3f, 0xe4, 0x81, 0xb5, 0x74, 0x70,
	0x67, 0xd1, 0x84, 0x4e, 0xa0, 0x66, 0x45, 0xa1, 0x60, 0x67, 0x8b, 0xfc, 0x75, 0xc5, 0xff, 0xe4,
	0x12, 0xfe, 0xbe, 0x52, 0x2e, 0xee, 0x50, 0xb5, 0x96, 0x8d, 0xe8, 0x18, 0x6a, 0x3e, 0x33, 0x05,
	0x27, 0xa7, 0xa7, 0xae, 0x65, 0xce, 0xda, 0xb4, 0x7c, 0x59, 0x9b, 0x82, 0x6c, 0xd1, 0xd5, 0xb7,
	0xda, 0xca, 0xdd, 0x0c, 0xae, 0xfa, 0x6c, 0x18, 0x03, 0xfa, 0xd3, 0x4e, 0x1d, 0xc2, 0x79, 0xc9,
	0x9e, 0x53, 0x2b, 0xd7, 0xa2, 0xce, 0x00, 0x33, 0xea, 0x37, 0xb0, 0x75, 0x4e, 0xa5, 0xb6, 0x43,
	0xcf, 0xd1, 0x1b, 0xd7, 0x41, 0xdf, 0x9c, 0x51, 0xba, 0xb6, 0x43, 0x67, 0xf8, 0x97, 0x70, 0xf3,
	0x62, 0xb8, 0x7e, 0x1d, 0x78, 0xed, 0x22, 0xf4, 0x0e, 0x54, 0xe8, 0x98, 0xfa, 0xc2, 0xf4, 0x98,
	0x63, 0x06, 0x44, 0x8c, 0x8c, 0x6a, 0x43, 0xdb, 0x2d, 0xe2, 0xb2, 0xb2, 0x3e, 0x61, 0xce, 0x11,
	0x11, 0x23, 0xf4, 0x08, 0x1a, 0xc4, 0x7b, 0x4d, 0x26, 0xa1, 0x0a, 0x9b, 0x4f, 0xba, 0x79, 0x4a,
	0x5c, 0x2f, 0xe2, 0x34, 0x34, 0x6a, 0x0d, 0x6d, 0xb7, 0x80, 0xb7, 0xe3, 0xb8, 0x27, 0xcc, 0x99,
	0x4b, 0xe6, 0xa3, 0x24, 0xa6, 0xfe, 0x02, 0xf2, 0x47, 0x64, 0xe2, 0x31, 0x62, 0xa3, 0x3b, 0x90,
	0x13, 0xf4, 0x4d, 0x3c, 0xbf, 0x8b, 0x9d, 0xa2, 0x3c, 0x67, 0x8e, 0xaf, 0x34, 0xb4, 0x5e, 0x06,
	0x2b, 0x07, 0x32, 0x60, 0xed, 0xc4, 0xf5, 0x09, 0x9f, 0xa8, 0xd1, 0x5c, 0xee, 0x65, 0x70, 0xb2,
	0xee, 0xe8, 0x90, 0x0f, 0x12, 0xca, 0xea, 0xaf, 0xef, 0xdf, 0x65, 0xb5, 0xfa, 0x5f, 0x59, 0xd8,
	0x58, 0x6a, 0x38, 0x84, 0x20, 0x37, 0x62, 0x61, 0xb2, 0x01, 0x56, 0xdf, 0xe8, 0xff, 0x90, 0x53,
	0x77, 0x5c, 0x59, 0xda, 0x14, 0x2b, 0x33, 0xfa, 0x02, 0x72, 0x21, 0xf5, 0xed, 0x64, 0xf2, 0x7e,
	0x74, 0x49, 0x21, 0x27, 0x37, 0xc1, 0x4a, 0x83, 0x1e, 0x42, 0x9e, 0x53, 0x8b, 0xba, 0x63, 0x9a,
	0x0c, 0xce, 0xab, 0xca, 0xa7, 0x32, 0xf4, 0x01, 0x94, 0x43, 0xca, 0xc7, 0xae, 0x45, 0x4d, 0x9f,
	0x9c, 0x51, 0x35, 0x32, 0x8b, 0xb8, 0x94, 0xd8, 0x06, 0xe4, 0x8c, 0xa2, 0x53, 0xd8, 0xe4, 0xf4,
	0xfb, 0x88, 0x86, 0x42, 0x26, 0xc1, 0xa6, 0x3c, 0x34, 0x05, 0x33, 0x89, 0x6d, 0x1b, 0x6b, 0x8d,
	0xec, 0x6e, 0xa9, 0xbd, 0x73, 0xf1, 0x9e, 0x36, 0xe5, 0x6a, 0xc0, 0x3d, 0x0b, 0x54, 0x51, 0x94,
	0xe4, 0xbd, 0xd7, 0x7e, 0xd4, 0xb2, 0xfa, 0x9f, 0x79, 0x5c, 0x4b, 0x80, 0x71, 0x58, 0x38, 0x64,
	0x7b, 0xb6, 0x8d, 0x1e, 0xc0, 0xad, 0x0b, 0xf6, 0xe1, 0xf4, 0x8c, 0x8d, 0xa9, 0x51, 0x68, 0x64,
	0x77, 0x8b, 0x78, 0x73, 0x59, 0x87, 0x95, 0x17, 0xfd, 0x0f, 0x8a, 0x72, 0xee, 0xca, 0x59, 0xd7,
	0x56, 0x23, 0xb7, 0x80, 0x0b, 0x51, 0x48, 0x65, 0x76, 0xda, 0x68, 0x1f, 0xaa, 0xf4, 0x4d, 0x40,
	0x2d, 0x41, 0x6d, 0x33, 0x14, 0x44, 0x44, 0x21, 0x0d, 0x8d, 0xa2, 0x3a, 0xfa, 0x66, 0x72, 0x74,
	0xf9, 0x8f, 0xd2, 0xec, 0xfb, 0xe2, 0xfe, 0x67, 0x58, 0xfe, 0xa8, 0x60, 0x7d, 0x2a, 0x78, 0x9e,
	0xc4, 0xd7, 0x7f, 0xd0, 0xa0, 0xb2, 0x38, 0x00, 0x67, 0x89, 0xd3, 0xfe, 0x5b, 0xe2, 0x56, 0xd4,
	0x49, 0xae, 0x9b, 0xb8, 0xfa, 0x0e, 0xe8, 0x98, 0xda, 0x6e, 0x38, 0x7f, 0x22, 0x1d, 0xb2, 0xaf,
	0xe8, 0x24, 0x29, 0x3e, 0xf9, 0x59, 0xc7, 0xb0, 0xb1, 0x34, 0x5a, 0x53, 0x19, 0xd7, 0xd2, 0x19,
	0xdf, 0x86, 0x22, 0x89, 0xc4, 0x88, 0x71, 0x57, 0xc4, 0x8d, 0x50, 0xc4, 0xe7, 0x86, 0xfa, 0x2f,
	0x1a, 0x54, 0x53, 0xf3, 0x54, 0x56, 0xf9, 0x39, 0x6e, 0xa1, 0xca, 0xa5, 0x19, 0x7d, 0x0a, 0x6b,
	0x16, 0xf3, 0x4f, 0x5d, 0x27, 0xf9, 0xe7, 0xd9, 0x4a, 0x8d, 0x8f, 0xe7, 0xea, 0x57, 0x4e, 0x76,
	0x5c, 0x1c, 0x88, 0x1e, 0x40, 0x59, 0xe6, 0xc5, 0x36, 0x13, 0x61, 0xdc, 0x20, 0x37, 0x52, 0xc2,
	0x3d, 0x7f, 0xd2, 0xcb, 0xe0, 0x92, 0x8a, 0xdd, 0x57, 0xa1, 0x9d, 0x75, 0x28, 0xc5, 0x22, 0x53,
	0x5a, 0x3b, 0x5b, 0x50, 0x99, 0x1f, 0x1f, 0x94, 0x27, 0x2d, 0xfc, 0x38, 0x57, 0x00, 0xbd, 0x84,
	0x11, 0x97, 0x0f, 0xb9, 0x30, 0x61, 0xee, 0x7e, 0x0b, 0xe5, 0xf8, 0x76, 0x71, 0x05, 0xa0, 0x12,
	0xe4, 0x8f, 0x07, 0x5f, 0x0e, 0x9e, 0x7d, 0x35, 0xd0, 0x33, 0x72, 0xd1, 0xeb, 0xee, 0x3d, 0x19,
	0xf6, 0x5e, 0xea, 0x1a, 0x5a, 0x87, 0xe2, 0xf1, 0x60, 0xba, 0x5c, 0x41, 0x65, 0x28, 0x1c, 0xe0,
	0xbd, 0xfe, 0xa0, 0x3f, 0x38, 0xd4, 0xb3, 0x32, 0x72, 0xd8, 0x7f, 0xda, 0x7d, 0x76, 0x3c, 0xd4,
	0x73, 0xca, 0xd5, 0x3d, 0xc4, 0x7b, 0x07, 0xdd, 0x03, 0x7d, 0xb5, 0xf3, 0xe0, 0xe7, 0x3f, 0x6e,
	0x6b, 0x70, 0xc7, 0x65, 0x71, 0xe6, 0x03, 0xce, 0xde, 0x4c, 0xd2, 0x45, 0xd0, 0xd1, 0xe7, 0x1e,
	0xf9, 0x48, 0x5e, 0xf9, 0x48, 0x3b, 0x59, 0x53, 0x77, 0xbf, 0xf7, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4c, 0xab, 0xb8, 0x9f, 0xb9, 0x0b, 0x00, 0x00,
}
