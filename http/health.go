package http

import (
	"context"
	"encoding/json"
	"net/http"
	"sync/atomic"

	"github.com/go-leo/errorx"
)

type ServingStatus int32

const (
	UNKNOWN     ServingStatus = 0
	SERVING     ServingStatus = 1
	NOT_SERVING ServingStatus = 2
)

func (ss ServingStatus) String() string {
	switch ss {
	case UNKNOWN:
		return "UNKNOWN"
	case SERVING:
		return "SERVING"
	case NOT_SERVING:
		return "NOT_SERVING"
	default:
		return "UNKNOWN"
	}
}

type HealthCheckResp struct {
	Status ServingStatus `json:"status,omitempty"`
}

func (resp *HealthCheckResp) GetStatus() ServingStatus {
	if resp != nil {
		return resp.Status
	}
	return UNKNOWN
}

type healthServer struct {
	status *atomic.Int32
}

func newHealthServer(ss ServingStatus) *healthServer {
	status := new(atomic.Int32)
	status.Store(int32(ss))
	return &healthServer{status: status}
}

func (s *healthServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	healthCheckResp := s.Check(request.Context())
	if healthCheckResp.Status == SERVING {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(errorx.Quiet(json.Marshal(healthCheckResp)))
		return
	}
	writer.WriteHeader(http.StatusServiceUnavailable)
	_, _ = writer.Write(errorx.Quiet(json.Marshal(healthCheckResp)))
}

func (s *healthServer) Check(_ context.Context) *HealthCheckResp {
	return &HealthCheckResp{Status: ServingStatus(s.status.Load())}
}

func (s *healthServer) Shutdown(_ context.Context) {
	s.status.Store(int32(NOT_SERVING))
}

func (s *healthServer) Resume(_ context.Context) {
	s.status.Store(int32(SERVING))
}
