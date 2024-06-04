// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package response

import (
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	errorx "github.com/go-leo/gox/errorx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc1 "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	io "io"
	http1 "net/http"
	url "net/url"
)

// =========================== endpoints ===========================

type ResponseService interface {
	OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error)
	HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error)
	HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error)
}

type ResponseEndpoints interface {
	OmittedResponse() endpoint.Endpoint
	StarResponse() endpoint.Endpoint
	NamedResponse() endpoint.Endpoint
	HttpBodyResponse() endpoint.Endpoint
	HttpBodyNamedResponse() endpoint.Endpoint
}

type responseEndpoints struct {
	svc         ResponseService
	middlewares []endpoint.Middleware
}

func (e *responseEndpoints) OmittedResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.OmittedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) StarResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.StarResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) NamedResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) HttpBodyResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *responseEndpoints) HttpBodyNamedResponse() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyNamedResponse(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewResponseEndpoints(svc ResponseService, middlewares ...endpoint.Middleware) ResponseEndpoints {
	return &responseEndpoints{svc: svc, middlewares: middlewares}
}

// =========================== cqrs ===========================

// =========================== grpc transports ===========================

type ResponseGrpcServerTransports interface {
	OmittedResponse() *grpc.Server
	StarResponse() *grpc.Server
	NamedResponse() *grpc.Server
	HttpBodyResponse() *grpc.Server
	HttpBodyNamedResponse() *grpc.Server
}

type ResponseGrpcClientTransports interface {
	OmittedResponse() *grpc.Client
	StarResponse() *grpc.Client
	NamedResponse() *grpc.Client
	HttpBodyResponse() *grpc.Client
	HttpBodyNamedResponse() *grpc.Client
}

type responseGrpcServerTransports struct {
	omittedResponse       *grpc.Server
	starResponse          *grpc.Server
	namedResponse         *grpc.Server
	httpBodyResponse      *grpc.Server
	httpBodyNamedResponse *grpc.Server
}

func (t *responseGrpcServerTransports) OmittedResponse() *grpc.Server {
	return t.omittedResponse
}

func (t *responseGrpcServerTransports) StarResponse() *grpc.Server {
	return t.starResponse
}

func (t *responseGrpcServerTransports) NamedResponse() *grpc.Server {
	return t.namedResponse
}

func (t *responseGrpcServerTransports) HttpBodyResponse() *grpc.Server {
	return t.httpBodyResponse
}

func (t *responseGrpcServerTransports) HttpBodyNamedResponse() *grpc.Server {
	return t.httpBodyNamedResponse
}

func NewResponseGrpcServerTransports(endpoints ResponseEndpoints) ResponseGrpcServerTransports {
	return &responseGrpcServerTransports{
		omittedResponse: grpc.NewServer(
			endpoints.OmittedResponse(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/OmittedResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
		),
		starResponse: grpc.NewServer(
			endpoints.StarResponse(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/StarResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
		),
		namedResponse: grpc.NewServer(
			endpoints.NamedResponse(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/NamedResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
		),
		httpBodyResponse: grpc.NewServer(
			endpoints.HttpBodyResponse(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/HttpBodyResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
		),
		httpBodyNamedResponse: grpc.NewServer(
			endpoints.HttpBodyNamedResponse(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.response.v1.Response/HttpBodyNamedResponse")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
		),
	}
}

type responseGrpcClientTransports struct {
	omittedResponse       *grpc.Client
	starResponse          *grpc.Client
	namedResponse         *grpc.Client
	httpBodyResponse      *grpc.Client
	httpBodyNamedResponse *grpc.Client
}

func (t *responseGrpcClientTransports) OmittedResponse() *grpc.Client {
	return t.omittedResponse
}

func (t *responseGrpcClientTransports) StarResponse() *grpc.Client {
	return t.starResponse
}

func (t *responseGrpcClientTransports) NamedResponse() *grpc.Client {
	return t.namedResponse
}

func (t *responseGrpcClientTransports) HttpBodyResponse() *grpc.Client {
	return t.httpBodyResponse
}

func (t *responseGrpcClientTransports) HttpBodyNamedResponse() *grpc.Client {
	return t.httpBodyNamedResponse
}

func NewResponseGrpcClientTransports(conn *grpc1.ClientConn) ResponseGrpcClientTransports {
	return &responseGrpcClientTransports{
		omittedResponse: grpc.NewClient(
			conn,
			"leo.example.response.v1.Response",
			"OmittedResponse",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			UserResponse{},
			grpc.ClientBefore(grpcx.ClientEndpointInjector("/leo.example.response.v1.Response/OmittedResponse")),
			grpc.ClientBefore(grpcx.ClientTransportInjector),
		),
		starResponse: grpc.NewClient(
			conn,
			"leo.example.response.v1.Response",
			"StarResponse",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			UserResponse{},
			grpc.ClientBefore(grpcx.ClientEndpointInjector("/leo.example.response.v1.Response/StarResponse")),
			grpc.ClientBefore(grpcx.ClientTransportInjector),
		),
		namedResponse: grpc.NewClient(
			conn,
			"leo.example.response.v1.Response",
			"NamedResponse",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			UserResponse{},
			grpc.ClientBefore(grpcx.ClientEndpointInjector("/leo.example.response.v1.Response/NamedResponse")),
			grpc.ClientBefore(grpcx.ClientTransportInjector),
		),
		httpBodyResponse: grpc.NewClient(
			conn,
			"leo.example.response.v1.Response",
			"HttpBodyResponse",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			httpbody.HttpBody{},
			grpc.ClientBefore(grpcx.ClientEndpointInjector("/leo.example.response.v1.Response/HttpBodyResponse")),
			grpc.ClientBefore(grpcx.ClientTransportInjector),
		),
		httpBodyNamedResponse: grpc.NewClient(
			conn,
			"leo.example.response.v1.Response",
			"HttpBodyNamedResponse",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			HttpBody{},
			grpc.ClientBefore(grpcx.ClientEndpointInjector("/leo.example.response.v1.Response/HttpBodyNamedResponse")),
			grpc.ClientBefore(grpcx.ClientTransportInjector),
		),
	}
}

type responseGrpcServer struct {
	omittedResponse       *grpc.Server
	starResponse          *grpc.Server
	namedResponse         *grpc.Server
	httpBodyResponse      *grpc.Server
	httpBodyNamedResponse *grpc.Server
}

func (s *responseGrpcServer) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.omittedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.starResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	ctx, rep, err := s.namedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*UserResponse), nil
}

func (s *responseGrpcServer) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	ctx, rep, err := s.httpBodyResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*httpbody.HttpBody), nil
}

func (s *responseGrpcServer) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	ctx, rep, err := s.httpBodyNamedResponse.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*HttpBody), nil
}

func NewResponseGrpcServer(transports ResponseGrpcServerTransports) ResponseService {
	return &responseGrpcServer{
		omittedResponse:       transports.OmittedResponse(),
		starResponse:          transports.StarResponse(),
		namedResponse:         transports.NamedResponse(),
		httpBodyResponse:      transports.HttpBodyResponse(),
		httpBodyNamedResponse: transports.HttpBodyNamedResponse(),
	}
}

type responseGrpcClient struct {
	omittedResponse       endpoint.Endpoint
	starResponse          endpoint.Endpoint
	namedResponse         endpoint.Endpoint
	httpBodyResponse      endpoint.Endpoint
	httpBodyNamedResponse endpoint.Endpoint
}

func (c *responseGrpcClient) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.omittedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseGrpcClient) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.starResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseGrpcClient) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.namedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseGrpcClient) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	rep, err := c.httpBodyResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*httpbody.HttpBody), nil
}

func (c *responseGrpcClient) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	rep, err := c.httpBodyNamedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*HttpBody), nil
}

func NewResponseGrpcClient(transports ResponseGrpcClientTransports, middlewares ...endpoint.Middleware) ResponseService {
	return &responseGrpcClient{
		omittedResponse:       endpointx.Chain(transports.OmittedResponse().Endpoint(), middlewares...),
		starResponse:          endpointx.Chain(transports.StarResponse().Endpoint(), middlewares...),
		namedResponse:         endpointx.Chain(transports.NamedResponse().Endpoint(), middlewares...),
		httpBodyResponse:      endpointx.Chain(transports.HttpBodyResponse().Endpoint(), middlewares...),
		httpBodyNamedResponse: endpointx.Chain(transports.HttpBodyNamedResponse().Endpoint(), middlewares...),
	}
}

// =========================== http transports ===========================

type ResponseHttpServerTransports interface {
	OmittedResponse() *http.Server
	StarResponse() *http.Server
	NamedResponse() *http.Server
	HttpBodyResponse() *http.Server
	HttpBodyNamedResponse() *http.Server
}

type ResponseHttpClientTransports interface {
	OmittedResponse() *http.Client
	StarResponse() *http.Client
	NamedResponse() *http.Client
	HttpBodyResponse() *http.Client
	HttpBodyNamedResponse() *http.Client
}

type responseHttpServerTransports struct {
	omittedResponse       *http.Server
	starResponse          *http.Server
	namedResponse         *http.Server
	httpBodyResponse      *http.Server
	httpBodyNamedResponse *http.Server
}

func (t *responseHttpServerTransports) OmittedResponse() *http.Server {
	return t.omittedResponse
}

func (t *responseHttpServerTransports) StarResponse() *http.Server {
	return t.starResponse
}

func (t *responseHttpServerTransports) NamedResponse() *http.Server {
	return t.namedResponse
}

func (t *responseHttpServerTransports) HttpBodyResponse() *http.Server {
	return t.httpBodyResponse
}

func (t *responseHttpServerTransports) HttpBodyNamedResponse() *http.Server {
	return t.httpBodyNamedResponse
}

func NewResponseHttpServerTransports(endpoints ResponseEndpoints) ResponseHttpServerTransports {
	return &responseHttpServerTransports{
		omittedResponse: http.NewServer(
			endpoints.OmittedResponse(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*UserResponse)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/OmittedResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		starResponse: http.NewServer(
			endpoints.StarResponse(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*UserResponse)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/StarResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		namedResponse: http.NewServer(
			endpoints.NamedResponse(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*UserResponse)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp.GetUser()); err != nil {
					return err
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/NamedResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		httpBodyResponse: http.NewServer(
			endpoints.HttpBodyResponse(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*httpbody.HttpBody)
				w.Header().Set("Content-Type", resp.GetContentType())
				for _, src := range resp.GetExtensions() {
					dst, err := anypb.UnmarshalNew(src, proto.UnmarshalOptions{})
					if err != nil {
						return err
					}
					metadata, ok := dst.(*structpb.Struct)
					if !ok {
						continue
					}
					for key, value := range metadata.GetFields() {
						w.Header().Add(key, string(errorx.Ignore(jsonx.Marshal(value))))
					}
				}
				w.WriteHeader(http1.StatusOK)
				if _, err := w.Write(resp.GetData()); err != nil {
					return err
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/HttpBodyResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		httpBodyNamedResponse: http.NewServer(
			endpoints.HttpBodyNamedResponse(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*HttpBody)
				w.Header().Set("Content-Type", resp.GetBody().GetContentType())
				for _, src := range resp.GetBody().GetExtensions() {
					dst, err := anypb.UnmarshalNew(src, proto.UnmarshalOptions{})
					if err != nil {
						return err
					}
					metadata, ok := dst.(*structpb.Struct)
					if !ok {
						continue
					}
					for key, value := range metadata.GetFields() {
						w.Header().Add(key, string(errorx.Ignore(jsonx.Marshal(value))))
					}
				}
				w.WriteHeader(http1.StatusOK)
				if _, err := w.Write(resp.GetBody().GetData()); err != nil {
					return err
				}
				return nil
			},
			http.ServerBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/HttpBodyNamedResponse")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
	}
}

type responseHttpClientTransports struct {
	omittedResponse       *http.Client
	starResponse          *http.Client
	namedResponse         *http.Client
	httpBodyResponse      *http.Client
	httpBodyNamedResponse *http.Client
}

func (t *responseHttpClientTransports) OmittedResponse() *http.Client {
	return t.omittedResponse
}

func (t *responseHttpClientTransports) StarResponse() *http.Client {
	return t.starResponse
}

func (t *responseHttpClientTransports) NamedResponse() *http.Client {
	return t.namedResponse
}

func (t *responseHttpClientTransports) HttpBodyResponse() *http.Client {
	return t.httpBodyResponse
}

func (t *responseHttpClientTransports) HttpBodyNamedResponse() *http.Client {
	return t.httpBodyNamedResponse
}

func NewResponseHttpClientTransports(scheme string, instance string) ResponseHttpClientTransports {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.response.v1.Response/OmittedResponse").Methods("POST").Path("/v1/omitted/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/StarResponse").Methods("POST").Path("/v1/star/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/NamedResponse").Methods("POST").Path("/v1/named/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyResponse").Methods("PUT").Path("/v1/http/body/omitted/response")
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyNamedResponse").Methods("PUT").Path("/v1/http/body/named/response")
	return &responseHttpClientTransports{
		omittedResponse: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*emptypb.Empty)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				path, err := router.Get("/leo.example.response.v1.Response/OmittedResponse").URLPath(pairs...)
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
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				if httpx.IsErrorResponse(r) {
					return nil, httpx.ErrorDecoder(ctx, r)
				}
				resp := &UserResponse{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/OmittedResponse")),
			http.ClientBefore(httpx.TransportInjector(httpx.HttpClient)),
		),
		starResponse: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*emptypb.Empty)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				path, err := router.Get("/leo.example.response.v1.Response/StarResponse").URLPath(pairs...)
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
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				if httpx.IsErrorResponse(r) {
					return nil, httpx.ErrorDecoder(ctx, r)
				}
				resp := &UserResponse{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/StarResponse")),
			http.ClientBefore(httpx.TransportInjector(httpx.HttpClient)),
		),
		namedResponse: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*emptypb.Empty)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				path, err := router.Get("/leo.example.response.v1.Response/NamedResponse").URLPath(pairs...)
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
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				if httpx.IsErrorResponse(r) {
					return nil, httpx.ErrorDecoder(ctx, r)
				}
				resp := &UserResponse{}
				if err := jsonx.NewDecoder(r.Body).Decode(&resp.User); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/NamedResponse")),
			http.ClientBefore(httpx.TransportInjector(httpx.HttpClient)),
		),
		httpBodyResponse: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*emptypb.Empty)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				path, err := router.Get("/leo.example.response.v1.Response/HttpBodyResponse").URLPath(pairs...)
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
				r, err := http1.NewRequestWithContext(ctx, "PUT", target.String(), body)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
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
			},
			http.ClientBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/HttpBodyResponse")),
			http.ClientBefore(httpx.TransportInjector(httpx.HttpClient)),
		),
		httpBodyNamedResponse: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*emptypb.Empty)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var pairs []string
				path, err := router.Get("/leo.example.response.v1.Response/HttpBodyNamedResponse").URLPath(pairs...)
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
				r, err := http1.NewRequestWithContext(ctx, "PUT", target.String(), body)
				if err != nil {
					return nil, err
				}
				return r, nil
			},
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
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
			},
			http.ClientBefore(httpx.EndpointInjector("/leo.example.response.v1.Response/HttpBodyNamedResponse")),
			http.ClientBefore(httpx.TransportInjector(httpx.HttpClient)),
		),
	}
}

func NewResponseHttpServerHandler(endpoints ResponseHttpServerTransports) http1.Handler {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.response.v1.Response/OmittedResponse").Methods("POST").Path("/v1/omitted/response").Handler(endpoints.OmittedResponse())
	router.NewRoute().Name("/leo.example.response.v1.Response/StarResponse").Methods("POST").Path("/v1/star/response").Handler(endpoints.StarResponse())
	router.NewRoute().Name("/leo.example.response.v1.Response/NamedResponse").Methods("POST").Path("/v1/named/response").Handler(endpoints.NamedResponse())
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyResponse").Methods("PUT").Path("/v1/http/body/omitted/response").Handler(endpoints.HttpBodyResponse())
	router.NewRoute().Name("/leo.example.response.v1.Response/HttpBodyNamedResponse").Methods("PUT").Path("/v1/http/body/named/response").Handler(endpoints.HttpBodyNamedResponse())
	return router
}

type responseHttpClient struct {
	omittedResponse       endpoint.Endpoint
	starResponse          endpoint.Endpoint
	namedResponse         endpoint.Endpoint
	httpBodyResponse      endpoint.Endpoint
	httpBodyNamedResponse endpoint.Endpoint
}

func (c *responseHttpClient) OmittedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.omittedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseHttpClient) StarResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.starResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseHttpClient) NamedResponse(ctx context.Context, request *emptypb.Empty) (*UserResponse, error) {
	rep, err := c.namedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*UserResponse), nil
}

func (c *responseHttpClient) HttpBodyResponse(ctx context.Context, request *emptypb.Empty) (*httpbody.HttpBody, error) {
	rep, err := c.httpBodyResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*httpbody.HttpBody), nil
}

func (c *responseHttpClient) HttpBodyNamedResponse(ctx context.Context, request *emptypb.Empty) (*HttpBody, error) {
	rep, err := c.httpBodyNamedResponse(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*HttpBody), nil
}

func NewResponseHttpClient(transports ResponseHttpClientTransports, middlewares ...endpoint.Middleware) ResponseService {
	return &responseHttpClient{
		omittedResponse:       endpointx.Chain(transports.OmittedResponse().Endpoint(), middlewares...),
		starResponse:          endpointx.Chain(transports.StarResponse().Endpoint(), middlewares...),
		namedResponse:         endpointx.Chain(transports.NamedResponse().Endpoint(), middlewares...),
		httpBodyResponse:      endpointx.Chain(transports.HttpBodyResponse().Endpoint(), middlewares...),
		httpBodyNamedResponse: endpointx.Chain(transports.HttpBodyNamedResponse().Endpoint(), middlewares...),
	}
}
