package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"log"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	md, ok := metadatax.FromIncomingContext(ctx)
	if ok {
		log.Printf("token: %v", md.Get("token"))
	}
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(*port))
	helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{}))
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
