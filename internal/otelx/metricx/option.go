package metricx

import (
	"go.opentelemetry.io/otel/attribute"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"

	"github.com/go-leo/leo/internal/otelx/resourcex"
)

type ViewOption struct {
	Criteria sdkmetric.Instrument
	Mask     sdkmetric.Stream
}

type options struct {
	// Service
	Service *resourcex.Service
	// Attributes
	Attributes []attribute.KeyValue
	// Resources
	Resources resourcex.ResourceFlag
	// HTTPOptions
	HTTPOptions *HTTPOptions
	// GRPCOptions
	GRPCOptions *GRPCOptions
	// PrometheusOptions
	PrometheusOptions *PrometheusOptions
	// WriterOptions
	WriterOptions *WriterOptions
	// ViewOptions
	ViewOptions []ViewOption
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {

}

type Option func(o *options)

func Service(svc *resourcex.Service) Option {
	return func(o *options) {
		o.Service = svc
	}
}

func Attributes(attrs ...attribute.KeyValue) Option {
	return func(o *options) {
		o.Attributes = attrs
	}
}

func Resources(res resourcex.ResourceFlag) Option {
	return func(o *options) {
		o.Resources = res
	}
}

func GRPC(gRPCOptions *GRPCOptions) Option {
	return func(o *options) {
		o.GRPCOptions = gRPCOptions
	}
}

func HTTP(httpOptions *HTTPOptions) Option {
	return func(o *options) {
		o.HTTPOptions = httpOptions
	}
}

func Prometheus(prometheusOptions *PrometheusOptions) Option {
	return func(o *options) {
		o.PrometheusOptions = prometheusOptions
	}
}

func Writer(writerOptions *WriterOptions) Option {
	return func(o *options) {
		o.WriterOptions = writerOptions
	}
}

func View(views ...ViewOption) Option {
	return func(o *options) {
		o.ViewOptions = append(o.ViewOptions, views...)
	}
}
