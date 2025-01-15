package httpx

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/sdx/stain"
	"github.com/go-leo/leo/v3/transportx/internal"
	"net/http"
	"time"
)

const (
	// HttpServer is the name of the http server transport.
	HttpServer = "http.server"
	// HttpClient is the name of the http client transport.
	HttpClient = "http.client"
)

const (
	kTimeoutKey = "X-Leo-Timeout"
	kStainKey   = "X-Leo-Stain"
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

type timeLimiterKey struct{}

func OutgoingTimeLimitInjector(ctx context.Context, request *http.Request) context.Context {
	if deadline, ok := ctx.Deadline(); ok {
		request.Header.Set(kTimeoutKey, internal.EncodeDuration(time.Until(deadline)))
	}
	return ctx
}

func IncomingTimeLimitInjector(ctx context.Context, request *http.Request) context.Context {
	if value := request.Header.Get(kTimeoutKey); value != "" {
		timeout, err := internal.DecodeTimeout(value)
		if err != nil {
			_ = logx.L().Log("error", err)
		}
		ctx, cancelFunc := context.WithTimeout(ctx, timeout)
		return context.WithValue(ctx, timeLimiterKey{}, cancelFunc)
	}
	ctx, cancelFunc := context.WithCancel(ctx)
	return context.WithValue(ctx, timeLimiterKey{}, cancelFunc)
}

func CancelInvoker(ctx context.Context, code int, r *http.Request) {
	cancelFunc, ok := ctx.Value(timeLimiterKey{}).(context.CancelFunc)
	if ok {
		cancelFunc()
	}
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
