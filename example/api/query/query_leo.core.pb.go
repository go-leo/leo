// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package query

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// =========================== endpoints ===========================

type QueryService interface {
	Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error)
}

type QueryEndpoints interface {
	Query() endpoint.Endpoint
}

type queryEndpoints struct {
	svc         QueryService
	middlewares []endpoint.Middleware
}

func (e *queryEndpoints) Query() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.Query(ctx, request.(*QueryRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewQueryEndpoints(svc QueryService, middlewares ...endpoint.Middleware) QueryEndpoints {
	return &queryEndpoints{svc: svc, middlewares: middlewares}
}

// =========================== cqrs ===========================

// =========================== grpc transports ===========================

type QueryGrpcServerTransports interface {
	Query() *grpc.Server
}

type QueryGrpcClientTransports interface {
	Query() *grpc.Client
}

type queryGrpcServerTransports struct {
	query *grpc.Server
}

func (t *queryGrpcServerTransports) Query() *grpc.Server {
	return t.query
}

func NewQueryGrpcServerTransports(endpoints QueryEndpoints, serverOptions ...grpc.ServerOption) QueryGrpcServerTransports {
	return &queryGrpcServerTransports{
		query: grpc.NewServer(
			endpoints.Query(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			serverOptions...,
		),
	}
}

type queryGrpcClientTransports struct {
	query *grpc.Client
}

func (t *queryGrpcClientTransports) Query() *grpc.Client {
	return t.query
}

func NewQueryGrpcClientTransports(conn *grpc1.ClientConn, clientOptions ...grpc.ClientOption) QueryGrpcClientTransports {
	return &queryGrpcClientTransports{
		query: grpc.NewClient(
			conn,
			"leo.example.query.v1.Query",
			"Query",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			clientOptions...,
		),
	}
}

type queryGrpcServer struct {
	query *grpc.Server
}

func (s *queryGrpcServer) Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.query.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewQueryGrpcServer(transports QueryGrpcServerTransports) QueryService {
	return &queryGrpcServer{
		query: transports.Query(),
	}
}

type queryGrpcClient struct {
	query endpoint.Endpoint
}

func (c *queryGrpcClient) Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error) {
	rep, err := c.query(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewQueryGrpcClient(transports QueryGrpcClientTransports, middlewares ...endpoint.Middleware) QueryService {
	return &queryGrpcClient{
		query: endpointx.Chain(transports.Query().Endpoint(), middlewares...),
	}
}
