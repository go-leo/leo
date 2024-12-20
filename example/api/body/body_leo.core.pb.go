// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package body

import (
	context "context"
	endpoint "github.com/go-kit/kit/endpoint"
	sd "github.com/go-kit/kit/sd"
	lb "github.com/go-kit/kit/sd/lb"
	log "github.com/go-kit/log"
	lazyloadx "github.com/go-leo/gox/syncx/lazyloadx"
	endpointx "github.com/go-leo/leo/v3/endpointx"
	sdx "github.com/go-leo/leo/v3/sdx"
	lbx "github.com/go-leo/leo/v3/sdx/lbx"
	stainx "github.com/go-leo/leo/v3/sdx/stainx"
	transportx "github.com/go-leo/leo/v3/transportx"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

type BodyService interface {
	StarBody(ctx context.Context, request *User) (*emptypb.Empty, error)
	NamedBody(ctx context.Context, request *UserRequest) (*emptypb.Empty, error)
	NonBody(ctx context.Context, request *emptypb.Empty) (*emptypb.Empty, error)
	HttpBodyStarBody(ctx context.Context, request *httpbody.HttpBody) (*emptypb.Empty, error)
	HttpBodyNamedBody(ctx context.Context, request *HttpBody) (*emptypb.Empty, error)
}

type BodyEndpoints interface {
	StarBody(ctx context.Context) endpoint.Endpoint
	NamedBody(ctx context.Context) endpoint.Endpoint
	NonBody(ctx context.Context) endpoint.Endpoint
	HttpBodyStarBody(ctx context.Context) endpoint.Endpoint
	HttpBodyNamedBody(ctx context.Context) endpoint.Endpoint
}

type BodyClientTransports interface {
	StarBody() transportx.ClientTransport
	NamedBody() transportx.ClientTransport
	NonBody() transportx.ClientTransport
	HttpBodyStarBody() transportx.ClientTransport
	HttpBodyNamedBody() transportx.ClientTransport
}
type BodyClientTransportsV2 interface {
	StarBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	NamedBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	NonBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	HttpBodyStarBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	HttpBodyNamedBody(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

type BodyFactories interface {
	StarBody(ctx context.Context) sd.Factory
	NamedBody(ctx context.Context) sd.Factory
	NonBody(ctx context.Context) sd.Factory
	HttpBodyStarBody(ctx context.Context) sd.Factory
	HttpBodyNamedBody(ctx context.Context) sd.Factory
}

type BodyEndpointers interface {
	StarBody(ctx context.Context, color string) (sd.Endpointer, error)
	NamedBody(ctx context.Context, color string) (sd.Endpointer, error)
	NonBody(ctx context.Context, color string) (sd.Endpointer, error)
	HttpBodyStarBody(ctx context.Context, color string) (sd.Endpointer, error)
	HttpBodyNamedBody(ctx context.Context, color string) (sd.Endpointer, error)
}

type BodyBalancers interface {
	StarBody(ctx context.Context) (lb.Balancer, error)
	NamedBody(ctx context.Context) (lb.Balancer, error)
	NonBody(ctx context.Context) (lb.Balancer, error)
	HttpBodyStarBody(ctx context.Context) (lb.Balancer, error)
	HttpBodyNamedBody(ctx context.Context) (lb.Balancer, error)
}

type bodyServerEndpoints struct {
	svc         BodyService
	middlewares []endpoint.Middleware
}

func (e *bodyServerEndpoints) StarBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.StarBody(ctx, request.(*User))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *bodyServerEndpoints) NamedBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedBody(ctx, request.(*UserRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *bodyServerEndpoints) NonBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NonBody(ctx, request.(*emptypb.Empty))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *bodyServerEndpoints) HttpBodyStarBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyStarBody(ctx, request.(*httpbody.HttpBody))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *bodyServerEndpoints) HttpBodyNamedBody(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.HttpBodyNamedBody(ctx, request.(*HttpBody))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func newBodyServerEndpoints(svc BodyService, middlewares ...endpoint.Middleware) BodyEndpoints {
	return &bodyServerEndpoints{svc: svc, middlewares: middlewares}
}

type bodyClientEndpoints struct {
	transports  BodyClientTransports
	middlewares []endpoint.Middleware
}

func (e *bodyClientEndpoints) StarBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.StarBody().Endpoint(ctx), e.middlewares...)
}
func (e *bodyClientEndpoints) NamedBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.NamedBody().Endpoint(ctx), e.middlewares...)
}
func (e *bodyClientEndpoints) NonBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.NonBody().Endpoint(ctx), e.middlewares...)
}
func (e *bodyClientEndpoints) HttpBodyStarBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.HttpBodyStarBody().Endpoint(ctx), e.middlewares...)
}
func (e *bodyClientEndpoints) HttpBodyNamedBody(ctx context.Context) endpoint.Endpoint {
	return endpointx.Chain(e.transports.HttpBodyNamedBody().Endpoint(ctx), e.middlewares...)
}
func newBodyClientEndpoints(transports BodyClientTransports, middlewares ...endpoint.Middleware) BodyEndpoints {
	return &bodyClientEndpoints{transports: transports, middlewares: middlewares}
}

type bodyFactories struct {
	transports BodyClientTransportsV2
}

func (f *bodyFactories) StarBody(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.StarBody(ctx, instance)
	}
}
func (f *bodyFactories) NamedBody(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.NamedBody(ctx, instance)
	}
}
func (f *bodyFactories) NonBody(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.NonBody(ctx, instance)
	}
}
func (f *bodyFactories) HttpBodyStarBody(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.HttpBodyStarBody(ctx, instance)
	}
}
func (f *bodyFactories) HttpBodyNamedBody(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.HttpBodyNamedBody(ctx, instance)
	}
}
func newBodyFactories(transports BodyClientTransportsV2) BodyFactories {
	return &bodyFactories{transports: transports}
}

type bodyEndpointers struct {
	target           string
	instancerFactory sdx.InstancerFactory
	factories        BodyFactories
	logger           log.Logger
	options          []sd.EndpointerOption
}

func (e *bodyEndpointers) StarBody(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.StarBody(ctx), e.logger, e.options...)
}
func (e *bodyEndpointers) NamedBody(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.NamedBody(ctx), e.logger, e.options...)
}
func (e *bodyEndpointers) NonBody(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.NonBody(ctx), e.logger, e.options...)
}
func (e *bodyEndpointers) HttpBodyStarBody(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.HttpBodyStarBody(ctx), e.logger, e.options...)
}
func (e *bodyEndpointers) HttpBodyNamedBody(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.HttpBodyNamedBody(ctx), e.logger, e.options...)
}
func newBodyEndpointers(
	target string,
	instancerFactory sdx.InstancerFactory,
	factories BodyFactories,
	logger log.Logger,
	options ...sd.EndpointerOption,
) BodyEndpointers {
	return &bodyEndpointers{
		target:           target,
		instancerFactory: instancerFactory,
		factories:        factories,
		logger:           logger,
		options:          options,
	}
}

type bodyBalancers struct {
	factory           lbx.BalancerFactory
	endpointer        BodyEndpointers
	starBody          lazyloadx.Group[lb.Balancer]
	namedBody         lazyloadx.Group[lb.Balancer]
	nonBody           lazyloadx.Group[lb.Balancer]
	httpBodyStarBody  lazyloadx.Group[lb.Balancer]
	httpBodyNamedBody lazyloadx.Group[lb.Balancer]
}

func (b *bodyBalancers) StarBody(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.starBody.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.StarBody))
	return balancer, err
}
func (b *bodyBalancers) NamedBody(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.namedBody.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.NamedBody))
	return balancer, err
}
func (b *bodyBalancers) NonBody(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.nonBody.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.NonBody))
	return balancer, err
}
func (b *bodyBalancers) HttpBodyStarBody(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.httpBodyStarBody.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.HttpBodyStarBody))
	return balancer, err
}
func (b *bodyBalancers) HttpBodyNamedBody(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.httpBodyNamedBody.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.HttpBodyNamedBody))
	return balancer, err
}
func newBodyBalancers(factory lbx.BalancerFactory, endpointer BodyEndpointers) BodyBalancers {
	return &bodyBalancers{
		factory:           factory,
		endpointer:        endpointer,
		starBody:          lazyloadx.Group[lb.Balancer]{},
		namedBody:         lazyloadx.Group[lb.Balancer]{},
		nonBody:           lazyloadx.Group[lb.Balancer]{},
		httpBodyStarBody:  lazyloadx.Group[lb.Balancer]{},
		httpBodyNamedBody: lazyloadx.Group[lb.Balancer]{},
	}
}
