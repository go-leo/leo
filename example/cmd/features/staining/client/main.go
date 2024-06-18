package main

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/netx/httpx/outgoing"
	"github.com/go-leo/leo/v3/example/api/helloworld"
)

func main() {
	go func() {
		for i := 0; i < 90; i++ {
			callApi("red")
		}
	}()

	go func() {
		for i := 0; i < 90; i++ {
			callApi("blue")
		}
	}()

	go func() {
		for i := 0; i < 90; i++ {
			callApi("yellow")
		}
	}()

	select {}

}

func callApi(color string) {
	req := &helloworld.HelloRequest{Name: "macbook"}
	receiver, err := outgoing.Sender().Post().
		URLString("http://localhost:8000/helloworld.Greeter/SayHello").
		JSONBody(req).
		Header("X-Color", color).Send(context.Background())
	if err != nil {
		panic(err)
	}
	var resp helloworld.HelloReply
	if err := receiver.JSONBody(&resp); err != nil {
		panic(err)
	}
	fmt.Println(resp.GetMessage())
}
