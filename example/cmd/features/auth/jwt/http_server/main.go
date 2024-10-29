package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/middleware/authx/jwtx"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	router := helloworld.AppendGreeterHttpRoutes(
		mux.NewRouter(),
		NewGreeterService(),
		jwtx.NewParser(
			func(token *jwt.Token) (interface{}, error) { return []byte("jwt_key_secret"), nil },
			jwt.SigningMethodHS256,
			jwtx.ClaimsFactory{Factory: jwtx.MapClaimsFactory{}},
		),
	)
	server := http.Server{Handler: router}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GreeterService struct {
}

func (g GreeterService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "hi " + request.GetName()}, nil
}

func NewGreeterService() helloworld.GreeterService {
	return &GreeterService{}
}
