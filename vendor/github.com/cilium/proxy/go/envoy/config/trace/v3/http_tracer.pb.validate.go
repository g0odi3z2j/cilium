// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v3/http_tracer.proto

package tracev3

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

// Validate checks the field values on Tracing with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tracing) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tracing with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TracingMultiError, or nil if none found.
func (m *Tracing) ValidateAll() error {
	return m.validate(true)
}

func (m *Tracing) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHttp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TracingValidationError{
					field:  "Http",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TracingValidationError{
					field:  "Http",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TracingValidationError{
				field:  "Http",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TracingMultiError(errors)
	}
	return nil
}

// TracingMultiError is an error wrapping multiple validation errors returned
// by Tracing.ValidateAll() if the designated constraints aren't met.
type TracingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TracingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TracingMultiError) AllErrors() []error { return m }

// TracingValidationError is the validation error returned by Tracing.Validate
// if the designated constraints aren't met.
type TracingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TracingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TracingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TracingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TracingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TracingValidationError) ErrorName() string { return "TracingValidationError" }

// Error satisfies the builtin error interface
func (e TracingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TracingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TracingValidationError{}

// Validate checks the field values on Tracing_Http with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Tracing_Http) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tracing_Http with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Tracing_HttpMultiError, or
// nil if none found.
func (m *Tracing_Http) ValidateAll() error {
	return m.validate(true)
}

func (m *Tracing_Http) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := Tracing_HttpValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	switch m.ConfigType.(type) {

	case *Tracing_Http_TypedConfig:

		if all {
			switch v := interface{}(m.GetTypedConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Tracing_HttpValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Tracing_HttpValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Tracing_HttpValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return Tracing_HttpMultiError(errors)
	}
	return nil
}

// Tracing_HttpMultiError is an error wrapping multiple validation errors
// returned by Tracing_Http.ValidateAll() if the designated constraints aren't met.
type Tracing_HttpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Tracing_HttpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Tracing_HttpMultiError) AllErrors() []error { return m }

// Tracing_HttpValidationError is the validation error returned by
// Tracing_Http.Validate if the designated constraints aren't met.
type Tracing_HttpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Tracing_HttpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Tracing_HttpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Tracing_HttpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Tracing_HttpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Tracing_HttpValidationError) ErrorName() string { return "Tracing_HttpValidationError" }

// Error satisfies the builtin error interface
func (e Tracing_HttpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing_Http.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Tracing_HttpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Tracing_HttpValidationError{}
