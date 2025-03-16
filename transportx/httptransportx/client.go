package httptransportx

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/fixed"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
)

type (
	ClientOptions interface {
		// Scheme returns the http scheme, http or https, default http
		Scheme() string
		// ClientTransportOptions returns the go-kit http transport client options.
		ClientTransportOptions() []httptransport.ClientOption
		// Builder returns the sdx.Builder.
		Builder() sdx.Builder
		// EndpointerOptions returns the sd.EndpointerOptions.
		EndpointerOptions() []sd.EndpointerOption
		// Logger returns the logger.
		Logger() log.Logger
		// BalancerFactory returns the lbx.BalancerFactory.
		BalancerFactory() lbx.BalancerFactory
		// UnmarshalOptions returns the protojson.UnmarshalOptions.
		UnmarshalOptions() protojson.UnmarshalOptions
		// MarshalOptions returns the protojson.MarshalOptions.
		MarshalOptions() protojson.MarshalOptions
		// Middlewares returns the go-kit endpoint middlewares.
		Middlewares() []endpoint.Middleware
	}

	clientOptions struct {
		scheme                 string
		clientTransportOptions []httptransport.ClientOption
		builder                sdx.Builder
		endpointerOptions      []sd.EndpointerOption
		logger                 log.Logger
		balancerFactory        lbx.BalancerFactory
		unmarshalOptions       protojson.UnmarshalOptions
		marshalOptions         protojson.MarshalOptions
		middlewares            []endpoint.Middleware
	}

	ClientOption func(o *clientOptions)
)

func (o *clientOptions) Apply(opts ...ClientOption) *clientOptions {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *clientOptions) Scheme() string {
	return o.scheme
}

func (o *clientOptions) ClientTransportOptions() []httptransport.ClientOption {
	return o.clientTransportOptions
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

func (o *clientOptions) UnmarshalOptions() protojson.UnmarshalOptions {
	return o.unmarshalOptions
}

func (o *clientOptions) MarshalOptions() protojson.MarshalOptions {
	return o.marshalOptions
}

func (o *clientOptions) Middlewares() []endpoint.Middleware {
	return o.middlewares
}

// WithScheme is a option that sets the http scheme, http or https, default http
func WithScheme(scheme string) ClientOption {
	return func(o *clientOptions) {
		o.scheme = scheme
	}
}

// WithClientTransportOption is a option that sets the go-kit http transport client options.
func WithClientTransportOption(options ...httptransport.ClientOption) ClientOption {
	return func(o *clientOptions) {
		o.clientTransportOptions = append(o.clientTransportOptions, options...)
	}
}

// WithInstancerBuilder is a option that sets the sd instancer factory.
func WithInstancerBuilder(builder sdx.Builder) ClientOption {
	return func(o *clientOptions) {
		o.builder = builder
	}
}

// WithEndpointerOption is a option that sets the endpointer EndpointerOptions.
func WithEndpointerOption(options ...sd.EndpointerOption) ClientOption {
	return func(o *clientOptions) {
		o.endpointerOptions = append(o.endpointerOptions, options...)
	}
}

// WithLogger is a option that sets the logger.
func Logger(logger log.Logger) ClientOption {
	return func(o *clientOptions) {
		o.logger = logger
	}
}

// WithBalancerFactory is a option that sets the balancer factory.
func WithBalancerFactory(factory lbx.BalancerFactory) ClientOption {
	return func(o *clientOptions) {
		o.balancerFactory = factory
	}
}

// WithUnmarshalOptions is a option that sets the protojson.UnmarshalOptions.
func WithUnmarshalOptions(opts protojson.UnmarshalOptions) ClientOption {
	return func(o *clientOptions) {
		o.unmarshalOptions = opts
	}
}

// WithMarshalOptions is a option that sets the protojson.MarshalOptions.
func WithMarshalOptions(opts protojson.MarshalOptions) ClientOption {
	return func(o *clientOptions) {
		o.marshalOptions = opts
	}
}

// WithMiddleware is a option that sets the go-kit endpoint middlewares.
func WithMiddleware(middlewares ...endpoint.Middleware) ClientOption {
	return func(o *clientOptions) {
		o.middlewares = append(o.middlewares, middlewares...)
	}
}

func NewClientOptions(opts ...ClientOption) ClientOptions {
	options := &clientOptions{
		scheme:                 "http",
		clientTransportOptions: nil,
		builder:                fixed.Builder{},
		endpointerOptions:      nil,
		logger:                 logx.New(os.Stdout, logx.JSON(), logx.Timestamp(), logx.Caller(0), logx.Sync()),
		balancerFactory:        lbx.PeakFirstFactory{},
		unmarshalOptions:       protojson.UnmarshalOptions{},
		marshalOptions:         protojson.MarshalOptions{},
		middlewares:            nil,
	}
	return options.Apply(opts...)
}
