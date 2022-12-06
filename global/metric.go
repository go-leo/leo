package global

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
)

// MeterProvider return otel MeterProvider.
// see github.com/go-leo/otelx.
// see github.com/open-telemetry/opentelemetry-go.
func MeterProvider() metric.MeterProvider {
	return global.MeterProvider()
}

// SetMeterProvider set otel MeterProvider.
// see github.com/go-leo/otelx.
// see github.com/open-telemetry/opentelemetry-go.
func SetMeterProvider(mp metric.MeterProvider) func() {
	prev := global.MeterProvider()
	global.SetMeterProvider(mp)
	return func() { SetMeterProvider(prev) }
}
