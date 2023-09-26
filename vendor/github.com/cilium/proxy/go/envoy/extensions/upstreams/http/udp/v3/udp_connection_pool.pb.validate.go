// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/upstreams/http/udp/v3/udp_connection_pool.proto

package udpv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on UdpConnectionPoolProto with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UdpConnectionPoolProto) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UdpConnectionPoolProto with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UdpConnectionPoolProtoMultiError, or nil if none found.
func (m *UdpConnectionPoolProto) ValidateAll() error {
	return m.validate(true)
}

func (m *UdpConnectionPoolProto) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UdpConnectionPoolProtoMultiError(errors)
	}

	return nil
}

// UdpConnectionPoolProtoMultiError is an error wrapping multiple validation
// errors returned by UdpConnectionPoolProto.ValidateAll() if the designated
// constraints aren't met.
type UdpConnectionPoolProtoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UdpConnectionPoolProtoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UdpConnectionPoolProtoMultiError) AllErrors() []error { return m }

// UdpConnectionPoolProtoValidationError is the validation error returned by
// UdpConnectionPoolProto.Validate if the designated constraints aren't met.
type UdpConnectionPoolProtoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpConnectionPoolProtoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpConnectionPoolProtoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpConnectionPoolProtoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpConnectionPoolProtoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpConnectionPoolProtoValidationError) ErrorName() string {
	return "UdpConnectionPoolProtoValidationError"
}

// Error satisfies the builtin error interface
func (e UdpConnectionPoolProtoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpConnectionPoolProto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpConnectionPoolProtoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpConnectionPoolProtoValidationError{}
