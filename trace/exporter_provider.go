package trace

import (
	"context"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type ExporterProvider interface {
	Exporter(ctx context.Context) (sdktrace.SpanExporter, error)
}
