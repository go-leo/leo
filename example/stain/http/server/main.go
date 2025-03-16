package main

import (
	"context"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/runner"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/gorilla/mux"
	"log"
)

type server struct {
	i     string
	color string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.color + " " + s.i + " Hello " + in.GetName()}, nil
}

func main() {
	var colors = []string{"red", "blue", "yellow", "black", "white"}
	var runners []runner.Runner
	for i := 0; i < 10; i++ {
		color := colors[i%len(colors)]
		httpSrv := httpserverx.NewServer(
			helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), &server{i: convx.ToString(i), color: color}),
			httpserverx.Instance("consul://localhost:8500/leo.example.sd.http?dc=dc1"),
			httpserverx.RegistrarBuilder(consulx.Builder{}),
			// 服务端标记(染色)
			httpserverx.Stain(color),
		)
		runners = append(runners, httpSrv)
	}
	if err := leo.NewApp(leo.Runner(runners...)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
