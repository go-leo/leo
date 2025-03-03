// Code generated by protoc-gen-go-leo. DO NOT EDIT.

package endpointsapis

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	httpx1 "github.com/go-leo/gox/netx/httpx"
	urlx "github.com/go-leo/gox/netx/urlx"
	strconvx "github.com/go-leo/gox/strconvx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	coder "github.com/go-leo/leo/v3/transportx/httpx/coder"
	mux "github.com/gorilla/mux"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
	strings "strings"
)

func appendWorkspacesHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").
		Methods(http.MethodGet).
		Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").
		Methods(http.MethodGet).
		Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").
		Methods(http.MethodPost).
		Path("/v1/projects/{project}/locations/{location}/workspaces")
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").
		Methods(http.MethodPatch).
		Path("/v1/projects/{project}/locations/{location}/Workspaces/{Workspac}")
	router.NewRoute().
		Name("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").
		Methods(http.MethodDelete).
		Path("/v1/projects/{project}/locations/{location}/workspaces/{workspac}")
	return router
}
func AppendWorkspacesHttpServerRoutes(router *mux.Router, svc WorkspacesService, opts ...httpx.ServerOption) *mux.Router {
	options := httpx.NewServerOptions(opts...)
	endpoints := &workspacesServerEndpoints{
		svc:         svc,
		middlewares: options.Middlewares(),
	}
	transports := &workspacesHttpServerTransports{
		endpoints: endpoints,
		requestDecoder: workspacesHttpServerRequestDecoder{
			unmarshalOptions: options.UnmarshalOptions(),
		},
		responseEncoder: workspacesHttpServerResponseEncoder{
			marshalOptions: options.MarshalOptions(),
		},
	}
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
	requestEncoder := &workspacesHttpClientRequestEncoder{
		router: appendWorkspacesHttpRoutes(mux.NewRouter()),
		scheme: options.Scheme(),
	}
	responseDecoder := &workspacesHttpClientResponseDecoder{}
	transports := &workspacesHttpClientTransports{
		clientOptions:   options.ClientTransportOptions(),
		middlewares:     options.Middlewares(),
		requestEncoder:  requestEncoder,
		responseDecoder: responseDecoder,
	}
	factories := &workspacesFactories{
		transports: transports,
	}
	endpointer := &workspacesEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &workspacesBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &workspacesClientEndpoints{
		balancers: balancers,
	}
	return &workspacesClientService{
		endpoints:     endpoints,
		transportName: httpx.HttpClient,
	}
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
	)
}

type workspacesHttpServerRequestDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder workspacesHttpServerRequestDecoder) ListWorkspaces() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &ListWorkspacesRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		req.Parent = fmt.Sprintf("projects/%s/locations/%s", vars.Get("project"), vars.Get("location"))
		queries := r.URL.Query()
		var queryErr error
		req.PageSize, queryErr = coder.DecodeForm[int32](queryErr, queries, "page_size", urlx.GetInt[int32])
		req.PageToken = queries.Get("page_token")
		if queryErr != nil {
			return nil, queryErr
		}
		return req, nil
	}
}
func (decoder workspacesHttpServerRequestDecoder) GetWorkspace() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &GetWorkspaceRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		req.Name = fmt.Sprintf("projects/%s/locations/%s/workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("workspac"))
		return req, nil
	}
}
func (decoder workspacesHttpServerRequestDecoder) CreateWorkspace() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &CreateWorkspaceRequest{}
		req.Workspace = &Workspace{}
		if err := coder.DecodeMessageFromRequest(ctx, r, req.GetWorkspace(), decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		vars := urlx.FormFromMap(mux.Vars(r))
		req.Parent = fmt.Sprintf("projects/%s/locations/%s", vars.Get("project"), vars.Get("location"))
		return req, nil
	}
}
func (decoder workspacesHttpServerRequestDecoder) UpdateWorkspace() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &UpdateWorkspaceRequest{}
		req.Workspace = &Workspace{}
		if err := coder.DecodeMessageFromRequest(ctx, r, req.GetWorkspace(), decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		vars := urlx.FormFromMap(mux.Vars(r))
		req.Name = fmt.Sprintf("projects/%s/locations/%s/Workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("Workspac"))
		return req, nil
	}
}
func (decoder workspacesHttpServerRequestDecoder) DeleteWorkspace() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &DeleteWorkspaceRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		req.Name = fmt.Sprintf("projects/%s/locations/%s/workspaces/%s", vars.Get("project"), vars.Get("location"), vars.Get("workspac"))
		return req, nil
	}
}

type workspacesHttpServerResponseEncoder struct {
	marshalOptions protojson.MarshalOptions
}

func (encoder workspacesHttpServerResponseEncoder) ListWorkspaces() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*ListWorkspacesResponse)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder workspacesHttpServerResponseEncoder) GetWorkspace() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*Workspace)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder workspacesHttpServerResponseEncoder) CreateWorkspace() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*Workspace)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder workspacesHttpServerResponseEncoder) UpdateWorkspace() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*Workspace)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder workspacesHttpServerResponseEncoder) DeleteWorkspace() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*emptypb.Empty)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
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

type workspacesHttpClientRequestEncoder struct {
	marshalOptions   protojson.MarshalOptions
	unmarshalOptions protojson.UnmarshalOptions
	router           *mux.Router
	scheme           string
}

func (encoder workspacesHttpClientRequestEncoder) ListWorkspaces(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*ListWorkspacesRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type, %T", obj)
		}
		_ = req
		method := http.MethodGet
		target := &url.URL{
			Scheme: encoder.scheme,
			Host:   instance,
		}
		header := http.Header{}
		var body bytes.Buffer
		var pairs []string
		namedPathParameter := req.GetParent()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 4 {
			return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
		}
		pairs = append(pairs,
			"project", namedPathValues[1],
			"location", namedPathValues[3],
		)
		path, err := encoder.router.Get("/google.example.endpointsapis.v1.Workspaces/ListWorkspaces").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		queries := url.Values{}
		queries["page_size"] = append(queries["page_size"], strconvx.FormatInt(req.GetPageSize(), 10))
		queries["page_token"] = append(queries["page_token"], req.GetPageToken())
		target.RawQuery = queries.Encode()
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx1.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder workspacesHttpClientRequestEncoder) GetWorkspace(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*GetWorkspaceRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type, %T", obj)
		}
		_ = req
		method := http.MethodGet
		target := &url.URL{
			Scheme: encoder.scheme,
			Host:   instance,
		}
		header := http.Header{}
		var body bytes.Buffer
		var pairs []string
		namedPathParameter := req.GetName()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 6 {
			return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
		}
		pairs = append(pairs,
			"project", namedPathValues[1],
			"location", namedPathValues[3],
			"workspac", namedPathValues[5],
		)
		path, err := encoder.router.Get("/google.example.endpointsapis.v1.Workspaces/GetWorkspace").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx1.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder workspacesHttpClientRequestEncoder) CreateWorkspace(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*CreateWorkspaceRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type, %T", obj)
		}
		_ = req
		method := http.MethodPost
		target := &url.URL{
			Scheme: encoder.scheme,
			Host:   instance,
		}
		header := http.Header{}
		var body bytes.Buffer
		if err := coder.EncodeMessageToRequest(ctx, req.GetWorkspace(), header, &body, encoder.marshalOptions); err != nil {
			return nil, err
		}
		var pairs []string
		namedPathParameter := req.GetParent()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 4 {
			return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
		}
		pairs = append(pairs,
			"project", namedPathValues[1],
			"location", namedPathValues[3],
		)
		path, err := encoder.router.Get("/google.example.endpointsapis.v1.Workspaces/CreateWorkspace").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx1.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder workspacesHttpClientRequestEncoder) UpdateWorkspace(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*UpdateWorkspaceRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type, %T", obj)
		}
		_ = req
		method := http.MethodPatch
		target := &url.URL{
			Scheme: encoder.scheme,
			Host:   instance,
		}
		header := http.Header{}
		var body bytes.Buffer
		if err := coder.EncodeMessageToRequest(ctx, req.GetWorkspace(), header, &body, encoder.marshalOptions); err != nil {
			return nil, err
		}
		var pairs []string
		namedPathParameter := req.GetName()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 6 {
			return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
		}
		pairs = append(pairs,
			"project", namedPathValues[1],
			"location", namedPathValues[3],
			"Workspac", namedPathValues[5],
		)
		path, err := encoder.router.Get("/google.example.endpointsapis.v1.Workspaces/UpdateWorkspace").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx1.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder workspacesHttpClientRequestEncoder) DeleteWorkspace(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*DeleteWorkspaceRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type, %T", obj)
		}
		_ = req
		method := http.MethodDelete
		target := &url.URL{
			Scheme: encoder.scheme,
			Host:   instance,
		}
		header := http.Header{}
		var body bytes.Buffer
		var pairs []string
		namedPathParameter := req.GetName()
		namedPathValues := strings.Split(namedPathParameter, "/")
		if len(namedPathValues) != 6 {
			return nil, fmt.Errorf("invalid named path parameter, %s", namedPathParameter)
		}
		pairs = append(pairs,
			"project", namedPathValues[1],
			"location", namedPathValues[3],
			"workspac", namedPathValues[5],
		)
		path, err := encoder.router.Get("/google.example.endpointsapis.v1.Workspaces/DeleteWorkspace").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx1.CopyHeader(r.Header, header)
		return r, nil
	}
}

type workspacesHttpClientResponseDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder workspacesHttpClientResponseDecoder) ListWorkspaces() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &ListWorkspacesResponse{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder workspacesHttpClientResponseDecoder) GetWorkspace() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &Workspace{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder workspacesHttpClientResponseDecoder) CreateWorkspace() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &Workspace{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder workspacesHttpClientResponseDecoder) UpdateWorkspace() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &Workspace{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder workspacesHttpClientResponseDecoder) DeleteWorkspace() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &emptypb.Empty{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
