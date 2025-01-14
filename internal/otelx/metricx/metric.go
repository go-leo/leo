package metricx

import (
	"context"
	"errors"

	"github.com/go-leo/leo/internal/otelx/resourcex"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type Metric struct {
	meterProvider *sdkmetric.MeterProvider
}

func NewMetric(ctx context.Context, opts ...Option) (*Metric, error) {
	o := new(options)
	o.apply(opts...)
	o.init()
	var ep ExporterProvider
	switch {
	case o.PrometheusOptions != nil:
		ep = o.PrometheusOptions
	case o.GRPCOptions != nil:
		ep = o.GRPCOptions
	case o.HTTPOptions != nil:
		ep = o.HTTPOptions
	case o.WriterOptions != nil:
		ep = o.WriterOptions
	default:
		return nil, errors.New("not found a metric ExporterProvider")
	}
	exporter, err := ep.Exporter(ctx)
	if err != nil {
		return nil, err
	}
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resourcex.NewResource(ctx, o.Service, o.Resources, o.Attributes...)),
		sdkmetric.WithReader(exporter),
		sdkmetric.WithView(newView()),
	)
	return &Metric{meterProvider: meterProvider}, nil
}

func allowedAttr(v ...string) attribute.Filter {
	m := make(map[string]struct{}, len(v))
	for _, s := range v {
		m[s] = struct{}{}
	}
	return func(kv attribute.KeyValue) bool {
		_, ok := m[string(kv.Key)]
		return ok
	}
}

func newView() sdkmetric.View {
	opt := ViewOption{
		Criteria: sdkmetric.Instrument{
			Scope: instrumentation.Scope{
				Name: "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
			},
		},
		Mask: sdkmetric.Stream{
			AttributeFilter: allowedAttr(
				"rpc.grpc.status.code",
				"rpc.method",
				"rpc.service",
				"rpc.system",
			),
		},
	}
	return sdkmetric.NewView(opt.Criteria, opt.Mask)
}

func (metric *Metric) MeterProvider() metric.MeterProvider {
	return metric.meterProvider
}
