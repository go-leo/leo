package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	transports, err := helloworld.NewGreeterGrpcClientTransports(
		"consul://localhost:8500/demo.grpc?dc=dc1",
		[]grpc1.DialOption{grpc1.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		panic(err)
	}
	client := helloworld.NewGreeterGrpcClient(transports)

	for i := 0; i < 90; i++ {
		callRpc(client)
	}
}

func callRpc(client helloworld.GreeterService) {
	ctx := context.Background()
	reply, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
