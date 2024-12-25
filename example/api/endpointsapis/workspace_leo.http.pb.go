// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package endpointsapis

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	urlx "github.com/go-leo/gox/netx/urlx"
	strconvx "github.com/go-leo/gox/strconvx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
	strings "strings"
)

func appendWorkspacesHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").Methods("GET").Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").Methods("GET").Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").Methods("POST").Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").Methods("PATCH").Path("/v1/projects/{project}/locations/{location}/Workspaces/{Workspac}")
	router.NewRoute().Name("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").Methods("DELETE").Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	return router
}
func AppendWorkspacesHttpRoutes(router *mux.Router, svc WorkspacesService, middlewares ...endpoint.Middleware) *mux.Router {
	transports := newWorkspacesHttpServerTransports(svc, middlewares...)
	router = appendWorkspacesHttpRoutes(router)
	router.Get("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").Handler(transports.ListWorkspaces())
	router.Get("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").Handler(transports.GetWorkspace())
	router.Get("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").Handler(transports.CreateWorkspace())
	router.Get("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").Handler(transports.UpdateWorkspace())
	router.Get("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").Handler(transports.DeleteWorkspace())
	return router
}

func NewWorkspacesHttpClient(target string, opts ...httpx.ClientOption) WorkspacesService {
	options := httpx.NewClientOptions(opts...)
	transports := newWorkspacesHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newWorkspacesClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newWorkspacesClientService(endpoints, httpx.HttpClient)
}

type WorkspacesHttpServerTransports interface {
	ListWorkspaces() http.Handler
	GetWorkspace() http.Handler
	CreateWorkspace() http.Handler
	UpdateWorkspace() http.Handler
	DeleteWorkspace() http.Handler
}

type WorkspacesHttpServerRequestDecoder interface {
	ListWorkspaces() http1.DecodeRequestFunc
	GetWorkspace() http1.DecodeRequestFunc
	CreateWorkspace() http1.DecodeRequestFunc
	UpdateWorkspace() http1.DecodeRequestFunc
	DeleteWorkspace() http1.DecodeRequestFunc
}

type WorkspacesHttpServerResponseEncoder interface {
	ListWorkspaces() http1.EncodeResponseFunc
	GetWorkspace() http1.EncodeResponseFunc
	CreateWorkspace() http1.EncodeResponseFunc
	UpdateWorkspace() http1.EncodeResponseFunc
	DeleteWorkspace() http1.EncodeResponseFunc
}

type WorkspacesHttpClientRequestEncoder interface {
	ListWorkspaces(instance string) http1.CreateRequestFunc
	GetWorkspace(instance string) http1.CreateRequestFunc
	CreateWorkspace(instance string) http1.CreateRequestFunc
	UpdateWorkspace(instance string) http1.CreateRequestFunc
	DeleteWorkspace(instance string) http1.CreateRequestFunc
}

type WorkspacesHttpClientResponseDecoder interface {
	ListWorkspaces() http1.DecodeResponseFunc
	GetWorkspace() http1.DecodeResponseFunc
	CreateWorkspace() http1.DecodeResponseFunc
	UpdateWorkspace() http1.DecodeResponseFunc
	DeleteWorkspace() http1.DecodeResponseFunc
}

type workspacesHttpServerTransports struct {
	endpoints       WorkspacesServerEndpoints
	requestDecoder  WorkspacesHttpServerRequestDecoder
	responseEncoder WorkspacesHttpServerResponseEncoder
}

func (t *workspacesHttpServerTransports) ListWorkspaces() http.Handler {
	return http1.NewServer(
		t.endpoints.ListWorkspaces(context.TODO()),
		t.requestDecoder.ListWorkspaces(),
		t.responseEncoder.ListWorkspaces(),
		http1.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *workspacesHttpServerTransports) GetWorkspace() http.Handler {
	return http1.NewServer(
		t.endpoints.GetWorkspace(context.TODO()),
		t.requestDecoder.GetWorkspace(),
		t.responseEncoder.GetWorkspace(),
		http1.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/GetWorkspace")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *workspacesHttpServerTransports) CreateWorkspace() http.Handler {
	return http1.NewServer(
		t.endpoints.CreateWorkspace(context.TODO()),
		t.requestDecoder.CreateWorkspace(),
		t.responseEncoder.CreateWorkspace(),
		http1.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *workspacesHttpServerTransports) UpdateWorkspace() http.Handler {
	return http1.NewServer(
		t.endpoints.UpdateWorkspace(context.TODO()),
		t.requestDecoder.UpdateWorkspace(),
		t.responseEncoder.UpdateWorkspace(),
		http1.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *workspacesHttpServerTransports) DeleteWorkspace() http.Handler {
	return http1.NewServer(
		t.endpoints.DeleteWorkspace(context.TODO()),
		t.requestDecoder.DeleteWorkspace(),
		t.responseEncoder.DeleteWorkspace(),
		http1.ServerBefore(httpx.EndpointInjector("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func newWorkspacesHttpServerTransports(svc WorkspacesService, middlewares ...endpoint.Middleware) WorkspacesHttpServerTransports {
	endpoints := newWorkspacesServerEndpoints(svc, middlewares...)
	return &workspacesHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  workspacesHttpServerRequestDecoder{},
		responseEncoder: workspacesHttpServerResponseEncoder{},
	}
}

type workspacesHttpServerRequestDecoder struct{}

func (workspacesHttpServerRequestDecoder) ListWorkspaces() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &ListWorkspacesRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.Parent = fmt.Sprintf("projects/%s/locations/%s", vars.Get("project"), vars.Get("location"))
		if varErr != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
		}
		queries := r.URL.Query()
		var queryErr error
		req.PageSize, queryErr = errorx.Break[int32](queryErr)(urlx.GetInt[int32](queries, "page_size"))
		req.PageToken = queries.Get("page_token")
		if queryErr != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(queryErr))
		}
		return req, nil
	}
}
func (workspacesHttpServerRequestDecoder) GetWorkspace() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &GetWorkspaceRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.Name = fmt.Sprintf("projects/%s/locations/%s/workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("workspac"))
		if varErr != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
		}
		return req, nil
	}
}
func (workspacesHttpServerRequestDecoder) CreateWorkspace() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &CreateWorkspaceRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(&req.Workspace); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.Parent = fmt.Sprintf("projects/%s/locations/%s", vars.Get("project"), vars.Get("location"))
		if varErr != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
		}
		return req, nil
	}
}
func (workspacesHttpServerRequestDecoder) UpdateWorkspace() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &UpdateWorkspaceRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(&req.Workspace); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.Name = fmt.Sprintf("projects/%s/locations/%s/Workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("Workspac"))
		if varErr != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
		}
		return req, nil
	}
}
func (workspacesHttpServerRequestDecoder) DeleteWorkspace() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &DeleteWorkspaceRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.Name = fmt.Sprintf("projects/%s/locations/%s/workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("workspac"))
		if varErr != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
		}
		return req, nil
	}
}

type workspacesHttpServerResponseEncoder struct{}

func (workspacesHttpServerResponseEncoder) ListWorkspaces() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*ListWorkspacesResponse)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}
func (workspacesHttpServerResponseEncoder) GetWorkspace() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*Workspace)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}
func (workspacesHttpServerResponseEncoder) CreateWorkspace() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*Workspace)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}
func (workspacesHttpServerResponseEncoder) UpdateWorkspace() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*Workspace)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}
func (workspacesHttpServerResponseEncoder) DeleteWorkspace() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*emptypb.Empty)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}
}

type workspacesHttpClientTransports struct {
	clientOptions   []http1.ClientOption
	middlewares     []endpoint.Middleware
	requestEncoder  WorkspacesHttpClientRequestEncoder
	responseDecoder WorkspacesHttpClientResponseDecoder
}

func (t *workspacesHttpClientTransports) ListWorkspaces(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.ListWorkspaces(instance),
		t.responseDecoder.ListWorkspaces(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *workspacesHttpClientTransports) GetWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.GetWorkspace(instance),
		t.responseDecoder.GetWorkspace(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *workspacesHttpClientTransports) CreateWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.CreateWorkspace(instance),
		t.responseDecoder.CreateWorkspace(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *workspacesHttpClientTransports) UpdateWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.UpdateWorkspace(instance),
		t.responseDecoder.UpdateWorkspace(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *workspacesHttpClientTransports) DeleteWorkspace(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.DeleteWorkspace(instance),
		t.responseDecoder.DeleteWorkspace(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newWorkspacesHttpClientTransports(scheme string, clientOptions []http1.ClientOption, middlewares []endpoint.Middleware) WorkspacesClientTransports {
	return &workspacesHttpClientTransports{
		clientOptions: clientOptions,
		middlewares:   middlewares,
		requestEncoder: workspacesHttpClientRequestEncoder{
			scheme: scheme,
			router: appendWorkspacesHttpRoutes(mux.NewRouter()),
		},
		responseDecoder: workspacesHttpClientResponseDecoder{},
	}
}

type workspacesHttpClientRequestEncoder struct {
	router *mux.Router
	scheme string
}

func (e workspacesHttpClientRequestEncoder) ListWorkspaces(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*ListWorkspacesRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var pairs []string
		namedPathParameter := req.GetParent()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 4 {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
		}
		pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3])
		path, err := e.router.Get("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").URLPath(pairs...)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		queries := url.Values{}
		queries["page_size"] = append(queries["page_size"], strconvx.FormatInt(req.GetPageSize(), 10))
		queries["page_token"] = append(queries["page_token"], req.GetPageToken())
		target := &url.URL{
			Scheme:   e.scheme,
			Host:     instance,
			Path:     path.Path,
			RawQuery: queries.Encode(),
		}
		r, err := http.NewRequestWithContext(ctx, "GET", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return r, nil
	}
}
func (e workspacesHttpClientRequestEncoder) GetWorkspace(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*GetWorkspaceRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var pairs []string
		namedPathParameter := req.GetName()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 6 {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
		}
		pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "workspac", namedPathValues[5])
		path, err := e.router.Get("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").URLPath(pairs...)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		queries := url.Values{}
		target := &url.URL{
			Scheme:   e.scheme,
			Host:     instance,
			Path:     path.Path,
			RawQuery: queries.Encode(),
		}
		r, err := http.NewRequestWithContext(ctx, "GET", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return r, nil
	}
}
func (e workspacesHttpClientRequestEncoder) CreateWorkspace(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*CreateWorkspaceRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var bodyBuf bytes.Buffer
		if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetWorkspace()); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		body = &bodyBuf
		contentType := "application/json; charset=utf-8"
		var pairs []string
		namedPathParameter := req.GetParent()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 4 {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
		}
		pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3])
		path, err := e.router.Get("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").URLPath(pairs...)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		queries := url.Values{}
		target := &url.URL{
			Scheme:   e.scheme,
			Host:     instance,
			Path:     path.Path,
			RawQuery: queries.Encode(),
		}
		r, err := http.NewRequestWithContext(ctx, "POST", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		r.Header.Set("Content-Type", contentType)
		return r, nil
	}
}
func (e workspacesHttpClientRequestEncoder) UpdateWorkspace(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*UpdateWorkspaceRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var bodyBuf bytes.Buffer
		if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetWorkspace()); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		body = &bodyBuf
		contentType := "application/json; charset=utf-8"
		var pairs []string
		namedPathParameter := req.GetName()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 6 {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
		}
		pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "Workspac", namedPathValues[5])
		path, err := e.router.Get("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").URLPath(pairs...)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		queries := url.Values{}
		target := &url.URL{
			Scheme:   e.scheme,
			Host:     instance,
			Path:     path.Path,
			RawQuery: queries.Encode(),
		}
		r, err := http.NewRequestWithContext(ctx, "PATCH", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		r.Header.Set("Content-Type", contentType)
		return r, nil
	}
}
func (e workspacesHttpClientRequestEncoder) DeleteWorkspace(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*DeleteWorkspaceRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var pairs []string
		namedPathParameter := req.GetName()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 6 {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
		}
		pairs = append(pairs, "project", namedPathValues[1], "location", namedPathValues[3], "workspac", namedPathValues[5])
		path, err := e.router.Get("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").URLPath(pairs...)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		queries := url.Values{}
		target := &url.URL{
			Scheme:   e.scheme,
			Host:     instance,
			Path:     path.Path,
			RawQuery: queries.Encode(),
		}
		r, err := http.NewRequestWithContext(ctx, "DELETE", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return r, nil
	}
}

type workspacesHttpClientResponseDecoder struct{}

func (workspacesHttpClientResponseDecoder) ListWorkspaces() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &ListWorkspacesResponse{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (workspacesHttpClientResponseDecoder) GetWorkspace() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &Workspace{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (workspacesHttpClientResponseDecoder) CreateWorkspace() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &Workspace{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (workspacesHttpClientResponseDecoder) UpdateWorkspace() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &Workspace{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (workspacesHttpClientResponseDecoder) DeleteWorkspace() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if httpx.IsErrorResponse(r) {
			return nil, httpx.ErrorDecoder(ctx, r)
		}
		resp := &emptypb.Empty{}
		if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
