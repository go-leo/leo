package httpx

import (
	"context"
	"encoding/json"
	"github.com/go-leo/leo/v3/transportx/httpx/internal/common"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

type ErrorEncoder func(ctx context.Context, err error, w http.ResponseWriter)

type ResponseTransformer func(ctx context.Context, resp proto.Message) proto.Message

func DefaultResponseTransformer(ctx context.Context, resp proto.Message) proto.Message { return resp }

func DefaultErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			contentType, body = common.JsonContentType, jsonBody
		}
	}
	w.Header().Set(common.ContentTypeKey, contentType)
	if headerer, ok := err.(interface{ Headers() http.Header }); ok {
		for k, values := range headerer.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(interface{ StatusCode() int }); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	_, err = w.Write(body)
	if err != nil {
		log.Default().Println("gors: response write error: ", err)
	}
}

func ResponseEncoder(ctx context.Context, w http.ResponseWriter, resp proto.Message, marshalOptions protojson.MarshalOptions) error {
	w.Header().Set(common.ContentTypeKey, common.JsonContentType)
	w.WriteHeader(http.StatusOK)
	data, err := marshalOptions.Marshal(resp)
	if err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}

func HttpBodyEncoder(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	w.Header().Set(common.ContentTypeKey, resp.GetContentType())
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp.GetData()); err != nil {
		return err
	}
	return nil
}

func HttpResponseEncoder(ctx context.Context, w http.ResponseWriter, resp *rpchttp.HttpResponse) error {
	for _, header := range resp.GetHeaders() {
		w.Header().Add(header.GetKey(), header.GetValue())
	}
	w.WriteHeader(int(resp.GetStatus()))
	if _, err := w.Write(resp.GetBody()); err != nil {
		return err
	}
	return nil
}
