package metricx

import (
	"context"

	prome "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

type PrometheusOptions struct {
	Registerer        prome.Registerer
	Aggregation       metric.AggregationSelector
	WithoutUnits      bool
	WithoutTargetInfo bool
	WithoutScopeInfo  bool
}

func (o *PrometheusOptions) Exporter(ctx context.Context) (metric.Reader, error) {
	var opts []prometheus.Option
	if o.Registerer != nil {
		opts = append(opts, prometheus.WithRegisterer(o.Registerer))
	}
	if o.Aggregation != nil {
		opts = append(opts, prometheus.WithAggregationSelector(o.Aggregation))
	}
	if o.WithoutTargetInfo {
		opts = append(opts, prometheus.WithoutTargetInfo())
	}
	if o.WithoutUnits {
		opts = append(opts, prometheus.WithoutUnits())
	}
	if o.WithoutScopeInfo {
		opts = append(opts, prometheus.WithoutScopeInfo())
	}
	exporter, err := prometheus.New(opts...)
	if err != nil {
		return nil, err
	}
	return exporter, nil
}
