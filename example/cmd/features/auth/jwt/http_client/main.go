package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/middleware/authx/jwtx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	client := helloworld.NewGreeterHttpClient(
		"127.0.0.1:8080",
		httpx.Middleware(jwtx.NewSigner("kid", []byte("jwt_key_secret"), jwt.SigningMethodHS256, jwt.MapClaims{"user": "go-leo"})),
	)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterHttpClient(
		"127.0.0.1:8080",
		httpx.Middleware(jwtx.NewSigner("kid", []byte("jwt_key_wrong_secret"), jwt.SigningMethodHS256, jwt.MapClaims{"user": "go-leo"})),
	)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
