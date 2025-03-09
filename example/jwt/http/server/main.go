package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"log"
)

var (
	port = flag.Int("port", 60051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	mdw := jwtx.Server(
		func(token *jwt.Token) (interface{}, error) { return []byte("jwt_key_secret"), nil },
		jwt.SigningMethodHS256,
	)
	router := mux.NewRouter()
	router = helloworld.AppendGreeterHttpServerRoutes(router, &server{}, httptransportx.Middleware(mdw))
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))
	if err := leo.NewApp(leo.Runner(httpSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
