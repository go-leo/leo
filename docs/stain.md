# 服务发现
流量染色功能，是在服务发现功能的基础的一次升级，请求带有特定的标记（颜色）时，会从目标服务列表中，选取一个与请求标记最匹配的服务地址，并将该请求转发到这个服务。

服务端标记染色后，在服务注册时，会将此标记当做服务的标签，服务发现时，会根据标签进行服务选择。

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
	i     string
	color string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.color + " " + s.i + " Hello " + in.GetName()}, nil
}

func main() {

	var colors = []string{"red", "blue", "yellow", "black", "white"}

	var runners []runner.Runner
	for i := 0; i < 10; i++ {
		color := colors[i%len(colors)]
		grpcSrv := grpcserverx.NewServer(
			grpcserverx.Instance("consul://localhost:8500/leo.example.sd.grpc?dc=dc1"),
			grpcserverx.RegistrarBuilder(consulx.Builder{}),
			// 服务端标记(染色)
			grpcserverx.Stain(color),
		)
		helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{i: convx.ToString(i), color: color}))
		runners = append(runners, grpcSrv)
	}
	if err := leo.NewApp(leo.Runner(runners...)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
注意
* `grpcserverx.Stain` 给服务端标记(染色)。

## gRPC 客户端
```go
package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/stainx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand/v2"
	"time"
)

func main() {
	client := helloworld.NewGreeterGrpcClient(
		"consul://localhost:8500/leo.example.sd.grpc?dc=dc1",
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		grpctransportx.WithInstancerBuilder(consulx.Builder{}),
		grpctransportx.WithBalancerFactory(lbx.RandomFactory{Seed: time.Now().Unix()}),
	)
	r := rand.New(rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Unix())))
	var colors = []string{"red", "blue", "yellow", "black", "white"}
	for i := 0; i < 100; i++ {
		ctx := context.Background()
		color := colors[r.IntN(len(colors))]
		// 客户端请求标记(染色)
		ctx = stainx.ColorInjector(ctx, color)
		r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: color})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}

```
注意
* `stainx.ColorInjector` 客户端请求标记(染色)

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

type server struct {
	i     string
	color string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.color + " " + s.i + " Hello " + in.GetName()}, nil
}

func main() {
	var colors = []string{"red", "blue", "yellow", "black", "white"}
	var runners []runner.Runner
	for i := 0; i < 10; i++ {
		color := colors[i%len(colors)]
		httpSrv := httpserverx.NewServer(
			helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), &server{i: convx.ToString(i), color: color}),
			httpserverx.Instance("consul://localhost:8500/leo.example.sd.http?dc=dc1"),
			httpserverx.RegistrarBuilder(consulx.Builder{}),
			// 服务端标记(染色)
			httpserverx.Stain(color),
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
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/stainx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"log"
	"math/rand/v2"
	"time"
)

func main() {
	client := helloworld.NewGreeterHttpClient(
		"consul://localhost:8500/leo.example.sd.http?dc=dc1",
		httptransportx.WithInstancerBuilder(consulx.Builder{}),
		httptransportx.WithBalancerFactory(lbx.RoundRobinFactory{}),
	)
	r := rand.New(rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Unix())))
	var colors = []string{"red", "blue", "yellow", "black", "white"}
	for i := 0; i < 100; i++ {
		ctx := context.Background()
		color := colors[r.IntN(len(colors))]
		// 客户端请求标记(染色)
		ctx = stainx.ColorInjector(ctx, color)
		r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: color})
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
	var colors = []string{"red", "blue", "yellow", "black", "white"}
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
			color := colors[i%len(colors)]
			httpSrv := httpserverx.NewServer(
				helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), client),
				httpserverx.Instance("consul://localhost:8500/leo.example.sd.http?dc=dc1"),
				httpserverx.RegistrarBuilder(consulx.Builder{}),
				// 服务端标记(染色)
				httpserverx.Stain(color),
			)
			runners = append(runners, httpSrv)
		}
		return leo.NewApp(leo.Runner(runners...)).Run(context.Background())
	})
	eg.Go(func() error {
		var runners []runner.Runner
		for i := 0; i < 10; i++ {
			color := colors[i%len(colors)]
			grpcSrv := grpcserverx.NewServer(
				grpcserverx.Instance("consul://localhost:8500/leo.example.sd.grpc?dc=dc1"),
				grpcserverx.RegistrarBuilder(consulx.Builder{}),
				// 服务端标记(染色)
				grpcserverx.Stain(color),
			)
			helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{i: convx.ToString(i), color: color}))
			runners = append(runners, grpcSrv)
		}
		return leo.NewApp(leo.Runner(runners...)).Run(context.Background())
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	i     string
	color string
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: s.color + " " + s.i + " Hello " + in.GetName()}, nil
}

```

# 代码
[stain](../example/stain)