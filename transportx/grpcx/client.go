package grpcx

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"google.golang.org/grpc"
)

type (
	ClientOptions interface {
		DialOptions() []grpc.DialOption
		ClientTransportOptions() []grpctransport.ClientOption
		Middlewares() []endpoint.Middleware
		Builder() sdx.Builder
		EndpointerOptions() []sd.EndpointerOption
		Logger() log.Logger
		BalancerFactory() lbx.BalancerFactory
	}

	clientOptions struct {
		dialOptions            []grpc.DialOption
		clientTransportOptions []grpctransport.ClientOption
		middlewares            []endpoint.Middleware
		builder                sdx.Builder
		endpointerOptions      []sd.EndpointerOption
		logger                 log.Logger
		balancerFactory        lbx.BalancerFactory
	}

	ClientOption func(o *clientOptions)
)

func (o *clientOptions) DialOptions() []grpc.DialOption {
	return o.dialOptions
}

func (o *clientOptions) ClientTransportOptions() []grpctransport.ClientOption {
	return o.clientTransportOptions
}

func (o *clientOptions) Middlewares() []endpoint.Middleware {
	return o.middlewares
}

func (o *clientOptions) Builder() sdx.Builder {
	return o.builder
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

// DialOptions is a option that sets the grpc dial options.
func DialOptions(options ...grpc.DialOption) ClientOption {
	return func(o *clientOptions) {
		o.dialOptions = append(o.dialOptions, options...)
	}
}

// ClientTransportOption is a option that sets the go-kit grpc transport client options.
func ClientTransportOption(options ...grpctransport.ClientOption) ClientOption {
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

// InstancerBuilder is a option that sets the sd instancer factory.
func InstancerBuilder(factory sdx.Builder) ClientOption {
	return func(o *clientOptions) {
		o.builder = factory
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
	var options = &clientOptions{
		dialOptions:            nil,
		clientTransportOptions: nil,
		middlewares:            nil,
		builder:                passthroughx.Builder{},
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
