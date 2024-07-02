package statusx

import (
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

const FailedCode codes.Code = 1717570208

func GrpcCodeFromCode(code codes.Code) codes.Code {
	if code == FailedCode {
		return codes.Unknown
	}
	return code
}

// HttpStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HttpStatusFromCode(code codes.Code) httpstatus.Code {
	switch code {
	case FailedCode:
		// code 1717570208: gRPC error code is Unknown, but http status is 200
		return httpstatus.Code_OK
	case codes.OK:
		return httpstatus.Code_OK
	case codes.Canceled:
		return httpstatus.Code_CLIENT_CLOSED_REQUEST
	case codes.Unknown:
		return httpstatus.Code_INTERNAL_SERVER_ERROR
	case codes.InvalidArgument:
		return httpstatus.Code_BAD_REQUEST
	case codes.DeadlineExceeded:
		return httpstatus.Code_GATEWAY_TIMEOUT
	case codes.NotFound:
		return httpstatus.Code_NOT_FOUND
	case codes.AlreadyExists:
		return httpstatus.Code_CONFLICT
	case codes.PermissionDenied:
		return httpstatus.Code_FORBIDDEN
	case codes.Unauthenticated:
		return httpstatus.Code_UNAUTHORIZED
	case codes.ResourceExhausted:
		return httpstatus.Code_TOO_MANY_REQUESTS
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return httpstatus.Code_BAD_REQUEST
	case codes.Aborted:
		return httpstatus.Code_CONFLICT
	case codes.OutOfRange:
		return httpstatus.Code_BAD_REQUEST
	case codes.Unimplemented:
		return httpstatus.Code_NOT_IMPLEMENTED
	case codes.Internal:
		return httpstatus.Code_INTERNAL_SERVER_ERROR
	case codes.Unavailable:
		return httpstatus.Code_SERVICE_UNAVAILABLE
	case codes.DataLoss:
		return httpstatus.Code_INTERNAL_SERVER_ERROR
	default:
		grpclog.Warningf("Unknown gRPC error code: %v", code)
		return httpstatus.Code_INTERNAL_SERVER_ERROR
	}
}
