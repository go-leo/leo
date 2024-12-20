package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"time"
)

func main() {
	client, err := helloworld.NewGreeterHttpClientV2(
		"http",
		"consul://localhost:8500/demo.http?dc=dc1",
		nil,
		nil,
		consulx.Factory{},
		nil,
		logx.L(),
		lbx.RandomFactory{},
	)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 90; i++ {
		callRpc(client)
		time.Sleep(100 * time.Millisecond)
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
