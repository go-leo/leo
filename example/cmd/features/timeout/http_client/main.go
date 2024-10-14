package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"log"
	"time"
)

func main() {
	transports, err := helloworld.NewGreeterHttpClientTransports("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	client := helloworld.NewGreeterHttpClient(transports)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	log.Printf("Greeting: %s", r.Message)
}
