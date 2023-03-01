package trace

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func TraceIDFromContext(ctx context.Context) (string, bool) {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.HasTraceID() {
		return spanContext.TraceID().String(), true
	}
	return "", false
}
