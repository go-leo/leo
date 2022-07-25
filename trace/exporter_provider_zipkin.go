package trace

import (
	"context"

	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-leo/leo/common/httpx"
)

type ZipkinOptions struct {
	URL string
}

func (o *ZipkinOptions) Exporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	return zipkin.New(o.URL, zipkin.WithClient(httpx.PooledClient()))
}
