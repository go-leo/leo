package tracex

import (
	"context"
	"net/http"
	"runtime"

	"go.opentelemetry.io/otel/exporters/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type JaegerOptions struct {
	Endpoint string
	Username string
	Password string
}

func (o *JaegerOptions) Exporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConnsPerHost = runtime.GOMAXPROCS(0) + 1
	endpoint := jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(o.Endpoint),
		jaeger.WithUsername(o.Username),
		jaeger.WithPassword(o.Password),
		jaeger.WithHTTPClient(&http.Client{Transport: transport}),
	)
	return jaeger.New(endpoint)
}
