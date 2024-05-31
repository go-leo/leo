// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package endpointsapis

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	urlx "github.com/go-leo/gox/netx/urlx"
	strconvx "github.com/go-leo/gox/strconvx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	transportx "github.com/go-leo/leo/v3/transportx"
	mux "github.com/gorilla/mux"
	grpc1 "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http1 "net/http"
	url "net/url"
	strings "strings"
)

// =========================== endpoints ===========================

type WorkspacesService interface {
	ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error)
	GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error)
	CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error)
	UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error)
	DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error)
}

type WorkspacesEndpoints interface {
	ListWorkspaces() endpoint.Endpoint
	GetWorkspace() endpoint.Endpoint
	CreateWorkspace() endpoint.Endpoint
	UpdateWorkspace() endpoint.Endpoint
	DeleteWorkspace() endpoint.Endpoint
}

type workspacesEndpoints struct {
	svc         WorkspacesService
	middlewares []endpoint.Middleware
}

func (e *workspacesEndpoints) ListWorkspaces() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.ListWorkspaces(ctx, request.(*ListWorkspacesRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *workspacesEndpoints) GetWorkspace() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.GetWorkspace(ctx, request.(*GetWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *workspacesEndpoints) CreateWorkspace() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.CreateWorkspace(ctx, request.(*CreateWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *workspacesEndpoints) UpdateWorkspace() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.UpdateWorkspace(ctx, request.(*UpdateWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *workspacesEndpoints) DeleteWorkspace() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.DeleteWorkspace(ctx, request.(*DeleteWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewWorkspacesEndpoints(svc WorkspacesService, middlewares ...endpoint.Middleware) WorkspacesEndpoints {
	return &workspacesEndpoints{svc: svc, middlewares: middlewares}
}

// =========================== cqrs ===========================

// =========================== grpc transports ===========================

type WorkspacesGrpcServerTransports interface {
	ListWorkspaces() *grpc.Server
	GetWorkspace() *grpc.Server
	CreateWorkspace() *grpc.Server
	UpdateWorkspace() *grpc.Server
	DeleteWorkspace() *grpc.Server
}

type WorkspacesGrpcClientTransports interface {
	ListWorkspaces() *grpc.Client
	GetWorkspace() *grpc.Client
	CreateWorkspace() *grpc.Client
	UpdateWorkspace() *grpc.Client
	DeleteWorkspace() *grpc.Client
}

type workspacesGrpcServerTransports struct {
	listWorkspaces  *grpc.Server
	getWorkspace    *grpc.Server
	createWorkspace *grpc.Server
	updateWorkspace *grpc.Server
	deleteWorkspace *grpc.Server
}

func (t *workspacesGrpcServerTransports) ListWorkspaces() *grpc.Server {
	return t.listWorkspaces
}

func (t *workspacesGrpcServerTransports) GetWorkspace() *grpc.Server {
	return t.getWorkspace
}

func (t *workspacesGrpcServerTransports) CreateWorkspace() *grpc.Server {
	return t.createWorkspace
}

func (t *workspacesGrpcServerTransports) UpdateWorkspace() *grpc.Server {
	return t.updateWorkspace
}

func (t *workspacesGrpcServerTransports) DeleteWorkspace() *grpc.Server {
	return t.deleteWorkspace
}

func NewWorkspacesGrpcServerTransports(endpoints WorkspacesEndpoints, serverOptions ...grpc.ServerOption) WorkspacesGrpcServerTransports {
	return &workspacesGrpcServerTransports{
		listWorkspaces: grpc.NewServer(
			endpoints.ListWorkspaces(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			append([]grpc.ServerOption{
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")
				}),
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcServer)
				}),
			}, serverOptions...)...,
		),
		getWorkspace: grpc.NewServer(
			endpoints.GetWorkspace(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			append([]grpc.ServerOption{
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/GetWorkspace")
				}),
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcServer)
				}),
			}, serverOptions...)...,
		),
		createWorkspace: grpc.NewServer(
			endpoints.CreateWorkspace(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			append([]grpc.ServerOption{
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")
				}),
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcServer)
				}),
			}, serverOptions...)...,
		),
		updateWorkspace: grpc.NewServer(
			endpoints.UpdateWorkspace(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			append([]grpc.ServerOption{
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")
				}),
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcServer)
				}),
			}, serverOptions...)...,
		),
		deleteWorkspace: grpc.NewServer(
			endpoints.DeleteWorkspace(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			append([]grpc.ServerOption{
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")
				}),
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcServer)
				}),
			}, serverOptions...)...,
		),
	}
}

type workspacesGrpcClientTransports struct {
	listWorkspaces  *grpc.Client
	getWorkspace    *grpc.Client
	createWorkspace *grpc.Client
	updateWorkspace *grpc.Client
	deleteWorkspace *grpc.Client
}

func (t *workspacesGrpcClientTransports) ListWorkspaces() *grpc.Client {
	return t.listWorkspaces
}

func (t *workspacesGrpcClientTransports) GetWorkspace() *grpc.Client {
	return t.getWorkspace
}

func (t *workspacesGrpcClientTransports) CreateWorkspace() *grpc.Client {
	return t.createWorkspace
}

func (t *workspacesGrpcClientTransports) UpdateWorkspace() *grpc.Client {
	return t.updateWorkspace
}

func (t *workspacesGrpcClientTransports) DeleteWorkspace() *grpc.Client {
	return t.deleteWorkspace
}

func NewWorkspacesGrpcClientTransports(conn *grpc1.ClientConn, clientOptions ...grpc.ClientOption) WorkspacesGrpcClientTransports {
	return &workspacesGrpcClientTransports{
		listWorkspaces: grpc.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"ListWorkspaces",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			ListWorkspacesResponse{},
			append([]grpc.ClientOption{
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")
				}),
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcClient)
				}),
			}, clientOptions...)...,
		),
		getWorkspace: grpc.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"GetWorkspace",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			Workspace{},
			append([]grpc.ClientOption{
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/GetWorkspace")
				}),
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcClient)
				}),
			}, clientOptions...)...,
		),
		createWorkspace: grpc.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"CreateWorkspace",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			Workspace{},
			append([]grpc.ClientOption{
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")
				}),
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcClient)
				}),
			}, clientOptions...)...,
		),
		updateWorkspace: grpc.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"UpdateWorkspace",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			Workspace{},
			append([]grpc.ClientOption{
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")
				}),
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcClient)
				}),
			}, clientOptions...)...,
		),
		deleteWorkspace: grpc.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"DeleteWorkspace",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			append([]grpc.ClientOption{
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")
				}),
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcClient)
				}),
			}, clientOptions...)...,
		),
	}
}

type workspacesGrpcServer struct {
	listWorkspaces  *grpc.Server
	getWorkspace    *grpc.Server
	createWorkspace *grpc.Server
	updateWorkspace *grpc.Server
	deleteWorkspace *grpc.Server
}

func (s *workspacesGrpcServer) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	ctx, rep, err := s.listWorkspaces.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*ListWorkspacesResponse), nil
}

func (s *workspacesGrpcServer) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.getWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *workspacesGrpcServer) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.createWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *workspacesGrpcServer) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	ctx, rep, err := s.updateWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*Workspace), nil
}

func (s *workspacesGrpcServer) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.deleteWorkspace.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesGrpcServer(transports WorkspacesGrpcServerTransports) WorkspacesService {
	return &workspacesGrpcServer{
		listWorkspaces:  transports.ListWorkspaces(),
		getWorkspace:    transports.GetWorkspace(),
		createWorkspace: transports.CreateWorkspace(),
		updateWorkspace: transports.UpdateWorkspace(),
		deleteWorkspace: transports.DeleteWorkspace(),
	}
}

type workspacesGrpcClient struct {
	listWorkspaces  endpoint.Endpoint
	getWorkspace    endpoint.Endpoint
	createWorkspace endpoint.Endpoint
	updateWorkspace endpoint.Endpoint
	deleteWorkspace endpoint.Endpoint
}

func (c *workspacesGrpcClient) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	rep, err := c.listWorkspaces(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListWorkspacesResponse), nil
}

func (c *workspacesGrpcClient) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	rep, err := c.getWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	rep, err := c.createWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	rep, err := c.updateWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesGrpcClient(transports WorkspacesGrpcClientTransports, middlewares ...endpoint.Middleware) WorkspacesService {
	return &workspacesGrpcClient{
		listWorkspaces:  endpointx.Chain(transports.ListWorkspaces().Endpoint(), middlewares...),
		getWorkspace:    endpointx.Chain(transports.GetWorkspace().Endpoint(), middlewares...),
		createWorkspace: endpointx.Chain(transports.CreateWorkspace().Endpoint(), middlewares...),
		updateWorkspace: endpointx.Chain(transports.UpdateWorkspace().Endpoint(), middlewares...),
		deleteWorkspace: endpointx.Chain(transports.DeleteWorkspace().Endpoint(), middlewares...),
	}
}

// =========================== http transports ===========================

type WorkspacesHttpServerTransports interface {
	ListWorkspaces() *http.Server
	GetWorkspace() *http.Server
	CreateWorkspace() *http.Server
	UpdateWorkspace() *http.Server
	DeleteWorkspace() *http.Server
}

type WorkspacesHttpClientTransports interface {
	ListWorkspaces() *http.Client
	GetWorkspace() *http.Client
	CreateWorkspace() *http.Client
	UpdateWorkspace() *http.Client
	DeleteWorkspace() *http.Client
}

type workspacesHttpServerTransports struct {
	listWorkspaces  *http.Server
	getWorkspace    *http.Server
	createWorkspace *http.Server
	updateWorkspace *http.Server
	deleteWorkspace *http.Server
}

func (t *workspacesHttpServerTransports) ListWorkspaces() *http.Server {
	return t.listWorkspaces
}

func (t *workspacesHttpServerTransports) GetWorkspace() *http.Server {
	return t.getWorkspace
}

func (t *workspacesHttpServerTransports) CreateWorkspace() *http.Server {
	return t.createWorkspace
}

func (t *workspacesHttpServerTransports) UpdateWorkspace() *http.Server {
	return t.updateWorkspace
}

func (t *workspacesHttpServerTransports) DeleteWorkspace() *http.Server {
	return t.deleteWorkspace
}

func NewWorkspacesHttpServerTransports(endpoints WorkspacesEndpoints, serverOptions ...http.ServerOption) WorkspacesHttpServerTransports {
	return &workspacesHttpServerTransports{
		listWorkspaces: http.NewServer(
			endpoints.ListWorkspaces(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &ListWorkspacesRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Parent = fmt.Sprintf("projects/%s/locations/%s", vars.Get("project"), vars.Get("location"))
				if varErr != nil {
					return nil, varErr
				}
				queries := r.URL.Query()
				var queryErr error
				req.PageSize, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "page_size"))
				req.PageToken = queries.Get("page_token")
				if queryErr != nil {
					return nil, queryErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*ListWorkspacesResponse)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			append([]http.ServerOption{
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")
				}),
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpServer)
				}),
			}, serverOptions...)...,
		),
		getWorkspace: http.NewServer(
			endpoints.GetWorkspace(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &GetWorkspaceRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("projects/%s/locations/%s/workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("workspac"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Workspace)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			append([]http.ServerOption{
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/GetWorkspace")
				}),
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpServer)
				}),
			}, serverOptions...)...,
		),
		createWorkspace: http.NewServer(
			endpoints.CreateWorkspace(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &CreateWorkspaceRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(&req.Workspace); err != nil {
					return nil, err
				}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Parent = fmt.Sprintf("projects/%s/locations/%s", vars.Get("project"), vars.Get("location"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Workspace)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			append([]http.ServerOption{
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")
				}),
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpServer)
				}),
			}, serverOptions...)...,
		),
		updateWorkspace: http.NewServer(
			endpoints.UpdateWorkspace(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &UpdateWorkspaceRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(&req.Workspace); err != nil {
					return nil, err
				}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("projects/%s/locations/%s/Workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("Workspac"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*Workspace)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			append([]http.ServerOption{
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")
				}),
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpServer)
				}),
			}, serverOptions...)...,
		),
		deleteWorkspace: http.NewServer(
			endpoints.DeleteWorkspace(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &DeleteWorkspaceRequest{}
				vars := urlx.FormFromMap(mux.Vars(r))
				var varErr error
				req.Name = fmt.Sprintf("projects/%s/locations/%s/workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("workspac"))
				if varErr != nil {
					return nil, varErr
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*emptypb.Empty)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			append([]http.ServerOption{
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")
				}),
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpServer)
				}),
			}, serverOptions...)...,
		),
	}
}

type workspacesHttpClientTransports struct {
	listWorkspaces  *http.Client
	getWorkspace    *http.Client
	createWorkspace *http.Client
	updateWorkspace *http.Client
	deleteWorkspace *http.Client
}

func (t *workspacesHttpClientTransports) ListWorkspaces() *http.Client {
	return t.listWorkspaces
}

func (t *workspacesHttpClientTransports) GetWorkspace() *http.Client {
	return t.getWorkspace
}

func (t *workspacesHttpClientTransports) CreateWorkspace() *http.Client {
	return t.createWorkspace
}

func (t *workspacesHttpClientTransports) UpdateWorkspace() *http.Client {
	return t.updateWorkspace
}

func (t *workspacesHttpClientTransports) DeleteWorkspace() *http.Client {
	return t.deleteWorkspace
}

func NewWorkspacesHttpClientTransports(scheme string, instance string, clientOptions ...http.ClientOption) WorkspacesHttpClientTransports {
	router := mux.NewRouter()
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").Methods("GET").Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").Methods("GET").Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").Methods("POST").Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").Methods("PATCH").Path("/v1/projects/{project}/locations/{location}/Workspaces/{Workspac}")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").Methods("DELETE").Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	return &workspacesHttpClientTransports{
		listWorkspaces: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*ListWorkspacesRequest)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				namedPathParameter := req.GetParent()
				namedPathValues := strings.Split(namedPathParameter, "/")
				if len(namedPathValues) != 4 {
					return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
				}
				pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3])
				path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").URLPath(pairs...)
				if err != nil {
					return nil, err
				}
				queries := url.Values{}
				queries["page_size"] = append(queries["page_size"], strconvx.FormatInt(req.GetPageSize(), 10))
				queries["page_token"] = append(queries["page_token"], req.GetPageToken())
				target := &url.URL{
					Scheme:   scheme,
					Host:     instance,
					Path:     path.Path,
					RawQuery: queries.Encode(),
				}
				r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				resp := &ListWorkspacesResponse{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			append([]http.ClientOption{
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")
				}),
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpClient)
				}),
			}, clientOptions...)...,
		),
		getWorkspace: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*GetWorkspaceRequest)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				namedPathParameter := req.GetName()
				namedPathValues := strings.Split(namedPathParameter, "/")
				if len(namedPathValues) != 6 {
					return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
				}
				pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "workspac", namedPathValues[5])
				path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").URLPath(pairs...)
				if err != nil {
					return nil, err
				}
				queries := url.Values{}
				target := &url.URL{
					Scheme:   scheme,
					Host:     instance,
					Path:     path.Path,
					RawQuery: queries.Encode(),
				}
				r, err := http1.NewRequestWithContext(ctx, "GET", target.String(), body)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				resp := &Workspace{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			append([]http.ClientOption{
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/GetWorkspace")
				}),
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpClient)
				}),
			}, clientOptions...)...,
		),
		createWorkspace: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*CreateWorkspaceRequest)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var bodyBuf bytes.Buffer
				if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetWorkspace()); err != nil {
					return nil, err
				}
				body = &bodyBuf
				contentType := "application/json; charset=utf-8"
				var pairs []string
				namedPathParameter := req.GetParent()
				namedPathValues := strings.Split(namedPathParameter, "/")
				if len(namedPathValues) != 4 {
					return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
				}
				pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3])
				path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").URLPath(pairs...)
				if err != nil {
					return nil, err
				}
				queries := url.Values{}
				target := &url.URL{
					Scheme:   scheme,
					Host:     instance,
					Path:     path.Path,
					RawQuery: queries.Encode(),
				}
				r, err := http1.NewRequestWithContext(ctx, "POST", target.String(), body)
				if err != nil {
					return nil, err
				}
				r.Header.Set("Content-Type", contentType)
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				resp := &Workspace{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			append([]http.ClientOption{
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")
				}),
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpClient)
				}),
			}, clientOptions...)...,
		),
		updateWorkspace: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*UpdateWorkspaceRequest)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var bodyBuf bytes.Buffer
				if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetWorkspace()); err != nil {
					return nil, err
				}
				body = &bodyBuf
				contentType := "application/json; charset=utf-8"
				var pairs []string
				namedPathParameter := req.GetName()
				namedPathValues := strings.Split(namedPathParameter, "/")
				if len(namedPathValues) != 6 {
					return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
				}
				pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "Workspac", namedPathValues[5])
				path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").URLPath(pairs...)
				if err != nil {
					return nil, err
				}
				queries := url.Values{}
				target := &url.URL{
					Scheme:   scheme,
					Host:     instance,
					Path:     path.Path,
					RawQuery: queries.Encode(),
				}
				r, err := http1.NewRequestWithContext(ctx, "PATCH", target.String(), body)
				if err != nil {
					return nil, err
				}
				r.Header.Set("Content-Type", contentType)
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				resp := &Workspace{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			append([]http.ClientOption{
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")
				}),
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpClient)
				}),
			}, clientOptions...)...,
		),
		deleteWorkspace: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*DeleteWorkspaceRequest)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				namedPathParameter := req.GetName()
				namedPathValues := strings.Split(namedPathParameter, "/")
				if len(namedPathValues) != 6 {
					return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
				}
				pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "workspac", namedPathValues[5])
				path, err := router.Get("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").URLPath(pairs...)
				if err != nil {
					return nil, err
				}
				queries := url.Values{}
				target := &url.URL{
					Scheme:   scheme,
					Host:     instance,
					Path:     path.Path,
					RawQuery: queries.Encode(),
				}
				r, err := http1.NewRequestWithContext(ctx, "DELETE", target.String(), body)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				resp := &emptypb.Empty{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			append([]http.ClientOption{
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")
				}),
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpClient)
				}),
			}, clientOptions...)...,
		),
	}
}

func NewWorkspacesHttpServerHandler(endpoints WorkspacesHttpServerTransports) http1.Handler {
	router := mux.NewRouter()
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").Methods("GET").Path("/v1/projects/{project}/locations/{location}/workspaces").Handler(endpoints.ListWorkspaces())
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").Methods("GET").Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}").Handler(endpoints.GetWorkspace())
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").Methods("POST").Path("/v1/projects/{project}/locations/{location}/workspaces").Handler(endpoints.CreateWorkspace())
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").Methods("PATCH").Path("/v1/projects/{project}/locations/{location}/Workspaces/{Workspac}").Handler(endpoints.UpdateWorkspace())
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").Methods("DELETE").Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}").Handler(endpoints.DeleteWorkspace())
	return router
}

type workspacesHttpClient struct {
	listWorkspaces  endpoint.Endpoint
	getWorkspace    endpoint.Endpoint
	createWorkspace endpoint.Endpoint
	updateWorkspace endpoint.Endpoint
	deleteWorkspace endpoint.Endpoint
}

func (c *workspacesHttpClient) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	rep, err := c.listWorkspaces(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListWorkspacesResponse), nil
}

func (c *workspacesHttpClient) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	rep, err := c.getWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesHttpClient) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	rep, err := c.createWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesHttpClient) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	rep, err := c.updateWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesHttpClient) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	rep, err := c.deleteWorkspace(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesHttpClient(transports WorkspacesHttpClientTransports, middlewares ...endpoint.Middleware) WorkspacesService {
	return &workspacesHttpClient{
		listWorkspaces:  endpointx.Chain(transports.ListWorkspaces().Endpoint(), middlewares...),
		getWorkspace:    endpointx.Chain(transports.GetWorkspace().Endpoint(), middlewares...),
		createWorkspace: endpointx.Chain(transports.CreateWorkspace().Endpoint(), middlewares...),
		updateWorkspace: endpointx.Chain(transports.UpdateWorkspace().Endpoint(), middlewares...),
		deleteWorkspace: endpointx.Chain(transports.DeleteWorkspace().Endpoint(), middlewares...),
	}
}