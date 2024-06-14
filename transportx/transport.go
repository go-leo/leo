package transportx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-leo/gox/backoff"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
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

	builders         []sdx.InstancerBuilder
	balancerFactory  lbx.BalancerFactory
	retryMax         int
	retryTimeout     time.Duration
	retryBackoffFunc backoff.BackoffFunc
	options          []sd.EndpointerOption

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

// Retry is a option that sets the retry options.
func Retry(max int, timeout time.Duration, backoff backoff.BackoffFunc) ClientTransportOption {
	return func(c *clientTransport) {
		c.retryMax = max
		c.retryTimeout = timeout
		c.retryBackoffFunc = backoff
	}
}

// EndpointerOption is a option that sets the endpointer options.
func EndpointerOption(options ...sd.EndpointerOption) ClientTransportOption {
	return func(c *clientTransport) {
		c.options = append(c.options, options...)
	}
}

func NewClientTransport(target string, factory sd.Factory, opts ...ClientTransportOption) (ClientTransport, error) {
	parsedTarget, err := sdx.ParseTarget(target)
	if err != nil {
		return nil, err
	}
	c := &clientTransport{
		target:           parsedTarget,
		factory:          factory,
		builders:         nil,
		balancerFactory:  lbx.RoundRobinFactory{},
		retryMax:         1,
		retryTimeout:     1 * time.Second,
		retryBackoffFunc: backoff.Zero(),
		options:          nil,
		clients:          sync.Map{},
		sfg:              singleflight.Group{},
	}
	for _, opt := range opts {
		opt(c)
	}
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
		builder := c.getInstancerBuilder(c.target.URL.Scheme)
		if builder == nil {
			return nil, statusx.ErrUnimplemented.WithMessage(fmt.Sprintf("could not get instancer builder for scheme: %q", c.target.URL.Scheme))
		}
		instancer, err := builder.Build(ctx, c.target, color)
		if err != nil {
			return nil, statusx.ErrFailedPrecondition.WithMessage(err.Error())
		}
		balancer := c.balancerFactory.New(ctx, sd.NewEndpointer(instancer, c.factory, logx.FromContext(ctx), c.options...))
		callback := func(n int, _ error) (bool, error) {
			time.Sleep(c.retryBackoffFunc(ctx, uint(n)))
			return n < c.retryMax, nil
		}
		return lb.RetryWithCallback(c.retryTimeout, balancer, callback), nil
	})
	result := <-resC
	if result.Err != nil {
		return endpointx.Error(result.Err)
	}
	if !result.Shared {
		c.clients.Store(key, result.Val)
	}
	return result.Val.(endpoint.Endpoint)
}

func (c *clientTransport) getInstancerBuilder(scheme string) sdx.InstancerBuilder {
	if scheme == "" {
		return &passthroughx.InstancerBuilder{}
	}
	for _, builder := range c.builders {
		if scheme == builder.Scheme() {
			return builder
		}
	}
	return nil
}
