package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	grpc1 "google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc1.NewServer()
	endpoints := helloworld.NewGreeterServerEndpoints(NewGreeterService())
	transports := helloworld.NewGreeterGrpcServerTransports(endpoints)
	service := helloworld.NewGreeterGrpcServer(transports)
	helloworld.RegisterGreeterServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GreeterService struct {
}

func (g GreeterService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, statusx.ErrResourceExhausted.
		WithMessage("too many requests").
		WithQuotaFailure(&errdetails.QuotaFailure{
			Violations: []*errdetails.QuotaFailure_Violation{
				{Subject: fmt.Sprintf("name:%s", request.Name), Description: "Limit one greeting per person"},
			},
		})
}

func NewGreeterService() helloworld.GreeterService {
	return &GreeterService{}
}
