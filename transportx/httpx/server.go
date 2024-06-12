package httpx

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

type Server struct {
	server *httptransport.Server
}

func NewServer(server *httptransport.Server) *Server {
	return &Server{server: server}
}
