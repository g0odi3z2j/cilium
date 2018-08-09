// Code generated by protoc-gen-validate
// source: cilium/accesslog.proto
// DO NOT EDIT!!!

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
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e KeyValueValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyValue.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = KeyValueValidationError{}

// Validate checks the field values on HttpLogEntry with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpLogEntry) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Timestamp

	// no validation rules for HttpProtocol

	// no validation rules for EntryType

	// no validation rules for PolicyName

	// no validation rules for CiliumRuleRef

	// no validation rules for SourceSecurityId

	// no validation rules for SourceAddress

	// no validation rules for DestinationAddress

	// no validation rules for Scheme

	// no validation rules for Host

	// no validation rules for Path

	// no validation rules for Method

	// no validation rules for Status

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpLogEntryValidationError{
					Field:  fmt.Sprintf("Headers[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	// no validation rules for IsIngress

	return nil
}

// HttpLogEntryValidationError is the validation error returned by
// HttpLogEntry.Validate if the designated constraints aren't met.
type HttpLogEntryValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HttpLogEntryValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpLogEntry.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HttpLogEntryValidationError{}
