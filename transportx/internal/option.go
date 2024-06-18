package internal

import (
	"github.com/go-kit/kit/sd"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/dnssrvx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"google.golang.org/grpc"
	"time"
)

type ClientTransportOptions struct {
	InstancerBuilders []sdx.InstancerBuilder
	BalancerFactory   lbx.BalancerFactory
	EndpointerOptions []sd.EndpointerOption
	DefaultScheme     string
	dialOptions       []grpc.DialOption
}

func (o *ClientTransportOptions) Init() *ClientTransportOptions {
	o.InstancerBuilders = []sdx.InstancerBuilder{
		&passthroughx.InstancerBuilder{},
		&dnssrvx.InstancerBuilder{TTL: 30 * time.Second},
		&consulx.InstancerBuilder{},
	}
	o.BalancerFactory = lbx.RoundRobinFactory{}
	o.EndpointerOptions = nil
	o.DefaultScheme = "passthrough"
	return o
}

func (o *ClientTransportOptions) Apply(opts ...ClientTransportOption) *ClientTransportOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type ClientTransportOption func(o *ClientTransportOptions)

// InstancerBuilder is a option that sets the instancer builder.
func InstancerBuilder(builders ...sdx.InstancerBuilder) ClientTransportOption {
	return func(o *ClientTransportOptions) {
		o.InstancerBuilders = append(o.InstancerBuilders, builders...)
	}
}

// BalancerFactory is a option that sets the balancer factory.
func BalancerFactory(factory lbx.BalancerFactory) ClientTransportOption {
	return func(o *ClientTransportOptions) {
		o.BalancerFactory = factory
	}
}

// EndpointerOption is a option that sets the endpointer EndpointerOptions.
func EndpointerOption(options ...sd.EndpointerOption) ClientTransportOption {
	return func(o *ClientTransportOptions) {
		o.EndpointerOptions = append(o.EndpointerOptions, options...)
	}
}

// DefaultScheme is a option that sets the endpointer EndpointerOptions.
func DefaultScheme(scheme string) ClientTransportOption {
	return func(o *ClientTransportOptions) {
		o.DefaultScheme = scheme
	}
}

func DialOptions(dialOptions ...grpc.DialOption) ClientTransportOption {
	return func(o *ClientTransportOptions) {
		o.dialOptions = append(o.dialOptions, dialOptions...)
	}
}
