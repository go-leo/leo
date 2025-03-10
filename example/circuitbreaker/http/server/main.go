package main

import (
	"context"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/gox/mathx/randx/v2"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/gorilla/mux"
	"log"
	"math/rand/v2"
)

// server is used to implement helloworld.GreeterServer.
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
	// basic 中间件
	router := mux.NewRouter()
	router = helloworld.AppendGreeterHttpServerRoutes(router, &server{r: errorx.Ignore(randx.NewPCG())})
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))
	if err := leo.NewApp(leo.Runner(httpSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
