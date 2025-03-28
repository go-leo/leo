// Code generated by protoc-gen-go-leo. DO NOT EDIT.

package demo

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	httpx "github.com/go-leo/gox/netx/httpx"
	urlx "github.com/go-leo/gox/netx/urlx"
	strconvx "github.com/go-leo/gox/strconvx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	metadatax "github.com/go-leo/leo/v3/metadatax"
	stainx "github.com/go-leo/leo/v3/stainx"
	timeoutx "github.com/go-leo/leo/v3/timeoutx"
	httptransportx "github.com/go-leo/leo/v3/transportx/httptransportx"
	coder "github.com/go-leo/leo/v3/transportx/httptransportx/coder"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
)

func appendDemoHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().
		Name("/leo.example.demo.v1.Demo/CreateUser").
		Methods(http.MethodPost).
		Path("/v1/user")
	router.NewRoute().
		Name("/leo.example.demo.v1.Demo/DeleteUser").
		Methods(http.MethodDelete).
		Path("/v1/user/{user_id}")
	router.NewRoute().
		Name("/leo.example.demo.v1.Demo/UpdateUser").
		Methods(http.MethodPut).
		Path("/v1/user/{user_id}")
	router.NewRoute().
		Name("/leo.example.demo.v1.Demo/GetUser").
		Methods(http.MethodGet).
		Path("/v1/user/{user_id}")
	router.NewRoute().
		Name("/leo.example.demo.v1.Demo/GetUsers").
		Methods(http.MethodGet).
		Path("/v1/users")
	router.NewRoute().
		Name("/leo.example.demo.v1.Demo/UploadUserAvatar").
		Methods(http.MethodPost).
		Path("/v1/user/{user_id}")
	router.NewRoute().
		Name("/leo.example.demo.v1.Demo/GetUserAvatar").
		Methods(http.MethodGet).
		Path("/v1/users/{user_id}")
	return router
}
func AppendDemoHttpServerRoutes(router *mux.Router, svc DemoService, opts ...httptransportx.ServerOption) *mux.Router {
	options := httptransportx.NewServerOptions(opts...)
	endpoints := &demoServerEndpoints{
		svc:         svc,
		middlewares: options.Middlewares(),
	}
	requestDecoder := demoHttpServerRequestDecoder{
		unmarshalOptions: options.UnmarshalOptions(),
	}
	responseEncoder := demoHttpServerResponseEncoder{
		marshalOptions: options.MarshalOptions(),
	}
	transports := &demoHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  requestDecoder,
		responseEncoder: responseEncoder,
	}
	router = appendDemoHttpRoutes(router)
	router.Get("/leo.example.demo.v1.Demo/CreateUser").Handler(transports.CreateUser())
	router.Get("/leo.example.demo.v1.Demo/DeleteUser").Handler(transports.DeleteUser())
	router.Get("/leo.example.demo.v1.Demo/UpdateUser").Handler(transports.UpdateUser())
	router.Get("/leo.example.demo.v1.Demo/GetUser").Handler(transports.GetUser())
	router.Get("/leo.example.demo.v1.Demo/GetUsers").Handler(transports.GetUsers())
	router.Get("/leo.example.demo.v1.Demo/UploadUserAvatar").Handler(transports.UploadUserAvatar())
	router.Get("/leo.example.demo.v1.Demo/GetUserAvatar").Handler(transports.GetUserAvatar())
	return router
}

func NewDemoHttpClient(target string, opts ...httptransportx.ClientOption) DemoService {
	options := httptransportx.NewClientOptions(opts...)
	requestEncoder := &demoHttpClientRequestEncoder{
		marshalOptions: options.MarshalOptions(),
		router:         appendDemoHttpRoutes(mux.NewRouter()),
		scheme:         options.Scheme(),
	}
	responseDecoder := &demoHttpClientResponseDecoder{
		unmarshalOptions: options.UnmarshalOptions(),
	}
	transports := &demoHttpClientTransports{
		clientOptions:   options.ClientTransportOptions(),
		middlewares:     options.Middlewares(),
		requestEncoder:  requestEncoder,
		responseDecoder: responseDecoder,
	}
	factories := &demoFactories{
		transports: transports,
	}
	endpointer := &demoEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &demoBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &demoClientEndpoints{
		balancers: balancers,
	}
	return &demoClientService{
		endpoints:     endpoints,
		transportName: httptransportx.HttpClient,
	}
}

type DemoHttpServerTransports interface {
	CreateUser() http.Handler
	DeleteUser() http.Handler
	UpdateUser() http.Handler
	GetUser() http.Handler
	GetUsers() http.Handler
	UploadUserAvatar() http.Handler
	GetUserAvatar() http.Handler
}

type DemoHttpServerRequestDecoder interface {
	CreateUser() http1.DecodeRequestFunc
	DeleteUser() http1.DecodeRequestFunc
	UpdateUser() http1.DecodeRequestFunc
	GetUser() http1.DecodeRequestFunc
	GetUsers() http1.DecodeRequestFunc
	UploadUserAvatar() http1.DecodeRequestFunc
	GetUserAvatar() http1.DecodeRequestFunc
}

type DemoHttpServerResponseEncoder interface {
	CreateUser() http1.EncodeResponseFunc
	DeleteUser() http1.EncodeResponseFunc
	UpdateUser() http1.EncodeResponseFunc
	GetUser() http1.EncodeResponseFunc
	GetUsers() http1.EncodeResponseFunc
	UploadUserAvatar() http1.EncodeResponseFunc
	GetUserAvatar() http1.EncodeResponseFunc
}

type DemoHttpClientRequestEncoder interface {
	CreateUser(instance string) http1.CreateRequestFunc
	DeleteUser(instance string) http1.CreateRequestFunc
	UpdateUser(instance string) http1.CreateRequestFunc
	GetUser(instance string) http1.CreateRequestFunc
	GetUsers(instance string) http1.CreateRequestFunc
	UploadUserAvatar(instance string) http1.CreateRequestFunc
	GetUserAvatar(instance string) http1.CreateRequestFunc
}

type DemoHttpClientResponseDecoder interface {
	CreateUser() http1.DecodeResponseFunc
	DeleteUser() http1.DecodeResponseFunc
	UpdateUser() http1.DecodeResponseFunc
	GetUser() http1.DecodeResponseFunc
	GetUsers() http1.DecodeResponseFunc
	UploadUserAvatar() http1.DecodeResponseFunc
	GetUserAvatar() http1.DecodeResponseFunc
}

type demoHttpServerTransports struct {
	endpoints       DemoServerEndpoints
	requestDecoder  DemoHttpServerRequestDecoder
	responseEncoder DemoHttpServerResponseEncoder
}

func (t *demoHttpServerTransports) CreateUser() http.Handler {
	return http1.NewServer(
		t.endpoints.CreateUser(context.TODO()),
		t.requestDecoder.CreateUser(),
		t.responseEncoder.CreateUser(),
		http1.ServerBefore(httptransportx.EndpointInjector("/leo.example.demo.v1.Demo/CreateUser")),
		http1.ServerBefore(httptransportx.ServerTransportInjector),
		http1.ServerBefore(metadatax.HttpIncomingInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(stainx.HttpIncomingInjector),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
	)
}

func (t *demoHttpServerTransports) DeleteUser() http.Handler {
	return http1.NewServer(
		t.endpoints.DeleteUser(context.TODO()),
		t.requestDecoder.DeleteUser(),
		t.responseEncoder.DeleteUser(),
		http1.ServerBefore(httptransportx.EndpointInjector("/leo.example.demo.v1.Demo/DeleteUser")),
		http1.ServerBefore(httptransportx.ServerTransportInjector),
		http1.ServerBefore(metadatax.HttpIncomingInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(stainx.HttpIncomingInjector),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
	)
}

func (t *demoHttpServerTransports) UpdateUser() http.Handler {
	return http1.NewServer(
		t.endpoints.UpdateUser(context.TODO()),
		t.requestDecoder.UpdateUser(),
		t.responseEncoder.UpdateUser(),
		http1.ServerBefore(httptransportx.EndpointInjector("/leo.example.demo.v1.Demo/UpdateUser")),
		http1.ServerBefore(httptransportx.ServerTransportInjector),
		http1.ServerBefore(metadatax.HttpIncomingInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(stainx.HttpIncomingInjector),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
	)
}

func (t *demoHttpServerTransports) GetUser() http.Handler {
	return http1.NewServer(
		t.endpoints.GetUser(context.TODO()),
		t.requestDecoder.GetUser(),
		t.responseEncoder.GetUser(),
		http1.ServerBefore(httptransportx.EndpointInjector("/leo.example.demo.v1.Demo/GetUser")),
		http1.ServerBefore(httptransportx.ServerTransportInjector),
		http1.ServerBefore(metadatax.HttpIncomingInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(stainx.HttpIncomingInjector),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
	)
}

func (t *demoHttpServerTransports) GetUsers() http.Handler {
	return http1.NewServer(
		t.endpoints.GetUsers(context.TODO()),
		t.requestDecoder.GetUsers(),
		t.responseEncoder.GetUsers(),
		http1.ServerBefore(httptransportx.EndpointInjector("/leo.example.demo.v1.Demo/GetUsers")),
		http1.ServerBefore(httptransportx.ServerTransportInjector),
		http1.ServerBefore(metadatax.HttpIncomingInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(stainx.HttpIncomingInjector),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
	)
}

func (t *demoHttpServerTransports) UploadUserAvatar() http.Handler {
	return http1.NewServer(
		t.endpoints.UploadUserAvatar(context.TODO()),
		t.requestDecoder.UploadUserAvatar(),
		t.responseEncoder.UploadUserAvatar(),
		http1.ServerBefore(httptransportx.EndpointInjector("/leo.example.demo.v1.Demo/UploadUserAvatar")),
		http1.ServerBefore(httptransportx.ServerTransportInjector),
		http1.ServerBefore(metadatax.HttpIncomingInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(stainx.HttpIncomingInjector),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
	)
}

func (t *demoHttpServerTransports) GetUserAvatar() http.Handler {
	return http1.NewServer(
		t.endpoints.GetUserAvatar(context.TODO()),
		t.requestDecoder.GetUserAvatar(),
		t.responseEncoder.GetUserAvatar(),
		http1.ServerBefore(httptransportx.EndpointInjector("/leo.example.demo.v1.Demo/GetUserAvatar")),
		http1.ServerBefore(httptransportx.ServerTransportInjector),
		http1.ServerBefore(metadatax.HttpIncomingInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(stainx.HttpIncomingInjector),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
	)
}

type demoHttpServerRequestDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder demoHttpServerRequestDecoder) CreateUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &CreateUserRequest{}
		if err := coder.DecodeMessageFromRequest(ctx, r, req, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return req, nil
	}
}
func (decoder demoHttpServerRequestDecoder) DeleteUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &DeleteUsersRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.UserId, varErr = coder.DecodeForm[uint64](varErr, vars, "user_id", urlx.GetUint)
		if varErr != nil {
			return nil, varErr
		}
		return req, nil
	}
}
func (decoder demoHttpServerRequestDecoder) UpdateUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &UpdateUserRequest{}
		req.User = &User{}
		if err := coder.DecodeMessageFromRequest(ctx, r, req.GetUser(), decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.UserId, varErr = coder.DecodeForm[uint64](varErr, vars, "user_id", urlx.GetUint)
		if varErr != nil {
			return nil, varErr
		}
		return req, nil
	}
}
func (decoder demoHttpServerRequestDecoder) GetUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &GetUserRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.UserId, varErr = coder.DecodeForm[uint64](varErr, vars, "user_id", urlx.GetUint)
		if varErr != nil {
			return nil, varErr
		}
		return req, nil
	}
}
func (decoder demoHttpServerRequestDecoder) GetUsers() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &GetUsersRequest{}
		queries := r.URL.Query()
		var queryErr error
		req.PageNo, queryErr = coder.DecodeForm[int32](queryErr, queries, "page_no", urlx.GetInt[int32])
		req.PageSize, queryErr = coder.DecodeForm[int32](queryErr, queries, "page_size", urlx.GetInt[int32])
		if queryErr != nil {
			return nil, queryErr
		}
		return req, nil
	}
}
func (decoder demoHttpServerRequestDecoder) UploadUserAvatar() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &UploadUserAvatarRequest{}
		req.Avatar = &httpbody.HttpBody{}
		if err := coder.DecodeHttpBodyFromRequest(ctx, r, req.GetAvatar()); err != nil {
			return nil, err
		}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.UserId, varErr = coder.DecodeForm[uint64](varErr, vars, "user_id", urlx.GetUint)
		if varErr != nil {
			return nil, varErr
		}
		return req, nil
	}
}
func (decoder demoHttpServerRequestDecoder) GetUserAvatar() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &GetUserAvatarRequest{}
		vars := urlx.FormFromMap(mux.Vars(r))
		var varErr error
		req.UserId, varErr = coder.DecodeForm[uint64](varErr, vars, "user_id", urlx.GetUint)
		if varErr != nil {
			return nil, varErr
		}
		return req, nil
	}
}

type demoHttpServerResponseEncoder struct {
	marshalOptions protojson.MarshalOptions
}

func (encoder demoHttpServerResponseEncoder) CreateUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*CreateUserResponse)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder demoHttpServerResponseEncoder) DeleteUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*emptypb.Empty)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder demoHttpServerResponseEncoder) UpdateUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*emptypb.Empty)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder demoHttpServerResponseEncoder) GetUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*GetUserResponse)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder demoHttpServerResponseEncoder) GetUsers() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*GetUsersResponse)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder demoHttpServerResponseEncoder) UploadUserAvatar() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*emptypb.Empty)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder demoHttpServerResponseEncoder) GetUserAvatar() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*httpbody.HttpBody)
		return coder.EncodeHttpBodyToResponse(ctx, w, resp)
	}
}

type demoHttpClientTransports struct {
	clientOptions   []http1.ClientOption
	middlewares     []endpoint.Middleware
	requestEncoder  DemoHttpClientRequestEncoder
	responseDecoder DemoHttpClientResponseDecoder
}

func (t *demoHttpClientTransports) CreateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(metadatax.HttpOutgoingInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(stainx.HttpOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.CreateUser(instance),
		t.responseDecoder.CreateUser(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *demoHttpClientTransports) DeleteUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(metadatax.HttpOutgoingInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(stainx.HttpOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.DeleteUser(instance),
		t.responseDecoder.DeleteUser(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *demoHttpClientTransports) UpdateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(metadatax.HttpOutgoingInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(stainx.HttpOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.UpdateUser(instance),
		t.responseDecoder.UpdateUser(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *demoHttpClientTransports) GetUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(metadatax.HttpOutgoingInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(stainx.HttpOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.GetUser(instance),
		t.responseDecoder.GetUser(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *demoHttpClientTransports) GetUsers(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(metadatax.HttpOutgoingInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(stainx.HttpOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.GetUsers(instance),
		t.responseDecoder.GetUsers(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *demoHttpClientTransports) UploadUserAvatar(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(metadatax.HttpOutgoingInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(stainx.HttpOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.UploadUserAvatar(instance),
		t.responseDecoder.UploadUserAvatar(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *demoHttpClientTransports) GetUserAvatar(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(metadatax.HttpOutgoingInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(stainx.HttpOutgoingInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.GetUserAvatar(instance),
		t.responseDecoder.GetUserAvatar(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

type demoHttpClientRequestEncoder struct {
	marshalOptions protojson.MarshalOptions
	router         *mux.Router
	scheme         string
}

func (encoder demoHttpClientRequestEncoder) CreateUser(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*CreateUserRequest)
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
		if err := coder.EncodeMessageToRequest(ctx, req, header, &body, encoder.marshalOptions); err != nil {
			return nil, err
		}
		var pairs []string
		path, err := encoder.router.Get("/leo.example.demo.v1.Demo/CreateUser").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder demoHttpClientRequestEncoder) DeleteUser(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*DeleteUsersRequest)
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
		pairs = append(pairs,
			"user_id", strconvx.FormatUint(req.GetUserId(), 10),
		)
		path, err := encoder.router.Get("/leo.example.demo.v1.Demo/DeleteUser").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder demoHttpClientRequestEncoder) UpdateUser(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*UpdateUserRequest)
		if !ok {
			return nil, fmt.Errorf("invalid request type, %T", obj)
		}
		_ = req
		method := http.MethodPut
		target := &url.URL{
			Scheme: encoder.scheme,
			Host:   instance,
		}
		header := http.Header{}
		var body bytes.Buffer
		if err := coder.EncodeMessageToRequest(ctx, req.GetUser(), header, &body, encoder.marshalOptions); err != nil {
			return nil, err
		}
		var pairs []string
		pairs = append(pairs,
			"user_id", strconvx.FormatUint(req.GetUserId(), 10),
		)
		path, err := encoder.router.Get("/leo.example.demo.v1.Demo/UpdateUser").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder demoHttpClientRequestEncoder) GetUser(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*GetUserRequest)
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
		pairs = append(pairs,
			"user_id", strconvx.FormatUint(req.GetUserId(), 10),
		)
		path, err := encoder.router.Get("/leo.example.demo.v1.Demo/GetUser").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder demoHttpClientRequestEncoder) GetUsers(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*GetUsersRequest)
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
		path, err := encoder.router.Get("/leo.example.demo.v1.Demo/GetUsers").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		queries := url.Values{}
		queries["page_no"] = append(queries["page_no"], strconvx.FormatInt(req.GetPageNo(), 10))
		queries["page_size"] = append(queries["page_size"], strconvx.FormatInt(req.GetPageSize(), 10))
		target.RawQuery = queries.Encode()
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder demoHttpClientRequestEncoder) UploadUserAvatar(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*UploadUserAvatarRequest)
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
		if err := coder.EncodeHttpBodyToRequest(ctx, req.GetAvatar(), header, &body); err != nil {
			return nil, err
		}
		var pairs []string
		pairs = append(pairs,
			"user_id", strconvx.FormatUint(req.GetUserId(), 10),
		)
		path, err := encoder.router.Get("/leo.example.demo.v1.Demo/UploadUserAvatar").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx.CopyHeader(r.Header, header)
		return r, nil
	}
}
func (encoder demoHttpClientRequestEncoder) GetUserAvatar(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*GetUserAvatarRequest)
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
		pairs = append(pairs,
			"user_id", strconvx.FormatUint(req.GetUserId(), 10),
		)
		path, err := encoder.router.Get("/leo.example.demo.v1.Demo/GetUserAvatar").URLPath(pairs...)
		if err != nil {
			return nil, err
		}
		target.Path = path.Path
		r, err := http.NewRequestWithContext(ctx, method, target.String(), &body)
		if err != nil {
			return nil, err
		}
		httpx.CopyHeader(r.Header, header)
		return r, nil
	}
}

type demoHttpClientResponseDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder demoHttpClientResponseDecoder) CreateUser() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &CreateUserResponse{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder demoHttpClientResponseDecoder) DeleteUser() http1.DecodeResponseFunc {
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
func (decoder demoHttpClientResponseDecoder) UpdateUser() http1.DecodeResponseFunc {
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
func (decoder demoHttpClientResponseDecoder) GetUser() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &GetUserResponse{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder demoHttpClientResponseDecoder) GetUsers() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &GetUsersResponse{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder demoHttpClientResponseDecoder) UploadUserAvatar() http1.DecodeResponseFunc {
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
func (decoder demoHttpClientResponseDecoder) GetUserAvatar() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &httpbody.HttpBody{}
		if err := coder.DecodeHttpBodyFromResponse(ctx, r, resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
