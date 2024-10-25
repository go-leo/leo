package statusx

import (
	"context"
	"errors"
	"fmt"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	grpcstatus "google.golang.org/grpc/status"
	"net/http"
	"net/url"
	"strconv"
)

func From(obj any) Api {
	switch err := obj.(type) {
	case *status:
		return err
	case Api:
		return err
	case *interstatusx.Error:
		return &status{err: err}
	case *http.Response:
		st := &status{}
		st.From(err)
		return st
	case *rpcstatus.Status:
		return fromRpcStatus(err)
	case *grpcstatus.Status:
		return fromRpcStatus(err.Proto())
	case interface{ GRPCStatus() *grpcstatus.Status }:
		return From(err.GRPCStatus())
	case error:
		return fromError(err)
	default:
		return ErrUnknown
	}
}

func FromGrpcError(err error) Api {
	if err == nil {
		return nil
	}
	if grpcStatus, ok := grpcstatus.FromError(err); ok {
		return fromRpcStatus(grpcStatus.Proto())
	}
	return ErrUnknown.With(Wrap(err))
}

func FromHttpError(err error) Api {
	if err == nil {
		return nil
	}

}

func fromError(err error) Api {
	if errors.Is(err, context.DeadlineExceeded) {
		return ErrDeadlineExceeded.With(Wrap(err))
	}
	if errors.Is(err, context.Canceled) {
		return ErrCanceled.With(Wrap(err))
	}
	if urlErr := new(url.Error); errors.As(err, &urlErr) {
		return ErrUnavailable.With(Message(strconv.Quote(fmt.Sprintf("%s %s: %s", urlErr.Op, urlErr.URL, urlErr.Err))), Wrap(urlErr))
	}
	if statusErr := new(status); errors.As(err, &statusErr) {
		return statusErr
	}
	if grpcStatus, ok := grpcstatus.FromError(err); ok {
		return From(grpcStatus)
	}
	return ErrUnknown.With(Wrap(err))
}

func fromRpcStatus(grpcProto *rpcstatus.Status) Api {
	st := &status{
		err: &interstatusx.Error{
			GrpcStatus: &rpcstatus.Status{
				Code:    grpcProto.GetCode(),
				Message: grpcProto.GetMessage(),
			},
		},
	}
	for _, value := range grpcProto.GetDetails() {
		switch {
		case value.MessageIs(&interstatusx.Cause{}):
			_ = value.UnmarshalTo(st.err.Cause)
		case value.MessageIs(&interstatusx.Detail{}):
			_ = value.UnmarshalTo(st.err.Detail)
		case value.MessageIs(&httpstatus.HttpResponse{}):
			_ = value.UnmarshalTo(st.err.HttpStatus)
		default:
			st.err.GrpcStatus.Details = append(st.err.GrpcStatus.Details, value)
		}
	}
	return st
}
