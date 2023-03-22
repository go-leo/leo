package ginhttp

import "sync/atomic"

type HealthServer struct {
	running *atomic.Bool
}

func (server *HealthServer) IsRunning() bool {
	return server.running.Load()
}

func (server *HealthServer) Shutdown() {
	server.running.Store(false)
}

func (server *HealthServer) Resume() {
	server.running.Store(true)
}
