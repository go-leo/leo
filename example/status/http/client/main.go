package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	api "github.com/go-leo/leo/v3/example/api/status"
	"log"
)

var (
	addr = flag.String("addr", "localhost:60051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	client := helloworld.NewGreeterHttpClient(*addr)
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: *name,
	})
	if err != nil {
		st, ok := api.IsInvalidName(err)
		if ok {
			log.Fatalf("could not greet: %v, identifier: %v, request info: %v", err, st.Identifier(), st.RequestInfo())
		}
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
