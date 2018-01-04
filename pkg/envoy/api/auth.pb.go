// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/auth.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Should we do white-list or black-list style access control.
type AuthAction_ActionType int32

const (
	// Request matches all rules are allowed, otherwise denied.
	AuthAction_ALLOW AuthAction_ActionType = 0
	// Request matches all rules or missing required auth fields are denied,
	// otherwise allowed.
	AuthAction_DENY AuthAction_ActionType = 1
)

var AuthAction_ActionType_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}
var AuthAction_ActionType_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x AuthAction_ActionType) String() string {
	return proto.EnumName(AuthAction_ActionType_name, int32(x))
}
func (AuthAction_ActionType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type AuthAction struct {
	ActionType AuthAction_ActionType `protobuf:"varint,1,opt,name=action_type,json=actionType,enum=envoy.api.v2.AuthAction_ActionType" json:"action_type,omitempty"`
	// List of rules
	Rule []*AuthAction_Rule `protobuf:"bytes,2,rep,name=rule" json:"rule,omitempty"`
}

func (m *AuthAction) Reset()                    { *m = AuthAction{} }
func (m *AuthAction) String() string            { return proto.CompactTextString(m) }
func (*AuthAction) ProtoMessage()               {}
func (*AuthAction) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *AuthAction) GetActionType() AuthAction_ActionType {
	if m != nil {
		return m.ActionType
	}
	return AuthAction_ALLOW
}

func (m *AuthAction) GetRule() []*AuthAction_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

// Logic AND that requires all rules match.
type AuthAction_AndRule struct {
	Rules []*AuthAction_Rule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *AuthAction_AndRule) Reset()                    { *m = AuthAction_AndRule{} }
func (m *AuthAction_AndRule) String() string            { return proto.CompactTextString(m) }
func (*AuthAction_AndRule) ProtoMessage()               {}
func (*AuthAction_AndRule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *AuthAction_AndRule) GetRules() []*AuthAction_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Logic OR that requires at least one rule matches.
type AuthAction_OrRule struct {
	Rules []*AuthAction_Rule `protobuf:"bytes,1,rep,name=rules" json:"rules,omitempty"`
}

func (m *AuthAction_OrRule) Reset()                    { *m = AuthAction_OrRule{} }
func (m *AuthAction_OrRule) String() string            { return proto.CompactTextString(m) }
func (*AuthAction_OrRule) ProtoMessage()               {}
func (*AuthAction_OrRule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

func (m *AuthAction_OrRule) GetRules() []*AuthAction_Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Check peer identity using X.509 certificate.
type AuthAction_X509Rule struct {
	// How to validate peer certificates.
	ValidationContext *CertificateValidationContext `protobuf:"bytes,3,opt,name=validation_context,json=validationContext" json:"validation_context,omitempty"`
}

func (m *AuthAction_X509Rule) Reset()                    { *m = AuthAction_X509Rule{} }
func (m *AuthAction_X509Rule) String() string            { return proto.CompactTextString(m) }
func (*AuthAction_X509Rule) ProtoMessage()               {}
func (*AuthAction_X509Rule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 2} }

func (m *AuthAction_X509Rule) GetValidationContext() *CertificateValidationContext {
	if m != nil {
		return m.ValidationContext
	}
	return nil
}

// Element type of AndRule/OrRule, it chooses among different type of rule.
type AuthAction_Rule struct {
	// Types that are valid to be assigned to RuleSpecifier:
	//	*AuthAction_Rule_AndRule
	//	*AuthAction_Rule_OrRule
	//	*AuthAction_Rule_X509Rule
	RuleSpecifier isAuthAction_Rule_RuleSpecifier `protobuf_oneof:"rule_specifier"`
}

func (m *AuthAction_Rule) Reset()                    { *m = AuthAction_Rule{} }
func (m *AuthAction_Rule) String() string            { return proto.CompactTextString(m) }
func (*AuthAction_Rule) ProtoMessage()               {}
func (*AuthAction_Rule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 3} }

type isAuthAction_Rule_RuleSpecifier interface {
	isAuthAction_Rule_RuleSpecifier()
}

type AuthAction_Rule_AndRule struct {
	AndRule *AuthAction_AndRule `protobuf:"bytes,1,opt,name=and_rule,json=andRule,oneof"`
}
type AuthAction_Rule_OrRule struct {
	OrRule *AuthAction_OrRule `protobuf:"bytes,2,opt,name=or_rule,json=orRule,oneof"`
}
type AuthAction_Rule_X509Rule struct {
	X509Rule *AuthAction_X509Rule `protobuf:"bytes,3,opt,name=x509_rule,json=x509Rule,oneof"`
}

func (*AuthAction_Rule_AndRule) isAuthAction_Rule_RuleSpecifier()  {}
func (*AuthAction_Rule_OrRule) isAuthAction_Rule_RuleSpecifier()   {}
func (*AuthAction_Rule_X509Rule) isAuthAction_Rule_RuleSpecifier() {}

func (m *AuthAction_Rule) GetRuleSpecifier() isAuthAction_Rule_RuleSpecifier {
	if m != nil {
		return m.RuleSpecifier
	}
	return nil
}

func (m *AuthAction_Rule) GetAndRule() *AuthAction_AndRule {
	if x, ok := m.GetRuleSpecifier().(*AuthAction_Rule_AndRule); ok {
		return x.AndRule
	}
	return nil
}

func (m *AuthAction_Rule) GetOrRule() *AuthAction_OrRule {
	if x, ok := m.GetRuleSpecifier().(*AuthAction_Rule_OrRule); ok {
		return x.OrRule
	}
	return nil
}

func (m *AuthAction_Rule) GetX509Rule() *AuthAction_X509Rule {
	if x, ok := m.GetRuleSpecifier().(*AuthAction_Rule_X509Rule); ok {
		return x.X509Rule
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AuthAction_Rule) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AuthAction_Rule_OneofMarshaler, _AuthAction_Rule_OneofUnmarshaler, _AuthAction_Rule_OneofSizer, []interface{}{
		(*AuthAction_Rule_AndRule)(nil),
		(*AuthAction_Rule_OrRule)(nil),
		(*AuthAction_Rule_X509Rule)(nil),
	}
}

func _AuthAction_Rule_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AuthAction_Rule)
	// rule_specifier
	switch x := m.RuleSpecifier.(type) {
	case *AuthAction_Rule_AndRule:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AndRule); err != nil {
			return err
		}
	case *AuthAction_Rule_OrRule:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OrRule); err != nil {
			return err
		}
	case *AuthAction_Rule_X509Rule:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.X509Rule); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AuthAction_Rule.RuleSpecifier has unexpected type %T", x)
	}
	return nil
}

func _AuthAction_Rule_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AuthAction_Rule)
	switch tag {
	case 1: // rule_specifier.and_rule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AuthAction_AndRule)
		err := b.DecodeMessage(msg)
		m.RuleSpecifier = &AuthAction_Rule_AndRule{msg}
		return true, err
	case 2: // rule_specifier.or_rule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AuthAction_OrRule)
		err := b.DecodeMessage(msg)
		m.RuleSpecifier = &AuthAction_Rule_OrRule{msg}
		return true, err
	case 3: // rule_specifier.x509_rule
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AuthAction_X509Rule)
		err := b.DecodeMessage(msg)
		m.RuleSpecifier = &AuthAction_Rule_X509Rule{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AuthAction_Rule_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AuthAction_Rule)
	// rule_specifier
	switch x := m.RuleSpecifier.(type) {
	case *AuthAction_Rule_AndRule:
		s := proto.Size(x.AndRule)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthAction_Rule_OrRule:
		s := proto.Size(x.OrRule)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AuthAction_Rule_X509Rule:
		s := proto.Size(x.X509Rule)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AuthAction)(nil), "envoy.api.v2.AuthAction")
	proto.RegisterType((*AuthAction_AndRule)(nil), "envoy.api.v2.AuthAction.AndRule")
	proto.RegisterType((*AuthAction_OrRule)(nil), "envoy.api.v2.AuthAction.OrRule")
	proto.RegisterType((*AuthAction_X509Rule)(nil), "envoy.api.v2.AuthAction.X509Rule")
	proto.RegisterType((*AuthAction_Rule)(nil), "envoy.api.v2.AuthAction.Rule")
	proto.RegisterEnum("envoy.api.v2.AuthAction_ActionType", AuthAction_ActionType_name, AuthAction_ActionType_value)
}

func init() { proto.RegisterFile("api/auth.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6b, 0xea, 0x40,
	0x10, 0xc7, 0x93, 0x67, 0xd4, 0x38, 0xbe, 0x27, 0xbe, 0x9c, 0x42, 0xa0, 0x34, 0xda, 0x8b, 0xf4,
	0x90, 0xda, 0x88, 0x07, 0x0b, 0x96, 0xa6, 0x5a, 0xf0, 0x20, 0x15, 0x96, 0xd2, 0xd6, 0x53, 0xd8,
	0x26, 0x2b, 0x2e, 0x48, 0x36, 0x24, 0x9b, 0xa0, 0x1f, 0xb4, 0x5f, 0xa1, 0x9f, 0xa3, 0x64, 0xd7,
	0x6a, 0x5b, 0x08, 0x94, 0x9e, 0x32, 0xcc, 0xcc, 0xef, 0x1f, 0xf6, 0xc7, 0x40, 0x0b, 0xc7, 0xf4,
	0x02, 0x67, 0x7c, 0xed, 0xc4, 0x09, 0xe3, 0xcc, 0xf8, 0x4b, 0xa2, 0x9c, 0xed, 0x1c, 0x1c, 0x53,
	0x27, 0x77, 0xad, 0x7f, 0xc5, 0x34, 0x0d, 0x53, 0x39, 0xec, 0xbe, 0x69, 0x00, 0x5e, 0xc6, 0xd7,
	0x5e, 0xc0, 0x29, 0x8b, 0x8c, 0x29, 0x34, 0xb1, 0xa8, 0x7c, 0xbe, 0x8b, 0x89, 0xa9, 0xda, 0x6a,
	0xaf, 0xe5, 0x9e, 0x39, 0x9f, 0x13, 0x9c, 0xe3, 0xba, 0x23, 0x3f, 0x0f, 0xbb, 0x98, 0x20, 0xc0,
	0x87, 0xda, 0xb8, 0x04, 0x2d, 0xc9, 0x36, 0xc4, 0xfc, 0x63, 0x57, 0x7a, 0x4d, 0xf7, 0xa4, 0x14,
	0x47, 0xd9, 0x86, 0x20, 0xb1, 0x6a, 0x5d, 0x43, 0xdd, 0x8b, 0xc2, 0xa2, 0x61, 0x0c, 0xa0, 0x5a,
	0xb4, 0x52, 0x53, 0xfd, 0x09, 0x2e, 0x77, 0xad, 0x31, 0xd4, 0x16, 0xc9, 0xef, 0x71, 0x02, 0xfa,
	0xf3, 0xb0, 0x3f, 0x12, 0x01, 0x4b, 0x30, 0x72, 0xbc, 0xa1, 0x21, 0x16, 0x1e, 0x02, 0x16, 0x71,
	0xb2, 0xe5, 0x66, 0xc5, 0x56, 0x7b, 0x4d, 0xf7, 0xfc, 0x6b, 0xda, 0x84, 0x24, 0x9c, 0xae, 0x68,
	0x80, 0x39, 0x79, 0x3c, 0x20, 0x13, 0x49, 0xa0, 0xff, 0xf9, 0xf7, 0x96, 0xf5, 0xaa, 0x82, 0x26,
	0xfe, 0x31, 0x06, 0x1d, 0x47, 0xa1, 0x2f, 0x2c, 0xa9, 0x22, 0xd9, 0x2e, 0x97, 0x2c, 0xbd, 0xcc,
	0x14, 0x54, 0xc7, 0x7b, 0x45, 0x57, 0x50, 0x67, 0x89, 0xbf, 0x77, 0x5c, 0xd0, 0xa7, 0xa5, 0xb4,
	0xb4, 0x32, 0x53, 0x50, 0x8d, 0x49, 0x3f, 0x37, 0xd0, 0xd8, 0x0e, 0xfb, 0x23, 0x49, 0xcb, 0x57,
	0x75, 0x4a, 0xe9, 0x0f, 0x29, 0x33, 0x05, 0xe9, 0xdb, 0x7d, 0x7d, 0xdb, 0x86, 0x56, 0x01, 0xfb,
	0x69, 0x4c, 0x02, 0xba, 0xa2, 0x24, 0xe9, 0x76, 0x00, 0x8e, 0xa7, 0x60, 0x34, 0xa0, 0xea, 0xcd,
	0xe7, 0x8b, 0xa7, 0xb6, 0x62, 0xe8, 0xa0, 0x4d, 0xef, 0xee, 0x97, 0x6d, 0xf5, 0xa5, 0x26, 0xee,
	0x6d, 0xf0, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x86, 0xa3, 0x19, 0x75, 0x9e, 0x02, 0x00, 0x00,
}
