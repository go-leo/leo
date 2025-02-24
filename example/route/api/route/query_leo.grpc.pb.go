// Code generated by protoc-gen-leo. DO NOT EDIT.

package route

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

func NewBoolQueryGrpcServer(svc BoolQueryService, middlewares ...endpoint.Middleware) BoolQueryServer {
	endpoints := &boolQueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &boolQueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &boolQueryGrpcServer{
		boolQuery: transports.BoolQuery(),
	}
}

func NewBoolQueryGrpcClient(target string, opts ...grpcx.ClientOption) BoolQueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &boolQueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &boolQueryFactories{
		transports: transports,
	}
	endpointer := &boolQueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &boolQueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &boolQueryClientEndpoints{
		balancers: balancers,
	}
	return &boolQueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type boolQueryGrpcServerTransports struct {
	endpoints BoolQueryServerEndpoints
}

func (t *boolQueryGrpcServerTransports) BoolQuery() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.BoolQuery(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.BoolQuery/BoolQuery")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type boolQueryGrpcServer struct {
	boolQuery grpc.Handler
}

func (s *boolQueryGrpcServer) BoolQuery(ctx context.Context, request *BoolQueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.boolQuery.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *boolQueryGrpcServer) mustEmbedUnimplementedBoolQueryServer() {}

type boolQueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *boolQueryGrpcClientTransports) BoolQuery(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.BoolQuery",
		"BoolQuery",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func NewInt32QueryGrpcServer(svc Int32QueryService, middlewares ...endpoint.Middleware) Int32QueryServer {
	endpoints := &int32QueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &int32QueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &int32QueryGrpcServer{
		int32Query: transports.Int32Query(),
	}
}

func NewInt32QueryGrpcClient(target string, opts ...grpcx.ClientOption) Int32QueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &int32QueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &int32QueryFactories{
		transports: transports,
	}
	endpointer := &int32QueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &int32QueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &int32QueryClientEndpoints{
		balancers: balancers,
	}
	return &int32QueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type int32QueryGrpcServerTransports struct {
	endpoints Int32QueryServerEndpoints
}

func (t *int32QueryGrpcServerTransports) Int32Query() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.Int32Query(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.Int32Query/Int32Query")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type int32QueryGrpcServer struct {
	int32Query grpc.Handler
}

func (s *int32QueryGrpcServer) Int32Query(ctx context.Context, request *Int32QueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.int32Query.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *int32QueryGrpcServer) mustEmbedUnimplementedInt32QueryServer() {}

type int32QueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *int32QueryGrpcClientTransports) Int32Query(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.Int32Query",
		"Int32Query",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func NewInt64QueryGrpcServer(svc Int64QueryService, middlewares ...endpoint.Middleware) Int64QueryServer {
	endpoints := &int64QueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &int64QueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &int64QueryGrpcServer{
		int64Query: transports.Int64Query(),
	}
}

func NewInt64QueryGrpcClient(target string, opts ...grpcx.ClientOption) Int64QueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &int64QueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &int64QueryFactories{
		transports: transports,
	}
	endpointer := &int64QueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &int64QueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &int64QueryClientEndpoints{
		balancers: balancers,
	}
	return &int64QueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type int64QueryGrpcServerTransports struct {
	endpoints Int64QueryServerEndpoints
}

func (t *int64QueryGrpcServerTransports) Int64Query() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.Int64Query(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.Int64Query/Int64Query")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type int64QueryGrpcServer struct {
	int64Query grpc.Handler
}

func (s *int64QueryGrpcServer) Int64Query(ctx context.Context, request *Int64QueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.int64Query.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *int64QueryGrpcServer) mustEmbedUnimplementedInt64QueryServer() {}

type int64QueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *int64QueryGrpcClientTransports) Int64Query(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.Int64Query",
		"Int64Query",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func NewUint32QueryGrpcServer(svc Uint32QueryService, middlewares ...endpoint.Middleware) Uint32QueryServer {
	endpoints := &uint32QueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &uint32QueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &uint32QueryGrpcServer{
		uint32Query: transports.Uint32Query(),
	}
}

func NewUint32QueryGrpcClient(target string, opts ...grpcx.ClientOption) Uint32QueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &uint32QueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &uint32QueryFactories{
		transports: transports,
	}
	endpointer := &uint32QueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &uint32QueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &uint32QueryClientEndpoints{
		balancers: balancers,
	}
	return &uint32QueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type uint32QueryGrpcServerTransports struct {
	endpoints Uint32QueryServerEndpoints
}

func (t *uint32QueryGrpcServerTransports) Uint32Query() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.Uint32Query(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.Uint32Query/Uint32Query")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type uint32QueryGrpcServer struct {
	uint32Query grpc.Handler
}

func (s *uint32QueryGrpcServer) Uint32Query(ctx context.Context, request *Uint32QueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.uint32Query.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *uint32QueryGrpcServer) mustEmbedUnimplementedUint32QueryServer() {}

type uint32QueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *uint32QueryGrpcClientTransports) Uint32Query(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.Uint32Query",
		"Uint32Query",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func NewUint64QueryGrpcServer(svc Uint64QueryService, middlewares ...endpoint.Middleware) Uint64QueryServer {
	endpoints := &uint64QueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &uint64QueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &uint64QueryGrpcServer{
		uint64Query: transports.Uint64Query(),
	}
}

func NewUint64QueryGrpcClient(target string, opts ...grpcx.ClientOption) Uint64QueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &uint64QueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &uint64QueryFactories{
		transports: transports,
	}
	endpointer := &uint64QueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &uint64QueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &uint64QueryClientEndpoints{
		balancers: balancers,
	}
	return &uint64QueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type uint64QueryGrpcServerTransports struct {
	endpoints Uint64QueryServerEndpoints
}

func (t *uint64QueryGrpcServerTransports) Uint64Query() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.Uint64Query(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.Uint64Query/Uint64Query")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type uint64QueryGrpcServer struct {
	uint64Query grpc.Handler
}

func (s *uint64QueryGrpcServer) Uint64Query(ctx context.Context, request *Uint64QueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.uint64Query.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *uint64QueryGrpcServer) mustEmbedUnimplementedUint64QueryServer() {}

type uint64QueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *uint64QueryGrpcClientTransports) Uint64Query(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.Uint64Query",
		"Uint64Query",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func NewFloatQueryGrpcServer(svc FloatQueryService, middlewares ...endpoint.Middleware) FloatQueryServer {
	endpoints := &floatQueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &floatQueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &floatQueryGrpcServer{
		floatQuery: transports.FloatQuery(),
	}
}

func NewFloatQueryGrpcClient(target string, opts ...grpcx.ClientOption) FloatQueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &floatQueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &floatQueryFactories{
		transports: transports,
	}
	endpointer := &floatQueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &floatQueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &floatQueryClientEndpoints{
		balancers: balancers,
	}
	return &floatQueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type floatQueryGrpcServerTransports struct {
	endpoints FloatQueryServerEndpoints
}

func (t *floatQueryGrpcServerTransports) FloatQuery() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.FloatQuery(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.FloatQuery/FloatQuery")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type floatQueryGrpcServer struct {
	floatQuery grpc.Handler
}

func (s *floatQueryGrpcServer) FloatQuery(ctx context.Context, request *FloatQueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.floatQuery.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *floatQueryGrpcServer) mustEmbedUnimplementedFloatQueryServer() {}

type floatQueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *floatQueryGrpcClientTransports) FloatQuery(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.FloatQuery",
		"FloatQuery",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func NewDoubleQueryGrpcServer(svc DoubleQueryService, middlewares ...endpoint.Middleware) DoubleQueryServer {
	endpoints := &doubleQueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &doubleQueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &doubleQueryGrpcServer{
		doubleQuery: transports.DoubleQuery(),
	}
}

func NewDoubleQueryGrpcClient(target string, opts ...grpcx.ClientOption) DoubleQueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &doubleQueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &doubleQueryFactories{
		transports: transports,
	}
	endpointer := &doubleQueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &doubleQueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &doubleQueryClientEndpoints{
		balancers: balancers,
	}
	return &doubleQueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type doubleQueryGrpcServerTransports struct {
	endpoints DoubleQueryServerEndpoints
}

func (t *doubleQueryGrpcServerTransports) DoubleQuery() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.DoubleQuery(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.DoubleQuery/DoubleQuery")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type doubleQueryGrpcServer struct {
	doubleQuery grpc.Handler
}

func (s *doubleQueryGrpcServer) DoubleQuery(ctx context.Context, request *DoubleQueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.doubleQuery.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *doubleQueryGrpcServer) mustEmbedUnimplementedDoubleQueryServer() {}

type doubleQueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *doubleQueryGrpcClientTransports) DoubleQuery(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.DoubleQuery",
		"DoubleQuery",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func NewStringQueryGrpcServer(svc StringQueryService, middlewares ...endpoint.Middleware) StringQueryServer {
	endpoints := &stringQueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &stringQueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &stringQueryGrpcServer{
		stringQuery: transports.StringQuery(),
	}
}

func NewStringQueryGrpcClient(target string, opts ...grpcx.ClientOption) StringQueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &stringQueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &stringQueryFactories{
		transports: transports,
	}
	endpointer := &stringQueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &stringQueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &stringQueryClientEndpoints{
		balancers: balancers,
	}
	return &stringQueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type stringQueryGrpcServerTransports struct {
	endpoints StringQueryServerEndpoints
}

func (t *stringQueryGrpcServerTransports) StringQuery() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.StringQuery(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.StringQuery/StringQuery")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type stringQueryGrpcServer struct {
	stringQuery grpc.Handler
}

func (s *stringQueryGrpcServer) StringQuery(ctx context.Context, request *StringQueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.stringQuery.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *stringQueryGrpcServer) mustEmbedUnimplementedStringQueryServer() {}

type stringQueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *stringQueryGrpcClientTransports) StringQuery(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.StringQuery",
		"StringQuery",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}

func NewEnumQueryGrpcServer(svc EnumQueryService, middlewares ...endpoint.Middleware) EnumQueryServer {
	endpoints := &enumQueryServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &enumQueryGrpcServerTransports{
		endpoints: endpoints,
	}
	return &enumQueryGrpcServer{
		enumQuery: transports.EnumQuery(),
	}
}

func NewEnumQueryGrpcClient(target string, opts ...grpcx.ClientOption) EnumQueryService {
	options := grpcx.NewClientOptions(opts...)
	transports := &enumQueryGrpcClientTransports{
		dialOptions:   options.DialOptions(),
		clientOptions: options.ClientTransportOptions(),
		middlewares:   options.Middlewares(),
	}
	factories := &enumQueryFactories{
		transports: transports,
	}
	endpointer := &enumQueryEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &enumQueryBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &enumQueryClientEndpoints{
		balancers: balancers,
	}
	return &enumQueryClientService{
		endpoints:     endpoints,
		transportName: grpcx.GrpcClient,
	}
}

type enumQueryGrpcServerTransports struct {
	endpoints EnumQueryServerEndpoints
}

func (t *enumQueryGrpcServerTransports) EnumQuery() grpc.Handler {
	return grpc.NewServer(
		t.endpoints.EnumQuery(context.TODO()),
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.route.query.EnumQuery/EnumQuery")),
		grpc.ServerBefore(grpcx.ServerTransportInjector),
		grpc.ServerBefore(grpcx.IncomingMetadataInjector),
		grpc.ServerBefore(grpcx.IncomingStainInjector),
	)
}

type enumQueryGrpcServer struct {
	enumQuery grpc.Handler
}

func (s *enumQueryGrpcServer) EnumQuery(ctx context.Context, request *EnumQueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.enumQuery.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *enumQueryGrpcServer) mustEmbedUnimplementedEnumQueryServer() {}

type enumQueryGrpcClientTransports struct {
	dialOptions   []grpc1.DialOption
	clientOptions []grpc.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *enumQueryGrpcClientTransports) EnumQuery(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
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
		"leo.example.route.query.EnumQuery",
		"EnumQuery",
		func(_ context.Context, v any) (any, error) { return v, nil },
		func(_ context.Context, v any) (any, error) { return v, nil },
		emptypb.Empty{},
		opts...)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), conn, nil
}
