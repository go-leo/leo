package statusx

import (
	"context"
	"errors"
	"github.com/go-leo/leo/v3/statusx/internal/statuspb"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
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
	case error:
		return fromError(st)
	default:
		return Unknown(Message("%s", obj)), false
	}
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

func fromRpcStatus(grpcProto *rpcstatus.Status) Status {
	st := newStatus(codes.Code(grpcProto.Code))
	st.err.GrpcStatus.Message = grpcProto.GetMessage()
	st.err.DetailInfo, st.err.HttpStatus = st.err.FromGrpcDetails(grpcProto.GetDetails())
	return st
}
