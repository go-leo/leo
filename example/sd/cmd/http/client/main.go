package main

import (
	"context"
	"flag"
	"github.com/go-leo/gox/netx/httpx/outgoing"
	"github.com/go-leo/leo/v3/example/sd/api"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:60051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	client := api.NewGreeterHttpClient(*addr)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	r, err := client.SayHello(ctx, &api.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
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
