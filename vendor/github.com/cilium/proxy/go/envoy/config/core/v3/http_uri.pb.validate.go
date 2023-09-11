// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/http_uri.proto

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

// Validate checks the field values on HttpUri with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HttpUri) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HttpUri with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HttpUriMultiError, or nil if none found.
func (m *HttpUri) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpUri) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUri()) < 1 {
		err := HttpUriValidationError{
			field:  "Uri",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTimeout() == nil {
		err := HttpUriValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetTimeout(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = HttpUriValidationError{
				field:  "Timeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gte := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur < gte {
				err := HttpUriValidationError{
					field:  "Timeout",
					reason: "value must be greater than or equal to 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	oneofHttpUpstreamTypePresent := false
	switch v := m.HttpUpstreamType.(type) {
	case *HttpUri_Cluster:
		if v == nil {
			err := HttpUriValidationError{
				field:  "HttpUpstreamType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofHttpUpstreamTypePresent = true

		if utf8.RuneCountInString(m.GetCluster()) < 1 {
			err := HttpUriValidationError{
				field:  "Cluster",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofHttpUpstreamTypePresent {
		err := HttpUriValidationError{
			field:  "HttpUpstreamType",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HttpUriMultiError(errors)
	}

	return nil
}

// HttpUriMultiError is an error wrapping multiple validation errors returned
// by HttpUri.ValidateAll() if the designated constraints aren't met.
type HttpUriMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpUriMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpUriMultiError) AllErrors() []error { return m }

// HttpUriValidationError is the validation error returned by HttpUri.Validate
// if the designated constraints aren't met.
type HttpUriValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpUriValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpUriValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpUriValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpUriValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpUriValidationError) ErrorName() string { return "HttpUriValidationError" }

// Error satisfies the builtin error interface
func (e HttpUriValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpUri.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpUriValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpUriValidationError{}
