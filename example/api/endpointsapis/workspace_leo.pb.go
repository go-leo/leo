// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package endpointsapis

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	urlx "github.com/go-leo/gox/netx/urlx"
	strconvx "github.com/go-leo/gox/strconvx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	grpc1 "google.golang.org/grpc"
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
	ListWorkspaces(ctx context.Context) endpoint.Endpoint
	GetWorkspace(ctx context.Context) endpoint.Endpoint
	CreateWorkspace(ctx context.Context) endpoint.Endpoint
	UpdateWorkspace(ctx context.Context) endpoint.Endpoint
	DeleteWorkspace(ctx context.Context) endpoint.Endpoint
}

type WorkspacesClientTransports interface {
	ListWorkspaces() transportx.ClientTransport
	GetWorkspace() transportx.ClientTransport
	CreateWorkspace() transportx.ClientTransport
	UpdateWorkspace() transportx.ClientTransport
	DeleteWorkspace() transportx.ClientTransport
}

type WorkspacesFactories interface {
	ListWorkspaces(middlewares ...endpoint.Middleware) sd.Factory
	GetWorkspace(middlewares ...endpoint.Middleware) sd.Factory
	CreateWorkspace(middlewares ...endpoint.Middleware) sd.Factory
	UpdateWorkspace(middlewares ...endpoint.Middleware) sd.Factory
	DeleteWorkspace(middlewares ...endpoint.Middleware) sd.Factory
}

type WorkspacesEndpointers interface {
	ListWorkspaces() sd.Endpointer
	GetWorkspace() sd.Endpointer
	CreateWorkspace() sd.Endpointer
	UpdateWorkspace() sd.Endpointer
	DeleteWorkspace() sd.Endpointer
}

type workspacesServerEndpoints struct {
	svc         WorkspacesService
	middlewares []endpoint.Middleware
}

func (e *workspacesServerEndpoints) ListWorkspaces(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.ListWorkspaces(ctx, request.(*ListWorkspacesRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *workspacesServerEndpoints) GetWorkspace(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.GetWorkspace(ctx, request.(*GetWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *workspacesServerEndpoints) CreateWorkspace(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.CreateWorkspace(ctx, request.(*CreateWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *workspacesServerEndpoints) UpdateWorkspace(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.UpdateWorkspace(ctx, request.(*UpdateWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *workspacesServerEndpoints) DeleteWorkspace(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.DeleteWorkspace(ctx, request.(*DeleteWorkspaceRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewWorkspacesServerEndpoints(svc WorkspacesService, middlewares ...endpoint.Middleware) WorkspacesEndpoints {
	return &workspacesServerEndpoints{svc: svc, middlewares: middlewares}
}

type workspacesClientEndpoints struct {
	transports  WorkspacesClientTransports
	middlewares []endpoint.Middleware
}

func (e *workspacesClientEndpoints) ListWorkspaces(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.ListWorkspaces().Endpoint(ctx), e.middlewares...)
}

func (e *workspacesClientEndpoints) GetWorkspace(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.GetWorkspace().Endpoint(ctx), e.middlewares...)
}

func (e *workspacesClientEndpoints) CreateWorkspace(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.CreateWorkspace().Endpoint(ctx), e.middlewares...)
}

func (e *workspacesClientEndpoints) UpdateWorkspace(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.UpdateWorkspace().Endpoint(ctx), e.middlewares...)
}

func (e *workspacesClientEndpoints) DeleteWorkspace(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.DeleteWorkspace().Endpoint(ctx), e.middlewares...)
}

func NewWorkspacesClientEndpoints(transports WorkspacesClientTransports, middlewares ...endpoint.Middleware) WorkspacesEndpoints {
	return &workspacesClientEndpoints{transports: transports, middlewares: middlewares}
}

// =========================== cqrs ===========================

// =========================== grpc server ===========================

type WorkspacesGrpcServerTransports interface {
	ListWorkspaces() *grpc.Server
	GetWorkspace() *grpc.Server
	CreateWorkspace() *grpc.Server
	UpdateWorkspace() *grpc.Server
	DeleteWorkspace() *grpc.Server
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

func NewWorkspacesGrpcServerTransports(endpoints WorkspacesEndpoints) WorkspacesGrpcServerTransports {
	return &workspacesGrpcServerTransports{
		listWorkspaces: grpc.NewServer(
			endpoints.ListWorkspaces(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		getWorkspace: grpc.NewServer(
			endpoints.GetWorkspace(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/GetWorkspace")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		createWorkspace: grpc.NewServer(
			endpoints.CreateWorkspace(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		updateWorkspace: grpc.NewServer(
			endpoints.UpdateWorkspace(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		deleteWorkspace: grpc.NewServer(
			endpoints.DeleteWorkspace(context.TODO()),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
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

// =========================== grpc client ===========================

type workspacesGrpcClientTransports struct {
	listWorkspaces  transportx.ClientTransport
	getWorkspace    transportx.ClientTransport
	createWorkspace transportx.ClientTransport
	updateWorkspace transportx.ClientTransport
	deleteWorkspace transportx.ClientTransport
}

func (t *workspacesGrpcClientTransports) ListWorkspaces() transportx.ClientTransport {
	return t.listWorkspaces
}

func (t *workspacesGrpcClientTransports) GetWorkspace() transportx.ClientTransport {
	return t.getWorkspace
}

func (t *workspacesGrpcClientTransports) CreateWorkspace() transportx.ClientTransport {
	return t.createWorkspace
}

func (t *workspacesGrpcClientTransports) UpdateWorkspace() transportx.ClientTransport {
	return t.updateWorkspace
}

func (t *workspacesGrpcClientTransports) DeleteWorkspace() transportx.ClientTransport {
	return t.deleteWorkspace
}

func NewWorkspacesGrpcClientTransports(conn *grpc1.ClientConn) WorkspacesClientTransports {
	return &workspacesGrpcClientTransports{
		listWorkspaces: grpcx.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"ListWorkspaces",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			ListWorkspacesResponse{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
		getWorkspace: grpcx.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"GetWorkspace",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			Workspace{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
		createWorkspace: grpcx.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"CreateWorkspace",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			Workspace{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
		updateWorkspace: grpcx.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"UpdateWorkspace",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			Workspace{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
		deleteWorkspace: grpcx.NewClient(
			conn,
			"google.example.endpointsapis.v1.Workspaces",
			"DeleteWorkspace",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
	}
}

type workspacesGrpcClient struct {
	endpoints WorkspacesEndpoints
}

func (c *workspacesGrpcClient) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.ListWorkspaces(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*ListWorkspacesResponse), nil
}

func (c *workspacesGrpcClient) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/GetWorkspace")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.GetWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.CreateWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.UpdateWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*Workspace), nil
}

func (c *workspacesGrpcClient) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.endpoints.DeleteWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesGrpcClient(endpoints WorkspacesEndpoints) WorkspacesService {
	return &workspacesGrpcClient{endpoints: endpoints}
}

// =========================== http server ===========================

type WorkspacesHttpServerTransports interface {
	ListWorkspaces() *http.Server
	GetWorkspace() *http.Server
	CreateWorkspace() *http.Server
	UpdateWorkspace() *http.Server
	DeleteWorkspace() *http.Server
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

func NewWorkspacesHttpServerTransports(endpoints WorkspacesEndpoints) WorkspacesHttpServerTransports {
	return &workspacesHttpServerTransports{
		listWorkspaces: http.NewServer(
			endpoints.ListWorkspaces(context.TODO()),
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
			http.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		getWorkspace: http.NewServer(
			endpoints.GetWorkspace(context.TODO()),
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
			http.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/GetWorkspace")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		createWorkspace: http.NewServer(
			endpoints.CreateWorkspace(context.TODO()),
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
			http.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		updateWorkspace: http.NewServer(
			endpoints.UpdateWorkspace(context.TODO()),
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
			http.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		deleteWorkspace: http.NewServer(
			endpoints.DeleteWorkspace(context.TODO()),
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
			http.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
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

// =========================== http client ===========================

type workspacesHttpClientTransports struct {
	listWorkspaces  transportx.ClientTransport
	getWorkspace    transportx.ClientTransport
	createWorkspace transportx.ClientTransport
	updateWorkspace transportx.ClientTransport
	deleteWorkspace transportx.ClientTransport
}

func (t *workspacesHttpClientTransports) ListWorkspaces() transportx.ClientTransport {
	return t.listWorkspaces
}

func (t *workspacesHttpClientTransports) GetWorkspace() transportx.ClientTransport {
	return t.getWorkspace
}

func (t *workspacesHttpClientTransports) CreateWorkspace() transportx.ClientTransport {
	return t.createWorkspace
}

func (t *workspacesHttpClientTransports) UpdateWorkspace() transportx.ClientTransport {
	return t.updateWorkspace
}

func (t *workspacesHttpClientTransports) DeleteWorkspace() transportx.ClientTransport {
	return t.deleteWorkspace
}

func NewWorkspacesHttpClientTransports(scheme string, instance string) WorkspacesClientTransports {
	router := mux.NewRouter()
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").Methods("GET").Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").Methods("GET").Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").Methods("POST").Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").Methods("PATCH").Path("/v1/projects/{project}/locations/{location}/Workspaces/{Workspac}")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").Methods("DELETE").Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	return &workspacesHttpClientTransports{
		listWorkspaces: httpx.NewClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
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
			func(ctx context.Context, r *http1.Response) (any, error) {
				if httpx.IsErrorResponse(r) {
					return nil, httpx.ErrorDecoder(ctx, r)
				}
				resp := &ListWorkspacesResponse{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.OutgoingMetadata),
		),
		getWorkspace: httpx.NewClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
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
			func(ctx context.Context, r *http1.Response) (any, error) {
				if httpx.IsErrorResponse(r) {
					return nil, httpx.ErrorDecoder(ctx, r)
				}
				resp := &Workspace{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.OutgoingMetadata),
		),
		createWorkspace: httpx.NewClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
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
			func(ctx context.Context, r *http1.Response) (any, error) {
				if httpx.IsErrorResponse(r) {
					return nil, httpx.ErrorDecoder(ctx, r)
				}
				resp := &Workspace{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.OutgoingMetadata),
		),
		updateWorkspace: httpx.NewClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
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
			func(ctx context.Context, r *http1.Response) (any, error) {
				if httpx.IsErrorResponse(r) {
					return nil, httpx.ErrorDecoder(ctx, r)
				}
				resp := &Workspace{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.OutgoingMetadata),
		),
		deleteWorkspace: httpx.NewClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
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
			func(ctx context.Context, r *http1.Response) (any, error) {
				if httpx.IsErrorResponse(r) {
					return nil, httpx.ErrorDecoder(ctx, r)
				}
				resp := &emptypb.Empty{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.OutgoingMetadata),
		),
	}
}

type workspacesHttpClient struct {
	endpoints WorkspacesEndpoints
}

func (c *workspacesHttpClient) ListWorkspaces(ctx context.Context, request *ListWorkspacesRequest) (*ListWorkspacesResponse, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.ListWorkspaces(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*ListWorkspacesResponse), nil
}

func (c *workspacesHttpClient) GetWorkspace(ctx context.Context, request *GetWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/GetWorkspace")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.GetWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesHttpClient) CreateWorkspace(ctx context.Context, request *CreateWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.CreateWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesHttpClient) UpdateWorkspace(ctx context.Context, request *UpdateWorkspaceRequest) (*Workspace, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.UpdateWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*Workspace), nil
}

func (c *workspacesHttpClient) DeleteWorkspace(ctx context.Context, request *DeleteWorkspaceRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.endpoints.DeleteWorkspace(ctx)(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewWorkspacesHttpClient(endpoints WorkspacesEndpoints) WorkspacesService {
	return &workspacesGrpcClient{endpoints: endpoints}
}
