package transportx

import (
	"context"
)

const (
	// GrpcServer is the name of the grpc server transport.
	GrpcServer = "grpc.server"

	// GrpcClient is the name of the grpc client transport.
	GrpcClient = "grpc.client"

	// HttpServer is the name of the http server transport.
	HttpServer = "http.server"

	// HttpClient is the name of the http client transport.
	HttpClient = "http.client"
)

type nameKey struct{}

// InjectName injects the name into the context.
func InjectName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, nameKey{}, name)
}

// ExtractName extracts the name from the context.
func ExtractName(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(nameKey{}).(string)
	return v, ok
}
