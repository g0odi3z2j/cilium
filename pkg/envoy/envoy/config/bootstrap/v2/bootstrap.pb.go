// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/bootstrap/v2/bootstrap.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v23 "github.com/cilium/cilium/pkg/envoy/envoy/api/v2"
import auth "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/auth"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import v2 "github.com/cilium/cilium/pkg/envoy/envoy/config/metrics/v2"
import v22 "github.com/cilium/cilium/pkg/envoy/envoy/config/ratelimit/v2"
import v21 "github.com/cilium/cilium/pkg/envoy/envoy/config/trace/v2"
import _ "github.com/gogo/protobuf/gogoproto"
import duration "github.com/golang/protobuf/ptypes/duration"
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

// Bootstrap :ref:`configuration overview <config_overview_v2_bootstrap>`.
type Bootstrap struct {
	// Node identity to present to the management server and for instance
	// identification purposes (e.g. in generated headers).
	Node *core.Node `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	// Statically specified resources.
	StaticResources *Bootstrap_StaticResources `protobuf:"bytes,2,opt,name=static_resources,json=staticResources" json:"static_resources,omitempty"`
	// xDS configuration sources.
	DynamicResources *Bootstrap_DynamicResources `protobuf:"bytes,3,opt,name=dynamic_resources,json=dynamicResources" json:"dynamic_resources,omitempty"`
	// Configuration for the cluster manager which owns all upstream clusters
	// within the server.
	ClusterManager *ClusterManager `protobuf:"bytes,4,opt,name=cluster_manager,json=clusterManager" json:"cluster_manager,omitempty"`
	// Optional file system path to search for startup flag files.
	FlagsPath string `protobuf:"bytes,5,opt,name=flags_path,json=flagsPath" json:"flags_path,omitempty"`
	// Optional set of stats sinks.
	StatsSinks []*v2.StatsSink `protobuf:"bytes,6,rep,name=stats_sinks,json=statsSinks" json:"stats_sinks,omitempty"`
	// Configuration for internal processing of stats.
	StatsConfig *v2.StatsConfig `protobuf:"bytes,13,opt,name=stats_config,json=statsConfig" json:"stats_config,omitempty"`
	// Optional duration between flushes to configured stats sinks. For
	// performance reasons Envoy latches counters and only flushes counters and
	// gauges at a periodic interval. If not specified the default is 5000ms (5
	// seconds).
	StatsFlushInterval *duration.Duration `protobuf:"bytes,7,opt,name=stats_flush_interval,json=statsFlushInterval" json:"stats_flush_interval,omitempty"`
	// Optional watchdog configuration.
	Watchdog *Watchdog `protobuf:"bytes,8,opt,name=watchdog" json:"watchdog,omitempty"`
	// Configuration for an external tracing provider. If not specified, no
	// tracing will be performed.
	Tracing *v21.Tracing `protobuf:"bytes,9,opt,name=tracing" json:"tracing,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService *v22.RateLimitServiceConfig `protobuf:"bytes,10,opt,name=rate_limit_service,json=rateLimitService" json:"rate_limit_service,omitempty"`
	// Configuration for the runtime configuration provider. If not specified, a
	// “null” provider will be used which will result in all defaults being used.
	Runtime *Runtime `protobuf:"bytes,11,opt,name=runtime" json:"runtime,omitempty"`
	// Configuration for the local administration HTTP server.
	Admin                *Admin   `protobuf:"bytes,12,opt,name=admin" json:"admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bootstrap) Reset()         { *m = Bootstrap{} }
func (m *Bootstrap) String() string { return proto.CompactTextString(m) }
func (*Bootstrap) ProtoMessage()    {}
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{0}
}
func (m *Bootstrap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap.Unmarshal(m, b)
}
func (m *Bootstrap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap.Marshal(b, m, deterministic)
}
func (dst *Bootstrap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap.Merge(dst, src)
}
func (m *Bootstrap) XXX_Size() int {
	return xxx_messageInfo_Bootstrap.Size(m)
}
func (m *Bootstrap) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap proto.InternalMessageInfo

func (m *Bootstrap) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Bootstrap) GetStaticResources() *Bootstrap_StaticResources {
	if m != nil {
		return m.StaticResources
	}
	return nil
}

func (m *Bootstrap) GetDynamicResources() *Bootstrap_DynamicResources {
	if m != nil {
		return m.DynamicResources
	}
	return nil
}

func (m *Bootstrap) GetClusterManager() *ClusterManager {
	if m != nil {
		return m.ClusterManager
	}
	return nil
}

func (m *Bootstrap) GetFlagsPath() string {
	if m != nil {
		return m.FlagsPath
	}
	return ""
}

func (m *Bootstrap) GetStatsSinks() []*v2.StatsSink {
	if m != nil {
		return m.StatsSinks
	}
	return nil
}

func (m *Bootstrap) GetStatsConfig() *v2.StatsConfig {
	if m != nil {
		return m.StatsConfig
	}
	return nil
}

func (m *Bootstrap) GetStatsFlushInterval() *duration.Duration {
	if m != nil {
		return m.StatsFlushInterval
	}
	return nil
}

func (m *Bootstrap) GetWatchdog() *Watchdog {
	if m != nil {
		return m.Watchdog
	}
	return nil
}

func (m *Bootstrap) GetTracing() *v21.Tracing {
	if m != nil {
		return m.Tracing
	}
	return nil
}

func (m *Bootstrap) GetRateLimitService() *v22.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func (m *Bootstrap) GetRuntime() *Runtime {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *Bootstrap) GetAdmin() *Admin {
	if m != nil {
		return m.Admin
	}
	return nil
}

type Bootstrap_StaticResources struct {
	// Static :ref:`Listeners <envoy_api_msg_Listener>`. These listeners are
	// available regardless of LDS configuration.
	Listeners []*v23.Listener `protobuf:"bytes,1,rep,name=listeners" json:"listeners,omitempty"`
	// If a network based configuration source is specified for :ref:`cds_config
	// <envoy_api_field_config.bootstrap.v2.Bootstrap.DynamicResources.cds_config>`, it's necessary
	// to have some initial cluster definitions available to allow Envoy to know
	// how to speak to the management server. These cluster definitions may not
	// use :ref:`EDS <arch_overview_dynamic_config_sds>` (i.e. they should be static
	// IP or DNS-based).
	Clusters []*v23.Cluster `protobuf:"bytes,2,rep,name=clusters" json:"clusters,omitempty"`
	// [#not-implemented-hide:]
	Secrets              []*auth.Secret `protobuf:"bytes,3,rep,name=secrets" json:"secrets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Bootstrap_StaticResources) Reset()         { *m = Bootstrap_StaticResources{} }
func (m *Bootstrap_StaticResources) String() string { return proto.CompactTextString(m) }
func (*Bootstrap_StaticResources) ProtoMessage()    {}
func (*Bootstrap_StaticResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{0, 0}
}
func (m *Bootstrap_StaticResources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap_StaticResources.Unmarshal(m, b)
}
func (m *Bootstrap_StaticResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap_StaticResources.Marshal(b, m, deterministic)
}
func (dst *Bootstrap_StaticResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap_StaticResources.Merge(dst, src)
}
func (m *Bootstrap_StaticResources) XXX_Size() int {
	return xxx_messageInfo_Bootstrap_StaticResources.Size(m)
}
func (m *Bootstrap_StaticResources) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap_StaticResources.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap_StaticResources proto.InternalMessageInfo

func (m *Bootstrap_StaticResources) GetListeners() []*v23.Listener {
	if m != nil {
		return m.Listeners
	}
	return nil
}

func (m *Bootstrap_StaticResources) GetClusters() []*v23.Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *Bootstrap_StaticResources) GetSecrets() []*auth.Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

type Bootstrap_DynamicResources struct {
	// All :ref:`Listeners <envoy_api_msg_Listener>` are provided by a single
	// :ref:`LDS <arch_overview_dynamic_config_lds>` configuration source.
	LdsConfig *core.ConfigSource `protobuf:"bytes,1,opt,name=lds_config,json=ldsConfig" json:"lds_config,omitempty"`
	// All post-bootstrap :ref:`Cluster <envoy_api_msg_Cluster>` definitions are
	// provided by a single :ref:`CDS <arch_overview_dynamic_config_cds>`
	// configuration source.
	CdsConfig *core.ConfigSource `protobuf:"bytes,2,opt,name=cds_config,json=cdsConfig" json:"cds_config,omitempty"`
	// A single :ref:`ADS <config_overview_v2_ads>` source may be optionally
	// specified. This must have :ref:`api_type
	// <envoy_api_field_core.ApiConfigSource.api_type>` :ref:`GRPC
	// <envoy_api_enum_value_core.ApiConfigSource.ApiType.GRPC>`. Only
	// :ref:`ConfigSources <envoy_api_msg_core.ConfigSource>` that have
	// the :ref:`ads <envoy_api_field_core.ConfigSource.ads>` field set will be
	// streamed on the ADS channel.
	AdsConfig *core.ApiConfigSource `protobuf:"bytes,3,opt,name=ads_config,json=adsConfig" json:"ads_config,omitempty"`
	// [#not-implemented-hide:] Hide from docs.
	DeprecatedV1         *Bootstrap_DynamicResources_DeprecatedV1 `protobuf:"bytes,4,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *Bootstrap_DynamicResources) Reset()         { *m = Bootstrap_DynamicResources{} }
func (m *Bootstrap_DynamicResources) String() string { return proto.CompactTextString(m) }
func (*Bootstrap_DynamicResources) ProtoMessage()    {}
func (*Bootstrap_DynamicResources) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{0, 1}
}
func (m *Bootstrap_DynamicResources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap_DynamicResources.Unmarshal(m, b)
}
func (m *Bootstrap_DynamicResources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap_DynamicResources.Marshal(b, m, deterministic)
}
func (dst *Bootstrap_DynamicResources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap_DynamicResources.Merge(dst, src)
}
func (m *Bootstrap_DynamicResources) XXX_Size() int {
	return xxx_messageInfo_Bootstrap_DynamicResources.Size(m)
}
func (m *Bootstrap_DynamicResources) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap_DynamicResources.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap_DynamicResources proto.InternalMessageInfo

func (m *Bootstrap_DynamicResources) GetLdsConfig() *core.ConfigSource {
	if m != nil {
		return m.LdsConfig
	}
	return nil
}

func (m *Bootstrap_DynamicResources) GetCdsConfig() *core.ConfigSource {
	if m != nil {
		return m.CdsConfig
	}
	return nil
}

func (m *Bootstrap_DynamicResources) GetAdsConfig() *core.ApiConfigSource {
	if m != nil {
		return m.AdsConfig
	}
	return nil
}

// Deprecated: Do not use.
func (m *Bootstrap_DynamicResources) GetDeprecatedV1() *Bootstrap_DynamicResources_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

// [#not-implemented-hide:] Hide from docs.
type Bootstrap_DynamicResources_DeprecatedV1 struct {
	// This is the global :ref:`SDS <arch_overview_dynamic_config_sds>` config
	// when using v1 REST for :ref:`CDS
	// <arch_overview_dynamic_config_cds>`/:ref:`EDS
	// <arch_overview_dynamic_config_sds>`.
	SdsConfig            *core.ConfigSource `protobuf:"bytes,1,opt,name=sds_config,json=sdsConfig" json:"sds_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Bootstrap_DynamicResources_DeprecatedV1) Reset() {
	*m = Bootstrap_DynamicResources_DeprecatedV1{}
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Bootstrap_DynamicResources_DeprecatedV1) ProtoMessage()    {}
func (*Bootstrap_DynamicResources_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{0, 1, 0}
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.Unmarshal(m, b)
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.Marshal(b, m, deterministic)
}
func (dst *Bootstrap_DynamicResources_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.Merge(dst, src)
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.Size(m)
}
func (m *Bootstrap_DynamicResources_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Bootstrap_DynamicResources_DeprecatedV1 proto.InternalMessageInfo

func (m *Bootstrap_DynamicResources_DeprecatedV1) GetSdsConfig() *core.ConfigSource {
	if m != nil {
		return m.SdsConfig
	}
	return nil
}

// Administration interface :ref:`operations documentation
// <operations_admin_interface>`.
type Admin struct {
	// The path to write the access log for the administration server. If no
	// access log is desired specify ‘/dev/null’.
	AccessLogPath string `protobuf:"bytes,1,opt,name=access_log_path,json=accessLogPath" json:"access_log_path,omitempty"`
	// The cpu profiler output path for the administration server. If no profile
	// path is specified, the default is ‘/var/log/envoy/envoy.prof’.
	ProfilePath string `protobuf:"bytes,2,opt,name=profile_path,json=profilePath" json:"profile_path,omitempty"`
	// The TCP address that the administration server will listen on.
	Address              *core.Address `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Admin) Reset()         { *m = Admin{} }
func (m *Admin) String() string { return proto.CompactTextString(m) }
func (*Admin) ProtoMessage()    {}
func (*Admin) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{1}
}
func (m *Admin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Admin.Unmarshal(m, b)
}
func (m *Admin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Admin.Marshal(b, m, deterministic)
}
func (dst *Admin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Admin.Merge(dst, src)
}
func (m *Admin) XXX_Size() int {
	return xxx_messageInfo_Admin.Size(m)
}
func (m *Admin) XXX_DiscardUnknown() {
	xxx_messageInfo_Admin.DiscardUnknown(m)
}

var xxx_messageInfo_Admin proto.InternalMessageInfo

func (m *Admin) GetAccessLogPath() string {
	if m != nil {
		return m.AccessLogPath
	}
	return ""
}

func (m *Admin) GetProfilePath() string {
	if m != nil {
		return m.ProfilePath
	}
	return ""
}

func (m *Admin) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

// Cluster manager :ref:`architecture overview <arch_overview_cluster_manager>`.
type ClusterManager struct {
	// Name of the local cluster (i.e., the cluster that owns the Envoy running
	// this configuration). In order to enable :ref:`zone aware routing
	// <arch_overview_load_balancing_zone_aware_routing>` this option must be set.
	// If *local_cluster_name* is defined then :ref:`clusters
	// <config_cluster_manager_clusters>` must be defined in the :ref:`Bootstrap
	// static cluster resources
	// <envoy_api_field_config.bootstrap.v2.Bootstrap.StaticResources.clusters>`. This is unrelated to
	// the :option:`--service-cluster` option which does not `affect zone aware
	// routing <https://github.com/envoyproxy/envoy/issues/774>`_.
	LocalClusterName string `protobuf:"bytes,1,opt,name=local_cluster_name,json=localClusterName" json:"local_cluster_name,omitempty"`
	// Optional global configuration for outlier detection.
	OutlierDetection *ClusterManager_OutlierDetection `protobuf:"bytes,2,opt,name=outlier_detection,json=outlierDetection" json:"outlier_detection,omitempty"`
	// Optional configuration used to bind newly established upstream connections.
	// This may be overridden on a per-cluster basis by upstream_bind_config in the cds_config.
	UpstreamBindConfig *core.BindConfig `protobuf:"bytes,3,opt,name=upstream_bind_config,json=upstreamBindConfig" json:"upstream_bind_config,omitempty"`
	// A management server endpoint to stream load stats to via
	// *StreamLoadStats*. This must have :ref:`api_type
	// <envoy_api_field_core.ApiConfigSource.api_type>` :ref:`GRPC
	// <envoy_api_enum_value_core.ApiConfigSource.ApiType.GRPC>`.
	LoadStatsConfig      *core.ApiConfigSource `protobuf:"bytes,4,opt,name=load_stats_config,json=loadStatsConfig" json:"load_stats_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClusterManager) Reset()         { *m = ClusterManager{} }
func (m *ClusterManager) String() string { return proto.CompactTextString(m) }
func (*ClusterManager) ProtoMessage()    {}
func (*ClusterManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{2}
}
func (m *ClusterManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterManager.Unmarshal(m, b)
}
func (m *ClusterManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterManager.Marshal(b, m, deterministic)
}
func (dst *ClusterManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterManager.Merge(dst, src)
}
func (m *ClusterManager) XXX_Size() int {
	return xxx_messageInfo_ClusterManager.Size(m)
}
func (m *ClusterManager) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterManager.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterManager proto.InternalMessageInfo

func (m *ClusterManager) GetLocalClusterName() string {
	if m != nil {
		return m.LocalClusterName
	}
	return ""
}

func (m *ClusterManager) GetOutlierDetection() *ClusterManager_OutlierDetection {
	if m != nil {
		return m.OutlierDetection
	}
	return nil
}

func (m *ClusterManager) GetUpstreamBindConfig() *core.BindConfig {
	if m != nil {
		return m.UpstreamBindConfig
	}
	return nil
}

func (m *ClusterManager) GetLoadStatsConfig() *core.ApiConfigSource {
	if m != nil {
		return m.LoadStatsConfig
	}
	return nil
}

type ClusterManager_OutlierDetection struct {
	// Specifies the path to the outlier event log.
	EventLogPath         string   `protobuf:"bytes,1,opt,name=event_log_path,json=eventLogPath" json:"event_log_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterManager_OutlierDetection) Reset()         { *m = ClusterManager_OutlierDetection{} }
func (m *ClusterManager_OutlierDetection) String() string { return proto.CompactTextString(m) }
func (*ClusterManager_OutlierDetection) ProtoMessage()    {}
func (*ClusterManager_OutlierDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{2, 0}
}
func (m *ClusterManager_OutlierDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterManager_OutlierDetection.Unmarshal(m, b)
}
func (m *ClusterManager_OutlierDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterManager_OutlierDetection.Marshal(b, m, deterministic)
}
func (dst *ClusterManager_OutlierDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterManager_OutlierDetection.Merge(dst, src)
}
func (m *ClusterManager_OutlierDetection) XXX_Size() int {
	return xxx_messageInfo_ClusterManager_OutlierDetection.Size(m)
}
func (m *ClusterManager_OutlierDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterManager_OutlierDetection.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterManager_OutlierDetection proto.InternalMessageInfo

func (m *ClusterManager_OutlierDetection) GetEventLogPath() string {
	if m != nil {
		return m.EventLogPath
	}
	return ""
}

// Envoy process watchdog configuration. When configured, this monitors for
// nonresponsive threads and kills the process after the configured thresholds.
type Watchdog struct {
	// The duration after which Envoy counts a nonresponsive thread in the
	// *server.watchdog_miss* statistic. If not specified the default is 200ms.
	MissTimeout *duration.Duration `protobuf:"bytes,1,opt,name=miss_timeout,json=missTimeout" json:"miss_timeout,omitempty"`
	// The duration after which Envoy counts a nonresponsive thread in the
	// *server.watchdog_mega_miss* statistic. If not specified the default is
	// 1000ms.
	MegamissTimeout *duration.Duration `protobuf:"bytes,2,opt,name=megamiss_timeout,json=megamissTimeout" json:"megamiss_timeout,omitempty"`
	// If a watched thread has been nonresponsive for this duration, assume a
	// programming error and kill the entire Envoy process. Set to 0 to disable
	// kill behavior. If not specified the default is 0 (disabled).
	KillTimeout *duration.Duration `protobuf:"bytes,3,opt,name=kill_timeout,json=killTimeout" json:"kill_timeout,omitempty"`
	// If at least two watched threads have been nonresponsive for at least this
	// duration assume a true deadlock and kill the entire Envoy process. Set to 0
	// to disable this behavior. If not specified the default is 0 (disabled).
	MultikillTimeout     *duration.Duration `protobuf:"bytes,4,opt,name=multikill_timeout,json=multikillTimeout" json:"multikill_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Watchdog) Reset()         { *m = Watchdog{} }
func (m *Watchdog) String() string { return proto.CompactTextString(m) }
func (*Watchdog) ProtoMessage()    {}
func (*Watchdog) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{3}
}
func (m *Watchdog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Watchdog.Unmarshal(m, b)
}
func (m *Watchdog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Watchdog.Marshal(b, m, deterministic)
}
func (dst *Watchdog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Watchdog.Merge(dst, src)
}
func (m *Watchdog) XXX_Size() int {
	return xxx_messageInfo_Watchdog.Size(m)
}
func (m *Watchdog) XXX_DiscardUnknown() {
	xxx_messageInfo_Watchdog.DiscardUnknown(m)
}

var xxx_messageInfo_Watchdog proto.InternalMessageInfo

func (m *Watchdog) GetMissTimeout() *duration.Duration {
	if m != nil {
		return m.MissTimeout
	}
	return nil
}

func (m *Watchdog) GetMegamissTimeout() *duration.Duration {
	if m != nil {
		return m.MegamissTimeout
	}
	return nil
}

func (m *Watchdog) GetKillTimeout() *duration.Duration {
	if m != nil {
		return m.KillTimeout
	}
	return nil
}

func (m *Watchdog) GetMultikillTimeout() *duration.Duration {
	if m != nil {
		return m.MultikillTimeout
	}
	return nil
}

// Runtime :ref:`configuration overview <config_runtime>`.
type Runtime struct {
	// The implementation assumes that the file system tree is accessed via a
	// symbolic link. An atomic link swap is used when a new tree should be
	// switched to. This parameter specifies the path to the symbolic link. Envoy
	// will watch the location for changes and reload the file system tree when
	// they happen.
	SymlinkRoot string `protobuf:"bytes,1,opt,name=symlink_root,json=symlinkRoot" json:"symlink_root,omitempty"`
	// Specifies the subdirectory to load within the root directory. This is
	// useful if multiple systems share the same delivery mechanism. Envoy
	// configuration elements can be contained in a dedicated subdirectory.
	Subdirectory string `protobuf:"bytes,2,opt,name=subdirectory" json:"subdirectory,omitempty"`
	// Specifies an optional subdirectory to load within the root directory. If
	// specified and the directory exists, configuration values within this
	// directory will override those found in the primary subdirectory. This is
	// useful when Envoy is deployed across many different types of servers.
	// Sometimes it is useful to have a per service cluster directory for runtime
	// configuration. See below for exactly how the override directory is used.
	OverrideSubdirectory string   `protobuf:"bytes,3,opt,name=override_subdirectory,json=overrideSubdirectory" json:"override_subdirectory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_bootstrap_e402d312e52c7a6a, []int{4}
}
func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (dst *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(dst, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetSymlinkRoot() string {
	if m != nil {
		return m.SymlinkRoot
	}
	return ""
}

func (m *Runtime) GetSubdirectory() string {
	if m != nil {
		return m.Subdirectory
	}
	return ""
}

func (m *Runtime) GetOverrideSubdirectory() string {
	if m != nil {
		return m.OverrideSubdirectory
	}
	return ""
}

func init() {
	proto.RegisterType((*Bootstrap)(nil), "envoy.config.bootstrap.v2.Bootstrap")
	proto.RegisterType((*Bootstrap_StaticResources)(nil), "envoy.config.bootstrap.v2.Bootstrap.StaticResources")
	proto.RegisterType((*Bootstrap_DynamicResources)(nil), "envoy.config.bootstrap.v2.Bootstrap.DynamicResources")
	proto.RegisterType((*Bootstrap_DynamicResources_DeprecatedV1)(nil), "envoy.config.bootstrap.v2.Bootstrap.DynamicResources.DeprecatedV1")
	proto.RegisterType((*Admin)(nil), "envoy.config.bootstrap.v2.Admin")
	proto.RegisterType((*ClusterManager)(nil), "envoy.config.bootstrap.v2.ClusterManager")
	proto.RegisterType((*ClusterManager_OutlierDetection)(nil), "envoy.config.bootstrap.v2.ClusterManager.OutlierDetection")
	proto.RegisterType((*Watchdog)(nil), "envoy.config.bootstrap.v2.Watchdog")
	proto.RegisterType((*Runtime)(nil), "envoy.config.bootstrap.v2.Runtime")
}

func init() {
	proto.RegisterFile("envoy/config/bootstrap/v2/bootstrap.proto", fileDescriptor_bootstrap_e402d312e52c7a6a)
}

var fileDescriptor_bootstrap_e402d312e52c7a6a = []byte{
	// 1162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x59, 0xdb, 0x69, 0xe2, 0x67, 0x37, 0x71, 0x46, 0x69, 0xbb, 0xb5, 0x68, 0x93, 0xba,
	0x45, 0x6a, 0x45, 0xb5, 0x56, 0x5c, 0x10, 0xa5, 0xaa, 0x8a, 0xea, 0x46, 0x45, 0x48, 0x25, 0x85,
	0x49, 0x05, 0x12, 0x97, 0xd5, 0x78, 0x77, 0xb2, 0x19, 0x65, 0x77, 0xc7, 0x9a, 0x99, 0x5d, 0x94,
	0x2b, 0x47, 0xc4, 0x01, 0x21, 0x0e, 0xdc, 0xf9, 0x2f, 0x38, 0xc1, 0x8d, 0xbf, 0x02, 0xce, 0xfc,
	0x17, 0x68, 0x7e, 0xac, 0x9d, 0x75, 0xdc, 0x24, 0x70, 0xdb, 0x7d, 0xf3, 0xfd, 0x7c, 0xdf, 0xec,
	0xcc, 0x7b, 0x6f, 0xe1, 0x01, 0xcd, 0x4b, 0x7e, 0x32, 0x8c, 0x78, 0x7e, 0xc8, 0x92, 0xe1, 0x84,
	0x73, 0x25, 0x95, 0x20, 0xd3, 0x61, 0x39, 0x9a, 0xbf, 0x04, 0x53, 0xc1, 0x15, 0x47, 0x37, 0x8d,
	0x34, 0xb0, 0xd2, 0x60, 0xbe, 0x5a, 0x8e, 0xfa, 0xdb, 0xd6, 0x85, 0x4c, 0x99, 0x06, 0x23, 0x2e,
	0xe8, 0x90, 0xc4, 0xb1, 0xa0, 0x52, 0x5a, 0xb6, 0xff, 0xee, 0x59, 0xc1, 0x84, 0x48, 0xba, 0x74,
	0x95, 0x14, 0xea, 0x68, 0x18, 0x51, 0xa1, 0xdc, 0xea, 0x7b, 0x67, 0x59, 0xbb, 0x87, 0x50, 0xf2,
	0x42, 0x44, 0x95, 0xc9, 0xf5, 0xba, 0x2c, 0x96, 0x4b, 0xe3, 0xe9, 0x2c, 0x7e, 0xa7, 0xf6, 0xe5,
	0x4a, 0x90, 0x88, 0x6a, 0x81, 0x79, 0x70, 0x92, 0xbb, 0x35, 0x49, 0x46, 0x95, 0x60, 0x91, 0xd4,
	0x22, 0xa9, 0x88, 0x92, 0x4b, 0x45, 0x82, 0x28, 0x9a, 0xb2, 0x8c, 0x29, 0x2d, 0x13, 0x69, 0x25,
	0xba, 0x9d, 0x70, 0x9e, 0xa4, 0x74, 0x68, 0xde, 0x26, 0xc5, 0xe1, 0x30, 0x2e, 0x04, 0x51, 0x8c,
	0xe7, 0x6e, 0xfd, 0x46, 0x49, 0x52, 0x16, 0x13, 0x45, 0x87, 0xd5, 0x83, 0x5b, 0xd8, 0x4a, 0x78,
	0xc2, 0xcd, 0xe3, 0x50, 0x3f, 0xd9, 0xe8, 0xe0, 0x87, 0x2e, 0xb4, 0xc7, 0xd5, 0x05, 0xa0, 0xf7,
	0xa1, 0x95, 0xf3, 0x98, 0xfa, 0xde, 0x8e, 0x77, 0xbf, 0x33, 0xba, 0x11, 0xd8, 0x7b, 0x22, 0x53,
	0x16, 0x94, 0xa3, 0x40, 0x9f, 0x57, 0xb0, 0xcf, 0x63, 0x8a, 0x8d, 0x08, 0x85, 0xd0, 0xd3, 0xbb,
	0x67, 0x51, 0x28, 0xa8, 0x3d, 0x3f, 0xe9, 0x37, 0x0c, 0xf8, 0x41, 0xf0, 0xd6, 0x0b, 0x0e, 0x66,
	0xc9, 0x82, 0x03, 0x03, 0xe3, 0x8a, 0xc5, 0x1b, 0xb2, 0x1e, 0x40, 0x13, 0xd8, 0x8c, 0x4f, 0x72,
	0x92, 0xd5, 0x32, 0x34, 0x4d, 0x86, 0x0f, 0x2f, 0x95, 0x61, 0xcf, 0xd2, 0xf3, 0x14, 0xbd, 0x78,
	0x21, 0x82, 0x30, 0x6c, 0x44, 0x69, 0x21, 0x15, 0x15, 0x61, 0x46, 0x72, 0x92, 0x50, 0xe1, 0xb7,
	0x4c, 0x86, 0x07, 0xe7, 0x64, 0x78, 0x61, 0x89, 0xcf, 0x2d, 0x80, 0xd7, 0xa3, 0xda, 0x3b, 0xba,
	0x05, 0x70, 0x98, 0x92, 0x44, 0x86, 0x53, 0xa2, 0x8e, 0xfc, 0x95, 0x1d, 0xef, 0x7e, 0x1b, 0xb7,
	0x4d, 0xe4, 0x0b, 0xa2, 0x8e, 0xd0, 0x0b, 0xe8, 0x98, 0x5b, 0x0f, 0x25, 0xcb, 0x8f, 0xa5, 0x7f,
	0x65, 0xa7, 0x79, 0xbf, 0x33, 0x1a, 0xd4, 0xd3, 0xb9, 0x0a, 0xd1, 0xc9, 0xf4, 0x31, 0xc9, 0x03,
	0x96, 0x1f, 0x63, 0x90, 0xd5, 0xa3, 0x44, 0x9f, 0x42, 0xd7, 0x9a, 0x58, 0xc0, 0xbf, 0x6a, 0x36,
	0x7d, 0xef, 0x7c, 0x97, 0x17, 0x26, 0x8c, 0x6d, 0x7a, 0xfb, 0x82, 0xbe, 0x84, 0x2d, 0x6b, 0x74,
	0x98, 0x16, 0xf2, 0x28, 0x64, 0xb9, 0xa2, 0xa2, 0x24, 0xa9, 0xbf, 0x6a, 0x0c, 0x6f, 0x06, 0xb6,
	0xdc, 0x82, 0xaa, 0xdc, 0x82, 0x3d, 0x57, 0x6e, 0xe3, 0xd6, 0x2f, 0x7f, 0x6f, 0x7b, 0x18, 0x19,
	0xf8, 0xa5, 0x66, 0x3f, 0x73, 0x28, 0xfa, 0x04, 0xd6, 0xbe, 0x25, 0x2a, 0x3a, 0x8a, 0x79, 0xe2,
	0xaf, 0x19, 0x9b, 0xbb, 0xe7, 0x1c, 0xe6, 0xd7, 0x4e, 0x8a, 0x67, 0x10, 0x7a, 0x0c, 0xab, 0xba,
	0x79, 0x58, 0x9e, 0xf8, 0x6d, 0xc3, 0xdf, 0xae, 0xf3, 0xb6, 0xb3, 0xca, 0x51, 0xf0, 0xc6, 0xaa,
	0x70, 0x25, 0x47, 0x21, 0x20, 0xdd, 0x37, 0xa1, 0x69, 0x9c, 0x50, 0x52, 0x51, 0xb2, 0x88, 0xfa,
	0x60, 0x4c, 0x76, 0xeb, 0x26, 0xb3, 0xfe, 0xd2, 0x46, 0x98, 0x28, 0xfa, 0x4a, 0xbf, 0x1c, 0x58,
	0xc4, 0x9d, 0x54, 0x4f, 0x2c, 0xc4, 0xd1, 0x53, 0x58, 0x15, 0x45, 0xae, 0x58, 0x46, 0xfd, 0x8e,
	0x71, 0x1d, 0x9c, 0xf3, 0x69, 0xd8, 0x2a, 0x71, 0x85, 0xa0, 0x97, 0xb0, 0x42, 0xe2, 0x8c, 0xe5,
	0x7e, 0xd7, 0xb0, 0x3b, 0xe7, 0xb0, 0xcf, 0xb5, 0x6e, 0xbc, 0xfe, 0xe7, 0x5f, 0xdb, 0xef, 0xfc,
	0xf6, 0xcf, 0xef, 0xcd, 0x95, 0xef, 0xbd, 0x46, 0xcf, 0xc3, 0x16, 0xef, 0xff, 0xe1, 0xc1, 0xc6,
	0x42, 0xfb, 0xa0, 0x27, 0xd0, 0x4e, 0x99, 0x54, 0x34, 0xa7, 0x42, 0xfa, 0x9e, 0x29, 0xaa, 0xeb,
	0xf5, 0x06, 0x7e, 0xe5, 0x96, 0xc7, 0x2d, 0xed, 0x8a, 0xe7, 0x72, 0xf4, 0x11, 0xac, 0xb9, 0x1a,
	0xd6, 0x2d, 0xac, 0xd1, 0x6b, 0x75, 0xd4, 0x55, 0xbc, 0x23, 0x67, 0x62, 0xf4, 0x31, 0xac, 0x4a,
	0x1a, 0x09, 0xaa, 0x74, 0x63, 0x36, 0x4d, 0xc1, 0xd4, 0x38, 0x3d, 0x81, 0x83, 0x03, 0xa3, 0x70,
	0x6c, 0xa5, 0xef, 0xff, 0xd4, 0x84, 0xde, 0x62, 0x83, 0xa2, 0x67, 0x00, 0x69, 0x3c, 0x2b, 0x6a,
	0x3b, 0x86, 0xb6, 0x97, 0x8c, 0x21, 0x7b, 0x3f, 0x07, 0x86, 0xc2, 0xed, 0x34, 0xae, 0xaa, 0xf9,
	0x19, 0x40, 0x34, 0xe7, 0x1b, 0x97, 0xe4, 0xa3, 0x19, 0xff, 0x1c, 0x80, 0xcc, 0xf9, 0x66, 0xed,
	0x86, 0x4f, 0xf3, 0xcf, 0xa7, 0xac, 0x6e, 0x41, 0x66, 0x16, 0xc7, 0x70, 0x35, 0xa6, 0x53, 0x41,
	0x23, 0xa2, 0x68, 0x1c, 0x96, 0xbb, 0x6e, 0x9e, 0x8c, 0xff, 0xd7, 0xc4, 0x0a, 0xf6, 0x66, 0x56,
	0x5f, 0xed, 0x8e, 0x1b, 0xbe, 0x87, 0xbb, 0xf1, 0xa9, 0x48, 0x7f, 0x1f, 0xba, 0xa7, 0x15, 0xfa,
	0xfb, 0xe5, 0x7f, 0x3f, 0x3f, 0x59, 0x6d, 0x7e, 0xf0, 0xab, 0x07, 0x2b, 0xa6, 0xf2, 0xd0, 0x2e,
	0x6c, 0x90, 0x28, 0xa2, 0x52, 0x86, 0x29, 0x4f, 0xec, 0x24, 0xd3, 0x76, 0xed, 0x71, 0x5b, 0x97,
	0x63, 0x4b, 0x34, 0x76, 0x3c, 0x7c, 0xd5, 0x2a, 0x5e, 0xf1, 0xc4, 0x0c, 0xb6, 0x3b, 0xd0, 0x9d,
	0x0a, 0x7e, 0xc8, 0x52, 0x6a, 0xf5, 0x0d, 0x33, 0xf9, 0x3a, 0x2e, 0x66, 0x24, 0x7b, 0xb0, 0xea,
	0x7e, 0xe7, 0xee, 0x70, 0xfb, 0xcb, 0x0e, 0xd7, 0x2a, 0xce, 0x14, 0x7f, 0x85, 0x0e, 0xbe, 0x6b,
	0xc2, 0x7a, 0x7d, 0x06, 0xa3, 0x87, 0x80, 0x52, 0x1e, 0x91, 0x34, 0xac, 0xa6, 0x79, 0x4e, 0x32,
	0xfb, 0x1f, 0x6b, 0xe3, 0x9e, 0x59, 0x71, 0xc0, 0x3e, 0xc9, 0x28, 0x4a, 0x60, 0x93, 0x17, 0x2a,
	0x65, 0x54, 0x84, 0x31, 0x55, 0x34, 0xd2, 0x03, 0xcd, 0x55, 0xcb, 0x93, 0x4b, 0xcf, 0xfd, 0xe0,
	0xb5, 0xb5, 0xd8, 0xab, 0x1c, 0x70, 0x8f, 0x2f, 0x44, 0xd0, 0x6b, 0xd8, 0x2a, 0xa6, 0x52, 0x09,
	0x4a, 0xb2, 0x70, 0xc2, 0xf2, 0xb8, 0x5e, 0x59, 0xb7, 0x96, 0x7c, 0xfc, 0x98, 0xe5, 0xb1, 0x9b,
	0x3e, 0xa8, 0x42, 0xe7, 0x31, 0xb4, 0x0f, 0x9b, 0x29, 0x27, 0x71, 0x58, 0x1b, 0xfe, 0xad, 0x4b,
	0xd7, 0xe9, 0x86, 0x86, 0x4f, 0xfd, 0x0b, 0xfa, 0x8f, 0xa1, 0xb7, 0xf8, 0x19, 0xe8, 0x1e, 0xac,
	0xd3, 0x92, 0xe6, 0x6a, 0xe1, 0xe6, 0x71, 0xd7, 0x44, 0xdd, 0x6d, 0x0f, 0x7e, 0x6e, 0xc0, 0x5a,
	0x35, 0xbb, 0xd1, 0x53, 0xe8, 0x66, 0x4c, 0xca, 0x50, 0x4f, 0x39, 0x5e, 0x28, 0x57, 0x79, 0x6f,
	0xff, 0x7b, 0xe0, 0x8e, 0x96, 0xbf, 0xb1, 0x6a, 0xb4, 0x07, 0xbd, 0x8c, 0x26, 0xa4, 0xe6, 0xd0,
	0xb8, 0xc8, 0x61, 0xa3, 0x42, 0x2a, 0x97, 0xa7, 0xd0, 0x3d, 0x66, 0x69, 0x3a, 0x73, 0x68, 0x5e,
	0xb8, 0x07, 0x2d, 0xaf, 0xe8, 0x97, 0xb0, 0x99, 0x15, 0xa9, 0x62, 0x35, 0x8b, 0xd6, 0x45, 0x16,
	0xbd, 0x19, 0xe3, 0x7c, 0x06, 0x3f, 0x7a, 0xb0, 0xea, 0xe6, 0x3e, 0x7a, 0x08, 0x5d, 0x79, 0x92,
	0xa5, 0x2c, 0x3f, 0x0e, 0x05, 0xe7, 0xea, 0x6c, 0x03, 0x75, 0xdc, 0x32, 0xe6, 0x5c, 0xa1, 0x01,
	0x74, 0x65, 0x31, 0x89, 0x99, 0xa0, 0x91, 0xe2, 0xe2, 0xc4, 0xb5, 0x4f, 0x2d, 0x86, 0x1e, 0xc1,
	0x35, 0x5e, 0x52, 0x21, 0x58, 0x4c, 0xc3, 0x9a, 0xb8, 0x69, 0xc4, 0x5b, 0xd5, 0xe2, 0xc1, 0xa9,
	0xb5, 0x71, 0xeb, 0x9b, 0x46, 0x39, 0x9a, 0x5c, 0x31, 0xbb, 0x7f, 0xf4, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x77, 0xb3, 0xf5, 0x2f, 0xa8, 0x0b, 0x00, 0x00,
}
