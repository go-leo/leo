package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	"log"
	"os"
)

func main() {
	transports, err := helloworld.NewGreeterHttpClientTransports("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	client := helloworld.NewGreeterHttpClient(transports)

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
