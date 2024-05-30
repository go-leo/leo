package statusx

/*
import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

var (
	// OK means no error.
	OK = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusOK,
			Reason: http.StatusText(http.StatusOK),
		},
		GrpcStatus: status.New(codes.OK, codes.OK.String()).Proto(),
	}).Froze()

	// OKBut means ok, but have business error
	OKBut = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusOK,
			Reason: http.StatusText(http.StatusOK),
		},
		GrpcStatus: status.New(codes.Unknown, codes.Unknown.String()).Proto(),
	}).Froze()

	// Canceled means request cancelled by the client.
	Canceled = (&Error{
		HttpStatus: &HttpStatus{
			Status: 499,
			Reason: "Canceled",
		},
		GrpcStatus: status.New(codes.Canceled, codes.Canceled.String()).Proto(),
	}).Froze()

	// Unknown means server error. typically a server bug.
	Unknown = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusInternalServerError,
			Reason: http.StatusText(http.StatusInternalServerError),
		},
		GrpcStatus: status.New(codes.Unknown, codes.Unknown.String()).Proto(),
	}).Froze()

	// InvalidArgument means client specified an invalid argument. Check error message and error details for more
	// information.
	InvalidArgument = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusBadRequest,
			Reason: http.StatusText(http.StatusBadRequest),
		},
		GrpcStatus: status.New(codes.InvalidArgument, codes.InvalidArgument.String()).Proto(),
	}).Froze()

	// DeadlineExceeded means request deadline exceeded. This will happen only if the caller sets a deadline that is
	// shorter than the method's default deadline (i.e. requested deadline is not enough for the server to process the
	// request) and the request did not finish within the deadline.
	DeadlineExceeded = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusGatewayTimeout,
			Reason: http.StatusText(http.StatusGatewayTimeout),
		},
		GrpcStatus: status.New(codes.DeadlineExceeded, codes.DeadlineExceeded.String()).Proto(),
	}).Froze()

	// NotFound means a specified resource is not found.
	NotFound = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusNotFound,
			Reason: http.StatusText(http.StatusNotFound),
		},
		GrpcStatus: status.New(codes.NotFound, codes.NotFound.String()).Proto(),
	}).Froze()

	// AlreadyExists means the resource that a client tried to create already exists.
	AlreadyExists = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusConflict,
			Reason: http.StatusText(http.StatusConflict),
		},
		GrpcStatus: status.New(codes.AlreadyExists, codes.AlreadyExists.String()).Proto(),
	}).Froze()

	// PermissionDenied means client does not have sufficient permission. This can happen because the OAuth token does
	// not have the right scopes, the client doesn't have permission, or the API has not been enabled.
	PermissionDenied = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusForbidden,
			Reason: http.StatusText(http.StatusForbidden),
		},
		GrpcStatus: status.New(codes.PermissionDenied, codes.PermissionDenied.String()).Proto(),
	}).Froze()

	// ResourceExhausted means either out of resource quota or reaching rate limiting. The client should look for
	// google.rpc.QuotaFailure error detail for more information.
	ResourceExhausted = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusTooManyRequests,
			Reason: http.StatusText(http.StatusTooManyRequests),
		},
		GrpcStatus: status.New(codes.ResourceExhausted, codes.ResourceExhausted.String()).Proto(),
	}).Froze()

	// FailedPrecondition means request can not be executed in the current system state, such as deleting a non-empty
	// directory.
	FailedPrecondition = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusBadRequest,
			Reason: http.StatusText(http.StatusBadRequest),
		},
		GrpcStatus: status.New(codes.FailedPrecondition, codes.FailedPrecondition.String()).Proto(),
	}).Froze()

	// Aborted means concurrency conflict, such as read-modify-write conflict.
	Aborted = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusConflict,
			Reason: http.StatusText(http.StatusConflict),
		},
		GrpcStatus: status.New(codes.Aborted, codes.Aborted.String()).Proto(),
	}).Froze()

	// OutOfRange means client specified an invalid range.
	OutOfRange = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusBadRequest,
			Reason: http.StatusText(http.StatusBadRequest),
		},
		GrpcStatus: status.New(codes.OutOfRange, codes.OutOfRange.String()).Proto(),
	}).Froze()

	// Unimplemented means api method not implemented by the server.
	Unimplemented = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusNotImplemented,
			Reason: http.StatusText(http.StatusNotImplemented),
		},
		GrpcStatus: status.New(codes.Unimplemented, codes.Unimplemented.String()).Proto(),
	}).Froze()

	// Internal means internal server error. typically a server bug.
	Internal = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusInternalServerError,
			Reason: http.StatusText(http.StatusInternalServerError),
		},
		GrpcStatus: status.New(codes.Internal, codes.Internal.String()).Proto(),
	}).Froze()

	// Unavailable means service unavailable. typically the server is down.
	Unavailable = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusServiceUnavailable,
			Reason: http.StatusText(http.StatusServiceUnavailable),
		},
		GrpcStatus: status.New(codes.Unavailable, codes.Unavailable.String()).Proto(),
	}).Froze()

	// DataLoss means unrecoverable data loss or data corruption. the client should report the error to the user.
	DataLoss = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusInternalServerError,
			Reason: http.StatusText(http.StatusInternalServerError),
		},
		GrpcStatus: status.New(codes.DataLoss, codes.DataLoss.String()).Proto(),
	}).Froze()

	// Unauthenticated means request not authenticated due to missing, invalid, or expired OAuth token.
	Unauthenticated = (&Error{
		HttpStatus: &HttpStatus{
			Status: http.StatusUnauthorized,
			Reason: http.StatusText(http.StatusUnauthorized),
		},
		GrpcStatus: status.New(codes.Unauthenticated, codes.Unauthenticated.String()).Proto(),
	}).Froze()
)
*/
