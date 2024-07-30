package statusx

import (
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"net/http"
)

const FailedCode codes.Code = 1717570208

// OK Not an error; returned on success.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: OK
// HTTP Mapping: 200 OK
var OK = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusOK)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.OK)},
	},
}

// Failed unlike ErrUnknown error, it just means business logic failed.
//
// For example, if client want sign up, but username already exists,
// it should return Failed("username already exist").
//
// GRPC Mapping: Unknown
// HTTP Mapping: 200 OK
var Failed = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusOK)},
		GrpcStatus: &rpcstatus.Status{Code: int32(FailedCode)},
	},
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
var ErrCanceled = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(499)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.Canceled)},
	},
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
var ErrUnknown = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusInternalServerError)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.Unknown)},
	},
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
var ErrInvalidArgument = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusBadRequest)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.InvalidArgument)},
	},
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
var ErrDeadlineExceeded = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusGatewayTimeout)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.DeadlineExceeded)},
	},
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
var ErrNotFound = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusNotFound)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.NotFound)},
	},
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
var ErrAlreadyExists = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusConflict)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.AlreadyExists)},
	},
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
var ErrPermissionDenied = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusForbidden)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.PermissionDenied)},
	},
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
var ErrResourceExhausted = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusTooManyRequests)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.ResourceExhausted)},
	},
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
var ErrFailedPrecondition = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusBadRequest)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.FailedPrecondition)},
	},
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
var ErrAborted = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusConflict)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.Aborted)},
	},
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
var ErrOutOfRange = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusBadRequest)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.OutOfRange)},
	},
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
var ErrUnimplemented = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusNotImplemented)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.Unimplemented)},
	},
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
// HTTP Mapping: 500 Internal Server Error
var ErrInternal = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusInternalServerError)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.Internal)},
	},
}

// ErrUnavailable The service is currently unavailable.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: Unavailable
// HTTP Mapping: 503 Service Unavailable
var ErrUnavailable = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusServiceUnavailable)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.Unavailable)},
	},
}

// ErrDataLoss Unrecoverable data loss or corruption.
//
// See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// See: [google api design]: https://cloud.google.com/apis/design/errors
// See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
// See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
//
// GRPC Mapping: DataLoss
// HTTP Mapping: 500 Internal Server Error
var ErrDataLoss = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusInternalServerError)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.DataLoss)},
	},
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
var ErrUnauthenticated = &Error{
	e: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{Status: int32(http.StatusUnauthorized)},
		GrpcStatus: &rpcstatus.Status{Code: int32(codes.Unauthenticated)},
	},
}
