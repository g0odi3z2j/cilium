// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/wasm/v3/wasm.proto

package wasmv3

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

// Validate checks the field values on CapabilityRestrictionConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CapabilityRestrictionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CapabilityRestrictionConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CapabilityRestrictionConfigMultiError, or nil if none found.
func (m *CapabilityRestrictionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *CapabilityRestrictionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetAllowedCapabilities()))
		i := 0
		for key := range m.GetAllowedCapabilities() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetAllowedCapabilities()[key]
			_ = val

			// no validation rules for AllowedCapabilities[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, CapabilityRestrictionConfigValidationError{
							field:  fmt.Sprintf("AllowedCapabilities[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, CapabilityRestrictionConfigValidationError{
							field:  fmt.Sprintf("AllowedCapabilities[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return CapabilityRestrictionConfigValidationError{
						field:  fmt.Sprintf("AllowedCapabilities[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return CapabilityRestrictionConfigMultiError(errors)
	}

	return nil
}

// CapabilityRestrictionConfigMultiError is an error wrapping multiple
// validation errors returned by CapabilityRestrictionConfig.ValidateAll() if
// the designated constraints aren't met.
type CapabilityRestrictionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CapabilityRestrictionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CapabilityRestrictionConfigMultiError) AllErrors() []error { return m }

// CapabilityRestrictionConfigValidationError is the validation error returned
// by CapabilityRestrictionConfig.Validate if the designated constraints
// aren't met.
type CapabilityRestrictionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CapabilityRestrictionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CapabilityRestrictionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CapabilityRestrictionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CapabilityRestrictionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CapabilityRestrictionConfigValidationError) ErrorName() string {
	return "CapabilityRestrictionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CapabilityRestrictionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCapabilityRestrictionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CapabilityRestrictionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CapabilityRestrictionConfigValidationError{}

// Validate checks the field values on SanitizationConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SanitizationConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SanitizationConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SanitizationConfigMultiError, or nil if none found.
func (m *SanitizationConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *SanitizationConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SanitizationConfigMultiError(errors)
	}

	return nil
}

// SanitizationConfigMultiError is an error wrapping multiple validation errors
// returned by SanitizationConfig.ValidateAll() if the designated constraints
// aren't met.
type SanitizationConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SanitizationConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SanitizationConfigMultiError) AllErrors() []error { return m }

// SanitizationConfigValidationError is the validation error returned by
// SanitizationConfig.Validate if the designated constraints aren't met.
type SanitizationConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SanitizationConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SanitizationConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SanitizationConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SanitizationConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SanitizationConfigValidationError) ErrorName() string {
	return "SanitizationConfigValidationError"
}

// Error satisfies the builtin error interface
func (e SanitizationConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSanitizationConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SanitizationConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SanitizationConfigValidationError{}

// Validate checks the field values on VmConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *VmConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VmConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in VmConfigMultiError, or nil
// if none found.
func (m *VmConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *VmConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for VmId

	// no validation rules for Runtime

	if all {
		switch v := interface{}(m.GetCode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, VmConfigValidationError{
					field:  "Code",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, VmConfigValidationError{
					field:  "Code",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VmConfigValidationError{
				field:  "Code",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetConfiguration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, VmConfigValidationError{
					field:  "Configuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, VmConfigValidationError{
					field:  "Configuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VmConfigValidationError{
				field:  "Configuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AllowPrecompiled

	// no validation rules for NackOnCodeCacheMiss

	if all {
		switch v := interface{}(m.GetEnvironmentVariables()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, VmConfigValidationError{
					field:  "EnvironmentVariables",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, VmConfigValidationError{
					field:  "EnvironmentVariables",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnvironmentVariables()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VmConfigValidationError{
				field:  "EnvironmentVariables",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return VmConfigMultiError(errors)
	}

	return nil
}

// VmConfigMultiError is an error wrapping multiple validation errors returned
// by VmConfig.ValidateAll() if the designated constraints aren't met.
type VmConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VmConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VmConfigMultiError) AllErrors() []error { return m }

// VmConfigValidationError is the validation error returned by
// VmConfig.Validate if the designated constraints aren't met.
type VmConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VmConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VmConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VmConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VmConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VmConfigValidationError) ErrorName() string { return "VmConfigValidationError" }

// Error satisfies the builtin error interface
func (e VmConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVmConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VmConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VmConfigValidationError{}

// Validate checks the field values on EnvironmentVariables with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EnvironmentVariables) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnvironmentVariables with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EnvironmentVariablesMultiError, or nil if none found.
func (m *EnvironmentVariables) ValidateAll() error {
	return m.validate(true)
}

func (m *EnvironmentVariables) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for KeyValues

	if len(errors) > 0 {
		return EnvironmentVariablesMultiError(errors)
	}

	return nil
}

// EnvironmentVariablesMultiError is an error wrapping multiple validation
// errors returned by EnvironmentVariables.ValidateAll() if the designated
// constraints aren't met.
type EnvironmentVariablesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnvironmentVariablesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnvironmentVariablesMultiError) AllErrors() []error { return m }

// EnvironmentVariablesValidationError is the validation error returned by
// EnvironmentVariables.Validate if the designated constraints aren't met.
type EnvironmentVariablesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnvironmentVariablesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnvironmentVariablesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnvironmentVariablesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnvironmentVariablesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnvironmentVariablesValidationError) ErrorName() string {
	return "EnvironmentVariablesValidationError"
}

// Error satisfies the builtin error interface
func (e EnvironmentVariablesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnvironmentVariables.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnvironmentVariablesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnvironmentVariablesValidationError{}

// Validate checks the field values on PluginConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PluginConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PluginConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PluginConfigMultiError, or
// nil if none found.
func (m *PluginConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *PluginConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for RootId

	if all {
		switch v := interface{}(m.GetConfiguration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PluginConfigValidationError{
					field:  "Configuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PluginConfigValidationError{
					field:  "Configuration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfiguration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PluginConfigValidationError{
				field:  "Configuration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FailOpen

	if all {
		switch v := interface{}(m.GetCapabilityRestrictionConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PluginConfigValidationError{
					field:  "CapabilityRestrictionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PluginConfigValidationError{
					field:  "CapabilityRestrictionConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCapabilityRestrictionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PluginConfigValidationError{
				field:  "CapabilityRestrictionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.Vm.(type) {
	case *PluginConfig_VmConfig:
		if v == nil {
			err := PluginConfigValidationError{
				field:  "Vm",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetVmConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PluginConfigValidationError{
						field:  "VmConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PluginConfigValidationError{
						field:  "VmConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetVmConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PluginConfigValidationError{
					field:  "VmConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return PluginConfigMultiError(errors)
	}

	return nil
}

// PluginConfigMultiError is an error wrapping multiple validation errors
// returned by PluginConfig.ValidateAll() if the designated constraints aren't met.
type PluginConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PluginConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PluginConfigMultiError) AllErrors() []error { return m }

// PluginConfigValidationError is the validation error returned by
// PluginConfig.Validate if the designated constraints aren't met.
type PluginConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PluginConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PluginConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PluginConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PluginConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PluginConfigValidationError) ErrorName() string { return "PluginConfigValidationError" }

// Error satisfies the builtin error interface
func (e PluginConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPluginConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PluginConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PluginConfigValidationError{}

// Validate checks the field values on WasmService with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WasmService) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WasmService with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WasmServiceMultiError, or
// nil if none found.
func (m *WasmService) ValidateAll() error {
	return m.validate(true)
}

func (m *WasmService) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WasmServiceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WasmServiceValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WasmServiceValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Singleton

	if len(errors) > 0 {
		return WasmServiceMultiError(errors)
	}

	return nil
}

// WasmServiceMultiError is an error wrapping multiple validation errors
// returned by WasmService.ValidateAll() if the designated constraints aren't met.
type WasmServiceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WasmServiceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WasmServiceMultiError) AllErrors() []error { return m }

// WasmServiceValidationError is the validation error returned by
// WasmService.Validate if the designated constraints aren't met.
type WasmServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WasmServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WasmServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WasmServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WasmServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WasmServiceValidationError) ErrorName() string { return "WasmServiceValidationError" }

// Error satisfies the builtin error interface
func (e WasmServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWasmService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WasmServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WasmServiceValidationError{}
