package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

var (
	addr = flag.String("addr", "localhost:60051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	// success
	mdw := jwtx.Client("test-kid", []byte("jwt_key_secret"), jwt.SigningMethodHS256)
	client := helloworld.NewGreeterHttpClient(*addr, httptransportx.WithMiddleware(mdw))
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	mdw = jwtx.Client("test-kid", []byte("jwt_key_wrong_secret"), jwt.SigningMethodHS256)
	client = helloworld.NewGreeterHttpClient(*addr, httptransportx.WithMiddleware(mdw))
	r, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}
