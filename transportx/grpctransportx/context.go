package grpctransportx

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-leo/leo/v3/endpointx"
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
		return endpointx.NameInjector(ctx, name)
	}
}

func ClientTransportInjector(ctx context.Context, md *metadata.MD) context.Context {
	return transportx.NameInjector(ctx, GrpcClient)
}

func ServerEndpointInjector(name string) grpctransport.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		return endpointx.NameInjector(ctx, name)
	}
}

func ServerTransportInjector(ctx context.Context, md metadata.MD) context.Context {
	return transportx.NameInjector(ctx, GrpcServer)
}
