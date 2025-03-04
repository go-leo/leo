package healthx

import (
	"google.golang.org/grpc/health/grpc_health_v1"
	"sync"
)

var (
	aggregator   Aggregator = defaultAggregator{}
	aggregatorMu sync.RWMutex
)

func RegisterAggregator(agg Aggregator) {
	aggregatorMu.Lock()
	aggregator = agg
	aggregatorMu.Unlock()
}

func GetAggregator() Aggregator {
	aggregatorMu.RLock()
	agg := aggregator
	aggregatorMu.RUnlock()
	return agg
}

// Aggregator aggregates multiple Statuses into one.
type Aggregator interface {
	Aggregate(statuses ...grpc_health_v1.HealthCheckResponse_ServingStatus) grpc_health_v1.HealthCheckResponse_ServingStatus
}

// defaultAggregator is the default Aggregator.
type defaultAggregator struct{}

func (defaultAggregator) Aggregate(statuses ...grpc_health_v1.HealthCheckResponse_ServingStatus) grpc_health_v1.HealthCheckResponse_ServingStatus {
	for _, status := range statuses {
		if status != grpc_health_v1.HealthCheckResponse_SERVING {
			return status
		}
	}
	return grpc_health_v1.HealthCheckResponse_SERVING
}
