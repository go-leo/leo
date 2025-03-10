package main

import (
	"context"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/mathx/randx/v2"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/statusx"
	"log"
	"math/rand/v2"
)

type server struct {
	r *rand.Rand
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	f := s.r.Float32()
	f = f + 0.0
	if f > 0.9 {
		return nil, statusx.Internal()
	}
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	service := helloworld.NewGreeterGrpcServer(&server{r: errorx.Ignore(randx.NewPCG())})
	helloworld.RegisterGreeterServer(grpcSrv, service)
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
