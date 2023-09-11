// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v3/skywalking.proto

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

// Validate checks the field values on SkyWalkingConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SkyWalkingConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SkyWalkingConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SkyWalkingConfigMultiError, or nil if none found.
func (m *SkyWalkingConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *SkyWalkingConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetGrpcService() == nil {
		err := SkyWalkingConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetGrpcService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SkyWalkingConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SkyWalkingConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SkyWalkingConfigValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetClientConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SkyWalkingConfigValidationError{
					field:  "ClientConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SkyWalkingConfigValidationError{
					field:  "ClientConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClientConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SkyWalkingConfigValidationError{
				field:  "ClientConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SkyWalkingConfigMultiError(errors)
	}

	return nil
}

// SkyWalkingConfigMultiError is an error wrapping multiple validation errors
// returned by SkyWalkingConfig.ValidateAll() if the designated constraints
// aren't met.
type SkyWalkingConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SkyWalkingConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SkyWalkingConfigMultiError) AllErrors() []error { return m }

// SkyWalkingConfigValidationError is the validation error returned by
// SkyWalkingConfig.Validate if the designated constraints aren't met.
type SkyWalkingConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SkyWalkingConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SkyWalkingConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SkyWalkingConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SkyWalkingConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SkyWalkingConfigValidationError) ErrorName() string { return "SkyWalkingConfigValidationError" }

// Error satisfies the builtin error interface
func (e SkyWalkingConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSkyWalkingConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SkyWalkingConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SkyWalkingConfigValidationError{}

// Validate checks the field values on ClientConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClientConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClientConfigMultiError, or
// nil if none found.
func (m *ClientConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ServiceName

	// no validation rules for InstanceName

	if all {
		switch v := interface{}(m.GetMaxCacheSize()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClientConfigValidationError{
					field:  "MaxCacheSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClientConfigValidationError{
					field:  "MaxCacheSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMaxCacheSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientConfigValidationError{
				field:  "MaxCacheSize",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.BackendTokenSpecifier.(type) {
	case *ClientConfig_BackendToken:
		if v == nil {
			err := ClientConfigValidationError{
				field:  "BackendTokenSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for BackendToken
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ClientConfigMultiError(errors)
	}

	return nil
}

// ClientConfigMultiError is an error wrapping multiple validation errors
// returned by ClientConfig.ValidateAll() if the designated constraints aren't met.
type ClientConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientConfigMultiError) AllErrors() []error { return m }

// ClientConfigValidationError is the validation error returned by
// ClientConfig.Validate if the designated constraints aren't met.
type ClientConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientConfigValidationError) ErrorName() string { return "ClientConfigValidationError" }

// Error satisfies the builtin error interface
func (e ClientConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientConfigValidationError{}
