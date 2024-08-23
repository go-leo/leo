package httpx

import (
	"context"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	interstatusx "github.com/go-leo/leo/v3/internal/statusx"
	"github.com/go-leo/leo/v3/statusx"
	"golang.org/x/exp/maps"
	httpstatus "google.golang.org/genproto/googleapis/rpc/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"net/http"
	"strings"
	"sync"
)

type StatusEncoder interface {
	Encode(ctx context.Context, err statusx.ErrorApi) (int, http.Header, []byte)
}

type StatusDecoder interface {
	Decode(ctx context.Context, status int, header http.Header, body []byte) statusx.ErrorApi
}

var (
	// defaultStatusEncoder is the default status encoder
	defaultStatusEncoder StatusEncoder = defaultStatusCoder{}

	// defaultStatusDecoder is the default status decoder
	defaultStatusDecoder StatusDecoder = defaultStatusCoder{}

	// defaultStatusLocker is the default status locker
	defaultStatusLocker sync.RWMutex
)

const (
	kStatusKey             = "X-Leo-Status"
	kStatusKeysKey         = "X-Leo-Status-Keys"
	kStatusDetailsKey      = "X-Leo-Status-Details"
	kStatusCauseMessageKey = "X-Leo-Status-Cause-Message"
	kStatusCauseErrorKey   = "X-Leo-Status-Cause-Error"
	kStatusDetailKey       = "X-Leo-Status-Detail"
	kStatusGrpcKey         = "X-Leo-Status-Grpc"

	kKitDefaultValue = "kit-default"
	kLeoDefaultValue = "leo-default"
)

func GetStatusEncoder() StatusEncoder {
	var c StatusEncoder
	defaultStatusLocker.RLock()
	c = defaultStatusEncoder
	defaultStatusLocker.RUnlock()
	return c
}

func SetStatusEncoder(c StatusEncoder) {
	defaultStatusLocker.Lock()
	defaultStatusEncoder = c
	defaultStatusLocker.Unlock()
}

func GetStatusDecoder() StatusDecoder {
	var c StatusDecoder
	defaultStatusLocker.RLock()
	c = defaultStatusDecoder
	defaultStatusLocker.RUnlock()
	return c
}

func SetStatusDecoder(c StatusDecoder) {
	defaultStatusLocker.Lock()
	defaultStatusDecoder = c
	defaultStatusLocker.Unlock()
}

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	// check if it is a status error
	statusErr, ok := statusx.FromError(err)
	if !ok {
		// failed to convert to status error, use go-kit default encoder
		w.Header().Set(kStatusKey, kKitDefaultValue)
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}

	if statusErr == nil {
		// status is nil, use go-kit default encoder
		w.Header().Set(kStatusKey, kKitDefaultValue)
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}

	// get leo status encoder
	encoder := GetStatusEncoder()
	if encoder == nil {
		// failed to get leo status encoder, use go-kit default encoder
		w.Header().Set(kStatusKey, kKitDefaultValue)
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}

	// encode status error
	statusCode, header, body := encoder.Encode(ctx, statusErr)
	w.Header().Set(kStatusKey, kLeoDefaultValue)
	// write response
	for key := range header {
		for _, value := range header.Values(key) {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

func IsErrorResponse(r *http.Response) bool {
	return r.Header.Get(kStatusKey) != ""
}

// ErrorDecoder decode error from http response
func ErrorDecoder(ctx context.Context, r *http.Response) error {
	// get error encoder
	errHeader := r.Header.Get(kStatusKey)
	body, _ := io.ReadAll(r.Body)
	// use go-kit default encoder
	if strings.EqualFold(errHeader, kKitDefaultValue) {
		return &ResponseError{statusCode: r.StatusCode, header: r.Header, body: body}
	}

	if strings.EqualFold(errHeader, kLeoDefaultValue) {
		decoder := GetStatusDecoder()
		if decoder == nil {
			// failed to get leo status encoder, use go-kit default encoder
			return &ResponseError{statusCode: r.StatusCode, header: r.Header, body: body}
		}
		if statusErr := decoder.Decode(ctx, r.StatusCode, r.Header, body); statusErr != nil {
			return statusErr
		}
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

type defaultStatusCoder struct{}

func (e defaultStatusCoder) Encode(ctx context.Context, errApi statusx.ErrorApi) (int, http.Header, []byte) {
	grpcProto, httpProto := errApi.Proto()

	headers := httpProto.GetHeaders()

	header := make(http.Header, len(headers))
	for _, h := range headers {
		header.Add(h.GetKey(), h.GetValue())
	}
	keys := maps.Keys(header)
	for _, key := range keys {
		header.Add(kStatusKeysKey, key)
	}

	// add cause error
	if cause := errApi.Unwrap(); cause != nil {
		if causeProto, ok := cause.(proto.Message); ok {
			causeAny, err := anypb.New(causeProto)
			if err != nil {
				panic(err)
			}
			causeAnyData, err := protojson.Marshal(causeAny)
			if err != nil {
				panic(err)
			}
			header.Set(kStatusCauseErrorKey, string(causeAnyData))
		} else {
			header.Set(kStatusCauseMessageKey, cause.Error())
		}
	}

	// add detail info
	detail := &interstatusx.Detail{
		ErrorInfo:           errApi.ErrorInfo(),
		RetryInfo:           errApi.RetryInfo(),
		DebugInfo:           errApi.DebugInfo(),
		QuotaFailure:        errApi.QuotaFailure(),
		PreconditionFailure: errApi.PreconditionFailure(),
		BadRequest:          errApi.BadRequest(),
		RequestInfo:         errApi.RequestInfo(),
		ResourceInfo:        errApi.ResourceInfo(),
		Help:                errApi.Help(),
		LocalizedMessage:    errApi.LocalizedMessage(),
	}
	detailData, err := protojson.Marshal(detail)
	if err != nil {
		panic(err)
	}
	header.Set(kStatusDetailKey, string(detailData))

	// add grpc status
	grpcProtoData, err := protojson.Marshal(grpcProto)
	if err != nil {
		panic(err)
	}
	header.Set(kStatusGrpcKey, string(grpcProtoData))

	return int(httpProto.GetStatus()), header, httpProto.GetBody()
}

func (d defaultStatusCoder) Decode(ctx context.Context, status int, header http.Header, body []byte) statusx.ErrorApi {
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
