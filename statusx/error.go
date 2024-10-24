package statusx

import (
	"errors"
	"fmt"
	"github.com/go-leo/gox/protox"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Api interface {
	error
	// With wraps the current Error with the given options and return new Error.
	With(opts ...Option) Api
	// GRPCStatus returns the gRPC Status.
	GRPCStatus() *grpcstatus.Status
	// HTTPStatus returns the HTTP Status.
	HTTPStatus() *httpstatus.HttpResponse
	// Proto return the gRPC and HTTP status protocol buffers.
	Proto() (*rpcstatus.Status, *httpstatus.HttpResponse)
	// Is implements future errors.Is functionality.
	Is(target error) bool
	// Equals checks if the current status is equal to the target status by
	// comparing gRPC status code and http status code.
	// It does not compare the details.
	Equals(target error) bool
	// Unwrap unwraps the cause error from the current Error.
	Unwrap() error
	// Message gets the message.
	Message() string
	// HttpHeader gets the http header info.
	HttpHeader() []*httpstatus.HttpHeader
	// HttpBody gets the http body.
	HttpBody() *wrapperspb.BytesValue
	// ErrorInfo gets the error info.
	ErrorInfo() *errdetails.ErrorInfo
	// RetryInfo gets the retry info.
	RetryInfo() *errdetails.RetryInfo
	// DebugInfo gets the debug info.
	DebugInfo() *errdetails.DebugInfo
	// QuotaFailure gets the quota failure info.
	QuotaFailure() *errdetails.QuotaFailure
	// PreconditionFailure gets the precondition failure info.
	PreconditionFailure() *errdetails.PreconditionFailure
	// BadRequest gets the bad request info.
	BadRequest() *errdetails.BadRequest
	// RequestInfo gets the request info.
	RequestInfo() *errdetails.RequestInfo
	// ResourceInfo gets the resource info.
	ResourceInfo() *errdetails.ResourceInfo
	// Help gets the help info.
	Help() *errdetails.Help
	// LocalizedMessage gets the localized message info.
	LocalizedMessage() *errdetails.LocalizedMessage
	// Details return additional details from the Error
	Details() []proto.Message
}

type status struct {
	err *interstatusx.Error
}

type Option func(st *status)

// Wrap wraps the cause error into the Error.
func Wrap(err error) Option {
	return func(st *status) {
		if err == nil {
			return
		}
		causeProto, ok := err.(proto.Message)
		if !ok {
			st.err.Cause = &interstatusx.Cause{Cause: &interstatusx.Cause_Message{Message: wrapperspb.String(fmt.Sprintf("%+v", err))}}
			return
		}
		causeAny, err := anypb.New(causeProto)
		if err != nil {
			panic(err)
		}
		st.err.Cause = &interstatusx.Cause{Cause: &interstatusx.Cause_Error{Error: causeAny}}
	}
}

func Message(format string, a ...any) Option {
	return func(st *status) {
		if len(a) <= 0 {
			st.err.GrpcStatus.Message = format
			return
		}
		st.err.GrpcStatus.Message = fmt.Sprintf(format, a...)
	}
}

// HttpHeader sets the http header info.
func HttpHeader(infos ...*httpstatus.HttpHeader) Option {
	return func(st *status) {
		st.err.HttpStatus.Headers = append(st.err.HttpStatus.Headers, infos...)
	}
}

// HttpBody sets the http body.
func HttpBody(info *wrapperspb.BytesValue) Option {
	return func(st *status) {
		st.err.HttpStatus.Body = info.GetValue()
	}
}

// ErrorInfo sets the error info.
func ErrorInfo(info *errdetails.ErrorInfo) Option {
	return func(st *status) {
		st.err.Detail.ErrorInfo = info
	}
}

// RetryInfo sets the retry info.
func RetryInfo(info *errdetails.RetryInfo) Option {
	return func(st *status) {
		st.err.Detail.RetryInfo = info
	}
}

// DebugInfo sets the debug info.
func DebugInfo(info *errdetails.DebugInfo) Option {
	return func(st *status) {
		st.err.Detail.DebugInfo = info
	}
}

// QuotaFailure sets the quota failure info.
func QuotaFailure(info *errdetails.QuotaFailure) Option {
	return func(st *status) {
		st.err.Detail.QuotaFailure = info
	}
}

// PreconditionFailure sets the precondition failure info.
func PreconditionFailure(info *errdetails.PreconditionFailure) Option {
	return func(st *status) {
		st.err.Detail.PreconditionFailure = info
	}
}

// BadRequest sets the bad request info.
func BadRequest(info *errdetails.BadRequest) Option {
	return func(st *status) {
		st.err.Detail.BadRequest = info
	}
}

// RequestInfo sets the request info.
func RequestInfo(info *errdetails.RequestInfo) Option {
	return func(st *status) {
		st.err.Detail.RequestInfo = info
	}
}

// ResourceInfo sets the resource info.
func ResourceInfo(info *errdetails.ResourceInfo) Option {
	return func(st *status) {
		st.err.Detail.ResourceInfo = info
	}
}

// Help sets the help info.
func Help(info *errdetails.Help) Option {
	return func(st *status) {
		st.err.Detail.Help = info
	}
}

// LocalizedMessage sets the localized message info.
func LocalizedMessage(info *errdetails.LocalizedMessage) Option {
	return func(st *status) {
		st.err.Detail.LocalizedMessage = info
	}
}

// Details adds additional details to the Error as protocol buffer messages.
func Details(details ...proto.Message) Option {
	return func(st *status) {
		for _, detail := range details {
			switch item := detail.(type) {
			case *interstatusx.Cause:
				st.err.Cause = item
			case *wrapperspb.StringValue:
				Message(item.GetValue())(st)
			case *httpstatus.HttpHeader:
				HttpHeader(item)(st)
			case *wrapperspb.BytesValue:
				HttpBody(item)(st)
			case *errdetails.ErrorInfo:
				ErrorInfo(item)(st)
			case *errdetails.RetryInfo:
				RetryInfo(item)(st)
			case *errdetails.DebugInfo:
				DebugInfo(item)(st)
			case *errdetails.QuotaFailure:
				QuotaFailure(item)(st)
			case *errdetails.PreconditionFailure:
				PreconditionFailure(item)(st)
			case *errdetails.BadRequest:
				BadRequest(item)(st)
			case *errdetails.RequestInfo:
				RequestInfo(item)(st)
			case *errdetails.ResourceInfo:
				ResourceInfo(item)(st)
			case *errdetails.Help:
				Help(item)(st)
			case *errdetails.LocalizedMessage:
				LocalizedMessage(item)(st)
			default:
				value, err := anypb.New(item)
				if err != nil {
					panic(err)
				}
				st.err.GrpcStatus.Details = append(st.err.GrpcStatus.Details, value)
			}
		}
	}
}

func (st *status) With(opts ...Option) Api {
	clonedSt := &status{
		err: protox.Clone(st.err),
	}
	for _, opt := range opts {
		opt(clonedSt)
	}
	return clonedSt
}

// Error wraps a pointer of a Error proto.
func (st *status) Error() string {
	grpcStatus := st.err.GetGrpcStatus()
	code := codes.Code(grpcStatus.GetCode())
	var message string
	if causeAny := st.err.GetCause(); causeAny != nil {
		message = st.causeMessage(causeAny)
	} else if errorInfo := st.err.GetDetail().GetErrorInfo(); errorInfo != nil {
		message = errorInfo.GetReason()
	} else {
		message = grpcStatus.GetMessage()
	}
	return fmt.Sprintf("statusx: code = %s, desc = %s", code, message)
}

func (st *status) causeMessage(causeAny *interstatusx.Cause) string {
	if causeProto := causeAny.GetError(); causeProto != nil {
		causeErr, _ := causeProto.UnmarshalNew()
		return causeErr.(error).Error()
	}
	if causeMsg := causeAny.GetMessage(); causeMsg != nil {
		return causeMsg.GetValue()
	}
	return ""
}

// GRPCStatus returns the gRPC Status.
func (st *status) GRPCStatus() *grpcstatus.Status {
	grpcStatus := st.err.GetGrpcStatus()

	// copy grpc status details
	details := make([]*anypb.Any, 0, len(grpcStatus.GetDetails())+3)

	// add cause info
	if st.err.GetCause() != nil {
		cause, err := anypb.New(st.err.GetCause())
		if err != nil {
			panic(err)
		}
		details = append(details, cause)
	}

	// add detail
	if st.err.GetDetail() != nil {
		detail, err := anypb.New(st.err.GetDetail())
		if err != nil {
			panic(err)
		}
		details = append(details, detail)
	}

	// add http status info
	if st.err.GetHttpStatus() != nil {
		httpStatus, err := anypb.New(st.err.GetHttpStatus())
		if err != nil {
			panic(err)
		}
		details = append(details, httpStatus)
	}

	// add grpc status details
	details = append(details, grpcStatus.GetDetails()...)

	// return new grpc status
	return grpcstatus.FromProto(&rpcstatus.Status{
		Code:    grpcStatus.GetCode(),
		Message: grpcStatus.GetMessage(),
		Details: details,
	})
}

// HTTPStatus returns the HTTP Status.
func (st *status) HTTPStatus() *httpstatus.HttpResponse {
	return protox.Clone(st.err.GetHttpStatus())
}

// Proto return the gRPC and HTTP status protocol buffers.
func (st *status) Proto() (*rpcstatus.Status, *httpstatus.HttpResponse) {
	return protox.Clone(st.err.GetGrpcStatus()), protox.Clone(st.err.GetHttpStatus())
}

// Is implements future errors.Is functionality.
func (st *status) Is(target error) bool {
	var targetErr *status
	if !errors.As(target, &targetErr) {
		return false
	}
	return proto.Equal(st.err, targetErr.err)
}

// Equals checks if the current status is equal to the target status by
// comparing gRPC status code and http status code.
func (st *status) Equals(target error) bool {
	var targetErr *status
	if !errors.As(target, &targetErr) {
		return false
	}
	if st.err.GetGrpcStatus().GetCode() != targetErr.err.GetGrpcStatus().GetCode() {
		return false
	}
	if st.err.GetHttpStatus().GetStatus() != targetErr.err.GetHttpStatus().GetStatus() {
		return false
	}
	return true
}

// Unwrap unwraps the cause error from the current Error.
func (st *status) Unwrap() error {
	cause := st.err.GetCause()

	// if no cause, return nil
	if cause == nil {
		return nil
	}

	causeAny := cause.GetError()
	// if no cause error, return message
	if causeAny == nil {
		return errors.New(cause.GetMessage().GetValue())
	}

	// unmarshal cause error
	causeProto, err := causeAny.UnmarshalNew()
	if err != nil {
		panic(err)
	}
	// must be error
	return causeProto.(error)
}

// Message gets the message.
func (st *status) Message() string {
	return st.err.GetGrpcStatus().GetMessage()

}

// HttpHeader gets the http header info.
func (st *status) HttpHeader() []*httpstatus.HttpHeader {
	return protox.CloneSlice(st.err.GetHttpStatus().GetHeaders())
}

// HttpBody gets the http body.
func (st *status) HttpBody() *wrapperspb.BytesValue {
	return wrapperspb.Bytes(st.err.GetHttpStatus().GetBody())
}

// ErrorInfo gets the error info.
func (st *status) ErrorInfo() *errdetails.ErrorInfo {
	return protox.Clone(st.err.GetDetail().GetErrorInfo())
}

// RetryInfo gets the retry info.
func (st *status) RetryInfo() *errdetails.RetryInfo {
	return protox.Clone(st.err.GetDetail().GetRetryInfo())
}

// DebugInfo gets the debug info.
func (st *status) DebugInfo() *errdetails.DebugInfo {
	return protox.Clone(st.err.GetDetail().GetDebugInfo())
}

// QuotaFailure gets the quota failure info.
func (st *status) QuotaFailure() *errdetails.QuotaFailure {
	return protox.Clone(st.err.GetDetail().GetQuotaFailure())
}

// PreconditionFailure gets the precondition failure info.
func (st *status) PreconditionFailure() *errdetails.PreconditionFailure {
	return protox.Clone(st.err.GetDetail().GetPreconditionFailure())
}

// BadRequest gets the bad request info.
func (st *status) BadRequest() *errdetails.BadRequest {
	return protox.Clone(st.err.GetDetail().GetBadRequest())
}

// RequestInfo gets the request info.
func (st *status) RequestInfo() *errdetails.RequestInfo {
	return protox.Clone(st.err.GetDetail().GetRequestInfo())
}

// ResourceInfo gets the resource info.
func (st *status) ResourceInfo() *errdetails.ResourceInfo {
	return protox.Clone(st.err.GetDetail().GetResourceInfo())
}

// Help gets the help info.
func (st *status) Help() *errdetails.Help {
	return protox.Clone(st.err.GetDetail().GetHelp())
}

// LocalizedMessage gets the localized message info.
func (st *status) LocalizedMessage() *errdetails.LocalizedMessage {
	return protox.Clone(st.err.GetDetail().GetLocalizedMessage())
}

// Details return additional details from the Error
func (st *status) Details() []proto.Message {
	details := st.err.GetGrpcStatus().GetDetails()
	messages := make([]proto.Message, 0, len(details))
	for _, anyDetail := range details {
		detail, err := anyDetail.UnmarshalNew()
		if err != nil {
			panic(err)
		}
		messages = append(messages, detail)
	}
	return messages
}
