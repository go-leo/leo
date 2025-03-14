package statusx

import (
	"context"
	"errors"
	"github.com/go-leo/leo/v3/proto/leo/status"
	"github.com/go-leo/leo/v3/statusx/internal/statuspb"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func From(obj any) (Status, bool) {
	switch st := obj.(type) {
	case *sampleStatus:
		return st, true
	case *statuspb.Error:
		return &sampleStatus{err: st}, true
	case Status:
		return st, true
	case *rpcstatus.Status:
		return fromRpcStatus(st), true
	case *grpcstatus.Status:
		return fromRpcStatus(st.Proto()), true
	case interface{ GRPCStatus() *grpcstatus.Status }:
		return fromRpcStatus(st.GRPCStatus().Proto()), true
	case *status.HttpBody:
		return fromHttpBody(st), true
	case *http.Response:
		return fromHttpResponse(st)
	case error:
		return fromError(st)
	default:
		return Unknown(Message("%+v", obj)), false
	}
}

func fromRpcStatus(grpcProto *rpcstatus.Status) Status {
	st := newStatus(codes.Code(grpcProto.Code))
	st.err.GrpcStatus.Message = grpcProto.GetMessage()
	details := statuspb.FromDetails(grpcProto.GetDetails())
	st.err.DetailInfo.MergeGrpcDetails(details)
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
	if ok {
		return fromRpcStatus(grpcStatus.Proto()), true
	}
	return Unknown(Message("%+v", err)), false
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
	keys := strings.Split(resp.Header.Get(kKey), kSeparator)
	if len(keys) == 0 {
		return fromHttpBody(body), true
	}
	header := make(http.Header)
	for _, key := range keys {
		values := resp.Header.Values(key)
		for _, value := range values {
			header.Add(key, value)
		}
	}
	return fromHttpBody(body, Headers(header)), true
}

func fromHttpBody(body *status.HttpBody, opts ...Option) Status {
	bodyStatus := body.GetError()
	options := append([]Option{Identifier(bodyStatus.GetIdentifier())}, opts...)
	st := newStatus(codes.Code(bodyStatus.GetStatus()), options...)
	st.err.GrpcStatus.Message = bodyStatus.GetMessage()
	details := statuspb.FromDetails(bodyStatus.GetDetails())
	st.err.DetailInfo.MergeHttpDetails(details)
	return st
}
