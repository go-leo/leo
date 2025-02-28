package main

import (
	"context"
	"flag"
	"github.com/go-leo/gox/netx/httpx/outgoing"
	"github.com/go-leo/leo/v3/example/status/api"
	"log"
)

var (
	addr = flag.String("addr", "localhost:60051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	client := api.NewGreeterHttpClient(*addr)
	r, err := client.SayHello(context.Background(), &api.HelloRequest{
		Name: *name,
	})
	if err != nil {
		st, ok := api.IsInvalidName(err)
		if ok {
			log.Fatalf("could not greet: %v, identifier: %v, request info: %v", err, st.Identifier(), st.RequestInfo())
		}
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func mainStd() {
	flag.Parse()
	receiver, err := outgoing.Post().
		URLString("http://" + *addr + "/v1/example/echo").
		JSONBody(&api.HelloRequest{
			Name: *name,
		}).Send(context.Background())
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}
	body, err := receiver.TextBody()
	if err != nil {
		log.Fatalf("failed to read response: %v", err)
	}
	if receiver.StatusCode() != 200 {
		log.Fatalf("could not greet: %v", body)
	}
	log.Printf("Greeting: %s", body)
}
