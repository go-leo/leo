package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/middleware/authx/basicx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// ok
	client := helloworld.NewGreeterGrpcClient(":9090",
		grpcx.DialOptions(grpc1.WithTransportCredentials(insecure.NewCredentials())),
		grpcx.Middleware(basicx.Middleware("soyacen", "123456", "basic auth example")),
	)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterGrpcClient(":9090",
		grpcx.DialOptions(grpc1.WithTransportCredentials(insecure.NewCredentials())),
		grpcx.Middleware(basicx.Middleware("soyacen", "654321", "basic auth example")),
	)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
