package staining

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"strings"
)

var (
	ErrMissMetadata = statusx.ErrInvalidArgument.WithMessage("missing metadata")
)

// Middleware is a middleware that get color info from incoming metadata and injects the color info into the context
// or append the color info to outgoing metadata.
func Middleware(key string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			name, ok := transportx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			switch name {
			case grpcx.GrpcServer, httpx.HttpServer:
				return handleIncoming(ctx, request, next, key)
			case grpcx.GrpcClient, httpx.HttpClient:
				return handleOutgoing(ctx, request, next, key)
			}
			return next(ctx, request)
		}
	}
}

func handleIncoming(ctx context.Context, request any, next endpoint.Endpoint, key string) (any, error) {
	md, ok := metadatax.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissMetadata
	}
	values := md.Values(key)
	colors := make([]*transportx.Color, 0, len(values))
	for _, value := range values {
		pair := strings.Split(value, "=")
		if len(pair) != 2 {
			continue
		}
		colorValues := strings.Split(pair[1], ",")
		colors = append(colors, &transportx.Color{Service: pair[0], Colors: colorValues})
	}
	ctx = transportx.InjectColors(ctx, colors)
	return next(ctx, request)
}

func handleOutgoing(ctx context.Context, request any, next endpoint.Endpoint, key string) (any, error) {
	colors, ok := transportx.ExtractColors(ctx)
	if !ok {
		return next(ctx, request)
	}
	values := make([]string, 0, len(colors))
	for _, c := range colors {
		values = append(values, c.Service+"="+strings.Join(c.Colors, ","))
	}
	ctx = metadatax.AppendToOutgoingContext(ctx, metadatax.FromMap(map[string][]string{key: values}))
	return next(ctx, request)
}
