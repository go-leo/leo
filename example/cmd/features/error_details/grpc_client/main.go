package main

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func main() {
	conn, err := grpc1.Dial(":9090", grpc1.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	factory := helloworld.NewGreeterGrpcClientFactory()
	endpointer := sd.NewEndpointer(instancer, factory, logger)
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(*retryMax, *retryTimeout, balancer)

	transports := helloworld.NewGreeterGrpcClientTransports(conn)
	helloworld.NewNewGreeterGrpcClientEndpoints(transports)
	client := helloworld.NewGreeterGrpcClient(transports)

	ctx := context.Background()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		statusErr, _ := statusx.FromError(err)
		failure := statusErr.QuotaFailure()
		log.Printf("Quota failure: %s", failure)
		os.Exit(1)
	}
	log.Printf("Greeting: %s", r.Message)
}
