package httpx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/gox/encodingx/jsonx"
	"github.com/go-leo/gox/strconvx"
	"github.com/go-leo/leo/v3/statusx"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
	"io"
	"net/http"
	"strings"
	"sync"
)

var (
	// defaultStatusEncoder is the default status encoder
	defaultStatusEncoder statusx.Encoder = statusEncoder{}

	// defaultStatusDecoder is the default status decoder
	defaultStatusDecoder statusx.Decoder = statusDecoder{}

	// defaultStatusLocker is the default status locker
	defaultStatusLocker sync.RWMutex
)

const (
	kContentTypeKey  = "Content-Type"
	kErrorEncoderKey = "X-Leo-Error-Encoder"
	kGrpcCodeKey     = "X-Leo-Grpc-Code"
	kGrpcMsgKey      = "X-Leo-Grpc-Msg"
	kStatusKeysKey   = "X-Leo-Status-Keys"

	kKitDefaultValue = "kit-default"
	kLeoDefaultValue = "leo-default"
	kJSONContentType = "application/json; charset=utf-8"
)

func GetStatusEncoder() statusx.Encoder {
	var c statusx.Encoder
	defaultStatusLocker.RLock()
	c = defaultStatusEncoder
	defaultStatusLocker.RUnlock()
	return c
}

func SetStatusEncoder(c statusx.Encoder) {
	defaultStatusLocker.Lock()
	defaultStatusEncoder = c
	defaultStatusLocker.Unlock()
}

func GetStatusDecoder() statusx.Decoder {
	var c statusx.Decoder
	defaultStatusLocker.RLock()
	c = defaultStatusDecoder
	defaultStatusLocker.RUnlock()
	return c
}

func SetStatusDecoder(c statusx.Decoder) {
	defaultStatusLocker.Lock()
	defaultStatusDecoder = c
	defaultStatusLocker.Unlock()
}

type statusEncoder struct{}

func (e statusEncoder) Encode(ctx context.Context, err *statusx.Error) (int, http.Header, []byte) {
	grpcProto, httpProto := err.Proto()

	header := httpProto.HttpHeader()
	header.Set(kContentTypeKey, kJSONContentType)
	header.Set(kErrorEncoderKey, kLeoDefaultValue)

	body := &bytes.Buffer{}
	bodyProto := err.HttpBody()
	if bodyProto != nil {
		header.Set(kGrpcCodeKey, strconvx.FormatInt(grpcProto.GetCode(), 10))
		header.Set(kGrpcMsgKey, grpcProto.GetMessage())
		if bodyAny, err := anypb.New(bodyProto); err != nil {
			_ = jsonx.NewEncoder(body).Encode(bodyProto)
		} else {
			_ = jsonx.NewEncoder(body).Encode(bodyAny)
		}
	} else {
		_ = jsonx.NewEncoder(body).Encode(grpcProto)
	}

	return int(httpProto.GetCode()), header, body.Bytes()
}

type statusDecoder struct {
}

func (d statusDecoder) Decode(ctx context.Context, status int, header http.Header, body []byte) *statusx.Error {
	errType := header.Get(kErrorEncoderKey)
	if !strings.EqualFold(errType, kLeoDefaultValue) {
		return nil
	}

	keys := header.Values(kStatusKeysKey)
	headers := make([]*httpstatus.Header, 0, len(keys))
	for _, key := range keys {
		for _, values := range header.Values(key) {
			headers = append(headers, &httpstatus.Header{Key: key, Value: values})
		}
	}

	grpcCode := header.Get(kGrpcCodeKey)
	if grpcCode != "" {
		grpcMsg := header.Get(kGrpcMsgKey)
		code, _ := strconvx.ParseUint[codes.Code](grpcCode, 10, 32)
		if status == http.StatusOK && code == codes.Unknown {
			code = statusx.FailedCode
		}

		bodyAny := &anypb.Any{}
		if err := jsonx.Unmarshal(body, &bodyAny); err == nil {
			bodyProto, err := bodyAny.UnmarshalNew()
			if err == nil {
				return statusx.NewError(code, grpcMsg).WithHttpHeader(headers...).WithHttpBody(bodyProto)
			}
		}

		bodyStruct := &structpb.Struct{}
		if err := jsonx.Unmarshal(body, &bodyStruct); err == nil {
			return statusx.NewError(code, grpcMsg).WithHttpHeader(headers...).WithHttpBody(bodyStruct)
		}

		return statusx.NewError(code, grpcMsg).WithHttpHeader(headers...)
	}
	var grpcProto *rpcstatus.Status
	_ = jsonx.Unmarshal(body, &grpcProto)
	code := codes.Code(grpcProto.GetCode())
	if status == http.StatusOK && code == codes.Unknown {
		code = statusx.FailedCode
	}
	return statusx.NewError(code, grpcProto.GetMessage()).WithHttpHeader(headers...)
}

func IsErrorResponse(r *http.Response) bool {
	return r.Header.Get(kErrorEncoderKey) != ""
}

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	var statusErr *statusx.Error
	if !errors.As(err, &statusErr) {
		w.Header().Set(kErrorEncoderKey, kKitDefaultValue)
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}

	if statusErr == nil {
		w.Header().Set(kErrorEncoderKey, kKitDefaultValue)
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}

	encoder := GetStatusEncoder()
	if encoder == nil {
		w.Header().Set(kErrorEncoderKey, kKitDefaultValue)
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}

	statusCode, header, body := encoder.Encode(ctx, statusErr)
	for key := range header {
		for _, value := range header.Values(key) {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

func ErrorDecoder(ctx context.Context, r *http.Response) error {
	errHeader := r.Header.Get(kErrorEncoderKey)

	if strings.EqualFold(errHeader, kKitDefaultValue) {
		body, _ := io.ReadAll(r.Body)
		return &ResponseError{statusCode: r.StatusCode, header: r.Header, body: body}
	}

	decoder := GetStatusDecoder()
	if decoder == nil {
		body, _ := io.ReadAll(r.Body)
		return &ResponseError{statusCode: r.StatusCode, header: r.Header, body: body}
	}

	body, _ := io.ReadAll(r.Body)

	statusErr := decoder.Decode(ctx, r.StatusCode, r.Header, body)
	if statusErr != nil {
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
	return fmt.Sprintf("rpc error: status-code = %d, body = %s", e.statusCode, string(e.body))
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
