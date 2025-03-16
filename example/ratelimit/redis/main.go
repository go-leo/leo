package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/ratelimitx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
	"log"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	client := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	// 令牌桶，每秒10个请求
	mdw := ratelimitx.Redis(redis_rate.NewLimiter(client), redis_rate.PerSecond(10), func(ctx context.Context) string { return "redis_limiter" })
	grpcServer := helloworld.NewGreeterGrpcServer(&server{},
		grpctransportx.Middleware(mdw),
	)
	helloworld.RegisterGreeterServer(grpcSrv, grpcServer)
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
