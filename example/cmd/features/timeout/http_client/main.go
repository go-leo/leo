package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"log"
	"time"
)

func main() {
	client := helloworld.NewGreeterHttpClient("127.0.0.1:8080")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Printf("Greeting: %s", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_, err = client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Printf("Greeting: %s", err)
	}
}
