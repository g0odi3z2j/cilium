// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/admin/v3/certs.proto

package adminv3

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

// Validate checks the field values on Certificates with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Certificates) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Certificates with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CertificatesMultiError, or
// nil if none found.
func (m *Certificates) ValidateAll() error {
	return m.validate(true)
}

func (m *Certificates) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCertificates() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CertificatesValidationError{
						field:  fmt.Sprintf("Certificates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CertificatesValidationError{
						field:  fmt.Sprintf("Certificates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CertificatesValidationError{
					field:  fmt.Sprintf("Certificates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CertificatesMultiError(errors)
	}

	return nil
}

// CertificatesMultiError is an error wrapping multiple validation errors
// returned by Certificates.ValidateAll() if the designated constraints aren't met.
type CertificatesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CertificatesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CertificatesMultiError) AllErrors() []error { return m }

// CertificatesValidationError is the validation error returned by
// Certificates.Validate if the designated constraints aren't met.
type CertificatesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificatesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificatesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificatesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificatesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificatesValidationError) ErrorName() string { return "CertificatesValidationError" }

// Error satisfies the builtin error interface
func (e CertificatesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificates.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificatesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificatesValidationError{}

// Validate checks the field values on Certificate with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Certificate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Certificate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CertificateMultiError, or
// nil if none found.
func (m *Certificate) ValidateAll() error {
	return m.validate(true)
}

func (m *Certificate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCaCert() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CertificateValidationError{
						field:  fmt.Sprintf("CaCert[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CertificateValidationError{
						field:  fmt.Sprintf("CaCert[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CertificateValidationError{
					field:  fmt.Sprintf("CaCert[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetCertChain() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CertificateValidationError{
						field:  fmt.Sprintf("CertChain[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CertificateValidationError{
						field:  fmt.Sprintf("CertChain[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CertificateValidationError{
					field:  fmt.Sprintf("CertChain[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CertificateMultiError(errors)
	}

	return nil
}

// CertificateMultiError is an error wrapping multiple validation errors
// returned by Certificate.ValidateAll() if the designated constraints aren't met.
type CertificateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CertificateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CertificateMultiError) AllErrors() []error { return m }

// CertificateValidationError is the validation error returned by
// Certificate.Validate if the designated constraints aren't met.
type CertificateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateValidationError) ErrorName() string { return "CertificateValidationError" }

// Error satisfies the builtin error interface
func (e CertificateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateValidationError{}

// Validate checks the field values on CertificateDetails with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CertificateDetails) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CertificateDetails with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CertificateDetailsMultiError, or nil if none found.
func (m *CertificateDetails) ValidateAll() error {
	return m.validate(true)
}

func (m *CertificateDetails) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	// no validation rules for SerialNumber

	for idx, item := range m.GetSubjectAltNames() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CertificateDetailsValidationError{
						field:  fmt.Sprintf("SubjectAltNames[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CertificateDetailsValidationError{
						field:  fmt.Sprintf("SubjectAltNames[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CertificateDetailsValidationError{
					field:  fmt.Sprintf("SubjectAltNames[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for DaysUntilExpiration

	if all {
		switch v := interface{}(m.GetValidFrom()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CertificateDetailsValidationError{
					field:  "ValidFrom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CertificateDetailsValidationError{
					field:  "ValidFrom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValidFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateDetailsValidationError{
				field:  "ValidFrom",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExpirationTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CertificateDetailsValidationError{
					field:  "ExpirationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CertificateDetailsValidationError{
					field:  "ExpirationTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpirationTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateDetailsValidationError{
				field:  "ExpirationTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOcspDetails()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CertificateDetailsValidationError{
					field:  "OcspDetails",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CertificateDetailsValidationError{
					field:  "OcspDetails",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOcspDetails()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateDetailsValidationError{
				field:  "OcspDetails",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CertificateDetailsMultiError(errors)
	}

	return nil
}

// CertificateDetailsMultiError is an error wrapping multiple validation errors
// returned by CertificateDetails.ValidateAll() if the designated constraints
// aren't met.
type CertificateDetailsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CertificateDetailsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CertificateDetailsMultiError) AllErrors() []error { return m }

// CertificateDetailsValidationError is the validation error returned by
// CertificateDetails.Validate if the designated constraints aren't met.
type CertificateDetailsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateDetailsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateDetailsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateDetailsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateDetailsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateDetailsValidationError) ErrorName() string {
	return "CertificateDetailsValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateDetailsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateDetails.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateDetailsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateDetailsValidationError{}

// Validate checks the field values on SubjectAlternateName with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubjectAlternateName) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubjectAlternateName with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubjectAlternateNameMultiError, or nil if none found.
func (m *SubjectAlternateName) ValidateAll() error {
	return m.validate(true)
}

func (m *SubjectAlternateName) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Name.(type) {
	case *SubjectAlternateName_Dns:
		if v == nil {
			err := SubjectAlternateNameValidationError{
				field:  "Name",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Dns
	case *SubjectAlternateName_Uri:
		if v == nil {
			err := SubjectAlternateNameValidationError{
				field:  "Name",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Uri
	case *SubjectAlternateName_IpAddress:
		if v == nil {
			err := SubjectAlternateNameValidationError{
				field:  "Name",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for IpAddress
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return SubjectAlternateNameMultiError(errors)
	}

	return nil
}

// SubjectAlternateNameMultiError is an error wrapping multiple validation
// errors returned by SubjectAlternateName.ValidateAll() if the designated
// constraints aren't met.
type SubjectAlternateNameMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubjectAlternateNameMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubjectAlternateNameMultiError) AllErrors() []error { return m }

// SubjectAlternateNameValidationError is the validation error returned by
// SubjectAlternateName.Validate if the designated constraints aren't met.
type SubjectAlternateNameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubjectAlternateNameValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubjectAlternateNameValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubjectAlternateNameValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubjectAlternateNameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubjectAlternateNameValidationError) ErrorName() string {
	return "SubjectAlternateNameValidationError"
}

// Error satisfies the builtin error interface
func (e SubjectAlternateNameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubjectAlternateName.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubjectAlternateNameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubjectAlternateNameValidationError{}

// Validate checks the field values on CertificateDetails_OcspDetails with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CertificateDetails_OcspDetails) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CertificateDetails_OcspDetails with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CertificateDetails_OcspDetailsMultiError, or nil if none found.
func (m *CertificateDetails_OcspDetails) ValidateAll() error {
	return m.validate(true)
}

func (m *CertificateDetails_OcspDetails) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetValidFrom()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CertificateDetails_OcspDetailsValidationError{
					field:  "ValidFrom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CertificateDetails_OcspDetailsValidationError{
					field:  "ValidFrom",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValidFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateDetails_OcspDetailsValidationError{
				field:  "ValidFrom",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExpiration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CertificateDetails_OcspDetailsValidationError{
					field:  "Expiration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CertificateDetails_OcspDetailsValidationError{
					field:  "Expiration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateDetails_OcspDetailsValidationError{
				field:  "Expiration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CertificateDetails_OcspDetailsMultiError(errors)
	}

	return nil
}

// CertificateDetails_OcspDetailsMultiError is an error wrapping multiple
// validation errors returned by CertificateDetails_OcspDetails.ValidateAll()
// if the designated constraints aren't met.
type CertificateDetails_OcspDetailsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CertificateDetails_OcspDetailsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CertificateDetails_OcspDetailsMultiError) AllErrors() []error { return m }

// CertificateDetails_OcspDetailsValidationError is the validation error
// returned by CertificateDetails_OcspDetails.Validate if the designated
// constraints aren't met.
type CertificateDetails_OcspDetailsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateDetails_OcspDetailsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateDetails_OcspDetailsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateDetails_OcspDetailsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateDetails_OcspDetailsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateDetails_OcspDetailsValidationError) ErrorName() string {
	return "CertificateDetails_OcspDetailsValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateDetails_OcspDetailsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateDetails_OcspDetails.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateDetails_OcspDetailsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateDetails_OcspDetailsValidationError{}
