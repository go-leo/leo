package main

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/statusx"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"net"
	"net/http"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	handler := helloworld.NewGreeterHttpServerHandler(NewGreeterService())
	server := http.Server{Handler: handler}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type GreeterService struct {
}

func (g GreeterService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return nil, statusx.Failed.With(
		statusx.Message("some thing wrong"),
		statusx.QuotaFailure(&errdetails.QuotaFailure{
			Violations: []*errdetails.QuotaFailure_Violation{
				{Subject: fmt.Sprintf("name:%s", request.Name), Description: "Limit one greeting per person"},
			},
		}),
		statusx.HttpBody(wrapperspb.Bytes(errorx.Ignore(protojson.Marshal(&helloworld.CodeMessage{Code: 4040001, Message: "some thing wrong"})))),
	)
}

func NewGreeterService() helloworld.GreeterService {
	return &GreeterService{}
}
