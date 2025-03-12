package main

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/netx/httpx/outgoing"
)

func main() {
	all()
	component("grpc")
	component("http")
	component("custom")
}

func all() {
	receiver, err := outgoing.Get().
		URLString("http://localhost:16060/health").
		Send(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(receiver.StatusCode())
	body, err := receiver.TextBody()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

}

func component(s string) {
	receiver, err := outgoing.Get().
		URLString("http://localhost:16060/health/" + s).
		Send(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(receiver.StatusCode())
	body, err := receiver.TextBody()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
