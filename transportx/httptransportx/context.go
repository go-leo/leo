package httptransportx

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/endpointx"
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

func ServerTransportInjector(ctx context.Context, request *http.Request) context.Context {
	return endpointx.InjectName(ctx, HttpServer)
}

type targetKey struct{}

func InjectTarget(ctx context.Context, target string) context.Context {
	return context.WithValue(ctx, targetKey{}, target)
}
