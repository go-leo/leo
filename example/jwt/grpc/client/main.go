package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// success
	// jwt 中间件
	mdw := jwtx.Client([]byte("jwt_key_secret"))
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	// 向ctx中注入jwt信息
	ctx := jwtx.NewContentWithClaims(context.Background(), jwt.MapClaims{"user_id": "123456"})
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	// jwt 中间件
	mdw = jwtx.Client([]byte("wrong_jwt_key_secret"))
	client = helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	// 向ctx中注入jwt信息
	ctx = jwtx.NewContentWithClaims(context.Background(), jwt.MapClaims{"user_id": "123456"})
	r, err = client.SayHello(ctx, &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}
