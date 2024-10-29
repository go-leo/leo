package statusx

import (
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"net/http"
)

/*
define status code
See: [gRPC documentation]: https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
See: [google api design]: https://cloud.google.com/apis/design/errors
See: [gRPC codes]: https://github.com/grpc/grpc-go/blob/master/codes/codes.go
See: [google rpc code]: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
See: [HTTP to gRPC Status Code Mapping]: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
*/
const kFailedCode codes.Code = 1717570208

// OK Not an error; returned on success.
//
// GRPC Code: OK
// HTTP Status: 200 OK
var OK Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusOK),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.OK),
		},
	},
}

// Failed unlike ErrUnknown error, it just means business logic failed.
//
// For example, if client want sign up, but username already exists,
// it can return Failed("username already exist").
//
// GRPC Code: Unknown
// HTTP Status: 200 OK
var Failed Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusOK),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(kFailedCode),
		},
	},
}

// ErrCanceled The operation was cancelled, typically by the caller.
//
// GRPC Code: Canceled
// HTTP Status: 499 Client Closed Request
var ErrCanceled Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(499),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.Canceled),
		},
	},
}

// ErrUnknown error.
//
// GRPC Code: Unknown
// HTTP Status: 500 Internal Server Error
var ErrUnknown Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusInternalServerError),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.Unknown),
		},
	},
}

// ErrInvalidArgument The client specified an invalid argument.
//
// GRPC Code: InvalidArgument
// HTTP Status: 400 Bad Request
var ErrInvalidArgument Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusBadRequest),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.InvalidArgument),
		},
	},
}

// ErrDeadlineExceeded The deadline expired before the operation could complete.
//
// GRPC Code: DeadlineExceeded
// HTTP Status: 504 Gateway Timeout
var ErrDeadlineExceeded Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusGatewayTimeout),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.DeadlineExceeded),
		},
	},
}

// ErrNotFound Some requested entity (e.g., file or directory) was not found.
//
// GRPC Code: NotFound
// HTTP Status: 404 Not Found
var ErrNotFound Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusNotFound),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.NotFound),
		},
	},
}

// ErrAlreadyExists The entity that a client attempted to create (e.g., file or directory)
// already exists.
//
// GRPC Code: AlreadyExists
// HTTP Status: 409 Conflict
var ErrAlreadyExists Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusConflict),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.AlreadyExists),
		},
	},
}

// ErrPermissionDenied The caller does not have permission to execute the specified
// operation.
//
// GRPC Code: PermissionDenied
// HTTP Status: 403 Forbidden
var ErrPermissionDenied Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusForbidden),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.PermissionDenied),
		},
	},
}

// ErrResourceExhausted Some resource has been exhausted, perhaps a per-user quota, or
// perhaps the entire file system is out of space.
//
// GRPC Code: ResourceExhausted
// HTTP Status: 429 Too Many Requests
var ErrResourceExhausted Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusTooManyRequests),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.ResourceExhausted),
		},
	},
}

// ErrFailedPrecondition The operation was rejected because the system is not in a state
// required for the operation's execution.
//
// GRPC Code: FailedPrecondition
// HTTP Status: 400 Bad Request
var ErrFailedPrecondition Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusBadRequest),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.FailedPrecondition),
		},
	},
}

// ErrAborted The operation was aborted, typically due to a concurrency issue such as
// a sequencer check failure or transaction abort.
//
// GRPC Code: Aborted
// HTTP Status: 409 Conflict
var ErrAborted Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusConflict),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.Aborted),
		},
	},
}

// ErrOutOfRange The operation was attempted past the valid range.
//
// GRPC Code: OutOfRange
// HTTP Status: 400 Bad Request
var ErrOutOfRange Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusBadRequest),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.OutOfRange),
		},
	},
}

// ErrUnimplemented The operation is not implemented or is not supported/enabled in this
// service.
//
// GRPC Code: Unimplemented
// HTTP Status: 501 Not Implemented
var ErrUnimplemented Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusNotImplemented),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.Unimplemented),
		},
	},
}

// ErrInternal internal errors.  This means that some invariants expected by the
// underlying system have been broken.  This error code is reserved
// for serious errors.
//
// GRPC Code: Internal
// HTTP Status: 500 Internal Server Error
var ErrInternal Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusInternalServerError),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.Internal),
		},
	},
}

// ErrUnavailable The service is currently unavailable.
//
// GRPC Code: Unavailable
// HTTP Status: 503 Service Unavailable
var ErrUnavailable Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusServiceUnavailable),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.Unavailable),
		},
	},
}

// ErrDataLoss Unrecoverable data loss or corruption.
//
// GRPC Code: DataLoss
// HTTP Status: 500 Internal Server Error
var ErrDataLoss Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusInternalServerError),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.DataLoss),
		},
	},
}

// ErrUnauthenticated The request does not have valid authentication credentials for the
// operation.
//
// GRPC Code: Unauthenticated
// HTTP Status: 401 Unauthorized
var ErrUnauthenticated Error = &status{
	err: &interstatusx.Error{
		HttpStatus: &httpstatus.HttpResponse{
			Status: int32(http.StatusUnauthorized),
		},
		GrpcStatus: &rpcstatus.Status{
			Code: int32(codes.Unauthenticated),
		},
	},
}
