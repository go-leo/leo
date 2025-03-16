package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/retryx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"log"
)

func main() {
	httpCli := helloworld.NewGreeterHttpClient(
		"localhost:60051",
		httptransportx.WithMiddleware(
			retryx.Middleware(retryx.MaxAttempts(3), retryx.Backoff(retryx.Exponential())),
		),
	)
	Call(httpCli)
}

func Call(grpcCli helloworld.GreeterService) {
	r, err := grpcCli.SayHello(context.Background(), &helloworld.HelloRequest{Name: "retry"})
	if err != nil {
		st, ok := statusx.From(err)
		if ok {
			log.Printf("could not greet: %v, retryInfo: %v", err, st.RetryInfo())
		}
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
