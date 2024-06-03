package statusx

import (
	"fmt"
	"github.com/go-leo/gox/protox"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/protoadapt"
	"google.golang.org/protobuf/types/known/anypb"
	"net/http"
)

// New like grpcstatus.New, but with http status additionally.
func New(c codes.Code, msg string) *grpcstatus.Status {
	switch c {
	case codes.OK:
		return OK(msg)
	case codes.Canceled:
		return Canceled(msg)
	case codes.Unknown:
		return Unknown(msg)
	case codes.InvalidArgument:
		return InvalidArgument(msg)
	case codes.DeadlineExceeded:
		return DeadlineExceeded(msg)
	case codes.NotFound:
		return NotFound(msg)
	case codes.AlreadyExists:
		return AlreadyExists(msg)
	case codes.PermissionDenied:
		return PermissionDenied(msg)
	case codes.ResourceExhausted:
		return ResourceExhausted(msg)
	case codes.FailedPrecondition:
		return FailedPrecondition(msg)
	case codes.Aborted:
		return Aborted(msg)
	case codes.OutOfRange:
		return OutOfRange(msg)
	case codes.Unimplemented:
		return Unimplemented(msg)
	case codes.Internal:
		return Internal(msg)
	case codes.Unavailable:
		return Unavailable(msg)
	case codes.DataLoss:
		return DataLoss(msg)
	case codes.Unauthenticated:
		return Unauthenticated(msg)
	}
	return WithHttpProto(grpcstatus.New(c, msg), &httpstatus.Status{Code: int32(httpstatus.Code_INTERNAL_SERVER_ERROR)})
}

// Newf like grpcstatus.Newf, but with http status additionally.
func Newf(c codes.Code, format string, a ...any) *grpcstatus.Status {
	return New(c, fmt.Sprintf(format, a...))
}

// Error like grpcstatus.Error, but with http status additionally.
func Error(c codes.Code, msg string) error {
	return New(c, msg).Err()
}

// Errorf like grpcstatus.Error, but with http status additionally.
func Errorf(c codes.Code, format string, a ...any) error {
	return Error(c, fmt.Sprintf(format, a...))
}

// Is compares the current status's gRPC and HTTP status codes with another target status.
// returns true if both codes are equal.
func Is(current *grpcstatus.Status, target *grpcstatus.Status) bool {
	currentGrpcProto, currentHttpProto := Proto(current)
	targetGrpcProto, targetHttpProto := Proto(target)
	return currentGrpcProto.GetCode() == targetGrpcProto.GetCode() && currentHttpProto.GetCode() == targetHttpProto.GetCode()
}

// Equals checks if the current status is equal to the target status by comparing their gRPC protocol buffer
// representations using proto.Equal.
func Equals(current *grpcstatus.Status, target *grpcstatus.Status) bool {
	currentGrpcProto, currentHttpProto := Proto(current)
	targetGrpcProto, targetHttpProto := Proto(target)
	return proto.Equal(currentGrpcProto, targetGrpcProto) && proto.Equal(currentHttpProto, targetHttpProto)
}

// WithHttpProto set the http status protocol buffers
func WithHttpProto(grpcStatus *grpcstatus.Status, httpStatus *httpstatus.Status) *grpcstatus.Status {
	return WithDetails(grpcStatus, httpStatus)
}

// Proto return the gRPC and HTTP status protocol buffers.
func Proto(grpcStatus *grpcstatus.Status) (*rpcstatus.Status, *httpstatus.Status) {
	grpcStatusPb := grpcStatus.Proto()
	if grpcStatusPb == nil {
		return nil, nil
	}

	details := grpcStatusPb.GetDetails()
	newDetails := make([]*anypb.Any, 0, len(details))
	var httpAnyProto *anypb.Any
	for _, detail := range grpcStatusPb.GetDetails() {
		if detail.GetTypeUrl() == httpstatus.AnyProto.GetTypeUrl() {
			httpAnyProto = detail
			continue
		}
		newDetails = append(newDetails, detail)
	}
	grpcStatusPb.Details = newDetails

	if httpAnyProto == nil {
		return grpcStatusPb, nil
	}

	httpProto, _ := httpAnyProto.UnmarshalNew()
	return grpcStatusPb, httpProto.(*httpstatus.Status)
}

// WithDetails adds additional details to the status as protocol buffer messages.
func WithDetails(grpcStatus *grpcstatus.Status, details ...proto.Message) *grpcstatus.Status {
	if grpcStatus.Code() == codes.OK {
		return grpcStatus
	}
	v1Details := make([]protoadapt.MessageV1, 0, len(details))
	for _, detail := range details {
		v1Details = append(v1Details, protoadapt.MessageV1Of(detail))
	}
	grpcStatus, _ = grpcStatus.WithDetails(v1Details...)
	return grpcStatus
}

// Details return additional details from the status
func Details(grpcStatus *grpcstatus.Status) []proto.Message {
	details := grpcStatus.Details()
	detailPbs := make([]proto.Message, 0, len(details))
	for _, detail := range details {
		detailPbs = append(detailPbs, detail.(proto.Message))
	}
	return detailPbs
}

// WithErrorInfo sets the error info.
func WithErrorInfo(grpcStatus *grpcstatus.Status, infos ...*errdetails.ErrorInfo) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// ErrorInfo gets the error info.
func ErrorInfo(grpcStatus *grpcstatus.Status) []*errdetails.ErrorInfo {
	return protox.ProtoSlice[[]*errdetails.ErrorInfo](Details(grpcStatus))
}

// WithRequestInfo sets the request info.
func WithRequestInfo(grpcStatus *grpcstatus.Status, infos ...*errdetails.RequestInfo) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// RequestInfo gets the request info.
func RequestInfo(grpcStatus *grpcstatus.Status) []*errdetails.RequestInfo {
	return protox.ProtoSlice[[]*errdetails.RequestInfo](Details(grpcStatus))
}

// WithDebugInfo sets the debug info.
func WithDebugInfo(grpcStatus *grpcstatus.Status, infos ...*errdetails.DebugInfo) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// DebugInfo gets the debug info.
func DebugInfo(grpcStatus *grpcstatus.Status) []*errdetails.DebugInfo {
	return protox.ProtoSlice[[]*errdetails.DebugInfo](Details(grpcStatus))
}

// WithLocalizedMessage sets the localized message info.
func WithLocalizedMessage(grpcStatus *grpcstatus.Status, infos ...*errdetails.LocalizedMessage) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// LocalizedMessage gets the localized message info.
func LocalizedMessage(grpcStatus *grpcstatus.Status) []*errdetails.LocalizedMessage {
	return protox.ProtoSlice[[]*errdetails.LocalizedMessage](Details(grpcStatus))
}

// WithBadRequest sets the bad request info.
func WithBadRequest(grpcStatus *grpcstatus.Status, infos ...*errdetails.BadRequest) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// BadRequest gets the bad request info.
func BadRequest(grpcStatus *grpcstatus.Status) []*errdetails.BadRequest {
	return protox.ProtoSlice[[]*errdetails.BadRequest](Details(grpcStatus))
}

// WithPreconditionFailure sets the precondition failure info.
func WithPreconditionFailure(grpcStatus *grpcstatus.Status, infos ...*errdetails.PreconditionFailure) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// PreconditionFailure gets the precondition failure info.
func PreconditionFailure(grpcStatus *grpcstatus.Status) []*errdetails.PreconditionFailure {
	return protox.ProtoSlice[[]*errdetails.PreconditionFailure](Details(grpcStatus))
}

// WithRetryInfo sets the retry info.
func WithRetryInfo(grpcStatus *grpcstatus.Status, infos ...*errdetails.RetryInfo) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// RetryInfo gets the retry info.
func RetryInfo(grpcStatus *grpcstatus.Status) []*errdetails.RetryInfo {
	return protox.ProtoSlice[[]*errdetails.RetryInfo](Details(grpcStatus))
}

// WithQuotaFailure sets the quota failure info.
func WithQuotaFailure(grpcStatus *grpcstatus.Status, infos ...*errdetails.QuotaFailure) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// QuotaFailure gets the quota failure info.
func QuotaFailure(grpcStatus *grpcstatus.Status) []*errdetails.QuotaFailure {
	return protox.ProtoSlice[[]*errdetails.QuotaFailure](Details(grpcStatus))
}

// WithResourceInfo sets the resource info.
func WithResourceInfo(grpcStatus *grpcstatus.Status, infos ...*errdetails.ResourceInfo) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// ResourceInfo gets the resource info.
func ResourceInfo(grpcStatus *grpcstatus.Status) []*errdetails.ResourceInfo {
	return protox.ProtoSlice[[]*errdetails.ResourceInfo](Details(grpcStatus))
}

// WithHelp sets the help info.
func WithHelp(grpcStatus *grpcstatus.Status, infos ...*errdetails.Help) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// Help gets the help info.
func Help(grpcStatus *grpcstatus.Status) []*errdetails.Help {
	return protox.ProtoSlice[[]*errdetails.Help](Details(grpcStatus))
}

// WithHttpHeader sets the http header info.
func WithHttpHeader(grpcStatus *grpcstatus.Status, infos ...*httpstatus.Header) *grpcstatus.Status {
	return WithDetails(grpcStatus, protox.MessageSlice(infos)...)
}

// HttpHeader gets the http header info.
func HttpHeader(grpcStatus *grpcstatus.Status) []*httpstatus.Header {
	return protox.ProtoSlice[[]*httpstatus.Header](Details(grpcStatus))
}

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromCode(code int32) int32 {
	switch codes.Code(code) {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return 499
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	default:
		grpclog.Warningf("Unknown gRPC error code: %v", code)
		return http.StatusInternalServerError
	}
}
