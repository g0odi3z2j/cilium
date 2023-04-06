// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.21.12
// source: envoy/extensions/transport_sockets/tls/v3/tls_spiffe_validator_config.proto

package tlsv3

import (
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
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

// Configuration specific to the `SPIFFE <https://github.com/spiffe/spiffe>`_ certificate validator.
//
// Example:
//
// .. validated-code-block:: yaml
//
//	:type-name: envoy.extensions.transport_sockets.tls.v3.CertificateValidationContext
//
//	custom_validator_config:
//	  name: envoy.tls.cert_validator.spiffe
//	  typed_config:
//	    "@type": type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.SPIFFECertValidatorConfig
//	    trust_domains:
//	    - name: foo.com
//	      trust_bundle:
//	        filename: "foo.pem"
//	    - name: envoy.com
//	      trust_bundle:
//	        filename: "envoy.pem"
//
// In this example, a presented peer certificate whose SAN matches “spiffe://foo.com/**“ is validated against
// the "foo.pem" x.509 certificate. All the trust bundles are isolated from each other, so no trust domain can mint
// a SVID belonging to another trust domain. That means, in this example, a SVID signed by “envoy.com“'s CA with “spiffe://foo.com/**“
// SAN would be rejected since Envoy selects the trust bundle according to the presented SAN before validate the certificate.
//
// Note that SPIFFE validator inherits and uses the following options from :ref:`CertificateValidationContext <envoy_v3_api_msg_extensions.transport_sockets.tls.v3.CertificateValidationContext>`.
//
// - :ref:`allow_expired_certificate <envoy_v3_api_field_extensions.transport_sockets.tls.v3.CertificateValidationContext.allow_expired_certificate>` to allow expired certificates.
// - :ref:`match_typed_subject_alt_names <envoy_v3_api_field_extensions.transport_sockets.tls.v3.CertificateValidationContext.match_typed_subject_alt_names>` to match **URI** SAN of certificates. Unlike the default validator, SPIFFE validator only matches **URI** SAN (which equals to SVID in SPIFFE terminology) and ignore other SAN types.
type SPIFFECertValidatorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field specifies trust domains used for validating incoming X.509-SVID(s).
	TrustDomains []*SPIFFECertValidatorConfig_TrustDomain `protobuf:"bytes,1,rep,name=trust_domains,json=trustDomains,proto3" json:"trust_domains,omitempty"`
}

func (x *SPIFFECertValidatorConfig) Reset() {
	*x = SPIFFECertValidatorConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SPIFFECertValidatorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPIFFECertValidatorConfig) ProtoMessage() {}

func (x *SPIFFECertValidatorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPIFFECertValidatorConfig.ProtoReflect.Descriptor instead.
func (*SPIFFECertValidatorConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescGZIP(), []int{0}
}

func (x *SPIFFECertValidatorConfig) GetTrustDomains() []*SPIFFECertValidatorConfig_TrustDomain {
	if x != nil {
		return x.TrustDomains
	}
	return nil
}

type SPIFFECertValidatorConfig_TrustDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the trust domain, “example.com“, “foo.bar.gov“ for example.
	// Note that this must *not* have "spiffe://" prefix.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specify a data source holding x.509 trust bundle used for validating incoming SVID(s) in this trust domain.
	TrustBundle *v3.DataSource `protobuf:"bytes,2,opt,name=trust_bundle,json=trustBundle,proto3" json:"trust_bundle,omitempty"`
}

func (x *SPIFFECertValidatorConfig_TrustDomain) Reset() {
	*x = SPIFFECertValidatorConfig_TrustDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SPIFFECertValidatorConfig_TrustDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPIFFECertValidatorConfig_TrustDomain) ProtoMessage() {}

func (x *SPIFFECertValidatorConfig_TrustDomain) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPIFFECertValidatorConfig_TrustDomain.ProtoReflect.Descriptor instead.
func (*SPIFFECertValidatorConfig_TrustDomain) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SPIFFECertValidatorConfig_TrustDomain) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SPIFFECertValidatorConfig_TrustDomain) GetTrustBundle() *v3.DataSource {
	if x != nil {
		return x.TrustBundle
	}
	return nil
}

var File_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto protoreflect.FileDescriptor

var file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x6c, 0x73, 0x5f,
	0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x19, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x43, 0x65, 0x72, 0x74,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x7f, 0x0a, 0x0d, 0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x43, 0x65, 0x72, 0x74, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x54, 0x72, 0x75,
	0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x0c, 0x74, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x1a, 0x6f, 0x0a, 0x0b, 0x54, 0x72, 0x75, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0c,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x42, 0xba, 0x01, 0x0a, 0x37, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x76, 0x33, 0x42, 0x1d, 0x54,
	0x6c, 0x73, 0x53, 0x70, 0x69, 0x66, 0x66, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33,
	0x3b, 0x74, 0x6c, 0x73, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescOnce sync.Once
	file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescData = file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDesc
)

func file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescGZIP() []byte {
	file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescData)
	})
	return file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDescData
}

var file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_goTypes = []interface{}{
	(*SPIFFECertValidatorConfig)(nil),             // 0: envoy.extensions.transport_sockets.tls.v3.SPIFFECertValidatorConfig
	(*SPIFFECertValidatorConfig_TrustDomain)(nil), // 1: envoy.extensions.transport_sockets.tls.v3.SPIFFECertValidatorConfig.TrustDomain
	(*v3.DataSource)(nil),                         // 2: envoy.config.core.v3.DataSource
}
var file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.transport_sockets.tls.v3.SPIFFECertValidatorConfig.trust_domains:type_name -> envoy.extensions.transport_sockets.tls.v3.SPIFFECertValidatorConfig.TrustDomain
	2, // 1: envoy.extensions.transport_sockets.tls.v3.SPIFFECertValidatorConfig.TrustDomain.trust_bundle:type_name -> envoy.config.core.v3.DataSource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_init() }
func file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_init() {
	if File_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SPIFFECertValidatorConfig); i {
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
		file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SPIFFECertValidatorConfig_TrustDomain); i {
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
			RawDescriptor: file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_msgTypes,
	}.Build()
	File_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto = out.File
	file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_rawDesc = nil
	file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_goTypes = nil
	file_envoy_extensions_transport_sockets_tls_v3_tls_spiffe_validator_config_proto_depIdxs = nil
}
