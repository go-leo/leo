package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-leo/leo/v3/example/timeout/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	deadline, ok := ctx.Deadline()
	log.Printf("timeout: %v, %v", deadline, ok)
	time.Sleep(10 * time.Second)
	return &api.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, api.NewGreeterGrpcServer(&server{}))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
