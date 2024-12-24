// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package response

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	proto "google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	io "io"
	http "net/http"
	url "net/url"
)

// =========================== http router ===========================

func appendResponseHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/leo.example.response.v1.Response/OmittedResponse").Methods("POST").Path("/v1/omitted/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/StarResponse").Methods("POST").Path("/v1/star/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/NamedResponse").Methods("POST").Path("/v1/named/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyResponse").Methods("PUT").Path("/v1/http/body/omitted/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyNamedResponse").Methods("PUT").Path("/v1/http/body/named/response")
	return router
}

// =========================== http server ===========================

type responseHttpServerTransports struct {
	endpoints ResponseServerEndpoints
}

func (t *responseHttpServerTransports) OmittedResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.OmittedResponse(context.TODO()),
		_Response_OmittedResponse_HttpServer_RequestDecoder,
		_Response_OmittedResponse_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/OmittedResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *responseHttpServerTransports) StarResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.StarResponse(context.TODO()),
		_Response_StarResponse_HttpServer_RequestDecoder,
		_Response_StarResponse_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/StarResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *responseHttpServerTransports) NamedResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.NamedResponse(context.TODO()),
		_Response_NamedResponse_HttpServer_RequestDecoder,
		_Response_NamedResponse_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/NamedResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *responseHttpServerTransports) HttpBodyResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.HttpBodyResponse(context.TODO()),
		_Response_HttpBodyResponse_HttpServer_RequestDecoder,
		_Response_HttpBodyResponse_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/HttpBodyResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *responseHttpServerTransports) HttpBodyNamedResponse() http.Handler {
	return http1.NewServer(
		t.endpoints.HttpBodyNamedResponse(context.TODO()),
		_Response_HttpBodyNamedResponse_HttpServer_RequestDecoder,
		_Response_HttpBodyNamedResponse_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/HttpBodyNamedResponse")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func AppendResponseHttpRoutes(router *mux.Router, svc ResponseService, middlewares ...endpoint.Middleware) *mux.Router {
	endpoints := newResponseServerEndpoints(svc, middlewares...)
	transports := &responseHttpServerTransports{endpoints: endpoints}
	router = appendResponseHttpRoutes(router)
	router.Get("/leo.example.response.v1.Response/OmittedResponse").Handler(transports.OmittedResponse())
	router.Get("/leo.example.response.v1.Response/StarResponse").Handler(transports.StarResponse())
	router.Get("/leo.example.response.v1.Response/NamedResponse").Handler(transports.NamedResponse())
	router.Get("/leo.example.response.v1.Response/HttpBodyResponse").Handler(transports.HttpBodyResponse())
	router.Get("/leo.example.response.v1.Response/HttpBodyNamedResponse").Handler(transports.HttpBodyNamedResponse())
	return router
}

// =========================== http client ===========================

type responseHttpClientTransports struct {
	scheme        string
	router        *mux.Router
	clientOptions []http1.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *responseHttpClientTransports) OmittedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_Response_OmittedResponse_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_Response_OmittedResponse_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) StarResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_Response_StarResponse_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_Response_StarResponse_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) NamedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_Response_NamedResponse_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_Response_NamedResponse_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) HttpBodyResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_Response_HttpBodyResponse_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_Response_HttpBodyResponse_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *responseHttpClientTransports) HttpBodyNamedResponse(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_Response_HttpBodyNamedResponse_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_Response_HttpBodyNamedResponse_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newResponseHttpClientTransports(scheme string, clientOptions []http1.ClientOption, middlewares []endpoint.Middleware) ResponseClientTransports {
	return &responseHttpClientTransports{
		scheme:        scheme,
		router:        appendResponseHttpRoutes(mux.NewRouter()),
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

func NewResponseHttpClient(target string, opts ...httpx.ClientOption) ResponseService {
	options := httpx.NewClientOptions(opts...)
	transports := newResponseHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newResponseClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newResponseClientService(endpoints, httpx.HttpClient)
}

// =========================== http coder ===========================

func _Response_OmittedResponse_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &emptypb.Empty{}
	return req, nil
}

func _Response_OmittedResponse_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
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
			path, err := router.Get("/leo.example.response.v1.Response/OmittedResponse").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			target := &url.URL{
				Scheme:   scheme,
				Host:     instance,
				Path:     path.Path,
				RawQuery: queries.Encode(),
			}
			r, err := http.NewRequestWithContext(ctx, "POST", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _Response_OmittedResponse_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*UserResponse)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _Response_OmittedResponse_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &UserResponse{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _Response_StarResponse_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &emptypb.Empty{}
	return req, nil
}

func _Response_StarResponse_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
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
			path, err := router.Get("/leo.example.response.v1.Response/StarResponse").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			target := &url.URL{
				Scheme:   scheme,
				Host:     instance,
				Path:     path.Path,
				RawQuery: queries.Encode(),
			}
			r, err := http.NewRequestWithContext(ctx, "POST", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _Response_StarResponse_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*UserResponse)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _Response_StarResponse_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &UserResponse{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _Response_NamedResponse_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &emptypb.Empty{}
	return req, nil
}

func _Response_NamedResponse_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
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
			path, err := router.Get("/leo.example.response.v1.Response/NamedResponse").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			target := &url.URL{
				Scheme:   scheme,
				Host:     instance,
				Path:     path.Path,
				RawQuery: queries.Encode(),
			}
			r, err := http.NewRequestWithContext(ctx, "POST", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _Response_NamedResponse_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*UserResponse)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp.GetUser()); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _Response_NamedResponse_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &UserResponse{}
	if err := jsonx.NewDecoder(r.Body).Decode(&resp.User); err != nil {
		return nil, err
	}
	return resp, nil
}

func _Response_HttpBodyResponse_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &emptypb.Empty{}
	return req, nil
}

func _Response_HttpBodyResponse_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
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
			path, err := router.Get("/leo.example.response.v1.Response/HttpBodyResponse").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			target := &url.URL{
				Scheme:   scheme,
				Host:     instance,
				Path:     path.Path,
				RawQuery: queries.Encode(),
			}
			r, err := http.NewRequestWithContext(ctx, "PUT", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _Response_HttpBodyResponse_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*httpbody.HttpBody)
	w.Header().Set("Content-Type", resp.GetContentType())
	for _, src := range resp.GetExtensions() {
		dst, err := anypb.UnmarshalNew(src, proto.UnmarshalOptions{})
		if err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		metadata, ok := dst.(*structpb.Struct)
		if !ok {
			continue
		}
		for key, value := range metadata.GetFields() {
			w.Header().Add(key, string(errorx.Ignore(jsonx.Marshal(value))))
		}
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp.GetData()); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _Response_HttpBodyResponse_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &httpbody.HttpBody{}
	resp.ContentType = r.Header.Get("Content-Type")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	resp.Data = body
	return resp, nil
}

func _Response_HttpBodyNamedResponse_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &emptypb.Empty{}
	return req, nil
}

func _Response_HttpBodyNamedResponse_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
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
			path, err := router.Get("/leo.example.response.v1.Response/HttpBodyNamedResponse").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			target := &url.URL{
				Scheme:   scheme,
				Host:     instance,
				Path:     path.Path,
				RawQuery: queries.Encode(),
			}
			r, err := http.NewRequestWithContext(ctx, "PUT", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _Response_HttpBodyNamedResponse_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*HttpBody)
	w.Header().Set("Content-Type", resp.GetBody().GetContentType())
	for _, src := range resp.GetBody().GetExtensions() {
		dst, err := anypb.UnmarshalNew(src, proto.UnmarshalOptions{})
		if err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		metadata, ok := dst.(*structpb.Struct)
		if !ok {
			continue
		}
		for key, value := range metadata.GetFields() {
			w.Header().Add(key, string(errorx.Ignore(jsonx.Marshal(value))))
		}
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp.GetBody().GetData()); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _Response_HttpBodyNamedResponse_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &HttpBody{}
	resp.Body = &httpbody.HttpBody{}
	resp.Body.ContentType = r.Header.Get("Content-Type")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Data = body
	return resp, nil
}
