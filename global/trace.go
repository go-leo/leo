package global

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// TracerProvider return otel TracerProvider.
// see github.com/go-leo/otelx.
// see github.com/open-telemetry/opentelemetry-go.
func TracerProvider() oteltrace.TracerProvider {
	return otel.GetTracerProvider()
}

// SetTracerProvider set otel TracerProvider.
// see github.com/go-leo/otelx.
// see github.com/open-telemetry/opentelemetry-go.
func SetTracerProvider(tp oteltrace.TracerProvider) func() {
	prev := otel.GetTracerProvider()
	otel.SetTracerProvider(tp)
	return func() { SetTracerProvider(prev) }
}

// TextMapPropagator return otel TextMapPropagator.
// see github.com/go-leo/otelx.
// see github.com/open-telemetry/opentelemetry-go.
func TextMapPropagator() propagation.TextMapPropagator {
	return otel.GetTextMapPropagator()
}

// SetTextMapPropagator set otel TextMapPropagator.
// see github.com/go-leo/otelx.
// see github.com/open-telemetry/opentelemetry-go.
func SetTextMapPropagator(propagator propagation.TextMapPropagator) func() {
	prev := otel.GetTextMapPropagator()
	otel.SetTextMapPropagator(propagator)
	return func() { SetTextMapPropagator(prev) }
}
