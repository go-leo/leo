package transportx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/dnssrvx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"github.com/go-leo/leo/v3/statusx"
	"golang.org/x/sync/singleflight"
	"strings"
	"sync"
	"time"
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

	builders        []sdx.InstancerBuilder
	balancerFactory lbx.BalancerFactory
	options         []sd.EndpointerOption
	defaultScheme   string

	clients sync.Map
	sfg     singleflight.Group
}

type ClientTransportOption func(c *clientTransport)

// InstancerBuilder is a option that sets the instancer builder.
func InstancerBuilder(builders ...sdx.InstancerBuilder) ClientTransportOption {
	return func(c *clientTransport) {
		c.builders = append(c.builders, builders...)
	}
}

// BalancerFactory is a option that sets the balancer factory.
func BalancerFactory(factory lbx.BalancerFactory) ClientTransportOption {
	return func(c *clientTransport) {
		c.balancerFactory = factory
	}
}

// EndpointerOption is a option that sets the endpointer options.
func EndpointerOption(options ...sd.EndpointerOption) ClientTransportOption {
	return func(c *clientTransport) {
		c.options = append(c.options, options...)
	}
}

// DefaultScheme is a option that sets the endpointer options.
func DefaultScheme(scheme string) ClientTransportOption {
	return func(c *clientTransport) {
		c.defaultScheme = scheme
	}
}

func NewClientTransport(target string, factory sd.Factory, opts ...ClientTransportOption) (ClientTransport, error) {
	c := &clientTransport{
		target:  nil,
		factory: factory,
		builder: nil,
		builders: []sdx.InstancerBuilder{
			&passthroughx.InstancerBuilder{},
			&dnssrvx.InstancerBuilder{TTL: 30 * time.Second},
			&consulx.InstancerBuilder{},
		},
		balancerFactory: lbx.RoundRobinFactory{},
		options:         nil,
		defaultScheme:   "passthrough",
		clients:         sync.Map{},
		sfg:             singleflight.Group{},
	}
	for _, opt := range opts {
		opt(c)
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

	canonicalTarget := c.defaultScheme + ":///" + target
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
	colors, ok := sdx.ExtractColors(ctx)
	if !ok {
		return c.endpoint(ctx, nil)
	}
	color, ok := colors.Find(c.target.Instance())
	if !ok {
		return c.endpoint(ctx, nil)
	}
	return c.endpoint(ctx, color)
}

func (c *clientTransport) endpoint(ctx context.Context, color *sdx.Color) endpoint.Endpoint {
	key := strings.Join(color.Color(), ",")
	value, ok := c.clients.Load(key)
	if ok {
		return value.(endpoint.Endpoint)
	}
	resC := c.sfg.DoChan(key, func() (interface{}, error) {
		instancer, err := c.builder.Build(ctx, c.target, color)
		if err != nil {
			return nil, nil
		}
		balancer := c.balancerFactory.New(ctx, sd.NewEndpointer(instancer, c.factory, logx.FromContext(ctx), c.options...))
		return balancer.Endpoint()
	})
	result := <-resC
	if result.Err != nil {
		return endpointx.Error(statusx.ErrFailedPrecondition.WithMessage(result.Err.Error()))
	}
	if !result.Shared {
		c.clients.Store(key, result.Val)
	}
	return result.Val.(endpoint.Endpoint)
}

// getInstancerBuilder gets the instancer builder for the given scheme.
// If the scheme is an empty string (""), the function returns an instance of passthroughx.InstancerBuilder.
// Otherwise, it iterates through a slice c.builders in reverse order (from the last element to the first).
// This slice presumably contains instances of structs that implement sdx.InstancerBuilder interface,
// each with a method Scheme() that returns a string representing the scheme it supports.
func (c *clientTransport) getInstancerBuilder(scheme string) sdx.InstancerBuilder {
	if scheme == "" {
		return &passthroughx.InstancerBuilder{}
	}
	for i := len(c.builders) - 1; i >= 0; i-- {
		if scheme == c.builders[i].Scheme() {
			return c.builders[i]
		}
	}
	return nil
}
