package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"log"
	"time"
)

func main() {
	client, err := helloworld.NewGreeterHttpClientV2(
		"http",
		"127.0.0.1:8080",
		nil,
		nil,
		passthroughx.Factory{},
		nil,
		logx.L(),
		lbx.RoundRobinFactory{},
	)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	log.Printf("Greeting: %s", r.Message)
}
