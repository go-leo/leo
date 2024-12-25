// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package cqrs

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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http "net/http"
	url "net/url"
)

// =========================== http router ===========================

func appendCQRSHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/pb.CQRS/CreateUser").Methods("POST").Path("/pb.CQRS/CreateUser")
	router.NewRoute().Name("/pb.CQRS/FindUser").Methods("POST").Path("/pb.CQRS/FindUser")
	return router
}
func AppendCQRSHttpRoutes(router *mux.Router, svc CQRSService, middlewares ...endpoint.Middleware) *mux.Router {
	transports := newCQRSHttpServerTransports(svc, middlewares...)
	router = appendCQRSHttpRoutes(router)
	router.Get("/pb.CQRS/CreateUser").Handler(transports.CreateUser())
	router.Get("/pb.CQRS/FindUser").Handler(transports.FindUser())
	return router
}

func NewCQRSHttpClient(target string, opts ...httpx.ClientOption) CQRSService {
	options := httpx.NewClientOptions(opts...)
	transports := newCQRSHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newCQRSClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newCQRSClientService(endpoints, httpx.HttpClient)
}

// =========================== http server ===========================

type CQRSHttpServerTransports interface {
	CreateUser() http.Handler
	FindUser() http.Handler
}

type CQRSHttpServerRequestDecoder interface {
	CreateUser() http1.DecodeRequestFunc
	FindUser() http1.DecodeRequestFunc
}

type CQRSHttpServerResponseEncoder interface {
	CreateUser() http1.EncodeResponseFunc
	FindUser() http1.EncodeResponseFunc
}

type cQRSHttpServerTransports struct {
	endpoints       CQRSServerEndpoints
	requestDecoder  CQRSHttpServerRequestDecoder
	responseEncoder CQRSHttpServerResponseEncoder
}

func (t *cQRSHttpServerTransports) CreateUser() http.Handler {
	return http1.NewServer(
		t.endpoints.CreateUser(context.TODO()),
		t.requestDecoder.CreateUser(),
		t.responseEncoder.CreateUser(),
		http1.ServerBefore(httpx.EndpointInjector("/pb.CQRS/CreateUser")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *cQRSHttpServerTransports) FindUser() http.Handler {
	return http1.NewServer(
		t.endpoints.FindUser(context.TODO()),
		t.requestDecoder.FindUser(),
		t.responseEncoder.FindUser(),
		http1.ServerBefore(httpx.EndpointInjector("/pb.CQRS/FindUser")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func newCQRSHttpServerTransports(svc CQRSService, middlewares ...endpoint.Middleware) CQRSHttpServerTransports {
	endpoints := newCQRSServerEndpoints(svc, middlewares...)
	return &cQRSHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  cQRSHttpServerRequestDecoder{},
		responseEncoder: cQRSHttpServerResponseEncoder{},
	}
}

type cQRSHttpServerRequestDecoder struct{}

func (cQRSHttpServerRequestDecoder) CreateUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &CreateUserRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return req, nil
	}
}
func (cQRSHttpServerRequestDecoder) FindUser() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &FindUserRequest{}
		if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
		}
		return req, nil
	}
}

type cQRSHttpServerResponseEncoder struct{}

func (cQRSHttpServerResponseEncoder) CreateUser() http1.EncodeResponseFunc {
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
func (cQRSHttpServerResponseEncoder) FindUser() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*GetUserResponse)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
			return statusx.ErrInternal.With(statusx.Wrap(err))
		}
		return nil
	}

}

// =========================== http client ===========================

type cQRSHttpClientTransports struct {
	scheme        string
	router        *mux.Router
	clientOptions []http1.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *cQRSHttpClientTransports) CreateUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_CQRS_CreateUser_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_CQRS_CreateUser_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *cQRSHttpClientTransports) FindUser(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_CQRS_FindUser_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_CQRS_FindUser_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newCQRSHttpClientTransports(scheme string, clientOptions []http1.ClientOption, middlewares []endpoint.Middleware) CQRSClientTransports {
	return &cQRSHttpClientTransports{
		scheme:        scheme,
		router:        appendCQRSHttpRoutes(mux.NewRouter()),
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

// =========================== http coder ===========================

func _CQRS_CreateUser_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*CreateUserRequest)
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
			path, err := router.Get("/pb.CQRS/CreateUser").URLPath(pairs...)
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
			r.Header.Set("Content-Type", contentType)
			return r, nil
		}
	}
}

func _CQRS_CreateUser_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _CQRS_FindUser_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*FindUserRequest)
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
			path, err := router.Get("/pb.CQRS/FindUser").URLPath(pairs...)
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
			r.Header.Set("Content-Type", contentType)
			return r, nil
		}
	}
}

func _CQRS_FindUser_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &GetUserResponse{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
