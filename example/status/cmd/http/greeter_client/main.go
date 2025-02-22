package main

import (
	"context"
	"flag"
	"github.com/go-leo/gox/netx/httpx/outgoing"
	helloworldpb "github.com/go-leo/status/example/api/helloworld"
	"log"
)

var (
	addr = flag.String("addr", "localhost:60051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	receiver, err := outgoing.Post().
		URLString("http://"+*addr+"/v1/example/echo").
		ObjectBody(&helloworldpb.HelloRequest{Name: *name}, func(a any) ([]byte, error) {
			return nil, nil
		}, "application/json").Send(context.Background())
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
