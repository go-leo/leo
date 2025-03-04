package healthx

import (
	"google.golang.org/grpc/health/grpc_health_v1"
	"net/http"
	"sync"
)

var (
	mapper   Mapper = defaultMapper{}
	mapperMu sync.RWMutex
)

func RegisterMapper(m Mapper) {
	mapperMu.Lock()
	defer mapperMu.Unlock()
	mapper = m
}

func GetMapper() Mapper {
	mapperMu.RLock()
	m := mapper
	mapperMu.RUnlock()
	return m
}

type Mapper interface {
	Map(status grpc_health_v1.HealthCheckResponse_ServingStatus) int
}

type defaultMapper struct{}

func (defaultMapper) Map(status grpc_health_v1.HealthCheckResponse_ServingStatus) int {
	if status == grpc_health_v1.HealthCheckResponse_SERVING {
		return http.StatusOK
	}
	return http.StatusServiceUnavailable
}
