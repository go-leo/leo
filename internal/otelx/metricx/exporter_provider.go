package metricx

import (
	"context"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

type ExporterProvider interface {
	Exporter(ctx context.Context) (sdkmetric.Reader, error)
}
