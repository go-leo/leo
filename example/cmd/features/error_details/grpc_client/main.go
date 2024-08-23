package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	transports, err := helloworld.NewGreeterGrpcClientTransports(":9090", transportx.GrpcDialOption(grpc1.WithTransportCredentials(insecure.NewCredentials())))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := helloworld.NewGreeterGrpcClient(transports)
	ctx := context.Background()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		statusErr := statusx.From(err)
		failure := statusErr.QuotaFailure()
		log.Printf("Quota failure: %s", failure)
		return
	}
	log.Printf("Greeting: %s", r.Message)
}
