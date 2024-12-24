package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	client := helloworld.NewGreeterGrpcClient(
		"consul://localhost:8500/demo.grpc?dc=dc1",
		grpcx.DialOptions(grpc1.WithTransportCredentials(insecure.NewCredentials())),
		grpcx.InstancerFactory(consulx.Factory{}),
		grpcx.BalancerFactory(lbx.RoundRobinFactory{}),
	)
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
