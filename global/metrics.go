package global

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

func GetMeterProvider() metric.MeterProvider {
	return otel.GetMeterProvider()
}

func SetMeterProvider(mp metric.MeterProvider) func() {
	prev := otel.GetMeterProvider()
	otel.SetMeterProvider(mp)
	return func() { SetMeterProvider(prev) }
}

func initMetric(ctx context.Context) error {
	metricConf := Configuration().Metrics
	if !metricConf.Enabled {
		return nil
	}
	m, err := metricConf.NewMetric(ctx)
	if err != nil {
		return err
	}
	SetMeterProvider(m.MeterProvider())
	return nil
}
