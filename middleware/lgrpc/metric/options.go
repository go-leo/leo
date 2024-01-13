package metric

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

type options struct {
	// Skips is a grpc full method array or url path array which metrics are collected.
	Skips            []string
	BucketBoundaries []float64
	MeterProvider    metric.MeterProvider
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func (o *options) init() {
	if o.MeterProvider == nil {
		o.MeterProvider = otel.GetMeterProvider()
	}
}

func WithSkips(skips ...string) Option {
	return func(o *options) {
		o.Skips = append(o.Skips, skips...)
	}
}

func WithBucketBoundaries(bounds ...float64) Option {
	return func(o *options) {
		o.BucketBoundaries = bounds
	}
}

func WithMeterProvider(p metric.MeterProvider) Option {
	return func(o *options) {
		o.MeterProvider = p
	}
}
