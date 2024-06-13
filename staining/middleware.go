package staining

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
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
	ctx = sdx.InjectColors(ctx, sdx.ParseColors(md.Values(key)))
	return next(ctx, request)
}

func handleOutgoing(ctx context.Context, request any, next endpoint.Endpoint, key string) (any, error) {
	colors, ok := sdx.ExtractColors(ctx)
	if !ok {
		return next(ctx, request)
	}
	ctx = metadatax.AppendToOutgoingContext(ctx, metadatax.FromMap(map[string][]string{key: colors.Pairs()}))
	return next(ctx, request)
}