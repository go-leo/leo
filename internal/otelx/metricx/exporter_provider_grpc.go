package metricx

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-leo/gox/stringx"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GRPCOptions struct {
	Endpoint               string
	Insecure               bool
	TLSConfig              *tls.Config
	Headers                map[string]string
	Compressor             string
	DialOptions            []grpc.DialOption
	GRPCConn               *grpc.ClientConn
	ReconnectionPeriod     time.Duration
	Retry                  *otlpmetricgrpc.RetryConfig
	Timeout                time.Duration
	ServiceConfig          string
	TemporalitySelector    sdkmetric.TemporalitySelector
	AggregationSelector    sdkmetric.AggregationSelector
	PeriodicReaderTimeout  time.Duration
	PeriodicReaderInterval time.Duration
}

func (o *GRPCOptions) Exporter(ctx context.Context) (sdkmetric.Reader, error) {
	var opts []otlpmetricgrpc.Option
	if stringx.IsNotBlank(o.Endpoint) {
		opts = append(opts, otlpmetricgrpc.WithEndpoint(o.Endpoint))
	}
	if o.Insecure {
		opts = append(opts, otlpmetricgrpc.WithInsecure())
	}
	if o.TLSConfig != nil {
		opts = append(opts, otlpmetricgrpc.WithTLSCredentials(credentials.NewTLS(o.TLSConfig)))
	}
	if len(o.Headers) > 0 {
		opts = append(opts, otlpmetricgrpc.WithHeaders(o.Headers))
	}
	if stringx.IsNotBlank(o.Compressor) {
		opts = append(opts, otlpmetricgrpc.WithCompressor(o.Compressor))
	}
	if len(o.DialOptions) > 0 {
		opts = append(opts, otlpmetricgrpc.WithDialOption(o.DialOptions...))
	}
	if o.GRPCConn != nil {
		opts = append(opts, otlpmetricgrpc.WithGRPCConn(o.GRPCConn))
	}
	if o.ReconnectionPeriod > 0 {
		opts = append(opts, otlpmetricgrpc.WithReconnectionPeriod(o.ReconnectionPeriod))
	}
	if o.Retry != nil {
		opts = append(opts, otlpmetricgrpc.WithRetry(*o.Retry))
	}
	if o.Timeout > 0 {
		opts = append(opts, otlpmetricgrpc.WithTimeout(o.Timeout))
	}
	if stringx.IsNotBlank(o.ServiceConfig) {
		opts = append(opts, otlpmetricgrpc.WithServiceConfig(o.ServiceConfig))
	}
	if o.TemporalitySelector != nil {
		opts = append(opts, otlpmetricgrpc.WithTemporalitySelector(o.TemporalitySelector))
	}
	if o.AggregationSelector != nil {
		opts = append(opts, otlpmetricgrpc.WithAggregationSelector(o.AggregationSelector))
	}
	exporter, err := otlpmetricgrpc.New(ctx, opts...)
	if err != nil {
		return nil, err
	}
	var prOpts []sdkmetric.PeriodicReaderOption
	if o.PeriodicReaderTimeout > 0 {
		prOpts = append(prOpts, sdkmetric.WithTimeout(o.PeriodicReaderTimeout))
	}
	if o.PeriodicReaderInterval > 0 {
		prOpts = append(prOpts, sdkmetric.WithInterval(o.PeriodicReaderInterval))
	}
	return sdkmetric.NewPeriodicReader(exporter, prOpts...), nil
}
