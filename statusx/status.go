package statusx

import (
	"encoding/json"
	"fmt"
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/leo/v3/proto/leo/status"
	"github.com/go-leo/leo/v3/statusx/internal/statuspb"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type Status interface {
	error

	// Identifier returns the identifier.
	Identifier() string

	// Code returns the status code.
	Code() codes.Code

	// Message returns the message.
	Message() string

	// GRPCStatus returns the gRPC Status.
	// see: https://github.com/grpc/grpc-go/blame/8528f4387f276518050f2b71a9dee1e3fb19d924/status/status.go#L100
	// type grpcstatus interface{ GRPCStatus() *Status }
	GRPCStatus() *grpcstatus.Status

	// HTTPStatus returns the HTTP Status.
	HTTPStatus() *httpstatus.HttpResponse

	// Is implements future errors.Is functionality.
	Is(target error) bool

	// StatusCode returns the http status code.
	StatusCode() int

	// Headers returns the http header info.
	Headers() http.Header

	// Marshaler implements json.Marshaler.
	json.Marshaler

	// ErrorInfo returns the error info.
	ErrorInfo() *errdetails.ErrorInfo

	// RetryInfo returns the retry info.
	RetryInfo() *errdetails.RetryInfo

	// DebugInfo returns the debug info.
	DebugInfo() *errdetails.DebugInfo

	// QuotaFailure returns the quota failure info.
	QuotaFailure() *errdetails.QuotaFailure

	// PreconditionFailure returns the precondition failure info.
	PreconditionFailure() *errdetails.PreconditionFailure

	// BadRequest returns the bad request info.
	BadRequest() *errdetails.BadRequest

	// RequestInfo returns the request info.
	RequestInfo() *errdetails.RequestInfo

	// ResourceInfo returns the resource info.
	ResourceInfo() *errdetails.ResourceInfo

	// Help returns the help info.
	Help() *errdetails.Help

	// LocalizedMessage returns the localized message info.
	LocalizedMessage() *errdetails.LocalizedMessage

	// Detail returns additional detail from the Status
	Detail() []proto.Message
}

var _ Status = (*sampleStatus)(nil)

type sampleStatus struct {
	err *statuspb.Error
}

func (st *sampleStatus) Error() string {
	return fmt.Sprintf("status: code = %s, status-code = %d, desc = %s", st.Code(), st.StatusCode(), st.Message())
}

func (st *sampleStatus) Identifier() string {
	return st.err.GetDetailInfo().GetIdentifier().GetValue()
}

func (st *sampleStatus) Code() codes.Code {
	if st == nil || st.err == nil {
		return codes.OK
	}
	return codes.Code(st.err.GetGrpcStatus().GetCode())
}

func (st *sampleStatus) Message() string {
	if st == nil || st.err == nil {
		return ""
	}
	return st.err.GetGrpcStatus().GetMessage()
}

func (st *sampleStatus) GRPCStatus() *grpcstatus.Status {
	grpcStatus := protox.Clone(st.err.GetGrpcStatus())
	grpcStatus.Details = st.err.ToGrpcDetails()
	return grpcstatus.FromProto(grpcStatus)
}

func (st *sampleStatus) HTTPStatus() *httpstatus.HttpResponse {
	httpStatus := protox.Clone(st.err.GetHttpStatus())
	httpStatus.Body, _ = st.MarshalJSON()
	return httpStatus
}

func (st *sampleStatus) Is(target error) bool {
	targetStatus, ok := From(target)
	if !ok {
		return false
	}
	return targetStatus.Code() == st.Code() &&
		targetStatus.StatusCode() == st.StatusCode() &&
		targetStatus.Identifier() == st.Identifier()
}

func (st *sampleStatus) StatusCode() int {
	return int(st.err.GetHttpStatus().GetStatus())
}

func (st *sampleStatus) Headers() http.Header {
	header := make(http.Header)
	for _, item := range st.err.GetHttpStatus().GetHeaders() {
		header.Add(item.GetKey(), item.GetValue())
	}
	return header
}

func (st *sampleStatus) MarshalJSON() ([]byte, error) {
	body := &status.HttpBody{
		Error: &status.HttpBody_Status{
			Code:    int32(st.StatusCode()),
			Message: st.Message(),
			Status:  code.Code(st.Code()),
			Details: st.err.HttpDetails(),
		},
	}
	return protojson.MarshalOptions{}.Marshal(body)
}

func (st *sampleStatus) ErrorInfo() *errdetails.ErrorInfo {
	return protox.Clone(st.err.GetDetailInfo().GetErrorInfo())
}

func (st *sampleStatus) RetryInfo() *errdetails.RetryInfo {
	return protox.Clone(st.err.GetDetailInfo().GetRetryInfo())
}

func (st *sampleStatus) DebugInfo() *errdetails.DebugInfo {
	return protox.Clone(st.err.GetDetailInfo().GetDebugInfo())
}

func (st *sampleStatus) QuotaFailure() *errdetails.QuotaFailure {
	return protox.Clone(st.err.GetDetailInfo().GetQuotaFailure())
}

func (st *sampleStatus) PreconditionFailure() *errdetails.PreconditionFailure {
	return protox.Clone(st.err.GetDetailInfo().GetPreconditionFailure())
}

func (st *sampleStatus) BadRequest() *errdetails.BadRequest {
	return protox.Clone(st.err.GetDetailInfo().GetBadRequest())
}

func (st *sampleStatus) RequestInfo() *errdetails.RequestInfo {
	return protox.Clone(st.err.GetDetailInfo().GetRequestInfo())
}

func (st *sampleStatus) ResourceInfo() *errdetails.ResourceInfo {
	return protox.Clone(st.err.GetDetailInfo().GetResourceInfo())
}

func (st *sampleStatus) Help() *errdetails.Help {
	return protox.Clone(st.err.GetDetailInfo().GetHelp())
}

func (st *sampleStatus) LocalizedMessage() *errdetails.LocalizedMessage {
	return protox.Clone(st.err.GetDetailInfo().GetLocalizedMessage())
}

func (st *sampleStatus) Detail() []proto.Message {
	var r []proto.Message
	for _, infoAny := range st.err.GetDetailInfo().GetDetails() {
		info, err := infoAny.UnmarshalNew()
		if err != nil {
			panic(err)
		}
		r = append(r, info)
	}
	return r
}
