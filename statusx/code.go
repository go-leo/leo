package statusx

import (
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

// OK Not an error; returned on success.
//
// HTTP Mapping: 200 OK
func OK(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.OK, msg), &httpstatus.Status{Code: int32(httpstatus.Code_OK)})
}

// Canceled The operation was cancelled, typically by the caller.
//
// HTTP Mapping: 499 Client Closed Request
func Canceled(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.Canceled, msg), &httpstatus.Status{Code: int32(httpstatus.Code_CLIENT_CLOSED_REQUEST)})
}

// Unknown error.  For example, this error may be returned when
// a `Status` value received from another address space belongs to
// an error space that is not known in this address space.  Also
// errors raised by APIs that do not return enough error information
// may be converted to this error.
//
// HTTP Mapping: 500 Internal Server Error
func Unknown(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.Unknown, msg), &httpstatus.Status{Code: int32(httpstatus.Code_INTERNAL_SERVER_ERROR)})
}

// InvalidArgument The client specified an invalid argument.  Note that this differs
// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
// that are problematic regardless of the state of the system
// (e.g., a malformed file name).
//
// HTTP Mapping: 400 Bad Request
func InvalidArgument(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.InvalidArgument, msg), &httpstatus.Status{Code: int32(httpstatus.Code_BAD_REQUEST)})
}

// DeadlineExceeded The deadline expired before the operation could complete. For operations
// that change the state of the system, this error may be returned
// even if the operation has completed successfully.  For example, a
// successful response from a server could have been delayed long
// enough for the deadline to expire.
//
// HTTP Mapping: 504 Gateway Timeout
func DeadlineExceeded(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.DeadlineExceeded, msg), &httpstatus.Status{Code: int32(httpstatus.Code_GATEWAY_TIMEOUT)})
}

// NotFound Some requested entity (e.g., file or directory) was not found.
//
// Note to server developers: if a request is denied for an entire class
// of users, such as gradual feature rollout or undocumented allowlist,
// `NOT_FOUND` may be used. If a request is denied for some users within
// a class of users, such as user-based access control, `PERMISSION_DENIED`
// must be used.
//
// HTTP Mapping: 404 Not Found
func NotFound(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.NotFound, msg), &httpstatus.Status{Code: int32(httpstatus.Code_NOT_FOUND)})
}

// AlreadyExists The entity that a client attempted to create (e.g., file or directory)
// already exists.
//
// HTTP Mapping: 409 Conflict
func AlreadyExists(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.AlreadyExists, msg), &httpstatus.Status{Code: int32(httpstatus.Code_CONFLICT)})
}

// PermissionDenied The caller does not have permission to execute the specified
// operation. `PERMISSION_DENIED` must not be used for rejections
// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
// instead for those errors). `PERMISSION_DENIED` must not be
// used if the caller can not be identified (use `UNAUTHENTICATED`
// instead for those errors). This error code does not imply the
// request is valid or the requested entity exists or satisfies
// other pre-conditions.
//
// HTTP Mapping: 403 Forbidden
func PermissionDenied(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.PermissionDenied, msg), &httpstatus.Status{Code: int32(httpstatus.Code_FORBIDDEN)})
}

// ResourceExhausted Some resource has been exhausted, perhaps a per-user quota, or
// perhaps the entire file system is out of space.
//
// HTTP Mapping: 429 Too Many Requests
func ResourceExhausted(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.ResourceExhausted, msg), &httpstatus.Status{Code: int32(httpstatus.Code_TOO_MANY_REQUESTS)})
}

// FailedPrecondition The operation was rejected because the system is not in a state
// required for the operation's execution.  For example, the directory
// to be deleted is non-empty, an rmdir operation is applied to
// a non-directory, etc.
//
// Service implementors can use the following guidelines to decide
// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
//
//	(a) Use `UNAVAILABLE` if the client can retry just the failing call.
//	(b) Use `ABORTED` if the client should retry at a higher level. For
//	    example, when a client-specified test-and-set fails, indicating the
//	    client should restart a read-modify-write sequence.
//	(c) Use `FAILED_PRECONDITION` if the client should not retry until
//	    the system state has been explicitly fixed. For example, if an "rmdir"
//	    fails because the directory is non-empty, `FAILED_PRECONDITION`
//	    should be returned since the client should not retry unless
//	    the files are deleted from the directory.
//
// HTTP Mapping: 400 Bad Request
func FailedPrecondition(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.FailedPrecondition, msg), &httpstatus.Status{Code: int32(httpstatus.Code_BAD_REQUEST)})
}

// Aborted The operation was aborted, typically due to a concurrency issue such as
// a sequencer check failure or transaction abort.
//
// See the guidelines above for deciding between FailedPrecondition,
// Aborted, and Unavailable.
//
// HTTP Mapping: 409 Conflict
func Aborted(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.Aborted, msg), &httpstatus.Status{Code: int32(httpstatus.Code_CONFLICT)})
}

// OutOfRange The operation was attempted past the valid range.  E.g., seeking or
// reading past end-of-file.
//
// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
// be fixed if the system state changes. For example, a 32-bit file
// system will generate `INVALID_ARGUMENT` if asked to read at an
// offset that is not in the range [0,2^32-1], but it will generate
// `OUT_OF_RANGE` if asked to read from an offset past the current
// file size.
//
// There is a fair bit of overlap between `FAILED_PRECONDITION` and
// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
// error) when it applies so that callers who are iterating through
// a space can easily look for an `OUT_OF_RANGE` error to detect when
// they are done.
//
// HTTP Mapping: 400 Bad Request
func OutOfRange(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.OutOfRange, msg), &httpstatus.Status{Code: int32(httpstatus.Code_BAD_REQUEST)})
}

// Unimplemented The operation is not implemented or is not supported/enabled in this
// service.
//
// HTTP Mapping: 501 Not Implemented
func Unimplemented(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.Unimplemented, msg), &httpstatus.Status{Code: int32(httpstatus.Code_NOT_IMPLEMENTED)})
}

// Internal Internal errors.  This means that some invariants expected by the
// underlying system have been broken.  This error code is reserved
// for serious errors.
//
// HTTP Mapping: 500 Internal Server Error
func Internal(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.Internal, msg), &httpstatus.Status{Code: int32(httpstatus.Code_INTERNAL_SERVER_ERROR)})
}

// Unavailable The service is currently unavailable.  This is most likely a
// transient condition, which can be corrected by retrying with
// a backoff. Note that it is not always safe to retry
// non-idempotent operations.
//
// See the guidelines above for deciding between `FAILED_PRECONDITION`,
// `ABORTED`, and `UNAVAILABLE`.
//
// HTTP Mapping: 503 Service Unavailable
// Aborted, and Unavailable.
func Unavailable(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.Unavailable, msg), &httpstatus.Status{Code: int32(httpstatus.Code_SERVICE_UNAVAILABLE)})
}

// DataLoss Unrecoverable data loss or corruption.
//
// HTTP Mapping: 500 Internal Server Error
func DataLoss(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.DataLoss, msg), &httpstatus.Status{Code: int32(httpstatus.Code_INTERNAL_SERVER_ERROR)})
}

// Unauthenticated The request does not have valid authentication credentials for the
// operation.
//
// HTTP Mapping: 401 Unauthorized
func Unauthenticated(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.Unauthenticated, msg), &httpstatus.Status{Code: int32(httpstatus.Code_UNAUTHORIZED)})
}

// OKButFailed is returned on success, but business logic failed.
//
// HTTP Mapping: 200 OK
func OKButFailed(msg string) *grpcstatus.Status {
	return WithHttpStatus(grpcstatus.New(codes.Unknown, msg), &httpstatus.Status{Code: int32(httpstatus.Code_OK)})
}
