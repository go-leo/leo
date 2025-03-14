package metadatax

import (
	"context"
	"net/http"
)

func HttpOutgoingInjector(ctx context.Context, request *http.Request) context.Context {
	incomingMD, ok := FromIncomingContext(ctx)
	if ok {
		for _, key := range incomingMD.Keys() {
			for _, value := range incomingMD.Values(key) {
				request.Header.Add(key, value)
			}
		}
	}
	md, ok := FromOutgoingContext(ctx)
	if ok {
		for _, key := range md.Keys() {
			for _, value := range md.Values(key) {
				request.Header.Add(key, value)
			}
		}
	}
	return ctx
}

func HttpIncomingInjector(ctx context.Context, request *http.Request) context.Context {
	return NewIncomingContext(ctx, FromHttpHeader(request.Header))
}

// AsHttpHeader Convert Metadata to http.Header
func AsHttpHeader(md Metadata) http.Header {
	res := http.Header{}
	for _, key := range md.Keys() {
		values := md.Values(key)
		for _, value := range values {
			res.Add(key, value)
		}
	}
	return res
}

// FromHttpHeader Convert http.Header to Metadata
//
// The keys should be in canonical form, as returned by http.CanonicalHeaderKey.
func FromHttpHeader(header http.Header) Metadata {
	return FromMap(header)
}
