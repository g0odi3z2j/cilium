// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/access_loggers/stream/v3/stream.proto

package streamv3

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

// Validate checks the field values on StdoutAccessLog with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StdoutAccessLog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StdoutAccessLog with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StdoutAccessLogMultiError, or nil if none found.
func (m *StdoutAccessLog) ValidateAll() error {
	return m.validate(true)
}

func (m *StdoutAccessLog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.AccessLogFormat.(type) {

	case *StdoutAccessLog_LogFormat:

		if m.GetLogFormat() == nil {
			err := StdoutAccessLogValidationError{
				field:  "LogFormat",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetLogFormat()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StdoutAccessLogValidationError{
						field:  "LogFormat",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StdoutAccessLogValidationError{
						field:  "LogFormat",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLogFormat()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StdoutAccessLogValidationError{
					field:  "LogFormat",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return StdoutAccessLogMultiError(errors)
	}
	return nil
}

// StdoutAccessLogMultiError is an error wrapping multiple validation errors
// returned by StdoutAccessLog.ValidateAll() if the designated constraints
// aren't met.
type StdoutAccessLogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StdoutAccessLogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StdoutAccessLogMultiError) AllErrors() []error { return m }

// StdoutAccessLogValidationError is the validation error returned by
// StdoutAccessLog.Validate if the designated constraints aren't met.
type StdoutAccessLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StdoutAccessLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StdoutAccessLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StdoutAccessLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StdoutAccessLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StdoutAccessLogValidationError) ErrorName() string { return "StdoutAccessLogValidationError" }

// Error satisfies the builtin error interface
func (e StdoutAccessLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStdoutAccessLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StdoutAccessLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StdoutAccessLogValidationError{}

// Validate checks the field values on StderrAccessLog with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StderrAccessLog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StderrAccessLog with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StderrAccessLogMultiError, or nil if none found.
func (m *StderrAccessLog) ValidateAll() error {
	return m.validate(true)
}

func (m *StderrAccessLog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.AccessLogFormat.(type) {

	case *StderrAccessLog_LogFormat:

		if m.GetLogFormat() == nil {
			err := StderrAccessLogValidationError{
				field:  "LogFormat",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetLogFormat()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, StderrAccessLogValidationError{
						field:  "LogFormat",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StderrAccessLogValidationError{
						field:  "LogFormat",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetLogFormat()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StderrAccessLogValidationError{
					field:  "LogFormat",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return StderrAccessLogMultiError(errors)
	}
	return nil
}

// StderrAccessLogMultiError is an error wrapping multiple validation errors
// returned by StderrAccessLog.ValidateAll() if the designated constraints
// aren't met.
type StderrAccessLogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StderrAccessLogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StderrAccessLogMultiError) AllErrors() []error { return m }

// StderrAccessLogValidationError is the validation error returned by
// StderrAccessLog.Validate if the designated constraints aren't met.
type StderrAccessLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StderrAccessLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StderrAccessLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StderrAccessLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StderrAccessLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StderrAccessLogValidationError) ErrorName() string { return "StderrAccessLogValidationError" }

// Error satisfies the builtin error interface
func (e StderrAccessLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStderrAccessLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StderrAccessLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StderrAccessLogValidationError{}
