package httpx

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
)

type (
	ClientOptions interface {
		Scheme() string
		ClientTransportOptions() []httptransport.ClientOption
		Middlewares() []endpoint.Middleware
		InstancerFactory() sdx.InstancerFactory
		EndpointerOptions() []sd.EndpointerOption
		Logger() log.Logger
		BalancerFactory() lbx.BalancerFactory
	}

	clientOptions struct {
		scheme                 string
		clientTransportOptions []httptransport.ClientOption
		middlewares            []endpoint.Middleware
		instancerFactory       sdx.InstancerFactory
		endpointerOptions      []sd.EndpointerOption
		logger                 log.Logger
		balancerFactory        lbx.BalancerFactory
	}

	ClientOption func(o *clientOptions)
)

func (o *clientOptions) Scheme() string {
	return o.scheme
}

func (o *clientOptions) ClientTransportOptions() []httptransport.ClientOption {
	return o.clientTransportOptions
}

func (o *clientOptions) Middlewares() []endpoint.Middleware {
	return o.middlewares
}

func (o *clientOptions) InstancerFactory() sdx.InstancerFactory {
	return o.instancerFactory
}

func (o *clientOptions) EndpointerOptions() []sd.EndpointerOption {
	return o.endpointerOptions
}

func (o *clientOptions) Logger() log.Logger {
	return o.logger
}

func (o *clientOptions) BalancerFactory() lbx.BalancerFactory {
	return o.balancerFactory
}

// Scheme is a option that sets the http scheme, http or https, default http
func Scheme(scheme string) ClientOption {
	return func(o *clientOptions) {
		o.scheme = scheme
	}
}

// ClientTransportOption is a option that sets the go-kit http transport client options.
func ClientTransportOption(options ...httptransport.ClientOption) ClientOption {
	return func(o *clientOptions) {
		o.clientTransportOptions = append(o.clientTransportOptions, options...)
	}
}

// Middleware is a option that sets the go-kit endpoint middlewares.
func Middleware(middlewares ...endpoint.Middleware) ClientOption {
	return func(o *clientOptions) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

// InstancerFactory is a option that sets the sd instancer factory.
func InstancerFactory(factory sdx.InstancerFactory) ClientOption {
	return func(o *clientOptions) {
		o.instancerFactory = factory
	}
}

// EndpointerOption is a option that sets the endpointer EndpointerOptions.
func EndpointerOption(options ...sd.EndpointerOption) ClientOption {
	return func(o *clientOptions) {
		o.endpointerOptions = append(o.endpointerOptions, options...)
	}
}

// Logger is a option that sets the logger.
func Logger(logger log.Logger) ClientOption {
	return func(o *clientOptions) {
		o.logger = logger
	}
}

// BalancerFactory is a option that sets the balancer factory.
func BalancerFactory(factory lbx.BalancerFactory) ClientOption {
	return func(o *clientOptions) {
		o.balancerFactory = factory
	}
}

func NewClientOptions(opts ...ClientOption) ClientOptions {
	options := &clientOptions{
		scheme:                 "http",
		clientTransportOptions: nil,
		middlewares:            nil,
		instancerFactory:       passthroughx.Factory{},
		endpointerOptions:      nil,
		logger:                 logx.L(),
		balancerFactory:        lbx.PeakFirstFactory{},
	}
	return options.Apply(opts...)
}

func (o *clientOptions) Apply(opts ...ClientOption) *clientOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}
