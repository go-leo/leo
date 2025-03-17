package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/gorilla/mux"
	"log"
	"time"
)

var (
	port = flag.Int("port", 60051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	deadline, ok := ctx.Deadline()
	log.Printf("timeout: %v, %v", deadline, ok)
	time.Sleep(10 * time.Second)
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	router := helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), &server{})
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(*port))
	if err := leo.NewApp(leo.Runner(httpSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
