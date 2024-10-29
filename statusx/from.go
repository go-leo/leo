package statusx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
	"net/url"
	"strconv"
)

func From(obj any) Error {
	switch err := obj.(type) {
	case *status:
		return err
	case Error:
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
	case codes.Code:
		return fromGrpcCode(err)
	case int:
		return fromHttpCode(err)
	case error:
		return fromError(err)
	default:
		return ErrUnknown
	}
}

func FromGrpcError(err error) Error {
	if err == nil {
		return nil
	}
	if grpcStatus, ok := grpcstatus.FromError(err); ok {
		return fromRpcStatus(grpcStatus.Proto())
	}
	return ErrUnknown.With(Wrap(err))
}

func fromError(err error) Error {
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
	return fromDefaultErrorEncoder(err)
}

func fromDefaultErrorEncoder(err error) Error {
	statusCode := http.StatusInternalServerError
	if sc, ok := err.(httptransport.StatusCoder); ok {
		statusCode = sc.StatusCode()
	}
	statusErr := fromHttpCode(statusCode)

	// body
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			contentType, body = "application/json; charset=utf-8", jsonBody
		}
	}

	// header
	header := make([]*httpstatus.HttpHeader, 0)
	header = append(header, &httpstatus.HttpHeader{
		Key:   "Content-Type",
		Value: contentType,
	})
	if headerer, ok := err.(httptransport.Headerer); ok {
		for key, values := range headerer.Headers() {
			for _, value := range values {
				header = append(header, &httpstatus.HttpHeader{
					Key:   key,
					Value: value,
				})
			}
		}
	}
	return statusErr.With(HttpHeader(header...), HttpBody(wrapperspb.Bytes(body)))
}

func fromRpcStatus(grpcProto *rpcstatus.Status) Error {
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
			st.err.Cause = new(interstatusx.Cause)
			_ = value.UnmarshalTo(st.err.Cause)
		case value.MessageIs(&interstatusx.Detail{}):
			st.err.Detail = new(interstatusx.Detail)
			_ = value.UnmarshalTo(st.err.Detail)
		case value.MessageIs(&httpstatus.HttpResponse{}):
			st.err.HttpStatus = new(httpstatus.HttpResponse)
			_ = value.UnmarshalTo(st.err.HttpStatus)
		default:
			st.err.GrpcStatus.Details = append(st.err.GrpcStatus.Details, value)
		}
	}
	return st
}

var kGrpcToHttpCode = map[codes.Code]Error{
	codes.OK:                 OK,
	kFailedCode:              Failed,
	codes.Canceled:           ErrCanceled,
	codes.Unknown:            ErrUnknown,
	codes.InvalidArgument:    ErrInvalidArgument,
	codes.DeadlineExceeded:   ErrDeadlineExceeded,
	codes.NotFound:           ErrNotFound,
	codes.AlreadyExists:      ErrAlreadyExists,
	codes.PermissionDenied:   ErrPermissionDenied,
	codes.ResourceExhausted:  ErrResourceExhausted,
	codes.FailedPrecondition: ErrFailedPrecondition,
	codes.Aborted:            ErrAborted,
	codes.OutOfRange:         ErrOutOfRange,
	codes.Unimplemented:      ErrUnimplemented,
	codes.Internal:           ErrInternal,
	codes.Unavailable:        ErrUnavailable,
	codes.DataLoss:           ErrDataLoss,
	codes.Unauthenticated:    ErrUnauthenticated,
}

// fromGrpcCode converts a gRPC status code to Error.
func fromGrpcCode(code codes.Code) Error {
	statusErr, ok := kGrpcToHttpCode[code]
	if ok {
		return statusErr
	}
	return ErrUnknown
}

var kHttpToGrpcCode = map[int]Error{
	http.StatusBadRequest:         ErrInternal,
	http.StatusUnauthorized:       ErrUnauthenticated,
	http.StatusForbidden:          ErrPermissionDenied,
	http.StatusNotFound:           ErrUnimplemented,
	http.StatusTooManyRequests:    ErrUnavailable,
	http.StatusBadGateway:         ErrUnavailable,
	http.StatusServiceUnavailable: ErrUnavailable,
	http.StatusGatewayTimeout:     ErrUnavailable,
}

// fromHttpCode converts an HTTP status code to Error.
// See: [HTTP to gRPC Status Code Mapping]: https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
func fromHttpCode(code int) Error {
	statusErr, ok := kHttpToGrpcCode[code]
	if ok {
		return statusErr
	}
	return ErrUnknown
}
