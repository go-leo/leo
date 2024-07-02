package statusx

import (
	"errors"
	"fmt"
	"github.com/go-leo/gox/protox"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Error struct {
	e *interstatusx.Error
}

// Error wraps a pointer of a Error proto.
func (e *Error) Error() string {
	return e.e.Error()
}

// GRPCStatus returns the gRPC Status.
func (e *Error) GRPCStatus() *grpcstatus.Status {
	return e.e.GRPCStatus()
}

// HTTPStatus returns the HTTP Status.
func (e *Error) HTTPStatus() *httpstatus.Status {
	return e.e.HTTPStatus()
}

// Proto return the gRPC and HTTP status protocol buffers.
func (e *Error) Proto() (*rpcstatus.Status, *httpstatus.Status) {
	return e.e.Proto()
}

// Is implements future errors.Is functionality.
func (e *Error) Is(target error) bool {
	var targetErr *Error
	if !errors.As(target, &targetErr) {
		return false
	}
	return e.e.Is(targetErr.e)
}

// Equals checks if the current status is equal to the target status by
// comparing gRPC status code and http status code.
// It does not compare the details.
func (e *Error) Equals(target error) bool {
	var targetErr *Error
	if !errors.As(target, &targetErr) {
		return false
	}
	return e.e.Equals(targetErr.e)
}

// Wrap wraps the cause error into the current Error.
func (e *Error) Wrap(cause error) *Error {
	if cause == nil {
		return e
	}
	return e.WithDetails(interstatusx.NewCause(cause))
}

// Unwrap unwraps the cause error from the current Error.
func (e *Error) Unwrap() error {
	var errs []error
	for _, detail := range e.Details() {
		cause, ok := detail.(*interstatusx.Cause)
		if !ok {
			continue
		}
		errs = append(errs, cause.Error())
	}
	return errors.Join(errs...)
}

func (e *Error) WithMessage(msg string) *Error {
	return &Error{e: e.e.WithMessage(msg)}
}

func (e *Error) WithMessagef(format string, a ...any) *Error {
	return e.WithMessage(fmt.Sprintf(format, a...))
}

// WithDetails adds additional details to the Error as protocol buffer messages.
func (e *Error) WithDetails(details ...proto.Message) *Error {
	if len(details) == 0 {
		return e
	}
	if e.e.GetGrpcStatus().GetCode() == int32(codes.OK) {
		return e
	}
	return &Error{e: e.e.WithDetails(details...)}
}

// Details return additional details from the Error
func (e *Error) Details() []proto.Message {
	return e.e.Details()
}

// WithHttpHeader sets the http header info.
func (e *Error) WithHttpHeader(infos ...*httpstatus.Header) *Error {
	if len(infos) == 0 {
		return e
	}
	return &Error{e: e.e.WithHttpHeader(infos...)}
}

// HttpHeader gets the http header info.
func (e *Error) HttpHeader() []*httpstatus.Header {
	return e.e.HttpHeader()
}

// WithHttpBody sets the http body.
func (e *Error) WithHttpBody(body proto.Message) *Error {
	if body == nil {
		return e
	}
	anyBody, err := anypb.New(body)
	if err != nil {
		return e
	}
	return &Error{e: e.e.WithHttpBody(anyBody)}
}

// HttpBody gets the http body.
func (e *Error) HttpBody() proto.Message {
	return e.e.HttpBody()
}

// WithErrorInfo sets the error info.
func (e *Error) WithErrorInfo(infos ...*errdetails.ErrorInfo) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// ErrorInfo gets the error info.
func (e *Error) ErrorInfo() []*errdetails.ErrorInfo {
	return protox.ProtoSlice[[]*errdetails.ErrorInfo](e.Details())
}

// WithRequestInfo sets the request info.
func (e *Error) WithRequestInfo(infos ...*errdetails.RequestInfo) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// RequestInfo gets the request info.
func (e *Error) RequestInfo() []*errdetails.RequestInfo {
	return protox.ProtoSlice[[]*errdetails.RequestInfo](e.Details())
}

// WithDebugInfo sets the debug info.
func (e *Error) WithDebugInfo(infos ...*errdetails.DebugInfo) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// DebugInfo gets the debug info.
func (e *Error) DebugInfo() []*errdetails.DebugInfo {
	return protox.ProtoSlice[[]*errdetails.DebugInfo](e.Details())
}

// WithLocalizedMessage sets the localized message info.
func (e *Error) WithLocalizedMessage(infos ...*errdetails.LocalizedMessage) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// LocalizedMessage gets the localized message info.
func (e *Error) LocalizedMessage() []*errdetails.LocalizedMessage {
	return protox.ProtoSlice[[]*errdetails.LocalizedMessage](e.Details())
}

// WithBadRequest sets the bad request info.
func (e *Error) WithBadRequest(infos ...*errdetails.BadRequest) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// BadRequest gets the bad request info.
func (e *Error) BadRequest() []*errdetails.BadRequest {
	return protox.ProtoSlice[[]*errdetails.BadRequest](e.Details())
}

// WithPreconditionFailure sets the precondition failure info.
func (e *Error) WithPreconditionFailure(infos ...*errdetails.PreconditionFailure) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// PreconditionFailure gets the precondition failure info.
func (e *Error) PreconditionFailure() []*errdetails.PreconditionFailure {
	return protox.ProtoSlice[[]*errdetails.PreconditionFailure](e.Details())
}

// WithRetryInfo sets the retry info.
func (e *Error) WithRetryInfo(infos ...*errdetails.RetryInfo) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// RetryInfo gets the retry info.
func (e *Error) RetryInfo() []*errdetails.RetryInfo {
	return protox.ProtoSlice[[]*errdetails.RetryInfo](e.Details())
}

// WithQuotaFailure sets the quota failure info.
func (e *Error) WithQuotaFailure(infos ...*errdetails.QuotaFailure) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// QuotaFailure gets the quota failure info.
func (e *Error) QuotaFailure() []*errdetails.QuotaFailure {
	return protox.ProtoSlice[[]*errdetails.QuotaFailure](e.Details())
}

// WithResourceInfo sets the resource info.
func (e *Error) WithResourceInfo(infos ...*errdetails.ResourceInfo) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// ResourceInfo gets the resource info.
func (e *Error) ResourceInfo() []*errdetails.ResourceInfo {
	return protox.ProtoSlice[[]*errdetails.ResourceInfo](e.Details())
}

// WithHelp sets the help info.
func (e *Error) WithHelp(infos ...*errdetails.Help) *Error {
	return e.WithDetails(protox.MessageSlice(infos)...)
}

// Help gets the help info.
func (e *Error) Help() []*errdetails.Help {
	return protox.ProtoSlice[[]*errdetails.Help](e.Details())
}
