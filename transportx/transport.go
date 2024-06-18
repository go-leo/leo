package transportx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx/internal"
	"golang.org/x/sync/singleflight"
	"sync"
)

// ClientTransport is a transport that can be used to invoke a remote endpoint.
type ClientTransport interface {
	// Endpoint returns a usable endpoint that invokes the remote endpoint.
	Endpoint(ctx context.Context) endpoint.Endpoint
}

// ClientTransport is a transport that can be used to invoke a remote endpoint.
type clientTransport struct {
	target  *sdx.Target
	factory sd.Factory
	builder sdx.InstancerBuilder

	options *internal.ClientTransportOptions

	clients sync.Map
	sfg     singleflight.Group
}

type ClientTransportOption = internal.ClientTransportOption

// InstancerBuilder is a option that sets the instancer builder.
func InstancerBuilder(builders ...sdx.InstancerBuilder) ClientTransportOption {
	return internal.InstancerBuilder(builders...)
}

// BalancerFactory is a option that sets the balancer factory.
func BalancerFactory(factory lbx.BalancerFactory) ClientTransportOption {
	return internal.BalancerFactory(factory)
}

// EndpointerOption is a option that sets the endpointer EndpointerOptions.
func EndpointerOption(options ...sd.EndpointerOption) ClientTransportOption {
	return internal.EndpointerOption(options...)
}

// DefaultScheme is a option that sets the endpointer EndpointerOptions.
func DefaultScheme(scheme string) ClientTransportOption {
	return internal.DefaultScheme(scheme)
}

func NewClientTransport(target string, factory sd.Factory, opts ...ClientTransportOption) (ClientTransport, error) {
	c := &clientTransport{
		target:  nil,
		factory: factory,
		builder: nil,
		options: new(internal.ClientTransportOptions).Init().Apply(opts...),
		clients: sync.Map{},
		sfg:     singleflight.Group{},
	}
	parsedTarget, err := sdx.ParseTarget(target)
	if err == nil {
		c.target = parsedTarget
		builder := c.getInstancerBuilder(c.target.URL.Scheme)
		if builder == nil {
			return nil, statusx.ErrUnimplemented.WithMessage(fmt.Sprintf("could not get instancer builder for scheme: %q", c.target.URL.Scheme))
		}
		c.builder = builder
		return c, nil
	}

	canonicalTarget := c.options.DefaultScheme + ":///" + target
	parsedTarget, err = sdx.ParseTarget(canonicalTarget)
	if err != nil {
		return nil, statusx.ErrUnimplemented.WithMessage(fmt.Sprintf("could not parse canonical target: %q", canonicalTarget))
	}
	c.target = parsedTarget
	builder := c.getInstancerBuilder(c.target.URL.Scheme)
	if builder == nil {
		return nil, statusx.ErrUnimplemented.WithMessage(fmt.Sprintf("could not get instancer builder for default scheme: %q", parsedTarget.URL.Scheme))
	}
	c.builder = builder
	return c, nil
}

func (c *clientTransport) Endpoint(ctx context.Context) endpoint.Endpoint {
	balancer, err := c.balancer(ctx)
	if err != nil {
		return endpointx.Error(statusx.ErrFailedPrecondition.WithMessage(err.Error()))
	}
	ep, err := balancer.Endpoint()
	if err != nil {
		return endpointx.Error(statusx.ErrFailedPrecondition.WithMessage(err.Error()))
	}
	return ep
}

func (c *clientTransport) balancer(ctx context.Context) (lb.Balancer, error) {
	var key string
	color, ok := sdx.ExtractColor(ctx)
	if ok {
		key = string(color)
	}
	value, ok := c.clients.Load(key)
	if ok {
		return value.(lb.Balancer), nil
	}
	resC := c.sfg.DoChan(key, func() (interface{}, error) {
		instancer, err := c.builder.Build(ctx, c.target)
		if err != nil {
			return nil, err
		}
		balancer := c.options.BalancerFactory.New(ctx, sd.NewEndpointer(instancer, c.factory, logx.FromContext(ctx), c.options.EndpointerOptions...))
		return balancer, nil
	})
	result := <-resC
	if result.Err != nil {
		return nil, result.Err
	}
	if !result.Shared {
		c.clients.Store(key, result.Val)
	}
	return result.Val.(lb.Balancer), nil
}

// getInstancerBuilder gets the instancer builder for the given scheme.
// If the scheme is an empty string (""), the function returns an instance of passthroughx.InstancerBuilder.
// Otherwise, it iterates through a slice c.InstancerBuilders in reverse order (from the last element to the first).
// This slice presumably contains instances of structs that implement sdx.InstancerBuilder interface,
// each with a method Scheme() that returns a string representing the scheme it supports.
func (c *clientTransport) getInstancerBuilder(scheme string) sdx.InstancerBuilder {
	if scheme == "" {
		return &passthroughx.InstancerBuilder{}
	}
	for i := len(c.options.InstancerBuilders) - 1; i >= 0; i-- {
		if scheme == c.options.InstancerBuilders[i].Scheme() {
			return c.options.InstancerBuilders[i]
		}
	}
	return nil
}
