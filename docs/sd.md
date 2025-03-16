# 服务发现
Leo 的服务发现功能基于 go-kit 的服务发现功能, 支持多种服务发现方式，包括 consul、etcdv3 等。可以快速扩展其他服务发现方式。

# 使用
## gRPC 服务端
```go
package main

import (
	"context"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/runner"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"log"
)

type server struct {
	i string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.i + " Hello " + in.GetName()}, nil
}

func main() {
	var runners []runner.Runner
	for i := 0; i < 10; i++ {
		grpcSrv := grpcserverx.NewServer(
			grpcserverx.Instance("consul://localhost:8500/leo.example.sd.grpc?dc=dc1"),
			grpcserverx.RegistrarBuilder(consulx.Builder{}),
		)
		helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{i: convx.ToString(i)}))
		runners = append(runners, grpcSrv)
	}
	if err := leo.NewApp(leo.Runner(runners...)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
注意
* 开启服务发现功能，需要配置 `grpcserverx.Instance` 和 `grpcserverx.Builder` 两个选项
* `grpcserverx.Instance` 是描述服务的Name,参考[gRPC Naming](https://github.com/grpc/grpc/blob/master/doc/naming.md)格式
* `grpcserverx.RegistrarBuilder` 配置服务注册的构建器

## gRPC 客户端
```go
package main

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	client := helloworld.NewGreeterGrpcClient(
		"consul://localhost:8500/leo.example.sd.grpc?dc=dc1",
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		grpctransportx.WithInstancerBuilder(consulx.Builder{}),
		grpctransportx.WithBalancerFactory(lbx.RandomFactory{Seed: time.Now().Unix()}),
	)
	for i := 0; i < 10; i++ {
		r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: randx.HexString(10)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
```
注意
* `grpctransportx.WithInstancerBuilder` 配置服务发现构建器
* `grpctransportx.WithBalancerFactory` 配置负载均衡器工厂

## HTTP 服务端
```go
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

// server is used to implement helloworld.GreeterServer.
type server struct {
	i string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.i + " Hello " + in.GetName()}, nil
}

func main() {
	var runners []runner.Runner
	for i := 0; i < 10; i++ {
		httpSrv := httpserverx.NewServer(
			helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), &server{i: convx.ToString(i)}),
			httpserverx.Instance("consul://localhost:8500/leo.example.sd.http?dc=dc1"),
			httpserverx.RegistrarBuilder(consulx.Builder{}),
		)
		runners = append(runners, httpSrv)
	}
	if err := leo.NewApp(leo.Runner(runners...)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
和 gRPC 服务端一样。

## HTTP 客户端
```go
package main

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"log"
)

func main() {
	client := helloworld.NewGreeterHttpClient(
		"consul://localhost:8500/leo.example.sd.http?dc=dc1",
		httptransportx.WithInstancerBuilder(consulx.Builder{}),
		httptransportx.WithBalancerFactory(lbx.RoundRobinFactory{}),
	)
	for i := 0; i < 100; i++ {
		r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: randx.HexString(10)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
```
和 gRPC 客户端一样。

## gRPC和HTTP合并
```go
package main

import (
	"context"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/runner"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		client := helloworld.NewGreeterGrpcClient(
			"consul://localhost:8500/leo.example.sd.grpc?dc=dc1",
			grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
			grpctransportx.WithInstancerBuilder(consulx.Builder{}),
			grpctransportx.WithBalancerFactory(lbx.RandomFactory{Seed: time.Now().Unix()}),
		)
		var runners []runner.Runner
		for i := 0; i < 10; i++ {
			httpSrv := httpserverx.NewServer(
				helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), client),
				httpserverx.Instance("consul://localhost:8500/leo.example.sd.http?dc=dc1"),
				httpserverx.RegistrarBuilder(consulx.Builder{}),
			)
			runners = append(runners, httpSrv)
		}
		return leo.NewApp(leo.Runner(runners...)).Run(context.Background())
	})
	eg.Go(func() error {
		var runners []runner.Runner
		for i := 0; i < 10; i++ {
			grpcSrv := grpcserverx.NewServer(
				grpcserverx.Instance("consul://localhost:8500/leo.example.sd.grpc?dc=dc1"),
				grpcserverx.RegistrarBuilder(consulx.Builder{}),
			)
			helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{i: convx.ToString(i)}))
			runners = append(runners, grpcSrv)
		}
		return leo.NewApp(leo.Runner(runners...)).Run(context.Background())
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	i string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.i + " Hello " + in.GetName()}, nil
}
```

# 代码
[sd](../example/sd)