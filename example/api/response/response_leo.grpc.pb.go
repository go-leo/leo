// Code generated by protoc-gen-leo-grpc. DO NOT EDIT.

package response

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

type responseGrpcServerTransports struct {
	endpoints ResponseServerEndpoints
}

func (t *responseGrpcServerTransports) OmittedResponse() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.OmittedResponse(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/OmittedResponse")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

func (t *responseGrpcServerTransports) StarResponse() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.StarResponse(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/StarResponse")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

func (t *responseGrpcServerTransports) NamedResponse() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.NamedResponse(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/NamedResponse")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

func (t *responseGrpcServerTransports) HttpBodyResponse() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.HttpBodyResponse(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/HttpBodyResponse")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

func (t *responseGrpcServerTransports) HttpBodyNamedResponse() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.HttpBodyNamedResponse(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/HttpBodyNamedResponse")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStain),
	)
}

type responseGrpcServer struct {
	omittedResponse       grpc.Handler
	starResponse          grpc.Handler
	namedResponse         grpc.Handler
	httpBodyResponse      grpc.Handler
	httpBodyNamedResponse grpc.Handler
}

func (s *responseGrpcServer) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.omittedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.starResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.namedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	ctx, rep, err := s.httpBodyResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*httpbody.HttpBody), nil
}

func (s *responseGrpcServer) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	ctx, rep, err := s.httpBodyNamedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*HttpBody), nil
}

func NewResponseGrpcServer(svc ResponseService, middlewares ...endpoint.Middleware) ResponseService {
	endpoints := newResponseServerEndpoints(svc, middlewares...)
	transports := &responseGrpcServerTransports{endpoints: endpoints}
	return &responseGrpcServer{
		omittedResponse:       transports.OmittedResponse(),
		starResponse:          transports.StarResponse(),
		namedResponse:         transports.NamedResponse(),
		httpBodyResponse:      transports.HttpBodyResponse(),
		httpBodyNamedResponse: transports.HttpBodyNamedResponse(),
	}
}

// =========================== grpc client ===========================

type responseGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *responseGrpcClientTransports) OmittedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.response.v1.Response",
		"OmittedResponse",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		UserResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *responseGrpcClientTransports) StarResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.response.v1.Response",
		"StarResponse",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		UserResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *responseGrpcClientTransports) NamedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.response.v1.Response",
		"NamedResponse",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		UserResponse{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *responseGrpcClientTransports) HttpBodyResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.response.v1.Response",
		"HttpBodyResponse",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		httpbody.HttpBody{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func (t *responseGrpcClientTransports) HttpBodyNamedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.response.v1.Response",
		"HttpBodyNamedResponse",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		HttpBody{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func newResponseGrpcClientTransports(
	dialOptions []grpc1.DialOption,
	clientOptions []grpc.ClientOption,
	middlewares []endpoint.Middleware,
) ResponseClientTransports {
	return &responseGrpcClientTransports{
		dialOptions:   dialOptions,
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

func NewResponseGrpcClient(target string, opts ...grpcx.ClientOption) ResponseService {
	options := grpcx.NewClientOptions(opts...)
	transports := newResponseGrpcClientTransports(options.DialOptions(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newResponseClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newResponseClientService(endpoints, grpcx.GrpcClient)
}
