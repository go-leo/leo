// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package path

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

type namedPathGrpcServerTransports struct {
	endpoints NamedPathServerEndpoints
}

func (t *namedPathGrpcServerTransports) NamedPathString() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.NamedPathString(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.NamedPath/NamedPathString")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

func (t *namedPathGrpcServerTransports) NamedPathOptString() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.NamedPathOptString(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.NamedPath/NamedPathOptString")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

func (t *namedPathGrpcServerTransports) NamedPathWrapString() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.NamedPathWrapString(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.NamedPath/NamedPathWrapString")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

func (t *namedPathGrpcServerTransports) EmbedNamedPathString() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.EmbedNamedPathString(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathString")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

func (t *namedPathGrpcServerTransports) EmbedNamedPathOptString() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.EmbedNamedPathOptString(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

func (t *namedPathGrpcServerTransports) EmbedNamedPathWrapString() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.EmbedNamedPathWrapString(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type namedPathGrpcServer struct {
	namedPathString          grpc.Handler
	namedPathOptString       grpc.Handler
	namedPathWrapString      grpc.Handler
	embedNamedPathString     grpc.Handler
	embedNamedPathOptString  grpc.Handler
	embedNamedPathWrapString grpc.Handler
}

func (s *namedPathGrpcServer) NamedPathString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.namedPathString.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *namedPathGrpcServer) NamedPathOptString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.namedPathOptString.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *namedPathGrpcServer) NamedPathWrapString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.namedPathWrapString.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *namedPathGrpcServer) EmbedNamedPathString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.embedNamedPathString.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *namedPathGrpcServer) EmbedNamedPathOptString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.embedNamedPathOptString.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *namedPathGrpcServer) EmbedNamedPathWrapString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.embedNamedPathWrapString.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewNamedPathGrpcServer(svc NamedPathService, middlewares ...endpoint.Middleware) NamedPathService {
	endpoints := newNamedPathServerEndpoints(svc, middlewares...)
	transports := &namedPathGrpcServerTransports{endpoints: endpoints}
	return &namedPathGrpcServer{
		namedPathString:          transports.NamedPathString(),
		namedPathOptString:       transports.NamedPathOptString(),
		namedPathWrapString:      transports.NamedPathWrapString(),
		embedNamedPathString:     transports.EmbedNamedPathString(),
		embedNamedPathOptString:  transports.EmbedNamedPathOptString(),
		embedNamedPathWrapString: transports.EmbedNamedPathWrapString(),
	}
}

// =========================== grpc client ===========================

type namedPathGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *namedPathGrpcClientTransports) NamedPathString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.path.v1.NamedPath",
		"NamedPathString",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *namedPathGrpcClientTransports) NamedPathOptString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.path.v1.NamedPath",
		"NamedPathOptString",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *namedPathGrpcClientTransports) NamedPathWrapString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.path.v1.NamedPath",
		"NamedPathWrapString",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *namedPathGrpcClientTransports) EmbedNamedPathString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.path.v1.NamedPath",
		"EmbedNamedPathString",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *namedPathGrpcClientTransports) EmbedNamedPathOptString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.path.v1.NamedPath",
		"EmbedNamedPathOptString",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *namedPathGrpcClientTransports) EmbedNamedPathWrapString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc1.NewClient(instance, t.dialOptions...)
	if err != nil {
		return nil, nil, err
	}
	opts := []grpc.ClientOption{
		grpc.ClientBefore(grpcx.OutgoingMetadataInjector),
		grpc.ClientBefore(grpcx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := grpc.NewClient(
		conn,
		"leo.example.path.v1.NamedPath",
		"EmbedNamedPathWrapString",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func newNamedPathGrpcClientTransports(
	dialOptions []grpc1.DialOption,
	clientOptions []grpc.ClientOption,
	middlewares []endpoint.Middleware,
) NamedPathClientTransports {
	return &namedPathGrpcClientTransports{
		dialOptions:   dialOptions,
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

func NewNamedPathGrpcClient(target string, opts ...grpcx.ClientOption) NamedPathService {
	options := grpcx.NewClientOptions(opts...)
	transports := newNamedPathGrpcClientTransports(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newNamedPathClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newNamedPathClientService(endpoints, grpcx.GrpcClient)
}
