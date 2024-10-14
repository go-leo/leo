package httpx

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/httpx/internal"
	"net/http"
	"time"
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

type timeoutHolderKey struct{}

type timeoutHolder struct {
	cancelFunc context.CancelFunc
}

func TimeoutController(ctx context.Context, request *http.Request) context.Context {
	name, _ := transportx.ExtractName(ctx)
	if name == HttpServer {
		return serverTimeoutController(ctx, request)
	} else if name == HttpClient {
		return clientTimeoutController(ctx, request)
	}
	return ctx
}

func clientTimeoutController(ctx context.Context, request *http.Request) context.Context {
	if deadline, ok := ctx.Deadline(); ok {
		request.Header.Set(kTimeoutKey, internal.EncodeDuration(time.Until(deadline)))
	}
	return ctx
}

func serverTimeoutController(ctx context.Context, request *http.Request) context.Context {
	if value := request.Header.Get(kTimeoutKey); value != "" {
		timeout, err := internal.DecodeTimeout(value)
		if err != nil {
			_ = logx.L().Log("error", err)
		}
		ctx, cancelFunc := context.WithTimeout(ctx, timeout)
		_ = cancelFunc
		return context.WithValue(ctx, timeoutHolderKey{}, &timeoutHolder{cancelFunc: cancelFunc})
	}
	ctx, cancelFunc := context.WithCancel(ctx)
	return context.WithValue(ctx, timeoutHolderKey{}, &timeoutHolder{cancelFunc: cancelFunc})
}

func CancelInvoker(ctx context.Context, code int, r *http.Request) {
	holder, ok := ctx.Value(timeoutHolderKey{}).(*timeoutHolder)
	if ok {
		holder.cancelFunc()
	}
}
