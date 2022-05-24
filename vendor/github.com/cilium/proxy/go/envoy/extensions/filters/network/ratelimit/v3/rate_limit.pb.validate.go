// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/ratelimit/v3/rate_limit.proto

package ratelimitv3

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

// Validate checks the field values on RateLimit with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RateLimit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimit with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RateLimitMultiError, or nil
// if none found.
func (m *RateLimit) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := RateLimitValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDomain()) < 1 {
		err := RateLimitValidationError{
			field:  "Domain",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetDescriptors()) < 1 {
		err := RateLimitValidationError{
			field:  "Descriptors",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetDescriptors() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RateLimitValidationError{
						field:  fmt.Sprintf("Descriptors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RateLimitValidationError{
						field:  fmt.Sprintf("Descriptors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitValidationError{
					field:  fmt.Sprintf("Descriptors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FailureModeDeny

	if m.GetRateLimitService() == nil {
		err := RateLimitValidationError{
			field:  "RateLimitService",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRateLimitService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitValidationError{
					field:  "RateLimitService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitValidationError{
					field:  "RateLimitService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRateLimitService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitValidationError{
				field:  "RateLimitService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RateLimitMultiError(errors)
	}
	return nil
}

// RateLimitMultiError is an error wrapping multiple validation errors returned
// by RateLimit.ValidateAll() if the designated constraints aren't met.
type RateLimitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitMultiError) AllErrors() []error { return m }

// RateLimitValidationError is the validation error returned by
// RateLimit.Validate if the designated constraints aren't met.
type RateLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitValidationError) ErrorName() string { return "RateLimitValidationError" }

// Error satisfies the builtin error interface
func (e RateLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitValidationError{}
