package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	client := helloworld.NewGreeterGrpcClient(":9090",
		grpcx.DialOptions(grpc1.WithTransportCredentials(insecure.NewCredentials())),
	)
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
