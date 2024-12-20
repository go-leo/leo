package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/middleware/authx/jwtx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/sdx/passthroughx"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	client, err := helloworld.NewGreeterHttpClientV2(
		"http",
		"127.0.0.1:8080",
		nil,
		[]endpoint.Middleware{jwtx.NewSigner("kid", []byte("jwt_key_secret"), jwt.SigningMethodHS256, jwt.MapClaims{"user": "go-leo"})},
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
		[]endpoint.Middleware{jwtx.NewSigner("kid", []byte("jwt_key_wrong_secret"), jwt.SigningMethodHS256, jwt.MapClaims{"user": "go-leo"})},
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
