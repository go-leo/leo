package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/transportx"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	transports, err := helloworld.NewGreeterGrpcClientTransports(":9090", transportx.GrpcDialOption(grpc1.WithTransportCredentials(insecure.NewCredentials())))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := helloworld.NewGreeterGrpcClient(transports)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	log.Printf("Greeting: %s", r.Message)
}
