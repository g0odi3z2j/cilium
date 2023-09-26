// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/grpc_field_extraction/v3/config.proto

package grpc_field_extractionv3

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

// Validate checks the field values on GrpcFieldExtractionConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GrpcFieldExtractionConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GrpcFieldExtractionConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GrpcFieldExtractionConfigMultiError, or nil if none found.
func (m *GrpcFieldExtractionConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *GrpcFieldExtractionConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDescriptorSet() == nil {
		err := GrpcFieldExtractionConfigValidationError{
			field:  "DescriptorSet",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDescriptorSet()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GrpcFieldExtractionConfigValidationError{
					field:  "DescriptorSet",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GrpcFieldExtractionConfigValidationError{
					field:  "DescriptorSet",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDescriptorSet()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GrpcFieldExtractionConfigValidationError{
				field:  "DescriptorSet",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	{
		sorted_keys := make([]string, len(m.GetExtractionsByMethod()))
		i := 0
		for key := range m.GetExtractionsByMethod() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetExtractionsByMethod()[key]
			_ = val

			// no validation rules for ExtractionsByMethod[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, GrpcFieldExtractionConfigValidationError{
							field:  fmt.Sprintf("ExtractionsByMethod[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, GrpcFieldExtractionConfigValidationError{
							field:  fmt.Sprintf("ExtractionsByMethod[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return GrpcFieldExtractionConfigValidationError{
						field:  fmt.Sprintf("ExtractionsByMethod[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return GrpcFieldExtractionConfigMultiError(errors)
	}

	return nil
}

// GrpcFieldExtractionConfigMultiError is an error wrapping multiple validation
// errors returned by GrpcFieldExtractionConfig.ValidateAll() if the
// designated constraints aren't met.
type GrpcFieldExtractionConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GrpcFieldExtractionConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GrpcFieldExtractionConfigMultiError) AllErrors() []error { return m }

// GrpcFieldExtractionConfigValidationError is the validation error returned by
// GrpcFieldExtractionConfig.Validate if the designated constraints aren't met.
type GrpcFieldExtractionConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcFieldExtractionConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcFieldExtractionConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcFieldExtractionConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcFieldExtractionConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcFieldExtractionConfigValidationError) ErrorName() string {
	return "GrpcFieldExtractionConfigValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcFieldExtractionConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcFieldExtractionConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcFieldExtractionConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcFieldExtractionConfigValidationError{}

// Validate checks the field values on FieldExtractions with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FieldExtractions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FieldExtractions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FieldExtractionsMultiError, or nil if none found.
func (m *FieldExtractions) ValidateAll() error {
	return m.validate(true)
}

func (m *FieldExtractions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetRequestFieldExtractions()))
		i := 0
		for key := range m.GetRequestFieldExtractions() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetRequestFieldExtractions()[key]
			_ = val

			// no validation rules for RequestFieldExtractions[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, FieldExtractionsValidationError{
							field:  fmt.Sprintf("RequestFieldExtractions[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, FieldExtractionsValidationError{
							field:  fmt.Sprintf("RequestFieldExtractions[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return FieldExtractionsValidationError{
						field:  fmt.Sprintf("RequestFieldExtractions[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return FieldExtractionsMultiError(errors)
	}

	return nil
}

// FieldExtractionsMultiError is an error wrapping multiple validation errors
// returned by FieldExtractions.ValidateAll() if the designated constraints
// aren't met.
type FieldExtractionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldExtractionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldExtractionsMultiError) AllErrors() []error { return m }

// FieldExtractionsValidationError is the validation error returned by
// FieldExtractions.Validate if the designated constraints aren't met.
type FieldExtractionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldExtractionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldExtractionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldExtractionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldExtractionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldExtractionsValidationError) ErrorName() string { return "FieldExtractionsValidationError" }

// Error satisfies the builtin error interface
func (e FieldExtractionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFieldExtractions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldExtractionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldExtractionsValidationError{}

// Validate checks the field values on RequestFieldValueDisposition with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RequestFieldValueDisposition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RequestFieldValueDisposition with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RequestFieldValueDispositionMultiError, or nil if none found.
func (m *RequestFieldValueDisposition) ValidateAll() error {
	return m.validate(true)
}

func (m *RequestFieldValueDisposition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Disposition.(type) {
	case *RequestFieldValueDisposition_DynamicMetadata:
		if v == nil {
			err := RequestFieldValueDispositionValidationError{
				field:  "Disposition",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for DynamicMetadata
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return RequestFieldValueDispositionMultiError(errors)
	}

	return nil
}

// RequestFieldValueDispositionMultiError is an error wrapping multiple
// validation errors returned by RequestFieldValueDisposition.ValidateAll() if
// the designated constraints aren't met.
type RequestFieldValueDispositionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestFieldValueDispositionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestFieldValueDispositionMultiError) AllErrors() []error { return m }

// RequestFieldValueDispositionValidationError is the validation error returned
// by RequestFieldValueDisposition.Validate if the designated constraints
// aren't met.
type RequestFieldValueDispositionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestFieldValueDispositionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestFieldValueDispositionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestFieldValueDispositionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestFieldValueDispositionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestFieldValueDispositionValidationError) ErrorName() string {
	return "RequestFieldValueDispositionValidationError"
}

// Error satisfies the builtin error interface
func (e RequestFieldValueDispositionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestFieldValueDisposition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestFieldValueDispositionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestFieldValueDispositionValidationError{}
