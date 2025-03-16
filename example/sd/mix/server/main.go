package main

import (
	"context"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/runner"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		client := helloworld.NewGreeterGrpcClient(
			"consul://localhost:8500/leo.example.sd.grpc?dc=dc1",
			grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
			grpctransportx.WithInstancerBuilder(consulx.Builder{}),
			grpctransportx.WithBalancerFactory(lbx.RandomFactory{Seed: time.Now().Unix()}),
		)
		var runners []runner.Runner
		for i := 0; i < 10; i++ {
			httpSrv := httpserverx.NewServer(
				helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), client),
				httpserverx.Instance("consul://localhost:8500/leo.example.sd.http?dc=dc1"),
				httpserverx.RegistrarBuilder(consulx.Builder{}),
			)
			runners = append(runners, httpSrv)
		}
		return leo.NewApp(leo.Runner(runners...)).Run(context.Background())
	})
	eg.Go(func() error {
		var runners []runner.Runner
		for i := 0; i < 10; i++ {
			grpcSrv := grpcserverx.NewServer(
				grpcserverx.Instance("consul://localhost:8500/leo.example.sd.grpc?dc=dc1"),
				grpcserverx.RegistrarBuilder(consulx.Builder{}),
			)
			helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{i: convx.ToString(i)}))
			runners = append(runners, grpcSrv)
		}
		return leo.NewApp(leo.Runner(runners...)).Run(context.Background())
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	i string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.i + " Hello " + in.GetName()}, nil
}
