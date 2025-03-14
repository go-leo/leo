# pprof
pprof是go语言官方提供的性能分析工具，它提供了多种分析方式

# 使用pprof
```go
func main() {
	// http server
	client := helloworld.NewGreeterGrpcClient("localhost:50051", grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	router := helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), client)
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))

	// grpc server
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{}))

	actuatorRouter := mux.NewRouter()
	// 添加pprof路由
	actuatorRouter = pprofx.Append(actuatorRouter)
	actuatorSrv := actuator.NewServer(16060, actuatorRouter)

	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv, actuatorSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

# 代码
[pprof](../example/pprof)