package main

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/gox/netx/httpx/outgoing"
	"github.com/go-leo/leo/v3/sdx/stain"
	helloworld "go.opencensus.io/examples/grpc/proto"
)

func main() {
	client := api.NewGreeterHttpClient("127.0.0.1:8000")
	var callApi = func(color string) {
		ctx := stain.InjectColor(context.Background(), color)
		resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: color})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp.GetMessage())
		}
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
			callApi("purple")
		}
	}
}

func mainStd() {
	var callApi = func(color string) {
		receiver, err := outgoing.Sender().Get().
			URLString("http://localhost:8000/helloworld/"+color+"-color").
			Header("X-Stain-Color", color).Send(context.Background())
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
