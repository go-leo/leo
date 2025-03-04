// Code generated by protoc-gen-go-leo. DO NOT EDIT.

package api

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	httpx1 "github.com/go-leo/gox/netx/httpx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	timeoutx "github.com/go-leo/leo/v3/timeoutx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	coder "github.com/go-leo/leo/v3/transportx/httpx/coder"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	http2 "google.golang.org/genproto/googleapis/rpc/http"
	protojson "google.golang.org/protobuf/encoding/protojson"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
)

func appendResponseHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().
		Name("/leo.example.route.response.Response/OmittedResponse").
		Methods(http.MethodPost).
		Path("/v1/omitted/response")
	router.NewRoute().
		Name("/leo.example.route.response.Response/StarResponse").
		Methods(http.MethodPost).
		Path("/v1/star/response")
	router.NewRoute().
		Name("/leo.example.route.response.Response/NamedResponse").
		Methods(http.MethodPost).
		Path("/v1/named/response")
	router.NewRoute().
		Name("/leo.example.route.response.Response/HttpBodyResponse").
		Methods(http.MethodPut).
		Path("/v1/http/body/omitted/response")
	router.NewRoute().
		Name("/leo.example.route.response.Response/HttpBodyNamedResponse").
		Methods(http.MethodPut).
		Path("/v1/http/body/named/response")
	router.NewRoute().
		Name("/leo.example.route.response.Response/HttpResponse").
		Methods(http.MethodGet).
		Path("/v1/http/response")
	return router
}
func AppendResponseHttpServerRoutes(router *mux.Router, svc ResponseService, opts ...httpx.ServerOption) *mux.Router {
	options := httpx.NewServerOptions(opts...)
	endpoints := &responseServerEndpoints{
		svc:         svc,
		middlewares: options.Middlewares(),
	}
	requestDecoder := responseHttpServerRequestDecoder{
		unmarshalOptions: options.UnmarshalOptions(),
	}
	responseEncoder := responseHttpServerResponseEncoder{
		marshalOptions: options.MarshalOptions(),
	}
	transports := &responseHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  requestDecoder,
		responseEncoder: responseEncoder,
	}
	router = appendResponseHttpRoutes(router)
	router.Get("/leo.example.route.response.Response/OmittedResponse").Handler(transports.OmittedResponse())
	router.Get("/leo.example.route.response.Response/StarResponse").Handler(transports.StarResponse())
	router.Get("/leo.example.route.response.Response/NamedResponse").Handler(transports.NamedResponse())
	router.Get("/leo.example.route.response.Response/HttpBodyResponse").Handler(transports.HttpBodyResponse())
	router.Get("/leo.example.route.response.Response/HttpBodyNamedResponse").Handler(transports.HttpBodyNamedResponse())
	router.Get("/leo.example.route.response.Response/HttpResponse").Handler(transports.HttpResponse())
	return router
}

func NewResponseHttpClient(target string, opts ...httpx.ClientOption) ResponseService {
	options := httpx.NewClientOptions(opts...)
	requestEncoder := &responseHttpClientRequestEncoder{
		marshalOptions: options.MarshalOptions(),
		router:         appendResponseHttpRoutes(mux.NewRouter()),
		scheme:         options.Scheme(),
	}
	responseDecoder := &responseHttpClientResponseDecoder{
		unmarshalOptions: options.UnmarshalOptions(),
	}
	transports := &responseHttpClientTransports{
		clientOptions:   options.ClientTransportOptions(),
		middlewares:     options.Middlewares(),
		requestEncoder:  requestEncoder,
		responseDecoder: responseDecoder,
	}
	factories := &responseFactories{
		transports: transports,
	}
	endpointer := &responseEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &responseBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &responseClientEndpoints{
		balancers: balancers,
	}
	return &responseClientService{
		endpoints:     endpoints,
		transportName: httpx.HttpClient,
	}
}

type ResponseHttpServerTransports interface {
	OmittedResponse() http.Handler
	StarResponse() http.Handler
	NamedResponse() http.Handler
	HttpBodyResponse() http.Handler
	HttpBodyNamedResponse() http.Handler
	HttpResponse() http.Handler
}

type ResponseHttpServerRequestDecoder interface {
	OmittedResponse() http1.DecodeRequestFunc
	StarResponse() http1.DecodeRequestFunc
	NamedResponse() http1.DecodeRequestFunc
	HttpBodyResponse() http1.DecodeRequestFunc
	HttpBodyNamedResponse() http1.DecodeRequestFunc
	HttpResponse() http1.DecodeRequestFunc
}

type ResponseHttpServerResponseEncoder interface {
	OmittedResponse() http1.EncodeResponseFunc
	StarResponse() http1.EncodeResponseFunc
	NamedResponse() http1.EncodeResponseFunc
	HttpBodyResponse() http1.EncodeResponseFunc
	HttpBodyNamedResponse() http1.EncodeResponseFunc
	HttpResponse() http1.EncodeResponseFunc
}

type ResponseHttpClientRequestEncoder interface {
	OmittedResponse(instance string) http1.CreateRequestFunc
	StarResponse(instance string) http1.CreateRequestFunc
	NamedResponse(instance string) http1.CreateRequestFunc
	HttpBodyResponse(instance string) http1.CreateRequestFunc
	HttpBodyNamedResponse(instance string) http1.CreateRequestFunc
	HttpResponse(instance string) http1.CreateRequestFunc
}

type ResponseHttpClientResponseDecoder interface {
	OmittedResponse() http1.DecodeResponseFunc
	StarResponse() http1.DecodeResponseFunc
	NamedResponse() http1.DecodeResponseFunc
	HttpBodyResponse() http1.DecodeResponseFunc
	HttpBodyNamedResponse() http1.DecodeResponseFunc
	HttpResponse() http1.DecodeResponseFunc
}

type responseHttpServerTransports struct {
	endpoints       ResponseServerEndpoints
	requestDecoder  ResponseHttpServerRequestDecoder
	responseEncoder ResponseHttpServerResponseEncoder
}

func (t *responseHttpServerTransports) OmittedResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.OmittedResponse(context.TODO()),
		t.requestDecoder.OmittedResponse(),
		t.responseEncoder.OmittedResponse(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.route.response.Response/OmittedResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
	)
}

func (t *responseHttpServerTransports) StarResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.StarResponse(context.TODO()),
		t.requestDecoder.StarResponse(),
		t.responseEncoder.StarResponse(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.route.response.Response/StarResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
	)
}

func (t *responseHttpServerTransports) NamedResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.NamedResponse(context.TODO()),
		t.requestDecoder.NamedResponse(),
		t.responseEncoder.NamedResponse(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.route.response.Response/NamedResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
	)
}

func (t *responseHttpServerTransports) HttpBodyResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.HttpBodyResponse(context.TODO()),
		t.requestDecoder.HttpBodyResponse(),
		t.responseEncoder.HttpBodyResponse(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.route.response.Response/HttpBodyResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
	)
}

func (t *responseHttpServerTransports) HttpBodyNamedResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.HttpBodyNamedResponse(context.TODO()),
		t.requestDecoder.HttpBodyNamedResponse(),
		t.responseEncoder.HttpBodyNamedResponse(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.route.response.Response/HttpBodyNamedResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
	)
}

func (t *responseHttpServerTransports) HttpResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.HttpResponse(context.TODO()),
		t.requestDecoder.HttpResponse(),
		t.responseEncoder.HttpResponse(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.route.response.Response/HttpResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(timeoutx.IncomingInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerErrorEncoder(coder.EncodeErrorToResponse),
		http1.ServerFinalizer(timeoutx.CancelInvoker),
	)
}

type responseHttpServerRequestDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder responseHttpServerRequestDecoder) OmittedResponse() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &emptypb.Empty{}
		return req, nil
	}
}
func (decoder responseHttpServerRequestDecoder) StarResponse() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &emptypb.Empty{}
		return req, nil
	}
}
func (decoder responseHttpServerRequestDecoder) NamedResponse() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &emptypb.Empty{}
		return req, nil
	}
}
func (decoder responseHttpServerRequestDecoder) HttpBodyResponse() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &emptypb.Empty{}
		return req, nil
	}
}
func (decoder responseHttpServerRequestDecoder) HttpBodyNamedResponse() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &emptypb.Empty{}
		return req, nil
	}
}
func (decoder responseHttpServerRequestDecoder) HttpResponse() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &emptypb.Empty{}
		return req, nil
	}
}

type responseHttpServerResponseEncoder struct {
	marshalOptions protojson.MarshalOptions
}

func (encoder responseHttpServerResponseEncoder) OmittedResponse() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*UserResponse)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder responseHttpServerResponseEncoder) StarResponse() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*UserResponse)
		return coder.EncodeMessageToResponse(ctx, w, resp, encoder.marshalOptions)
	}
}
func (encoder responseHttpServerResponseEncoder) NamedResponse() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*UserResponse)
		return coder.EncodeMessageToResponse(ctx, w, resp.GetUser(), encoder.marshalOptions)
	}
}
func (encoder responseHttpServerResponseEncoder) HttpBodyResponse() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*httpbody.HttpBody)
		return coder.EncodeHttpBodyToResponse(ctx, w, resp)
	}
}
func (encoder responseHttpServerResponseEncoder) HttpBodyNamedResponse() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*HttpBody)
		return coder.EncodeHttpBodyToResponse(ctx, w, resp.GetBody())
	}
}
func (encoder responseHttpServerResponseEncoder) HttpResponse() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*http2.HttpResponse)
		return coder.EncodeHttpResponseToResponse(ctx, w, resp)
	}
}

type responseHttpClientTransports struct {
	clientOptions   []http1.ClientOption
	middlewares     []endpoint.Middleware
	requestEncoder  ResponseHttpClientRequestEncoder
	responseDecoder ResponseHttpClientResponseDecoder
}

func (t *responseHttpClientTransports) OmittedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.OmittedResponse(instance),
		t.responseDecoder.OmittedResponse(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) StarResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.StarResponse(instance),
		t.responseDecoder.StarResponse(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) NamedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.NamedResponse(instance),
		t.responseDecoder.NamedResponse(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) HttpBodyResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.HttpBodyResponse(instance),
		t.responseDecoder.HttpBodyResponse(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) HttpBodyNamedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.HttpBodyNamedResponse(instance),
		t.responseDecoder.HttpBodyNamedResponse(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) HttpResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(timeoutx.OutgoingInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.HttpResponse(instance),
		t.responseDecoder.HttpResponse(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

type responseHttpClientRequestEncoder struct {
	marshalOptions protojson.MarshalOptions
	router         *mux.Router
	scheme         string
}

func (encoder responseHttpClientRequestEncoder) OmittedResponse(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*emptypb.Empty)
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
		var pairs []string
		path, err := encoder.router.Get("/leo.example.route.response.Response/OmittedResponse").URLPath(pairs...)
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
func (encoder responseHttpClientRequestEncoder) StarResponse(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*emptypb.Empty)
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
		var pairs []string
		path, err := encoder.router.Get("/leo.example.route.response.Response/StarResponse").URLPath(pairs...)
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
func (encoder responseHttpClientRequestEncoder) NamedResponse(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*emptypb.Empty)
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
		var pairs []string
		path, err := encoder.router.Get("/leo.example.route.response.Response/NamedResponse").URLPath(pairs...)
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
func (encoder responseHttpClientRequestEncoder) HttpBodyResponse(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*emptypb.Empty)
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
		var pairs []string
		path, err := encoder.router.Get("/leo.example.route.response.Response/HttpBodyResponse").URLPath(pairs...)
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
func (encoder responseHttpClientRequestEncoder) HttpBodyNamedResponse(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*emptypb.Empty)
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
		var pairs []string
		path, err := encoder.router.Get("/leo.example.route.response.Response/HttpBodyNamedResponse").URLPath(pairs...)
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
func (encoder responseHttpClientRequestEncoder) HttpResponse(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*emptypb.Empty)
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
		path, err := encoder.router.Get("/leo.example.route.response.Response/HttpResponse").URLPath(pairs...)
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

type responseHttpClientResponseDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder responseHttpClientResponseDecoder) OmittedResponse() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &UserResponse{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder responseHttpClientResponseDecoder) StarResponse() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &UserResponse{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder responseHttpClientResponseDecoder) NamedResponse() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &UserResponse{}
		if resp.User == nil {
			resp.User = &User{}
		}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp.User, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder responseHttpClientResponseDecoder) HttpBodyResponse() http1.DecodeResponseFunc {
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
func (decoder responseHttpClientResponseDecoder) HttpBodyNamedResponse() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &HttpBody{}
		if resp.Body == nil {
			resp.Body = &httpbody.HttpBody{}
		}
		if err := coder.DecodeHttpBodyFromResponse(ctx, r, resp.Body); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
func (decoder responseHttpClientResponseDecoder) HttpResponse() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		if r.StatusCode != http.StatusOK {
			return nil, coder.DecodeErrorFromResponse(ctx, r)
		}
		resp := &http2.HttpResponse{}
		if err := coder.DecodeHttpResponseFromResponse(ctx, r, resp); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
