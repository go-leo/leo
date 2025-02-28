package coder

import (
	"context"
	"errors"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/leo/v3/statusx"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"net/url"
)

// ===== server request decoder =====

func DecodeMessageFromRequest(ctx context.Context, r *http.Request, req proto.Message, unmarshalOptions protojson.UnmarshalOptions) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := unmarshalOptions.Unmarshal(data, req); err != nil {
		return err
	}
	return nil
}

func DecodeHttpBodyFromRequest(ctx context.Context, r *http.Request, body *httpbody.HttpBody) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	body.Data = data
	body.ContentType = r.Header.Get(ContentTypeKey)
	return nil
}

func DecodeHttpRequestFromRequest(ctx context.Context, r *http.Request, req *rpchttp.HttpRequest) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	req.Method = r.Method
	req.Uri = r.URL.String()
	for key, values := range r.Header {
		for _, value := range values {
			req.Headers = append(req.Headers, &rpchttp.HttpHeader{Key: key, Value: value})
		}
	}
	req.Body = data
	return nil
}

type FormGetter[T any] func(form url.Values, key string) (T, error)

func DecodeForm[T any](pre error, form url.Values, key string, f FormGetter[T]) (T, error) {
	return errorx.Break[T](pre)(func() (T, error) { return f(form, key) })
}

// ===== client response decoder =====

// DecodeMessageFromResponse decodes the proto.Message from the http.Response.
func DecodeMessageFromResponse(ctx context.Context, r *http.Response, resp proto.Message, unmarshalOptions protojson.UnmarshalOptions) (err error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Join(err, r.Body.Close())
	}
	if err := unmarshalOptions.Unmarshal(data, resp); err != nil {
		return errors.Join(err, r.Body.Close())
	}
	return r.Body.Close()
}

// DecodeHttpBodyFromResponse decodes the httpbody.HttpBody from the http.Response.
func DecodeHttpBodyFromResponse(ctx context.Context, r *http.Response, resp *httpbody.HttpBody) error {
	resp.ContentType = r.Header.Get("Content-Type")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	resp.Data = body
	return nil
}

// DecodeHttpResponseFromResponse decodes the http.HttpResponse from the http.Response.
func DecodeHttpResponseFromResponse(ctx context.Context, r *http.Response, resp *rpchttp.HttpResponse) error {
	resp.Status = int32(r.StatusCode)
	resp.Reason = http.StatusText(r.StatusCode)
	resp.Headers = make([]*rpchttp.HttpHeader, 0, len(r.Header))
	for key, values := range r.Header {
		for _, value := range values {
			elems := &rpchttp.HttpHeader{
				Key:   key,
				Value: value,
			}
			resp.Headers = append(resp.Headers, elems)
		}
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.Join(err, r.Body.Close())
	}
	resp.Body = data
	return r.Body.Close()
}

func DecodeErrorFromResponse(ctx context.Context, r *http.Response) error {
	st, _ := statusx.From(r)
	return st
}
