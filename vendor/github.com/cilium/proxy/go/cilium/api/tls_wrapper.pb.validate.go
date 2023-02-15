// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cilium/api/tls_wrapper.proto

package cilium

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

// Validate checks the field values on UpstreamTlsWrapperContext with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpstreamTlsWrapperContext) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpstreamTlsWrapperContext with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpstreamTlsWrapperContextMultiError, or nil if none found.
func (m *UpstreamTlsWrapperContext) ValidateAll() error {
	return m.validate(true)
}

func (m *UpstreamTlsWrapperContext) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpstreamTlsWrapperContextMultiError(errors)
	}
	return nil
}

// UpstreamTlsWrapperContextMultiError is an error wrapping multiple validation
// errors returned by UpstreamTlsWrapperContext.ValidateAll() if the
// designated constraints aren't met.
type UpstreamTlsWrapperContextMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpstreamTlsWrapperContextMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpstreamTlsWrapperContextMultiError) AllErrors() []error { return m }

// UpstreamTlsWrapperContextValidationError is the validation error returned by
// UpstreamTlsWrapperContext.Validate if the designated constraints aren't met.
type UpstreamTlsWrapperContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamTlsWrapperContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamTlsWrapperContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamTlsWrapperContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamTlsWrapperContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamTlsWrapperContextValidationError) ErrorName() string {
	return "UpstreamTlsWrapperContextValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamTlsWrapperContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamTlsWrapperContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamTlsWrapperContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamTlsWrapperContextValidationError{}

// Validate checks the field values on DownstreamTlsWrapperContext with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownstreamTlsWrapperContext) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownstreamTlsWrapperContext with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownstreamTlsWrapperContextMultiError, or nil if none found.
func (m *DownstreamTlsWrapperContext) ValidateAll() error {
	return m.validate(true)
}

func (m *DownstreamTlsWrapperContext) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DownstreamTlsWrapperContextMultiError(errors)
	}
	return nil
}

// DownstreamTlsWrapperContextMultiError is an error wrapping multiple
// validation errors returned by DownstreamTlsWrapperContext.ValidateAll() if
// the designated constraints aren't met.
type DownstreamTlsWrapperContextMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownstreamTlsWrapperContextMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownstreamTlsWrapperContextMultiError) AllErrors() []error { return m }

// DownstreamTlsWrapperContextValidationError is the validation error returned
// by DownstreamTlsWrapperContext.Validate if the designated constraints
// aren't met.
type DownstreamTlsWrapperContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownstreamTlsWrapperContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownstreamTlsWrapperContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownstreamTlsWrapperContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownstreamTlsWrapperContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownstreamTlsWrapperContextValidationError) ErrorName() string {
	return "DownstreamTlsWrapperContextValidationError"
}

// Error satisfies the builtin error interface
func (e DownstreamTlsWrapperContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownstreamTlsWrapperContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownstreamTlsWrapperContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownstreamTlsWrapperContextValidationError{}
