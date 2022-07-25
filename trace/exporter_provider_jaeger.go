package trace

import (
	"context"

	"go.opentelemetry.io/otel/exporters/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-leo/leo/common/httpx"
)

type JaegerOptions struct {
	Endpoint string
	Username string
	Password string
}

func (o *JaegerOptions) Exporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	endpoint := jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(o.Endpoint),
		jaeger.WithUsername(o.Username),
		jaeger.WithPassword(o.Password),
		jaeger.WithHTTPClient(httpx.PooledClient()),
	)
	return jaeger.New(endpoint)
}
