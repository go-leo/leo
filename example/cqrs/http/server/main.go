package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/cqrs/v1"
	"github.com/go-leo/leo/v3/example/cqrs/handler"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	cqrsService, err := cqrs.NewCqrsCqrsService(handler.NewCommandHandler(), handler.NewCommandEmptyHandler(), handler.NewQueryHandler(), handler.NewQueryOneOfHandler())
	if err != nil {
		panic(err)
	}
	httpSrv := httpserverx.NewServer(
		cqrs.AppendCqrsHttpServerRoutes(mux.NewRouter(), cqrsService),
		httpserverx.Port(60051),
	)
	if err := leo.NewApp(leo.Runner(httpSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
