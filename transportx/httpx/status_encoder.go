package httpx

import (
	"context"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	"github.com/go-leo/leo/v3/statusx"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"net/http"
	"strings"
)

func IsErrorResponse(r *http.Response) bool {
	return r.Header.Get(kStatusCoderKey) != ""
}

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	// check if it is a status error
	statusErr := statusx.From(err)
	if statusErr == nil {
		// status is nil, use go-kit default encoder
		w.Header().Set(kStatusCoderKey, kKitCoderValue)
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}

	// encode status error
	statusCode, header, body := encode(ctx, statusErr)
	w.Header().Set(kStatusCoderKey, kLeoCoderValue)
	// write response
	for key := range header {
		for _, value := range header.Values(key) {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

// ErrorDecoder decode error from http response
func ErrorDecoder(ctx context.Context, r *http.Response) error {
	// get error encoder
	errHeader := r.Header.Get(kStatusCoderKey)
	body, _ := io.ReadAll(r.Body)
	// use go-kit default encoder
	if strings.EqualFold(errHeader, kKitDefaultValue) {
		return &ResponseError{statusCode: r.StatusCode, header: r.Header, body: body}
	}
	if !strings.EqualFold(errHeader, kLeoDefaultValue) {
		return &ResponseError{statusCode: r.StatusCode, header: r.Header, body: body}
	}
	if statusErr := decode(ctx, r.StatusCode, r.Header, body); statusErr != nil {
		return statusErr
	}
	return &ResponseError{statusCode: r.StatusCode, header: r.Header, body: body}
}

type ResponseError struct {
	statusCode int
	header     http.Header
	body       []byte
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("httpx: status-code = %d", e.statusCode)
}

func (e *ResponseError) StatusCode() int {
	return e.statusCode
}

func (e *ResponseError) Header() http.Header {
	return e.header
}

func (e *ResponseError) Body() []byte {
	return e.body
}

func decode(ctx context.Context, status int, header http.Header, body []byte) statusx.Api {
	keys := header.Values(kStatusKeysKey)
	headers := make([]*httpstatus.HttpHeader, 0, len(keys))
	for _, key := range keys {
		for _, value := range header.Values(key) {
			headers = append(headers, &httpstatus.HttpHeader{Key: key, Value: value})
		}
	}

	var cause *interstatusx.Cause
	if message := header.Get(kStatusCauseMessageKey); message != "" {
		cause = &interstatusx.Cause{Cause: &interstatusx.Cause_Message{Message: wrapperspb.String(message)}}
	} else if causeProtoData := header.Get(kStatusCauseErrorKey); causeProtoData != "" {
		var causeAny anypb.Any
		if err := protojson.Unmarshal([]byte(causeProtoData), &causeAny); err != nil {
			panic(err)
		}
		cause = &interstatusx.Cause{Cause: &interstatusx.Cause_Error{Error: &causeAny}}
	}

	var detail *interstatusx.Detail
	if detailData := header.Get(kStatusDetailKey); detailData != "" {
		detail = &interstatusx.Detail{}
		if err := protojson.Unmarshal([]byte(detailData), detail); err != nil {
			panic(err)
		}
	}

	var grpcProto *rpcstatus.Status
	if grpcProtoData := header.Get(kStatusGrpcKey); grpcProtoData != "" {
		grpcProto = &rpcstatus.Status{}
		if err := protojson.Unmarshal([]byte(grpcProtoData), grpcProto); err != nil {
			panic(err)
		}
	}

	err := &interstatusx.Error{
		Cause:  cause,
		Detail: detail,
		HttpStatus: &httpstatus.HttpResponse{
			Status:  int32(status),
			Reason:  http.StatusText(status),
			Headers: headers,
			Body:    body,
		},
		GrpcStatus: grpcProto,
	}
	return statusx.From(err)
}
