package httpx

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/sdx/stain"
	"net/http"
)

const (
	// HttpServer is the name of the http server transport.
	HttpServer = "http.server"
	// HttpClient is the name of the http client transport.
	HttpClient = "http.client"
)

const (
	kStainKey = "X-Leo-Stain"
)

func EndpointInjector(name string) httptransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		return endpointx.InjectName(ctx, name)
	}
}

func ServerTransportInjector(ctx context.Context, request *http.Request) context.Context {
	return endpointx.InjectName(ctx, HttpServer)
}

func OutgoingMetadataInjector(ctx context.Context, request *http.Request) context.Context {
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

func IncomingMetadataInjector(ctx context.Context, request *http.Request) context.Context {
	return metadatax.NewIncomingContext(ctx, metadatax.FromHttpHeader(request.Header))
}

type targetKey struct{}

func InjectTarget(ctx context.Context, target string) context.Context {
	return context.WithValue(ctx, targetKey{}, target)
}

func OutgoingStainInjector(ctx context.Context, request *http.Request) context.Context {
	color, ok := stain.ExtractColor(ctx)
	if !ok {
		return ctx
	}
	request.Header.Set(kStainKey, color)
	return ctx
}

func IncomingStainInjector(ctx context.Context, request *http.Request) context.Context {
	values := request.Header.Values(kStainKey)
	if values == nil || len(values) == 0 {
		return ctx
	}
	return stain.InjectColor(ctx, values[0])
}
