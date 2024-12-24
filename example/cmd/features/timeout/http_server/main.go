package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	router := helloworld.AppendGreeterHttpRoutes(
		mux.NewRouter(),
		NewGreeterService(),
	)
	server := http.Server{Handler: router}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
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
