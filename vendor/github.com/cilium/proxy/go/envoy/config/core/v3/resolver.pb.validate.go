// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/resolver.proto

package corev3

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

// Validate checks the field values on DnsResolverOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DnsResolverOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsResolverOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DnsResolverOptionsMultiError, or nil if none found.
func (m *DnsResolverOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsResolverOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UseTcpForDnsLookups

	// no validation rules for NoDefaultSearchDomain

	if len(errors) > 0 {
		return DnsResolverOptionsMultiError(errors)
	}
	return nil
}

// DnsResolverOptionsMultiError is an error wrapping multiple validation errors
// returned by DnsResolverOptions.ValidateAll() if the designated constraints
// aren't met.
type DnsResolverOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsResolverOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsResolverOptionsMultiError) AllErrors() []error { return m }

// DnsResolverOptionsValidationError is the validation error returned by
// DnsResolverOptions.Validate if the designated constraints aren't met.
type DnsResolverOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsResolverOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsResolverOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsResolverOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsResolverOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsResolverOptionsValidationError) ErrorName() string {
	return "DnsResolverOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e DnsResolverOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsResolverOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsResolverOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsResolverOptionsValidationError{}

// Validate checks the field values on DnsResolutionConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DnsResolutionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DnsResolutionConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DnsResolutionConfigMultiError, or nil if none found.
func (m *DnsResolutionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *DnsResolutionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetResolvers()) < 1 {
		err := DnsResolutionConfigValidationError{
			field:  "Resolvers",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetResolvers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DnsResolutionConfigValidationError{
						field:  fmt.Sprintf("Resolvers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DnsResolutionConfigValidationError{
						field:  fmt.Sprintf("Resolvers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsResolutionConfigValidationError{
					field:  fmt.Sprintf("Resolvers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetDnsResolverOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DnsResolutionConfigValidationError{
					field:  "DnsResolverOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DnsResolutionConfigValidationError{
					field:  "DnsResolverOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDnsResolverOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsResolutionConfigValidationError{
				field:  "DnsResolverOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DnsResolutionConfigMultiError(errors)
	}
	return nil
}

// DnsResolutionConfigMultiError is an error wrapping multiple validation
// errors returned by DnsResolutionConfig.ValidateAll() if the designated
// constraints aren't met.
type DnsResolutionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DnsResolutionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DnsResolutionConfigMultiError) AllErrors() []error { return m }

// DnsResolutionConfigValidationError is the validation error returned by
// DnsResolutionConfig.Validate if the designated constraints aren't met.
type DnsResolutionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsResolutionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsResolutionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsResolutionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsResolutionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsResolutionConfigValidationError) ErrorName() string {
	return "DnsResolutionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e DnsResolutionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsResolutionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsResolutionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsResolutionConfigValidationError{}
