package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/example/api/demo"
	"github.com/go-leo/leo/v3/example/internal/demo/assembler"
	"github.com/go-leo/leo/v3/example/internal/demo/command"
	"github.com/go-leo/leo/v3/example/internal/demo/query"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
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
	service := demo.NewDemoHTTPServer(endpoints, []httptransport.ServerOption{})
	server := http.Server{Handler: service}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
