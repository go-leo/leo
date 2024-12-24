// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package path

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
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	io "io"
)

type NamedPathService interface {
	NamedPathString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
	NamedPathOptString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
	NamedPathWrapString(ctx context.Context, request *NamedPathRequest) (*emptypb.Empty, error)
	EmbedNamedPathString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
	EmbedNamedPathOptString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
	EmbedNamedPathWrapString(ctx context.Context, request *EmbedNamedPathRequest) (*emptypb.Empty, error)
}

type NamedPathEndpoints interface {
	NamedPathString(ctx context.Context) endpoint.Endpoint
	NamedPathOptString(ctx context.Context) endpoint.Endpoint
	NamedPathWrapString(ctx context.Context) endpoint.Endpoint
	EmbedNamedPathString(ctx context.Context) endpoint.Endpoint
	EmbedNamedPathOptString(ctx context.Context) endpoint.Endpoint
	EmbedNamedPathWrapString(ctx context.Context) endpoint.Endpoint
}

type NamedPathClientEndpoints interface {
	NamedPathString(ctx context.Context) (endpoint.Endpoint, error)
	NamedPathOptString(ctx context.Context) (endpoint.Endpoint, error)
	NamedPathWrapString(ctx context.Context) (endpoint.Endpoint, error)
	EmbedNamedPathString(ctx context.Context) (endpoint.Endpoint, error)
	EmbedNamedPathOptString(ctx context.Context) (endpoint.Endpoint, error)
	EmbedNamedPathWrapString(ctx context.Context) (endpoint.Endpoint, error)
}

type NamedPathClientTransports interface {
	NamedPathString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	NamedPathOptString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	NamedPathWrapString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	EmbedNamedPathString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	EmbedNamedPathOptString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
	EmbedNamedPathWrapString(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

type NamedPathFactories interface {
	NamedPathString(ctx context.Context) sd.Factory
	NamedPathOptString(ctx context.Context) sd.Factory
	NamedPathWrapString(ctx context.Context) sd.Factory
	EmbedNamedPathString(ctx context.Context) sd.Factory
	EmbedNamedPathOptString(ctx context.Context) sd.Factory
	EmbedNamedPathWrapString(ctx context.Context) sd.Factory
}

type NamedPathEndpointers interface {
	NamedPathString(ctx context.Context, color string) (sd.Endpointer, error)
	NamedPathOptString(ctx context.Context, color string) (sd.Endpointer, error)
	NamedPathWrapString(ctx context.Context, color string) (sd.Endpointer, error)
	EmbedNamedPathString(ctx context.Context, color string) (sd.Endpointer, error)
	EmbedNamedPathOptString(ctx context.Context, color string) (sd.Endpointer, error)
	EmbedNamedPathWrapString(ctx context.Context, color string) (sd.Endpointer, error)
}

type NamedPathBalancers interface {
	NamedPathString(ctx context.Context) (lb.Balancer, error)
	NamedPathOptString(ctx context.Context) (lb.Balancer, error)
	NamedPathWrapString(ctx context.Context) (lb.Balancer, error)
	EmbedNamedPathString(ctx context.Context) (lb.Balancer, error)
	EmbedNamedPathOptString(ctx context.Context) (lb.Balancer, error)
	EmbedNamedPathWrapString(ctx context.Context) (lb.Balancer, error)
}

type namedPathServerEndpoints struct {
	svc         NamedPathService
	middlewares []endpoint.Middleware
}

func (e *namedPathServerEndpoints) NamedPathString(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedPathString(ctx, request.(*NamedPathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *namedPathServerEndpoints) NamedPathOptString(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedPathOptString(ctx, request.(*NamedPathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *namedPathServerEndpoints) NamedPathWrapString(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.NamedPathWrapString(ctx, request.(*NamedPathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *namedPathServerEndpoints) EmbedNamedPathString(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.EmbedNamedPathString(ctx, request.(*EmbedNamedPathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *namedPathServerEndpoints) EmbedNamedPathOptString(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.EmbedNamedPathOptString(ctx, request.(*EmbedNamedPathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func (e *namedPathServerEndpoints) EmbedNamedPathWrapString(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.EmbedNamedPathWrapString(ctx, request.(*EmbedNamedPathRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func newNamedPathServerEndpoints(svc NamedPathService, middlewares ...endpoint.Middleware) NamedPathEndpoints {
	return &namedPathServerEndpoints{svc: svc, middlewares: middlewares}
}

type namedPathClientEndpoints struct {
	balancers NamedPathBalancers
}

func (e *namedPathClientEndpoints) NamedPathString(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.NamedPathString(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *namedPathClientEndpoints) NamedPathOptString(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.NamedPathOptString(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *namedPathClientEndpoints) NamedPathWrapString(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.NamedPathWrapString(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *namedPathClientEndpoints) EmbedNamedPathString(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.EmbedNamedPathString(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *namedPathClientEndpoints) EmbedNamedPathOptString(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.EmbedNamedPathOptString(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func (e *namedPathClientEndpoints) EmbedNamedPathWrapString(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.EmbedNamedPathWrapString(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func newNamedPathClientEndpoints(
	target string,
	transports NamedPathClientTransports,
	instancerFactory sdx.InstancerFactory,
	endpointerOptions []sd.EndpointerOption,
	balancerFactory lbx.BalancerFactory,
	logger log.Logger,
) NamedPathClientEndpoints {
	factories := newNamedPathFactories(transports)
	endpointers := newNamedPathEndpointers(target, instancerFactory, factories, logger, endpointerOptions...)
	balancers := newNamedPathBalancers(balancerFactory, endpointers)
	return &namedPathClientEndpoints{balancers: balancers}
}

type namedPathFactories struct {
	transports NamedPathClientTransports
}

func (f *namedPathFactories) NamedPathString(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.NamedPathString(ctx, instance)
	}
}
func (f *namedPathFactories) NamedPathOptString(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.NamedPathOptString(ctx, instance)
	}
}
func (f *namedPathFactories) NamedPathWrapString(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.NamedPathWrapString(ctx, instance)
	}
}
func (f *namedPathFactories) EmbedNamedPathString(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.EmbedNamedPathString(ctx, instance)
	}
}
func (f *namedPathFactories) EmbedNamedPathOptString(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.EmbedNamedPathOptString(ctx, instance)
	}
}
func (f *namedPathFactories) EmbedNamedPathWrapString(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.EmbedNamedPathWrapString(ctx, instance)
	}
}
func newNamedPathFactories(transports NamedPathClientTransports) NamedPathFactories {
	return &namedPathFactories{transports: transports}
}

type namedPathEndpointers struct {
	target           string
	instancerFactory sdx.InstancerFactory
	factories        NamedPathFactories
	logger           log.Logger
	options          []sd.EndpointerOption
}

func (e *namedPathEndpointers) NamedPathString(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.NamedPathString(ctx), e.logger, e.options...)
}
func (e *namedPathEndpointers) NamedPathOptString(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.NamedPathOptString(ctx), e.logger, e.options...)
}
func (e *namedPathEndpointers) NamedPathWrapString(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.NamedPathWrapString(ctx), e.logger, e.options...)
}
func (e *namedPathEndpointers) EmbedNamedPathString(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.EmbedNamedPathString(ctx), e.logger, e.options...)
}
func (e *namedPathEndpointers) EmbedNamedPathOptString(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.EmbedNamedPathOptString(ctx), e.logger, e.options...)
}
func (e *namedPathEndpointers) EmbedNamedPathWrapString(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.EmbedNamedPathWrapString(ctx), e.logger, e.options...)
}
func newNamedPathEndpointers(
	target string,
	instancerFactory sdx.InstancerFactory,
	factories NamedPathFactories,
	logger log.Logger,
	options ...sd.EndpointerOption,
) NamedPathEndpointers {
	return &namedPathEndpointers{
		target:           target,
		instancerFactory: instancerFactory,
		factories:        factories,
		logger:           logger,
		options:          options,
	}
}

type namedPathBalancers struct {
	factory                  lbx.BalancerFactory
	endpointer               NamedPathEndpointers
	namedPathString          lazyloadx.Group[lb.Balancer]
	namedPathOptString       lazyloadx.Group[lb.Balancer]
	namedPathWrapString      lazyloadx.Group[lb.Balancer]
	embedNamedPathString     lazyloadx.Group[lb.Balancer]
	embedNamedPathOptString  lazyloadx.Group[lb.Balancer]
	embedNamedPathWrapString lazyloadx.Group[lb.Balancer]
}

func (b *namedPathBalancers) NamedPathString(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.namedPathString.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.NamedPathString))
	return balancer, err
}
func (b *namedPathBalancers) NamedPathOptString(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.namedPathOptString.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.NamedPathOptString))
	return balancer, err
}
func (b *namedPathBalancers) NamedPathWrapString(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.namedPathWrapString.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.NamedPathWrapString))
	return balancer, err
}
func (b *namedPathBalancers) EmbedNamedPathString(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.embedNamedPathString.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.EmbedNamedPathString))
	return balancer, err
}
func (b *namedPathBalancers) EmbedNamedPathOptString(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.embedNamedPathOptString.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.EmbedNamedPathOptString))
	return balancer, err
}
func (b *namedPathBalancers) EmbedNamedPathWrapString(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.embedNamedPathWrapString.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.EmbedNamedPathWrapString))
	return balancer, err
}
func newNamedPathBalancers(factory lbx.BalancerFactory, endpointer NamedPathEndpointers) NamedPathBalancers {
	return &namedPathBalancers{
		factory:                  factory,
		endpointer:               endpointer,
		namedPathString:          lazyloadx.Group[lb.Balancer]{},
		namedPathOptString:       lazyloadx.Group[lb.Balancer]{},
		namedPathWrapString:      lazyloadx.Group[lb.Balancer]{},
		embedNamedPathString:     lazyloadx.Group[lb.Balancer]{},
		embedNamedPathOptString:  lazyloadx.Group[lb.Balancer]{},
		embedNamedPathWrapString: lazyloadx.Group[lb.Balancer]{},
	}
}
