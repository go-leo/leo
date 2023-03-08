package global

import (
	"context"
	"sync"

	otelprometheus "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

var (
	prometheusExporter *otelprometheus.Exporter
	meterLocker        sync.RWMutex
)

func GetMeterProvider() metric.MeterProvider {
	return global.MeterProvider()
}

func SetMeterProvider(mp metric.MeterProvider) func() {
	prev := global.MeterProvider()
	global.SetMeterProvider(mp)
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
