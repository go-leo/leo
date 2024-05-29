package transportx

import (
	"context"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/transport/http"
	grpcmetadata "google.golang.org/grpc/metadata"
	http1 "net/http"
)

type nameKey struct{}

func NewContext(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, nameKey{}, name)
}

func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(nameKey{}).(string)
	return v, ok
}

func InjectToHttp(name string) http.RequestFunc {
	return func(ctx context.Context, request *http1.Request) context.Context {
		return NewContext(ctx, name)
	}
}

func InjectToGrpcServer(name string) grpc.ServerRequestFunc {
	return func(ctx context.Context, md grpcmetadata.MD) context.Context {
		return NewContext(ctx, name)
	}
}

func InjectToGrpcClient(name string) grpc.ClientRequestFunc {
	return func(ctx context.Context, md *grpcmetadata.MD) context.Context {
		return NewContext(ctx, name)
	}
}
