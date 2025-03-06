package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/cqrs/v1"
	"github.com/go-leo/leo/v3/example/cqrs/cq"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	cqrsService, err := cqrs.NewCqrsCqrsService(cq.NewCommandHandler(), cq.NewCommandEmptyHandler(), cq.NewQueryHandler(), cq.NewQueryOneOfHandler())
	if err != nil {
		panic(err)
	}
	httpSrv := httpserverx.NewServer(
		cqrs.AppendCqrsHttpServerRoutes(mux.NewRouter(), cqrsService),
		httpserverx.Instance("consul://localhost:8500/leo.example.cqrs.http?dc=dc1"),
		httpserverx.RegistrarBuilder(consulx.Builder{}),
	)
	if err := leo.NewApp(leo.Runner(httpSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
