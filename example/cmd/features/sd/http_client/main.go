package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/httpx"
)

func main() {
	client := helloworld.NewGreeterHttpClient(
		"consul://localhost:8500/demo.http?dc=dc1",
		httpx.InstancerFactory(consulx.Factory{}),
		httpx.BalancerFactory(lbx.RoundRobinFactory{}),
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
