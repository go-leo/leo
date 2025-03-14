package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/metadatax"
	"log"
)

func main() {
	flag.Parse()
	client := helloworld.NewGreeterHttpClient("localhost:60051")
	ctx := context.Background()
	ctx = metadatax.AppendOutgoingContext(ctx, metadatax.Pairs("token", "1234567890"))
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
