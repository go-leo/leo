package stainx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc/metadata"
	"net/http"
)

const (
	kStainKey = "X-Leo-Stain"
)

type colorKey struct{}

// ColorInjector injects the colors into the context.
func ColorInjector(ctx context.Context, color string) context.Context {
	return context.WithValue(ctx, colorKey{}, color)
}

// ColorExtractor extracts the colors from the context.
func ColorExtractor(ctx context.Context) (string, bool) {
	color, ok := ctx.Value(colorKey{}).(string)
	return color, ok
}

func MatchColor(ctx context.Context, color string) bool {
	v, ok := ColorExtractor(ctx)
	if !ok {
		return false
	}
	return v == color
}

func WithColor(color string, do endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return do(ColorInjector(ctx, color), request)
	}
}

func GrpcOutgoingInjector(ctx context.Context, grpcMD *metadata.MD) context.Context {
	color, ok := ColorExtractor(ctx)
	if !ok {
		return ctx
	}
	grpcMD.Set(kStainKey, color)
	return ctx
}

func GrpcIncomingInjector(ctx context.Context, md metadata.MD) context.Context {
	values := md.Get(kStainKey)
	if values == nil || len(values) == 0 {
		return ctx
	}
	return ColorInjector(ctx, values[0])
}

func HttpOutgoingInjector(ctx context.Context, request *http.Request) context.Context {
	color, ok := ColorExtractor(ctx)
	if !ok {
		return ctx
	}
	request.Header.Set(kStainKey, color)
	return ctx
}

func HttpIncomingInjector(ctx context.Context, request *http.Request) context.Context {
	values := request.Header.Values(kStainKey)
	if values == nil || len(values) == 0 {
		return ctx
	}
	return ColorInjector(ctx, values[0])
}
