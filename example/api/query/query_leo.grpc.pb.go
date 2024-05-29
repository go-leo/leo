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

type queryGRPCServer struct {
	query grpc.Handler
}

func (s *queryGRPCServer) Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.query.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewQueryGRPCServer(
	endpoints interface {
		Query() endpoint.Endpoint
	},
	opts []grpc.ServerOption,
	middlewares ...endpoint.Middleware,
) interface {
	Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error)
} {
	return &queryGRPCServer{
		query: grpc.NewServer(
			endpointx.Chain(endpoints.Query(), middlewares...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
	}
}

type queryGRPCClient struct {
	query endpoint.Endpoint
}

func (c *queryGRPCClient) Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error) {
	rep, err := c.query(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewQueryGRPCClient(
	conn *grpc1.ClientConn,
	opts []grpc.ClientOption,
	middlewares ...endpoint.Middleware,
) interface {
	Query(ctx context.Context, request *QueryRequest) (*emptypb.Empty, error)
} {
	return &queryGRPCClient{
		query: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.query.v1.Query",
				"Query",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			middlewares...),
	}
}
