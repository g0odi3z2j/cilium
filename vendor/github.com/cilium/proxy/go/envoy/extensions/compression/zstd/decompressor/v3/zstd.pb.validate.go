// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/compression/zstd/decompressor/v3/zstd.proto

package decompressorv3

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

// Validate checks the field values on Zstd with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Zstd) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Zstd with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ZstdMultiError, or nil if none found.
func (m *Zstd) ValidateAll() error {
	return m.validate(true)
}

func (m *Zstd) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDictionaries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ZstdValidationError{
						field:  fmt.Sprintf("Dictionaries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ZstdValidationError{
						field:  fmt.Sprintf("Dictionaries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ZstdValidationError{
					field:  fmt.Sprintf("Dictionaries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if wrapper := m.GetChunkSize(); wrapper != nil {

		if val := wrapper.GetValue(); val < 4096 || val > 65536 {
			err := ZstdValidationError{
				field:  "ChunkSize",
				reason: "value must be inside range [4096, 65536]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ZstdMultiError(errors)
	}

	return nil
}

// ZstdMultiError is an error wrapping multiple validation errors returned by
// Zstd.ValidateAll() if the designated constraints aren't met.
type ZstdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ZstdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ZstdMultiError) AllErrors() []error { return m }

// ZstdValidationError is the validation error returned by Zstd.Validate if the
// designated constraints aren't met.
type ZstdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ZstdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ZstdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ZstdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ZstdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ZstdValidationError) ErrorName() string { return "ZstdValidationError" }

// Error satisfies the builtin error interface
func (e ZstdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sZstd.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ZstdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ZstdValidationError{}
