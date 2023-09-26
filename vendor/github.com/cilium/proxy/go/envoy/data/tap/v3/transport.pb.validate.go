// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/data/tap/v3/transport.proto

package tapv3

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

// Validate checks the field values on Connection with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Connection) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Connection with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConnectionMultiError, or
// nil if none found.
func (m *Connection) ValidateAll() error {
	return m.validate(true)
}

func (m *Connection) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocalAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConnectionValidationError{
					field:  "LocalAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConnectionValidationError{
					field:  "LocalAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocalAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionValidationError{
				field:  "LocalAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRemoteAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConnectionValidationError{
					field:  "RemoteAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConnectionValidationError{
					field:  "RemoteAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRemoteAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionValidationError{
				field:  "RemoteAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConnectionMultiError(errors)
	}

	return nil
}

// ConnectionMultiError is an error wrapping multiple validation errors
// returned by Connection.ValidateAll() if the designated constraints aren't met.
type ConnectionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectionMultiError) AllErrors() []error { return m }

// ConnectionValidationError is the validation error returned by
// Connection.Validate if the designated constraints aren't met.
type ConnectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectionValidationError) ErrorName() string { return "ConnectionValidationError" }

// Error satisfies the builtin error interface
func (e ConnectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectionValidationError{}

// Validate checks the field values on SocketEvent with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SocketEvent) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketEvent with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SocketEventMultiError, or
// nil if none found.
func (m *SocketEvent) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketEvent) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SocketEventValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SocketEventValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SocketEventValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.EventSelector.(type) {
	case *SocketEvent_Read_:
		if v == nil {
			err := SocketEventValidationError{
				field:  "EventSelector",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetRead()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SocketEventValidationError{
						field:  "Read",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SocketEventValidationError{
						field:  "Read",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRead()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SocketEventValidationError{
					field:  "Read",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *SocketEvent_Write_:
		if v == nil {
			err := SocketEventValidationError{
				field:  "EventSelector",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetWrite()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SocketEventValidationError{
						field:  "Write",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SocketEventValidationError{
						field:  "Write",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetWrite()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SocketEventValidationError{
					field:  "Write",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *SocketEvent_Closed_:
		if v == nil {
			err := SocketEventValidationError{
				field:  "EventSelector",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetClosed()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SocketEventValidationError{
						field:  "Closed",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SocketEventValidationError{
						field:  "Closed",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetClosed()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SocketEventValidationError{
					field:  "Closed",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return SocketEventMultiError(errors)
	}

	return nil
}

// SocketEventMultiError is an error wrapping multiple validation errors
// returned by SocketEvent.ValidateAll() if the designated constraints aren't met.
type SocketEventMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketEventMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketEventMultiError) AllErrors() []error { return m }

// SocketEventValidationError is the validation error returned by
// SocketEvent.Validate if the designated constraints aren't met.
type SocketEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketEventValidationError) ErrorName() string { return "SocketEventValidationError" }

// Error satisfies the builtin error interface
func (e SocketEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketEventValidationError{}

// Validate checks the field values on SocketBufferedTrace with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SocketBufferedTrace) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketBufferedTrace with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SocketBufferedTraceMultiError, or nil if none found.
func (m *SocketBufferedTrace) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketBufferedTrace) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TraceId

	if all {
		switch v := interface{}(m.GetConnection()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SocketBufferedTraceValidationError{
					field:  "Connection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SocketBufferedTraceValidationError{
					field:  "Connection",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConnection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SocketBufferedTraceValidationError{
				field:  "Connection",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetEvents() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SocketBufferedTraceValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SocketBufferedTraceValidationError{
						field:  fmt.Sprintf("Events[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SocketBufferedTraceValidationError{
					field:  fmt.Sprintf("Events[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for ReadTruncated

	// no validation rules for WriteTruncated

	if len(errors) > 0 {
		return SocketBufferedTraceMultiError(errors)
	}

	return nil
}

// SocketBufferedTraceMultiError is an error wrapping multiple validation
// errors returned by SocketBufferedTrace.ValidateAll() if the designated
// constraints aren't met.
type SocketBufferedTraceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketBufferedTraceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketBufferedTraceMultiError) AllErrors() []error { return m }

// SocketBufferedTraceValidationError is the validation error returned by
// SocketBufferedTrace.Validate if the designated constraints aren't met.
type SocketBufferedTraceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketBufferedTraceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketBufferedTraceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketBufferedTraceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketBufferedTraceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketBufferedTraceValidationError) ErrorName() string {
	return "SocketBufferedTraceValidationError"
}

// Error satisfies the builtin error interface
func (e SocketBufferedTraceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketBufferedTrace.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketBufferedTraceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketBufferedTraceValidationError{}

// Validate checks the field values on SocketStreamedTraceSegment with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SocketStreamedTraceSegment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketStreamedTraceSegment with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SocketStreamedTraceSegmentMultiError, or nil if none found.
func (m *SocketStreamedTraceSegment) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketStreamedTraceSegment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TraceId

	switch v := m.MessagePiece.(type) {
	case *SocketStreamedTraceSegment_Connection:
		if v == nil {
			err := SocketStreamedTraceSegmentValidationError{
				field:  "MessagePiece",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetConnection()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SocketStreamedTraceSegmentValidationError{
						field:  "Connection",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SocketStreamedTraceSegmentValidationError{
						field:  "Connection",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetConnection()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SocketStreamedTraceSegmentValidationError{
					field:  "Connection",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *SocketStreamedTraceSegment_Event:
		if v == nil {
			err := SocketStreamedTraceSegmentValidationError{
				field:  "MessagePiece",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetEvent()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SocketStreamedTraceSegmentValidationError{
						field:  "Event",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SocketStreamedTraceSegmentValidationError{
						field:  "Event",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SocketStreamedTraceSegmentValidationError{
					field:  "Event",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return SocketStreamedTraceSegmentMultiError(errors)
	}

	return nil
}

// SocketStreamedTraceSegmentMultiError is an error wrapping multiple
// validation errors returned by SocketStreamedTraceSegment.ValidateAll() if
// the designated constraints aren't met.
type SocketStreamedTraceSegmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketStreamedTraceSegmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketStreamedTraceSegmentMultiError) AllErrors() []error { return m }

// SocketStreamedTraceSegmentValidationError is the validation error returned
// by SocketStreamedTraceSegment.Validate if the designated constraints aren't met.
type SocketStreamedTraceSegmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketStreamedTraceSegmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketStreamedTraceSegmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketStreamedTraceSegmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketStreamedTraceSegmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketStreamedTraceSegmentValidationError) ErrorName() string {
	return "SocketStreamedTraceSegmentValidationError"
}

// Error satisfies the builtin error interface
func (e SocketStreamedTraceSegmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketStreamedTraceSegment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketStreamedTraceSegmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketStreamedTraceSegmentValidationError{}

// Validate checks the field values on SocketEvent_Read with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SocketEvent_Read) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketEvent_Read with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SocketEvent_ReadMultiError, or nil if none found.
func (m *SocketEvent_Read) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketEvent_Read) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SocketEvent_ReadValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SocketEvent_ReadValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SocketEvent_ReadValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SocketEvent_ReadMultiError(errors)
	}

	return nil
}

// SocketEvent_ReadMultiError is an error wrapping multiple validation errors
// returned by SocketEvent_Read.ValidateAll() if the designated constraints
// aren't met.
type SocketEvent_ReadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketEvent_ReadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketEvent_ReadMultiError) AllErrors() []error { return m }

// SocketEvent_ReadValidationError is the validation error returned by
// SocketEvent_Read.Validate if the designated constraints aren't met.
type SocketEvent_ReadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketEvent_ReadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketEvent_ReadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketEvent_ReadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketEvent_ReadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketEvent_ReadValidationError) ErrorName() string { return "SocketEvent_ReadValidationError" }

// Error satisfies the builtin error interface
func (e SocketEvent_ReadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketEvent_Read.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketEvent_ReadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketEvent_ReadValidationError{}

// Validate checks the field values on SocketEvent_Write with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SocketEvent_Write) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketEvent_Write with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SocketEvent_WriteMultiError, or nil if none found.
func (m *SocketEvent_Write) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketEvent_Write) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SocketEvent_WriteValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SocketEvent_WriteValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SocketEvent_WriteValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for EndStream

	if len(errors) > 0 {
		return SocketEvent_WriteMultiError(errors)
	}

	return nil
}

// SocketEvent_WriteMultiError is an error wrapping multiple validation errors
// returned by SocketEvent_Write.ValidateAll() if the designated constraints
// aren't met.
type SocketEvent_WriteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketEvent_WriteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketEvent_WriteMultiError) AllErrors() []error { return m }

// SocketEvent_WriteValidationError is the validation error returned by
// SocketEvent_Write.Validate if the designated constraints aren't met.
type SocketEvent_WriteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketEvent_WriteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketEvent_WriteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketEvent_WriteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketEvent_WriteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketEvent_WriteValidationError) ErrorName() string {
	return "SocketEvent_WriteValidationError"
}

// Error satisfies the builtin error interface
func (e SocketEvent_WriteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketEvent_Write.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketEvent_WriteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketEvent_WriteValidationError{}

// Validate checks the field values on SocketEvent_Closed with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SocketEvent_Closed) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketEvent_Closed with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SocketEvent_ClosedMultiError, or nil if none found.
func (m *SocketEvent_Closed) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketEvent_Closed) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SocketEvent_ClosedMultiError(errors)
	}

	return nil
}

// SocketEvent_ClosedMultiError is an error wrapping multiple validation errors
// returned by SocketEvent_Closed.ValidateAll() if the designated constraints
// aren't met.
type SocketEvent_ClosedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketEvent_ClosedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketEvent_ClosedMultiError) AllErrors() []error { return m }

// SocketEvent_ClosedValidationError is the validation error returned by
// SocketEvent_Closed.Validate if the designated constraints aren't met.
type SocketEvent_ClosedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketEvent_ClosedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketEvent_ClosedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketEvent_ClosedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketEvent_ClosedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketEvent_ClosedValidationError) ErrorName() string {
	return "SocketEvent_ClosedValidationError"
}

// Error satisfies the builtin error interface
func (e SocketEvent_ClosedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketEvent_Closed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketEvent_ClosedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketEvent_ClosedValidationError{}
