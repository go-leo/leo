package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	"log"
	"os"
)

func main() {
	transports := helloworld.NewGreeterHttpClientTransports("http", "127.0.0.1:8080")
	client := helloworld.NewGreeterHttpClient(transports)

	ctx := context.Background()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		statusErr, _ := statusx.FromError(err)
		body := statusErr.HttpBody()
		log.Printf("body: %s", body)
		os.Exit(1)
	}
	log.Printf("Greeting: %s", r.Message)

}
