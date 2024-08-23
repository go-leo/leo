package statusx

import (
	"context"
	"errors"
	"fmt"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"net/url"
	"strconv"
)

// FromGrpcStatus returns an error representing s.
func FromGrpcStatus(s *grpcstatus.Status) ErrorApi {
	// s may be nil
	if s == nil {
		return nil
	}

	var errs []error
	var cause *interstatusx.Cause
	var detail *interstatusx.Detail
	var httpProto *httpstatus.HttpResponse
	details := make([]*anypb.Any, 0, len(s.Details()))
	for _, value := range s.Details() {
		switch item := value.(type) {
		case *interstatusx.Cause:
			// cause info
			cause = item
		case *interstatusx.Detail:
			// detail info
			detail = item
		case *httpstatus.HttpResponse:
			// http info
			httpProto = item
		case error:
			// error
			errs = append(errs, item)
		default:
			// other details
			message, ok := item.(proto.Message)
			if !ok {
				panic(fmt.Errorf("statusx: failed to convert value to proto.Message, %#v", item))
			}
			elem, err := anypb.New(message)
			if err != nil {
				panic(err)
			}
			details = append(details, elem)
		}
	}
	if len(errs) > 0 {
		panic(errors.Join(errs...))
	}
	return &status{
		err: &interstatusx.Error{
			Cause:      cause,
			Detail:     detail,
			HttpStatus: httpProto,
			GrpcStatus: &rpcstatus.Status{
				Code:    int32(s.Code()),
				Message: s.Message(),
				Details: details,
			},
		},
	}
}

func FromGrpcError(err error) ErrorApi {
	if err == nil {
		return nil
	}
	grpcStatus, _ := grpcstatus.FromError(err)
	return FromGrpcStatus(grpcStatus)
}

func From(obj any) ErrorApi {
	if obj == nil {
		return nil
	}
	switch err := obj.(type) {
	case *interstatusx.Error:
		return &status{err: err}
	case error:
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
			return FromGrpcStatus(grpcStatus)
		}
		return ErrUnknown.With(Wrap(err))
	}
	return ErrUnknown
}
