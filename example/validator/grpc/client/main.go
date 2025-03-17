package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/go-leo/leo/v3/validatorx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 添加 validator 中间件
	mdw := validatorx.Middleware()
	c := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		grpctransportx.WithMiddleware(mdw),
	)
	Call(c, "")
	Call(c, "leo")
}

func Call(c helloworld.GreeterService, name string) {
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
	} else {
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
