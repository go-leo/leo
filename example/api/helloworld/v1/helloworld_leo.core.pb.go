// Code generated by protoc-gen-go-leo. DO NOT EDIT.

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
	stainx "github.com/go-leo/leo/v3/stainx"
	transportx "github.com/go-leo/leo/v3/transportx"
	io "io"
)

// GreeterService is a service
type GreeterService interface {
	SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error)
}

// GreeterServerEndpoints is server endpoints
type GreeterServerEndpoints interface {
	SayHello(ctx context.Context) endpoint.Endpoint
}

// GreeterClientEndpoints is client endpoints
type GreeterClientEndpoints interface {
	SayHello(ctx context.Context) (endpoint.Endpoint, error)
}

// GreeterClientTransports is client transports
type GreeterClientTransports interface {
	SayHello(ctx context.Context, instance string) (endpoint.Endpoint, io.Closer, error)
}

// GreeterFactories is client factories
type GreeterFactories interface {
	SayHello(ctx context.Context) sd.Factory
}

// GreeterEndpointers is client endpointers
type GreeterEndpointers interface {
	SayHello(ctx context.Context, color string) (sd.Endpointer, error)
}

// GreeterBalancers is client balancers
type GreeterBalancers interface {
	SayHello(ctx context.Context) (lb.Balancer, error)
}

// greeterServerEndpoints implements GreeterServerEndpoints
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

// greeterFactories implements GreeterFactories
type greeterFactories struct {
	transports GreeterClientTransports
}

func (f *greeterFactories) SayHello(ctx context.Context) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return f.transports.SayHello(ctx, instance)
	}
}

// greeterEndpointers implements GreeterEndpointers
type greeterEndpointers struct {
	target    string
	builder   sdx.Builder
	factories GreeterFactories
	logger    log.Logger
	options   []sd.EndpointerOption
}

func (e *greeterEndpointers) SayHello(ctx context.Context, color string) (sd.Endpointer, error) {
	return sdx.NewEndpointer(ctx, e.target, color, e.builder, e.factories.SayHello(ctx), e.logger, e.options...)
}

// greeterBalancers implements GreeterBalancers
type greeterBalancers struct {
	factory    lbx.BalancerFactory
	endpointer GreeterEndpointers
	sayHello   lazyloadx.Group[lb.Balancer]
}

func (b *greeterBalancers) SayHello(ctx context.Context) (lb.Balancer, error) {
	color, _ := stainx.ColorExtractor(ctx)
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

// greeterClientEndpoints implements GreeterClientEndpoints
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

// greeterClientService implements GreeterClientService
type greeterClientService struct {
	endpoints     GreeterClientEndpoints
	transportName string
}

func (c *greeterClientService) SayHello(ctx context.Context, request *HelloRequest) (*HelloReply, error) {
	ctx = endpointx.NameInjector(ctx, "/helloworld.Greeter/SayHello")
	ctx = transportx.NameInjector(ctx, c.transportName)
	endpoint, err := c.endpoints.SayHello(ctx)
	if err != nil {
		return nil, err
	}
	rep, err := endpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	return rep.(*HelloReply), nil
}
