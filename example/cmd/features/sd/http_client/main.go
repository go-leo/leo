package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld"
)

func main() {
	transports, err := helloworld.NewGreeterHttpClientTransports(
		"consul://localhost:8500/demo.http?dc=dc1",
		"http")
	if err != nil {
		panic(err)
	}

	client := helloworld.NewGreeterHttpClient(transports)

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
