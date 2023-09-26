// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/retry/host/omit_host_metadata/v3/omit_host_metadata_config.proto

package omit_host_metadatav3

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

// Validate checks the field values on OmitHostMetadataConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OmitHostMetadataConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OmitHostMetadataConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OmitHostMetadataConfigMultiError, or nil if none found.
func (m *OmitHostMetadataConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *OmitHostMetadataConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMetadataMatch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OmitHostMetadataConfigValidationError{
					field:  "MetadataMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OmitHostMetadataConfigValidationError{
					field:  "MetadataMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadataMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OmitHostMetadataConfigValidationError{
				field:  "MetadataMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OmitHostMetadataConfigMultiError(errors)
	}

	return nil
}

// OmitHostMetadataConfigMultiError is an error wrapping multiple validation
// errors returned by OmitHostMetadataConfig.ValidateAll() if the designated
// constraints aren't met.
type OmitHostMetadataConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OmitHostMetadataConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OmitHostMetadataConfigMultiError) AllErrors() []error { return m }

// OmitHostMetadataConfigValidationError is the validation error returned by
// OmitHostMetadataConfig.Validate if the designated constraints aren't met.
type OmitHostMetadataConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OmitHostMetadataConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OmitHostMetadataConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OmitHostMetadataConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OmitHostMetadataConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OmitHostMetadataConfigValidationError) ErrorName() string {
	return "OmitHostMetadataConfigValidationError"
}

// Error satisfies the builtin error interface
func (e OmitHostMetadataConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOmitHostMetadataConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OmitHostMetadataConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OmitHostMetadataConfigValidationError{}
