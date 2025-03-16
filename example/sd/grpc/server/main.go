package main

import (
	"context"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/runner"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"log"
)

type server struct {
	i string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.i + " Hello " + in.GetName()}, nil
}

func main() {
	var runners []runner.Runner
	for i := 0; i < 10; i++ {
		grpcSrv := grpcserverx.NewServer(
			grpcserverx.Instance("consul://localhost:8500/leo.example.sd.grpc?dc=dc1"),
			grpcserverx.RegistrarBuilder(consulx.Builder{}),
		)
		helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{i: convx.ToString(i)}))
		runners = append(runners, grpcSrv)
	}
	if err := leo.NewApp(leo.Runner(runners...)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
