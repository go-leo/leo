package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/middleware/authx/basicx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
)

func main() {
	// ok
	client, err := helloworld.NewGreeterHttpClientV2(
		"http",
		"127.0.0.1:8080",
		nil,
		[]endpoint.Middleware{basicx.Middleware("soyacen", "123456", "basic auth example")},
		passthroughx.Factory{},
		nil,
		logx.L(),
		lbx.RoundRobinFactory{},
	)
	if err != nil {
		panic(err)
	}
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client, err = helloworld.NewGreeterHttpClientV2(
		"http",
		"127.0.0.1:8080",
		nil,
		[]endpoint.Middleware{basicx.Middleware("soyacen", "654321", "basic auth example")},
		passthroughx.Factory{},
		nil,
		logx.L(),
		lbx.RoundRobinFactory{},
	)
	if err != nil {
		panic(err)
	}
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
