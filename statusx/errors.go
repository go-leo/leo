package statusx

import (
	"google.golang.org/grpc/codes"
)

// OK Not an error; returned on success.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: OK
// HTTP Mapping: 200 OK
func OK(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.OK, msg, args...)
	}
	return NewError(codes.OK, msg)
}

var kOK = NewError(codes.OK, "")

func IsOK(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kOK)
}

// Failed unlike ErrUnknown error, it just means business logic failed.
//
// For example, if client want sign up, but username already exists,
// it should return Failed("username already exist").
//
// GRPC Mapping: Unknown
// HTTP Mapping: 200 OK
func Failed(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(FailedCode, msg, args...)
	}
	return NewError(FailedCode, msg)
}

var kFailed = NewError(FailedCode, "")

func IsFailed(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kFailed)
}

// ErrCanceled The operation was cancelled, typically by the caller.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: Canceled
// HTTP Mapping: 499 Client Closed Request
func ErrCanceled(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.Canceled, msg, args...)
	}
	return NewError(codes.Canceled, msg)
}

var kErrCanceled = NewError(codes.Canceled, "")

func IsCanceled(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrCanceled)
}

// ErrUnknown error.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: Unknown
// HTTP Mapping: 500 ErrInternal Server Error
func ErrUnknown(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.Unknown, msg, args...)
	}
	return NewError(codes.Unknown, msg)
}

var kErrUnknown = NewError(codes.Unknown, "")

func IsUnknown(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrUnknown)
}

// ErrInvalidArgument The client specified an invalid argument.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: InvalidArgument
// HTTP Mapping: 400 Bad Request
func ErrInvalidArgument(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.InvalidArgument, msg, args...)
	}
	return NewError(codes.InvalidArgument, msg)
}

var kErrInvalidArgument = NewError(codes.InvalidArgument, "")

func IsInvalidArgument(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrInvalidArgument)
}

// ErrDeadlineExceeded The deadline expired before the operation could complete.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: DeadlineExceeded
// HTTP Mapping: 504 Gateway Timeout
func ErrDeadlineExceeded(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.DeadlineExceeded, msg, args...)
	}
	return NewError(codes.DeadlineExceeded, msg)
}

var kErrDeadlineExceeded = NewError(codes.DeadlineExceeded, "")

func IsDeadlineExceeded(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrDeadlineExceeded)
}

// ErrNotFound Some requested entity (e.g., file or directory) was not found.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: NotFound
// HTTP Mapping: 404 Not Found
func ErrNotFound(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.NotFound, msg, args...)
	}
	return NewError(codes.NotFound, msg)
}

var kErrNotFound = NewError(codes.NotFound, "")

func IsNotFound(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrNotFound)
}

// ErrAlreadyExists The entity that a client attempted to create (e.g., file or directory)
// already exists.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: AlreadyExists
// HTTP Mapping: 409 Conflict
func ErrAlreadyExists(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.AlreadyExists, msg, args...)
	}
	return NewError(codes.AlreadyExists, msg)
}

var kErrAlreadyExists = NewError(codes.AlreadyExists, "")

func IsAlreadyExists(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrAlreadyExists)
}

// ErrPermissionDenied The caller does not have permission to execute the specified
// operation.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: PermissionDenied
// HTTP Mapping: 403 Forbidden
func ErrPermissionDenied(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.PermissionDenied, msg, args...)
	}
	return NewError(codes.PermissionDenied, msg)
}

var kErrPermissionDenied = NewError(codes.PermissionDenied, "")

func IsPermissionDenied(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrPermissionDenied)
}

// ErrResourceExhausted Some resource has been exhausted, perhaps a per-user quota, or
// perhaps the entire file system is out of space.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: ResourceExhausted
// HTTP Mapping: 429 Too Many Requests
func ErrResourceExhausted(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.ResourceExhausted, msg, args...)
	}
	return NewError(codes.ResourceExhausted, msg)
}

var kErrResourceExhausted = NewError(codes.ResourceExhausted, "")

func IsResourceExhausted(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrResourceExhausted)
}

// ErrFailedPrecondition The operation was rejected because the system is not in a state
// required for the operation's execution.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: FailedPrecondition
// HTTP Mapping: 400 Bad Request
func ErrFailedPrecondition(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.FailedPrecondition, msg, args...)
	}
	return NewError(codes.FailedPrecondition, msg)
}

var kErrFailedPrecondition = NewError(codes.FailedPrecondition, "")

func IsFailedPrecondition(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrFailedPrecondition)
}

// ErrAborted The operation was aborted, typically due to a concurrency issue such as
// a sequencer check failure or transaction abort.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: Aborted
// HTTP Mapping: 409 Conflict
func ErrAborted(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.Aborted, msg, args...)
	}
	return NewError(codes.Aborted, msg)
}

var kErrAborted = NewError(codes.Aborted, "")

func IsAborted(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrAborted)
}

// ErrOutOfRange The operation was attempted past the valid range.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: OutOfRange
// HTTP Mapping: 400 Bad Request
func ErrOutOfRange(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.OutOfRange, msg, args...)
	}
	return NewError(codes.OutOfRange, msg)
}

var kErrOutOfRange = NewError(codes.OutOfRange, "")

func IsOutOfRange(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrOutOfRange)
}

// ErrUnimplemented The operation is not implemented or is not supported/enabled in this
// service.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: Unimplemented
// HTTP Mapping: 501 Not Implemented
func ErrUnimplemented(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.Unimplemented, msg, args...)
	}
	return NewError(codes.Unimplemented, msg)
}

var kErrUnimplemented = NewError(codes.Unimplemented, "")

func IsUnimplemented(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrUnimplemented)
}

// ErrInternal internal errors.  This means that some invariants expected by the
// underlying system have been broken.  This error code is reserved
// for serious errors.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: Internal
// HTTP Mapping: 500 ErrInternal Server Error
func ErrInternal(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.Internal, msg, args...)
	}
	return NewError(codes.Internal, msg)
}

var kErrInternal = NewError(codes.Internal, "")

func IsInternal(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrInternal)
}

// ErrUnavailable The service is currently unavailable.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: Unavailable
// HTTP Mapping: 503 Service ErrUnavailable
func ErrUnavailable(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.Unavailable, msg, args...)
	}
	return NewError(codes.Unavailable, msg)
}

var kErrUnavailable = NewError(codes.Unavailable, "")

func IsUnavailable(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrUnavailable)
}

// ErrDataLoss Unrecoverable data loss or corruption.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: DataLoss
// HTTP Mapping: 500 ErrInternal Server Error
func ErrDataLoss(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.DataLoss, msg, args...)
	}
	return NewError(codes.DataLoss, msg)
}

var kErrDataLoss = NewError(codes.DataLoss, "")

func IsDataLoss(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrDataLoss)
}

// ErrUnauthenticated The request does not have valid authentication credentials for the
// operation.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: Unauthenticated
// HTTP Mapping: 401 Unauthorized
func ErrUnauthenticated(msg string, args ...any) *Error {
	if len(args) > 0 {
		return NewErrorf(codes.Unauthenticated, msg, args...)
	}
	return NewError(codes.Unauthenticated, msg)
}

var kErrUnauthenticated = NewError(codes.Unauthenticated, "")

func IsUnauthenticated(err error) (*Error, bool) {
	e, ok := FromError(err)
	if !ok {
		return nil, false
	}
	return e, e.Equals(kErrUnauthenticated)
}
