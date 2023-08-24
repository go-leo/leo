package metric

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/embedded"
	"go.opentelemetry.io/otel/metric/noop"
	"google.golang.org/grpc"
	"testing"
	"time"
)

type testMeterProvider struct{ embedded.MeterProvider }

var _ metric.MeterProvider = &testMeterProvider{}

func (*testMeterProvider) Meter(_ string, _ ...metric.MeterOption) metric.Meter {
	return noop.NewMeterProvider().Meter("")
}

func TestUnaryServerInterceptor(t *testing.T) {
	meterProvider := &testMeterProvider{}
	otel.SetMeterProvider(meterProvider)

	// Create an instance of your UnaryServerInterceptor
	interceptor := UnaryServerInterceptor()

	// Create a context and request for testing
	ctx := context.Background()
	req := struct{}{}

	// Call the interceptor with a mock handler
	mockHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		// Simulate some delay
		time.Sleep(50 * time.Millisecond)
		return "response", nil
	}
	info := &grpc.UnaryServerInfo{}
	_, _ = interceptor(ctx, req, info, mockHandler)

}
