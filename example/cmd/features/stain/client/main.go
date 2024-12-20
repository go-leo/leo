package main

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/gox/netx/httpx/outgoing"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"github.com/go-leo/leo/v3/sdx/stainx"
)

func main() {
	client, err := helloworld.NewGreeterHttpClientV2(
		"http",
		"127.0.0.1:8000",
		nil,
		nil,
		passthroughx.Factory{},
		nil,
		logx.L(),
		lbx.RandomFactory{},
	)
	if err != nil {
		panic(err)
	}

	var callApi = func(color string) {
		ctx := stainx.InjectColor(context.Background(), color)
		resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: color + "-color"})
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.GetMessage())
	}
	for i := 0; i < 900; i++ {
		r := randx.Int()
		if r%4 == 0 {
			callApi("red")
		} else if r%4 == 1 {
			callApi("blue")
		} else if r%4 == 2 {
			callApi("yellow")
		} else {
			callApi("")
		}
	}
}

func mainStd() {
	var callApi = func(color string) {
		receiver, err := outgoing.Sender().Get().
			URLString("http://localhost:8000/helloworld/"+color+"-color").
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
	for i := 0; i < 900; i++ {
		r := randx.Int()
		if r%4 == 0 {
			callApi("red")
		} else if r%4 == 1 {
			callApi("blue")
		} else if r%4 == 2 {
			callApi("yellow")
		} else {
			callApi("")
		}
	}
	select {}

}
