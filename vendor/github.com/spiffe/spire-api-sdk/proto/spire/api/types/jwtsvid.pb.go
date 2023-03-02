// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: spire/api/types/jwtsvid.proto

package types

import (
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

// JWT SPIFFE Verifiable Identity Document. It contains the raw JWT token
// as well as a few denormalized fields for convenience.
type JWTSVID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The serialized JWT token.
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// The SPIFFE ID of the JWT-SVID.
	Id *SPIFFEID `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Expiration timestamp (seconds since Unix epoch).
	ExpiresAt int64 `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Issuance timestamp (seconds since Unix epoch).
	IssuedAt int64 `protobuf:"varint,4,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
}

func (x *JWTSVID) Reset() {
	*x = JWTSVID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_api_types_jwtsvid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTSVID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTSVID) ProtoMessage() {}

func (x *JWTSVID) ProtoReflect() protoreflect.Message {
	mi := &file_spire_api_types_jwtsvid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTSVID.ProtoReflect.Descriptor instead.
func (*JWTSVID) Descriptor() ([]byte, []int) {
	return file_spire_api_types_jwtsvid_proto_rawDescGZIP(), []int{0}
}

func (x *JWTSVID) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JWTSVID) GetId() *SPIFFEID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *JWTSVID) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

func (x *JWTSVID) GetIssuedAt() int64 {
	if x != nil {
		return x.IssuedAt
	}
	return 0
}

var File_spire_api_types_jwtsvid_proto protoreflect.FileDescriptor

var file_spire_api_types_jwtsvid_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6a, 0x77, 0x74, 0x73, 0x76, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x1e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x86, 0x01, 0x0a, 0x07, 0x4a, 0x57, 0x54, 0x53, 0x56, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x53, 0x50, 0x49, 0x46, 0x46, 0x45, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_api_types_jwtsvid_proto_rawDescOnce sync.Once
	file_spire_api_types_jwtsvid_proto_rawDescData = file_spire_api_types_jwtsvid_proto_rawDesc
)

func file_spire_api_types_jwtsvid_proto_rawDescGZIP() []byte {
	file_spire_api_types_jwtsvid_proto_rawDescOnce.Do(func() {
		file_spire_api_types_jwtsvid_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_api_types_jwtsvid_proto_rawDescData)
	})
	return file_spire_api_types_jwtsvid_proto_rawDescData
}

var file_spire_api_types_jwtsvid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spire_api_types_jwtsvid_proto_goTypes = []interface{}{
	(*JWTSVID)(nil),  // 0: spire.api.types.JWTSVID
	(*SPIFFEID)(nil), // 1: spire.api.types.SPIFFEID
}
var file_spire_api_types_jwtsvid_proto_depIdxs = []int32{
	1, // 0: spire.api.types.JWTSVID.id:type_name -> spire.api.types.SPIFFEID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spire_api_types_jwtsvid_proto_init() }
func file_spire_api_types_jwtsvid_proto_init() {
	if File_spire_api_types_jwtsvid_proto != nil {
		return
	}
	file_spire_api_types_spiffeid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spire_api_types_jwtsvid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTSVID); i {
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
			RawDescriptor: file_spire_api_types_jwtsvid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spire_api_types_jwtsvid_proto_goTypes,
		DependencyIndexes: file_spire_api_types_jwtsvid_proto_depIdxs,
		MessageInfos:      file_spire_api_types_jwtsvid_proto_msgTypes,
	}.Build()
	File_spire_api_types_jwtsvid_proto = out.File
	file_spire_api_types_jwtsvid_proto_rawDesc = nil
	file_spire_api_types_jwtsvid_proto_goTypes = nil
	file_spire_api_types_jwtsvid_proto_depIdxs = nil
}
