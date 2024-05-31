package statusx

import (
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"
)

// Status is the status interface.
type Status interface {
	// Clone clones the status.
	Clone() Status

	// Is only needs to compare both the grpc status code and the http status code.
	Is(target Status) bool

	// Equals compares not only the grpc status code and the http status code, but also the details.
	Equals(target Status) bool

	// String returns the string.
	String() string

	// GrpcProto returns the grpc proto.
	GrpcProto() *rpcstatus.Status

	// HttpProto returns the httpstatus proto.
	HttpProto() *httpstatus.Status

	// Message returns the message.
	Message() string

	// WithMessage sets the message.
	WithMessage(msg string, args ...any) Status

	// Err returns the error.
	Err() error

	// WithErr sets the error.
	WithErr(err error) Status

	// WithDetails sets the details.
	WithDetails(details ...proto.Message) Status

	// Details gets the details.
	Details() []proto.Message

	// WithErrorInfo sets the error info.
	WithErrorInfo(infos ...*errdetails.ErrorInfo) Status

	// ErrorInfo gets the error info.
	ErrorInfo() []*errdetails.ErrorInfo

	// WithRequestInfo sets the request info.
	WithRequestInfo(infos ...*errdetails.RequestInfo) Status

	// RequestInfo gets the request info.
	RequestInfo() []*errdetails.RequestInfo

	// WithDebugInfo sets the debug info.
	WithDebugInfo(infos ...*errdetails.DebugInfo) Status

	// DebugInfo gets the debug info.
	DebugInfo() []*errdetails.DebugInfo

	// WithLocalizedMessage sets the localized message info.
	WithLocalizedMessage(infos ...*errdetails.LocalizedMessage) Status

	// LocalizedMessage gets the localized message info.
	LocalizedMessage() []*errdetails.LocalizedMessage

	// WithHelp sets the help info.
	WithHelp(infos ...*errdetails.Help) Status

	// Help gets the help info.
	Help() []*errdetails.Help

	// WithBadRequest sets the bad request info.
	WithBadRequest(infos ...*errdetails.BadRequest) Status

	// BadRequest gets the bad request info.
	BadRequest() []*errdetails.BadRequest

	// WithPreconditionFailure sets the precondition failure info.
	WithPreconditionFailure(infos ...*errdetails.PreconditionFailure) Status

	// PreconditionFailure gets the precondition failure info.
	PreconditionFailure() []*errdetails.PreconditionFailure

	// WithRetryInfo sets the retry info.
	WithRetryInfo(infos ...*errdetails.RetryInfo) Status

	// RetryInfo gets the retry info.
	RetryInfo() []*errdetails.RetryInfo

	// WithQuotaFailure sets the quota failure info.
	WithQuotaFailure(infos ...*errdetails.QuotaFailure) Status

	// QuotaFailure gets the quota failure info.
	QuotaFailure() []*errdetails.QuotaFailure

	// WithResourceInfo sets the resource info.
	WithResourceInfo(infos ...*errdetails.ResourceInfo) Status

	// ResourceInfo gets the resource info.
	ResourceInfo() []*errdetails.ResourceInfo
}
