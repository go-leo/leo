package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"github.com/go-leo/leo/v3/statusx"
	"log"
	"os"
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

	ctx := context.Background()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		statusErr := statusx.From(err)
		failure := statusErr.QuotaFailure()
		log.Printf("Quota failure: %s", failure)
		body := statusErr.HttpBody()
		log.Printf("body: %s", body)
		os.Exit(1)
	}
	log.Printf("Greeting: %s", r.Message)

}
