// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cqrs/v1/cqrs.proto

package cqrs

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

// Validate checks the field values on QueryRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueryRequestMultiError, or
// nil if none found.
func (m *QueryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return QueryRequestMultiError(errors)
	}

	return nil
}

// QueryRequestMultiError is an error wrapping multiple validation errors
// returned by QueryRequest.ValidateAll() if the designated constraints aren't met.
type QueryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryRequestMultiError) AllErrors() []error { return m }

// QueryRequestValidationError is the validation error returned by
// QueryRequest.Validate if the designated constraints aren't met.
type QueryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryRequestValidationError) ErrorName() string { return "QueryRequestValidationError" }

// Error satisfies the builtin error interface
func (e QueryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryRequestValidationError{}

// Validate checks the field values on QueryReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueryReplyMultiError, or
// nil if none found.
func (m *QueryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return QueryReplyMultiError(errors)
	}

	return nil
}

// QueryReplyMultiError is an error wrapping multiple validation errors
// returned by QueryReply.ValidateAll() if the designated constraints aren't met.
type QueryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryReplyMultiError) AllErrors() []error { return m }

// QueryReplyValidationError is the validation error returned by
// QueryReply.Validate if the designated constraints aren't met.
type QueryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryReplyValidationError) ErrorName() string { return "QueryReplyValidationError" }

// Error satisfies the builtin error interface
func (e QueryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryReplyValidationError{}

// Validate checks the field values on CommandRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommandRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommandRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommandRequestMultiError,
// or nil if none found.
func (m *CommandRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CommandRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return CommandRequestMultiError(errors)
	}

	return nil
}

// CommandRequestMultiError is an error wrapping multiple validation errors
// returned by CommandRequest.ValidateAll() if the designated constraints
// aren't met.
type CommandRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommandRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommandRequestMultiError) AllErrors() []error { return m }

// CommandRequestValidationError is the validation error returned by
// CommandRequest.Validate if the designated constraints aren't met.
type CommandRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommandRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommandRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommandRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommandRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommandRequestValidationError) ErrorName() string { return "CommandRequestValidationError" }

// Error satisfies the builtin error interface
func (e CommandRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommandRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommandRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommandRequestValidationError{}

// Validate checks the field values on CommandReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommandReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommandReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommandReplyMultiError, or
// nil if none found.
func (m *CommandReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CommandReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CommandReplyMultiError(errors)
	}

	return nil
}

// CommandReplyMultiError is an error wrapping multiple validation errors
// returned by CommandReply.ValidateAll() if the designated constraints aren't met.
type CommandReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommandReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommandReplyMultiError) AllErrors() []error { return m }

// CommandReplyValidationError is the validation error returned by
// CommandReply.Validate if the designated constraints aren't met.
type CommandReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommandReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommandReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommandReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommandReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommandReplyValidationError) ErrorName() string { return "CommandReplyValidationError" }

// Error satisfies the builtin error interface
func (e CommandReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommandReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommandReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommandReplyValidationError{}

// Validate checks the field values on QueryOneOfRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QueryOneOfRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryOneOfRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryOneOfRequestMultiError, or nil if none found.
func (m *QueryOneOfRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryOneOfRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return QueryOneOfRequestMultiError(errors)
	}

	return nil
}

// QueryOneOfRequestMultiError is an error wrapping multiple validation errors
// returned by QueryOneOfRequest.ValidateAll() if the designated constraints
// aren't met.
type QueryOneOfRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryOneOfRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryOneOfRequestMultiError) AllErrors() []error { return m }

// QueryOneOfRequestValidationError is the validation error returned by
// QueryOneOfRequest.Validate if the designated constraints aren't met.
type QueryOneOfRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryOneOfRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryOneOfRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryOneOfRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryOneOfRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryOneOfRequestValidationError) ErrorName() string {
	return "QueryOneOfRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QueryOneOfRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryOneOfRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryOneOfRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryOneOfRequestValidationError{}

// Validate checks the field values on QueryOneOfReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QueryOneOfReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryOneOfReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryOneOfReplyMultiError, or nil if none found.
func (m *QueryOneOfReply) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryOneOfReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Data.(type) {
	case *QueryOneOfReply_Name:
		if v == nil {
			err := QueryOneOfReplyValidationError{
				field:  "Data",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Name
	case *QueryOneOfReply_Id:
		if v == nil {
			err := QueryOneOfReplyValidationError{
				field:  "Data",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Id
	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return QueryOneOfReplyMultiError(errors)
	}

	return nil
}

// QueryOneOfReplyMultiError is an error wrapping multiple validation errors
// returned by QueryOneOfReply.ValidateAll() if the designated constraints
// aren't met.
type QueryOneOfReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryOneOfReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryOneOfReplyMultiError) AllErrors() []error { return m }

// QueryOneOfReplyValidationError is the validation error returned by
// QueryOneOfReply.Validate if the designated constraints aren't met.
type QueryOneOfReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryOneOfReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryOneOfReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryOneOfReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryOneOfReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryOneOfReplyValidationError) ErrorName() string { return "QueryOneOfReplyValidationError" }

// Error satisfies the builtin error interface
func (e QueryOneOfReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryOneOfReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryOneOfReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryOneOfReplyValidationError{}
