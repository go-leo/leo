package grpcx

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/transportx"
	"google.golang.org/grpc/metadata"
)

const (
	// GrpcServer is the name of the grpc server transport.
	GrpcServer = "grpc.server"

	// GrpcClient is the name of the grpc client transport.
	GrpcClient = "grpc.client"
)

func ClientEndpointInjector(name string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		return endpointx.InjectName(ctx, name)
	}
}

func ClientTransportInjector(ctx context.Context, md *metadata.MD) context.Context {
	return transportx.InjectName(ctx, GrpcClient)
}

func ServerEndpointInjector(name string) grpctransport.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		return endpointx.InjectName(ctx, name)
	}
}

func ServerTransportInjector(ctx context.Context, md metadata.MD) context.Context {
	return transportx.InjectName(ctx, GrpcServer)
}

func OutgoingMetadataInjector(ctx context.Context, grpcMD *metadata.MD) context.Context {
	md, ok := metadatax.FromOutgoingContext(ctx)
	if !ok {
		return ctx
	}
	for _, key := range md.Keys() {
		grpcMD.Set(key, md.Values(key)...)
	}
	return ctx
}

func IncomingMetadataInjector(ctx context.Context, md metadata.MD) context.Context {
	return metadatax.NewIncomingContext(ctx, metadatax.FromGrpcMetadata(md))
}
