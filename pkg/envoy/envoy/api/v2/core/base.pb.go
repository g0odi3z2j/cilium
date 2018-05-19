// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/base.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _struct "github.com/golang/protobuf/ptypes/struct"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Envoy supports :ref:`upstream priority routing
// <arch_overview_http_routing_priority>` both at the route and the virtual
// cluster level. The current priority implementation uses different connection
// pool and circuit breaking settings for each priority level. This means that
// even for HTTP/2 requests, two physical connections will be used to an
// upstream host. In the future Envoy will likely support true HTTP/2 priority
// over a single upstream connection.
type RoutingPriority int32

const (
	RoutingPriority_DEFAULT RoutingPriority = 0
	RoutingPriority_HIGH    RoutingPriority = 1
)

var RoutingPriority_name = map[int32]string{
	0: "DEFAULT",
	1: "HIGH",
}
var RoutingPriority_value = map[string]int32{
	"DEFAULT": 0,
	"HIGH":    1,
}

func (x RoutingPriority) String() string {
	return proto.EnumName(RoutingPriority_name, int32(x))
}
func (RoutingPriority) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{0}
}

// HTTP request method.
type RequestMethod int32

const (
	RequestMethod_METHOD_UNSPECIFIED RequestMethod = 0
	RequestMethod_GET                RequestMethod = 1
	RequestMethod_HEAD               RequestMethod = 2
	RequestMethod_POST               RequestMethod = 3
	RequestMethod_PUT                RequestMethod = 4
	RequestMethod_DELETE             RequestMethod = 5
	RequestMethod_CONNECT            RequestMethod = 6
	RequestMethod_OPTIONS            RequestMethod = 7
	RequestMethod_TRACE              RequestMethod = 8
)

var RequestMethod_name = map[int32]string{
	0: "METHOD_UNSPECIFIED",
	1: "GET",
	2: "HEAD",
	3: "POST",
	4: "PUT",
	5: "DELETE",
	6: "CONNECT",
	7: "OPTIONS",
	8: "TRACE",
}
var RequestMethod_value = map[string]int32{
	"METHOD_UNSPECIFIED": 0,
	"GET":                1,
	"HEAD":               2,
	"POST":               3,
	"PUT":                4,
	"DELETE":             5,
	"CONNECT":            6,
	"OPTIONS":            7,
	"TRACE":              8,
}

func (x RequestMethod) String() string {
	return proto.EnumName(RequestMethod_name, int32(x))
}
func (RequestMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{1}
}

// Identifies location of where either Envoy runs or where upstream hosts run.
type Locality struct {
	// Region this :ref:`zone <envoy_api_field_core.Locality.zone>` belongs to.
	Region string `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	// Defines the local service zone where Envoy is running. Though optional, it
	// should be set if discovery service routing is used and the discovery
	// service exposes :ref:`zone data <config_cluster_manager_sds_api_host_az>`,
	// either in this message or via :option:`--service-zone`. The meaning of zone
	// is context dependent, e.g. `Availability Zone (AZ)
	// <https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html>`_
	// on AWS, `Zone <https://cloud.google.com/compute/docs/regions-zones/>`_ on
	// GCP, etc.
	Zone string `protobuf:"bytes,2,opt,name=zone" json:"zone,omitempty"`
	// When used for locality of upstream hosts, this field further splits zone
	// into smaller chunks of sub-zones so they can be load balanced
	// independently.
	SubZone              string   `protobuf:"bytes,3,opt,name=sub_zone,json=subZone" json:"sub_zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Locality) Reset()         { *m = Locality{} }
func (m *Locality) String() string { return proto.CompactTextString(m) }
func (*Locality) ProtoMessage()    {}
func (*Locality) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{0}
}
func (m *Locality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Locality.Unmarshal(m, b)
}
func (m *Locality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Locality.Marshal(b, m, deterministic)
}
func (dst *Locality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Locality.Merge(dst, src)
}
func (m *Locality) XXX_Size() int {
	return xxx_messageInfo_Locality.Size(m)
}
func (m *Locality) XXX_DiscardUnknown() {
	xxx_messageInfo_Locality.DiscardUnknown(m)
}

var xxx_messageInfo_Locality proto.InternalMessageInfo

func (m *Locality) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *Locality) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Locality) GetSubZone() string {
	if m != nil {
		return m.SubZone
	}
	return ""
}

// Identifies a specific Envoy instance. The node identifier is presented to the
// management server, which may use this identifier to distinguish per Envoy
// configuration for serving.
type Node struct {
	// An opaque node identifier for the Envoy node. This also provides the local
	// service node name. It should be set if any of the following features are
	// used: :ref:`statsd <arch_overview_statistics>`, :ref:`CDS
	// <config_cluster_manager_cds>`, and :ref:`HTTP tracing
	// <arch_overview_tracing>`, either in this message or via
	// :option:`--service-node`.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Defines the local service cluster name where Envoy is running. Though
	// optional, it should be set if any of the following features are used:
	// :ref:`statsd <arch_overview_statistics>`, :ref:`health check cluster
	// verification <config_cluster_manager_cluster_hc_service_name>`,
	// :ref:`runtime override directory <config_runtime_override_subdirectory>`,
	// :ref:`user agent addition <config_http_conn_man_add_user_agent>`,
	// :ref:`HTTP global rate limiting <config_http_filters_rate_limit>`,
	// :ref:`CDS <config_cluster_manager_cds>`, and :ref:`HTTP tracing
	// <arch_overview_tracing>`, either in this message or via
	// :option:`--service-cluster`.
	Cluster string `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	// Opaque metadata extending the node identifier. Envoy will pass this
	// directly to the management server.
	Metadata *_struct.Struct `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	// Locality specifying where the Envoy instance is running.
	Locality *Locality `protobuf:"bytes,4,opt,name=locality" json:"locality,omitempty"`
	// This is motivated by informing a management server during canary which
	// version of Envoy is being tested in a heterogeneous fleet. This will be set
	// by Envoy in management server RPCs.
	BuildVersion         string   `protobuf:"bytes,5,opt,name=build_version,json=buildVersion" json:"build_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{1}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Node) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Node) GetLocality() *Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *Node) GetBuildVersion() string {
	if m != nil {
		return m.BuildVersion
	}
	return ""
}

// Metadata provides additional inputs to filters based on matched listeners,
// filter chains, routes and endpoints. It is structured as a map from filter
// name (in reverse DNS format) to metadata specific to the filter. Metadata
// key-values for a filter are merged as connection and request handling occurs,
// with later values for the same key overriding earlier values.
//
// An example use of metadata is providing additional values to
// http_connection_manager in the envoy.http_connection_manager.access_log
// namespace.
//
// For load balancing, Metadata provides a means to subset cluster endpoints.
// Endpoints have a Metadata object associated and routes contain a Metadata
// object to match against. There are some well defined metadata used today for
// this purpose:
//
// * ``{"envoy.lb": {"canary": <bool> }}`` This indicates the canary status of an
//   endpoint and is also used during header processing
//   (x-envoy-upstream-canary) and for stats purposes.
type Metadata struct {
	// Key is the reverse DNS filter name, e.g. com.acme.widget. The envoy.*
	// namespace is reserved for Envoy's built-in filters.
	FilterMetadata       map[string]*_struct.Struct `protobuf:"bytes,1,rep,name=filter_metadata,json=filterMetadata" json:"filter_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{2}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetFilterMetadata() map[string]*_struct.Struct {
	if m != nil {
		return m.FilterMetadata
	}
	return nil
}

// Runtime derived uint32 with a default when not specified.
type RuntimeUInt32 struct {
	// Default value if runtime value is not available.
	DefaultValue uint32 `protobuf:"varint,2,opt,name=default_value,json=defaultValue" json:"default_value,omitempty"`
	// Runtime key to get value for comparison. This value is used if defined.
	RuntimeKey           string   `protobuf:"bytes,3,opt,name=runtime_key,json=runtimeKey" json:"runtime_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RuntimeUInt32) Reset()         { *m = RuntimeUInt32{} }
func (m *RuntimeUInt32) String() string { return proto.CompactTextString(m) }
func (*RuntimeUInt32) ProtoMessage()    {}
func (*RuntimeUInt32) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{3}
}
func (m *RuntimeUInt32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RuntimeUInt32.Unmarshal(m, b)
}
func (m *RuntimeUInt32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RuntimeUInt32.Marshal(b, m, deterministic)
}
func (dst *RuntimeUInt32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RuntimeUInt32.Merge(dst, src)
}
func (m *RuntimeUInt32) XXX_Size() int {
	return xxx_messageInfo_RuntimeUInt32.Size(m)
}
func (m *RuntimeUInt32) XXX_DiscardUnknown() {
	xxx_messageInfo_RuntimeUInt32.DiscardUnknown(m)
}

var xxx_messageInfo_RuntimeUInt32 proto.InternalMessageInfo

func (m *RuntimeUInt32) GetDefaultValue() uint32 {
	if m != nil {
		return m.DefaultValue
	}
	return 0
}

func (m *RuntimeUInt32) GetRuntimeKey() string {
	if m != nil {
		return m.RuntimeKey
	}
	return ""
}

// Header name/value pair.
type HeaderValue struct {
	// Header name.
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Header value.
	//
	// The same :ref:`format specifier <config_access_log_format>` as used for
	// :ref:`HTTP access logging <config_access_log>` applies here, however
	// unknown header values are replaced with the empty string instead of `-`.
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderValue) Reset()         { *m = HeaderValue{} }
func (m *HeaderValue) String() string { return proto.CompactTextString(m) }
func (*HeaderValue) ProtoMessage()    {}
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{4}
}
func (m *HeaderValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderValue.Unmarshal(m, b)
}
func (m *HeaderValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderValue.Marshal(b, m, deterministic)
}
func (dst *HeaderValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValue.Merge(dst, src)
}
func (m *HeaderValue) XXX_Size() int {
	return xxx_messageInfo_HeaderValue.Size(m)
}
func (m *HeaderValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValue.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValue proto.InternalMessageInfo

func (m *HeaderValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HeaderValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Header name/value pair plus option to control append behavior.
type HeaderValueOption struct {
	// Header name/value pair that this option applies to.
	Header *HeaderValue `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Should the value be appended? If true (default), the value is appended to
	// existing values.
	Append               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=append" json:"append,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HeaderValueOption) Reset()         { *m = HeaderValueOption{} }
func (m *HeaderValueOption) String() string { return proto.CompactTextString(m) }
func (*HeaderValueOption) ProtoMessage()    {}
func (*HeaderValueOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{5}
}
func (m *HeaderValueOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderValueOption.Unmarshal(m, b)
}
func (m *HeaderValueOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderValueOption.Marshal(b, m, deterministic)
}
func (dst *HeaderValueOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderValueOption.Merge(dst, src)
}
func (m *HeaderValueOption) XXX_Size() int {
	return xxx_messageInfo_HeaderValueOption.Size(m)
}
func (m *HeaderValueOption) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderValueOption.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderValueOption proto.InternalMessageInfo

func (m *HeaderValueOption) GetHeader() *HeaderValue {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HeaderValueOption) GetAppend() *wrappers.BoolValue {
	if m != nil {
		return m.Append
	}
	return nil
}

// Data source consisting of either a file or an inline value.
type DataSource struct {
	// Types that are valid to be assigned to Specifier:
	//	*DataSource_Filename
	//	*DataSource_InlineBytes
	//	*DataSource_InlineString
	Specifier            isDataSource_Specifier `protobuf_oneof:"specifier"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DataSource) Reset()         { *m = DataSource{} }
func (m *DataSource) String() string { return proto.CompactTextString(m) }
func (*DataSource) ProtoMessage()    {}
func (*DataSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{6}
}
func (m *DataSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSource.Unmarshal(m, b)
}
func (m *DataSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSource.Marshal(b, m, deterministic)
}
func (dst *DataSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSource.Merge(dst, src)
}
func (m *DataSource) XXX_Size() int {
	return xxx_messageInfo_DataSource.Size(m)
}
func (m *DataSource) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSource.DiscardUnknown(m)
}

var xxx_messageInfo_DataSource proto.InternalMessageInfo

type isDataSource_Specifier interface {
	isDataSource_Specifier()
}

type DataSource_Filename struct {
	Filename string `protobuf:"bytes,1,opt,name=filename,oneof"`
}
type DataSource_InlineBytes struct {
	InlineBytes []byte `protobuf:"bytes,2,opt,name=inline_bytes,json=inlineBytes,proto3,oneof"`
}
type DataSource_InlineString struct {
	InlineString string `protobuf:"bytes,3,opt,name=inline_string,json=inlineString,oneof"`
}

func (*DataSource_Filename) isDataSource_Specifier()     {}
func (*DataSource_InlineBytes) isDataSource_Specifier()  {}
func (*DataSource_InlineString) isDataSource_Specifier() {}

func (m *DataSource) GetSpecifier() isDataSource_Specifier {
	if m != nil {
		return m.Specifier
	}
	return nil
}

func (m *DataSource) GetFilename() string {
	if x, ok := m.GetSpecifier().(*DataSource_Filename); ok {
		return x.Filename
	}
	return ""
}

func (m *DataSource) GetInlineBytes() []byte {
	if x, ok := m.GetSpecifier().(*DataSource_InlineBytes); ok {
		return x.InlineBytes
	}
	return nil
}

func (m *DataSource) GetInlineString() string {
	if x, ok := m.GetSpecifier().(*DataSource_InlineString); ok {
		return x.InlineString
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DataSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DataSource_OneofMarshaler, _DataSource_OneofUnmarshaler, _DataSource_OneofSizer, []interface{}{
		(*DataSource_Filename)(nil),
		(*DataSource_InlineBytes)(nil),
		(*DataSource_InlineString)(nil),
	}
}

func _DataSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DataSource)
	// specifier
	switch x := m.Specifier.(type) {
	case *DataSource_Filename:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Filename)
	case *DataSource_InlineBytes:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.InlineBytes)
	case *DataSource_InlineString:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.InlineString)
	case nil:
	default:
		return fmt.Errorf("DataSource.Specifier has unexpected type %T", x)
	}
	return nil
}

func _DataSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DataSource)
	switch tag {
	case 1: // specifier.filename
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Specifier = &DataSource_Filename{x}
		return true, err
	case 2: // specifier.inline_bytes
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Specifier = &DataSource_InlineBytes{x}
		return true, err
	case 3: // specifier.inline_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Specifier = &DataSource_InlineString{x}
		return true, err
	default:
		return false, nil
	}
}

func _DataSource_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DataSource)
	// specifier
	switch x := m.Specifier.(type) {
	case *DataSource_Filename:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Filename)))
		n += len(x.Filename)
	case *DataSource_InlineBytes:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.InlineBytes)))
		n += len(x.InlineBytes)
	case *DataSource_InlineString:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.InlineString)))
		n += len(x.InlineString)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Configuration for transport socket in :ref:`listeners <config_listeners>` and
// :ref:`clusters <config_cluster_manager_cluster>`. If the configuration is
// empty, a default transport socket implementation and configuration will be
// chosen based on the platform and existence of tls_context.
type TransportSocket struct {
	// The name of the transport socket to instantiate. The name must match a supported transport
	// socket implementation.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Implementation specific configuration which depends on the implementation being instantiated.
	// See the supported transport socket implementations for further documentation.
	Config               *_struct.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TransportSocket) Reset()         { *m = TransportSocket{} }
func (m *TransportSocket) String() string { return proto.CompactTextString(m) }
func (*TransportSocket) ProtoMessage()    {}
func (*TransportSocket) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_797415309f0aa448, []int{7}
}
func (m *TransportSocket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransportSocket.Unmarshal(m, b)
}
func (m *TransportSocket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransportSocket.Marshal(b, m, deterministic)
}
func (dst *TransportSocket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransportSocket.Merge(dst, src)
}
func (m *TransportSocket) XXX_Size() int {
	return xxx_messageInfo_TransportSocket.Size(m)
}
func (m *TransportSocket) XXX_DiscardUnknown() {
	xxx_messageInfo_TransportSocket.DiscardUnknown(m)
}

var xxx_messageInfo_TransportSocket proto.InternalMessageInfo

func (m *TransportSocket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TransportSocket) GetConfig() *_struct.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Locality)(nil), "envoy.api.v2.core.Locality")
	proto.RegisterType((*Node)(nil), "envoy.api.v2.core.Node")
	proto.RegisterType((*Metadata)(nil), "envoy.api.v2.core.Metadata")
	proto.RegisterMapType((map[string]*_struct.Struct)(nil), "envoy.api.v2.core.Metadata.FilterMetadataEntry")
	proto.RegisterType((*RuntimeUInt32)(nil), "envoy.api.v2.core.RuntimeUInt32")
	proto.RegisterType((*HeaderValue)(nil), "envoy.api.v2.core.HeaderValue")
	proto.RegisterType((*HeaderValueOption)(nil), "envoy.api.v2.core.HeaderValueOption")
	proto.RegisterType((*DataSource)(nil), "envoy.api.v2.core.DataSource")
	proto.RegisterType((*TransportSocket)(nil), "envoy.api.v2.core.TransportSocket")
	proto.RegisterEnum("envoy.api.v2.core.RoutingPriority", RoutingPriority_name, RoutingPriority_value)
	proto.RegisterEnum("envoy.api.v2.core.RequestMethod", RequestMethod_name, RequestMethod_value)
}

func init() { proto.RegisterFile("envoy/api/v2/core/base.proto", fileDescriptor_base_797415309f0aa448) }

var fileDescriptor_base_797415309f0aa448 = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x8f, 0x13, 0xe7, 0xdf, 0x4b, 0xb3, 0xf5, 0x0e, 0xab, 0xdd, 0x12, 0x96, 0xaa, 0x32, 0x07,
	0xaa, 0x4a, 0xd8, 0x28, 0x15, 0x7f, 0xc4, 0xad, 0x69, 0xdc, 0x4d, 0x44, 0x9b, 0x04, 0xc7, 0x5d,
	0xa1, 0x5e, 0xc2, 0xc4, 0x9e, 0x78, 0x47, 0xeb, 0x7a, 0xcc, 0x78, 0x1c, 0x94, 0xbd, 0x70, 0xe1,
	0xc0, 0xf7, 0xe0, 0x82, 0xc4, 0x17, 0x40, 0x9c, 0x2a, 0xf1, 0x4d, 0xb8, 0xed, 0xb7, 0x40, 0x33,
	0x9e, 0x94, 0x42, 0x2b, 0xf6, 0xf6, 0xe6, 0xf7, 0xe7, 0xcd, 0xef, 0x3d, 0x8f, 0xe1, 0x39, 0x49,
	0xd7, 0x6c, 0xe3, 0xe2, 0x8c, 0xba, 0xeb, 0xbe, 0x1b, 0x32, 0x4e, 0xdc, 0x25, 0xce, 0x89, 0x93,
	0x71, 0x26, 0x18, 0x7a, 0xac, 0x58, 0x07, 0x67, 0xd4, 0x59, 0xf7, 0x1d, 0xc9, 0xf6, 0x9e, 0xc7,
	0x8c, 0xc5, 0x09, 0x71, 0x95, 0x60, 0x59, 0xac, 0xdc, 0x5c, 0xf0, 0x22, 0x14, 0xa5, 0xa1, 0xb7,
	0xff, 0x5f, 0xf6, 0x07, 0x8e, 0xb3, 0x8c, 0xf0, 0x5c, 0xf3, 0xcf, 0xd6, 0x38, 0xa1, 0x11, 0x16,
	0xc4, 0xdd, 0x16, 0x9a, 0x78, 0x12, 0xb3, 0x98, 0xa9, 0xd2, 0x95, 0x55, 0x89, 0xda, 0xdf, 0x40,
	0xeb, 0x9c, 0x85, 0x38, 0xa1, 0x62, 0x83, 0x9e, 0x42, 0x83, 0x93, 0x98, 0xb2, 0x74, 0xcf, 0x38,
	0x30, 0x0e, 0xdb, 0xbe, 0x3e, 0x21, 0x04, 0xe6, 0x1b, 0x96, 0x92, 0xbd, 0xaa, 0x42, 0x55, 0x8d,
	0xde, 0x87, 0x56, 0x5e, 0x2c, 0x17, 0x0a, 0xaf, 0x29, 0xbc, 0x99, 0x17, 0xcb, 0x2b, 0x96, 0x12,
	0xfb, 0x4f, 0x03, 0xcc, 0x09, 0x8b, 0x08, 0x7a, 0x04, 0x55, 0x1a, 0xe9, 0x5e, 0x55, 0x1a, 0xa1,
	0x3d, 0x68, 0x86, 0x49, 0x91, 0x0b, 0xc2, 0x75, 0xab, 0xed, 0x11, 0x1d, 0x43, 0xeb, 0x9a, 0x08,
	0x1c, 0x61, 0x81, 0x55, 0xb7, 0x4e, 0xff, 0x99, 0x53, 0xce, 0xe9, 0x6c, 0xe7, 0x74, 0xe6, 0x6a,
	0x0b, 0xfe, 0xad, 0x10, 0x7d, 0x01, 0xad, 0x44, 0x47, 0xdf, 0x33, 0x95, 0xe9, 0x03, 0xe7, 0xde,
	0x36, 0x9d, 0xed, 0x74, 0xfe, 0xad, 0x18, 0x7d, 0x04, 0xdd, 0x65, 0x41, 0x93, 0x68, 0xb1, 0x26,
	0x3c, 0x97, 0xe3, 0xd6, 0x55, 0x9a, 0x1d, 0x05, 0xbe, 0x2c, 0x31, 0xfb, 0xc6, 0x80, 0xd6, 0xc5,
	0xf6, 0xaa, 0x6f, 0x61, 0x77, 0x45, 0x13, 0x41, 0xf8, 0xe2, 0x36, 0xa6, 0x71, 0x50, 0x3b, 0xec,
	0xf4, 0xdd, 0x07, 0x6e, 0xdc, 0xba, 0x9c, 0x33, 0x65, 0xd9, 0x1e, 0xbd, 0x54, 0xf0, 0x8d, 0xff,
	0x68, 0xf5, 0x2f, 0xb0, 0x77, 0x05, 0xef, 0x3d, 0x20, 0x43, 0x16, 0xd4, 0x5e, 0x93, 0x8d, 0xde,
	0x9d, 0x2c, 0xd1, 0x27, 0x50, 0x5f, 0xe3, 0xa4, 0x28, 0xbf, 0xc2, 0xff, 0xec, 0xa7, 0x54, 0x7d,
	0x55, 0xfd, 0xd2, 0xb0, 0xbf, 0x83, 0xae, 0x5f, 0xa4, 0x82, 0x5e, 0x93, 0xcb, 0x71, 0x2a, 0x8e,
	0xfb, 0x72, 0xf0, 0x88, 0xac, 0x70, 0x91, 0x88, 0xc5, 0x3f, 0xbd, 0xba, 0xfe, 0x8e, 0x06, 0x5f,
	0x4a, 0x0c, 0x1d, 0x41, 0x87, 0x97, 0xae, 0x85, 0x8c, 0xa0, 0x3e, 0xee, 0xa0, 0xfd, 0xc7, 0xdb,
	0x9b, 0x9a, 0xc9, 0xab, 0x07, 0x86, 0x0f, 0x9a, 0xfd, 0x9a, 0x6c, 0xec, 0xcf, 0xa0, 0x33, 0x22,
	0x38, 0x22, 0xbc, 0xb4, 0xde, 0x4f, 0xfd, 0xe4, 0x6e, 0xea, 0xb6, 0x0e, 0x67, 0xff, 0x08, 0x8f,
	0xef, 0xd8, 0xa6, 0x99, 0x90, 0xaf, 0xec, 0x73, 0x68, 0xbc, 0x52, 0xa0, 0xf2, 0x77, 0xfa, 0xfb,
	0x0f, 0xac, 0xf6, 0x8e, 0xcb, 0xd7, 0x6a, 0xd4, 0x87, 0x86, 0xfc, 0x01, 0xd2, 0x48, 0x6f, 0xa6,
	0x77, 0x6f, 0x33, 0x03, 0xc6, 0x12, 0xed, 0x29, 0x95, 0xf6, 0x6f, 0x06, 0xc0, 0x10, 0x0b, 0x3c,
	0x67, 0x05, 0x0f, 0x09, 0xfa, 0x18, 0x5a, 0x2b, 0x9a, 0x90, 0x14, 0x5f, 0x93, 0x32, 0xfc, 0x9d,
	0x79, 0x47, 0x15, 0xff, 0x96, 0x44, 0x0e, 0xec, 0xd0, 0x34, 0xa1, 0x29, 0x59, 0x2c, 0x37, 0x82,
	0xe4, 0xea, 0xc6, 0x1d, 0x2d, 0x7e, 0x53, 0xb5, 0xa4, 0xb8, 0x53, 0x0a, 0x06, 0x92, 0x47, 0x9f,
	0x42, 0x57, 0xeb, 0x73, 0xc1, 0x69, 0x1a, 0xdf, 0xdb, 0xe6, 0xa8, 0xe2, 0xeb, 0x8e, 0x73, 0x25,
	0x18, 0x20, 0x68, 0xe7, 0x19, 0x09, 0xe9, 0x8a, 0x12, 0x8e, 0xea, 0xbf, 0xbf, 0xbd, 0xa9, 0x19,
	0x36, 0x86, 0xdd, 0x80, 0xe3, 0x34, 0xcf, 0x18, 0x17, 0x73, 0x16, 0xbe, 0x26, 0x02, 0x7d, 0x08,
	0xe6, 0x83, 0x69, 0x7d, 0x05, 0x23, 0x17, 0x1a, 0x21, 0x4b, 0x57, 0x34, 0x7e, 0xd7, 0x6b, 0xd1,
	0xb2, 0xa3, 0x43, 0xd8, 0xf5, 0x59, 0x21, 0x68, 0x1a, 0xcf, 0x38, 0x65, 0x5c, 0xfe, 0x25, 0x1d,
	0x68, 0x0e, 0xbd, 0xb3, 0x93, 0xcb, 0xf3, 0xc0, 0xaa, 0xa0, 0x16, 0x98, 0xa3, 0xf1, 0x8b, 0x91,
	0x65, 0x1c, 0xfd, 0x64, 0x40, 0xd7, 0x27, 0xdf, 0x17, 0x24, 0x17, 0x17, 0x44, 0xbc, 0x62, 0x11,
	0x7a, 0x0a, 0xe8, 0xc2, 0x0b, 0x46, 0xd3, 0xe1, 0xe2, 0x72, 0x32, 0x9f, 0x79, 0xa7, 0xe3, 0xb3,
	0xb1, 0x37, 0xb4, 0x2a, 0xa8, 0x09, 0xb5, 0x17, 0x5e, 0x60, 0x19, 0xca, 0xec, 0x9d, 0x0c, 0xad,
	0xaa, 0xac, 0x66, 0xd3, 0x79, 0x60, 0xd5, 0x24, 0x39, 0xbb, 0x0c, 0x2c, 0x13, 0x01, 0x34, 0x86,
	0xde, 0xb9, 0x17, 0x78, 0x56, 0x5d, 0x5e, 0x79, 0x3a, 0x9d, 0x4c, 0xbc, 0xd3, 0xc0, 0x6a, 0xc8,
	0xc3, 0x74, 0x16, 0x8c, 0xa7, 0x93, 0xb9, 0xd5, 0x44, 0x6d, 0xa8, 0x07, 0xfe, 0xc9, 0xa9, 0x67,
	0xb5, 0x7a, 0xe6, 0xcf, 0xbf, 0xec, 0x57, 0x06, 0xf0, 0xeb, 0x5f, 0xfb, 0xc6, 0x95, 0x29, 0x5f,
	0xc5, 0xb2, 0xa1, 0xa6, 0x3a, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xcc, 0x88, 0xee, 0x6a,
	0x05, 0x00, 0x00,
}
