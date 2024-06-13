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

type clientTransport struct {
	target           *sdx.Target
	factory          sd.Factory
	builders         []sdx.InstancerBuilder
	balancerFactory  lbx.BalancerFactory
	retryMax         int
	retryTimeout     time.Duration
	retryBackoffFunc backoff.BackoffFunc
	options          []sd.EndpointerOption
	clients          sync.Map
	sfg              singleflight.Group
}

func NewClient(target string, factory sd.Factory) (ClientTransport, error) {
	parsedTarget, err := sdx.ParseTarget(target)
	if err != nil {
		return nil, err
	}
	c := &clientTransport{
		target:           parsedTarget,
		factory:          factory,
		builders:         nil,
		balancerFactory:  nil,
		retryMax:         0,
		retryTimeout:     0,
		retryBackoffFunc: nil,
		options:          nil,
		clients:          sync.Map{},
		sfg:              singleflight.Group{},
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
	v, err, _ := c.sfg.Do(key, func() (interface{}, error) {
		var instancer sd.Instancer
		builder := c.getInstancerBuilder(c.target.URL.Scheme)
		if builder == nil {
			return nil, statusx.ErrUnimplemented.WithMessage(fmt.Sprintf("could not get instancer builder for scheme: %q", c.target.URL.Scheme))
		}
		instancer, err := builder.Build(ctx, c.target, color)
		if err != nil {
			return nil, err
		}
		ep := lb.RetryWithCallback(
			c.retryTimeout,
			c.balancerFactory.New(ctx, sd.NewEndpointer(instancer, c.factory, logx.FromContext(ctx), c.options...)),
			func(n int, received error) (keepTrying bool, replacement error) {
				time.Sleep(c.retryBackoffFunc(context.TODO(), uint(n)))
				return n < c.retryMax, nil
			})
		c.clients.Store(key, ep)
		return ep, nil
	})
	if err != nil {
		return endpointx.Error(err)
	}
	return v.(endpoint.Endpoint)
}

func (c *clientTransport) getInstancerBuilder(scheme string) sdx.InstancerBuilder {
	if scheme == "" {
		return passthroughx.NewInstancerBuilder()
	}
	for _, builder := range c.builders {
		if scheme == builder.Scheme() {
			return builder
		}
	}
	return nil
}
