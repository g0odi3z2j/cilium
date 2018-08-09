// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/config_dump.proto

package envoy_admin_v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v21 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2"
import v2 "github.com/cilium/cilium/pkg/envoy/envoy/config/bootstrap/v2"
import _ "github.com/gogo/protobuf/gogoproto"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The :ref:`/config_dump <operations_admin_interface_config_dump>` admin endpoint uses this wrapper
// message to maintain and serve arbitrary configuration information from any component in Envoy.
type ConfigDump struct {
	// This map is serialized and dumped in its entirety at the
	// :ref:`/config_dump <operations_admin_interface_config_dump>` endpoint.
	//
	// Keys are a short descriptor of the config object they map to. The following keys (and the
	// messages they map to) are currently supported:
	//
	// * *bootstrap*: :ref:`BootstrapConfigDump <envoy_api_msg_admin.v2alpha.BootstrapConfigDump>`
	// * *listeners*: :ref:`ListenersConfigDump <envoy_api_msg_admin.v2alpha.ListenersConfigDump>`
	// * *clusters*: :ref:`ClustersConfigDump <envoy_api_msg_admin.v2alpha.ClustersConfigDump>`
	// * *routes*:  :ref:`RoutesConfigDump <envoy_api_msg_admin.v2alpha.RoutesConfigDump>`
	Configs              map[string]*any.Any `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConfigDump) Reset()         { *m = ConfigDump{} }
func (m *ConfigDump) String() string { return proto.CompactTextString(m) }
func (*ConfigDump) ProtoMessage()    {}
func (*ConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{0}
}
func (m *ConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigDump.Unmarshal(m, b)
}
func (m *ConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigDump.Marshal(b, m, deterministic)
}
func (dst *ConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigDump.Merge(dst, src)
}
func (m *ConfigDump) XXX_Size() int {
	return xxx_messageInfo_ConfigDump.Size(m)
}
func (m *ConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigDump proto.InternalMessageInfo

func (m *ConfigDump) GetConfigs() map[string]*any.Any {
	if m != nil {
		return m.Configs
	}
	return nil
}

// This message describes the bootstrap configuration that Envoy was started with. This includes
// any CLI overrides that were merged. Bootstrap configuration information can be used to recreate
// the static portions of an Envoy configuration by reusing the output as the bootstrap
// configuration for another Envoy.
type BootstrapConfigDump struct {
	Bootstrap *v2.Bootstrap `protobuf:"bytes,1,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	// The timestamp when the BootstrapConfig was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BootstrapConfigDump) Reset()         { *m = BootstrapConfigDump{} }
func (m *BootstrapConfigDump) String() string { return proto.CompactTextString(m) }
func (*BootstrapConfigDump) ProtoMessage()    {}
func (*BootstrapConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{1}
}
func (m *BootstrapConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootstrapConfigDump.Unmarshal(m, b)
}
func (m *BootstrapConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootstrapConfigDump.Marshal(b, m, deterministic)
}
func (dst *BootstrapConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapConfigDump.Merge(dst, src)
}
func (m *BootstrapConfigDump) XXX_Size() int {
	return xxx_messageInfo_BootstrapConfigDump.Size(m)
}
func (m *BootstrapConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapConfigDump proto.InternalMessageInfo

func (m *BootstrapConfigDump) GetBootstrap() *v2.Bootstrap {
	if m != nil {
		return m.Bootstrap
	}
	return nil
}

func (m *BootstrapConfigDump) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Envoy's listener manager fills this message with all currently known listeners. Listener
// configuration information can be used to recreate an Envoy configuration by populating all
// listeners as static listeners or by returning them in a LDS response.
type ListenersConfigDump struct {
	// This is the :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` in the
	// last processed LDS discovery response. If there are only static bootstrap listeners, this field
	// will be "".
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The statically loaded listener configs.
	StaticListeners []*ListenersConfigDump_StaticListener `protobuf:"bytes,2,rep,name=static_listeners,json=staticListeners,proto3" json:"static_listeners,omitempty"`
	// The dynamically loaded active listeners. These are listeners that are available to service
	// data plane traffic.
	DynamicActiveListeners []*ListenersConfigDump_DynamicListener `protobuf:"bytes,3,rep,name=dynamic_active_listeners,json=dynamicActiveListeners,proto3" json:"dynamic_active_listeners,omitempty"`
	// The dynamically loaded warming listeners. These are listeners that are currently undergoing
	// warming in preparation to service data plane traffic. Note that if attempting to recreate an
	// Envoy configuration from a configuration dump, the warming listeners should generally be
	// discarded.
	DynamicWarmingListeners []*ListenersConfigDump_DynamicListener `protobuf:"bytes,4,rep,name=dynamic_warming_listeners,json=dynamicWarmingListeners,proto3" json:"dynamic_warming_listeners,omitempty"`
	// The dynamically loaded draining listeners. These are listeners that are currently undergoing
	// draining in preparation to stop servicing data plane traffic. Note that if attempting to
	// recreate an Envoy configuration from a configuration dump, the draining listeners should
	// generally be discarded.
	DynamicDrainingListeners []*ListenersConfigDump_DynamicListener `protobuf:"bytes,5,rep,name=dynamic_draining_listeners,json=dynamicDrainingListeners,proto3" json:"dynamic_draining_listeners,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                               `json:"-"`
	XXX_unrecognized         []byte                                 `json:"-"`
	XXX_sizecache            int32                                  `json:"-"`
}

func (m *ListenersConfigDump) Reset()         { *m = ListenersConfigDump{} }
func (m *ListenersConfigDump) String() string { return proto.CompactTextString(m) }
func (*ListenersConfigDump) ProtoMessage()    {}
func (*ListenersConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{2}
}
func (m *ListenersConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenersConfigDump.Unmarshal(m, b)
}
func (m *ListenersConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenersConfigDump.Marshal(b, m, deterministic)
}
func (dst *ListenersConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenersConfigDump.Merge(dst, src)
}
func (m *ListenersConfigDump) XXX_Size() int {
	return xxx_messageInfo_ListenersConfigDump.Size(m)
}
func (m *ListenersConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenersConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_ListenersConfigDump proto.InternalMessageInfo

func (m *ListenersConfigDump) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *ListenersConfigDump) GetStaticListeners() []*ListenersConfigDump_StaticListener {
	if m != nil {
		return m.StaticListeners
	}
	return nil
}

func (m *ListenersConfigDump) GetDynamicActiveListeners() []*ListenersConfigDump_DynamicListener {
	if m != nil {
		return m.DynamicActiveListeners
	}
	return nil
}

func (m *ListenersConfigDump) GetDynamicWarmingListeners() []*ListenersConfigDump_DynamicListener {
	if m != nil {
		return m.DynamicWarmingListeners
	}
	return nil
}

func (m *ListenersConfigDump) GetDynamicDrainingListeners() []*ListenersConfigDump_DynamicListener {
	if m != nil {
		return m.DynamicDrainingListeners
	}
	return nil
}

// Describes a statically loaded cluster.
type ListenersConfigDump_StaticListener struct {
	// The listener config.
	Listener *v21.Listener `protobuf:"bytes,1,opt,name=listener,proto3" json:"listener,omitempty"`
	// The timestamp when the Listener was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListenersConfigDump_StaticListener) Reset()         { *m = ListenersConfigDump_StaticListener{} }
func (m *ListenersConfigDump_StaticListener) String() string { return proto.CompactTextString(m) }
func (*ListenersConfigDump_StaticListener) ProtoMessage()    {}
func (*ListenersConfigDump_StaticListener) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{2, 0}
}
func (m *ListenersConfigDump_StaticListener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenersConfigDump_StaticListener.Unmarshal(m, b)
}
func (m *ListenersConfigDump_StaticListener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenersConfigDump_StaticListener.Marshal(b, m, deterministic)
}
func (dst *ListenersConfigDump_StaticListener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenersConfigDump_StaticListener.Merge(dst, src)
}
func (m *ListenersConfigDump_StaticListener) XXX_Size() int {
	return xxx_messageInfo_ListenersConfigDump_StaticListener.Size(m)
}
func (m *ListenersConfigDump_StaticListener) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenersConfigDump_StaticListener.DiscardUnknown(m)
}

var xxx_messageInfo_ListenersConfigDump_StaticListener proto.InternalMessageInfo

func (m *ListenersConfigDump_StaticListener) GetListener() *v21.Listener {
	if m != nil {
		return m.Listener
	}
	return nil
}

func (m *ListenersConfigDump_StaticListener) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Describes a dynamically loaded cluster via the LDS API.
type ListenersConfigDump_DynamicListener struct {
	// This is the per-resource version information. This version is currently taken from the
	// :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` field at the time
	// that the listener was loaded. In the future, discrete per-listener versions may be supported
	// by the API.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The listener config.
	Listener *v21.Listener `protobuf:"bytes,2,opt,name=listener,proto3" json:"listener,omitempty"`
	// The timestamp when the Listener was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListenersConfigDump_DynamicListener) Reset()         { *m = ListenersConfigDump_DynamicListener{} }
func (m *ListenersConfigDump_DynamicListener) String() string { return proto.CompactTextString(m) }
func (*ListenersConfigDump_DynamicListener) ProtoMessage()    {}
func (*ListenersConfigDump_DynamicListener) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{2, 1}
}
func (m *ListenersConfigDump_DynamicListener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenersConfigDump_DynamicListener.Unmarshal(m, b)
}
func (m *ListenersConfigDump_DynamicListener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenersConfigDump_DynamicListener.Marshal(b, m, deterministic)
}
func (dst *ListenersConfigDump_DynamicListener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenersConfigDump_DynamicListener.Merge(dst, src)
}
func (m *ListenersConfigDump_DynamicListener) XXX_Size() int {
	return xxx_messageInfo_ListenersConfigDump_DynamicListener.Size(m)
}
func (m *ListenersConfigDump_DynamicListener) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenersConfigDump_DynamicListener.DiscardUnknown(m)
}

var xxx_messageInfo_ListenersConfigDump_DynamicListener proto.InternalMessageInfo

func (m *ListenersConfigDump_DynamicListener) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *ListenersConfigDump_DynamicListener) GetListener() *v21.Listener {
	if m != nil {
		return m.Listener
	}
	return nil
}

func (m *ListenersConfigDump_DynamicListener) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Envoy's cluster manager fills this message with all currently known clusters. Cluster
// configuration information can be used to recreate an Envoy configuration by populating all
// clusters as static clusters or by returning them in a CDS response.
type ClustersConfigDump struct {
	// This is the :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` in the
	// last processed CDS discovery response. If there are only static bootstrap clusters, this field
	// will be "".
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The statically loaded cluster configs.
	StaticClusters []*ClustersConfigDump_StaticCluster `protobuf:"bytes,2,rep,name=static_clusters,json=staticClusters,proto3" json:"static_clusters,omitempty"`
	// The dynamically loaded active clusters. These are clusters that are available to service
	// data plane traffic.
	DynamicActiveClusters []*ClustersConfigDump_DynamicCluster `protobuf:"bytes,3,rep,name=dynamic_active_clusters,json=dynamicActiveClusters,proto3" json:"dynamic_active_clusters,omitempty"`
	// The dynamically loaded warming clusters. These are clusters that are currently undergoing
	// warming in preparation to service data plane traffic. Note that if attempting to recreate an
	// Envoy configuration from a configuration dump, the warming clusters should generally be
	// discarded.
	DynamicWarmingClusters []*ClustersConfigDump_DynamicCluster `protobuf:"bytes,4,rep,name=dynamic_warming_clusters,json=dynamicWarmingClusters,proto3" json:"dynamic_warming_clusters,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                             `json:"-"`
	XXX_unrecognized       []byte                               `json:"-"`
	XXX_sizecache          int32                                `json:"-"`
}

func (m *ClustersConfigDump) Reset()         { *m = ClustersConfigDump{} }
func (m *ClustersConfigDump) String() string { return proto.CompactTextString(m) }
func (*ClustersConfigDump) ProtoMessage()    {}
func (*ClustersConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{3}
}
func (m *ClustersConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClustersConfigDump.Unmarshal(m, b)
}
func (m *ClustersConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClustersConfigDump.Marshal(b, m, deterministic)
}
func (dst *ClustersConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClustersConfigDump.Merge(dst, src)
}
func (m *ClustersConfigDump) XXX_Size() int {
	return xxx_messageInfo_ClustersConfigDump.Size(m)
}
func (m *ClustersConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_ClustersConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_ClustersConfigDump proto.InternalMessageInfo

func (m *ClustersConfigDump) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *ClustersConfigDump) GetStaticClusters() []*ClustersConfigDump_StaticCluster {
	if m != nil {
		return m.StaticClusters
	}
	return nil
}

func (m *ClustersConfigDump) GetDynamicActiveClusters() []*ClustersConfigDump_DynamicCluster {
	if m != nil {
		return m.DynamicActiveClusters
	}
	return nil
}

func (m *ClustersConfigDump) GetDynamicWarmingClusters() []*ClustersConfigDump_DynamicCluster {
	if m != nil {
		return m.DynamicWarmingClusters
	}
	return nil
}

// Describes a statically loaded cluster.
type ClustersConfigDump_StaticCluster struct {
	// The cluster config.
	Cluster *v21.Cluster `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// The timestamp when the Cluster was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClustersConfigDump_StaticCluster) Reset()         { *m = ClustersConfigDump_StaticCluster{} }
func (m *ClustersConfigDump_StaticCluster) String() string { return proto.CompactTextString(m) }
func (*ClustersConfigDump_StaticCluster) ProtoMessage()    {}
func (*ClustersConfigDump_StaticCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{3, 0}
}
func (m *ClustersConfigDump_StaticCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClustersConfigDump_StaticCluster.Unmarshal(m, b)
}
func (m *ClustersConfigDump_StaticCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClustersConfigDump_StaticCluster.Marshal(b, m, deterministic)
}
func (dst *ClustersConfigDump_StaticCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClustersConfigDump_StaticCluster.Merge(dst, src)
}
func (m *ClustersConfigDump_StaticCluster) XXX_Size() int {
	return xxx_messageInfo_ClustersConfigDump_StaticCluster.Size(m)
}
func (m *ClustersConfigDump_StaticCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_ClustersConfigDump_StaticCluster.DiscardUnknown(m)
}

var xxx_messageInfo_ClustersConfigDump_StaticCluster proto.InternalMessageInfo

func (m *ClustersConfigDump_StaticCluster) GetCluster() *v21.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClustersConfigDump_StaticCluster) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Describes a dynamically loaded cluster via the CDS API.
type ClustersConfigDump_DynamicCluster struct {
	// This is the per-resource version information. This version is currently taken from the
	// :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` field at the time
	// that the cluster was loaded. In the future, discrete per-cluster versions may be supported by
	// the API.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The cluster config.
	Cluster *v21.Cluster `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// The timestamp when the Cluster was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClustersConfigDump_DynamicCluster) Reset()         { *m = ClustersConfigDump_DynamicCluster{} }
func (m *ClustersConfigDump_DynamicCluster) String() string { return proto.CompactTextString(m) }
func (*ClustersConfigDump_DynamicCluster) ProtoMessage()    {}
func (*ClustersConfigDump_DynamicCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{3, 1}
}
func (m *ClustersConfigDump_DynamicCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClustersConfigDump_DynamicCluster.Unmarshal(m, b)
}
func (m *ClustersConfigDump_DynamicCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClustersConfigDump_DynamicCluster.Marshal(b, m, deterministic)
}
func (dst *ClustersConfigDump_DynamicCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClustersConfigDump_DynamicCluster.Merge(dst, src)
}
func (m *ClustersConfigDump_DynamicCluster) XXX_Size() int {
	return xxx_messageInfo_ClustersConfigDump_DynamicCluster.Size(m)
}
func (m *ClustersConfigDump_DynamicCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_ClustersConfigDump_DynamicCluster.DiscardUnknown(m)
}

var xxx_messageInfo_ClustersConfigDump_DynamicCluster proto.InternalMessageInfo

func (m *ClustersConfigDump_DynamicCluster) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *ClustersConfigDump_DynamicCluster) GetCluster() *v21.Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClustersConfigDump_DynamicCluster) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Envoy's RDS implementation fills this message with all currently loaded routes, as described by
// their RouteConfiguration objects. Static routes configured in the bootstrap configuration are
// separated from those configured dynamically via RDS. Route configuration information can be used
// to recreate an Envoy configuration by populating all routes as static routes or by returning them
// in RDS responses.
type RoutesConfigDump struct {
	// The statically loaded route configs.
	StaticRouteConfigs []*RoutesConfigDump_StaticRouteConfig `protobuf:"bytes,2,rep,name=static_route_configs,json=staticRouteConfigs,proto3" json:"static_route_configs,omitempty"`
	// The dynamically loaded route configs.
	DynamicRouteConfigs  []*RoutesConfigDump_DynamicRouteConfig `protobuf:"bytes,3,rep,name=dynamic_route_configs,json=dynamicRouteConfigs,proto3" json:"dynamic_route_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *RoutesConfigDump) Reset()         { *m = RoutesConfigDump{} }
func (m *RoutesConfigDump) String() string { return proto.CompactTextString(m) }
func (*RoutesConfigDump) ProtoMessage()    {}
func (*RoutesConfigDump) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{4}
}
func (m *RoutesConfigDump) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesConfigDump.Unmarshal(m, b)
}
func (m *RoutesConfigDump) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesConfigDump.Marshal(b, m, deterministic)
}
func (dst *RoutesConfigDump) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesConfigDump.Merge(dst, src)
}
func (m *RoutesConfigDump) XXX_Size() int {
	return xxx_messageInfo_RoutesConfigDump.Size(m)
}
func (m *RoutesConfigDump) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesConfigDump.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesConfigDump proto.InternalMessageInfo

func (m *RoutesConfigDump) GetStaticRouteConfigs() []*RoutesConfigDump_StaticRouteConfig {
	if m != nil {
		return m.StaticRouteConfigs
	}
	return nil
}

func (m *RoutesConfigDump) GetDynamicRouteConfigs() []*RoutesConfigDump_DynamicRouteConfig {
	if m != nil {
		return m.DynamicRouteConfigs
	}
	return nil
}

type RoutesConfigDump_StaticRouteConfig struct {
	// The route config.
	RouteConfig *v21.RouteConfiguration `protobuf:"bytes,1,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// The timestamp when the Route was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RoutesConfigDump_StaticRouteConfig) Reset()         { *m = RoutesConfigDump_StaticRouteConfig{} }
func (m *RoutesConfigDump_StaticRouteConfig) String() string { return proto.CompactTextString(m) }
func (*RoutesConfigDump_StaticRouteConfig) ProtoMessage()    {}
func (*RoutesConfigDump_StaticRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{4, 0}
}
func (m *RoutesConfigDump_StaticRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.Unmarshal(m, b)
}
func (m *RoutesConfigDump_StaticRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.Marshal(b, m, deterministic)
}
func (dst *RoutesConfigDump_StaticRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.Merge(dst, src)
}
func (m *RoutesConfigDump_StaticRouteConfig) XXX_Size() int {
	return xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.Size(m)
}
func (m *RoutesConfigDump_StaticRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesConfigDump_StaticRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesConfigDump_StaticRouteConfig proto.InternalMessageInfo

func (m *RoutesConfigDump_StaticRouteConfig) GetRouteConfig() *v21.RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *RoutesConfigDump_StaticRouteConfig) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type RoutesConfigDump_DynamicRouteConfig struct {
	// This is the per-resource version information. This version is currently taken from the
	// :ref:`version_info <envoy_api_field_DiscoveryResponse.version_info>` field at the time that
	// the route configuration was loaded.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// The route config.
	RouteConfig *v21.RouteConfiguration `protobuf:"bytes,2,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// The timestamp when the Route was last updated.
	LastUpdated          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RoutesConfigDump_DynamicRouteConfig) Reset()         { *m = RoutesConfigDump_DynamicRouteConfig{} }
func (m *RoutesConfigDump_DynamicRouteConfig) String() string { return proto.CompactTextString(m) }
func (*RoutesConfigDump_DynamicRouteConfig) ProtoMessage()    {}
func (*RoutesConfigDump_DynamicRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_dump_8839bd1bf1fe5fcc, []int{4, 1}
}
func (m *RoutesConfigDump_DynamicRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.Unmarshal(m, b)
}
func (m *RoutesConfigDump_DynamicRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.Marshal(b, m, deterministic)
}
func (dst *RoutesConfigDump_DynamicRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.Merge(dst, src)
}
func (m *RoutesConfigDump_DynamicRouteConfig) XXX_Size() int {
	return xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.Size(m)
}
func (m *RoutesConfigDump_DynamicRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesConfigDump_DynamicRouteConfig proto.InternalMessageInfo

func (m *RoutesConfigDump_DynamicRouteConfig) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *RoutesConfigDump_DynamicRouteConfig) GetRouteConfig() *v21.RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *RoutesConfigDump_DynamicRouteConfig) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigDump)(nil), "envoy.admin.v2alpha.ConfigDump")
	proto.RegisterMapType((map[string]*any.Any)(nil), "envoy.admin.v2alpha.ConfigDump.ConfigsEntry")
	proto.RegisterType((*BootstrapConfigDump)(nil), "envoy.admin.v2alpha.BootstrapConfigDump")
	proto.RegisterType((*ListenersConfigDump)(nil), "envoy.admin.v2alpha.ListenersConfigDump")
	proto.RegisterType((*ListenersConfigDump_StaticListener)(nil), "envoy.admin.v2alpha.ListenersConfigDump.StaticListener")
	proto.RegisterType((*ListenersConfigDump_DynamicListener)(nil), "envoy.admin.v2alpha.ListenersConfigDump.DynamicListener")
	proto.RegisterType((*ClustersConfigDump)(nil), "envoy.admin.v2alpha.ClustersConfigDump")
	proto.RegisterType((*ClustersConfigDump_StaticCluster)(nil), "envoy.admin.v2alpha.ClustersConfigDump.StaticCluster")
	proto.RegisterType((*ClustersConfigDump_DynamicCluster)(nil), "envoy.admin.v2alpha.ClustersConfigDump.DynamicCluster")
	proto.RegisterType((*RoutesConfigDump)(nil), "envoy.admin.v2alpha.RoutesConfigDump")
	proto.RegisterType((*RoutesConfigDump_StaticRouteConfig)(nil), "envoy.admin.v2alpha.RoutesConfigDump.StaticRouteConfig")
	proto.RegisterType((*RoutesConfigDump_DynamicRouteConfig)(nil), "envoy.admin.v2alpha.RoutesConfigDump.DynamicRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/admin/v2alpha/config_dump.proto", fileDescriptor_config_dump_8839bd1bf1fe5fcc)
}

var fileDescriptor_config_dump_8839bd1bf1fe5fcc = []byte{
	// 762 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0x1f, 0x27, 0x69, 0xfb, 0x74, 0x92, 0xa7, 0xed, 0xb3, 0x69, 0xd3, 0xd4, 0x97, 0x86,
	0x0a, 0xa4, 0x82, 0x90, 0x2d, 0x05, 0x01, 0x15, 0x12, 0x87, 0xbe, 0x20, 0x81, 0xd4, 0x03, 0x32,
	0x20, 0x8e, 0xd1, 0x36, 0x76, 0x52, 0x8b, 0x64, 0xd7, 0xf2, 0xae, 0x0d, 0x41, 0x48, 0x1c, 0xf8,
	0x10, 0xdc, 0xb8, 0x70, 0xe0, 0xc2, 0x89, 0x2b, 0x77, 0xc4, 0x8d, 0x6f, 0xc0, 0x67, 0x41, 0xb1,
	0x67, 0x1d, 0xbf, 0x84, 0x2a, 0x69, 0x73, 0xdb, 0xcc, 0xce, 0xce, 0xef, 0xbf, 0xb3, 0xff, 0x71,
	0xe0, 0x86, 0xc3, 0x42, 0x3e, 0x32, 0xa9, 0x3d, 0x74, 0x99, 0x19, 0xb6, 0xe9, 0xc0, 0x3b, 0xa7,
	0x66, 0x97, 0xb3, 0x9e, 0xdb, 0xef, 0xd8, 0xc1, 0xd0, 0x33, 0x3c, 0x9f, 0x4b, 0x4e, 0xea, 0x51,
	0x9a, 0x11, 0xa5, 0x19, 0x98, 0xa6, 0x37, 0xf0, 0xac, 0xe7, 0x9a, 0x61, 0xdb, 0xec, 0xda, 0x22,
	0x4e, 0xce, 0xc5, 0x07, 0x7f, 0x89, 0xfb, 0x49, 0xfc, 0x66, 0x1c, 0x8f, 0xa9, 0xe6, 0x19, 0xe7,
	0x52, 0x48, 0x9f, 0x7a, 0xe3, 0xa4, 0xe4, 0x07, 0xa6, 0xee, 0xf4, 0x39, 0xef, 0x0f, 0x1c, 0x33,
	0xfa, 0x75, 0x16, 0xf4, 0x4c, 0xca, 0x46, 0xb8, 0xb5, 0x9b, 0xdf, 0x92, 0xee, 0xd0, 0x11, 0x92,
	0xaa, 0x3b, 0xe8, 0x9b, 0x7d, 0xde, 0xe7, 0xd1, 0xd2, 0x1c, 0xaf, 0xe2, 0xe8, 0xde, 0x57, 0x0d,
	0xe0, 0x38, 0x22, 0x9f, 0x04, 0x43, 0x8f, 0x9c, 0xc2, 0x4a, 0xac, 0x43, 0x34, 0xb5, 0x56, 0x79,
	0xbf, 0xda, 0xbe, 0x6d, 0x4c, 0xb9, 0xba, 0x31, 0x39, 0x81, 0x4b, 0xf1, 0x88, 0x49, 0x7f, 0x74,
	0x54, 0xf9, 0xf9, 0x7b, 0xf7, 0x1f, 0x4b, 0x95, 0xd0, 0x9f, 0x42, 0x2d, 0xbd, 0x4d, 0x36, 0xa0,
	0xfc, 0xca, 0x19, 0x35, 0xb5, 0x96, 0xb6, 0xbf, 0x6a, 0x8d, 0x97, 0xe4, 0x16, 0x2c, 0x85, 0x74,
	0x10, 0x38, 0xcd, 0x52, 0x4b, 0xdb, 0xaf, 0xb6, 0x37, 0x8d, 0xf8, 0x16, 0x86, 0xba, 0x85, 0x71,
	0xc8, 0x46, 0x56, 0x9c, 0xf2, 0xa0, 0x74, 0xa0, 0xed, 0x7d, 0xd2, 0xa0, 0x7e, 0xa4, 0x9a, 0x92,
	0xd2, 0xfd, 0x18, 0x56, 0x93, 0x5e, 0x45, 0xf5, 0xab, 0xed, 0xeb, 0xa8, 0x3c, 0x16, 0x63, 0x4c,
	0x5a, 0x19, 0xb6, 0x8d, 0xa4, 0x04, 0x2a, 0x9e, 0x1c, 0x26, 0x0f, 0xa1, 0x36, 0xa0, 0x42, 0x76,
	0x02, 0xcf, 0xa6, 0xd2, 0xb1, 0x51, 0x98, 0x5e, 0x10, 0xf6, 0x5c, 0xb5, 0xd7, 0xaa, 0x8e, 0xf3,
	0x5f, 0xc4, 0xe9, 0x7b, 0x3f, 0x96, 0xa1, 0x7e, 0xea, 0x0a, 0xe9, 0x30, 0xc7, 0x17, 0x29, 0x81,
	0xd7, 0xa0, 0x16, 0x3a, 0xbe, 0x70, 0x39, 0xeb, 0xb8, 0xac, 0xc7, 0xb1, 0x07, 0x55, 0x8c, 0x3d,
	0x61, 0x3d, 0x4e, 0xce, 0x61, 0x43, 0x48, 0x2a, 0xdd, 0x6e, 0x67, 0xa0, 0x0a, 0x34, 0x4b, 0xd1,
	0x23, 0xdc, 0x9f, 0xfa, 0x08, 0x53, 0x30, 0xc6, 0xb3, 0xa8, 0x80, 0xda, 0xc1, 0xdb, 0xad, 0x8b,
	0x4c, 0x54, 0x90, 0x37, 0xd0, 0xb4, 0x47, 0x8c, 0x0e, 0xdd, 0x6e, 0x87, 0x76, 0xa5, 0x1b, 0x3a,
	0x29, 0x62, 0x39, 0x22, 0x1e, 0xcc, 0x4c, 0x3c, 0x89, 0x0b, 0xe5, 0x90, 0x0d, 0xac, 0x7f, 0x18,
	0x95, 0x9f, 0x90, 0xdf, 0xc2, 0x8e, 0x22, 0xbf, 0xa6, 0xfe, 0xd0, 0x65, 0xfd, 0x14, 0xba, 0xb2,
	0x10, 0xf4, 0x36, 0x02, 0x5e, 0xc6, 0xf5, 0x27, 0xec, 0x77, 0xa0, 0x2b, 0xb6, 0xed, 0x53, 0x97,
	0x65, 0xe1, 0x4b, 0x0b, 0x81, 0xab, 0xbe, 0x9e, 0x20, 0x20, 0x39, 0xa9, 0x7f, 0xd0, 0x60, 0x2d,
	0xfb, 0x3a, 0xa4, 0x0d, 0xff, 0x2a, 0x3e, 0x7a, 0xb6, 0xa1, 0xf0, 0x9e, 0x3b, 0xb6, 0xa9, 0xca,
	0xb4, 0x92, 0xbc, 0x2b, 0xda, 0x53, 0xff, 0xa2, 0xc1, 0x7a, 0x4e, 0xf9, 0x2c, 0xd6, 0x4c, 0x2b,
	0x2d, 0x5d, 0x52, 0x69, 0x79, 0xbe, 0x41, 0xfa, 0xb6, 0x04, 0xe4, 0x78, 0x10, 0x08, 0x39, 0xf7,
	0x1c, 0xd9, 0x80, 0x86, 0xef, 0x74, 0xf1, 0x3c, 0x8e, 0xd1, 0xdd, 0xe9, 0xdf, 0xb2, 0x02, 0x04,
	0xa7, 0x08, 0x37, 0xf0, 0x65, 0xd7, 0x44, 0x3a, 0x28, 0x88, 0x84, 0xed, 0xdc, 0x0c, 0x25, 0xb4,
	0x78, 0x84, 0xee, 0xcd, 0x4a, 0xc3, 0xf7, 0xc8, 0xe2, 0xb6, 0x32, 0x03, 0x94, 0x50, 0xc3, 0xc9,
	0xe4, 0xaa, 0xf9, 0x49, 0xb0, 0x95, 0x05, 0x60, 0x1b, 0xd9, 0xe1, 0x51, 0xc7, 0xf4, 0xf7, 0xf0,
	0x5f, 0xa6, 0x29, 0xc4, 0x84, 0x15, 0x04, 0xa3, 0x75, 0xb7, 0xb2, 0x86, 0xc0, 0x3c, 0x4b, 0x65,
	0x5d, 0xd5, 0xb8, 0x9f, 0x35, 0x58, 0xcb, 0x2a, 0x9e, 0xc5, 0x0a, 0x29, 0x95, 0xa5, 0x4b, 0xa9,
	0x9c, 0xd3, 0xb4, 0xbf, 0x2a, 0xb0, 0x61, 0xf1, 0x40, 0x3a, 0x69, 0xcb, 0x72, 0xd8, 0x44, 0x3f,
	0xfa, 0xe3, 0xad, 0x8e, 0xfa, 0x83, 0xbd, 0xe8, 0xdb, 0x9e, 0x2f, 0x82, 0x96, 0x8c, 0xc2, 0x71,
	0x14, 0x1f, 0x8c, 0x88, 0xfc, 0x86, 0x20, 0x3e, 0x28, 0xf7, 0xe4, 0x88, 0x17, 0x7d, 0xdb, 0x0b,
	0x44, 0xec, 0x76, 0x11, 0x59, 0xb7, 0x0b, 0x3b, 0x42, 0xff, 0xa8, 0xc1, 0xff, 0x05, 0x8d, 0xe4,
	0x18, 0x6a, 0x69, 0x05, 0x68, 0x95, 0x56, 0xf6, 0x11, 0x52, 0x07, 0x02, 0x9f, 0x4a, 0x97, 0x33,
	0xab, 0xea, 0xa7, 0x8a, 0x5c, 0xd1, 0x39, 0xdf, 0x35, 0x20, 0xc5, 0xbb, 0xcc, 0xe2, 0x9e, 0xbc,
	0xfa, 0xd2, 0x22, 0xd4, 0xcf, 0xe7, 0xa8, 0xb3, 0xe5, 0x28, 0xe1, 0xce, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xbb, 0x43, 0xe8, 0x93, 0xa9, 0x0a, 0x00, 0x00,
}
