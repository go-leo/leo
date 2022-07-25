package otelx

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/credentials"
)

func NewWriterTracerProvider(writer io.Writer, samplingRate float64, attrs ...attribute.KeyValue) trace.TracerProvider {
	exporter, _ := stdouttrace.New(
		stdouttrace.WithWriter(writer),
	)
	tp := NewTracerProvider(exporter, sdktrace.ParentBased(newSample(samplingRate)), Resource(attrs...))
	return tp
}

func NewJeagerTracerProvider(endpoint string, samplingRate float64, attrs ...attribute.KeyValue) (trace.TracerProvider, error) {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return nil, fmt.Errorf("failed new jager exporter, %w", err)
	}
	tp := NewTracerProvider(exporter, sdktrace.ParentBased(newSample(samplingRate)), Resource(attrs...))
	return tp, nil
}

func NewZipkinTracerProvider(url string, samplingRate float64, attrs ...attribute.KeyValue) (trace.TracerProvider, error) {
	exporter, err := zipkin.New(url)
	if err != nil {
		return nil, fmt.Errorf("failed new zipkin exporter, %w", err)
	}
	tp := NewTracerProvider(exporter, sdktrace.ParentBased(newSample(samplingRate)), Resource(attrs...))
	return tp, nil
}

func NewGRPCTracerProvider(ctx context.Context, options struct {
	endpoint           string
	insecure           bool
	tlsConfig          *tls.Config
	ReconnectionPeriod time.Duration
	compressor         string
	headers            map[string]string
}, samplingRate float64, attrs ...attribute.KeyValue) (trace.TracerProvider, error) {
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(options.endpoint),
	}
	if options.insecure {
		opts = append(opts, otlptracegrpc.WithInsecure())
	}
	if options.tlsConfig != nil {
		opts = append(opts, otlptracegrpc.WithTLSCredentials(credentials.NewTLS(options.tlsConfig)))
	}
	if len(options.headers) > 0 {
		opts = append(opts, otlptracegrpc.WithHeaders(options.headers))
	}
	exporter, err := otlptracegrpc.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed new grpc exporter, %w", err)
	}
	tp := NewTracerProvider(exporter, sdktrace.ParentBased(newSample(samplingRate)), Resource(attrs...))
	return tp, nil
}

func NewHTTPTracerProvider(ctx context.Context, url string, headers map[string]string, insecure bool, tlsConfig *tls.Config, samplingRate float64, attrs ...attribute.KeyValue) (trace.TracerProvider, error) {
	opts := []otlptracehttp.Option{
		otlptracehttp.WithEndpoint(url),
	}
	if insecure {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	if tlsConfig != nil {
		opts = append(opts, otlptracehttp.WithTLSClientConfig(tlsConfig))
	}
	if len(headers) > 0 {
		opts = append(opts, otlptracehttp.WithHeaders(headers))
	}
	exporter, err := otlptracehttp.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed new zipkin exporter, %w", err)
	}
	tp := NewTracerProvider(exporter, sdktrace.ParentBased(newSample(samplingRate)), Resource(attrs...))
	return tp, nil
}

func NewTracerProvider(exporter sdktrace.SpanExporter, sampler sdktrace.Sampler, resource *resource.Resource) *sdktrace.TracerProvider {
	return sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter), sdktrace.WithSampler(sampler), sdktrace.WithResource(resource))
}

func newSample(samplingRate float64, attrs ...attribute.KeyValue) sdktrace.Sampler {
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
