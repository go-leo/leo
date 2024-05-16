// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package path

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type pathGRPCServer struct {
	boolPath grpc.Handler

	int32Path grpc.Handler

	int64Path grpc.Handler

	uint32Path grpc.Handler

	uint64Path grpc.Handler

	floatPath grpc.Handler

	doublePath grpc.Handler

	stringPath grpc.Handler

	enumPath grpc.Handler
}

func (s *pathGRPCServer) BoolPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.boolPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGRPCServer) Int32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.int32Path.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGRPCServer) Int64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.int64Path.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGRPCServer) Uint32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.uint32Path.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGRPCServer) Uint64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.uint64Path.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGRPCServer) FloatPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.floatPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGRPCServer) DoublePath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.doublePath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGRPCServer) StringPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.stringPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *pathGRPCServer) EnumPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.enumPath.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewPathGRPCServer(
	endpoints interface {
		BoolPath() endpoint.Endpoint
		Int32Path() endpoint.Endpoint
		Int64Path() endpoint.Endpoint
		Uint32Path() endpoint.Endpoint
		Uint64Path() endpoint.Endpoint
		FloatPath() endpoint.Endpoint
		DoublePath() endpoint.Endpoint
		StringPath() endpoint.Endpoint
		EnumPath() endpoint.Endpoint
	},
	mdw []endpoint.Middleware,
	opts ...grpc.ServerOption,
) interface {
	BoolPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Int32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Int64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Uint32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Uint64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	FloatPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	DoublePath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	StringPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	EnumPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
} {
	return &pathGRPCServer{
		boolPath: grpc.NewServer(
			endpointx.Chain(endpoints.BoolPath(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		int32Path: grpc.NewServer(
			endpointx.Chain(endpoints.Int32Path(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		int64Path: grpc.NewServer(
			endpointx.Chain(endpoints.Int64Path(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		uint32Path: grpc.NewServer(
			endpointx.Chain(endpoints.Uint32Path(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		uint64Path: grpc.NewServer(
			endpointx.Chain(endpoints.Uint64Path(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		floatPath: grpc.NewServer(
			endpointx.Chain(endpoints.FloatPath(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		doublePath: grpc.NewServer(
			endpointx.Chain(endpoints.DoublePath(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		stringPath: grpc.NewServer(
			endpointx.Chain(endpoints.StringPath(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		enumPath: grpc.NewServer(
			endpointx.Chain(endpoints.EnumPath(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
	}
}

type pathGRPCClient struct {
	boolPath   endpoint.Endpoint
	int32Path  endpoint.Endpoint
	int64Path  endpoint.Endpoint
	uint32Path endpoint.Endpoint
	uint64Path endpoint.Endpoint
	floatPath  endpoint.Endpoint
	doublePath endpoint.Endpoint
	stringPath endpoint.Endpoint
	enumPath   endpoint.Endpoint
}

func (c *pathGRPCClient) BoolPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.boolPath(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGRPCClient) Int32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.int32Path(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGRPCClient) Int64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.int64Path(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGRPCClient) Uint32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.uint32Path(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGRPCClient) Uint64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.uint64Path(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGRPCClient) FloatPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.floatPath(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGRPCClient) DoublePath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.doublePath(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGRPCClient) StringPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.stringPath(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *pathGRPCClient) EnumPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error) {
	rep, err := c.enumPath(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewPathGRPCClient(
	conn *grpc1.ClientConn,
	mdw []endpoint.Middleware,
	opts ...grpc.ClientOption,
) interface {
	BoolPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Int32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Int64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Uint32Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	Uint64Path(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	FloatPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	DoublePath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	StringPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
	EnumPath(ctx context.Context, request *PathRequest) (*emptypb.Empty, error)
} {
	return &pathGRPCClient{
		boolPath: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"BoolPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		int32Path: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"Int32Path",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		int64Path: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"Int64Path",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		uint32Path: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"Uint32Path",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		uint64Path: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"Uint64Path",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		floatPath: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"FloatPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		doublePath: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"DoublePath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		stringPath: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"StringPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
		enumPath: endpointx.Chain(
			grpc.NewClient(
				conn,
				"leo.example.path.v1.Path",
				"EnumPath",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
	}
}