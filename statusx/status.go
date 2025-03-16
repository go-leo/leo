package statusx

import (
	"encoding/json"
	"fmt"
	"github.com/go-leo/gox/protox"
	"github.com/go-leo/leo/v3/statusx/internal/statuspb"
	"github.com/go-leo/leo/v3/statusx/internal/util"
	"golang.org/x/exp/maps"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"net/http"
	"strings"
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

	// Extra returns additional detail from the Status
	Extra() proto.Message
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
	grpcStatus.Details = statuspb.ToGrpcDetails(st.err.GetDetailInfo())
	return grpcstatus.FromProto(grpcStatus)
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
	return util.ToHttpStatusCode(st.Code())
}

func (st *sampleStatus) Headers() http.Header {
	headers := st.err.GetDetailInfo().GetHeader().GetHeaders()
	header := make(http.Header, len(headers))
	keys := make(map[string]struct{}, len(headers))
	for _, item := range headers {
		header.Add(item.GetKey(), item.GetValue())
		keys[item.GetKey()] = struct{}{}
	}
	header.Add(kKey, strings.Join(maps.Keys(keys), kSeparator))
	return header
}

func (st *sampleStatus) MarshalJSON() ([]byte, error) {
	return marshalHttpBody(st)
}

func (st *sampleStatus) ErrorInfo() *errdetails.ErrorInfo {
	return st.err.GetDetailInfo().GetErrorInfo()
}

func (st *sampleStatus) RetryInfo() *errdetails.RetryInfo {
	return st.err.GetDetailInfo().GetRetryInfo()
}

func (st *sampleStatus) DebugInfo() *errdetails.DebugInfo {
	return st.err.GetDetailInfo().GetDebugInfo()
}

func (st *sampleStatus) QuotaFailure() *errdetails.QuotaFailure {
	return st.err.GetDetailInfo().GetQuotaFailure()
}

func (st *sampleStatus) PreconditionFailure() *errdetails.PreconditionFailure {
	return st.err.GetDetailInfo().GetPreconditionFailure()
}

func (st *sampleStatus) BadRequest() *errdetails.BadRequest {
	return st.err.GetDetailInfo().GetBadRequest()
}

func (st *sampleStatus) RequestInfo() *errdetails.RequestInfo {
	return st.err.GetDetailInfo().GetRequestInfo()
}

func (st *sampleStatus) ResourceInfo() *errdetails.ResourceInfo {
	return st.err.GetDetailInfo().GetResourceInfo()
}

func (st *sampleStatus) Help() *errdetails.Help {
	return st.err.GetDetailInfo().GetHelp()
}

func (st *sampleStatus) LocalizedMessage() *errdetails.LocalizedMessage {
	return st.err.GetDetailInfo().GetLocalizedMessage()
}

func (st *sampleStatus) Extra() proto.Message {
	detail := st.err.GetDetailInfo().GetExtra()
	if detail == nil {
		return nil
	}
	info, err := detail.UnmarshalNew()
	if err != nil {
		panic(err)
	}
	return info
}
