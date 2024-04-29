package main

import (
	"github.com/go-leo/cqrs/example/api/demo"
	"github.com/go-leo/cqrs/example/internal/demo/assembler"
	"github.com/go-leo/cqrs/example/internal/demo/command"
	"github.com/go-leo/cqrs/example/internal/demo/query"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bus, err := demo.NewDemoServiceBus(
		command.NewCreateUser(),
		command.NewUpdateUser(),
		query.NewGetUser(),
		query.NewGetUsers(),
		command.NewDeleteUser(),
		query.NewAsyncGetUsers(),
		command.NewAsyncDeleteUsers(),
	)
	if err != nil {
		panic(err)
	}
	service := demo.NewDemoServiceCQRSService(bus, assembler.NewDemoServiceAssembler())
	demo.RegisterDemoServiceServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
