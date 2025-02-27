// Code generated by protoc-gen-leo. DO NOT EDIT.

package helloworld

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	endpoint "github.com/go-kit/kit/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	httpx1 "github.com/go-leo/gox/netx/httpx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	httpx "github.com/go-leo/leo/v3/transportx/httpx"
	coder "github.com/go-leo/leo/v3/transportx/httpx/coder"
	mux "github.com/gorilla/mux"
	protojson "google.golang.org/protobuf/encoding/protojson"
	io "io"
	http "net/http"
	url "net/url"
)

func appendGreeterHttpRoutes(router *mux.Router) *mux.Router {
	router.NewRoute().
		Name("/helloworld.Greeter/SayHello").
		Methods(http.MethodPost).
		Path("/v1/example/echo")
	return router
}
func AppendGreeterHttpServerRoutes(router *mux.Router, svc GreeterService, middlewares ...endpoint.Middleware) *mux.Router {
	endpoints := &greeterServerEndpoints{
		svc:         svc,
		middlewares: middlewares,
	}
	transports := &greeterHttpServerTransports{
		endpoints:       endpoints,
		requestDecoder:  greeterHttpServerRequestDecoder{},
		responseEncoder: greeterHttpServerResponseEncoder{},
	}
	router = appendGreeterHttpRoutes(router)
	router.Get("/helloworld.Greeter/SayHello").Handler(transports.SayHello())
	return router
}

func NewGreeterHttpClient(target string, opts ...httpx.ClientOption) GreeterService {
	options := httpx.NewClientOptions(opts...)
	requestEncoder := &greeterHttpClientRequestEncoder{
		router: appendGreeterHttpRoutes(mux.NewRouter()),
		scheme: options.Scheme(),
	}
	responseDecoder := &greeterHttpClientResponseDecoder{}
	transports := &greeterHttpClientTransports{
		clientOptions:   options.ClientTransportOptions(),
		middlewares:     options.Middlewares(),
		requestEncoder:  requestEncoder,
		responseDecoder: responseDecoder,
	}
	factories := &greeterFactories{
		transports: transports,
	}
	endpointer := &greeterEndpointers{
		target:    target,
		builder:   options.Builder(),
		factories: factories,
		logger:    options.Logger(),
		options:   options.EndpointerOptions(),
	}
	balancers := &greeterBalancers{
		factory:    options.BalancerFactory(),
		endpointer: endpointer,
	}
	endpoints := &greeterClientEndpoints{
		balancers: balancers,
	}
	return &greeterClientService{
		endpoints:     endpoints,
		transportName: httpx.HttpClient,
	}
}

type GreeterHttpServerTransports interface {
	SayHello() http.Handler
}

type GreeterHttpServerRequestDecoder interface {
	SayHello() http1.DecodeRequestFunc
}

type GreeterHttpServerResponseEncoder interface {
	SayHello() http1.EncodeResponseFunc
}

type GreeterHttpClientRequestEncoder interface {
	SayHello(instance string) http1.CreateRequestFunc
}

type GreeterHttpClientResponseDecoder interface {
	SayHello() http1.DecodeResponseFunc
}

type greeterHttpServerTransports struct {
	endpoints       GreeterServerEndpoints
	requestDecoder  GreeterHttpServerRequestDecoder
	responseEncoder GreeterHttpServerResponseEncoder
}

func (t *greeterHttpServerTransports) SayHello() http.Handler {
	return http1.NewServer(
		t.endpoints.SayHello(context.TODO()),
		t.requestDecoder.SayHello(),
		t.responseEncoder.SayHello(),
		http1.ServerBefore(httpx.EndpointInjector("/helloworld.Greeter/SayHello")),
		http1.ServerBefore(httpx.ServerTransportInjector),
		http1.ServerBefore(httpx.IncomingMetadataInjector),
		http1.ServerBefore(httpx.IncomingTimeLimitInjector),
		http1.ServerBefore(httpx.IncomingStainInjector),
		http1.ServerFinalizer(httpx.CancelInvoker),
	)
}

type greeterHttpServerRequestDecoder struct {
	unmarshalOptions protojson.UnmarshalOptions
}

func (decoder greeterHttpServerRequestDecoder) SayHello() http1.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (any, error) {
		req := &HelloRequest{}
		if err := coder.DecodeMessageFromRequest(ctx, r, req, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return req, nil
	}
}

type greeterHttpServerResponseEncoder struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer coder.ResponseTransformer
}

func (encoder greeterHttpServerResponseEncoder) SayHello() http1.EncodeResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter, obj any) error {
		resp := obj.(*HelloReply)
		return coder.EncodeMessageToResponse(ctx, w, encoder.responseTransformer(ctx, resp), encoder.marshalOptions)
	}
}

type greeterHttpClientTransports struct {
	clientOptions   []http1.ClientOption
	middlewares     []endpoint.Middleware
	requestEncoder  GreeterHttpClientRequestEncoder
	responseDecoder GreeterHttpClientResponseDecoder
}

func (t *greeterHttpClientTransports) SayHello(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error) {
	opts := []http1.ClientOption{
		http1.ClientBefore(httpx.OutgoingMetadataInjector),
		http1.ClientBefore(httpx.OutgoingTimeLimitInjector),
		http1.ClientBefore(httpx.OutgoingStainInjector),
	}
	opts = append(opts, t.clientOptions...)
	client := http1.NewExplicitClient(
		t.requestEncoder.SayHello(instance),
		t.responseDecoder.SayHello(),
		opts...,
	)
	return endpointx.Chain(client.Endpoint(), t.middlewares...), nil, nil
}

type greeterHttpClientRequestEncoder struct {
	marshalOptions   protojson.MarshalOptions
	unmarshalOptions protojson.UnmarshalOptions
	router           *mux.Router
	scheme           string
}

func (encoder greeterHttpClientRequestEncoder) SayHello(instance string) http1.CreateRequestFunc {
	return func(ctx context.Context, obj any) (*http.Request, error) {
		if obj == nil {
			return nil, errors.New("request is nil")
		}
		req, ok := obj.(*HelloRequest)
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
		path, err := encoder.router.Get("/helloworld.Greeter/SayHello").URLPath(pairs...)
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

type greeterHttpClientResponseDecoder struct {
	marshalOptions      protojson.MarshalOptions
	unmarshalOptions    protojson.UnmarshalOptions
	responseTransformer coder.ResponseTransformer
}

func (decoder greeterHttpClientResponseDecoder) SayHello() http1.DecodeResponseFunc {
	return func(ctx context.Context, r *http.Response) (any, error) {
		resp := &HelloReply{}
		if err := coder.DecodeMessageFromResponse(ctx, r, resp, decoder.unmarshalOptions); err != nil {
			return nil, err
		}
		return resp, nil
	}
}
