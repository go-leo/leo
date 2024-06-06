package statusx

import (
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

const FailedCode codes.Code = 1717570208

func GrpcCodeFromCode(code codes.Code) int32 {
	if code == FailedCode {
		return int32(codes.Unknown)
	}
	return int32(code)
}

// HttpStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HttpStatusFromCode(code codes.Code) int32 {
	switch code {
	case FailedCode:
		// code 1717570208: gRPC error code is Unknown, but http status is 200
		return int32(httpstatus.Code_OK)
	case codes.OK:
		return int32(httpstatus.Code_OK)
	case codes.Canceled:
		return int32(httpstatus.Code_CLIENT_CLOSED_REQUEST)
	case codes.Unknown:
		return int32(httpstatus.Code_INTERNAL_SERVER_ERROR)
	case codes.InvalidArgument:
		return int32(httpstatus.Code_BAD_REQUEST)
	case codes.DeadlineExceeded:
		return int32(httpstatus.Code_GATEWAY_TIMEOUT)
	case codes.NotFound:
		return int32(httpstatus.Code_NOT_FOUND)
	case codes.AlreadyExists:
		return int32(httpstatus.Code_CONFLICT)
	case codes.PermissionDenied:
		return int32(httpstatus.Code_FORBIDDEN)
	case codes.Unauthenticated:
		return int32(httpstatus.Code_UNAUTHORIZED)
	case codes.ResourceExhausted:
		return int32(httpstatus.Code_TOO_MANY_REQUESTS)
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return int32(httpstatus.Code_BAD_REQUEST)
	case codes.Aborted:
		return int32(httpstatus.Code_CONFLICT)
	case codes.OutOfRange:
		return int32(httpstatus.Code_BAD_REQUEST)
	case codes.Unimplemented:
		return int32(httpstatus.Code_NOT_IMPLEMENTED)
	case codes.Internal:
		return int32(httpstatus.Code_INTERNAL_SERVER_ERROR)
	case codes.Unavailable:
		return int32(httpstatus.Code_SERVICE_UNAVAILABLE)
	case codes.DataLoss:
		return int32(httpstatus.Code_INTERNAL_SERVER_ERROR)
	default:
		grpclog.Warningf("Unknown gRPC error code: %v", code)
		return int32(httpstatus.Code_INTERNAL_SERVER_ERROR)
	}
}
