package httpx

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/transportx"
	"net/http"
)

const (
	// HttpServer is the name of the http server transport.
	HttpServer = "http.server"

	// HttpClient is the name of the http client transport.
	HttpClient = "http.client"
)

func EndpointInjector(name string) httptransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		return endpointx.InjectName(ctx, name)
	}
}

func TransportInjector(name string) httptransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		return transportx.InjectName(ctx, name)
	}
}

func OutgoingMetadata(ctx context.Context, request *http.Request) context.Context {
	md, ok := metadatax.FromOutgoingContext(ctx)
	if !ok {
		return ctx
	}
	for _, key := range md.Keys() {
		for _, value := range md.Values(key) {
			request.Header.Add(key, value)
		}
	}
	return ctx
}
