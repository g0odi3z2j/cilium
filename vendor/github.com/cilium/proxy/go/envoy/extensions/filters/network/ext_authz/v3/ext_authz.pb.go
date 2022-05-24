// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.19.1
// source: envoy/extensions/filters/network/ext_authz/v3/ext_authz.proto

package ext_authzv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	v31 "github.com/cilium/proxy/go/envoy/type/matcher/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// External Authorization filter calls out to an external service over the
// gRPC Authorization API defined by
// :ref:`CheckRequest <envoy_v3_api_msg_service.auth.v3.CheckRequest>`.
// A failed check will cause this filter to close the TCP connection.
// [#next-free-field: 8]
type ExtAuthz struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The external authorization gRPC service configuration.
	// The default timeout is set to 200ms by this filter.
	GrpcService *v3.GrpcService `protobuf:"bytes,2,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// The filter's behaviour in case the external authorization service does
	// not respond back. When it is set to true, Envoy will also allow traffic in case of
	// communication failure between authorization service and the proxy.
	// Defaults to false.
	FailureModeAllow bool `protobuf:"varint,3,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	// Specifies if the peer certificate is sent to the external service.
	//
	// When this field is true, Envoy will include the peer X.509 certificate, if available, in the
	// :ref:`certificate<envoy_v3_api_field_service.auth.v3.AttributeContext.Peer.certificate>`.
	IncludePeerCertificate bool `protobuf:"varint,4,opt,name=include_peer_certificate,json=includePeerCertificate,proto3" json:"include_peer_certificate,omitempty"`
	// API version for ext_authz transport protocol. This describes the ext_authz gRPC endpoint and
	// version of Check{Request,Response} used on the wire.
	TransportApiVersion v3.ApiVersion `protobuf:"varint,5,opt,name=transport_api_version,json=transportApiVersion,proto3,enum=envoy.config.core.v3.ApiVersion" json:"transport_api_version,omitempty"`
	// Specifies if the filter is enabled with metadata matcher.
	// If this field is not specified, the filter will be enabled for all requests.
	FilterEnabledMetadata *v31.MetadataMatcher `protobuf:"bytes,6,opt,name=filter_enabled_metadata,json=filterEnabledMetadata,proto3" json:"filter_enabled_metadata,omitempty"`
	// Optional labels that will be passed to :ref:`labels<envoy_v3_api_field_service.auth.v3.AttributeContext.Peer.labels>` in
	// :ref:`destination<envoy_v3_api_field_service.auth.v3.AttributeContext.destination>`.
	// The labels will be read from :ref:`metadata<envoy_v3_api_msg_config.core.v3.Node>` with the specified key.
	BootstrapMetadataLabelsKey string `protobuf:"bytes,7,opt,name=bootstrap_metadata_labels_key,json=bootstrapMetadataLabelsKey,proto3" json:"bootstrap_metadata_labels_key,omitempty"`
}

func (x *ExtAuthz) Reset() {
	*x = ExtAuthz{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtAuthz) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtAuthz) ProtoMessage() {}

func (x *ExtAuthz) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtAuthz.ProtoReflect.Descriptor instead.
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDescGZIP(), []int{0}
}

func (x *ExtAuthz) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *ExtAuthz) GetGrpcService() *v3.GrpcService {
	if x != nil {
		return x.GrpcService
	}
	return nil
}

func (x *ExtAuthz) GetFailureModeAllow() bool {
	if x != nil {
		return x.FailureModeAllow
	}
	return false
}

func (x *ExtAuthz) GetIncludePeerCertificate() bool {
	if x != nil {
		return x.IncludePeerCertificate
	}
	return false
}

func (x *ExtAuthz) GetTransportApiVersion() v3.ApiVersion {
	if x != nil {
		return x.TransportApiVersion
	}
	return v3.ApiVersion_AUTO
}

func (x *ExtAuthz) GetFilterEnabledMetadata() *v31.MetadataMatcher {
	if x != nil {
		return x.FilterEnabledMetadata
	}
	return nil
}

func (x *ExtAuthz) GetBootstrapMetadataLabelsKey() string {
	if x != nil {
		return x.BootstrapMetadataLabelsKey
	}
	return ""
}

var File_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x33, 0x2f,
	0x65, 0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x33, 0x1a, 0x28,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9f, 0x04, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x12,
	0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x44, 0x0a, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x38, 0x0a,
	0x18, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x16, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x50, 0x65, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x5e, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x70, 0x69,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x17, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x15, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x1d, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a,
	0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x4b, 0x65, 0x79, 0x3a, 0x38, 0x9a, 0xc5, 0x88, 0x1e,
	0x33, 0x0a, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65,
	0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x7a, 0x42, 0xb8, 0x01, 0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x33, 0x42, 0x0d, 0x45, 0x78, 0x74, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x65,
	0x78, 0x74, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x33, 0x3b, 0x65, 0x78, 0x74, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDescData = file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDesc
)

func file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDescData
}

var file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_goTypes = []interface{}{
	(*ExtAuthz)(nil),            // 0: envoy.extensions.filters.network.ext_authz.v3.ExtAuthz
	(*v3.GrpcService)(nil),      // 1: envoy.config.core.v3.GrpcService
	(v3.ApiVersion)(0),          // 2: envoy.config.core.v3.ApiVersion
	(*v31.MetadataMatcher)(nil), // 3: envoy.type.matcher.v3.MetadataMatcher
}
var file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.network.ext_authz.v3.ExtAuthz.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	2, // 1: envoy.extensions.filters.network.ext_authz.v3.ExtAuthz.transport_api_version:type_name -> envoy.config.core.v3.ApiVersion
	3, // 2: envoy.extensions.filters.network.ext_authz.v3.ExtAuthz.filter_enabled_metadata:type_name -> envoy.type.matcher.v3.MetadataMatcher
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_init() }
func file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_init() {
	if File_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtAuthz); i {
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
			RawDescriptor: file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto = out.File
	file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_rawDesc = nil
	file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_goTypes = nil
	file_envoy_extensions_filters_network_ext_authz_v3_ext_authz_proto_depIdxs = nil
}
