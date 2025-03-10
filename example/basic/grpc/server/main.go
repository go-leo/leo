package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"log"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// basic 中间件
	mdw := basicx.Server("ubuntu", "mint")
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	service := helloworld.NewGreeterGrpcServer(&server{}, grpctransportx.Middleware(mdw))
	helloworld.RegisterGreeterServer(grpcSrv, service)
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
