// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/load_balancing_policies/common/v3/common.proto

package commonv3

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

// Validate checks the field values on LocalityLbConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LocalityLbConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalityLbConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LocalityLbConfigMultiError, or nil if none found.
func (m *LocalityLbConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalityLbConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofLocalityConfigSpecifierPresent := false
	switch v := m.LocalityConfigSpecifier.(type) {
	case *LocalityLbConfig_ZoneAwareLbConfig_:
		if v == nil {
			err := LocalityLbConfigValidationError{
				field:  "LocalityConfigSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofLocalityConfigSpecifierPresent = true

		if all {
			switch v := interface{}(m.GetZoneAwareLbConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalityLbConfigValidationError{
						field:  "ZoneAwareLbConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalityLbConfigValidationError{
						field:  "ZoneAwareLbConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetZoneAwareLbConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityLbConfigValidationError{
					field:  "ZoneAwareLbConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LocalityLbConfig_LocalityWeightedLbConfig_:
		if v == nil {
			err := LocalityLbConfigValidationError{
				field:  "LocalityConfigSpecifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofLocalityConfigSpecifierPresent = true

		if all {
			switch v := interface{}(m.GetLocalityWeightedLbConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalityLbConfigValidationError{
						field:  "LocalityWeightedLbConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalityLbConfigValidationError{
						field:  "LocalityWeightedLbConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLocalityWeightedLbConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityLbConfigValidationError{
					field:  "LocalityWeightedLbConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofLocalityConfigSpecifierPresent {
		err := LocalityLbConfigValidationError{
			field:  "LocalityConfigSpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LocalityLbConfigMultiError(errors)
	}

	return nil
}

// LocalityLbConfigMultiError is an error wrapping multiple validation errors
// returned by LocalityLbConfig.ValidateAll() if the designated constraints
// aren't met.
type LocalityLbConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalityLbConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalityLbConfigMultiError) AllErrors() []error { return m }

// LocalityLbConfigValidationError is the validation error returned by
// LocalityLbConfig.Validate if the designated constraints aren't met.
type LocalityLbConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityLbConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityLbConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityLbConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityLbConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityLbConfigValidationError) ErrorName() string { return "LocalityLbConfigValidationError" }

// Error satisfies the builtin error interface
func (e LocalityLbConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityLbConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityLbConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityLbConfigValidationError{}

// Validate checks the field values on SlowStartConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SlowStartConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SlowStartConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SlowStartConfigMultiError, or nil if none found.
func (m *SlowStartConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *SlowStartConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSlowStartWindow()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SlowStartConfigValidationError{
					field:  "SlowStartWindow",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SlowStartConfigValidationError{
					field:  "SlowStartWindow",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSlowStartWindow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SlowStartConfigValidationError{
				field:  "SlowStartWindow",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAggression()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SlowStartConfigValidationError{
					field:  "Aggression",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SlowStartConfigValidationError{
					field:  "Aggression",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAggression()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SlowStartConfigValidationError{
				field:  "Aggression",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMinWeightPercent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SlowStartConfigValidationError{
					field:  "MinWeightPercent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SlowStartConfigValidationError{
					field:  "MinWeightPercent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMinWeightPercent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SlowStartConfigValidationError{
				field:  "MinWeightPercent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SlowStartConfigMultiError(errors)
	}

	return nil
}

// SlowStartConfigMultiError is an error wrapping multiple validation errors
// returned by SlowStartConfig.ValidateAll() if the designated constraints
// aren't met.
type SlowStartConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SlowStartConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SlowStartConfigMultiError) AllErrors() []error { return m }

// SlowStartConfigValidationError is the validation error returned by
// SlowStartConfig.Validate if the designated constraints aren't met.
type SlowStartConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SlowStartConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SlowStartConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SlowStartConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SlowStartConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SlowStartConfigValidationError) ErrorName() string { return "SlowStartConfigValidationError" }

// Error satisfies the builtin error interface
func (e SlowStartConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSlowStartConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SlowStartConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SlowStartConfigValidationError{}

// Validate checks the field values on ConsistentHashingLbConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ConsistentHashingLbConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConsistentHashingLbConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConsistentHashingLbConfigMultiError, or nil if none found.
func (m *ConsistentHashingLbConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ConsistentHashingLbConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UseHostnameForHashing

	if wrapper := m.GetHashBalanceFactor(); wrapper != nil {

		if wrapper.GetValue() < 100 {
			err := ConsistentHashingLbConfigValidationError{
				field:  "HashBalanceFactor",
				reason: "value must be greater than or equal to 100",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ConsistentHashingLbConfigMultiError(errors)
	}

	return nil
}

// ConsistentHashingLbConfigMultiError is an error wrapping multiple validation
// errors returned by ConsistentHashingLbConfig.ValidateAll() if the
// designated constraints aren't met.
type ConsistentHashingLbConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsistentHashingLbConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsistentHashingLbConfigMultiError) AllErrors() []error { return m }

// ConsistentHashingLbConfigValidationError is the validation error returned by
// ConsistentHashingLbConfig.Validate if the designated constraints aren't met.
type ConsistentHashingLbConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsistentHashingLbConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsistentHashingLbConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsistentHashingLbConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsistentHashingLbConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsistentHashingLbConfigValidationError) ErrorName() string {
	return "ConsistentHashingLbConfigValidationError"
}

// Error satisfies the builtin error interface
func (e ConsistentHashingLbConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsistentHashingLbConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsistentHashingLbConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsistentHashingLbConfigValidationError{}

// Validate checks the field values on LocalityLbConfig_ZoneAwareLbConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *LocalityLbConfig_ZoneAwareLbConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalityLbConfig_ZoneAwareLbConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// LocalityLbConfig_ZoneAwareLbConfigMultiError, or nil if none found.
func (m *LocalityLbConfig_ZoneAwareLbConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalityLbConfig_ZoneAwareLbConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRoutingEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalityLbConfig_ZoneAwareLbConfigValidationError{
					field:  "RoutingEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalityLbConfig_ZoneAwareLbConfigValidationError{
					field:  "RoutingEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRoutingEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityLbConfig_ZoneAwareLbConfigValidationError{
				field:  "RoutingEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMinClusterSize()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalityLbConfig_ZoneAwareLbConfigValidationError{
					field:  "MinClusterSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalityLbConfig_ZoneAwareLbConfigValidationError{
					field:  "MinClusterSize",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMinClusterSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityLbConfig_ZoneAwareLbConfigValidationError{
				field:  "MinClusterSize",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FailTrafficOnPanic

	if len(errors) > 0 {
		return LocalityLbConfig_ZoneAwareLbConfigMultiError(errors)
	}

	return nil
}

// LocalityLbConfig_ZoneAwareLbConfigMultiError is an error wrapping multiple
// validation errors returned by
// LocalityLbConfig_ZoneAwareLbConfig.ValidateAll() if the designated
// constraints aren't met.
type LocalityLbConfig_ZoneAwareLbConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalityLbConfig_ZoneAwareLbConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalityLbConfig_ZoneAwareLbConfigMultiError) AllErrors() []error { return m }

// LocalityLbConfig_ZoneAwareLbConfigValidationError is the validation error
// returned by LocalityLbConfig_ZoneAwareLbConfig.Validate if the designated
// constraints aren't met.
type LocalityLbConfig_ZoneAwareLbConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityLbConfig_ZoneAwareLbConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityLbConfig_ZoneAwareLbConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityLbConfig_ZoneAwareLbConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityLbConfig_ZoneAwareLbConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityLbConfig_ZoneAwareLbConfigValidationError) ErrorName() string {
	return "LocalityLbConfig_ZoneAwareLbConfigValidationError"
}

// Error satisfies the builtin error interface
func (e LocalityLbConfig_ZoneAwareLbConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityLbConfig_ZoneAwareLbConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityLbConfig_ZoneAwareLbConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityLbConfig_ZoneAwareLbConfigValidationError{}

// Validate checks the field values on
// LocalityLbConfig_LocalityWeightedLbConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LocalityLbConfig_LocalityWeightedLbConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// LocalityLbConfig_LocalityWeightedLbConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// LocalityLbConfig_LocalityWeightedLbConfigMultiError, or nil if none found.
func (m *LocalityLbConfig_LocalityWeightedLbConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalityLbConfig_LocalityWeightedLbConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LocalityLbConfig_LocalityWeightedLbConfigMultiError(errors)
	}

	return nil
}

// LocalityLbConfig_LocalityWeightedLbConfigMultiError is an error wrapping
// multiple validation errors returned by
// LocalityLbConfig_LocalityWeightedLbConfig.ValidateAll() if the designated
// constraints aren't met.
type LocalityLbConfig_LocalityWeightedLbConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalityLbConfig_LocalityWeightedLbConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalityLbConfig_LocalityWeightedLbConfigMultiError) AllErrors() []error { return m }

// LocalityLbConfig_LocalityWeightedLbConfigValidationError is the validation
// error returned by LocalityLbConfig_LocalityWeightedLbConfig.Validate if the
// designated constraints aren't met.
type LocalityLbConfig_LocalityWeightedLbConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityLbConfig_LocalityWeightedLbConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityLbConfig_LocalityWeightedLbConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityLbConfig_LocalityWeightedLbConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityLbConfig_LocalityWeightedLbConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityLbConfig_LocalityWeightedLbConfigValidationError) ErrorName() string {
	return "LocalityLbConfig_LocalityWeightedLbConfigValidationError"
}

// Error satisfies the builtin error interface
func (e LocalityLbConfig_LocalityWeightedLbConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityLbConfig_LocalityWeightedLbConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityLbConfig_LocalityWeightedLbConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityLbConfig_LocalityWeightedLbConfigValidationError{}
