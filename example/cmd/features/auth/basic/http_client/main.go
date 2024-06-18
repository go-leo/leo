package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld"
)

func main() {
	transports, err := helloworld.NewGreeterHttpClientTransports("127.0.0.1:8080", "http")
	if err != nil {
		panic(err)
	}

	// ok
	client := helloworld.NewGreeterHttpClient(
		transports,
		basicx.Middleware("soyacen", "123456", "basic auth example"),
	)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterHttpClient(
		transports,
		basicx.Middleware("soyacen", "654321", "basic auth example"),
	)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
