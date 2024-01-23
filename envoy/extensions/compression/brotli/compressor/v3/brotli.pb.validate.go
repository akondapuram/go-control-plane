//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/compression/brotli/compressor/v3/brotli.proto

package compressorv3

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

// Validate checks the field values on Brotli with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Brotli) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Brotli with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BrotliMultiError, or nil if none found.
func (m *Brotli) ValidateAll() error {
	return m.validate(true)
}

func (m *Brotli) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if wrapper := m.GetQuality(); wrapper != nil {

		if wrapper.GetValue() > 11 {
			err := BrotliValidationError{
				field:  "Quality",
				reason: "value must be less than or equal to 11",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if _, ok := Brotli_EncoderMode_name[int32(m.GetEncoderMode())]; !ok {
		err := BrotliValidationError{
			field:  "EncoderMode",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if wrapper := m.GetWindowBits(); wrapper != nil {

		if val := wrapper.GetValue(); val < 10 || val > 24 {
			err := BrotliValidationError{
				field:  "WindowBits",
				reason: "value must be inside range [10, 24]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if wrapper := m.GetInputBlockBits(); wrapper != nil {

		if val := wrapper.GetValue(); val < 16 || val > 24 {
			err := BrotliValidationError{
				field:  "InputBlockBits",
				reason: "value must be inside range [16, 24]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if wrapper := m.GetChunkSize(); wrapper != nil {

		if val := wrapper.GetValue(); val < 4096 || val > 65536 {
			err := BrotliValidationError{
				field:  "ChunkSize",
				reason: "value must be inside range [4096, 65536]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for DisableLiteralContextModeling

	if len(errors) > 0 {
		return BrotliMultiError(errors)
	}

	return nil
}

// BrotliMultiError is an error wrapping multiple validation errors returned by
// Brotli.ValidateAll() if the designated constraints aren't met.
type BrotliMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BrotliMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BrotliMultiError) AllErrors() []error { return m }

// BrotliValidationError is the validation error returned by Brotli.Validate if
// the designated constraints aren't met.
type BrotliValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BrotliValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BrotliValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BrotliValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BrotliValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BrotliValidationError) ErrorName() string { return "BrotliValidationError" }

// Error satisfies the builtin error interface
func (e BrotliValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBrotli.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BrotliValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BrotliValidationError{}
