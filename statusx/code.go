package statusx

import (
	"google.golang.org/grpc/codes"
)

// OK not an error; returned on success.
func OK(opts ...Option) Status {
	return newStatus(codes.OK, opts...)
}

// Canceled error indicates the operation was canceled (typically by the caller).
func Canceled(opts ...Option) Status {
	return newStatus(codes.Canceled, opts...)
}

// Unknown error.
func Unknown(opts ...Option) Status {
	return newStatus(codes.Unknown, opts...)
}

// InvalidArgument error indicates client specified an invalid argument.
func InvalidArgument(opts ...Option) Status {
	return newStatus(codes.InvalidArgument, opts...)
}

// DeadlineExceeded error means operation expired before completion.
func DeadlineExceeded(opts ...Option) Status {
	return newStatus(codes.DeadlineExceeded, opts...)
}

// NotFound error means some requested entity (e.g., file or directory) was not found.
func NotFound(opts ...Option) Status {
	return newStatus(codes.NotFound, opts...)
}

// AlreadyExists error means an attempt to create an entity failed because one already exists.
func AlreadyExists(opts ...Option) Status {
	return newStatus(codes.AlreadyExists, opts...)
}

// PermissionDenied error indicates the caller does not have permission to execute the specified
// operation.
func PermissionDenied(opts ...Option) Status {
	return newStatus(codes.PermissionDenied, opts...)
}

// ResourceExhausted error indicates some resource has been exhausted, perhaps a per-user quota, or
// perhaps the entire file system is out of space.
func ResourceExhausted(opts ...Option) Status {
	return newStatus(codes.ResourceExhausted, opts...)
}

// FailedPrecondition error indicates the operation was rejected because the system is not in a state
// required for the operation's execution.
func FailedPrecondition(opts ...Option) Status {
	return newStatus(codes.FailedPrecondition, opts...)
}

// Aborted error indicates the operation was aborted, typically due to a concurrency issue such as
// a sequencer check failure or transaction abort.
func Aborted(opts ...Option) Status {
	return newStatus(codes.Aborted, opts...)
}

// OutOfRange error means the operation was attempted past the valid range.
func OutOfRange(opts ...Option) Status {
	return newStatus(codes.OutOfRange, opts...)
}

// Unimplemented error indicates the operation is not implemented or is not supported/enabled in this
// service.
func Unimplemented(opts ...Option) Status {
	return newStatus(codes.Unimplemented, opts...)
}

// Internal errors means that some invariants expected by the underlying system have been broken.
// This error code is reserved for serious errors.
func Internal(opts ...Option) Status {
	return newStatus(codes.Internal, opts...)
}

// Unavailable error indicates the service is currently unavailable.
func Unavailable(opts ...Option) Status {
	return newStatus(codes.Unavailable, opts...)
}

// DataLoss error indicates unrecoverable data loss or corruption.
func DataLoss(opts ...Option) Status {
	return newStatus(codes.DataLoss, opts...)
}

// Unauthenticated error indicates the request does not have valid authentication credentials for the
// operation.
func Unauthenticated(opts ...Option) Status {
	return newStatus(codes.Unauthenticated, opts...)
}
