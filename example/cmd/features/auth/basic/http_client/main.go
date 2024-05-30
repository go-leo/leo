package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld"
)

func main() {
	flag.Parse()
	client := helloworld.NewGreeterHttpClient(helloworld.NewGreeterHttpClientTransports("http", "127.0.0.1:8080"))
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
