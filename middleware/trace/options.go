package trace

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type options struct {
	Propagators    propagation.TextMapPropagator
	TracerProvider trace.TracerProvider
	Skips          map[string]struct{}
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
	if o.Skips == nil {
		o.Skips = make(map[string]struct{})
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

func Skips(skips ...string) Option {
	return func(o *options) {
		if o.Skips == nil {
			o.Skips = make(map[string]struct{})
		}
		for _, skip := range skips {
			o.Skips[skip] = struct{}{}
		}
	}
}
