package grpcx

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type Server struct {
	server *grpctransport.Server
}

func NewServer(server *grpctransport.Server) *Server {
	return &Server{server: server}
}
