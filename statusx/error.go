package statusx

import (
	"errors"
	"fmt"
	"github.com/go-leo/gox/protox"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"golang.org/x/exp/slices"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Error struct {
	e      *interstatusx.Error
	frozen bool
}

// Error wraps a pointer of a Error proto.
func (e *Error) freeze() *Error {
	e.frozen = true
	return e
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
	return protox.Clone(e.e.GetHttpStatus())
}

// Proto return the gRPC and HTTP status protocol buffers.
func (e *Error) Proto() (*rpcstatus.Status, *httpstatus.Status) {
	return protox.Clone(e.e.GetGrpcStatus()), protox.Clone(e.e.GetHttpStatus())
}

// Is implements future errors.Is functionality.
func (e *Error) Is(target error) bool {
	return e.e.Is(target)
}

// Equals checks if the current status is equal to the target status by
// comparing gRPC status code and http status code.
func (e *Error) Equals(target error) bool {
	if e == nil && target == nil {
		return true
	}
	tse, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.e.GetGrpcStatus().GetCode() == tse.e.GetGrpcStatus().GetCode() &&
		e.e.GetHttpStatus().GetCode() == tse.e.GetHttpStatus().GetCode()
}

// Wrap wraps the cause error into the current Error.
func (e *Error) Wrap(cause error) *Error {
	if cause == nil {
		return e
	}
	if causeProto, ok := cause.(proto.Message); ok {
		if causeAny, err := anypb.New(causeProto); err == nil {
			return e.WithDetails(&interstatusx.Cause{Cause: &interstatusx.Cause_Error{Error: causeAny}})
		}
	}
	return e.WithDetails(&interstatusx.Cause{Cause: &interstatusx.Cause_Message{Message: cause.Error()}})
}

// Unwrap unwraps the cause error from the current Error.
func (e *Error) Unwrap() error {
	var errs []error
	for _, detail := range e.Details() {
		causeDetail, ok := detail.(*interstatusx.Cause)
		if !ok {
			continue
		}
		switch {
		case causeDetail.GetError() != nil:
			causeProto, err := causeDetail.GetError().UnmarshalNew()
			if err != nil {
				continue
			}
			errs = append(errs, causeProto.(error))
			continue
		case len(causeDetail.GetMessage()) > 0:
			errs = append(errs, errors.New(causeDetail.GetMessage()))
			continue
		}
	}
	return errors.Join(errs...)
}

func (e *Error) WithMessage(msg string) *Error {
	cloned := protox.Clone(e.e)
	cloned.GrpcStatus.Message = msg
	return &Error{e: cloned}
}

func (e *Error) WithMessagef(format string, a ...any) *Error {
	cloned := protox.Clone(e.e)
	cloned.GrpcStatus.Message = fmt.Sprintf(format, a...)
	return &Error{e: cloned}
}

// WithDetails adds additional details to the Error as protocol buffer messages.
func (e *Error) WithDetails(details ...proto.Message) *Error {
	if len(details) == 0 {
		return e
	}
	if e.e.GetGrpcStatus().GetCode() == int32(codes.OK) {
		return e
	}
	anyDetails := make([]*anypb.Any, 0, len(details))
	for _, detail := range details {
		anyDetail, err := anypb.New(detail)
		if err != nil {
			continue
		}
		anyDetails = append(anyDetails, anyDetail)
	}
	cloned := protox.Clone(e.e)
	cloned.GrpcStatus.Details = append(cloned.GrpcStatus.Details, anyDetails...)
	return &Error{e: cloned}
}

// Details return additional details from the Error
func (e *Error) Details() []proto.Message {
	detailPbs := make([]proto.Message, 0, len(e.e.GetGrpcStatus().GetDetails()))
	for _, anyDetail := range e.e.GetGrpcStatus().GetDetails() {
		detail, err := anyDetail.UnmarshalNew()
		if err != nil {
			continue
		}
		detailPbs = append(detailPbs, detail)
	}
	return detailPbs
}

// WithHttpHeader sets the http header info.
func (e *Error) WithHttpHeader(infos ...*httpstatus.Header) *Error {
	if len(infos) == 0 {
		return e
	}
	cloned := protox.Clone(e.e)
	cloned.HttpStatus.Headers = append(cloned.HttpStatus.Headers, infos...)
	return &Error{e: cloned}
}

// HttpHeader gets the http header info.
func (e *Error) HttpHeader() []*httpstatus.Header {
	return slices.Clone(e.e.GetHttpStatus().GetHeaders())
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
	cloned := protox.Clone(e.e)
	cloned.HttpStatus.Body = anyBody
	return &Error{e: cloned}
}

// HttpBody gets the http body.
func (e *Error) HttpBody() proto.Message {
	if e.e.GetHttpStatus().GetBody() == nil {
		return nil
	}
	body, err := e.e.GetHttpStatus().GetBody().UnmarshalNew()
	if err != nil {
		return nil
	}
	return body
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
