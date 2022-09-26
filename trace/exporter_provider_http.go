package trace

import (
	"context"
	"crypto/tls"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-leo/stringx"
)

type HTTPOptions struct {
	Endpoint    string
	Insecure    bool
	TLSConfig   *tls.Config
	Headers     map[string]string
	Compression otlptracehttp.Compression
	Retry       *otlptracehttp.RetryConfig
	Timeout     time.Duration
	URLPath     string
}

func (o *HTTPOptions) Exporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(o.Endpoint),
	}
	if o.Insecure {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	if o.TLSConfig != nil {
		opts = append(opts, otlptracehttp.WithTLSClientConfig(o.TLSConfig))
	}
	if len(o.Headers) > 0 {
		opts = append(opts, otlptracehttp.WithHeaders(o.Headers))
	}
	if o.Compression > 0 {
		opts = append(opts, otlptracehttp.WithCompression(o.Compression))
	}
	if o.Retry != nil {
		opts = append(opts, otlptracehttp.WithRetry(*o.Retry))
	}
	if o.Timeout > 0 {
		opts = append(opts, otlptracehttp.WithTimeout(o.Timeout))
	}
	if stringx.IsNotBlank(o.URLPath) {
		opts = append(opts, otlptracehttp.WithURLPath(o.URLPath))
	}
	return otlptracehttp.New(ctx, opts...)
}
