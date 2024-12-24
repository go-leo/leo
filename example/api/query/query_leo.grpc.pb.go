// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package query

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

// =========================== grpc server ===========================

type queryGrpcServerTransports struct {
	endpoints QueryServerEndpoints
}

func (t *queryGrpcServerTransports) Query() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.Query(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.query.v1.Query/Query")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

type queryGrpcServer struct {
	query grpc.Handler
}

func (s *queryGrpcServer) Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.query.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewQueryGrpcServer(svc QueryService, middlewares ...endpoint.Middleware) QueryService {
	endpoints := newQueryServerEndpoints(svc, middlewares...)
	transports := &queryGrpcServerTransports{endpoints: endpoints}
	return &queryGrpcServer{
		query: transports.Query(),
	}
}

// =========================== grpc client ===========================

type queryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *queryGrpcClientTransports) Query(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStain),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.query.v1.Query",
		"Query",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func newQueryGrpcClientTransports(
	dialOptions []grpc1.DialOption,
	clientOptions []grpc.ClientOption,
	middlewares []endpoint.Middleware,
) QueryClientTransports {
	return &queryGrpcClientTransports{
		dialOptions:   dialOptions,
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

func NewQueryGrpcClient(target string, opts ...grpcx.ClientOption) QueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := newQueryGrpcClientTransports(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newQueryClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newQueryClientService(endpoints, grpcx.GrpcClient)
}
