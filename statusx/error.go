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
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Error struct {
	e *interstatusx.Error
}

type options struct {
	Cause   *interstatusx.Cause
	Message string

	HttpHeader []*httpstatus.HttpHeader
	HttpBody   *wrapperspb.BytesValue

	ErrorInfo           *errdetails.ErrorInfo
	RetryInfo           *errdetails.RetryInfo
	DebugInfo           *errdetails.DebugInfo
	QuotaFailure        *errdetails.QuotaFailure
	PreconditionFailure *errdetails.PreconditionFailure
	BadRequest          *errdetails.BadRequest
	RequestInfo         *errdetails.RequestInfo
	ResourceInfo        *errdetails.ResourceInfo
	Help                *errdetails.Help
	LocalizedMessage    *errdetails.LocalizedMessage

	Details []*anypb.Any
}

type Option func(o *options)

// Wrap wraps the cause error into the Error.
func Wrap(err error) Option {
	return func(o *options) {
		if err == nil {
			return
		}
		causeProto, ok := err.(proto.Message)
		if !ok {
			o.Cause = &interstatusx.Cause{Cause: &interstatusx.Cause_Message{Message: fmt.Sprintf("%+v", err)}}
			return
		}
		causeAny, err := anypb.New(causeProto)
		if err != nil {
			panic(err)
		}
		o.Cause = &interstatusx.Cause{Cause: &interstatusx.Cause_Error{Error: causeAny}}
	}
}

func Message(format string, a ...any) Option {
	return func(o *options) {
		if len(a) <= 0 {
			o.Message = format
			return
		}
		o.Message = fmt.Sprintf(format, a...)
	}
}

// HttpHeader sets the http header info.
func HttpHeader(infos ...*httpstatus.HttpHeader) Option {
	return func(o *options) {
		o.HttpHeader = append(o.HttpHeader, infos...)
	}
}

// HttpBody sets the http body.
func HttpBody(info *wrapperspb.BytesValue) Option {
	return func(o *options) {
		o.HttpBody = info
	}
}

// ErrorInfo sets the error info.
func ErrorInfo(info *errdetails.ErrorInfo) Option {
	return func(o *options) {
		o.ErrorInfo = info
	}
}

// RetryInfo sets the retry info.
func RetryInfo(info *errdetails.RetryInfo) Option {
	return func(o *options) {
		o.RetryInfo = info
	}
}

// DebugInfo sets the debug info.
func DebugInfo(info *errdetails.DebugInfo) Option {
	return func(o *options) {
		o.DebugInfo = info
	}
}

// QuotaFailure sets the quota failure info.
func QuotaFailure(info *errdetails.QuotaFailure) Option {
	return func(o *options) {
		o.QuotaFailure = info
	}
}

// PreconditionFailure sets the precondition failure info.
func PreconditionFailure(info *errdetails.PreconditionFailure) Option {
	return func(o *options) {
		o.PreconditionFailure = info
	}
}

// BadRequest sets the bad request info.
func BadRequest(info *errdetails.BadRequest) Option {
	return func(o *options) {
		o.BadRequest = info
	}
}

// RequestInfo sets the request info.
func RequestInfo(info *errdetails.RequestInfo) Option {
	return func(o *options) {
		o.RequestInfo = info
	}
}

// ResourceInfo sets the resource info.
func ResourceInfo(info *errdetails.ResourceInfo) Option {
	return func(o *options) {
		o.ResourceInfo = info
	}
}

// Help sets the help info.
func Help(info *errdetails.Help) Option {
	return func(o *options) {
		o.Help = info
	}
}

// LocalizedMessage sets the localized message info.
func LocalizedMessage(info *errdetails.LocalizedMessage) Option {
	return func(o *options) {
		o.LocalizedMessage = info
	}
}

// Details adds additional details to the Error as protocol buffer messages.
func Details(details ...proto.Message) Option {
	return func(o *options) {
		for _, detail := range details {
			switch item := detail.(type) {
			case *interstatusx.Cause:
				o.Cause = item
			case *wrapperspb.StringValue:
				Message(item.GetValue())(o)
			case *httpstatus.HttpHeader:
				HttpHeader(item)(o)
			case *wrapperspb.BytesValue:
				HttpBody(item)(o)
			case *errdetails.ErrorInfo:
				ErrorInfo(item)(o)
			case *errdetails.RetryInfo:
				RetryInfo(item)(o)
			case *errdetails.DebugInfo:
				DebugInfo(item)(o)
			case *errdetails.QuotaFailure:
				QuotaFailure(item)(o)
			case *errdetails.PreconditionFailure:
				PreconditionFailure(item)(o)
			case *errdetails.BadRequest:
				BadRequest(item)(o)
			case *errdetails.RequestInfo:
				RequestInfo(item)(o)
			case *errdetails.ResourceInfo:
				ResourceInfo(item)(o)
			case *errdetails.Help:
				Help(item)(o)
			case *errdetails.LocalizedMessage:
				LocalizedMessage(item)(o)
			default:
				value, err := anypb.New(item)
				if err != nil {
					panic(err)
				}
				o.Details = append(o.Details, value)
			}
		}
	}
}

func (e *Error) With(opts ...Option) *Error {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return &Error{
		e: &interstatusx.Error{
			Cause: o.Cause,
			Detail: &interstatusx.Detail{
				ErrorInfo:           o.ErrorInfo,
				RetryInfo:           o.RetryInfo,
				DebugInfo:           o.DebugInfo,
				QuotaFailure:        o.QuotaFailure,
				PreconditionFailure: o.PreconditionFailure,
				BadRequest:          o.BadRequest,
				RequestInfo:         o.RequestInfo,
				ResourceInfo:        o.ResourceInfo,
				Help:                o.Help,
				LocalizedMessage:    o.LocalizedMessage,
			},
			HttpStatus: &httpstatus.HttpResponse{
				Status:  e.e.HttpStatus.GetStatus(),
				Reason:  e.e.HttpStatus.GetReason(),
				Headers: o.HttpHeader,
				Body:    o.HttpBody.GetValue(),
			},
			GrpcStatus: &rpcstatus.Status{
				Code:    e.e.GetGrpcStatus().GetCode(),
				Message: o.Message,
				Details: o.Details,
			},
		},
	}
}

// Error wraps a pointer of a Error proto.
func (e *Error) Error() string {
	return fmt.Sprintf("statusx: code = %s desc = %s", codes.Code(e.e.GetGrpcStatus().GetCode()), e.e.GetGrpcStatus().GetMessage())
}

// GRPCStatus returns the gRPC Status.
func (e *Error) GRPCStatus() *grpcstatus.Status {
	grpcStatus := e.e.GetGrpcStatus()
	details := make([]*anypb.Any, 0, len(grpcStatus.GetDetails())+3)
	if e.e.GetCause() != nil {
		cause, err := anypb.New(e.e.GetCause())
		if err != nil {
			panic(err)
		}
		details = append(details, cause)
	}
	if e.e.GetDetail() != nil {
		detail, err := anypb.New(e.e.GetDetail())
		if err != nil {
			panic(err)
		}
		details = append(details, detail)
	}
	if e.e.GetHttpStatus() != nil {
		httpStatus, err := anypb.New(e.e.GetHttpStatus())
		if err != nil {
			panic(err)
		}
		details = append(details, httpStatus)
	}
	details = append(details, grpcStatus.GetDetails()...)
	return grpcstatus.FromProto(&rpcstatus.Status{
		Code:    grpcStatus.GetCode(),
		Message: grpcStatus.GetMessage(),
		Details: details,
	})
}

// HTTPStatus returns the HTTP Status.
func (e *Error) HTTPStatus() *httpstatus.HttpResponse {
	return protox.Clone(e.e.GetHttpStatus())
}

// Proto return the gRPC and HTTP status protocol buffers.
func (e *Error) Proto() (*rpcstatus.Status, *httpstatus.HttpResponse) {
	return protox.Clone(e.e.GetGrpcStatus()), protox.Clone(e.e.GetHttpStatus())
}

// Is implements future errors.Is functionality.
func (e *Error) Is(target error) bool {
	var targetErr *Error
	if !errors.As(target, &targetErr) {
		return false
	}
	return proto.Equal(e.e, targetErr.e)
}

// Equals checks if the current status is equal to the target status by
// comparing gRPC status code and http status code.
// It does not compare the details.
func (e *Error) Equals(target error) bool {
	var targetErr *Error
	if !errors.As(target, &targetErr) {
		return false
	}
	if e.e.GetGrpcStatus().GetCode() != targetErr.e.GetGrpcStatus().GetCode() {
		return false
	}
	if e.e.GetHttpStatus().GetStatus() != targetErr.e.GetHttpStatus().GetStatus() {
		return false
	}
	return true
}

// Unwrap unwraps the cause error from the current Error.
func (e *Error) Unwrap() error {
	cause := e.e.GetCause()
	if cause == nil {
		return nil
	}
	causeAny := cause.GetError()
	if causeAny == nil {
		return errors.New(cause.GetMessage())
	}
	causeProto, err := causeAny.UnmarshalNew()
	if err != nil {
		panic(err)
	}
	if err, ok := causeProto.(error); ok {
		return err
	}
	js, err := protojson.Marshal(causeProto)
	if err != nil {
		panic(err)
	}
	return errors.New(string(js))
}

// Message gets the message.
func (e *Error) Message() string {
	return e.e.GetGrpcStatus().GetMessage()

}

// HttpHeader gets the http header info.
func (e *Error) HttpHeader() []*httpstatus.HttpHeader {
	return protox.CloneSlice(e.e.GetHttpStatus().GetHeaders())
}

// HttpBody gets the http body.
func (e *Error) HttpBody() []byte {
	return e.e.GetHttpStatus().GetBody()
}

// ErrorInfo gets the error info.
func (e *Error) ErrorInfo() *errdetails.ErrorInfo {
	return protox.Clone(e.e.GetDetail().GetErrorInfo())
}

// RetryInfo gets the retry info.
func (e *Error) RetryInfo() *errdetails.RetryInfo {
	return protox.Clone(e.e.GetDetail().GetRetryInfo())
}

// DebugInfo gets the debug info.
func (e *Error) DebugInfo() *errdetails.DebugInfo {
	return protox.Clone(e.e.GetDetail().GetDebugInfo())
}

// QuotaFailure gets the quota failure info.
func (e *Error) QuotaFailure() *errdetails.QuotaFailure {
	return protox.Clone(e.e.GetDetail().GetQuotaFailure())
}

// PreconditionFailure gets the precondition failure info.
func (e *Error) PreconditionFailure() *errdetails.PreconditionFailure {
	return protox.Clone(e.e.GetDetail().GetPreconditionFailure())
}

// BadRequest gets the bad request info.
func (e *Error) BadRequest() *errdetails.BadRequest {
	return protox.Clone(e.e.GetDetail().GetBadRequest())
}

// RequestInfo gets the request info.
func (e *Error) RequestInfo() *errdetails.RequestInfo {
	return protox.Clone(e.e.GetDetail().GetRequestInfo())
}

// ResourceInfo gets the resource info.
func (e *Error) ResourceInfo() *errdetails.ResourceInfo {
	return protox.Clone(e.e.GetDetail().GetResourceInfo())
}

// Help gets the help info.
func (e *Error) Help() *errdetails.Help {
	return protox.Clone(e.e.GetDetail().GetHelp())
}

// LocalizedMessage gets the localized message info.
func (e *Error) LocalizedMessage() *errdetails.LocalizedMessage {
	return protox.Clone(e.e.GetDetail().GetLocalizedMessage())
}

// Details return additional details from the Error
func (e *Error) Details() []proto.Message {
	details := e.e.GetGrpcStatus().GetDetails()
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
