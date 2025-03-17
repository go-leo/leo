package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"log"
	"time"
)

func main() {
	client := helloworld.NewGreeterHttpClient("localhost:60051")
	// 设置超时时间为1秒
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
