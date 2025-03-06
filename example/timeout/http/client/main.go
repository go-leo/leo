package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:60051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	client := helloworld.NewGreeterHttpClient(*addr)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
