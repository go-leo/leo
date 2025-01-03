// Code generated by protoc-gen-leo-core. DO NOT EDIT.

package helloworld

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
	statusx "github.com/go-leo/leo/v3/statusx"
	transportx "github.com/go-leo/leo/v3/transportx"
	io "io"
)

type GreeterService interface {
	// Query
	SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error)
}

type GreeterServerEndpoints interface {
	SayHello(ctx context.Context) endpoint.Endpoint
}

type GreeterClientEndpoints interface {
	SayHello(ctx context.Context) (endpoint.Endpoint, error)
}

type GreeterClientTransports interface {
	SayHello(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

type GreeterFactories interface {
	SayHello(ctx context.Context) sd.Factory
}

type GreeterEndpointers interface {
	SayHello(ctx context.Context, color string) (sd.Endpointer, error)
}

type GreeterBalancers interface {
	SayHello(ctx context.Context) (lb.Balancer, error)
}

type greeterServerEndpoints struct {
	svc         GreeterService
	middlewares []endpoint.Middleware
}

func (e *greeterServerEndpoints) SayHello(context.Context) endpoint.Endpoint {
	component := func(ctx context.Context, request any) (any, error) {
		return e.svc.SayHello(ctx, request.(*HelloRequest))
	}
	return endpointx.Chain(component, e.middlewares...)
}
func newGreeterServerEndpoints(svc GreeterService, middlewares ...endpoint.Middleware) GreeterServerEndpoints {
	return &greeterServerEndpoints{svc: svc, middlewares: middlewares}
}

type greeterClientEndpoints struct {
	balancers GreeterBalancers
}

func (e *greeterClientEndpoints) SayHello(ctx context.Context) (endpoint.Endpoint, error) {
	balancer, err := e.balancers.SayHello(ctx)
	if err != nil {
		return nil, err
	}
	return balancer.Endpoint()
}
func newGreeterClientEndpoints(
	target string,
	transports GreeterClientTransports,
	instancerFactory sdx.InstancerFactory,
	endpointerOptions []sd.EndpointerOption,
	balancerFactory lbx.BalancerFactory,
	logger log.Logger,
) GreeterClientEndpoints {
	factories := newGreeterFactories(transports)
	endpointers := newGreeterEndpointers(target, instancerFactory, factories, logger, endpointerOptions...)
	balancers := newGreeterBalancers(balancerFactory, endpointers)
	return &greeterClientEndpoints{balancers: balancers}
}

type greeterFactories struct {
	transports GreeterClientTransports
}

func (f *greeterFactories) SayHello(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.SayHello(ctx, instance)
	}
}
func newGreeterFactories(transports GreeterClientTransports) GreeterFactories {
	return &greeterFactories{transports: transports}
}

type greeterEndpointers struct {
	target           string
	instancerFactory sdx.InstancerFactory
	factories        GreeterFactories
	logger           log.Logger
	options          []sd.EndpointerOption
}

func (e *greeterEndpointers) SayHello(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.instancerFactory, e.factories.SayHello(ctx), e.logger, e.options...)
}
func newGreeterEndpointers(
	target string,
	instancerFactory sdx.InstancerFactory,
	factories GreeterFactories,
	logger log.Logger,
	options ...sd.EndpointerOption,
) GreeterEndpointers {
	return &greeterEndpointers{
		target:           target,
		instancerFactory: instancerFactory,
		factories:        factories,
		logger:           logger,
		options:          options,
	}
}

type greeterBalancers struct {
	factory    lbx.BalancerFactory
	endpointer GreeterEndpointers
	sayHello   lazyloadx.Group[lb.Balancer]
}

func (b *greeterBalancers) SayHello(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ExtractColor(ctx)
	balancer, err, _ := b.sayHello.LoadOrNew(color, lbx.NewBalancer(ctx, b.factory, b.endpointer.SayHello))
	return balancer, err
}
func newGreeterBalancers(factory lbx.BalancerFactory, endpointer GreeterEndpointers) GreeterBalancers {
	return &greeterBalancers{
		factory:    factory,
		endpointer: endpointer,
		sayHello:   lazyloadx.Group[lb.Balancer]{},
	}
}

type greeterClientService struct {
	endpoints     GreeterClientEndpoints
	transportName string
}

func (c *greeterClientService) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	ctx = endpointx.InjectName(ctx, "/helloworld.Greeter/SayHello")
	ctx = transportx.InjectName(ctx, c.transportName)
	endpoint, err := c.endpoints.SayHello(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, statusx.From(err)
	}
	return rep.(*HelloReply), nil
}
func newGreeterClientService(endpoints GreeterClientEndpoints, transportName string) GreeterService {
	return &greeterClientService{endpoints: endpoints, transportName: transportName}
}
