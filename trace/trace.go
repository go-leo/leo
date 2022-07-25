package trace

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type Trace struct {
	tracerProvider    trace.TracerProvider
	textMapPropagator propagation.TextMapPropagator
}

func New(ctx context.Context, opts ...Option) (*Trace, error) {
	o := new(options)
	o.apply(opts...)
	o.init()
	var ep ExporterProvider
	switch {
	case o.HTTPOptions != nil:
		ep = o.HTTPOptions
	case o.GRPCOptions != nil:
		ep = o.GRPCOptions
	case o.JaegerOptions != nil:
		ep = o.JaegerOptions
	case o.ZipkinOptions != nil:
		ep = o.ZipkinOptions
	case o.WriterOptions != nil:
		ep = o.WriterOptions
	default:
		return nil, errors.New("not found a validate exporter provider")
	}
	exporter, err := ep.Exporter(ctx)
	if err != nil {
		return nil, err
	}
	sampler := sdktrace.ParentBased(newSample(o.SampleRate))
	r := newResource(ctx, o.Attributes...)
	provider := newTracerProvider(exporter, sampler, r)
	propagator := propagation.NewCompositeTextMapPropagator(
		propagation.Baggage{},
		propagation.TraceContext{},
	)
	return &Trace{tracerProvider: provider, textMapPropagator: propagator}, nil
}

func (trace *Trace) TracerProvider() trace.TracerProvider {
	return trace.tracerProvider
}

func (trace *Trace) TextMapPropagator() propagation.TextMapPropagator {
	return trace.textMapPropagator
}

func newTracerProvider(exporter sdktrace.SpanExporter, sampler sdktrace.Sampler, resource *resource.Resource) *sdktrace.TracerProvider {
	return sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter), sdktrace.WithSampler(sampler), sdktrace.WithResource(resource))
}

func newSample(samplingRate float64) sdktrace.Sampler {
	var sampler sdktrace.Sampler
	switch {
	case samplingRate >= 1:
		sampler = sdktrace.AlwaysSample()
	case samplingRate <= 0:
		sampler = sdktrace.NeverSample()
	default:
		sampler = sdktrace.TraceIDRatioBased(samplingRate)
	}
	return sampler
}

func newResource(ctx context.Context, attrs ...attribute.KeyValue) *resource.Resource {
	attributes, err := resource.New(
		ctx,
		resource.WithAttributes(attrs...),
		resource.WithHost(),
		//resource.WithOS(),
		resource.WithProcess(),
		resource.WithContainer(),
	)
	if err != nil {
		return resource.Default()
	}
	return attributes
}
