package health

import (
	"context"
	"sync"
)

type Server struct {
	mu       sync.RWMutex
	shutdown bool
	status   ServingStatus
}

func NewServer() *Server {
	return &Server{status: SERVING}
}

func (s *Server) Check(_ context.Context) *HealthCheckResponse {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return &HealthCheckResponse{Status: s.status}
}

func (s *Server) Shutdown() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.shutdown = true
	s.setServingStatusLocked(NOT_SERVING)
}

func (s *Server) Resume() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.shutdown = false
	s.setServingStatusLocked(SERVING)
}

func (s *Server) setServingStatusLocked(servingStatus ServingStatus) {
	s.status = servingStatus
}
