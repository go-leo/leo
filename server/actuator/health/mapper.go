package health

import (
	"github.com/go-leo/leo/v3/healthx"
	"net/http"
	"sync"
)

var (
	mapper HttpStatusMapper = httpStatusMapper{
		healthx.Unknown:        http.StatusOK,
		healthx.Serving:        http.StatusOK,
		healthx.NotServing:     http.StatusServiceUnavailable,
		healthx.ServiceUnknown: http.StatusServiceUnavailable,
	}
	mapperMu sync.RWMutex
)

func RegisterHttpStatusMapper(m HttpStatusMapper) {
	mapperMu.Lock()
	defer mapperMu.Unlock()
	mapper = m
}

func getHttpStatusMapper() HttpStatusMapper {
	mapperMu.RLock()
	m := mapper
	mapperMu.RUnlock()
	return m
}

type HttpStatusMapper interface {
	MapStatus(status healthx.Status) int
}

type httpStatusMapper map[healthx.Status]int

func (mapper httpStatusMapper) MapStatus(status healthx.Status) int {
	return mapper[status]
}
