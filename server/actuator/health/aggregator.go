package health

import (
	"github.com/go-leo/leo/v3/healthx"
	"sync"
)

var (
	aggregator StatusAggregator = statusAggregator{
		healthx.ServiceUnknown,
		healthx.NotServing,
		healthx.Serving,
		healthx.Unknown,
	}
	aggregatorMu sync.RWMutex
)

func RegisterStatusAggregator(agg StatusAggregator) {
	mapperMu.Lock()
	defer mapperMu.Unlock()
	aggregator = agg
}

func getStatusAggregator() StatusAggregator {
	mapperMu.RLock()
	agg := aggregator
	mapperMu.RUnlock()
	return agg
}

// StatusAggregator aggregates multiple Statuses into one.
type StatusAggregator interface {
	AggregateStatus(statuses ...healthx.Status) healthx.Status
}

// statusAggregator is an ordered list of statuses.
type statusAggregator []healthx.Status

func (agg statusAggregator) AggregateStatus(statuses ...healthx.Status) healthx.Status {
	for _, orderStatus := range agg {
		for _, status := range statuses {
			if orderStatus.Code() == status.Code() {
				return status
			}
		}
	}
	return healthx.Unknown
}
