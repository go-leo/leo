// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc "github.com/go-kit/kit/transport/grpc"
	http "github.com/go-kit/kit/transport/http"
	jsonx "github.com/go-leo/gox/encodingx/jsonx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	transportx "github.com/go-leo/leo/v3/transportx"
	mux "github.com/gorilla/mux"
	grpc1 "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	io "io"
	http1 "net/http"
	url "net/url"
)

// =========================== endpoints ===========================

type GreeterService interface {
	SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error)
}

type GreeterEndpoints interface {
	SayHello() endpoint.Endpoint
}

type greeterEndpoints struct {
	svc         GreeterService
	middlewares []endpoint.Middleware
}

func (e *greeterEndpoints) SayHello() endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.SayHello(ctx, request.(*HelloRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}

func NewGreeterEndpoints(svc GreeterService, middlewares ...endpoint.Middleware) GreeterEndpoints {
	return &greeterEndpoints{svc: svc, middlewares: middlewares}
}

// =========================== cqrs ===========================

// =========================== grpc transports ===========================

type GreeterGrpcServerTransports interface {
	SayHello() *grpc.Server
}

type GreeterGrpcClientTransports interface {
	SayHello() *grpc.Client
}

type greeterGrpcServerTransports struct {
	sayHello *grpc.Server
}

func (t *greeterGrpcServerTransports) SayHello() *grpc.Server {
	return t.sayHello
}

func NewGreeterGrpcServerTransports(endpoints GreeterEndpoints, serverOptions ...grpc.ServerOption) GreeterGrpcServerTransports {
	return &greeterGrpcServerTransports{
		sayHello: grpc.NewServer(
			endpoints.SayHello(),
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			append([]grpc.ServerOption{
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/helloworld.Greeter/SayHello")
				}),
				grpc.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcServer)
				}),
			}, serverOptions...)...,
		),
	}
}

type greeterGrpcClientTransports struct {
	sayHello *grpc.Client
}

func (t *greeterGrpcClientTransports) SayHello() *grpc.Client {
	return t.sayHello
}

func NewGreeterGrpcClientTransports(conn *grpc1.ClientConn, clientOptions ...grpc.ClientOption) GreeterGrpcClientTransports {
	return &greeterGrpcClientTransports{
		sayHello: grpc.NewClient(
			conn,
			"helloworld.Greeter",
			"SayHello",
			func(_ context.Context, v any) (any, error) { return v, nil },
			func(_ context.Context, v any) (any, error) { return v, nil },
			HelloReply{},
			append([]grpc.ClientOption{
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return endpointx.InjectName(ctx, "/helloworld.Greeter/SayHello")
				}),
				grpc.ClientBefore(func(ctx context.Context, md *metadata.MD) context.Context {
					return transportx.InjectName(ctx, transportx.GrpcClient)
				}),
			}, clientOptions...)...,
		),
	}
}

type greeterGrpcServer struct {
	sayHello *grpc.Server
}

func (s *greeterGrpcServer) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	ctx, rep, err := s.sayHello.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	_ = ctx
	return rep.(*HelloReply), nil
}

func NewGreeterGrpcServer(transports GreeterGrpcServerTransports) GreeterService {
	return &greeterGrpcServer{
		sayHello: transports.SayHello(),
	}
}

type greeterGrpcClient struct {
	sayHello endpoint.Endpoint
}

func (c *greeterGrpcClient) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	rep, err := c.sayHello(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*HelloReply), nil
}

func NewGreeterGrpcClient(transports GreeterGrpcClientTransports, middlewares ...endpoint.Middleware) GreeterService {
	return &greeterGrpcClient{
		sayHello: endpointx.Chain(transports.SayHello().Endpoint(), middlewares...),
	}
}

// =========================== http transports ===========================

type GreeterHttpServerTransports interface {
	SayHello() *http.Server
}

type GreeterHttpClientTransports interface {
	SayHello() *http.Client
}

type greeterHttpServerTransports struct {
	sayHello *http.Server
}

func (t *greeterHttpServerTransports) SayHello() *http.Server {
	return t.sayHello
}

func NewGreeterHttpServerTransports(endpoints GreeterEndpoints, serverOptions ...http.ServerOption) GreeterHttpServerTransports {
	return &greeterHttpServerTransports{
		sayHello: http.NewServer(
			endpoints.SayHello(),
			func(ctx context.Context, r *http1.Request) (any, error) {
				req := &HelloRequest{}
				if err := jsonx.NewDecoder(r.Body).Decode(req); err != nil {
					return nil, err
				}
				return req, nil
			},
			func(ctx context.Context, w http1.ResponseWriter, obj any) error {
				resp := obj.(*HelloReply)
				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http1.StatusOK)
				if err := jsonx.NewEncoder(w).Encode(resp); err != nil {
					return err
				}
				return nil
			},
			append([]http.ServerOption{
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/helloworld.Greeter/SayHello")
				}),
				http.ServerBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpServer)
				}),
			}, serverOptions...)...,
		),
	}
}

type greeterHttpClientTransports struct {
	sayHello *http.Client
}

func (t *greeterHttpClientTransports) SayHello() *http.Client {
	return t.sayHello
}

func NewGreeterHttpClientTransports(scheme string, instance string, clientOptions ...http.ClientOption) GreeterHttpClientTransports {
	router := mux.NewRouter()
	router.NewRoute().Name("/helloworld.Greeter/SayHello").Methods("POST").Path("/helloworld.Greeter/SayHello")
	return &greeterHttpClientTransports{
		sayHello: http.NewExplicitClient(
			func(ctx context.Context, obj interface{}) (*http1.Request, error) {
				if obj == nil {
					return nil, errors.New("request object is nil")
				}
				req, ok := obj.(*HelloRequest)
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
				path, err := router.Get("/helloworld.Greeter/SayHello").URLPath(pairs...)
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
			func(ctx context.Context, r *http1.Response) (interface{}, error) {
				resp := &HelloReply{}
				if err := jsonx.NewDecoder(r.Body).Decode(resp); err != nil {
					return nil, err
				}
				return resp, nil
			},
			append([]http.ClientOption{
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return endpointx.InjectName(ctx, "/helloworld.Greeter/SayHello")
				}),
				http.ClientBefore(func(ctx context.Context, request *http1.Request) context.Context {
					return transportx.InjectName(ctx, transportx.HttpClient)
				}),
			}, clientOptions...)...,
		),
	}
}

func NewGreeterHttpServerHandler(endpoints GreeterHttpServerTransports) http1.Handler {
	router := mux.NewRouter()
	router.NewRoute().Name("/helloworld.Greeter/SayHello").Methods("POST").Path("/helloworld.Greeter/SayHello").Handler(endpoints.SayHello())
	return router
}

type greeterHttpClient struct {
	sayHello endpoint.Endpoint
}

func (c *greeterHttpClient) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	rep, err := c.sayHello(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*HelloReply), nil
}

func NewGreeterHttpClient(transports GreeterHttpClientTransports, middlewares ...endpoint.Middleware) GreeterService {
	return &greeterHttpClient{
		sayHello: endpointx.Chain(transports.SayHello().Endpoint(), middlewares...),
	}
}