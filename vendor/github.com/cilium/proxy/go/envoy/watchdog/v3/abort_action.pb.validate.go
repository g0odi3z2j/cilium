// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/watchdog/v3/abort_action.proto

package watchdogv3

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

// Validate checks the field values on AbortActionConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AbortActionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AbortActionConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AbortActionConfigMultiError, or nil if none found.
func (m *AbortActionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *AbortActionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetWaitDuration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AbortActionConfigValidationError{
					field:  "WaitDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AbortActionConfigValidationError{
					field:  "WaitDuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWaitDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AbortActionConfigValidationError{
				field:  "WaitDuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AbortActionConfigMultiError(errors)
	}

	return nil
}

// AbortActionConfigMultiError is an error wrapping multiple validation errors
// returned by AbortActionConfig.ValidateAll() if the designated constraints
// aren't met.
type AbortActionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AbortActionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AbortActionConfigMultiError) AllErrors() []error { return m }

// AbortActionConfigValidationError is the validation error returned by
// AbortActionConfig.Validate if the designated constraints aren't met.
type AbortActionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AbortActionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AbortActionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AbortActionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AbortActionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AbortActionConfigValidationError) ErrorName() string {
	return "AbortActionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e AbortActionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAbortActionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AbortActionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AbortActionConfigValidationError{}
