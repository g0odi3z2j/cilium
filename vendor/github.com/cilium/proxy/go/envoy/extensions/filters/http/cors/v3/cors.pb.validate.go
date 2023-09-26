// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/cors/v3/cors.proto

package corsv3

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

// Validate checks the field values on Cors with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Cors) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cors with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CorsMultiError, or nil if none found.
func (m *Cors) ValidateAll() error {
	return m.validate(true)
}

func (m *Cors) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CorsMultiError(errors)
	}

	return nil
}

// CorsMultiError is an error wrapping multiple validation errors returned by
// Cors.ValidateAll() if the designated constraints aren't met.
type CorsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CorsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CorsMultiError) AllErrors() []error { return m }

// CorsValidationError is the validation error returned by Cors.Validate if the
// designated constraints aren't met.
type CorsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CorsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CorsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CorsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CorsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CorsValidationError) ErrorName() string { return "CorsValidationError" }

// Error satisfies the builtin error interface
func (e CorsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCors.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CorsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CorsValidationError{}

// Validate checks the field values on CorsPolicy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CorsPolicy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CorsPolicy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CorsPolicyMultiError, or
// nil if none found.
func (m *CorsPolicy) ValidateAll() error {
	return m.validate(true)
}

func (m *CorsPolicy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAllowOriginStringMatch() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CorsPolicyValidationError{
						field:  fmt.Sprintf("AllowOriginStringMatch[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CorsPolicyValidationError{
						field:  fmt.Sprintf("AllowOriginStringMatch[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CorsPolicyValidationError{
					field:  fmt.Sprintf("AllowOriginStringMatch[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for AllowMethods

	// no validation rules for AllowHeaders

	// no validation rules for ExposeHeaders

	// no validation rules for MaxAge

	if all {
		switch v := interface{}(m.GetAllowCredentials()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CorsPolicyValidationError{
					field:  "AllowCredentials",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CorsPolicyValidationError{
					field:  "AllowCredentials",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAllowCredentials()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CorsPolicyValidationError{
				field:  "AllowCredentials",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFilterEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CorsPolicyValidationError{
					field:  "FilterEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CorsPolicyValidationError{
					field:  "FilterEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilterEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CorsPolicyValidationError{
				field:  "FilterEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetShadowEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CorsPolicyValidationError{
					field:  "ShadowEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CorsPolicyValidationError{
					field:  "ShadowEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShadowEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CorsPolicyValidationError{
				field:  "ShadowEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAllowPrivateNetworkAccess()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CorsPolicyValidationError{
					field:  "AllowPrivateNetworkAccess",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CorsPolicyValidationError{
					field:  "AllowPrivateNetworkAccess",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAllowPrivateNetworkAccess()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CorsPolicyValidationError{
				field:  "AllowPrivateNetworkAccess",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CorsPolicyMultiError(errors)
	}

	return nil
}

// CorsPolicyMultiError is an error wrapping multiple validation errors
// returned by CorsPolicy.ValidateAll() if the designated constraints aren't met.
type CorsPolicyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CorsPolicyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CorsPolicyMultiError) AllErrors() []error { return m }

// CorsPolicyValidationError is the validation error returned by
// CorsPolicy.Validate if the designated constraints aren't met.
type CorsPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CorsPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CorsPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CorsPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CorsPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CorsPolicyValidationError) ErrorName() string { return "CorsPolicyValidationError" }

// Error satisfies the builtin error interface
func (e CorsPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCorsPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CorsPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CorsPolicyValidationError{}
