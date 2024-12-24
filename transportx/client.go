package transportx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/dnssrvx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/stainx"
	"github.com/go-leo/leo/v3/statusx"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
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
	factory sdx.Factory
	builder sdx.InstancerBuilder

	options *clientTransportOptions

	clients lazyloadx.Group[lb.Balancer]
	sfg     singleflight.Group
}

type clientTransportOptions struct {
	InstancerBuilders []sdx.InstancerBuilder
	BalancerFactory   lbx.BalancerFactory
	EndpointerOptions []sd.EndpointerOption
	DefaultScheme     string
	FactoryArgs       any
}

func (o *clientTransportOptions) Init() *clientTransportOptions {
	o.InstancerBuilders = []sdx.InstancerBuilder{
		&dnssrvx.InstancerBuilder{TTL: 30 * time.Second},
		&consulx.InstancerBuilder{},
	}
	o.BalancerFactory = lbx.RoundRobinFactory{}
	o.EndpointerOptions = nil
	o.DefaultScheme = "passthrough"
	return o
}

func (o *clientTransportOptions) Apply(opts ...ClientTransportOption) *clientTransportOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type ClientTransportOption func(o *clientTransportOptions)

// HttpScheme is a option that sets the http scheme, http or https, default http
func HttpScheme(scheme string) ClientTransportOption {
	return func(o *clientTransportOptions) {
		if o.FactoryArgs == nil {
			o.FactoryArgs = scheme
			return
		}
		_, ok := o.FactoryArgs.(string)
		if !ok {
			panic("factory args have already been set")
		}
		o.FactoryArgs = scheme
	}
}

// GrpcDialOption is a option that sets the grpc dial options.
func GrpcDialOption(options ...grpc.DialOption) ClientTransportOption {
	return func(o *clientTransportOptions) {
		if o.FactoryArgs == nil {
			dialOptions := make([]grpc.DialOption, 0, len(options))
			dialOptions = append(dialOptions, options...)
			o.FactoryArgs = dialOptions
			return
		}
		dialOptions, ok := o.FactoryArgs.([]grpc.DialOption)
		if !ok {
			panic("factory args have already been set")
		}
		dialOptions = append(dialOptions, options...)
		o.FactoryArgs = dialOptions
	}
}

// InstancerBuilder is a option that sets the instancer builder.
func InstancerBuilder(builders ...sdx.InstancerBuilder) ClientTransportOption {
	return func(o *clientTransportOptions) {
		o.InstancerBuilders = append(o.InstancerBuilders, builders...)
	}
}

// BalancerFactory is a option that sets the balancer factory.
func BalancerFactory(factory lbx.BalancerFactory) ClientTransportOption {
	return func(o *clientTransportOptions) {
		o.BalancerFactory = factory
	}
}

// EndpointerOption is a option that sets the endpointer EndpointerOptions.
func EndpointerOption(options ...sd.EndpointerOption) ClientTransportOption {
	return func(o *clientTransportOptions) {
		o.EndpointerOptions = append(o.EndpointerOptions, options...)
	}
}

// DefaultScheme is a option that sets the endpointer EndpointerOptions.
func DefaultScheme(scheme string) ClientTransportOption {
	return func(o *clientTransportOptions) {
		o.DefaultScheme = scheme
	}
}

func NewClientTransport(target string, factory sdx.Factory, opts ...ClientTransportOption) (ClientTransport, error) {
	c := &clientTransport{
		target:  nil,
		factory: factory,
		builder: nil,
		options: new(clientTransportOptions).Init().Apply(opts...),
		clients: lazyloadx.Group[lb.Balancer]{},
		sfg:     singleflight.Group{},
	}
	parsedTarget, err := sdx.ParseTarget(target)
	if err == nil {
		c.target = parsedTarget
		builder := c.getInstancerBuilder(c.target.URL.Scheme)
		if builder == nil {
			return nil, statusx.ErrUnimplemented.With(statusx.Message("could not get instancer builder for scheme: %q", c.target.URL.Scheme))
		}
		c.builder = builder
		return c, nil
	}

	canonicalTarget := c.options.DefaultScheme + ":///" + target
	parsedTarget, err = sdx.ParseTarget(canonicalTarget)
	if err != nil {
		return nil, statusx.ErrUnimplemented.With(statusx.Message("could not parse canonical target: %q", canonicalTarget))
	}
	c.target = parsedTarget
	builder := c.getInstancerBuilder(c.target.URL.Scheme)
	if builder == nil {
		return nil, statusx.ErrUnimplemented.With(statusx.Message("could not get instancer builder for default scheme: %q", parsedTarget.URL.Scheme))
	}
	c.builder = builder
	return c, nil
}

func (c *clientTransport) Endpoint(ctx context.Context) endpoint.Endpoint {
	balancer, err := c.balancer(ctx)
	if err != nil {
		return endpointx.Error(statusx.ErrFailedPrecondition.With(statusx.Wrap(err)))
	}
	ep, err := balancer.Endpoint()
	if err != nil {
		return endpointx.Error(statusx.ErrFailedPrecondition.With(statusx.Wrap(err)))
	}
	return ep
}

func (c *clientTransport) balancer(ctx context.Context) (lb.Balancer, error) {
	var key string
	if color, ok := stainx.ExtractColor(ctx); ok {
		key = color
	}
	value, err, _ := c.clients.LoadOrNew(key, func(key string) (lb.Balancer, error) {
		instancer, err := c.builder.Build(ctx, c.target)
		if err != nil {
			return nil, err
		}
		factory := c.factory(ctx, c.options.FactoryArgs)
		endpointer := sd.NewEndpointer(instancer, factory, logx.FromContext(ctx), c.options.EndpointerOptions...)
		balancer := c.options.BalancerFactory.New(ctx, endpointer)
		return balancer, nil
	})
	if err != nil {
		return nil, err
	}
	return value.(lb.Balancer), nil
}

// getInstancerBuilder gets the instancer builder for the given scheme.
// If the scheme is an empty string (""), the function returns an instance of passthroughx.InstancerBuilder.
// Otherwise, it iterates through a slice c.InstancerBuilders in reverse order (from the last element to the first).
// This slice presumably contains instances of structs that implement sdx.InstancerBuilder interface,
// each with a method Scheme() that returns a string representing the scheme it supports.
func (c *clientTransport) getInstancerBuilder(scheme string) sdx.InstancerBuilder {
	if scheme == "" {
		return nil
	}
	for i := len(c.options.InstancerBuilders) - 1; i >= 0; i-- {
		if scheme == c.options.InstancerBuilders[i].Scheme() {
			return c.options.InstancerBuilders[i]
		}
	}
	return nil
}
