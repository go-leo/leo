package httpx

import (
	"context"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/leo/v3/transportx/httpx/internal/common"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"net/url"
)

func RequestDecoder(ctx context.Context, r *http.Request, req proto.Message, unmarshalOptions protojson.UnmarshalOptions) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := unmarshalOptions.Unmarshal(data, req); err != nil {
		return err
	}
	return nil
}

func HttpBodyDecoder(ctx context.Context, r *http.Request, body *httpbody.HttpBody) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	body.Data = data
	body.ContentType = r.Header.Get(common.ContentTypeKey)
	return nil
}

func HttpRequestDecoder(ctx context.Context, r *http.Request, request *rpchttp.HttpRequest) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	request.Method = r.Method
	request.Uri = r.URL.String()
	for key, values := range r.Header {
		for _, value := range values {
			request.Headers = append(request.Headers, &rpchttp.HttpHeader{Key: key, Value: value})
		}
	}
	request.Body = data
	return nil
}

type FormGetter[T any] func(form url.Values, key string) (T, error)

func FormDecoder[T any](pre error, form url.Values, key string, f FormGetter[T]) (T, error) {
	return errorx.Break[T](pre)(func() (T, error) { return f(form, key) })
}
