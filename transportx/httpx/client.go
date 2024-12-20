package httpx

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/dnssrvx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"time"
)

type ClientOptions interface {
	Scheme() string
	Target() string
	ClientOptions() []httptransport.ClientOption
	Middlewares() []endpoint.Middleware
	InstancerFactory() sdx.InstancerFactory
	EndpointerOptions() []sd.EndpointerOption
	Logger() log.Logger
	BalancerFactory() lbx.BalancerFactory
}

type clientOptions struct {
	scheme            string
	target            string
	clientOptions     []httptransport.ClientOption
	middlewares       []endpoint.Middleware
	instancerFactory  sdx.InstancerFactory
	endpointerOptions []sd.EndpointerOption
	logger            log.Logger
	balancerFactory   lbx.BalancerFactory
}

func (o *clientOptions) Scheme() string {
	//TODO implement me
	panic("implement me")
}

func (o *clientOptions) Target() string {
	//TODO implement me
	panic("implement me")
}

func (o *clientOptions) ClientOptions() []httptransport.ClientOption {
	//TODO implement me
	panic("implement me")
}

func (o *clientOptions) Middlewares() []endpoint.Middleware {
	//TODO implement me
	panic("implement me")
}

func (o *clientOptions) InstancerFactory() sdx.InstancerFactory {
	//TODO implement me
	panic("implement me")
}

func (o *clientOptions) EndpointerOptions() []sd.EndpointerOption {
	//TODO implement me
	panic("implement me")
}

func (o *clientOptions) Logger() log.Logger {
	//TODO implement me
	panic("implement me")
}

func (o *clientOptions) BalancerFactory() lbx.BalancerFactory {
	//TODO implement me
	panic("implement me")
}

func NewClientOptions(opts ...ClientTransportOption) ClientOptions {
	return &clientOptions{}
}

func (o *clientOptions) Init() *clientOptions {
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

func (o *clientOptions) Apply(opts ...ClientTransportOption) *clientOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type ClientTransportOption func(o *clientOptions)

// HttpScheme is a option that sets the http scheme, http or https, default http
func HttpScheme(scheme string) ClientTransportOption {
	return func(o *clientOptions) {
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

// InstancerBuilder is a option that sets the instancer builder.
func InstancerBuilder(builders ...sdx.InstancerBuilder) ClientTransportOption {
	return func(o *clientOptions) {
		o.InstancerBuilders = append(o.InstancerBuilders, builders...)
	}
}

// BalancerFactory is a option that sets the balancer factory.
func BalancerFactory(factory lbx.BalancerFactory) ClientTransportOption {
	return func(o *clientOptions) {
		o.BalancerFactory = factory
	}
}

// EndpointerOption is a option that sets the endpointer EndpointerOptions.
func EndpointerOption(options ...sd.EndpointerOption) ClientTransportOption {
	return func(o *clientOptions) {
		o.EndpointerOptions = append(o.EndpointerOptions, options...)
	}
}

// DefaultScheme is a option that sets the endpointer EndpointerOptions.
func DefaultScheme(scheme string) ClientTransportOption {
	return func(o *clientOptions) {
		o.DefaultScheme = scheme
	}
}

func Middleware(middlewares ...endpoint.Middleware) ClientTransportOption {
	return func(o *clientOptions) {
		o.Middlewares = append(o.Middlewares, middlewares...)
	}
}
