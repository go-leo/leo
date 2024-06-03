package statusx

import (
	"context"
	"github.com/go-leo/gox/encodingx/jsonx"
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	grpcstatus "google.golang.org/grpc/status"
	"net/http"
	"sync"
)

type Encoder interface {
	Encode(ctx context.Context, grpcStatus *grpcstatus.Status) (int, http.Header, string, []byte)
}

func Marshal(ctx context.Context, encoder Encoder, grpcStatus *grpcstatus.Status, w http.ResponseWriter) {
	statusCode, header, contentType, body := encoder.Encode(ctx, grpcStatus)
	w.Header().Set("Content-Type", contentType)
	for key, values := range header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

var defaultEncoder Encoder = encoder{}
var defaultEncoderLocker sync.RWMutex

func GetEncoder() Encoder {
	var m Encoder
	defaultEncoderLocker.RLock()
	m = defaultEncoder
	defaultEncoderLocker.RUnlock()
	return m
}

func SetEncoder(m Encoder) {
	defaultEncoderLocker.Lock()
	defaultEncoder = m
	defaultEncoderLocker.Unlock()
}

var _ Encoder = (*encoder)(nil)

type encoder struct{}

func (e encoder) Encode(ctx context.Context, grpcStatus *grpcstatus.Status) (int, http.Header, string, []byte) {
	grpcProto, httpProto := Proto(grpcStatus)
	if grpcProto == nil {
		return http.StatusOK, nil, "application/json; charset=utf-8", []byte("{}")
	}
	if httpProto == nil {
		httpProto = &httpstatus.Status{Code: HTTPStatusFromCode(grpcProto.GetCode())}
	}

	header := make(http.Header)
	for _, h := range HttpHeader(grpcStatus) {
		header.Add(h.Key, h.Value)
	}

	body, _ := jsonx.Marshal(grpcProto)

	return int(httpProto.Code), header, "application/json; charset=utf-8", body
}
