package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	c := helloworld.NewGreeterGrpcClient("localhost:50051", grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	ctx := context.Background()
	ctx = metadatax.NewOutgoingContext(ctx, metadatax.Pairs("token", "1234567890"))
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
