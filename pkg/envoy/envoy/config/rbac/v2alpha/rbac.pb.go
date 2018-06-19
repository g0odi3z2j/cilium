// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v2alpha/rbac.proto

package v2alpha

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/core"
import route "github.com/cilium/cilium/pkg/envoy/envoy/api/v2/route"
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

// Should we do safe-list or block-list style access control?
type RBAC_Action int32

const (
	// The policies grant access to principals. The rest is denied. This is safe-list style
	// access control. This is the default type.
	RBAC_ALLOW RBAC_Action = 0
	// The policies deny access to principals. The rest is allowed. This is block-list style
	// access control.
	RBAC_DENY RBAC_Action = 1
)

var RBAC_Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}
var RBAC_Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x RBAC_Action) String() string {
	return proto.EnumName(RBAC_Action_name, int32(x))
}
func (RBAC_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_rbac_713dfdae9a03bbb7, []int{0, 0}
}

// Role Based Access Control (RBAC) provides service-level and method-level access control for a
// service. RBAC policies are additive. The policies are examined in order. A request is allowed
// once a matching policy is found (suppose the `action` is ALLOW).
//
// Here is an example of RBAC configuration. It has two policies:
//
// * Service account "cluster.local/ns/default/sa/admin" has full access (empty permission entry
//   means full access) to the service.
//
// * Any user (empty principal entry means any user) can read ("GET") the service at paths with
//   prefix "/products" or suffix "/reviews" when request header "version" set to either "v1" or
//   "v2".
//
//  .. code-block:: yaml
//
//   action: ALLOW
//   policies:
//     "service-admin":
//       permissions:
//         - any: true
//       principals:
//         - authenticated: { name: "cluster.local/ns/default/sa/admin" }
//         - authenticated: { name: "cluster.local/ns/default/sa/superuser" }
//     "product-viewer":
//       permissions:
//           - and_rules:
//               rules:
//                 - header: { name: ":method", exact_match: "GET" }
//                 - header: { name: ":path", regex_match: "/products(/.*)?" }
//                 - or_rules:
//                     rules:
//                       - destination_port: 80
//                       - destination_port: 443
//       principals:
//         - any: true
//
type RBAC struct {
	// The action to take if a policy matches. The request is allowed if and only if:
	//
	//   * `action` is "ALLOWED" and at least one policy matches
	//   * `action` is "DENY" and none of the policies match
	Action RBAC_Action `protobuf:"varint,1,opt,name=action,proto3,enum=envoy.config.rbac.v2alpha.RBAC_Action" json:"action,omitempty"`
	// Maps from policy name to policy. A match occurs when at least one policy matches the request.
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_713dfdae9a03bbb7, []int{0}
}
func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (dst *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(dst, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetAction() RBAC_Action {
	if m != nil {
		return m.Action
	}
	return RBAC_ALLOW
}

func (m *RBAC) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

// Policy specifies a role and the principals that are assigned/denied the role. A policy matches if
// and only if at least one of its permissions match the action taking place AND at least one of its
// principals match the downstream.
type Policy struct {
	// Required. The set of permissions that define a role. Each permission is matched with OR
	// semantics. To match all actions for this policy, a single Permission with the `any` field set
	// to true should be used.
	Permissions []*Permission `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Required. The set of principals that are assigned/denied the role based on “action”. Each
	// principal is matched with OR semantics. To match all downstreams for this policy, a single
	// Principal with the `any` field set to true should be used.
	Principals           []*Principal `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_713dfdae9a03bbb7, []int{1}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

// Permission defines an action (or actions) that a principal can take.
type Permission struct {
	// Types that are valid to be assigned to Rule:
	//	*Permission_AndRules
	//	*Permission_OrRules
	//	*Permission_Any
	//	*Permission_Header
	//	*Permission_DestinationIp
	//	*Permission_DestinationPort
	Rule                 isPermission_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_713dfdae9a03bbb7, []int{2}
}
func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (dst *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(dst, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

type isPermission_Rule interface {
	isPermission_Rule()
}

type Permission_AndRules struct {
	AndRules *Permission_Set `protobuf:"bytes,1,opt,name=and_rules,json=andRules,proto3,oneof"`
}
type Permission_OrRules struct {
	OrRules *Permission_Set `protobuf:"bytes,2,opt,name=or_rules,json=orRules,proto3,oneof"`
}
type Permission_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}
type Permission_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}
type Permission_DestinationIp struct {
	DestinationIp *core.CidrRange `protobuf:"bytes,5,opt,name=destination_ip,json=destinationIp,proto3,oneof"`
}
type Permission_DestinationPort struct {
	DestinationPort uint32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,proto3,oneof"`
}

func (*Permission_AndRules) isPermission_Rule()        {}
func (*Permission_OrRules) isPermission_Rule()         {}
func (*Permission_Any) isPermission_Rule()             {}
func (*Permission_Header) isPermission_Rule()          {}
func (*Permission_DestinationIp) isPermission_Rule()   {}
func (*Permission_DestinationPort) isPermission_Rule() {}

func (m *Permission) GetRule() isPermission_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Permission) GetAndRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_AndRules); ok {
		return x.AndRules
	}
	return nil
}

func (m *Permission) GetOrRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_OrRules); ok {
		return x.OrRules
	}
	return nil
}

func (m *Permission) GetAny() bool {
	if x, ok := m.GetRule().(*Permission_Any); ok {
		return x.Any
	}
	return false
}

func (m *Permission) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetRule().(*Permission_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Permission) GetDestinationIp() *core.CidrRange {
	if x, ok := m.GetRule().(*Permission_DestinationIp); ok {
		return x.DestinationIp
	}
	return nil
}

func (m *Permission) GetDestinationPort() uint32 {
	if x, ok := m.GetRule().(*Permission_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Permission) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Permission_OneofMarshaler, _Permission_OneofUnmarshaler, _Permission_OneofSizer, []interface{}{
		(*Permission_AndRules)(nil),
		(*Permission_OrRules)(nil),
		(*Permission_Any)(nil),
		(*Permission_Header)(nil),
		(*Permission_DestinationIp)(nil),
		(*Permission_DestinationPort)(nil),
	}
}

func _Permission_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Permission)
	// rule
	switch x := m.Rule.(type) {
	case *Permission_AndRules:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AndRules); err != nil {
			return err
		}
	case *Permission_OrRules:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OrRules); err != nil {
			return err
		}
	case *Permission_Any:
		t := uint64(0)
		if x.Any {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Permission_Header:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Header); err != nil {
			return err
		}
	case *Permission_DestinationIp:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DestinationIp); err != nil {
			return err
		}
	case *Permission_DestinationPort:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.DestinationPort))
	case nil:
	default:
		return fmt.Errorf("Permission.Rule has unexpected type %T", x)
	}
	return nil
}

func _Permission_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Permission)
	switch tag {
	case 1: // rule.and_rules
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Permission_Set)
		err := b.DecodeMessage(msg)
		m.Rule = &Permission_AndRules{msg}
		return true, err
	case 2: // rule.or_rules
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Permission_Set)
		err := b.DecodeMessage(msg)
		m.Rule = &Permission_OrRules{msg}
		return true, err
	case 3: // rule.any
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Rule = &Permission_Any{x != 0}
		return true, err
	case 4: // rule.header
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(route.HeaderMatcher)
		err := b.DecodeMessage(msg)
		m.Rule = &Permission_Header{msg}
		return true, err
	case 5: // rule.destination_ip
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.CidrRange)
		err := b.DecodeMessage(msg)
		m.Rule = &Permission_DestinationIp{msg}
		return true, err
	case 6: // rule.destination_port
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Rule = &Permission_DestinationPort{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Permission_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Permission)
	// rule
	switch x := m.Rule.(type) {
	case *Permission_AndRules:
		s := proto.Size(x.AndRules)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Permission_OrRules:
		s := proto.Size(x.OrRules)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Permission_Any:
		n += 1 // tag and wire
		n += 1
	case *Permission_Header:
		s := proto.Size(x.Header)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Permission_DestinationIp:
		s := proto.Size(x.DestinationIp)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Permission_DestinationPort:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.DestinationPort))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Used in the `and_rules` and `or_rules` fields in the `rule` oneof. Depending on the context,
// each are applied with the associated behavior.
type Permission_Set struct {
	Rules                []*Permission `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Permission_Set) Reset()         { *m = Permission_Set{} }
func (m *Permission_Set) String() string { return proto.CompactTextString(m) }
func (*Permission_Set) ProtoMessage()    {}
func (*Permission_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_713dfdae9a03bbb7, []int{2, 0}
}
func (m *Permission_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission_Set.Unmarshal(m, b)
}
func (m *Permission_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission_Set.Marshal(b, m, deterministic)
}
func (dst *Permission_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission_Set.Merge(dst, src)
}
func (m *Permission_Set) XXX_Size() int {
	return xxx_messageInfo_Permission_Set.Size(m)
}
func (m *Permission_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Permission_Set proto.InternalMessageInfo

func (m *Permission_Set) GetRules() []*Permission {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Principal defines an identity or a group of identities for a downstream subject.
type Principal struct {
	// Types that are valid to be assigned to Identifier:
	//	*Principal_AndIds
	//	*Principal_OrIds
	//	*Principal_Any
	//	*Principal_Authenticated_
	//	*Principal_SourceIp
	//	*Principal_Header
	Identifier           isPrincipal_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_713dfdae9a03bbb7, []int{3}
}
func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (dst *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(dst, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

type isPrincipal_Identifier interface {
	isPrincipal_Identifier()
}

type Principal_AndIds struct {
	AndIds *Principal_Set `protobuf:"bytes,1,opt,name=and_ids,json=andIds,proto3,oneof"`
}
type Principal_OrIds struct {
	OrIds *Principal_Set `protobuf:"bytes,2,opt,name=or_ids,json=orIds,proto3,oneof"`
}
type Principal_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}
type Principal_Authenticated_ struct {
	Authenticated *Principal_Authenticated `protobuf:"bytes,4,opt,name=authenticated,proto3,oneof"`
}
type Principal_SourceIp struct {
	SourceIp *core.CidrRange `protobuf:"bytes,5,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}
type Principal_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,6,opt,name=header,proto3,oneof"`
}

func (*Principal_AndIds) isPrincipal_Identifier()         {}
func (*Principal_OrIds) isPrincipal_Identifier()          {}
func (*Principal_Any) isPrincipal_Identifier()            {}
func (*Principal_Authenticated_) isPrincipal_Identifier() {}
func (*Principal_SourceIp) isPrincipal_Identifier()       {}
func (*Principal_Header) isPrincipal_Identifier()         {}

func (m *Principal) GetIdentifier() isPrincipal_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Principal) GetAndIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_AndIds); ok {
		return x.AndIds
	}
	return nil
}

func (m *Principal) GetOrIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_OrIds); ok {
		return x.OrIds
	}
	return nil
}

func (m *Principal) GetAny() bool {
	if x, ok := m.GetIdentifier().(*Principal_Any); ok {
		return x.Any
	}
	return false
}

func (m *Principal) GetAuthenticated() *Principal_Authenticated {
	if x, ok := m.GetIdentifier().(*Principal_Authenticated_); ok {
		return x.Authenticated
	}
	return nil
}

func (m *Principal) GetSourceIp() *core.CidrRange {
	if x, ok := m.GetIdentifier().(*Principal_SourceIp); ok {
		return x.SourceIp
	}
	return nil
}

func (m *Principal) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Header); ok {
		return x.Header
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Principal) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Principal_OneofMarshaler, _Principal_OneofUnmarshaler, _Principal_OneofSizer, []interface{}{
		(*Principal_AndIds)(nil),
		(*Principal_OrIds)(nil),
		(*Principal_Any)(nil),
		(*Principal_Authenticated_)(nil),
		(*Principal_SourceIp)(nil),
		(*Principal_Header)(nil),
	}
}

func _Principal_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Principal)
	// identifier
	switch x := m.Identifier.(type) {
	case *Principal_AndIds:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AndIds); err != nil {
			return err
		}
	case *Principal_OrIds:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OrIds); err != nil {
			return err
		}
	case *Principal_Any:
		t := uint64(0)
		if x.Any {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Principal_Authenticated_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Authenticated); err != nil {
			return err
		}
	case *Principal_SourceIp:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SourceIp); err != nil {
			return err
		}
	case *Principal_Header:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Header); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Principal.Identifier has unexpected type %T", x)
	}
	return nil
}

func _Principal_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Principal)
	switch tag {
	case 1: // identifier.and_ids
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Principal_Set)
		err := b.DecodeMessage(msg)
		m.Identifier = &Principal_AndIds{msg}
		return true, err
	case 2: // identifier.or_ids
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Principal_Set)
		err := b.DecodeMessage(msg)
		m.Identifier = &Principal_OrIds{msg}
		return true, err
	case 3: // identifier.any
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Identifier = &Principal_Any{x != 0}
		return true, err
	case 4: // identifier.authenticated
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Principal_Authenticated)
		err := b.DecodeMessage(msg)
		m.Identifier = &Principal_Authenticated_{msg}
		return true, err
	case 5: // identifier.source_ip
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.CidrRange)
		err := b.DecodeMessage(msg)
		m.Identifier = &Principal_SourceIp{msg}
		return true, err
	case 6: // identifier.header
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(route.HeaderMatcher)
		err := b.DecodeMessage(msg)
		m.Identifier = &Principal_Header{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Principal_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Principal)
	// identifier
	switch x := m.Identifier.(type) {
	case *Principal_AndIds:
		s := proto.Size(x.AndIds)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Principal_OrIds:
		s := proto.Size(x.OrIds)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Principal_Any:
		n += 1 // tag and wire
		n += 1
	case *Principal_Authenticated_:
		s := proto.Size(x.Authenticated)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Principal_SourceIp:
		s := proto.Size(x.SourceIp)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Principal_Header:
		s := proto.Size(x.Header)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Used in the `and_ids` and `or_ids` fields in the `identifier` oneof. Depending on the context,
// each are applied with the associated behavior.
type Principal_Set struct {
	Ids                  []*Principal `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Principal_Set) Reset()         { *m = Principal_Set{} }
func (m *Principal_Set) String() string { return proto.CompactTextString(m) }
func (*Principal_Set) ProtoMessage()    {}
func (*Principal_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_713dfdae9a03bbb7, []int{3, 0}
}
func (m *Principal_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Set.Unmarshal(m, b)
}
func (m *Principal_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Set.Marshal(b, m, deterministic)
}
func (dst *Principal_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Set.Merge(dst, src)
}
func (m *Principal_Set) XXX_Size() int {
	return xxx_messageInfo_Principal_Set.Size(m)
}
func (m *Principal_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Set proto.InternalMessageInfo

func (m *Principal_Set) GetIds() []*Principal {
	if m != nil {
		return m.Ids
	}
	return nil
}

// Authentication attributes for a downstream.
type Principal_Authenticated struct {
	// The name of the principal. If set, the URI SAN is used from the certificate, otherwise the
	// subject field is used. If unset, it applies to any user that is authenticated.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Principal_Authenticated) Reset()         { *m = Principal_Authenticated{} }
func (m *Principal_Authenticated) String() string { return proto.CompactTextString(m) }
func (*Principal_Authenticated) ProtoMessage()    {}
func (*Principal_Authenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_rbac_713dfdae9a03bbb7, []int{3, 1}
}
func (m *Principal_Authenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Authenticated.Unmarshal(m, b)
}
func (m *Principal_Authenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Authenticated.Marshal(b, m, deterministic)
}
func (dst *Principal_Authenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Authenticated.Merge(dst, src)
}
func (m *Principal_Authenticated) XXX_Size() int {
	return xxx_messageInfo_Principal_Authenticated.Size(m)
}
func (m *Principal_Authenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Authenticated.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Authenticated proto.InternalMessageInfo

func (m *Principal_Authenticated) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*RBAC)(nil), "envoy.config.rbac.v2alpha.RBAC")
	proto.RegisterMapType((map[string]*Policy)(nil), "envoy.config.rbac.v2alpha.RBAC.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "envoy.config.rbac.v2alpha.Policy")
	proto.RegisterType((*Permission)(nil), "envoy.config.rbac.v2alpha.Permission")
	proto.RegisterType((*Permission_Set)(nil), "envoy.config.rbac.v2alpha.Permission.Set")
	proto.RegisterType((*Principal)(nil), "envoy.config.rbac.v2alpha.Principal")
	proto.RegisterType((*Principal_Set)(nil), "envoy.config.rbac.v2alpha.Principal.Set")
	proto.RegisterType((*Principal_Authenticated)(nil), "envoy.config.rbac.v2alpha.Principal.Authenticated")
	proto.RegisterEnum("envoy.config.rbac.v2alpha.RBAC_Action", RBAC_Action_name, RBAC_Action_value)
}

func init() {
	proto.RegisterFile("envoy/config/rbac/v2alpha/rbac.proto", fileDescriptor_rbac_713dfdae9a03bbb7)
}

var fileDescriptor_rbac_713dfdae9a03bbb7 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x4a,
	0x14, 0xc6, 0xe3, 0x3f, 0x71, 0x9d, 0x13, 0xa5, 0x37, 0x9a, 0xbb, 0xb8, 0xbe, 0x11, 0x85, 0x10,
	0x0a, 0x0a, 0x48, 0x38, 0x92, 0x59, 0x50, 0x51, 0x09, 0x29, 0x2e, 0x85, 0x44, 0x2a, 0xa5, 0x72,
	0x17, 0x88, 0x2e, 0xa8, 0xa6, 0xf6, 0xb4, 0x19, 0x48, 0x67, 0xac, 0xf1, 0x24, 0x52, 0x1e, 0x02,
	0x09, 0xf1, 0x22, 0x6c, 0x11, 0xab, 0x3e, 0x07, 0x6f, 0xd0, 0xa7, 0x28, 0x1a, 0x8f, 0x53, 0xe2,
	0x05, 0x6d, 0xda, 0x4d, 0x34, 0xf1, 0x7c, 0xdf, 0x6f, 0xce, 0x9c, 0xef, 0xd8, 0xb0, 0x4e, 0xd8,
	0x94, 0xcf, 0x7a, 0x31, 0x67, 0xc7, 0xf4, 0xa4, 0x27, 0x8e, 0x70, 0xdc, 0x9b, 0x06, 0x78, 0x9c,
	0x8e, 0x70, 0xfe, 0xc7, 0x4f, 0x05, 0x97, 0x1c, 0xfd, 0x9f, 0xab, 0x7c, 0xad, 0xf2, 0xf3, 0x8d,
	0x42, 0xd5, 0xfa, 0x6f, 0x8a, 0xc7, 0x34, 0xc1, 0x92, 0xf4, 0xe6, 0x0b, 0xed, 0x69, 0xdd, 0xd3,
	0x64, 0x9c, 0xd2, 0xde, 0x34, 0xe8, 0xc5, 0x5c, 0x90, 0x1e, 0x4e, 0x12, 0x41, 0xb2, 0xac, 0x10,
	0xdc, 0x2d, 0x09, 0x04, 0x9f, 0x48, 0xa2, 0x7f, 0xf5, 0x7e, 0xe7, 0xab, 0x09, 0x76, 0x14, 0xf6,
	0xb7, 0xd0, 0x4b, 0x70, 0x70, 0x2c, 0x29, 0x67, 0x9e, 0xd1, 0x36, 0xba, 0xab, 0xc1, 0x23, 0xff,
	0xaf, 0xe5, 0xf8, 0xca, 0xe0, 0xf7, 0x73, 0x75, 0x54, 0xb8, 0xd0, 0x10, 0xdc, 0x94, 0x8f, 0x69,
	0x4c, 0x49, 0xe6, 0x99, 0x6d, 0xab, 0x5b, 0x0f, 0x9e, 0x5e, 0x47, 0xd8, 0x2b, 0xf4, 0xdb, 0x4c,
	0x8a, 0x59, 0x74, 0x69, 0x6f, 0x7d, 0x84, 0x46, 0x69, 0x0b, 0x35, 0xc1, 0xfa, 0x4c, 0x66, 0x79,
	0x61, 0xb5, 0x48, 0x2d, 0xd1, 0x73, 0xa8, 0x4e, 0xf1, 0x78, 0x42, 0x3c, 0xb3, 0x6d, 0x74, 0xeb,
	0xc1, 0xfd, 0x2b, 0x8e, 0xca, 0x51, 0xb3, 0x48, 0xeb, 0x5f, 0x98, 0x1b, 0x46, 0x67, 0x0d, 0x1c,
	0x5d, 0x3c, 0xaa, 0x41, 0xb5, 0xbf, 0xb3, 0xf3, 0xee, 0x7d, 0xb3, 0x82, 0x5c, 0xb0, 0x5f, 0x6d,
	0xef, 0x7e, 0x68, 0x1a, 0x9d, 0xef, 0x06, 0x38, 0xda, 0x84, 0xf6, 0xa1, 0x9e, 0x12, 0x71, 0x4a,
	0xb3, 0x8c, 0x72, 0x96, 0x79, 0x46, 0x7e, 0xaf, 0x87, 0x57, 0x1d, 0x76, 0xa9, 0x0e, 0xe1, 0xe7,
	0xf9, 0x99, 0x55, 0xfd, 0x66, 0x98, 0xae, 0x11, 0x2d, 0x52, 0xd0, 0x1e, 0x40, 0x2a, 0x28, 0x8b,
	0x69, 0x8a, 0xc7, 0xf3, 0x5e, 0xad, 0x5f, 0xc5, 0x9c, 0x8b, 0x4b, 0xc8, 0x05, 0x46, 0xe7, 0x97,
	0x05, 0xf0, 0xe7, 0x64, 0x34, 0x80, 0x1a, 0x66, 0xc9, 0xa1, 0x98, 0x8c, 0x49, 0x96, 0x37, 0xad,
	0x1e, 0x3c, 0x5e, 0xaa, 0x66, 0x7f, 0x9f, 0xc8, 0x41, 0x25, 0x72, 0x31, 0x4b, 0x22, 0x65, 0x46,
	0xaf, 0xc1, 0xe5, 0xa2, 0x00, 0x99, 0x37, 0x07, 0xad, 0x70, 0xa1, 0x39, 0x6b, 0x60, 0x61, 0x36,
	0xf3, 0xac, 0xb6, 0xd1, 0x75, 0xc3, 0x9a, 0xba, 0x85, 0xfd, 0xc9, 0x74, 0x8d, 0x41, 0x25, 0x52,
	0xcf, 0xd1, 0x26, 0x38, 0x23, 0x82, 0x13, 0x22, 0x3c, 0xbb, 0x14, 0x27, 0x4e, 0xa9, 0x3f, 0x0d,
	0x7c, 0x3d, 0xaf, 0x83, 0x5c, 0xf1, 0x16, 0xcb, 0x78, 0x44, 0xc4, 0xa0, 0x12, 0x15, 0x16, 0xb4,
	0x0d, 0xab, 0x09, 0xc9, 0x24, 0x65, 0x58, 0x45, 0x7a, 0x48, 0x53, 0xaf, 0x9a, 0x43, 0xee, 0x94,
	0x21, 0xea, 0xdd, 0xf0, 0xb7, 0x68, 0x22, 0x22, 0xcc, 0x4e, 0xc8, 0xa0, 0x12, 0x35, 0x16, 0x5c,
	0xc3, 0x14, 0x6d, 0x40, 0x73, 0x11, 0x93, 0x72, 0x21, 0x3d, 0xa7, 0x6d, 0x74, 0x1b, 0x61, 0x5d,
	0xd5, 0xeb, 0x3c, 0xb1, 0xbd, 0x8b, 0x0b, 0x6b, 0x50, 0x89, 0xfe, 0x59, 0x90, 0xed, 0x71, 0x21,
	0x5b, 0xbb, 0x60, 0xed, 0x13, 0x89, 0xde, 0x40, 0x75, 0xde, 0xf1, 0x5b, 0x4e, 0x89, 0xf6, 0x87,
	0x0d, 0xb0, 0xd5, 0x02, 0x55, 0x7f, 0x9c, 0x9f, 0x59, 0x46, 0xe7, 0x8b, 0x0d, 0xb5, 0xcb, 0x11,
	0x40, 0x5b, 0xb0, 0xa2, 0xb2, 0xa5, 0xc9, 0x3c, 0xd9, 0xee, 0x32, 0x93, 0x53, 0xe4, 0xe1, 0x60,
	0x96, 0x0c, 0x93, 0x0c, 0xf5, 0xc1, 0xe1, 0x22, 0x67, 0x98, 0x37, 0x66, 0x54, 0xb9, 0x50, 0x88,
	0x6b, 0x12, 0x3d, 0x80, 0x06, 0x9e, 0xc8, 0x11, 0x61, 0x92, 0xc6, 0x58, 0x92, 0xa4, 0x08, 0x36,
	0x58, 0xea, 0xa0, 0xfe, 0xa2, 0x53, 0x25, 0x55, 0x42, 0xa1, 0x4d, 0xa8, 0x65, 0x7c, 0x22, 0x62,
	0xb2, 0x7c, 0xd6, 0xae, 0x36, 0x0c, 0xd3, 0x85, 0x51, 0x73, 0x6e, 0x3c, 0x6a, 0xad, 0xa1, 0x4e,
	0x3a, 0x04, 0x4b, 0xf7, 0xff, 0x76, 0x6f, 0xae, 0x32, 0xb7, 0x1e, 0x40, 0xa3, 0x74, 0x4d, 0x84,
	0xc0, 0x66, 0xf8, 0x94, 0x14, 0x1f, 0xb9, 0x7c, 0x1d, 0xfe, 0x0b, 0x40, 0x13, 0x25, 0x39, 0xa6,
	0x44, 0x14, 0xf3, 0x10, 0xd6, 0x0e, 0x56, 0x0a, 0xfe, 0x91, 0x93, 0x7f, 0xc3, 0x9f, 0xfd, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x99, 0x71, 0x76, 0xe3, 0x60, 0x06, 0x00, 0x00,
}
