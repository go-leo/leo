package httpx

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	"google.golang.org/grpc/status"
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

func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	st, ok := status.FromError(err)
	if !ok {
		httptransport.DefaultErrorEncoder(ctx, err, w)
		return
	}
	statusx.Marshal(ctx, statusx.GetEncoder(), st, w)
}
