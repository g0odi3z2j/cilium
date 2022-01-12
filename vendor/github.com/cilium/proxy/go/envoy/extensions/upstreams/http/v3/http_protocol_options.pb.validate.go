// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/upstreams/http/v3/http_protocol_options.proto

package httpv3

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

// Validate checks the field values on HttpProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HttpProtocolOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HttpProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HttpProtocolOptionsMultiError, or nil if none found.
func (m *HttpProtocolOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpProtocolOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommonHttpProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptionsValidationError{
					field:  "CommonHttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptionsValidationError{
					field:  "CommonHttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonHttpProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptionsValidationError{
				field:  "CommonHttpProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpstreamHttpProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptionsValidationError{
					field:  "UpstreamHttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptionsValidationError{
					field:  "UpstreamHttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpstreamHttpProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptionsValidationError{
				field:  "UpstreamHttpProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.UpstreamProtocolOptions.(type) {

	case *HttpProtocolOptions_ExplicitHttpConfig_:

		if all {
			switch v := interface{}(m.GetExplicitHttpConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HttpProtocolOptionsValidationError{
						field:  "ExplicitHttpConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HttpProtocolOptionsValidationError{
						field:  "ExplicitHttpConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetExplicitHttpConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptionsValidationError{
					field:  "ExplicitHttpConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpProtocolOptions_UseDownstreamProtocolConfig:

		if all {
			switch v := interface{}(m.GetUseDownstreamProtocolConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HttpProtocolOptionsValidationError{
						field:  "UseDownstreamProtocolConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HttpProtocolOptionsValidationError{
						field:  "UseDownstreamProtocolConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUseDownstreamProtocolConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptionsValidationError{
					field:  "UseDownstreamProtocolConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpProtocolOptions_AutoConfig:

		if all {
			switch v := interface{}(m.GetAutoConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HttpProtocolOptionsValidationError{
						field:  "AutoConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HttpProtocolOptionsValidationError{
						field:  "AutoConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAutoConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptionsValidationError{
					field:  "AutoConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := HttpProtocolOptionsValidationError{
			field:  "UpstreamProtocolOptions",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return HttpProtocolOptionsMultiError(errors)
	}
	return nil
}

// HttpProtocolOptionsMultiError is an error wrapping multiple validation
// errors returned by HttpProtocolOptions.ValidateAll() if the designated
// constraints aren't met.
type HttpProtocolOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpProtocolOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpProtocolOptionsMultiError) AllErrors() []error { return m }

// HttpProtocolOptionsValidationError is the validation error returned by
// HttpProtocolOptions.Validate if the designated constraints aren't met.
type HttpProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptionsValidationError) ErrorName() string {
	return "HttpProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptionsValidationError{}

// Validate checks the field values on HttpProtocolOptions_ExplicitHttpConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *HttpProtocolOptions_ExplicitHttpConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// HttpProtocolOptions_ExplicitHttpConfig with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// HttpProtocolOptions_ExplicitHttpConfigMultiError, or nil if none found.
func (m *HttpProtocolOptions_ExplicitHttpConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpProtocolOptions_ExplicitHttpConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.ProtocolConfig.(type) {

	case *HttpProtocolOptions_ExplicitHttpConfig_HttpProtocolOptions:

		if all {
			switch v := interface{}(m.GetHttpProtocolOptions()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HttpProtocolOptions_ExplicitHttpConfigValidationError{
						field:  "HttpProtocolOptions",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HttpProtocolOptions_ExplicitHttpConfigValidationError{
						field:  "HttpProtocolOptions",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHttpProtocolOptions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptions_ExplicitHttpConfigValidationError{
					field:  "HttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions:

		if all {
			switch v := interface{}(m.GetHttp2ProtocolOptions()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HttpProtocolOptions_ExplicitHttpConfigValidationError{
						field:  "Http2ProtocolOptions",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HttpProtocolOptions_ExplicitHttpConfigValidationError{
						field:  "Http2ProtocolOptions",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHttp2ProtocolOptions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptions_ExplicitHttpConfigValidationError{
					field:  "Http2ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HttpProtocolOptions_ExplicitHttpConfig_Http3ProtocolOptions:

		if all {
			switch v := interface{}(m.GetHttp3ProtocolOptions()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HttpProtocolOptions_ExplicitHttpConfigValidationError{
						field:  "Http3ProtocolOptions",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HttpProtocolOptions_ExplicitHttpConfigValidationError{
						field:  "Http3ProtocolOptions",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHttp3ProtocolOptions()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HttpProtocolOptions_ExplicitHttpConfigValidationError{
					field:  "Http3ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		err := HttpProtocolOptions_ExplicitHttpConfigValidationError{
			field:  "ProtocolConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return HttpProtocolOptions_ExplicitHttpConfigMultiError(errors)
	}
	return nil
}

// HttpProtocolOptions_ExplicitHttpConfigMultiError is an error wrapping
// multiple validation errors returned by
// HttpProtocolOptions_ExplicitHttpConfig.ValidateAll() if the designated
// constraints aren't met.
type HttpProtocolOptions_ExplicitHttpConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpProtocolOptions_ExplicitHttpConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpProtocolOptions_ExplicitHttpConfigMultiError) AllErrors() []error { return m }

// HttpProtocolOptions_ExplicitHttpConfigValidationError is the validation
// error returned by HttpProtocolOptions_ExplicitHttpConfig.Validate if the
// designated constraints aren't met.
type HttpProtocolOptions_ExplicitHttpConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) ErrorName() string {
	return "HttpProtocolOptions_ExplicitHttpConfigValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptions_ExplicitHttpConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions_ExplicitHttpConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptions_ExplicitHttpConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptions_ExplicitHttpConfigValidationError{}

// Validate checks the field values on
// HttpProtocolOptions_UseDownstreamHttpConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HttpProtocolOptions_UseDownstreamHttpConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// HttpProtocolOptions_UseDownstreamHttpConfig with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// HttpProtocolOptions_UseDownstreamHttpConfigMultiError, or nil if none found.
func (m *HttpProtocolOptions_UseDownstreamHttpConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpProtocolOptions_UseDownstreamHttpConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHttpProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
					field:  "HttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
					field:  "HttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttpProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
				field:  "HttpProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHttp2ProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
					field:  "Http2ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
					field:  "Http2ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttp2ProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
				field:  "Http2ProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHttp3ProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
					field:  "Http3ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
					field:  "Http3ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttp3ProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_UseDownstreamHttpConfigValidationError{
				field:  "Http3ProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HttpProtocolOptions_UseDownstreamHttpConfigMultiError(errors)
	}
	return nil
}

// HttpProtocolOptions_UseDownstreamHttpConfigMultiError is an error wrapping
// multiple validation errors returned by
// HttpProtocolOptions_UseDownstreamHttpConfig.ValidateAll() if the designated
// constraints aren't met.
type HttpProtocolOptions_UseDownstreamHttpConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpProtocolOptions_UseDownstreamHttpConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpProtocolOptions_UseDownstreamHttpConfigMultiError) AllErrors() []error { return m }

// HttpProtocolOptions_UseDownstreamHttpConfigValidationError is the validation
// error returned by HttpProtocolOptions_UseDownstreamHttpConfig.Validate if
// the designated constraints aren't met.
type HttpProtocolOptions_UseDownstreamHttpConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) ErrorName() string {
	return "HttpProtocolOptions_UseDownstreamHttpConfigValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptions_UseDownstreamHttpConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions_UseDownstreamHttpConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptions_UseDownstreamHttpConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptions_UseDownstreamHttpConfigValidationError{}

// Validate checks the field values on HttpProtocolOptions_AutoHttpConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *HttpProtocolOptions_AutoHttpConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HttpProtocolOptions_AutoHttpConfig
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// HttpProtocolOptions_AutoHttpConfigMultiError, or nil if none found.
func (m *HttpProtocolOptions_AutoHttpConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpProtocolOptions_AutoHttpConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHttpProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptions_AutoHttpConfigValidationError{
					field:  "HttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptions_AutoHttpConfigValidationError{
					field:  "HttpProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttpProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_AutoHttpConfigValidationError{
				field:  "HttpProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHttp2ProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptions_AutoHttpConfigValidationError{
					field:  "Http2ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptions_AutoHttpConfigValidationError{
					field:  "Http2ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttp2ProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_AutoHttpConfigValidationError{
				field:  "Http2ProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHttp3ProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptions_AutoHttpConfigValidationError{
					field:  "Http3ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptions_AutoHttpConfigValidationError{
					field:  "Http3ProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHttp3ProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_AutoHttpConfigValidationError{
				field:  "Http3ProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAlternateProtocolsCacheOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpProtocolOptions_AutoHttpConfigValidationError{
					field:  "AlternateProtocolsCacheOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpProtocolOptions_AutoHttpConfigValidationError{
					field:  "AlternateProtocolsCacheOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlternateProtocolsCacheOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpProtocolOptions_AutoHttpConfigValidationError{
				field:  "AlternateProtocolsCacheOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HttpProtocolOptions_AutoHttpConfigMultiError(errors)
	}
	return nil
}

// HttpProtocolOptions_AutoHttpConfigMultiError is an error wrapping multiple
// validation errors returned by
// HttpProtocolOptions_AutoHttpConfig.ValidateAll() if the designated
// constraints aren't met.
type HttpProtocolOptions_AutoHttpConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpProtocolOptions_AutoHttpConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpProtocolOptions_AutoHttpConfigMultiError) AllErrors() []error { return m }

// HttpProtocolOptions_AutoHttpConfigValidationError is the validation error
// returned by HttpProtocolOptions_AutoHttpConfig.Validate if the designated
// constraints aren't met.
type HttpProtocolOptions_AutoHttpConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpProtocolOptions_AutoHttpConfigValidationError) ErrorName() string {
	return "HttpProtocolOptions_AutoHttpConfigValidationError"
}

// Error satisfies the builtin error interface
func (e HttpProtocolOptions_AutoHttpConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpProtocolOptions_AutoHttpConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpProtocolOptions_AutoHttpConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpProtocolOptions_AutoHttpConfigValidationError{}
