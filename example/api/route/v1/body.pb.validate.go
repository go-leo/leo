// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: route/v1/body.proto

package route

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

// Validate checks the field values on BodyRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BodyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BodyRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BodyRequestMultiError, or
// nil if none found.
func (m *BodyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BodyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BodyRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BodyRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BodyRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BodyRequestMultiError(errors)
	}

	return nil
}

// BodyRequestMultiError is an error wrapping multiple validation errors
// returned by BodyRequest.ValidateAll() if the designated constraints aren't met.
type BodyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BodyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BodyRequestMultiError) AllErrors() []error { return m }

// BodyRequestValidationError is the validation error returned by
// BodyRequest.Validate if the designated constraints aren't met.
type BodyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BodyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BodyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BodyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BodyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BodyRequestValidationError) ErrorName() string { return "BodyRequestValidationError" }

// Error satisfies the builtin error interface
func (e BodyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBodyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BodyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BodyRequestValidationError{}

// Validate checks the field values on HttpBodyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *HttpBodyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HttpBodyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HttpBodyRequestMultiError, or nil if none found.
func (m *HttpBodyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpBodyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBody()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpBodyRequestValidationError{
					field:  "Body",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpBodyRequestValidationError{
					field:  "Body",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpBodyRequestValidationError{
				field:  "Body",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HttpBodyRequestMultiError(errors)
	}

	return nil
}

// HttpBodyRequestMultiError is an error wrapping multiple validation errors
// returned by HttpBodyRequest.ValidateAll() if the designated constraints
// aren't met.
type HttpBodyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpBodyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpBodyRequestMultiError) AllErrors() []error { return m }

// HttpBodyRequestValidationError is the validation error returned by
// HttpBodyRequest.Validate if the designated constraints aren't met.
type HttpBodyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpBodyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpBodyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpBodyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpBodyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpBodyRequestValidationError) ErrorName() string { return "HttpBodyRequestValidationError" }

// Error satisfies the builtin error interface
func (e HttpBodyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpBodyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpBodyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpBodyRequestValidationError{}

// Validate checks the field values on BodyRequest_User with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BodyRequest_User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BodyRequest_User with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BodyRequest_UserMultiError, or nil if none found.
func (m *BodyRequest_User) ValidateAll() error {
	return m.validate(true)
}

func (m *BodyRequest_User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Email

	// no validation rules for Phone

	// no validation rules for Address

	if len(errors) > 0 {
		return BodyRequest_UserMultiError(errors)
	}

	return nil
}

// BodyRequest_UserMultiError is an error wrapping multiple validation errors
// returned by BodyRequest_User.ValidateAll() if the designated constraints
// aren't met.
type BodyRequest_UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BodyRequest_UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BodyRequest_UserMultiError) AllErrors() []error { return m }

// BodyRequest_UserValidationError is the validation error returned by
// BodyRequest_User.Validate if the designated constraints aren't met.
type BodyRequest_UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BodyRequest_UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BodyRequest_UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BodyRequest_UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BodyRequest_UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BodyRequest_UserValidationError) ErrorName() string { return "BodyRequest_UserValidationError" }

// Error satisfies the builtin error interface
func (e BodyRequest_UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBodyRequest_User.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BodyRequest_UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BodyRequest_UserValidationError{}
