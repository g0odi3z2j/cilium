// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/listener/tls_inspector/v3/tls_inspector.proto

package tls_inspectorv3

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

// Validate checks the field values on TlsInspector with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TlsInspector) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TlsInspector with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TlsInspectorMultiError, or
// nil if none found.
func (m *TlsInspector) ValidateAll() error {
	return m.validate(true)
}

func (m *TlsInspector) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEnableJa3Fingerprinting()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TlsInspectorValidationError{
					field:  "EnableJa3Fingerprinting",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TlsInspectorValidationError{
					field:  "EnableJa3Fingerprinting",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnableJa3Fingerprinting()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TlsInspectorValidationError{
				field:  "EnableJa3Fingerprinting",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TlsInspectorMultiError(errors)
	}
	return nil
}

// TlsInspectorMultiError is an error wrapping multiple validation errors
// returned by TlsInspector.ValidateAll() if the designated constraints aren't met.
type TlsInspectorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TlsInspectorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TlsInspectorMultiError) AllErrors() []error { return m }

// TlsInspectorValidationError is the validation error returned by
// TlsInspector.Validate if the designated constraints aren't met.
type TlsInspectorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TlsInspectorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TlsInspectorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TlsInspectorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TlsInspectorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TlsInspectorValidationError) ErrorName() string { return "TlsInspectorValidationError" }

// Error satisfies the builtin error interface
func (e TlsInspectorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTlsInspector.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TlsInspectorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TlsInspectorValidationError{}
