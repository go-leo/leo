package http

import (
	"context"
	"sync/atomic"

	"github.com/gin-gonic/gin"
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
	status      *atomic.Int32
	OKStatus    int
	NotOKStatus int
}

func newHealthServer(ss ServingStatus, okStatus, notOKStatus int) *healthServer {
	status := new(atomic.Int32)
	status.Store(int32(ss))
	return &healthServer{status: status, OKStatus: okStatus, NotOKStatus: notOKStatus}
}

func (s *healthServer) HandlerFunc(c *gin.Context) {
	healthCheckResp := s.Check(c.Request.Context())
	if healthCheckResp.Status == SERVING {
		c.JSON(s.OKStatus, healthCheckResp)
		return
	}
	c.JSON(s.NotOKStatus, healthCheckResp)
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
