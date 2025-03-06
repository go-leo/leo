package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/middleware/authx/basicx"
	"github.com/go-leo/leo/v3/transportx/httpx"
)

func main() {
	// ok
	client := helloworld.NewGreeterHttpClient(
		"127.0.0.1:8080",
		httpx.Middleware(basicx.Middleware("soyacen", "123456", "basic auth example")),
	)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterHttpClient(
		"127.0.0.1:8080",
		httpx.Middleware(basicx.Middleware("soyacen", "654321", "basic auth example")),
	)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
