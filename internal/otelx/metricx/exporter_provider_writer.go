package metricx

import (
	"context"
	"io"
	"time"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type WriterOptions struct {
	// Writer 标准输入或者文件
	Writer                 io.Writer
	NewEncoder             func(writer io.Writer) stdoutmetric.Encoder
	TemporalitySelector    sdkmetric.TemporalitySelector
	AggregationSelector    sdkmetric.AggregationSelector
	PeriodicReaderTimeout  time.Duration
	PeriodicReaderInterval time.Duration
}

func (o *WriterOptions) Exporter(ctx context.Context) (sdkmetric.Reader, error) {
	var opts []stdoutmetric.Option
	if o.Writer != nil && o.NewEncoder != nil {
		enc := o.NewEncoder(o.Writer)
		opts = append(opts, stdoutmetric.WithEncoder(enc))
	}
	if o.TemporalitySelector != nil {
		opts = append(opts, stdoutmetric.WithTemporalitySelector(o.TemporalitySelector))
	}
	if o.AggregationSelector != nil {
		opts = append(opts, stdoutmetric.WithAggregationSelector(o.AggregationSelector))
	}
	exporter, err := stdoutmetric.New(opts...)
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
