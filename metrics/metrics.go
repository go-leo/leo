package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	otelprometheus "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	"go.opentelemetry.io/otel/sdk/metric/export"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"

	"github.com/go-leo/leo/common/otelx"
)

type Metrics struct {
	controller         *controller.Controller
	prometheusExporter *otelprometheus.Exporter
	exporter           export.Exporter
}

func New(opts ...Option) (*Metrics, error) {
	o := new(options)
	o.apply(opts...)
	o.init()
	// CheckpointerFactory 是一种接口，用来创建配置好的Checkpointer实例
	checkpointerFactory := processor.NewFactory(
		o.AggregatorSelector,
		o.TemporalitySelector,
		processor.WithMemory(true),
	)

	ctrlOpts := []controller.Option{
		controller.WithResource(otelx.Resource(o.Attributes...)),
	}
	if o.CollectPeriod > 0 {
		ctrlOpts = append(ctrlOpts, controller.WithCollectPeriod(o.CollectPeriod))
	}
	if o.CollectTimeout > 0 {
		ctrlOpts = append(ctrlOpts, controller.WithCollectTimeout(o.CollectTimeout))
	}
	if o.PushTimeout > 0 {
		ctrlOpts = append(ctrlOpts, controller.WithPushTimeout(o.PushTimeout))
	}
	ctrl := controller.New(checkpointerFactory, ctrlOpts...)

	var prometheusExporter *otelprometheus.Exporter
	if o.Prometheus {
		// 注册prometheus官方的GoCollector和ProcessCollector
		registry := prometheus.NewRegistry()
		_ = registry.Register(collectors.NewGoCollector())
		_ = registry.Register(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		config := otelprometheus.Config{Registry: registry}
		if len(o.Boundaries) > 0 {
			config.DefaultHistogramBoundaries = o.Boundaries
		}
		exporter, err := otelprometheus.New(config, ctrl)
		if err != nil {
			return nil, err
		}
		prometheusExporter = exporter
	}

	if err := runtime.Start(
		runtime.WithMeterProvider(ctrl),
		runtime.WithMinimumReadMemStatsInterval(time.Second),
	); err != nil {
		return nil, err
	}

	return &Metrics{controller: ctrl, prometheusExporter: prometheusExporter}, nil
}

func (metrics *Metrics) MeterProvider() metric.MeterProvider {
	return metrics.controller
}

func (metrics *Metrics) PrometheusExporter() *otelprometheus.Exporter {
	return metrics.prometheusExporter
}

func (metrics *Metrics) Exporter() export.Exporter {
	return metrics.exporter
}
