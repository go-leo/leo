// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package endpointsapis

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	endpointx "github.com/go-leo/kitx/endpointx"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type gRPCWorkspacesServer struct {
	listWorkspaces grpc.Handler

	getWorkspace grpc.Handler

	createWorkspace grpc.Handler

	updateWorkspace grpc.Handler

	deleteWorkspace grpc.Handler
}

func (s *gRPCWorkspacesServer) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	ctx, rep, err := s.listWorkspaces.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*ListWorkspacesResponse), nil
}

func (s *gRPCWorkspacesServer) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.getWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *gRPCWorkspacesServer) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.createWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *gRPCWorkspacesServer) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.updateWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *gRPCWorkspacesServer) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.deleteWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesGRPCServer(
	endpoints interface {
		ListWorkspaces() endpoint.Endpoint
		GetWorkspace() endpoint.Endpoint
		CreateWorkspace() endpoint.Endpoint
		UpdateWorkspace() endpoint.Endpoint
		DeleteWorkspace() endpoint.Endpoint
	},
	mdw []endpoint.Middleware,
	opts ...grpc.ServerOption,
) interface {
	ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error)
	CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error)
	UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error)
	DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error)
} {
	return &gRPCWorkspacesServer{
		listWorkspaces: grpc.NewServer(
			endpointx.Chain(endpoints.ListWorkspaces(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		getWorkspace: grpc.NewServer(
			endpointx.Chain(endpoints.GetWorkspace(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		createWorkspace: grpc.NewServer(
			endpointx.Chain(endpoints.CreateWorkspace(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		updateWorkspace: grpc.NewServer(
			endpointx.Chain(endpoints.UpdateWorkspace(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
		deleteWorkspace: grpc.NewServer(
			endpointx.Chain(endpoints.DeleteWorkspace(), mdw...),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			opts...,
		),
	}
}

type gRPCWorkspacesClient struct {
	listWorkspaces  endpoint.Endpoint
	getWorkspace    endpoint.Endpoint
	createWorkspace endpoint.Endpoint
	updateWorkspace endpoint.Endpoint
	deleteWorkspace endpoint.Endpoint
}

func (c *gRPCWorkspacesClient) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	rep, err := c.listWorkspaces(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListWorkspacesResponse), nil
}

func (c *gRPCWorkspacesClient) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	rep, err := c.getWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *gRPCWorkspacesClient) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	rep, err := c.createWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *gRPCWorkspacesClient) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	rep, err := c.updateWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *gRPCWorkspacesClient) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesGRPCClient(
	conn *grpc1.ClientConn,
	mdw []endpoint.Middleware,
	opts ...grpc.ClientOption,
) interface {
	ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error)
	CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error)
	UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error)
	DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error)
} {
	return &gRPCWorkspacesClient{
		listWorkspaces: endpointx.Chain(
			grpc.NewClient(
				conn,
				"google.example.endpointsapis.v1.Workspaces",
				"ListWorkspaces",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				ListWorkspacesResponse{},
				opts...,
			).Endpoint(),
			mdw...),
		getWorkspace: endpointx.Chain(
			grpc.NewClient(
				conn,
				"google.example.endpointsapis.v1.Workspaces",
				"GetWorkspace",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				Workspace{},
				opts...,
			).Endpoint(),
			mdw...),
		createWorkspace: endpointx.Chain(
			grpc.NewClient(
				conn,
				"google.example.endpointsapis.v1.Workspaces",
				"CreateWorkspace",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				Workspace{},
				opts...,
			).Endpoint(),
			mdw...),
		updateWorkspace: endpointx.Chain(
			grpc.NewClient(
				conn,
				"google.example.endpointsapis.v1.Workspaces",
				"UpdateWorkspace",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				Workspace{},
				opts...,
			).Endpoint(),
			mdw...),
		deleteWorkspace: endpointx.Chain(
			grpc.NewClient(
				conn,
				"google.example.endpointsapis.v1.Workspaces",
				"DeleteWorkspace",
				func(_ context.Context, v any) (any, error) { return v, nil },
				func(_ context.Context, v any) (any, error) { return v, nil },
				emptypb.Empty{},
				opts...,
			).Endpoint(),
			mdw...),
	}
}