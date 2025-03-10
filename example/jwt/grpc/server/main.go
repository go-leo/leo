package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	// 从上下文中获取jwt信息
	claims, _ := jwtx.ClaimsFromContext(ctx)
	fmt.Println(claims)
	// 从上下文中获取jwt token
	token, _ := jwtx.TokenFromContext(ctx)
	fmt.Println(token)
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// jwt 中间件
	mdw := jwtx.Server(func(token *jwt.Token) (interface{}, error) { return []byte("jwt_key_secret"), nil })
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	service := helloworld.NewGreeterGrpcServer(&server{}, grpctransportx.Middleware(mdw))
	helloworld.RegisterGreeterServer(grpcSrv, service)
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
