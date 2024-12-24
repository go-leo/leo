package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	grpc1 "google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc1.NewServer()
	service := helloworld.NewGreeterGrpcServer(NewGreeterService())
	helloworld.RegisterGreeterServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GreeterService struct {
}

func (g GreeterService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	select {
	case <-ctx.Done():
		return nil, statusx.ErrDeadlineExceeded.With(statusx.Message("ctx Done"))
	case <-time.After(5 * time.Second):
		return nil, statusx.ErrDeadlineExceeded.With(statusx.Message("超时"))
	case <-time.After(2000 * time.Second):

	}
	return &helloworld.HelloReply{Message: "hi " + request.GetName()}, nil
}

func NewGreeterService() helloworld.GreeterService {
	return &GreeterService{}
}
