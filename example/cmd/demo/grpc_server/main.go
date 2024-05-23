package main

import (
	grpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-leo/leo/v3/example/api/demo"
	"github.com/go-leo/leo/v3/example/internal/demo/assembler"
	"github.com/go-leo/leo/v3/example/internal/demo/command"
	"github.com/go-leo/leo/v3/example/internal/demo/query"
	grpc1 "google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc1.NewServer()
	bus, err := demo.NewDemoBus(
		command.NewCreateUser(),
		command.NewDeleteUser(),
		command.NewUpdateUser(),
		query.NewGetUser(),
		query.NewGetUsers(),
		command.NewUploadUserAvatar(),
		query.NewGetUserAvatar(),
		query.NewPushUsers(),
	)
	if err != nil {
		log.Fatalf("failed to new bus: %v", err)
	}
	demoAssembler := assembler.NewDemoAssembler()
	cqrsService := demo.NewDemoCQRSService(bus, demoAssembler)
	endpoints := demo.NewDemoEndpoints(cqrsService)
	service := demo.NewDemoGRPCServer(endpoints, []grpc.ServerOption{})
	demo.RegisterDemoServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
