package statusx

import (
	"google.golang.org/grpc/codes"
)

// OK Not an error; returned on success.
//
// GRPC Mapping: OK
// HTTP Mapping: 200 OK
var OK = NewError(codes.OK, "")

// Failed unlike ErrUnknown error, it just means business logic failed.
//
// For example, if client want sign up, but username already exists,
// it should return Failed("username already exist").
//
// GRPC Mapping: Unknown
// HTTP Mapping: 200 OK
var Failed = NewError(FailedCode, "")

// ErrCanceled The operation was cancelled, typically by the caller.
//
// For example, client Application cancelled the request.
//
// GRPC Mapping: Canceled
// HTTP Mapping: 499 Client Closed Request
var ErrCanceled = NewError(codes.Canceled, "")

// ErrUnknown error.  For example, this error may be returned when
// a `Status` value received from another address space belongs to
// an error space that is not known in this address space.  Also
// errors raised by APIs that do not return enough error information
// may be converted to this error.
//
// For example, server side application throws an exception (or does
// something other than returning a Status code to terminate an RPC)
//
// # Another example, error parsing returned status
//
// GRPC Mapping: Unknown
// HTTP Mapping: 500 ErrInternal Server Error
var ErrUnknown = NewError(codes.Unknown, "")

// ErrInvalidArgument The client specified an invalid argument.  Note that this differs
// from ErrFailedPrecondition.  ErrInvalidArgument indicates arguments
// that are problematic regardless of the state of the system
// (e.g., a malformed file name).
//
// GRPC Mapping: InvalidArgument
// HTTP Mapping: 400 Bad Request
var ErrInvalidArgument = NewError(codes.InvalidArgument, "")

// ErrDeadlineExceeded The deadline expired before the operation could complete. For operations
// that change the state of the system, this error may be returned
// even if the operation has completed successfully.  For example, a
// successful response from a server could have been delayed long
// enough for the deadline to expire.
//
// For example, deadline expires before server returns status.
//
// Another example, no response received before Deadline expires. This may occur either when the
// client is unable to send the request to the server or when the server fails to respond in time.
//
// GRPC Mapping: DeadlineExceeded
// HTTP Mapping: 504 Gateway Timeout
var ErrDeadlineExceeded = NewError(codes.DeadlineExceeded, "")

// ErrNotFound Some requested entity (e.g., file or directory) was not found.
//
// Note to server developers: if a request is denied for an entire class
// of users, such as gradual feature rollout or undocumented allowlist,
// ErrNotFound may be used. If a request is denied for some users within
// a class of users, such as user-based access control, ErrPermissionDenied
// must be used.
//
// GRPC Mapping: NotFound
// HTTP Mapping: 404 Not Found
var ErrNotFound = NewError(codes.NotFound, "")

// ErrAlreadyExists The entity that a client attempted to create (e.g., file or directory)
// already exists.
//
// GRPC Mapping: AlreadyExists
// HTTP Mapping: 409 Conflict
var ErrAlreadyExists = NewError(codes.AlreadyExists, "")

// ErrPermissionDenied The caller does not have permission to execute the specified
// operation. ErrPermissionDenied must not be used for rejections
// caused by exhausting some resource (use ErrResourceExhausted
// instead for those errors). ErrPermissionDenied must not be
// used if the caller can not be identified (use ErrUnauthenticated
// instead for those errors). This error code does not imply the
// request is valid or the requested entity exists or satisfies
// other pre-conditions.
//
// GRPC Mapping: PermissionDenied
// HTTP Mapping: 403 Forbidden
var ErrPermissionDenied = NewError(codes.PermissionDenied, "")

// ErrResourceExhausted Some resource has been exhausted, perhaps a per-user quota, or
// perhaps the entire file system is out of space.
//
// Example 1, server temporarily out of resources (e.g., Flow-control resource limits reached)
//
// Example 2, client does not have enough memory to hold the server response.
//
// Example 3, Sent or received message was larger than configured limit.
//
// GRPC Mapping: ResourceExhausted
// HTTP Mapping: 429 Too Many Requests
var ErrResourceExhausted = NewError(codes.ResourceExhausted, "")

// ErrFailedPrecondition The operation was rejected because the system is not in a state
// required for the operation's execution.  For example, the directory
// to be deleted is non-empty, an rmdir operation is applied to
// a non-directory, etc.
//
// Service implementors can use the following guidelines to decide
// between ErrFailedPrecondition, ErrAborted, and ErrUnavailable:
//
//	(a) Use ErrUnavailable if the client can retry just the failing call.
//	(b) Use ErrAborted if the client should retry at a higher level. For
//	    example, when a client-specified test-and-set fails, indicating the
//	    client should restart a read-modify-write sequence.
//	(c) Use ErrFailedPrecondition if the client should not retry until
//	    the system state has been explicitly fixed. For example, if an "rmdir"
//	    fails because the directory is non-empty, ErrFailedPrecondition
//	    should be returned since the client should not retry unless
//	    the files are deleted from the directory.
//
// GRPC Mapping: FailedPrecondition
// HTTP Mapping: 400 Bad Request
var ErrFailedPrecondition = NewError(codes.FailedPrecondition, "")

// ErrAborted The operation was aborted, typically due to a concurrency issue such as
// a sequencer check failure or transaction abort.
//
// See the guidelines above for deciding between ErrFailedPrecondition,
// ErrAborted, and ErrUnavailable.
//
// GRPC Mapping: Aborted
// HTTP Mapping: 409 Conflict
var ErrAborted = NewError(codes.Aborted, "")

// ErrOutOfRange The operation was attempted past the valid range.  E.g., seeking or
// reading past end-of-file.
//
// Unlike ErrInvalidArgument, this error indicates a problem that may
// be fixed if the system state changes. For example, a 32-bit file
// system will generate ErrInvalidArgument if asked to read at an
// offset that is not in the range [0,2^32-1], but it will generate
// ErrOutOfRange if asked to read from an offset past the current
// file size.
//
// There is a fair bit of overlap between ErrFailedPrecondition and
// ErrOutOfRange.  We recommend using ErrOutOfRange (the more specific
// error) when it applies so that callers who are iterating through
// a space can easily look for an ErrOutOfRange error to detect when
// they are done.
//
// GRPC Mapping: OutOfRange
// HTTP Mapping: 400 Bad Request
var ErrOutOfRange = NewError(codes.OutOfRange, "")

// ErrUnimplemented The operation is not implemented or is not supported/enabled in this
// service.
//
// # Example 1, method not found at server.
//
// # Example 2, compression mechanism used by client not supported at server.
//
// # Example 3, request cardinality violation (method requires exactly one request
// but client sent some other number of requests or server sent some other number of responses)
//
// GRPC Mapping: Unimplemented
// HTTP Mapping: 501 Not Implemented
var ErrUnimplemented = NewError(codes.Unimplemented, "")

// ErrInternal internal errors.  This means that some invariants expected by the
// underlying system have been broken.  This error code is reserved
// for serious errors.
//
// # Example 1, could not decompress, but compression algorithm supported.
//
// # Example 2, flow-control protocol violation.
//
// # Example 3, error parsing request or response proto.
//
// GRPC Mapping: Internal
// HTTP Mapping: 500 ErrInternal Server Error
var ErrInternal = NewError(codes.Internal, "")

// ErrUnavailable The service is currently unavailable.  This is most likely a
// transient condition, which can be corrected by retrying with
// a backoff. Note that it is not always safe to retry
// non-idempotent operations.
//
// # Example 1, server shutting down.
//
// # Example 2, some data transmitted (e.g., request metadata written to TCP connection) before connection breaks
//
// # Example 3, keepalive watchdog times out
//
// See the guidelines above for deciding between ErrFailedPrecondition,
// ErrAborted, and ErrUnavailable.
//
// GRPC Mapping: Unavailable
// HTTP Mapping: 503 Service ErrUnavailable
var ErrUnavailable = NewError(codes.Unavailable, "")

// ErrDataLoss Unrecoverable data loss or corruption.
//
// GRPC Mapping: DataLoss
// HTTP Mapping: 500 ErrInternal Server Error
var ErrDataLoss = NewError(codes.DataLoss, "")

// ErrUnauthenticated The request does not have valid authentication credentials for the
// operation.
//
// For example, incorrect Auth metadata ( Credentials failed to get metadata, Incompatible credentials set on channel
// and call, Invalid host set in :authority metadata, etc.)
//
// GRPC Mapping: Unauthenticated
// HTTP Mapping: 401 Unauthorized
var ErrUnauthenticated = NewError(codes.Unauthenticated, "")
