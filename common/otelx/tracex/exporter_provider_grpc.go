package tracex

import (
	"context"
	"crypto/tls"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/go-leo/gox/stringx"
)

type GRPCOptions struct {
	Endpoint           string
	Insecure           bool
	TLSConfig          *tls.Config
	Headers            map[string]string
	Compressor         string
	DialOptions        []grpc.DialOption
	GRPCConn           *grpc.ClientConn
	ReconnectionPeriod time.Duration
	Retry              *otlptracegrpc.RetryConfig
	Timeout            time.Duration
	ServiceConfig      string
}

func (o *GRPCOptions) Exporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	var opts []otlptracegrpc.Option
	if stringx.IsNotBlank(o.Endpoint) {
		opts = append(opts, otlptracegrpc.WithEndpoint(o.Endpoint))
	}
	if o.Insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}
	if o.TLSConfig != nil {
		opts = append(opts, otlptracegrpc.WithTLSCredentials(credentials.NewTLS(o.TLSConfig)))
	}
	if len(o.Headers) > 0 {
		opts = append(opts, otlptracegrpc.WithHeaders(o.Headers))
	}
	if stringx.IsNotBlank(o.Compressor) {
		opts = append(opts, otlptracegrpc.WithCompressor(o.Compressor))
	}
	if len(o.DialOptions) > 0 {
		opts = append(opts, otlptracegrpc.WithDialOption(o.DialOptions...))
	}
	if o.GRPCConn != nil {
		opts = append(opts, otlptracegrpc.WithGRPCConn(o.GRPCConn))
	}
	if o.ReconnectionPeriod > 0 {
		opts = append(opts, otlptracegrpc.WithReconnectionPeriod(o.ReconnectionPeriod))
	}
	if o.Retry != nil {
		opts = append(opts, otlptracegrpc.WithRetry(*o.Retry))
	}
	if o.Timeout > 0 {
		opts = append(opts, otlptracegrpc.WithTimeout(o.Timeout))
	}
	if stringx.IsNotBlank(o.ServiceConfig) {
		opts = append(opts, otlptracegrpc.WithServiceConfig(o.ServiceConfig))
	}
	return otlptracegrpc.New(ctx, opts...)
}
