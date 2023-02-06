// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.4
// source: envoy/config/endpoint/v3/load_report.proto

package endpointv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// These are stats Envoy reports to the management server at a frequency defined by
// :ref:`LoadStatsResponse.load_reporting_interval<envoy_v3_api_field_service.load_stats.v3.LoadStatsResponse.load_reporting_interval>`.
// Stats per upstream region/zone and optionally per subzone.
// [#next-free-field: 9]
type UpstreamLocalityStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of zone, region and optionally endpoint group these metrics were
	// collected from. Zone and region names could be empty if unknown.
	Locality *v3.Locality `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	// The total number of requests successfully completed by the endpoints in the
	// locality.
	TotalSuccessfulRequests uint64 `protobuf:"varint,2,opt,name=total_successful_requests,json=totalSuccessfulRequests,proto3" json:"total_successful_requests,omitempty"`
	// The total number of unfinished requests
	TotalRequestsInProgress uint64 `protobuf:"varint,3,opt,name=total_requests_in_progress,json=totalRequestsInProgress,proto3" json:"total_requests_in_progress,omitempty"`
	// The total number of requests that failed due to errors at the endpoint,
	// aggregated over all endpoints in the locality.
	TotalErrorRequests uint64 `protobuf:"varint,4,opt,name=total_error_requests,json=totalErrorRequests,proto3" json:"total_error_requests,omitempty"`
	// The total number of requests that were issued by this Envoy since
	// the last report. This information is aggregated over all the
	// upstream endpoints in the locality.
	TotalIssuedRequests uint64 `protobuf:"varint,8,opt,name=total_issued_requests,json=totalIssuedRequests,proto3" json:"total_issued_requests,omitempty"`
	// Stats for multi-dimensional load balancing.
	LoadMetricStats []*EndpointLoadMetricStats `protobuf:"bytes,5,rep,name=load_metric_stats,json=loadMetricStats,proto3" json:"load_metric_stats,omitempty"`
	// Endpoint granularity stats information for this locality. This information
	// is populated if the Server requests it by setting
	// :ref:`LoadStatsResponse.report_endpoint_granularity<envoy_v3_api_field_service.load_stats.v3.LoadStatsResponse.report_endpoint_granularity>`.
	UpstreamEndpointStats []*UpstreamEndpointStats `protobuf:"bytes,7,rep,name=upstream_endpoint_stats,json=upstreamEndpointStats,proto3" json:"upstream_endpoint_stats,omitempty"`
	// [#not-implemented-hide:] The priority of the endpoint group these metrics
	// were collected from.
	Priority uint32 `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *UpstreamLocalityStats) Reset() {
	*x = UpstreamLocalityStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamLocalityStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamLocalityStats) ProtoMessage() {}

func (x *UpstreamLocalityStats) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamLocalityStats.ProtoReflect.Descriptor instead.
func (*UpstreamLocalityStats) Descriptor() ([]byte, []int) {
	return file_envoy_config_endpoint_v3_load_report_proto_rawDescGZIP(), []int{0}
}

func (x *UpstreamLocalityStats) GetLocality() *v3.Locality {
	if x != nil {
		return x.Locality
	}
	return nil
}

func (x *UpstreamLocalityStats) GetTotalSuccessfulRequests() uint64 {
	if x != nil {
		return x.TotalSuccessfulRequests
	}
	return 0
}

func (x *UpstreamLocalityStats) GetTotalRequestsInProgress() uint64 {
	if x != nil {
		return x.TotalRequestsInProgress
	}
	return 0
}

func (x *UpstreamLocalityStats) GetTotalErrorRequests() uint64 {
	if x != nil {
		return x.TotalErrorRequests
	}
	return 0
}

func (x *UpstreamLocalityStats) GetTotalIssuedRequests() uint64 {
	if x != nil {
		return x.TotalIssuedRequests
	}
	return 0
}

func (x *UpstreamLocalityStats) GetLoadMetricStats() []*EndpointLoadMetricStats {
	if x != nil {
		return x.LoadMetricStats
	}
	return nil
}

func (x *UpstreamLocalityStats) GetUpstreamEndpointStats() []*UpstreamEndpointStats {
	if x != nil {
		return x.UpstreamEndpointStats
	}
	return nil
}

func (x *UpstreamLocalityStats) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

// [#next-free-field: 8]
type UpstreamEndpointStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upstream host address.
	Address *v3.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Opaque and implementation dependent metadata of the
	// endpoint. Envoy will pass this directly to the management server.
	Metadata *structpb.Struct `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The total number of requests successfully completed by the endpoints in the
	// locality. These include non-5xx responses for HTTP, where errors
	// originate at the client and the endpoint responded successfully. For gRPC,
	// the grpc-status values are those not covered by total_error_requests below.
	TotalSuccessfulRequests uint64 `protobuf:"varint,2,opt,name=total_successful_requests,json=totalSuccessfulRequests,proto3" json:"total_successful_requests,omitempty"`
	// The total number of unfinished requests for this endpoint.
	TotalRequestsInProgress uint64 `protobuf:"varint,3,opt,name=total_requests_in_progress,json=totalRequestsInProgress,proto3" json:"total_requests_in_progress,omitempty"`
	// The total number of requests that failed due to errors at the endpoint.
	// For HTTP these are responses with 5xx status codes and for gRPC the
	// grpc-status values:
	//
	//   - DeadlineExceeded
	//   - Unimplemented
	//   - Internal
	//   - Unavailable
	//   - Unknown
	//   - DataLoss
	TotalErrorRequests uint64 `protobuf:"varint,4,opt,name=total_error_requests,json=totalErrorRequests,proto3" json:"total_error_requests,omitempty"`
	// The total number of requests that were issued to this endpoint
	// since the last report. A single TCP connection, HTTP or gRPC
	// request or stream is counted as one request.
	TotalIssuedRequests uint64 `protobuf:"varint,7,opt,name=total_issued_requests,json=totalIssuedRequests,proto3" json:"total_issued_requests,omitempty"`
	// Stats for multi-dimensional load balancing.
	LoadMetricStats []*EndpointLoadMetricStats `protobuf:"bytes,5,rep,name=load_metric_stats,json=loadMetricStats,proto3" json:"load_metric_stats,omitempty"`
}

func (x *UpstreamEndpointStats) Reset() {
	*x = UpstreamEndpointStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpstreamEndpointStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpstreamEndpointStats) ProtoMessage() {}

func (x *UpstreamEndpointStats) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpstreamEndpointStats.ProtoReflect.Descriptor instead.
func (*UpstreamEndpointStats) Descriptor() ([]byte, []int) {
	return file_envoy_config_endpoint_v3_load_report_proto_rawDescGZIP(), []int{1}
}

func (x *UpstreamEndpointStats) GetAddress() *v3.Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *UpstreamEndpointStats) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpstreamEndpointStats) GetTotalSuccessfulRequests() uint64 {
	if x != nil {
		return x.TotalSuccessfulRequests
	}
	return 0
}

func (x *UpstreamEndpointStats) GetTotalRequestsInProgress() uint64 {
	if x != nil {
		return x.TotalRequestsInProgress
	}
	return 0
}

func (x *UpstreamEndpointStats) GetTotalErrorRequests() uint64 {
	if x != nil {
		return x.TotalErrorRequests
	}
	return 0
}

func (x *UpstreamEndpointStats) GetTotalIssuedRequests() uint64 {
	if x != nil {
		return x.TotalIssuedRequests
	}
	return 0
}

func (x *UpstreamEndpointStats) GetLoadMetricStats() []*EndpointLoadMetricStats {
	if x != nil {
		return x.LoadMetricStats
	}
	return nil
}

type EndpointLoadMetricStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the metric; may be empty.
	MetricName string `protobuf:"bytes,1,opt,name=metric_name,json=metricName,proto3" json:"metric_name,omitempty"`
	// Number of calls that finished and included this metric.
	NumRequestsFinishedWithMetric uint64 `protobuf:"varint,2,opt,name=num_requests_finished_with_metric,json=numRequestsFinishedWithMetric,proto3" json:"num_requests_finished_with_metric,omitempty"`
	// Sum of metric values across all calls that finished with this metric for
	// load_reporting_interval.
	TotalMetricValue float64 `protobuf:"fixed64,3,opt,name=total_metric_value,json=totalMetricValue,proto3" json:"total_metric_value,omitempty"`
}

func (x *EndpointLoadMetricStats) Reset() {
	*x = EndpointLoadMetricStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointLoadMetricStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointLoadMetricStats) ProtoMessage() {}

func (x *EndpointLoadMetricStats) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointLoadMetricStats.ProtoReflect.Descriptor instead.
func (*EndpointLoadMetricStats) Descriptor() ([]byte, []int) {
	return file_envoy_config_endpoint_v3_load_report_proto_rawDescGZIP(), []int{2}
}

func (x *EndpointLoadMetricStats) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *EndpointLoadMetricStats) GetNumRequestsFinishedWithMetric() uint64 {
	if x != nil {
		return x.NumRequestsFinishedWithMetric
	}
	return 0
}

func (x *EndpointLoadMetricStats) GetTotalMetricValue() float64 {
	if x != nil {
		return x.TotalMetricValue
	}
	return 0
}

// Per cluster load stats. Envoy reports these stats a management server in a
// :ref:`LoadStatsRequest<envoy_v3_api_msg_service.load_stats.v3.LoadStatsRequest>`
// Next ID: 7
// [#next-free-field: 7]
type ClusterStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the cluster.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// The eds_cluster_config service_name of the cluster.
	// It's possible that two clusters send the same service_name to EDS,
	// in that case, the management server is supposed to do aggregation on the load reports.
	ClusterServiceName string `protobuf:"bytes,6,opt,name=cluster_service_name,json=clusterServiceName,proto3" json:"cluster_service_name,omitempty"`
	// Need at least one.
	UpstreamLocalityStats []*UpstreamLocalityStats `protobuf:"bytes,2,rep,name=upstream_locality_stats,json=upstreamLocalityStats,proto3" json:"upstream_locality_stats,omitempty"`
	// Cluster-level stats such as total_successful_requests may be computed by
	// summing upstream_locality_stats. In addition, below there are additional
	// cluster-wide stats.
	//
	// The total number of dropped requests. This covers requests
	// deliberately dropped by the drop_overload policy and circuit breaking.
	TotalDroppedRequests uint64 `protobuf:"varint,3,opt,name=total_dropped_requests,json=totalDroppedRequests,proto3" json:"total_dropped_requests,omitempty"`
	// Information about deliberately dropped requests for each category specified
	// in the DropOverload policy.
	DroppedRequests []*ClusterStats_DroppedRequests `protobuf:"bytes,5,rep,name=dropped_requests,json=droppedRequests,proto3" json:"dropped_requests,omitempty"`
	// Period over which the actual load report occurred. This will be guaranteed to include every
	// request reported. Due to system load and delays between the *LoadStatsRequest* sent from Envoy
	// and the *LoadStatsResponse* message sent from the management server, this may be longer than
	// the requested load reporting interval in the *LoadStatsResponse*.
	LoadReportInterval *durationpb.Duration `protobuf:"bytes,4,opt,name=load_report_interval,json=loadReportInterval,proto3" json:"load_report_interval,omitempty"`
}

func (x *ClusterStats) Reset() {
	*x = ClusterStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStats) ProtoMessage() {}

func (x *ClusterStats) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStats.ProtoReflect.Descriptor instead.
func (*ClusterStats) Descriptor() ([]byte, []int) {
	return file_envoy_config_endpoint_v3_load_report_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterStats) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ClusterStats) GetClusterServiceName() string {
	if x != nil {
		return x.ClusterServiceName
	}
	return ""
}

func (x *ClusterStats) GetUpstreamLocalityStats() []*UpstreamLocalityStats {
	if x != nil {
		return x.UpstreamLocalityStats
	}
	return nil
}

func (x *ClusterStats) GetTotalDroppedRequests() uint64 {
	if x != nil {
		return x.TotalDroppedRequests
	}
	return 0
}

func (x *ClusterStats) GetDroppedRequests() []*ClusterStats_DroppedRequests {
	if x != nil {
		return x.DroppedRequests
	}
	return nil
}

func (x *ClusterStats) GetLoadReportInterval() *durationpb.Duration {
	if x != nil {
		return x.LoadReportInterval
	}
	return nil
}

type ClusterStats_DroppedRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier for the policy specifying the drop.
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	// Total number of deliberately dropped requests for the category.
	DroppedCount uint64 `protobuf:"varint,2,opt,name=dropped_count,json=droppedCount,proto3" json:"dropped_count,omitempty"`
}

func (x *ClusterStats_DroppedRequests) Reset() {
	*x = ClusterStats_DroppedRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterStats_DroppedRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStats_DroppedRequests) ProtoMessage() {}

func (x *ClusterStats_DroppedRequests) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_endpoint_v3_load_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStats_DroppedRequests.ProtoReflect.Descriptor instead.
func (*ClusterStats_DroppedRequests) Descriptor() ([]byte, []int) {
	return file_envoy_config_endpoint_v3_load_report_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ClusterStats_DroppedRequests) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ClusterStats_DroppedRequests) GetDroppedCount() uint64 {
	if x != nil {
		return x.DroppedCount
	}
	return 0
}

var File_envoy_config_endpoint_v3_load_report_proto protoreflect.FileDescriptor

var file_envoy_config_endpoint_v3_load_report_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x04, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x19, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x5d, 0x0a, 0x11, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x2e,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0f, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x67, 0x0a, 0x17, 0x75, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x15, 0x75, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x3a, 0x32, 0x9a,
	0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x22, 0xf7, 0x03, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3a, 0x0a, 0x19, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x5d, 0x0a, 0x11, 0x6c, 0x6f, 0x61, 0x64,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x0f, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x3a, 0x32, 0x9a, 0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x17,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x21, 0x6e, 0x75, 0x6d, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x1d, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x34, 0x9a, 0xc5, 0x88, 0x1e, 0x2f, 0x0a, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0x89, 0x05, 0x0a, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x71, 0x0a, 0x17, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76,
	0x33, 0x2e, 0x55, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x15, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44,
	0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x61,
	0x0a, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x44, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x52, 0x0f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x4b, 0x0a, 0x14, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x1a, 0x96,
	0x01, 0x0a, 0x0f, 0x44, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x72, 0x6f, 0x70, 0x70,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x39, 0x9a, 0xc5,
	0x88, 0x1e, 0x34, 0x0a, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x3a, 0x29, 0x9a, 0xc5, 0x88, 0x1e, 0x24, 0x0a, 0x22,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x42, 0x8f, 0x01, 0x0a, 0x26, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0x42, 0x0f, 0x4c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x76,
	0x33, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_endpoint_v3_load_report_proto_rawDescOnce sync.Once
	file_envoy_config_endpoint_v3_load_report_proto_rawDescData = file_envoy_config_endpoint_v3_load_report_proto_rawDesc
)

func file_envoy_config_endpoint_v3_load_report_proto_rawDescGZIP() []byte {
	file_envoy_config_endpoint_v3_load_report_proto_rawDescOnce.Do(func() {
		file_envoy_config_endpoint_v3_load_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_endpoint_v3_load_report_proto_rawDescData)
	})
	return file_envoy_config_endpoint_v3_load_report_proto_rawDescData
}

var file_envoy_config_endpoint_v3_load_report_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_envoy_config_endpoint_v3_load_report_proto_goTypes = []interface{}{
	(*UpstreamLocalityStats)(nil),        // 0: envoy.config.endpoint.v3.UpstreamLocalityStats
	(*UpstreamEndpointStats)(nil),        // 1: envoy.config.endpoint.v3.UpstreamEndpointStats
	(*EndpointLoadMetricStats)(nil),      // 2: envoy.config.endpoint.v3.EndpointLoadMetricStats
	(*ClusterStats)(nil),                 // 3: envoy.config.endpoint.v3.ClusterStats
	(*ClusterStats_DroppedRequests)(nil), // 4: envoy.config.endpoint.v3.ClusterStats.DroppedRequests
	(*v3.Locality)(nil),                  // 5: envoy.config.core.v3.Locality
	(*v3.Address)(nil),                   // 6: envoy.config.core.v3.Address
	(*structpb.Struct)(nil),              // 7: google.protobuf.Struct
	(*durationpb.Duration)(nil),          // 8: google.protobuf.Duration
}
var file_envoy_config_endpoint_v3_load_report_proto_depIdxs = []int32{
	5, // 0: envoy.config.endpoint.v3.UpstreamLocalityStats.locality:type_name -> envoy.config.core.v3.Locality
	2, // 1: envoy.config.endpoint.v3.UpstreamLocalityStats.load_metric_stats:type_name -> envoy.config.endpoint.v3.EndpointLoadMetricStats
	1, // 2: envoy.config.endpoint.v3.UpstreamLocalityStats.upstream_endpoint_stats:type_name -> envoy.config.endpoint.v3.UpstreamEndpointStats
	6, // 3: envoy.config.endpoint.v3.UpstreamEndpointStats.address:type_name -> envoy.config.core.v3.Address
	7, // 4: envoy.config.endpoint.v3.UpstreamEndpointStats.metadata:type_name -> google.protobuf.Struct
	2, // 5: envoy.config.endpoint.v3.UpstreamEndpointStats.load_metric_stats:type_name -> envoy.config.endpoint.v3.EndpointLoadMetricStats
	0, // 6: envoy.config.endpoint.v3.ClusterStats.upstream_locality_stats:type_name -> envoy.config.endpoint.v3.UpstreamLocalityStats
	4, // 7: envoy.config.endpoint.v3.ClusterStats.dropped_requests:type_name -> envoy.config.endpoint.v3.ClusterStats.DroppedRequests
	8, // 8: envoy.config.endpoint.v3.ClusterStats.load_report_interval:type_name -> google.protobuf.Duration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_envoy_config_endpoint_v3_load_report_proto_init() }
func file_envoy_config_endpoint_v3_load_report_proto_init() {
	if File_envoy_config_endpoint_v3_load_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_endpoint_v3_load_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamLocalityStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_endpoint_v3_load_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpstreamEndpointStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_endpoint_v3_load_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointLoadMetricStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_endpoint_v3_load_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_config_endpoint_v3_load_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterStats_DroppedRequests); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_endpoint_v3_load_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_endpoint_v3_load_report_proto_goTypes,
		DependencyIndexes: file_envoy_config_endpoint_v3_load_report_proto_depIdxs,
		MessageInfos:      file_envoy_config_endpoint_v3_load_report_proto_msgTypes,
	}.Build()
	File_envoy_config_endpoint_v3_load_report_proto = out.File
	file_envoy_config_endpoint_v3_load_report_proto_rawDesc = nil
	file_envoy_config_endpoint_v3_load_report_proto_goTypes = nil
	file_envoy_config_endpoint_v3_load_report_proto_depIdxs = nil
}
