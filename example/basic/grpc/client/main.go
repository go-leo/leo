package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// success
	// basic 中间件
	mdw := basicx.Client("ubuntu", "mint")
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	// basic 中间件
	mdw = basicx.Client("ubuntu", "redhat")
	client = helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	r, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}
