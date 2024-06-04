package httpx

import (
	"bytes"
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/gox/encodingx/jsonx"
	"github.com/go-leo/gox/strconvx"
	"github.com/go-leo/leo/v3/statusx"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
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

func (e statusEncoder) Encode(ctx context.Context, grpcStatus *grpcstatus.Status) (*httpstatus.Status, []*httpstatus.Header, []byte) {
	grpcProto, httpProto := statusx.Proto(grpcStatus)
	if httpProto == nil {
		httpProto = &httpstatus.Status{Code: statusx.HTTPStatusFromCode(grpcProto.GetCode())}
	}

	httpHeader := statusx.HttpHeader(grpcStatus)
	httpHeader = append(httpHeader, &httpstatus.Header{Key: kContentTypeKey, Value: kJSONContentType})
	httpHeader = append(httpHeader, &httpstatus.Header{Key: kErrorEncoderKey, Value: kLeoDefaultValue})

	body := &bytes.Buffer{}
	httpResult := statusx.HttpResult(grpcStatus)
	if httpResult != nil {
		httpHeader = append(httpHeader,
			&httpstatus.Header{Key: kGrpcCodeKey, Value: strconvx.FormatInt(grpcProto.GetCode(), 10)},
			&httpstatus.Header{Key: kGrpcMsgKey, Value: grpcProto.GetMessage()},
		)
		_ = jsonx.NewEncoder(body).Encode(httpResult)
	} else {
		_ = jsonx.NewEncoder(body).Encode(grpcProto)
	}

	return httpProto, httpHeader, body.Bytes()
}

type statusDecoder struct {
}

func (d statusDecoder) Decode(ctx context.Context, httpProto *httpstatus.Status, header []*httpstatus.Header, body []byte) *grpcstatus.Status {
	var errType string
	var grpcCode string
	var grpcMsg string
	for _, h := range header {
		key := h.GetKey()
		if strings.EqualFold(key, kErrorEncoderKey) {
			errType = h.GetValue()
		}
		if strings.EqualFold(key, kGrpcCodeKey) {
			grpcCode = h.Value
		}
		if strings.EqualFold(key, kGrpcMsgKey) {
			grpcMsg = h.Value
		}
	}
	if !strings.EqualFold(errType, kLeoDefaultValue) {
		return nil
	}
	if grpcCode != "" {
		code, _ := strconvx.ParseInt[int32](grpcCode, 10, 32)
		return statusx.WithDetails(grpcstatus.New(codes.Code(code), grpcMsg), httpProto)
	}
	var grpcProto *rpcstatus.Status
	_ = jsonx.Unmarshal(body, &grpcProto)
	return statusx.WithHttpStatus(grpcstatus.FromProto(grpcProto), httpProto)
}

func IsErrorResponse(r *http.Response) bool {
	return r.Header.Get(kErrorEncoderKey) != ""
}

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	grpcStatus, ok := grpcstatus.FromError(err)
	if !ok {
		w.Header().Set(kErrorEncoderKey, kKitDefaultValue)
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}

	if grpcStatus == nil {
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

	httpProto, header, body := encoder.Encode(ctx, grpcStatus)
	for _, h := range header {
		w.Header().Add(h.GetKey(), h.GetValue())
	}
	w.WriteHeader(int(httpProto.GetCode()))
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

	header := make([]*httpstatus.Header, 0, len(r.Header))
	for key, values := range r.Header {
		for _, value := range values {
			header = append(header, &httpstatus.Header{Key: key, Value: value})
		}
	}

	body, _ := io.ReadAll(r.Body)

	grpcStatus := decoder.Decode(ctx, &httpstatus.Status{Code: int32(r.StatusCode)}, header, body)

	if grpcStatus == nil {
		body, _ := io.ReadAll(r.Body)
		return &ResponseError{statusCode: r.StatusCode, header: r.Header, body: body}
	}
	return grpcStatus.Err()
}

type ResponseError struct {
	statusCode int
	header     http.Header
	body       []byte
}

func (e *ResponseError) Error() string {
	return "response error"
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
