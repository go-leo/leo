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
)

// FromGrpcStatus returns an error representing s.
func FromGrpcStatus(s *grpcstatus.Status) *Error {
	if s == nil {
		return nil
	}
	details := make([]*anypb.Any, 0, len(s.Details()))
	var errs []error
	var cause *interstatusx.Cause
	var detail *interstatusx.Detail
	var httpProto *httpstatus.HttpResponse
	for _, value := range s.Details() {
		switch item := value.(type) {
		case *interstatusx.Cause:
			cause = item
		case *interstatusx.Detail:
			detail = item
		case *httpstatus.HttpResponse:
			httpProto = item
		case error:
			errs = append(errs, item)
		default:
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
	return &Error{
		e: &interstatusx.Error{
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

func FromError(err error) (*Error, bool) {
	if err == nil {
		return nil, true
	}
	var statusErr *Error
	if errors.As(err, &statusErr) {
		return statusErr, true
	}
	grpcStatus, ok := grpcstatus.FromError(err)
	if ok {
		return FromGrpcStatus(grpcStatus), true
	}
	return nil, false
}

func FromGrpcError(err error) *Error {
	grpcStatus, _ := grpcstatus.FromError(err)
	return FromGrpcStatus(grpcStatus)
}

// FromContextError converts a context error to Error
func FromContextError(err error) *Error {
	if err == nil {
		return nil
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return ErrDeadlineExceeded.With(Wrap(err))
	}
	if errors.Is(err, context.Canceled) {
		return ErrCanceled.With(Wrap(err))
	}
	return ErrUnknown.With(Wrap(err))
}
