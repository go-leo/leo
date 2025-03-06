package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/middleware/authx/jwtx"
	"github.com/golang-jwt/jwt/v4"
	grpc1 "google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc1.NewServer()
	service := helloworld.NewGreeterGrpcServer(
		NewGreeterService(),
		jwtx.NewParser(
			func(token *jwt.Token) (interface{}, error) { return []byte("jwt_key_secret"), nil },
			jwt.SigningMethodHS256,
			jwtx.ClaimsFactory{Factory: jwtx.MapClaimsFactory{}},
		),
	)
	helloworld.RegisterGreeterServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
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
