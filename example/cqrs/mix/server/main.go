package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/cqrs/v1"
	"github.com/go-leo/leo/v3/example/cqrs/handler"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		client := cqrs.NewCqrsGrpcClient(
			"localhost:50051",
			grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		)
		httpSrv := httpserverx.NewServer(
			cqrs.AppendCqrsHttpServerRoutes(mux.NewRouter(), client),
			httpserverx.Port(60051),
		)
		return leo.NewApp(leo.Runner(httpSrv)).Run(context.Background())
	})
	eg.Go(func() error {
		cqrsService, err := cqrs.NewCqrsCqrsService(handler.NewCommandHandler(), handler.NewCommandEmptyHandler(), handler.NewQueryHandler(), handler.NewQueryOneOfHandler())
		if err != nil {
			panic(err)
		}
		grpcSrv := grpcserverx.NewServer(
			grpcserverx.Port(50051),
		)
		cqrs.RegisterCqrsServer(grpcSrv, cqrs.NewCqrsGrpcServer(cqrsService))
		return leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background())
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
