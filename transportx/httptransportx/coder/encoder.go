package coder

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/statusx"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

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

func EncodeErrorToResponse(ctx context.Context, err error, w http.ResponseWriter) {
	st, ok := statusx.From(err)
	if !ok {
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}
	httptransport.DefaultErrorEncoder(ctx, st, w)
}

// ===== client request encoder =====

// EncodeMessageToRequest encodes the proto.Message to the io.Reader.
func EncodeMessageToRequest(ctx context.Context, req proto.Message, header http.Header, body io.Writer, marshalOptions protojson.MarshalOptions) error {
	data, err := marshalOptions.Marshal(req)
	if err != nil {
		return err
	}
	if _, err = body.Write(data); err != nil {
		return err
	}
	header.Set(ContentTypeKey, JsonContentType)
	return nil
}

func EncodeHttpBodyToRequest(ctx context.Context, req *httpbody.HttpBody, header http.Header, body io.Writer) error {
	if _, err := body.Write(req.GetData()); err != nil {
		return err
	}
	header.Set(ContentTypeKey, req.GetContentType())
	return nil
}

func EncodeHttpRequestToRequest(ctx context.Context, req *rpchttp.HttpRequest, header http.Header, body io.Writer) error {
	if _, err := body.Write(req.GetBody()); err != nil {
		return err
	}
	for _, httpHeader := range req.GetHeaders() {
		header.Add(httpHeader.GetKey(), httpHeader.GetValue())
	}
	return nil
}
