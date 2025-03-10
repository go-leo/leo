# Basic Auth

`authx/basicx` provides basic auth middleware

## Usage
### gRPC
#### Server
```go
func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc1.NewServer()
	endpoints := helloworld.NewGreeterEndpoints(
		NewGreeterService(),
		basicx.Middleware("soyacen", "123456", "basic auth example"),
	)
	transports := helloworld.NewGreeterGrpcServerTransports(endpoints)
	service := helloworld.NewGreeterGrpcServer(transports)
	helloworld.RegisterGreeterServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

#### Client
```go
package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// success
	mdw := jwtx.Client("test-kid", []byte("jwt_key_secret"))
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	mdw = jwtx.Client("test-kid", []byte("wrong_jwt_key_secret"))
	client = helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	r, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}
```

### Http
#### Server
```go
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	endpoints := helloworld.NewGreeterEndpoints(
		NewGreeterService(),
		basicx.Middleware("soyacen", "123456", "basic auth example"),
	)
	transports := helloworld.NewGreeterHttpServerTransports(endpoints)
	handler := helloworld.NewGreeterHttpServerHandler(transports)
	server := http.Server{Handler: handler}
	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
#### Client
```go

func main() {
	transports := helloworld.NewGreeterHttpClientTransports("http", "127.0.0.1:8080")

	// ok
	client := helloworld.NewGreeterHttpClient(
		transports,
		basicx.Middleware("soyacen", "123456", "basic auth example"),
	)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterHttpClient(
		transports,
		basicx.Middleware("soyacen", "654321", "basic auth example"),
	)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
```

## Example
[basic](..%2F..%2Fexample%2Fcmd%2Ffeatures%2Fauth%2Fbasic)