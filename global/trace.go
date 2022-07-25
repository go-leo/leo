package global

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func TracerProvider() oteltrace.TracerProvider {
	return otel.GetTracerProvider()
}

func SetTracerProvider(tp oteltrace.TracerProvider) func() {
	prev := otel.GetTracerProvider()
	otel.SetTracerProvider(tp)
	return func() { SetTracerProvider(prev) }
}

func SetTextMapPropagator(propagator propagation.TextMapPropagator) func() {
	prev := otel.GetTextMapPropagator()
	otel.SetTextMapPropagator(propagator)
	return func() { SetTextMapPropagator(prev) }
}

func initTrace(ctx context.Context) error {
	traceConf := Configuration().Trace
	if !traceConf.Enabled {
		return nil
	}

	t, err := traceConf.NewTrace(ctx)
	if err != nil {
		return err
	}
	if t == nil {
		return nil
	}
	_ = SetTracerProvider(t.TracerProvider())
	_ = SetTextMapPropagator(t.TextMapPropagator())
	return nil
}
