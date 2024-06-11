// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package body

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
	endpointx "github.com/go-leo/leo/v3/endpointx"
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	grpcx "github.com/go-leo/leo/v3/transportx/grpcx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	mux "github.com/gorilla/mux"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc1 "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
	http1 "net/http"
	url "net/url"
)

// =========================== endpoints ===========================

type BodyService interface {
	StarBody(ctx context.Context, request *User) (*emptypb.Empty, error)
	NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error)
	NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error)
	HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error)
	HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error)
}

type BodyEndpoints interface {
	StarBody() endpoint.Endpoint
	NamedBody() endpoint.Endpoint
	NonBody() endpoint.Endpoint
	HttpBodyStarBody() endpoint.Endpoint
	HttpBodyNamedBody() endpoint.Endpoint
}

type BodyTransports interface {
	StarBody() transportx.Transport
	NamedBody() transportx.Transport
	NonBody() transportx.Transport
	HttpBodyStarBody() transportx.Transport
	HttpBodyNamedBody() transportx.Transport
}

type BodyFactories interface {
	StarBody(middlewares ...endpoint.Middleware) sd.Factory
	NamedBody(middlewares ...endpoint.Middleware) sd.Factory
	NonBody(middlewares ...endpoint.Middleware) sd.Factory
	HttpBodyStarBody(middlewares ...endpoint.Middleware) sd.Factory
	HttpBodyNamedBody(middlewares ...endpoint.Middleware) sd.Factory
}

type BodyEndpointers interface {
	StarBody() sd.Endpointer
	NamedBody() sd.Endpointer
	NonBody() sd.Endpointer
	HttpBodyStarBody() sd.Endpointer
	HttpBodyNamedBody() sd.Endpointer
}

type bodyEndpoints struct {
	svc         BodyService
	middlewares []endpoint.Middleware
}

func (e *bodyEndpoints) StarBody() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.StarBody(ctx, request.(*User))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *bodyEndpoints) NamedBody() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedBody(ctx, request.(*UserRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *bodyEndpoints) NonBody() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NonBody(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *bodyEndpoints) HttpBodyStarBody() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyStarBody(ctx, request.(*httpbody.HttpBody))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func (e *bodyEndpoints) HttpBodyNamedBody() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyNamedBody(ctx, request.(*HttpBody))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewBodyEndpoints(svc BodyService, middlewares ...endpoint.Middleware) BodyEndpoints {
	return &bodyEndpoints{svc: svc, middlewares: middlewares}
}

// =========================== cqrs ===========================

// =========================== grpc server ===========================

type BodyGrpcServerTransports interface {
	StarBody() *grpc.Server
	NamedBody() *grpc.Server
	NonBody() *grpc.Server
	HttpBodyStarBody() *grpc.Server
	HttpBodyNamedBody() *grpc.Server
}

type bodyGrpcServerTransports struct {
	starBody          *grpc.Server
	namedBody         *grpc.Server
	nonBody           *grpc.Server
	httpBodyStarBody  *grpc.Server
	httpBodyNamedBody *grpc.Server
}

func (t *bodyGrpcServerTransports) StarBody() *grpc.Server {
	return t.starBody
}

func (t *bodyGrpcServerTransports) NamedBody() *grpc.Server {
	return t.namedBody
}

func (t *bodyGrpcServerTransports) NonBody() *grpc.Server {
	return t.nonBody
}

func (t *bodyGrpcServerTransports) HttpBodyStarBody() *grpc.Server {
	return t.httpBodyStarBody
}

func (t *bodyGrpcServerTransports) HttpBodyNamedBody() *grpc.Server {
	return t.httpBodyNamedBody
}

func NewBodyGrpcServerTransports(endpoints BodyEndpoints) BodyGrpcServerTransports {
	return &bodyGrpcServerTransports{
		starBody: grpc.NewServer(
			endpoints.StarBody(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/StarBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		namedBody: grpc.NewServer(
			endpoints.NamedBody(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/NamedBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		nonBody: grpc.NewServer(
			endpoints.NonBody(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/NonBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		httpBodyStarBody: grpc.NewServer(
			endpoints.HttpBodyStarBody(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/HttpBodyStarBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
		httpBodyNamedBody: grpc.NewServer(
			endpoints.HttpBodyNamedBody(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			grpc.ServerBefore(grpcx.ServerEndpointInjector("/leo.example.body.v1.Body/HttpBodyNamedBody")),
			grpc.ServerBefore(grpcx.ServerTransportInjector),
			grpc.ServerBefore(grpcx.IncomingMetadata),
		),
	}
}

type bodyGrpcServer struct {
	starBody          *grpc.Server
	namedBody         *grpc.Server
	nonBody           *grpc.Server
	httpBodyStarBody  *grpc.Server
	httpBodyNamedBody *grpc.Server
}

func (s *bodyGrpcServer) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	ctx, rep, err := s.starBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	ctx, rep, err := s.namedBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	ctx, rep, err := s.nonBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpBodyStarBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func (s *bodyGrpcServer) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	ctx, rep, err := s.httpBodyNamedBody.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*emptypb.Empty), nil
}

func NewBodyGrpcServer(transports BodyGrpcServerTransports) BodyService {
	return &bodyGrpcServer{
		starBody:          transports.StarBody(),
		namedBody:         transports.NamedBody(),
		nonBody:           transports.NonBody(),
		httpBodyStarBody:  transports.HttpBodyStarBody(),
		httpBodyNamedBody: transports.HttpBodyNamedBody(),
	}
}

// =========================== grpc client ===========================

type bodyGrpcClientTransports struct {
	starBody          transportx.Transport
	namedBody         transportx.Transport
	nonBody           transportx.Transport
	httpBodyStarBody  transportx.Transport
	httpBodyNamedBody transportx.Transport
}

func (t *bodyGrpcClientTransports) StarBody() transportx.Transport {
	return t.starBody
}

func (t *bodyGrpcClientTransports) NamedBody() transportx.Transport {
	return t.namedBody
}

func (t *bodyGrpcClientTransports) NonBody() transportx.Transport {
	return t.nonBody
}

func (t *bodyGrpcClientTransports) HttpBodyStarBody() transportx.Transport {
	return t.httpBodyStarBody
}

func (t *bodyGrpcClientTransports) HttpBodyNamedBody() transportx.Transport {
	return t.httpBodyNamedBody
}

func NewBodyGrpcClientTransports(conn *grpc1.ClientConn) BodyTransports {
	return &bodyGrpcClientTransports{
		starBody: grpc.NewClient(
			conn,
			"leo.example.body.v1.Body",
			"StarBody",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
		namedBody: grpc.NewClient(
			conn,
			"leo.example.body.v1.Body",
			"NamedBody",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
		nonBody: grpc.NewClient(
			conn,
			"leo.example.body.v1.Body",
			"NonBody",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
		httpBodyStarBody: grpc.NewClient(
			conn,
			"leo.example.body.v1.Body",
			"HttpBodyStarBody",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
		httpBodyNamedBody: grpc.NewClient(
			conn,
			"leo.example.body.v1.Body",
			"HttpBodyNamedBody",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			emptypb.Empty{},
			grpc.ClientBefore(grpcx.OutgoingMetadata),
		),
	}
}

type bodyGrpcClientEndpoints struct {
	transports  BodyTransports
	middlewares []endpoint.Middleware
}

func (e *bodyGrpcClientEndpoints) StarBody() endpoint.Endpoint {
	return endpointx.Chain(e.transports.StarBody().Endpoint(), e.middlewares...)
}

func (e *bodyGrpcClientEndpoints) NamedBody() endpoint.Endpoint {
	return endpointx.Chain(e.transports.NamedBody().Endpoint(), e.middlewares...)
}

func (e *bodyGrpcClientEndpoints) NonBody() endpoint.Endpoint {
	return endpointx.Chain(e.transports.NonBody().Endpoint(), e.middlewares...)
}

func (e *bodyGrpcClientEndpoints) HttpBodyStarBody() endpoint.Endpoint {
	return endpointx.Chain(e.transports.HttpBodyStarBody().Endpoint(), e.middlewares...)
}

func (e *bodyGrpcClientEndpoints) HttpBodyNamedBody() endpoint.Endpoint {
	return endpointx.Chain(e.transports.HttpBodyNamedBody().Endpoint(), e.middlewares...)
}

func NewBodyGrpcClientEndpoints(transports BodyTransports, middlewares ...endpoint.Middleware) BodyEndpoints {
	return &bodyGrpcClientEndpoints{transports: transports, middlewares: middlewares}
}

type bodyGrpcClient struct {
	starBody          endpoint.Endpoint
	namedBody         endpoint.Endpoint
	nonBody           endpoint.Endpoint
	httpBodyStarBody  endpoint.Endpoint
	httpBodyNamedBody endpoint.Endpoint
}

func (c *bodyGrpcClient) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/StarBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.starBody(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGrpcClient) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/NamedBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.namedBody(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGrpcClient) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/NonBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.nonBody(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGrpcClient) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/HttpBodyStarBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.httpBodyStarBody(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyGrpcClient) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/HttpBodyNamedBody")
	ctx = transportx.InjectName(ctx, grpcx.GrpcClient)
	rep, err := c.httpBodyNamedBody(ctx, request)
	if err != nil {
		return nil, statusx.FromGrpcError(err)
	}
	return rep.(*emptypb.Empty), nil
}

func NewBodyGrpcClient(endpoints BodyEndpoints) BodyService {
	return &bodyGrpcClient{
		starBody:          endpoints.StarBody(),
		namedBody:         endpoints.NamedBody(),
		nonBody:           endpoints.NonBody(),
		httpBodyStarBody:  endpoints.HttpBodyStarBody(),
		httpBodyNamedBody: endpoints.HttpBodyNamedBody(),
	}
}

type bodyGrpcClientFactories struct {
	opts []grpc1.DialOption
}

func (f *bodyGrpcClientFactories) StarBody(middlewares ...endpoint.Middleware) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc1.NewClient(instance, f.opts...)
		if err != nil {
			return nil, nil, err
		}
		transports := NewBodyGrpcClientTransports(conn)
		endpoints := NewBodyGrpcClientEndpoints(transports, middlewares...)
		return endpoints.StarBody(), conn, nil
	}
}

func (f *bodyGrpcClientFactories) NamedBody(middlewares ...endpoint.Middleware) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc1.NewClient(instance, f.opts...)
		if err != nil {
			return nil, nil, err
		}
		transports := NewBodyGrpcClientTransports(conn)
		endpoints := NewBodyGrpcClientEndpoints(transports, middlewares...)
		return endpoints.NamedBody(), conn, nil
	}
}

func (f *bodyGrpcClientFactories) NonBody(middlewares ...endpoint.Middleware) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc1.NewClient(instance, f.opts...)
		if err != nil {
			return nil, nil, err
		}
		transports := NewBodyGrpcClientTransports(conn)
		endpoints := NewBodyGrpcClientEndpoints(transports, middlewares...)
		return endpoints.NonBody(), conn, nil
	}
}

func (f *bodyGrpcClientFactories) HttpBodyStarBody(middlewares ...endpoint.Middleware) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc1.NewClient(instance, f.opts...)
		if err != nil {
			return nil, nil, err
		}
		transports := NewBodyGrpcClientTransports(conn)
		endpoints := NewBodyGrpcClientEndpoints(transports, middlewares...)
		return endpoints.HttpBodyStarBody(), conn, nil
	}
}

func (f *bodyGrpcClientFactories) HttpBodyNamedBody(middlewares ...endpoint.Middleware) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc1.NewClient(instance, f.opts...)
		if err != nil {
			return nil, nil, err
		}
		transports := NewBodyGrpcClientTransports(conn)
		endpoints := NewBodyGrpcClientEndpoints(transports, middlewares...)
		return endpoints.HttpBodyNamedBody(), conn, nil
	}
}

func NewBodyGrpcClientFactories(opts ...grpc1.DialOption) BodyFactories {
	return &bodyGrpcClientFactories{opts: opts}
}

// =========================== http server ===========================

type BodyHttpServerTransports interface {
	StarBody() *http.Server
	NamedBody() *http.Server
	NonBody() *http.Server
	HttpBodyStarBody() *http.Server
	HttpBodyNamedBody() *http.Server
}

type bodyHttpServerTransports struct {
	starBody          *http.Server
	namedBody         *http.Server
	nonBody           *http.Server
	httpBodyStarBody  *http.Server
	httpBodyNamedBody *http.Server
}

func (t *bodyHttpServerTransports) StarBody() *http.Server {
	return t.starBody
}

func (t *bodyHttpServerTransports) NamedBody() *http.Server {
	return t.namedBody
}

func (t *bodyHttpServerTransports) NonBody() *http.Server {
	return t.nonBody
}

func (t *bodyHttpServerTransports) HttpBodyStarBody() *http.Server {
	return t.httpBodyStarBody
}

func (t *bodyHttpServerTransports) HttpBodyNamedBody() *http.Server {
	return t.httpBodyNamedBody
}

func NewBodyHttpServerTransports(endpoints BodyEndpoints) BodyHttpServerTransports {
	return &bodyHttpServerTransports{
		starBody: http.NewServer(
			endpoints.StarBody(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &User{}
				if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
					return nil, err
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
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/StarBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		namedBody: http.NewServer(
			endpoints.NamedBody(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &UserRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(&req.User); err != nil {
					return nil, err
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
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/NamedBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		nonBody: http.NewServer(
			endpoints.NonBody(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &emptypb.Empty{}
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
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/NonBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		httpBodyStarBody: http.NewServer(
			endpoints.HttpBodyStarBody(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &httpbody.HttpBody{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				req.Data = body
				req.ContentType = r.Header.Get("Content-Type")
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
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/HttpBodyStarBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
		httpBodyNamedBody: http.NewServer(
			endpoints.HttpBodyNamedBody(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &HttpBody{}
				req.Body = &httpbody.HttpBody{}
				body, err := io.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				req.Body.Data = body
				req.Body.ContentType = r.Header.Get("Content-Type")
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
			http.ServerBefore(httpx.EndpointInjector("/leo.example.body.v1.Body/HttpBodyNamedBody")),
			http.ServerBefore(httpx.TransportInjector(httpx.HttpServer)),
			http.ServerBefore(httpx.IncomingMetadata),
			http.ServerErrorEncoder(httpx.ErrorEncoder),
		),
	}
}

func NewBodyHttpServerHandler(endpoints BodyHttpServerTransports) http1.Handler {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.body.v1.Body/StarBody").Methods("POST").Path("/v1/star/body").Handler(endpoints.StarBody())
	router.NewRoute().Name("/leo.example.body.v1.Body/NamedBody").Methods("POST").Path("/v1/named/body").Handler(endpoints.NamedBody())
	router.NewRoute().Name("/leo.example.body.v1.Body/NonBody").Methods("GET").Path("/v1/user_body").Handler(endpoints.NonBody())
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyStarBody").Methods("PUT").Path("/v1/http/body/star/body").Handler(endpoints.HttpBodyStarBody())
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyNamedBody").Methods("PUT").Path("/v1/http/body/named/body").Handler(endpoints.HttpBodyNamedBody())
	return router
}

// =========================== http client ===========================

type bodyHttpClientTransports struct {
	starBody          transportx.Transport
	namedBody         transportx.Transport
	nonBody           transportx.Transport
	httpBodyStarBody  transportx.Transport
	httpBodyNamedBody transportx.Transport
}

func (t *bodyHttpClientTransports) StarBody() transportx.Transport {
	return t.starBody
}

func (t *bodyHttpClientTransports) NamedBody() transportx.Transport {
	return t.namedBody
}

func (t *bodyHttpClientTransports) NonBody() transportx.Transport {
	return t.nonBody
}

func (t *bodyHttpClientTransports) HttpBodyStarBody() transportx.Transport {
	return t.httpBodyStarBody
}

func (t *bodyHttpClientTransports) HttpBodyNamedBody() transportx.Transport {
	return t.httpBodyNamedBody
}

func NewBodyHttpClientTransports(scheme string, instance string) BodyTransports {
	router := mux.NewRouter()
	router.NewRoute().Name("/leo.example.body.v1.Body/StarBody").Methods("POST").Path("/v1/star/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/NamedBody").Methods("POST").Path("/v1/named/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/NonBody").Methods("GET").Path("/v1/user_body")
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyStarBody").Methods("PUT").Path("/v1/http/body/star/body")
	router.NewRoute().Name("/leo.example.body.v1.Body/HttpBodyNamedBody").Methods("PUT").Path("/v1/http/body/named/body")
	return &bodyHttpClientTransports{
		starBody: http.NewExplicitClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*User)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var bodyBuf bytes.Buffer
				if err := jsonx.NewEncoder(&bodyBuf).Encode(req); err != nil {
					return nil, err
				}
				body = &bodyBuf
				contentType := "application/json; charset=utf-8"
				var pairs []string
				path, err := router.Get("/leo.example.body.v1.Body/StarBody").URLPath(pairs...)
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
				resp := &emptypb.Empty{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.OutgoingMetadata),
		),
		namedBody: http.NewExplicitClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*UserRequest)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				var bodyBuf bytes.Buffer
				if err := jsonx.NewEncoder(&bodyBuf).Encode(req.GetUser()); err != nil {
					return nil, err
				}
				body = &bodyBuf
				contentType := "application/json; charset=utf-8"
				var pairs []string
				path, err := router.Get("/leo.example.body.v1.Body/NamedBody").URLPath(pairs...)
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
				resp := &emptypb.Empty{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.OutgoingMetadata),
		),
		nonBody: http.NewExplicitClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
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
				path, err := router.Get("/leo.example.body.v1.Body/NonBody").URLPath(pairs...)
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
				resp := &emptypb.Empty{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			http.ClientBefore(httpx.OutgoingMetadata),
		),
		httpBodyStarBody: http.NewExplicitClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*httpbody.HttpBody)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				body = bytes.NewReader(req.GetData())
				contentType := req.GetContentType()
				var pairs []string
				path, err := router.Get("/leo.example.body.v1.Body/HttpBodyStarBody").URLPath(pairs...)
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
				r.Header.Set("Content-Type", contentType)
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
		httpBodyNamedBody: http.NewExplicitClient(
			func(ctx context.Context, obj any) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*HttpBody)
				if !ok {
					return nil, fmt.Errorf("invalid request object type, %T", obj)
				}
				_ = req
				var body io.Reader
				body = bytes.NewReader(req.GetBody().GetData())
				contentType := req.GetBody().GetContentType()
				var pairs []string
				path, err := router.Get("/leo.example.body.v1.Body/HttpBodyNamedBody").URLPath(pairs...)
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
				r.Header.Set("Content-Type", contentType)
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

type bodyHttpClient struct {
	starBody          endpoint.Endpoint
	namedBody         endpoint.Endpoint
	nonBody           endpoint.Endpoint
	httpBodyStarBody  endpoint.Endpoint
	httpBodyNamedBody endpoint.Endpoint
}

func (c *bodyHttpClient) StarBody(ctx context.Context, request *User) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/StarBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.starBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyHttpClient) NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/NamedBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.namedBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyHttpClient) NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/NonBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.nonBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyHttpClient) HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/HttpBodyStarBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.httpBodyStarBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func (c *bodyHttpClient) HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error) {
	ctx = endpointx.InjectName(ctx, "/leo.example.body.v1.Body/HttpBodyNamedBody")
	ctx = transportx.InjectName(ctx, httpx.HttpClient)
	rep, err := c.httpBodyNamedBody(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*emptypb.Empty), nil
}

func NewBodyHttpClient(transports BodyTransports, middlewares ...endpoint.Middleware) BodyService {
	return &bodyHttpClient{
		starBody:          endpointx.Chain(transports.StarBody().Endpoint(), middlewares...),
		namedBody:         endpointx.Chain(transports.NamedBody().Endpoint(), middlewares...),
		nonBody:           endpointx.Chain(transports.NonBody().Endpoint(), middlewares...),
		httpBodyStarBody:  endpointx.Chain(transports.HttpBodyStarBody().Endpoint(), middlewares...),
		httpBodyNamedBody: endpointx.Chain(transports.HttpBodyNamedBody().Endpoint(), middlewares...),
	}
}
