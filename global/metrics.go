package global

import (
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

func GetPrometheusExporter() *otelprometheus.Exporter {
	meterLocker.RLock()
	defer meterLocker.RUnlock()
	return prometheusExporter
}

func SetPrometheusExporter(e *otelprometheus.Exporter) func() {
	meterLocker.RLock()
	defer meterLocker.RUnlock()
	prev := prometheusExporter
	prometheusExporter = e
	SetMeterProvider(e.MeterProvider())
	return func() { SetPrometheusExporter(prev) }
}

func initMetric() error {
	metricConf := Configuration().Metrics
	if !metricConf.Enabled {
		return nil
	}

	m, err := metricConf.NewMetric()
	if err != nil {
		return err
	}

	exporter := m.PrometheusExporter()
	if exporter != nil {
		SetPrometheusExporter(exporter)
		return nil
	}

	SetMeterProvider(m.MeterProvider())
	return nil
}
