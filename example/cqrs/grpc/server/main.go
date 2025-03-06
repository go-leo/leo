package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/cqrs/v1"
	"github.com/go-leo/leo/v3/example/cqrs/cq"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"log"
)

func main() {
	cqrsService, err := cqrs.NewCqrsCqrsService(cq.NewCommandHandler(), cq.NewCommandEmptyHandler(), cq.NewQueryHandler(), cq.NewQueryOneOfHandler())
	if err != nil {
		panic(err)
	}
	grpcSrv := grpcserverx.NewServer(
		grpcserverx.Instance("consul://localhost:8500/leo.example.cqrs.grpc?dc=dc1"),
		grpcserverx.Builder(consulx.Builder{}),
	)
	cqrs.RegisterCqrsServer(grpcSrv, cqrs.NewCqrsGrpcServer(cqrsService))
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
