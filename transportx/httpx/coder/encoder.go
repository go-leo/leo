package coder

import (
	"context"
	"github.com/go-leo/leo/v3/transportx/httpx/internal/common"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

type ResponseTransformer func(ctx context.Context, resp proto.Message) proto.Message

func DefaultResponseTransformer(ctx context.Context, resp proto.Message) proto.Message { return resp }

// ===== server response encoder =====

// EncodeResponseToResponse encodes the proto.Message to the http.ResponseWriter.
func EncodeResponseToResponse(ctx context.Context, w http.ResponseWriter, resp proto.Message, marshalOptions protojson.MarshalOptions) error {
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

// EncodeHttpBodyToResponse encodes the httpbody.HttpBody to the http.ResponseWriter.
func EncodeHttpBodyToResponse(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	w.Header().Set(common.ContentTypeKey, resp.GetContentType())
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp.GetData()); err != nil {
		return err
	}
	return nil
}

// EncodeHttpResponseToResponse encodes the http.HttpResponse to the http.ResponseWriter.
func EncodeHttpResponseToResponse(ctx context.Context, w http.ResponseWriter, resp *rpchttp.HttpResponse) error {
	for _, header := range resp.GetHeaders() {
		w.Header().Add(header.GetKey(), header.GetValue())
	}
	w.WriteHeader(int(resp.GetStatus()))
	if _, err := w.Write(resp.GetBody()); err != nil {
		return err
	}
	return nil
}
