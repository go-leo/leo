# package authx/jwtx

`package authx/basicx` provides basic auth middleware

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

func main() {
	conn, err := grpc1.Dial(":9090", grpc1.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	transports := helloworld.NewGreeterGrpcClientTransports(conn)
	
	// ok
	client := helloworld.NewGreeterGrpcClient(
		transports,
		basicx.Middleware("soyacen", "123456", "basic auth example"),
	)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterGrpcClient(
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