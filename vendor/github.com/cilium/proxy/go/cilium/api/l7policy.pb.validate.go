// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cilium/api/l7policy.proto

package cilium

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _l_7_policy_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on L7Policy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *L7Policy) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessLogPath

	// no validation rules for PolicyName

	// no validation rules for Denied_403Body

	if v, ok := interface{}(m.GetIsIngress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return L7PolicyValidationError{
				field:  "IsIngress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// L7PolicyValidationError is the validation error returned by
// L7Policy.Validate if the designated constraints aren't met.
type L7PolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e L7PolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e L7PolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e L7PolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e L7PolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e L7PolicyValidationError) ErrorName() string { return "L7PolicyValidationError" }

// Error satisfies the builtin error interface
func (e L7PolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sL7Policy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = L7PolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = L7PolicyValidationError{}
