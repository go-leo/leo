package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-leo/leo/v3/example/status/api"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

var (
	port = flag.Int("port", 60051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

func (s *server) SayHello(_ context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	if in.GetName() == "" {
		return nil, api.ErrInvalidName(statusx.RequestInfo(uuid.NewString(), in.GetName()))
	}
	log.Printf("Received: %v", in.GetName())
	return &api.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := &http.Server{
		Handler: api.AppendGreeterHttpServerRoutes(mux.NewRouter(), &server{}),
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
