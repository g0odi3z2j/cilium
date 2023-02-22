// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/address.proto

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

// Validate checks the field values on Pipe with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Pipe) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pipe with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PipeMultiError, or nil if none found.
func (m *Pipe) ValidateAll() error {
	return m.validate(true)
}

func (m *Pipe) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		err := PipeValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMode() > 511 {
		err := PipeValidationError{
			field:  "Mode",
			reason: "value must be less than or equal to 511",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PipeMultiError(errors)
	}
	return nil
}

// PipeMultiError is an error wrapping multiple validation errors returned by
// Pipe.ValidateAll() if the designated constraints aren't met.
type PipeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PipeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PipeMultiError) AllErrors() []error { return m }

// PipeValidationError is the validation error returned by Pipe.Validate if the
// designated constraints aren't met.
type PipeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PipeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PipeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PipeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PipeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PipeValidationError) ErrorName() string { return "PipeValidationError" }

// Error satisfies the builtin error interface
func (e PipeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPipe.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PipeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PipeValidationError{}

// Validate checks the field values on EnvoyInternalAddress with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EnvoyInternalAddress) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnvoyInternalAddress with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EnvoyInternalAddressMultiError, or nil if none found.
func (m *EnvoyInternalAddress) ValidateAll() error {
	return m.validate(true)
}

func (m *EnvoyInternalAddress) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EndpointId

	switch m.AddressNameSpecifier.(type) {

	case *EnvoyInternalAddress_ServerListenerName:
		// no validation rules for ServerListenerName

	default:
		err := EnvoyInternalAddressValidationError{
			field:  "AddressNameSpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return EnvoyInternalAddressMultiError(errors)
	}
	return nil
}

// EnvoyInternalAddressMultiError is an error wrapping multiple validation
// errors returned by EnvoyInternalAddress.ValidateAll() if the designated
// constraints aren't met.
type EnvoyInternalAddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnvoyInternalAddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnvoyInternalAddressMultiError) AllErrors() []error { return m }

// EnvoyInternalAddressValidationError is the validation error returned by
// EnvoyInternalAddress.Validate if the designated constraints aren't met.
type EnvoyInternalAddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnvoyInternalAddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnvoyInternalAddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnvoyInternalAddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnvoyInternalAddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnvoyInternalAddressValidationError) ErrorName() string {
	return "EnvoyInternalAddressValidationError"
}

// Error satisfies the builtin error interface
func (e EnvoyInternalAddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnvoyInternalAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnvoyInternalAddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnvoyInternalAddressValidationError{}

// Validate checks the field values on SocketAddress with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SocketAddress) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketAddress with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SocketAddressMultiError, or
// nil if none found.
func (m *SocketAddress) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketAddress) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := SocketAddress_Protocol_name[int32(m.GetProtocol())]; !ok {
		err := SocketAddressValidationError{
			field:  "Protocol",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAddress()) < 1 {
		err := SocketAddressValidationError{
			field:  "Address",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ResolverName

	// no validation rules for Ipv4Compat

	switch m.PortSpecifier.(type) {

	case *SocketAddress_PortValue:

		if m.GetPortValue() > 65535 {
			err := SocketAddressValidationError{
				field:  "PortValue",
				reason: "value must be less than or equal to 65535",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *SocketAddress_NamedPort:
		// no validation rules for NamedPort

	default:
		err := SocketAddressValidationError{
			field:  "PortSpecifier",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return SocketAddressMultiError(errors)
	}
	return nil
}

// SocketAddressMultiError is an error wrapping multiple validation errors
// returned by SocketAddress.ValidateAll() if the designated constraints
// aren't met.
type SocketAddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketAddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketAddressMultiError) AllErrors() []error { return m }

// SocketAddressValidationError is the validation error returned by
// SocketAddress.Validate if the designated constraints aren't met.
type SocketAddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketAddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketAddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketAddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketAddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketAddressValidationError) ErrorName() string { return "SocketAddressValidationError" }

// Error satisfies the builtin error interface
func (e SocketAddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketAddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketAddressValidationError{}

// Validate checks the field values on TcpKeepalive with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TcpKeepalive) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TcpKeepalive with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TcpKeepaliveMultiError, or
// nil if none found.
func (m *TcpKeepalive) ValidateAll() error {
	return m.validate(true)
}

func (m *TcpKeepalive) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetKeepaliveProbes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpKeepaliveValidationError{
					field:  "KeepaliveProbes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpKeepaliveValidationError{
					field:  "KeepaliveProbes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKeepaliveProbes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpKeepaliveValidationError{
				field:  "KeepaliveProbes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetKeepaliveTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpKeepaliveValidationError{
					field:  "KeepaliveTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpKeepaliveValidationError{
					field:  "KeepaliveTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKeepaliveTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpKeepaliveValidationError{
				field:  "KeepaliveTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetKeepaliveInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpKeepaliveValidationError{
					field:  "KeepaliveInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpKeepaliveValidationError{
					field:  "KeepaliveInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKeepaliveInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpKeepaliveValidationError{
				field:  "KeepaliveInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TcpKeepaliveMultiError(errors)
	}
	return nil
}

// TcpKeepaliveMultiError is an error wrapping multiple validation errors
// returned by TcpKeepalive.ValidateAll() if the designated constraints aren't met.
type TcpKeepaliveMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TcpKeepaliveMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TcpKeepaliveMultiError) AllErrors() []error { return m }

// TcpKeepaliveValidationError is the validation error returned by
// TcpKeepalive.Validate if the designated constraints aren't met.
type TcpKeepaliveValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpKeepaliveValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpKeepaliveValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpKeepaliveValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpKeepaliveValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpKeepaliveValidationError) ErrorName() string { return "TcpKeepaliveValidationError" }

// Error satisfies the builtin error interface
func (e TcpKeepaliveValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpKeepalive.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpKeepaliveValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpKeepaliveValidationError{}

// Validate checks the field values on ExtraSourceAddress with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExtraSourceAddress) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExtraSourceAddress with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExtraSourceAddressMultiError, or nil if none found.
func (m *ExtraSourceAddress) ValidateAll() error {
	return m.validate(true)
}

func (m *ExtraSourceAddress) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAddress() == nil {
		err := ExtraSourceAddressValidationError{
			field:  "Address",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExtraSourceAddressValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExtraSourceAddressValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtraSourceAddressValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSocketOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExtraSourceAddressValidationError{
					field:  "SocketOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExtraSourceAddressValidationError{
					field:  "SocketOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSocketOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtraSourceAddressValidationError{
				field:  "SocketOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ExtraSourceAddressMultiError(errors)
	}
	return nil
}

// ExtraSourceAddressMultiError is an error wrapping multiple validation errors
// returned by ExtraSourceAddress.ValidateAll() if the designated constraints
// aren't met.
type ExtraSourceAddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExtraSourceAddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExtraSourceAddressMultiError) AllErrors() []error { return m }

// ExtraSourceAddressValidationError is the validation error returned by
// ExtraSourceAddress.Validate if the designated constraints aren't met.
type ExtraSourceAddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtraSourceAddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtraSourceAddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtraSourceAddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtraSourceAddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtraSourceAddressValidationError) ErrorName() string {
	return "ExtraSourceAddressValidationError"
}

// Error satisfies the builtin error interface
func (e ExtraSourceAddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtraSourceAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtraSourceAddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtraSourceAddressValidationError{}

// Validate checks the field values on BindConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BindConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BindConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BindConfigMultiError, or
// nil if none found.
func (m *BindConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *BindConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSourceAddress() == nil {
		err := BindConfigValidationError{
			field:  "SourceAddress",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSourceAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BindConfigValidationError{
					field:  "SourceAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BindConfigValidationError{
					field:  "SourceAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSourceAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BindConfigValidationError{
				field:  "SourceAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFreebind()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BindConfigValidationError{
					field:  "Freebind",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BindConfigValidationError{
					field:  "Freebind",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFreebind()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BindConfigValidationError{
				field:  "Freebind",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSocketOptions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BindConfigValidationError{
						field:  fmt.Sprintf("SocketOptions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BindConfigValidationError{
						field:  fmt.Sprintf("SocketOptions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BindConfigValidationError{
					field:  fmt.Sprintf("SocketOptions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetExtraSourceAddresses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BindConfigValidationError{
						field:  fmt.Sprintf("ExtraSourceAddresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BindConfigValidationError{
						field:  fmt.Sprintf("ExtraSourceAddresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BindConfigValidationError{
					field:  fmt.Sprintf("ExtraSourceAddresses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetAdditionalSourceAddresses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BindConfigValidationError{
						field:  fmt.Sprintf("AdditionalSourceAddresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BindConfigValidationError{
						field:  fmt.Sprintf("AdditionalSourceAddresses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BindConfigValidationError{
					field:  fmt.Sprintf("AdditionalSourceAddresses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BindConfigMultiError(errors)
	}
	return nil
}

// BindConfigMultiError is an error wrapping multiple validation errors
// returned by BindConfig.ValidateAll() if the designated constraints aren't met.
type BindConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BindConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BindConfigMultiError) AllErrors() []error { return m }

// BindConfigValidationError is the validation error returned by
// BindConfig.Validate if the designated constraints aren't met.
type BindConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BindConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BindConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BindConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BindConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BindConfigValidationError) ErrorName() string { return "BindConfigValidationError" }

// Error satisfies the builtin error interface
func (e BindConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBindConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BindConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BindConfigValidationError{}

// Validate checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Address) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AddressMultiError, or nil if none found.
func (m *Address) ValidateAll() error {
	return m.validate(true)
}

func (m *Address) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Address.(type) {

	case *Address_SocketAddress:

		if all {
			switch v := interface{}(m.GetSocketAddress()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "SocketAddress",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "SocketAddress",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSocketAddress()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddressValidationError{
					field:  "SocketAddress",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Address_Pipe:

		if all {
			switch v := interface{}(m.GetPipe()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "Pipe",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "Pipe",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPipe()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddressValidationError{
					field:  "Pipe",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Address_EnvoyInternalAddress:

		if all {
			switch v := interface{}(m.GetEnvoyInternalAddress()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "EnvoyInternalAddress",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AddressValidationError{
						field:  "EnvoyInternalAddress",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEnvoyInternalAddress()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddressValidationError{
					field:  "EnvoyInternalAddress",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := AddressValidationError{
			field:  "Address",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return AddressMultiError(errors)
	}
	return nil
}

// AddressMultiError is an error wrapping multiple validation errors returned
// by Address.ValidateAll() if the designated constraints aren't met.
type AddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressMultiError) AllErrors() []error { return m }

// AddressValidationError is the validation error returned by Address.Validate
// if the designated constraints aren't met.
type AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressValidationError) ErrorName() string { return "AddressValidationError" }

// Error satisfies the builtin error interface
func (e AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressValidationError{}

// Validate checks the field values on CidrRange with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CidrRange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CidrRange with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CidrRangeMultiError, or nil
// if none found.
func (m *CidrRange) ValidateAll() error {
	return m.validate(true)
}

func (m *CidrRange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAddressPrefix()) < 1 {
		err := CidrRangeValidationError{
			field:  "AddressPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if wrapper := m.GetPrefixLen(); wrapper != nil {

		if wrapper.GetValue() > 128 {
			err := CidrRangeValidationError{
				field:  "PrefixLen",
				reason: "value must be less than or equal to 128",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return CidrRangeMultiError(errors)
	}
	return nil
}

// CidrRangeMultiError is an error wrapping multiple validation errors returned
// by CidrRange.ValidateAll() if the designated constraints aren't met.
type CidrRangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CidrRangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CidrRangeMultiError) AllErrors() []error { return m }

// CidrRangeValidationError is the validation error returned by
// CidrRange.Validate if the designated constraints aren't met.
type CidrRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CidrRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CidrRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CidrRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CidrRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CidrRangeValidationError) ErrorName() string { return "CidrRangeValidationError" }

// Error satisfies the builtin error interface
func (e CidrRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCidrRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CidrRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CidrRangeValidationError{}
