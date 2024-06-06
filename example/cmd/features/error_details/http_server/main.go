package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	endpoints := helloworld.NewGreeterEndpoints(
		NewGreeterService(),
	)
	transports := helloworld.NewGreeterHttpServerTransports(endpoints)
	handler := helloworld.NewGreeterHttpServerHandler(transports)
	server := http.Server{Handler: handler}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GreeterService struct {
}

func (g GreeterService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, statusx.Failed.
		WithMessage("some thing wrong").
		WithHttpBody(&helloworld.CodeMessage{Code: 4040001, Message: "some thing wrong"})
}

func NewGreeterService() helloworld.GreeterService {
	return &GreeterService{}
}
