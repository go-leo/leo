package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/cqrs/v1"
	"github.com/go-leo/leo/v3/example/cqrs/handler"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"log"
)

func main() {
	cqrsService, err := cqrs.NewCqrsCqrsService(
		handler.NewCommandHandler(),
		handler.NewCommandEmptyHandler(),
		handler.NewQueryHandler(),
		handler.NewQueryOneOfHandler(),
	)
	if err != nil {
		panic(err)
	}
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	cqrs.RegisterCqrsServer(grpcSrv, cqrs.NewCqrsGrpcServer(cqrsService))
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
