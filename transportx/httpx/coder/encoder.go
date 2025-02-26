package coder

import (
	"bytes"
	"context"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

type ResponseTransformer func(ctx context.Context, resp proto.Message) proto.Message

func DefaultResponseTransformer(ctx context.Context, resp proto.Message) proto.Message { return resp }

// ===== server response encoder =====

// EncodeMessageToResponse encodes the proto.Message to the http.ResponseWriter.
func EncodeMessageToResponse(ctx context.Context, w http.ResponseWriter, resp proto.Message, marshalOptions protojson.MarshalOptions) error {
	w.Header().Set(ContentTypeKey, JsonContentType)
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
	w.Header().Set(ContentTypeKey, resp.GetContentType())
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

// ===== client request encoder =====

// EncodeMessageToRequest encodes the proto.Message to the io.Reader.
func EncodeMessageToRequest(ctx context.Context, req proto.Message, marshalOptions protojson.MarshalOptions) (io.Reader, string, error) {
	var bodyBuf bytes.Buffer
	data, err := marshalOptions.Marshal(req)
	if err != nil {
		return nil, "", err
	}
	if _, err = bodyBuf.Write(data); err != nil {
		return nil, "", err
	}
	return &bodyBuf, JsonContentType, nil
}

func EncodeHttpBodyToRequest(ctx context.Context, req *httpbody.HttpBody) (io.Reader, string) {
	return bytes.NewReader(req.GetData()), req.GetContentType()
}
