package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	client := helloworld.NewGreeterGrpcClient("localhost:50051", grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	router := helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), client)
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))

	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{}))

	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	md, ok := metadatax.FromIncomingContext(ctx)
	if ok {
		log.Printf("token: %v", md.Get("token"))
	}
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
