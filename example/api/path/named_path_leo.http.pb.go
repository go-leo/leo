// Code generated by protoc-gen-leo-http. DO NOT EDIT.

package path

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	urlx "github.com/go-leo/gox/netx/urlx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	io "io"
	http "net/http"
	url "net/url"
	strings "strings"
)

// =========================== http router ===========================

func appendNamedPathHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/NamedPathString").Methods("GET").Path("/v1/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/NamedPathOptString").Methods("GET").Path("/v1/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/NamedPathWrapString").Methods("GET").Path("/v1/wrap_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/EmbedNamedPathString").Methods("GET").Path("/v1/embed/string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").Methods("GET").Path("/v1/embed/opt_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	router.NewRoute().Name("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").Methods("GET").Path("/v1/embed/wrap_string/classes/{class}/shelves/{shelf}/books/{book}/families/{family}")
	return router
}

// =========================== http server ===========================

type namedPathHttpServerTransports struct {
	endpoints NamedPathServerEndpoints
}

func (t *namedPathHttpServerTransports) NamedPathString() http.Handler {
	return http1.NewServer(
		t.endpoints.NamedPathString(context.TODO()),
		_NamedPath_NamedPathString_HttpServer_RequestDecoder,
		_NamedPath_NamedPathString_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/NamedPathString")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *namedPathHttpServerTransports) NamedPathOptString() http.Handler {
	return http1.NewServer(
		t.endpoints.NamedPathOptString(context.TODO()),
		_NamedPath_NamedPathOptString_HttpServer_RequestDecoder,
		_NamedPath_NamedPathOptString_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/NamedPathOptString")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *namedPathHttpServerTransports) NamedPathWrapString() http.Handler {
	return http1.NewServer(
		t.endpoints.NamedPathWrapString(context.TODO()),
		_NamedPath_NamedPathWrapString_HttpServer_RequestDecoder,
		_NamedPath_NamedPathWrapString_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/NamedPathWrapString")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *namedPathHttpServerTransports) EmbedNamedPathString() http.Handler {
	return http1.NewServer(
		t.endpoints.EmbedNamedPathString(context.TODO()),
		_NamedPath_EmbedNamedPathString_HttpServer_RequestDecoder,
		_NamedPath_EmbedNamedPathString_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathString")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *namedPathHttpServerTransports) EmbedNamedPathOptString() http.Handler {
	return http1.NewServer(
		t.endpoints.EmbedNamedPathOptString(context.TODO()),
		_NamedPath_EmbedNamedPathOptString_HttpServer_RequestDecoder,
		_NamedPath_EmbedNamedPathOptString_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func (t *namedPathHttpServerTransports) EmbedNamedPathWrapString() http.Handler {
	return http1.NewServer(
		t.endpoints.EmbedNamedPathWrapString(context.TODO()),
		_NamedPath_EmbedNamedPathWrapString_HttpServer_RequestDecoder,
		_NamedPath_EmbedNamedPathWrapString_HttpServer_ResponseEncoder,
		http1.ServerBefore(httpx.EndpointInjector("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
		http1.ServerErrorEncoder(httpx.ErrorEncoder),
	)
}

func AppendNamedPathHttpRoutes(router *mux.Router, svc NamedPathService, middlewares ...endpoint.Middleware) *mux.Router {
	endpoints := newNamedPathServerEndpoints(svc, middlewares...)
	transports := &namedPathHttpServerTransports{endpoints: endpoints}
	router = appendNamedPathHttpRoutes(router)
	router.Get("/leo.example.path.v1.NamedPath/NamedPathString").Handler(transports.NamedPathString())
	router.Get("/leo.example.path.v1.NamedPath/NamedPathOptString").Handler(transports.NamedPathOptString())
	router.Get("/leo.example.path.v1.NamedPath/NamedPathWrapString").Handler(transports.NamedPathWrapString())
	router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathString").Handler(transports.EmbedNamedPathString())
	router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").Handler(transports.EmbedNamedPathOptString())
	router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").Handler(transports.EmbedNamedPathWrapString())
	return router
}

// =========================== http client ===========================

type namedPathHttpClientTransports struct {
	scheme        string
	router        *mux.Router
	clientOptions []http1.ClientOption
	middlewares   []endpoint.Middleware
}

func (t *namedPathHttpClientTransports) NamedPathString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_NamedPath_NamedPathString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_NamedPathString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) NamedPathOptString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_NamedPath_NamedPathOptString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_NamedPathOptString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) NamedPathWrapString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_NamedPath_NamedPathWrapString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_NamedPathWrapString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) EmbedNamedPathString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_NamedPath_EmbedNamedPathString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_EmbedNamedPathString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) EmbedNamedPathOptString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_NamedPath_EmbedNamedPathOptString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_EmbedNamedPathOptString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func (t *namedPathHttpClientTransports) EmbedNamedPathWrapString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		_NamedPath_EmbedNamedPathWrapString_HttpClient_RequestEncoder(t.router)(t.scheme, instance),
		_NamedPath_EmbedNamedPathWrapString_HttpClient_ResponseDecoder,
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

func newNamedPathHttpClientTransports(scheme string, clientOptions []http1.ClientOption, middlewares []endpoint.Middleware) NamedPathClientTransports {
	return &namedPathHttpClientTransports{
		scheme:        scheme,
		router:        appendNamedPathHttpRoutes(mux.NewRouter()),
		clientOptions: clientOptions,
		middlewares:   middlewares,
	}
}

func NewNamedPathHttpClient(target string, opts ...httpx.ClientOption) NamedPathService {
	options := httpx.NewClientOptions(opts...)
	transports := newNamedPathHttpClientTransports(options.Scheme(), options.ClientTransportOptions(), options.Middlewares())
	endpoints := newNamedPathClientEndpoints(target, transports, options.InstancerFactory(), options.EndpointerOptions(), options.BalancerFactory(), options.Logger())
	return newNamedPathClientService(endpoints, httpx.HttpClient)
}

// =========================== http coder ===========================

func _NamedPath_NamedPathString_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &NamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.String_ = fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family"))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	queries := r.URL.Query()
	var queryErr error
	req.OptString = proto.String(queries.Get("opt_string"))
	req.WrapString = wrapperspb.String(queries.Get("wrap_string"))
	if queryErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(queryErr))
	}
	return req, nil
}

func _NamedPath_NamedPathString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*NamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetString_()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathString").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			queries["opt_string"] = append(queries["opt_string"], req.GetOptString())
			queries["wrap_string"] = append(queries["wrap_string"], req.GetWrapString().GetValue())
			target := &url.URL{
				Scheme:   scheme,
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
}

func _NamedPath_NamedPathString_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_NamedPathString_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_NamedPathOptString_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &NamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.OptString = proto.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	queries := r.URL.Query()
	var queryErr error
	req.String_ = queries.Get("string")
	req.WrapString = wrapperspb.String(queries.Get("wrap_string"))
	if queryErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(queryErr))
	}
	return req, nil
}

func _NamedPath_NamedPathOptString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*NamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetOptString()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathOptString").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			queries["string"] = append(queries["string"], req.GetString_())
			queries["wrap_string"] = append(queries["wrap_string"], req.GetWrapString().GetValue())
			target := &url.URL{
				Scheme:   scheme,
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
}

func _NamedPath_NamedPathOptString_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_NamedPathOptString_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_NamedPathWrapString_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &NamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	req.WrapString = wrapperspb.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	queries := r.URL.Query()
	var queryErr error
	req.String_ = queries.Get("string")
	req.OptString = proto.String(queries.Get("opt_string"))
	if queryErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(queryErr))
	}
	return req, nil
}

func _NamedPath_NamedPathWrapString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*NamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetWrapString().GetValue()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/NamedPathWrapString").URLPath(pairs...)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			queries := url.Values{}
			queries["string"] = append(queries["string"], req.GetString_())
			queries["opt_string"] = append(queries["opt_string"], req.GetOptString())
			target := &url.URL{
				Scheme:   scheme,
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
}

func _NamedPath_NamedPathWrapString_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_NamedPathWrapString_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_EmbedNamedPathString_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &EmbedNamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	if req.Embed == nil {
		req.Embed = &NamedPathRequest{}
	}
	req.Embed.String_ = fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family"))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	return req, nil
}

func _NamedPath_EmbedNamedPathString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*EmbedNamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetEmbed().GetString_()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathString").URLPath(pairs...)
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
			r, err := http.NewRequestWithContext(ctx, "GET", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _NamedPath_EmbedNamedPathString_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_EmbedNamedPathString_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_EmbedNamedPathOptString_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &EmbedNamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	if req.Embed == nil {
		req.Embed = &NamedPathRequest{}
	}
	req.Embed.OptString = proto.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	return req, nil
}

func _NamedPath_EmbedNamedPathOptString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*EmbedNamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetEmbed().GetOptString()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathOptString").URLPath(pairs...)
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
			r, err := http.NewRequestWithContext(ctx, "GET", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _NamedPath_EmbedNamedPathOptString_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_EmbedNamedPathOptString_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func _NamedPath_EmbedNamedPathWrapString_HttpServer_RequestDecoder(ctx context.Context, r *http.Request) (any, error) {
	req := &EmbedNamedPathRequest{}
	vars := urlx.FormFromMap(mux.Vars(r))
	var varErr error
	if req.Embed == nil {
		req.Embed = &NamedPathRequest{}
	}
	req.Embed.WrapString = wrapperspb.String(fmt.Sprintf("classes/%s/shelves/%s/books/%s/families/%s", vars.Get("class"), vars.Get("shelf"), vars.Get("book"), vars.Get("family")))
	if varErr != nil {
		return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(varErr))
	}
	return req, nil
}

func _NamedPath_EmbedNamedPathWrapString_HttpClient_RequestEncoder(router *mux.Router) func(scheme string, instance string) http1.CreateRequestFunc {
	return func(scheme string, instance string) http1.CreateRequestFunc {
		return func(ctx context.Context, obj any) (*http.Request, error) {
			if obj == nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("request is nil"))
			}
			req, ok := obj.(*EmbedNamedPathRequest)
			if !ok {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid request type, %T", obj))
			}
			_ = req
			var body io.Reader
			var pairs []string
			namedPathParameter := req.GetEmbed().GetWrapString().GetValue()
			namedPathValues := strings.Split(namedPathParameter, "/")
			if len(namedPathValues) != 8 {
				return nil, statusx.ErrInvalidArgument.With(statusx.Message("invalid named path parameter, %s", namedPathParameter))
			}
			pairs = append(pairs, "class", namedPathValues[1], "shelf", namedPathValues[3], "book", namedPathValues[5], "family", namedPathValues[7])
			path, err := router.Get("/leo.example.path.v1.NamedPath/EmbedNamedPathWrapString").URLPath(pairs...)
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
			r, err := http.NewRequestWithContext(ctx, "GET", target.String(), body)
			if err != nil {
				return nil, statusx.ErrInvalidArgument.With(statusx.Wrap(err))
			}
			return r, nil
		}
	}
}

func _NamedPath_EmbedNamedPathWrapString_HttpServer_ResponseEncoder(ctx context.Context, w http.ResponseWriter, obj any) error {
	resp := obj.(*emptypb.Empty)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
		return statusx.ErrInternal.With(statusx.Wrap(err))
	}
	return nil
}

func _NamedPath_EmbedNamedPathWrapString_HttpClient_ResponseDecoder(ctx context.Context, r *http.Response) (any, error) {
	if httpx.IsErrorResponse(r) {
		return nil, httpx.ErrorDecoder(ctx, r)
	}
	resp := &emptypb.Empty{}
	if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
