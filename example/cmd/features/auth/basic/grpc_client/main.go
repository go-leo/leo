package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	transports, err := helloworld.NewGreeterGrpcClientTransports(":9090", []grpc1.DialOption{grpc1.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		panic(err)
	}

	// ok
	endpoints := helloworld.NewGreeterClientEndpoints(
		transports,
		basicx.Middleware("soyacen", "123456", "basic auth example"),
	)
	client := helloworld.NewGreeterGrpcClient(endpoints)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	endpoints = helloworld.NewGreeterClientEndpoints(
		transports,
		basicx.Middleware("soyacen", "654321", "basic auth example"),
	)
	client = helloworld.NewGreeterGrpcClient(endpoints)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
