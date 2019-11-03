// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cilium/api/accesslog.proto

package cilium

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _accesslog_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on KeyValue with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *KeyValue) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Key

	// no validation rules for Value

	return nil
}

// KeyValueValidationError is the validation error returned by
// KeyValue.Validate if the designated constraints aren't met.
type KeyValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyValueValidationError) ErrorName() string { return "KeyValueValidationError" }

// Error satisfies the builtin error interface
func (e KeyValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyValueValidationError{}

// Validate checks the field values on HttpLogEntry with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpLogEntry) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for HttpProtocol

	// no validation rules for Scheme

	// no validation rules for Host

	// no validation rules for Path

	// no validation rules for Method

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpLogEntryValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Status

	for idx, item := range m.GetMissingHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpLogEntryValidationError{
					field:  fmt.Sprintf("MissingHeaders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetRejectedHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpLogEntryValidationError{
					field:  fmt.Sprintf("RejectedHeaders[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HttpLogEntryValidationError is the validation error returned by
// HttpLogEntry.Validate if the designated constraints aren't met.
type HttpLogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpLogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpLogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpLogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpLogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpLogEntryValidationError) ErrorName() string { return "HttpLogEntryValidationError" }

// Error satisfies the builtin error interface
func (e HttpLogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpLogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpLogEntryValidationError{}

// Validate checks the field values on L7LogEntry with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *L7LogEntry) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Proto

	// no validation rules for Fields

	return nil
}

// L7LogEntryValidationError is the validation error returned by
// L7LogEntry.Validate if the designated constraints aren't met.
type L7LogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e L7LogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e L7LogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e L7LogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e L7LogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e L7LogEntryValidationError) ErrorName() string { return "L7LogEntryValidationError" }

// Error satisfies the builtin error interface
func (e L7LogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sL7LogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = L7LogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = L7LogEntryValidationError{}

// Validate checks the field values on LogEntry with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *LogEntry) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Timestamp

	// no validation rules for IsIngress

	// no validation rules for EntryType

	// no validation rules for PolicyName

	// no validation rules for CiliumRuleRef

	// no validation rules for SourceSecurityId

	// no validation rules for DestinationSecurityId

	// no validation rules for SourceAddress

	// no validation rules for DestinationAddress

	// no validation rules for HttpProtocol

	// no validation rules for Scheme

	// no validation rules for Host

	// no validation rules for Path

	// no validation rules for Method

	// no validation rules for Status

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LogEntryValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.L7.(type) {

	case *LogEntry_Http:

		if v, ok := interface{}(m.GetHttp()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LogEntryValidationError{
					field:  "Http",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *LogEntry_GenericL7:

		if v, ok := interface{}(m.GetGenericL7()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LogEntryValidationError{
					field:  "GenericL7",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LogEntryValidationError is the validation error returned by
// LogEntry.Validate if the designated constraints aren't met.
type LogEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogEntryValidationError) ErrorName() string { return "LogEntryValidationError" }

// Error satisfies the builtin error interface
func (e LogEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogEntryValidationError{}
