package statusx

import (
	"context"
	"errors"
	"github.com/go-leo/leo/v3/statusx/internal/statuspb"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"io"
	"net/http"
	"net/url"
)

func From(obj any) (Status, bool) {
	switch st := obj.(type) {
	case *sampleStatus:
		return st, true
	case *statuspb.Error:
		return &sampleStatus{
			err: st,
		}, true
	case Status:
		return st, true
	case *rpcstatus.Status:
		return fromRpcStatus(st), true
	case *grpcstatus.Status:
		return fromRpcStatus(st.Proto()), true
	case interface{ GRPCStatus() *grpcstatus.Status }:
		return fromRpcStatus(st.GRPCStatus().Proto()), true
	case *http.Response:
		return fromHttpResponse(st)
	case error:
		return fromError(st)
	default:
		return Unknown(Message("%s", obj)), false
	}
}

func fromRpcStatus(grpcProto *rpcstatus.Status) Status {
	st := newStatus(codes.Code(grpcProto.Code))
	st.err.GrpcStatus.Message = grpcProto.GetMessage()
	st.err.DetailInfo = statuspb.FromDetails(grpcProto.GetDetails())
	return st
}

func fromError(err error) (Status, bool) {
	if errors.Is(err, context.DeadlineExceeded) {
		return DeadlineExceeded(), true
	}
	if errors.Is(err, context.Canceled) {
		return Canceled(), true
	}
	if urlErr := new(url.Error); errors.As(err, &urlErr) {
		return Unavailable(), true
	}
	if statusErr := new(sampleStatus); errors.As(err, &statusErr) {
		return statusErr, true
	}
	grpcStatus, ok := grpcstatus.FromError(err)
	return fromRpcStatus(grpcStatus.Proto()), ok
}

func fromHttpResponse(resp *http.Response) (Status, bool) {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Unknown(Message(string(data)), Headers(resp.Header)), false
	}
	body, err := unmarshalHttpBody(data)
	if err != nil {
		return Unknown(Message(string(data)), Headers(resp.Header)), false
	}
	bodyStatus := body.GetError()
	st := newStatus(codes.Code(bodyStatus.GetStatus()), Headers(resp.Header), Identifier(bodyStatus.GetIdentifier()))
	st.err.GrpcStatus.Message = bodyStatus.GetMessage()
	details := statuspb.FromDetails(bodyStatus.GetDetails())
	st.err.DetailInfo.ErrorInfo = details.GetErrorInfo()
	st.err.DetailInfo.RetryInfo = details.GetRetryInfo()
	st.err.DetailInfo.DebugInfo = details.GetDebugInfo()
	st.err.DetailInfo.QuotaFailure = details.GetQuotaFailure()
	st.err.DetailInfo.PreconditionFailure = details.GetPreconditionFailure()
	st.err.DetailInfo.BadRequest = details.GetBadRequest()
	st.err.DetailInfo.RequestInfo = details.GetRequestInfo()
	st.err.DetailInfo.ResourceInfo = details.GetResourceInfo()
	st.err.DetailInfo.Help = details.GetHelp()
	st.err.DetailInfo.LocalizedMessage = details.GetLocalizedMessage()
	st.err.DetailInfo.Extra = details.GetExtra()
	return st, true
}
