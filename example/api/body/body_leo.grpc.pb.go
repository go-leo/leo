// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package body

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

// =========================== grpc server ===========================

type bodyGrpcServerTransports struct {
	endpoints BodyServerEndpoints
}

func (t *bodyGrpcServerTransports) StarBody() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.StarBody(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/StarBody")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

func (t *bodyGrpcServerTransports) NamedBody() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.NamedBody(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/NamedBody")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

func (t *bodyGrpcServerTransports) NonBody() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.NonBody(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/NonBody")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

func (t *bodyGrpcServerTransports) HttpBodyStarBody() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.HttpBodyStarBody(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/HttpBodyStarBody")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

func (t *bodyGrpcServerTransports) HttpBodyNamedBody() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.HttpBodyNamedBody(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/HttpBodyNamedBody")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

type bodyGrpcServer struct {
	starBody          grpc.Handler
	namedBody         grpc.Handler
	nonBody           grpc.Handler
	httpBodyStarBody  grpc.Handler
	httpBodyNamedBody grpc.Handler
}

func (s *bodyGrpcServer) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	ctx, rep, err := s.starBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.namedBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	ctx, rep, err := s.nonBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpBodyStarBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpBodyNamedBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewBodyGrpcServer(svc BodyService, middlewares ...endpoint.Middleware) BodyService {
	endpoints := newBodyServerEndpoints(svc, middlewares...)
	transports := &bodyGrpcServerTransports{endpoints: endpoints}
	return &bodyGrpcServer{
		starBody:          transports.StarBody(),
		namedBody:         transports.NamedBody(),
		nonBody:           transports.NonBody(),
		httpBodyStarBody:  transports.HttpBodyStarBody(),
		httpBodyNamedBody: transports.HttpBodyNamedBody(),
	}
}

// =========================== grpc client ===========================

type bodyGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *bodyGrpcClientTransports) StarBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.body.v1.Body",
		"StarBody",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *bodyGrpcClientTransports) NamedBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.body.v1.Body",
		"NamedBody",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *bodyGrpcClientTransports) NonBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.body.v1.Body",
		"NonBody",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *bodyGrpcClientTransports) HttpBodyStarBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.body.v1.Body",
		"HttpBodyStarBody",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *bodyGrpcClientTransports) HttpBodyNamedBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.body.v1.Body",
		"HttpBodyNamedBody",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func newBodyGrpcClientTransports(
	dialOptions []grpc1.DialOption,
	clientOptions []grpc.ClientOption,
	middlewares []endpoint.Middleware,
) BodyClientTransports {
	return &bodyGrpcClientTransports{
		dialOptions:   dialOptions,
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

func NewBodyGrpcClient(target string, opts ...grpcx.ClientOption) BodyService {
	options := grpcx.NewClientOptions(opts...)
	transports := newBodyGrpcClientTransports(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newBodyClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newBodyClientService(endpoints, grpcx.GrpcClient)
}
