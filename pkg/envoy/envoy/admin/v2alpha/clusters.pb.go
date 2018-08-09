// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/clusters.proto

package envoy_admin_v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import _type "github.com/cilium/cilium/pkg/envoy/envoy/type"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Admin endpoint uses this wrapper for `/clusters` to display cluster status information.
// See :ref:`/clusters <operations_admin_interface_clusters>` for more information.
type Clusters struct {
	// Mapping from cluster name to each cluster's status.
	ClusterStatuses      []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_clusters_247b8a9840bb24bd, []int{0}
}
func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (dst *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(dst, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusterStatuses() []*ClusterStatus {
	if m != nil {
		return m.ClusterStatuses
	}
	return nil
}

// Details an individual cluster's current status.
type ClusterStatus struct {
	// Name of the cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Denotes whether this cluster was added via API or configured statically.
	AddedViaApi bool `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	// The success rate threshold used in the last interval. The threshold is used to eject hosts
	// based on their success rate. See
	// :ref:`Cluster outlier detection <arch_overview_outlier_detection>` statistics
	//
	// Note: this field may be omitted in any of the three following cases:
	//
	// 1. There were not enough hosts with enough request volume to proceed with success rate based
	//    outlier ejection.
	// 2. The threshold is computed to be < 0 because a negative value implies that there was no
	//    threshold for that interval.
	// 3. Outlier detection is not enabled for this cluster.
	SuccessRateEjectionThreshold *_type.Percent `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	// Mapping from host address to the host's current status.
	HostStatuses         []*HostStatus `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_clusters_247b8a9840bb24bd, []int{1}
}
func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStatus.Unmarshal(m, b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStatus.Marshal(b, m, deterministic)
}
func (dst *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(dst, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterStatus.Size(m)
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func (m *ClusterStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterStatus) GetAddedViaApi() bool {
	if m != nil {
		return m.AddedViaApi
	}
	return false
}

func (m *ClusterStatus) GetSuccessRateEjectionThreshold() *_type.Percent {
	if m != nil {
		return m.SuccessRateEjectionThreshold
	}
	return nil
}

func (m *ClusterStatus) GetHostStatuses() []*HostStatus {
	if m != nil {
		return m.HostStatuses
	}
	return nil
}

// Current state of a particular host.
type HostStatus struct {
	// Address of this host.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Mapping from the name of the statistic to the current value.
	Stats map[string]*SimpleMetric `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The host's current health status.
	HealthStatus *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	// Request success rate for this host over the last calculated interval.
	//
	// Note: the message will not be present if host did not have enough request volume to calculate
	// success rate or the cluster did not have enough hosts to run through success rate outlier
	// ejection.
	SuccessRate          *_type.Percent `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_clusters_247b8a9840bb24bd, []int{2}
}
func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostStatus.Unmarshal(m, b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
}
func (dst *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(dst, src)
}
func (m *HostStatus) XXX_Size() int {
	return xxx_messageInfo_HostStatus.Size(m)
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *HostStatus) GetStats() map[string]*SimpleMetric {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *HostStatus) GetHealthStatus() *HostHealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return nil
}

func (m *HostStatus) GetSuccessRate() *_type.Percent {
	if m != nil {
		return m.SuccessRate
	}
	return nil
}

// Health status for a host.
type HostHealthStatus struct {
	// The host is currently failing active health checks.
	FailedActiveHealthCheck bool `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	// The host is currently considered an outlier and has been ejected.
	FailedOutlierCheck bool `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	// Health status as reported by EDS. Note: only HEALTHY and UNHEALTHY are currently supported
	// here.
	// TODO(mrice32): pipe through remaining EDS health status possibilities.
	EdsHealthStatus      core.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"eds_health_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HostHealthStatus) Reset()         { *m = HostHealthStatus{} }
func (m *HostHealthStatus) String() string { return proto.CompactTextString(m) }
func (*HostHealthStatus) ProtoMessage()    {}
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_clusters_247b8a9840bb24bd, []int{3}
}
func (m *HostHealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostHealthStatus.Unmarshal(m, b)
}
func (m *HostHealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostHealthStatus.Marshal(b, m, deterministic)
}
func (dst *HostHealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostHealthStatus.Merge(dst, src)
}
func (m *HostHealthStatus) XXX_Size() int {
	return xxx_messageInfo_HostHealthStatus.Size(m)
}
func (m *HostHealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostHealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostHealthStatus proto.InternalMessageInfo

func (m *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if m != nil {
		return m.FailedActiveHealthCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedOutlierCheck() bool {
	if m != nil {
		return m.FailedOutlierCheck
	}
	return false
}

func (m *HostHealthStatus) GetEdsHealthStatus() core.HealthStatus {
	if m != nil {
		return m.EdsHealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func init() {
	proto.RegisterType((*Clusters)(nil), "envoy.admin.v2alpha.Clusters")
	proto.RegisterType((*ClusterStatus)(nil), "envoy.admin.v2alpha.ClusterStatus")
	proto.RegisterType((*HostStatus)(nil), "envoy.admin.v2alpha.HostStatus")
	proto.RegisterMapType((map[string]*SimpleMetric)(nil), "envoy.admin.v2alpha.HostStatus.StatsEntry")
	proto.RegisterType((*HostHealthStatus)(nil), "envoy.admin.v2alpha.HostHealthStatus")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/clusters.proto", fileDescriptor_clusters_247b8a9840bb24bd)
}

var fileDescriptor_clusters_247b8a9840bb24bd = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0x8a, 0xd3, 0x4e,
	0x14, 0xc6, 0x49, 0xdb, 0xfd, 0xff, 0xeb, 0xc9, 0xd6, 0xad, 0xb3, 0x82, 0xa1, 0x08, 0xed, 0x06,
	0x85, 0xe2, 0x45, 0x22, 0x51, 0x54, 0xf4, 0xc6, 0xb2, 0x2e, 0x2c, 0xca, 0xa2, 0x64, 0x45, 0x50,
	0x2f, 0xc2, 0x38, 0x73, 0x24, 0xe3, 0xa6, 0x49, 0x98, 0x99, 0x06, 0xfa, 0x92, 0xbe, 0x85, 0x17,
	0xbe, 0x85, 0x64, 0x66, 0xba, 0xcd, 0xba, 0x5d, 0xbc, 0xea, 0x64, 0xce, 0xef, 0x9b, 0x7e, 0xdf,
	0x39, 0x33, 0x10, 0x62, 0xd9, 0x54, 0xeb, 0x98, 0xf2, 0xa5, 0x28, 0xe3, 0x26, 0xa1, 0x45, 0x9d,
	0xd3, 0x98, 0x15, 0x2b, 0xa5, 0x51, 0xaa, 0xa8, 0x96, 0x95, 0xae, 0xc8, 0xa1, 0x61, 0x22, 0xc3,
	0x44, 0x8e, 0x99, 0x1c, 0xed, 0x12, 0x2e, 0x51, 0x4b, 0xc1, 0x9c, 0x6e, 0x32, 0x75, 0x48, 0x2d,
	0xe2, 0x26, 0x89, 0x59, 0x25, 0x31, 0xa6, 0x9c, 0x4b, 0x54, 0x1b, 0xe0, 0xc1, 0x75, 0x20, 0x47,
	0x5a, 0xe8, 0x3c, 0x63, 0x39, 0xb2, 0x0b, 0x47, 0x05, 0x96, 0xd2, 0xeb, 0x1a, 0xe3, 0x1a, 0x25,
	0xc3, 0x52, 0xdb, 0x4a, 0xf8, 0x19, 0x86, 0xc7, 0xce, 0x2a, 0x39, 0x83, 0xb1, 0xb3, 0x9d, 0x29,
	0x4d, 0xf5, 0x4a, 0xa1, 0x0a, 0xbc, 0x59, 0x7f, 0xee, 0x27, 0x61, 0xb4, 0xc3, 0x7f, 0xe4, 0x84,
	0xe7, 0x86, 0x4d, 0x0f, 0x58, 0xf7, 0x13, 0x55, 0xf8, 0xdb, 0x83, 0xd1, 0x15, 0x84, 0x10, 0x18,
	0x94, 0x74, 0x89, 0x81, 0x37, 0xf3, 0xe6, 0xb7, 0x52, 0xb3, 0x26, 0x21, 0x8c, 0x28, 0xe7, 0xc8,
	0xb3, 0x46, 0xd0, 0x8c, 0xd6, 0x22, 0xe8, 0xcd, 0xbc, 0xf9, 0x30, 0xf5, 0xcd, 0xe6, 0x27, 0x41,
	0x17, 0xb5, 0x20, 0x5f, 0x60, 0xaa, 0x56, 0x8c, 0xa1, 0x52, 0x99, 0xa4, 0x1a, 0x33, 0xfc, 0x81,
	0x4c, 0x8b, 0xaa, 0xcc, 0x74, 0x2e, 0x51, 0xe5, 0x55, 0xc1, 0x83, 0xfe, 0xcc, 0x9b, 0xfb, 0xc9,
	0xa1, 0xf3, 0xd9, 0x06, 0x8d, 0x3e, 0xd8, 0xa0, 0xe9, 0x7d, 0xa7, 0x4d, 0xa9, 0xc6, 0x13, 0xa7,
	0xfc, 0xb8, 0x11, 0x92, 0x37, 0x30, 0xca, 0x2b, 0xa5, 0xb7, 0x89, 0x07, 0x26, 0xf1, 0x74, 0x67,
	0xe2, 0xd3, 0x4a, 0x69, 0x17, 0x77, 0x3f, 0xbf, 0x5c, 0xa3, 0x0a, 0x7f, 0xf5, 0x00, 0xb6, 0x45,
	0xf2, 0x14, 0xfe, 0x77, 0x63, 0x32, 0x59, 0xfd, 0x64, 0xb2, 0x39, 0xae, 0x16, 0x51, 0x93, 0x44,
	0xed, 0x9c, 0xa2, 0x85, 0x25, 0xd2, 0x0d, 0x4a, 0x5e, 0xc3, 0x5e, 0xeb, 0x42, 0x05, 0x3d, 0x63,
	0xe1, 0xd1, 0x3f, 0x2c, 0x44, 0xed, 0x8f, 0x3a, 0x29, 0xb5, 0x5c, 0xa7, 0x56, 0x48, 0xde, 0xc2,
	0xc8, 0x4d, 0xdf, 0xc6, 0x71, 0x6d, 0x79, 0x78, 0xe3, 0x49, 0xa7, 0x86, 0xbe, 0x8c, 0xd4, 0xf9,
	0x22, 0xcf, 0x60, 0xbf, 0xdb, 0xf4, 0x60, 0x70, 0x73, 0x87, 0xfd, 0x4e, 0x87, 0x27, 0x5f, 0x01,
	0xb6, 0xc6, 0xc8, 0x18, 0xfa, 0x17, 0xb8, 0x76, 0x13, 0x6f, 0x97, 0xe4, 0x39, 0xec, 0x35, 0xb4,
	0x58, 0xa1, 0x19, 0xb4, 0x9f, 0x1c, 0xed, 0xf4, 0x76, 0x2e, 0x96, 0x75, 0x81, 0x67, 0xe6, 0x2d,
	0xa4, 0x96, 0x7f, 0xd9, 0x7b, 0xe1, 0x85, 0x3f, 0x3d, 0x18, 0xff, 0xed, 0x9b, 0xbc, 0x82, 0xc9,
	0x77, 0x2a, 0x0a, 0xe4, 0x19, 0x65, 0x5a, 0x34, 0x98, 0x75, 0x5f, 0x80, 0xf9, 0xeb, 0x61, 0x7a,
	0xcf, 0x12, 0x0b, 0x03, 0x58, 0xf5, 0x71, 0x5b, 0x26, 0x8f, 0xe1, 0xae, 0x13, 0x57, 0x2b, 0x5d,
	0x08, 0x94, 0x4e, 0x66, 0xaf, 0x21, 0xb1, 0xb5, 0xf7, 0xb6, 0x64, 0x15, 0xef, 0xe0, 0x0e, 0x72,
	0x95, 0x5d, 0x6f, 0xf4, 0xed, 0xed, 0xad, 0xe9, 0x8c, 0xf9, 0x4a, 0x8b, 0x0f, 0x90, 0xab, 0xee,
	0xc6, 0xb7, 0xff, 0xcc, 0x33, 0x7c, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x25, 0x6b, 0x02, 0x9f,
	0x45, 0x04, 0x00, 0x00,
}
