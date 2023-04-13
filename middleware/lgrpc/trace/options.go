package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type Skip func(ctx context.Context, method string) bool

type options struct {
	Propagators    propagation.TextMapPropagator
	TracerProvider trace.TracerProvider
	Skips          []Skip
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Propagators == nil {
		o.Propagators = otel.GetTextMapPropagator()
	}
	if o.TracerProvider == nil {
		o.TracerProvider = otel.GetTracerProvider()
	}
}

type Option func(o *options)

func Propagators(propagators propagation.TextMapPropagator) Option {
	return func(o *options) {
		o.Propagators = propagators
	}
}

func TracerProvider(tracerProvider trace.TracerProvider) Option {
	return func(o *options) {
		o.TracerProvider = tracerProvider
	}
}

func Skips(skips ...Skip) Option {
	return func(o *options) {
		o.Skips = append(o.Skips, skips...)
	}
}
