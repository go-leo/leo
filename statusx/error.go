package statusx

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// ErrorAPI is the error with error details.
// see: https://cloud.google.com/apis/design/errors
type ErrorAPI interface {

	// Error gets the error message.
	error

	// Froze freezes the error.
	Froze() ErrorAPI

	// Wrap wraps the error with the given error.
	Wrap(err error) ErrorAPI

	// Unwrap unwraps the error.
	Unwrap() error

	// Is checks whether the error is the same as the target error.
	Is(target error) bool

	// HTTPStatus gets the http status code.
	HTTPStatus() int

	// GRPCStatus gets the grpc status.
	GRPCStatus() *grpcstatus.Status

	// WithDetails sets the details.
	WithDetails(details ...proto.Message) ErrorAPI

	// Details gets the details.
	Details() []proto.Message

	// WithErrorInfo sets the error info.
	WithErrorInfo(infos ...*errdetails.ErrorInfo) ErrorAPI

	// ErrorInfo gets the error info.
	ErrorInfo() []*errdetails.ErrorInfo

	// WithRetryInfo sets the retry info.
	WithRetryInfo(infos ...*errdetails.RetryInfo) ErrorAPI

	// RetryInfo gets the retry info.
	RetryInfo() []*errdetails.RetryInfo

	// WithDebugInfo sets the debug info.
	WithDebugInfo(infos ...*errdetails.DebugInfo) ErrorAPI

	// DebugInfo gets the debug info.
	DebugInfo() []*errdetails.DebugInfo

	// WithQuotaFailure sets the quota failure info.
	WithQuotaFailure(infos ...*errdetails.QuotaFailure) ErrorAPI

	// QuotaFailure gets the quota failure info.
	QuotaFailure() []*errdetails.QuotaFailure

	// WithPreconditionFailure sets the precondition failure info.
	WithPreconditionFailure(infos ...*errdetails.PreconditionFailure) ErrorAPI

	// PreconditionFailure gets the precondition failure info.
	PreconditionFailure() []*errdetails.PreconditionFailure

	// WithBadRequest sets the bad request info.
	WithBadRequest(infos ...*errdetails.BadRequest) ErrorAPI

	// BadRequest gets the bad request info.
	BadRequest() []*errdetails.BadRequest

	// WithRequestInfo sets the request info.
	WithRequestInfo(infos ...*errdetails.RequestInfo) ErrorAPI

	// RequestInfo gets the request info.
	RequestInfo() []*errdetails.RequestInfo

	// WithResourceInfo sets the resource info.
	WithResourceInfo(infos ...*errdetails.ResourceInfo) ErrorAPI

	// ResourceInfo gets the resource info.
	ResourceInfo() []*errdetails.ResourceInfo

	// WithHelp sets the help info.
	WithHelp(infos ...*errdetails.Help) ErrorAPI

	// Help gets the help info.
	Help() []*errdetails.Help

	// WithLocalizedMessage sets the localized message info.
	WithLocalizedMessage(infos ...*errdetails.LocalizedMessage) ErrorAPI

	// LocalizedMessage gets the localized message info.
	LocalizedMessage() []*errdetails.LocalizedMessage
}
