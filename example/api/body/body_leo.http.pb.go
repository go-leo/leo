// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package body

import (
	bytes "bytes"
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
)

func appendBodyHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/leo.example.body.v1.Body/StarBody").Methods("POST").Path("/v1/star/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/NamedBody").Methods("POST").Path("/v1/named/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/NonBody").Methods("GET").Path("/v1/user_body")
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyStarBody").Methods("PUT").Path("/v1/http/body/star/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyNamedBody").Methods("PUT").Path("/v1/http/body/named/body")
	return router
}
func AppendBodyHttpServerRoutes(router *mux.Router, svc BodyService, middlewares ...endpoint.Middleware) *mux.Router {
	transports := newBodyHttpServerTransports(svc, middlewares...)
	router = appendBodyHttpRoutes(router)
	router.Get("/leo.example.body.v1.Body/StarBody").Handler(transports.StarBody())
	router.Get("/leo.example.body.v1.Body/NamedBody").Handler(transports.NamedBody())
	router.Get("/leo.example.body.v1.Body/NonBody").Handler(transports.NonBody())
	router.Get("/leo.example.body.v1.Body/HttpBodyStarBody").Handler(transports.HttpBodyStarBody())
	router.Get("/leo.example.body.v1.Body/HttpBodyNamedBody").Handler(transports.HttpBodyNamedBody())
	return router
}

func NewBodyHttpClient(target string, opts ...httpx.ClientOption) BodyService {
	options := httpx.NewClientOptions(opts...)
	transports := newBodyHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newBodyClientEndpoints(target, transports, options.Builder(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newBodyClientService(endpoints, httpx.HttpClient)
}

type BodyHttpServerTransports interface {
	StarBody() http.Handler
	NamedBody() http.Handler
	NonBody() http.Handler
	HttpBodyStarBody() http.Handler
	HttpBodyNamedBody() http.Handler
}

type BodyHttpServerRequestDecoder interface {
	StarBody() http1.DecodeRequestFunc
	NamedBody() http1.DecodeRequestFunc
	NonBody() http1.DecodeRequestFunc
	HttpBodyStarBody() http1.DecodeRequestFunc
	HttpBodyNamedBody() http1.DecodeRequestFunc
}

type BodyHttpServerResponseEncoder interface {
	StarBody() http1.EncodeResponseFunc
	NamedBody() http1.EncodeResponseFunc
	NonBody() http1.EncodeResponseFunc
	HttpBodyStarBody() http1.EncodeResponseFunc
	HttpBodyNamedBody() http1.EncodeResponseFunc
}

type BodyHttpClientRequestEncoder interface {
	StarBody(instance string) http1.CreateRequestFunc
	NamedBody(instance string) http1.CreateRequestFunc
	NonBody(instance string) http1.CreateRequestFunc
	HttpBodyStarBody(instance string) http1.CreateRequestFunc
	HttpBodyNamedBody(instance string) http1.CreateRequestFunc
}

type BodyHttpClientResponseDecoder interface {
	StarBody() http1.DecodeResponseFunc
	NamedBody() http1.DecodeResponseFunc
	NonBody() http1.DecodeResponseFunc
	HttpBodyStarBody() http1.DecodeResponseFunc
	HttpBodyNamedBody() http1.DecodeResponseFunc
}

type bodyHttpServerTransports struct {
	endpoints       BodyServerEndpoints
	requestDecoder  BodyHttpServerRequestDecoder
	responseEncoder BodyHttpServerResponseEncoder
}

func (t *bodyHttpServerTransports) StarBody() http.Handler {
	return http1.NewServer(
		t.endpoints.StarBody(context.TODO()),
		t.requestDecoder.StarBody(),
		t.responseEncoder.StarBody(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/StarBody")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *bodyHttpServerTransports) NamedBody() http.Handler {
	return http1.NewServer(
		t.endpoints.NamedBody(context.TODO()),
		t.requestDecoder.NamedBody(),
		t.responseEncoder.NamedBody(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/NamedBody")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *bodyHttpServerTransports) NonBody() http.Handler {
	return http1.NewServer(
		t.endpoints.NonBody(context.TODO()),
		t.requestDecoder.NonBody(),
		t.responseEncoder.NonBody(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/NonBody")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *bodyHttpServerTransports) HttpBodyStarBody() http.Handler {
	return http1.NewServer(
		t.endpoints.HttpBodyStarBody(context.TODO()),
		t.requestDecoder.HttpBodyStarBody(),
		t.responseEncoder.HttpBodyStarBody(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/HttpBodyStarBody")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *bodyHttpServerTransports) HttpBodyNamedBody() http.Handler {
	return http1.NewServer(
		t.endpoints.HttpBodyNamedBody(context.TODO()),
		t.requestDecoder.HttpBodyNamedBody(),
		t.responseEncoder.HttpBodyNamedBody(),
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/HttpBodyNamedBody")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func newBodyHttpServerTransports(svc BodyService, middlewares ...endpoint.Middleware) BodyHttpServerTransports {
	endpoints := newBodyServerEndpoints(svc, middlewares...)
	return &bodyHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  bodyHttpServerRequestDecoder{},
		responseEncoder: bodyHttpServerResponseEncoder{},
	}
}

type bodyHttpServerRequestDecoder struct{}

func (bodyHttpServerRequestDecoder) StarBody() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &User{}
		if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return req, nil
	}
}
func (bodyHttpServerRequestDecoder) NamedBody() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &UserRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(&req.User); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return req, nil
	}
}
func (bodyHttpServerRequestDecoder) NonBody() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &emptypb.Empty{}
		return req, nil
	}
}
func (bodyHttpServerRequestDecoder) HttpBodyStarBody() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &httpbody.HttpBody{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		req.Data = body
		req.ContentType = r.Header.Get("Content-Type")
		return req, nil
	}
}
func (bodyHttpServerRequestDecoder) HttpBodyNamedBody() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &HttpBody{}
		req.Body = &httpbody.HttpBody{}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		req.Body.Data = body
		req.Body.ContentType = r.Header.Get("Content-Type")
		return req, nil
	}
}

type bodyHttpServerResponseEncoder struct{}

func (bodyHttpServerResponseEncoder) StarBody() http1.EncodeResponseFunc {
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
func (bodyHttpServerResponseEncoder) NamedBody() http1.EncodeResponseFunc {
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
func (bodyHttpServerResponseEncoder) NonBody() http1.EncodeResponseFunc {
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
func (bodyHttpServerResponseEncoder) HttpBodyStarBody() http1.EncodeResponseFunc {
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
func (bodyHttpServerResponseEncoder) HttpBodyNamedBody() http1.EncodeResponseFunc {
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

type bodyHttpClientTransports struct {
	clientOptions   []http1.ClientOption
	middlewares     []endpoint.Middleware
	requestEncoder  BodyHttpClientRequestEncoder
	responseDecoder BodyHttpClientResponseDecoder
}

func (t *bodyHttpClientTransports) StarBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.StarBody(instance),
		t.responseDecoder.StarBody(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *bodyHttpClientTransports) NamedBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.NamedBody(instance),
		t.responseDecoder.NamedBody(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *bodyHttpClientTransports) NonBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.NonBody(instance),
		t.responseDecoder.NonBody(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *bodyHttpClientTransports) HttpBodyStarBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.HttpBodyStarBody(instance),
		t.responseDecoder.HttpBodyStarBody(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *bodyHttpClientTransports) HttpBodyNamedBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.HttpBodyNamedBody(instance),
		t.responseDecoder.HttpBodyNamedBody(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newBodyHttpClientTransports(scheme string, clientOptions []http1.ClientOption, middlewares []endpoint.Middleware) BodyClientTransports {
	return &bodyHttpClientTransports{
		clientOptions: clientOptions,
		middlewares:   middlewares,
		requestEncoder: bodyHttpClientRequestEncoder{
			scheme: scheme,
			router: appendBodyHttpRoutes(mux.NewRouter()),
		},
		responseDecoder: bodyHttpClientResponseDecoder{},
	}
}

type bodyHttpClientRequestEncoder struct {
	router *mux.Router
	scheme string
}

func (e bodyHttpClientRequestEncoder) StarBody(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*User)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var bodyBuf bytes.Buffer
		if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		body = &bodyBuf
		contentType := "application/json; charset=utf-8"
		var pairs []string
		path, err := e.router.Get("/leo.example.body.v1.Body/StarBody").URLPath(pairs...)
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
func (e bodyHttpClientRequestEncoder) NamedBody(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*UserRequest)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var bodyBuf bytes.Buffer
		if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetUser()); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		body = &bodyBuf
		contentType := "application/json; charset=utf-8"
		var pairs []string
		path, err := e.router.Get("/leo.example.body.v1.Body/NamedBody").URLPath(pairs...)
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
func (e bodyHttpClientRequestEncoder) NonBody(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*emptypb.Empty)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		var pairs []string
		path, err := e.router.Get("/leo.example.body.v1.Body/NonBody").URLPath(pairs...)
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
func (e bodyHttpClientRequestEncoder) HttpBodyStarBody(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*httpbody.HttpBody)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		body = bytes.NewReader(req.GetData())
		contentType := req.GetContentType()
		var pairs []string
		path, err := e.router.Get("/leo.example.body.v1.Body/HttpBodyStarBody").URLPath(pairs...)
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
		r, err := http.NewRequestWithContext(ctx, "PUT", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		r.Header.Set("Content-Type", contentType)
		return r, nil
	}
}
func (e bodyHttpClientRequestEncoder) HttpBodyNamedBody(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
		}
		req, ok := obj.(*HttpBody)
		if !ok {
			return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
		}
		_ = req
		var body io.Reader
		body = bytes.NewReader(req.GetBody().GetData())
		contentType := req.GetBody().GetContentType()
		var pairs []string
		path, err := e.router.Get("/leo.example.body.v1.Body/HttpBodyNamedBody").URLPath(pairs...)
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
		r, err := http.NewRequestWithContext(ctx, "PUT", target.String(), body)
		if err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		r.Header.Set("Content-Type", contentType)
		return r, nil
	}
}

type bodyHttpClientResponseDecoder struct{}

func (bodyHttpClientResponseDecoder) StarBody() http1.DecodeResponseFunc {
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
func (bodyHttpClientResponseDecoder) NamedBody() http1.DecodeResponseFunc {
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
func (bodyHttpClientResponseDecoder) NonBody() http1.DecodeResponseFunc {
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
func (bodyHttpClientResponseDecoder) HttpBodyStarBody() http1.DecodeResponseFunc {
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
func (bodyHttpClientResponseDecoder) HttpBodyNamedBody() http1.DecodeResponseFunc {
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
