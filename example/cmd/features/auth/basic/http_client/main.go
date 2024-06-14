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
	endpoints := helloworld.NewGreeterClientEndpoints(
		transports,
		basicx.Middleware("soyacen", "123456", "basic auth example"),
	)
	client := helloworld.NewGreeterHttpClient(endpoints)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	endpoints = helloworld.NewGreeterClientEndpoints(
		transports,
		basicx.Middleware("soyacen", "654321", "basic auth example"),
	)
	client = helloworld.NewGreeterHttpClient(endpoints)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
