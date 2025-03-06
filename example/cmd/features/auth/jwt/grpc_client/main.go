package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/middleware/authx/jwtx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/golang-jwt/jwt/v4"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// ok
	client := helloworld.NewGreeterGrpcClient(":9090",
		grpcx.DialOptions(grpc1.WithTransportCredentials(insecure.NewCredentials())),
		grpcx.Middleware(jwtx.NewSigner("kid", []byte("jwt_key_secret"), jwt.SigningMethodHS256, jwt.MapClaims{"user": "go-leo"})),
	)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterGrpcClient(":9090",
		grpcx.DialOptions(grpc1.WithTransportCredentials(insecure.NewCredentials())),
		grpcx.Middleware(jwtx.NewSigner("kid", []byte("jwt_key_wrong_secret"), jwt.SigningMethodHS256, jwt.MapClaims{"user": "go-leo"})),
	)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
