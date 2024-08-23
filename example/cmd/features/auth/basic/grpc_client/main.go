package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/transportx"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	transports, err := helloworld.NewGreeterGrpcClientTransports(":9090", transportx.GrpcDialOption(grpc1.WithTransportCredentials(insecure.NewCredentials())))
	if err != nil {
		panic(err)
	}

	// ok
	client := helloworld.NewGreeterGrpcClient(transports, basicx.Middleware("soyacen", "123456", "basic auth example"))
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterGrpcClient(transports, basicx.Middleware("soyacen", "654321", "basic auth example"))
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
