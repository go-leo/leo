package main

import (
	"context"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/sdx/lbx"
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

	ctx := context.TODO()
	factories := helloworld.NewGreeterGrpcClientFactories()
	endpointers := helloworld.NewGreeterEndpointers(factories, nil, nil, nil)
	endpointer := endpointers.SayHello(ctx, nil)
	factory := lbx.RandomFactory{}
	balancer := factory.New(ctx, endpointer)
	endpoint := lb.Retry(3, 10*1000, balancer)

	resp, err := endpoint(context.Background(), nil)

	transports := helloworld.NewGreeterGrpcClientTransports(conn)
	client := helloworld.NewGreeterGrpcClient(helloworld.NewGreeterGrpcClientEndpoints(transports))

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
