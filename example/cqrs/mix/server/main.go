package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/cqrs/v1"
	"github.com/go-leo/leo/v3/example/cqrs/cq"
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
		client := cqrs.NewCqrsGrpcClient(
			"consul://localhost:8500/leo.example.cqrs.grpc?dc=dc1",
			grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
			grpctransportx.WithInstancerBuilder(consulx.Builder{}),
			grpctransportx.WithBalancerFactory(lbx.RandomFactory{Seed: time.Now().Unix()}),
		)
		httpSrv := httpserverx.NewServer(
			cqrs.AppendCqrsHttpServerRoutes(mux.NewRouter(), client),
			httpserverx.Instance("consul://localhost:8500/leo.example.cqrs.http?dc=dc1"),
			httpserverx.RegistrarBuilder(consulx.Builder{}),
		)
		return leo.NewApp(leo.Runner(httpSrv)).Run(context.Background())
	})
	eg.Go(func() error {
		cqrsService, err := cqrs.NewCqrsCqrsService(cq.NewCommandHandler(), cq.NewCommandEmptyHandler(), cq.NewQueryHandler(), cq.NewQueryOneOfHandler())
		if err != nil {
			panic(err)
		}
		grpcSrv := grpcserverx.NewServer(
			grpcserverx.Instance("consul://localhost:8500/leo.example.cqrs.grpc?dc=dc1"),
			grpcserverx.Builder(consulx.Builder{}),
		)
		cqrs.RegisterCqrsServer(grpcSrv, cqrs.NewCqrsGrpcServer(cqrsService))
		return leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background())
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
