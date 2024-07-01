package server

import "github.com/go-leo/leo/v3/runner"

type Server interface {
	runner.StartStopper
}

//type Server struct {
//	s *health.Server
//}
//
//func (s *Server) Start(ctx context.Context) error {
//	s.s.Resume()
//	return nil
//}
//
//func (s *Server) Stop(ctx context.Context) error {
//	s.s.Shutdown()
//	return nil
//}
//
//func NewServer() *Server {
//	return newServer(health.NewServer())
//}
//
//func newServer(s *health.Server) *Server {
//	return &Server{s: s}
//}
