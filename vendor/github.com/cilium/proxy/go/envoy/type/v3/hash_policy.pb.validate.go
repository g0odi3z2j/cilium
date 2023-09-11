// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/v3/hash_policy.proto

package typev3

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

// Validate checks the field values on HashPolicy with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HashPolicy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HashPolicy with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HashPolicyMultiError, or
// nil if none found.
func (m *HashPolicy) ValidateAll() error {
	return m.validate(true)
}

func (m *HashPolicy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofPolicySpecifierPresent := false
	switch v := m.PolicySpecifier.(type) {
	case *HashPolicy_SourceIp_:
		if v == nil {
			err := HashPolicyValidationError{
				field:  "PolicySpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofPolicySpecifierPresent = true

		if all {
			switch v := interface{}(m.GetSourceIp()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HashPolicyValidationError{
						field:  "SourceIp",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HashPolicyValidationError{
						field:  "SourceIp",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSourceIp()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HashPolicyValidationError{
					field:  "SourceIp",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HashPolicy_FilterState_:
		if v == nil {
			err := HashPolicyValidationError{
				field:  "PolicySpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofPolicySpecifierPresent = true

		if all {
			switch v := interface{}(m.GetFilterState()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HashPolicyValidationError{
						field:  "FilterState",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HashPolicyValidationError{
						field:  "FilterState",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFilterState()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HashPolicyValidationError{
					field:  "FilterState",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofPolicySpecifierPresent {
		err := HashPolicyValidationError{
			field:  "PolicySpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HashPolicyMultiError(errors)
	}

	return nil
}

// HashPolicyMultiError is an error wrapping multiple validation errors
// returned by HashPolicy.ValidateAll() if the designated constraints aren't met.
type HashPolicyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HashPolicyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HashPolicyMultiError) AllErrors() []error { return m }

// HashPolicyValidationError is the validation error returned by
// HashPolicy.Validate if the designated constraints aren't met.
type HashPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HashPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HashPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HashPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HashPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HashPolicyValidationError) ErrorName() string { return "HashPolicyValidationError" }

// Error satisfies the builtin error interface
func (e HashPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHashPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HashPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HashPolicyValidationError{}

// Validate checks the field values on HashPolicy_SourceIp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HashPolicy_SourceIp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HashPolicy_SourceIp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HashPolicy_SourceIpMultiError, or nil if none found.
func (m *HashPolicy_SourceIp) ValidateAll() error {
	return m.validate(true)
}

func (m *HashPolicy_SourceIp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return HashPolicy_SourceIpMultiError(errors)
	}

	return nil
}

// HashPolicy_SourceIpMultiError is an error wrapping multiple validation
// errors returned by HashPolicy_SourceIp.ValidateAll() if the designated
// constraints aren't met.
type HashPolicy_SourceIpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HashPolicy_SourceIpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HashPolicy_SourceIpMultiError) AllErrors() []error { return m }

// HashPolicy_SourceIpValidationError is the validation error returned by
// HashPolicy_SourceIp.Validate if the designated constraints aren't met.
type HashPolicy_SourceIpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HashPolicy_SourceIpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HashPolicy_SourceIpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HashPolicy_SourceIpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HashPolicy_SourceIpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HashPolicy_SourceIpValidationError) ErrorName() string {
	return "HashPolicy_SourceIpValidationError"
}

// Error satisfies the builtin error interface
func (e HashPolicy_SourceIpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHashPolicy_SourceIp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HashPolicy_SourceIpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HashPolicy_SourceIpValidationError{}

// Validate checks the field values on HashPolicy_FilterState with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HashPolicy_FilterState) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HashPolicy_FilterState with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HashPolicy_FilterStateMultiError, or nil if none found.
func (m *HashPolicy_FilterState) ValidateAll() error {
	return m.validate(true)
}

func (m *HashPolicy_FilterState) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := HashPolicy_FilterStateValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HashPolicy_FilterStateMultiError(errors)
	}

	return nil
}

// HashPolicy_FilterStateMultiError is an error wrapping multiple validation
// errors returned by HashPolicy_FilterState.ValidateAll() if the designated
// constraints aren't met.
type HashPolicy_FilterStateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HashPolicy_FilterStateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HashPolicy_FilterStateMultiError) AllErrors() []error { return m }

// HashPolicy_FilterStateValidationError is the validation error returned by
// HashPolicy_FilterState.Validate if the designated constraints aren't met.
type HashPolicy_FilterStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HashPolicy_FilterStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HashPolicy_FilterStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HashPolicy_FilterStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HashPolicy_FilterStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HashPolicy_FilterStateValidationError) ErrorName() string {
	return "HashPolicy_FilterStateValidationError"
}

// Error satisfies the builtin error interface
func (e HashPolicy_FilterStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHashPolicy_FilterState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HashPolicy_FilterStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HashPolicy_FilterStateValidationError{}
