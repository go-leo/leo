# package authx/jwtx

`package authx/jwtx` provides jwt auth middleware

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
		jwtx.NewParser(
			func(token *jwt.Token) (interface{}, error) { return []byte("jwt_key_secret"), nil },
			jwt.SigningMethodHS256,
			jwtx.ClaimsFactory{Factory: jwtx.MapClaimsFactory{}},
		),
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
		jwtx.NewSigner("kid", []byte("jwt_key_secret"), jwt.SigningMethodHS256, jwt.MapClaims{"user": "go-leo"}),
	)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterGrpcClient(
		transports,
		jwtx.NewSigner("kid", []byte("jwt_key_wrong_secret"), jwt.SigningMethodHS256, jwt.MapClaims{"user": "go-leo"}),
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
		jwtx.NewParser(
			func(token *jwt.Token) (interface{}, error) { return []byte("jwt_key_secret"), nil },
			jwt.SigningMethodHS256,
			jwtx.ClaimsFactory{Factory: jwtx.MapClaimsFactory{}},
		),
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

```

## Example
[jwt](..%2F..%2Fexample%2Fcmd%2Ffeatures%2Fauth%2Fjwt)