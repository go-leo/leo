package metricx

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/go-leo/gox/stringx"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type HTTPOptions struct {
	Endpoint               string
	URLPath                string
	Compression            otlpmetrichttp.Compression
	TLSConfig              *tls.Config
	Insecure               bool
	Headers                map[string]string
	Timeout                time.Duration
	Retry                  *otlpmetrichttp.RetryConfig
	TemporalitySelector    sdkmetric.TemporalitySelector
	AggregationSelector    sdkmetric.AggregationSelector
	PeriodicReaderTimeout  time.Duration
	PeriodicReaderInterval time.Duration
}

func (o *HTTPOptions) Exporter(ctx context.Context) (sdkmetric.Reader, error) {
	var opts []otlpmetrichttp.Option
	if stringx.IsNotBlank(o.Endpoint) {
		opts = append(opts, otlpmetrichttp.WithEndpoint(o.Endpoint))
	}
	if stringx.IsNotBlank(o.URLPath) {
		opts = append(opts, otlpmetrichttp.WithURLPath(o.URLPath))
	}
	if o.Compression > 0 {
		opts = append(opts, otlpmetrichttp.WithCompression(o.Compression))
	}
	if o.TLSConfig != nil {
		opts = append(opts, otlpmetrichttp.WithTLSClientConfig(o.TLSConfig))
	}
	if o.Insecure {
		opts = append(opts, otlpmetrichttp.WithInsecure())
	}
	if len(o.Headers) > 0 {
		opts = append(opts, otlpmetrichttp.WithHeaders(o.Headers))
	}
	if o.Timeout > 0 {
		opts = append(opts, otlpmetrichttp.WithTimeout(o.Timeout))
	}
	if o.Retry != nil {
		opts = append(opts, otlpmetrichttp.WithRetry(*o.Retry))
	}
	if o.TemporalitySelector != nil {
		opts = append(opts, otlpmetrichttp.WithTemporalitySelector(o.TemporalitySelector))
	}
	if o.AggregationSelector != nil {
		opts = append(opts, otlpmetrichttp.WithAggregationSelector(o.AggregationSelector))
	}
	exporter, err := otlpmetrichttp.New(ctx, opts...)
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
