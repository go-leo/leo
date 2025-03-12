# 健康检查
Leo支持gRPC和HTTP健康检查方式，它支持自定义其他系统（比如redis、mysql等）的健康检查。

# 用法
## 创建grpc和Http服务
```go
	// http server
	client := helloworld.NewGreeterGrpcClient("localhost:50051", grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	router := helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), client)
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))

	// grpc server
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{}))

```

## 创建Actuator服务，并且添加健康检查路由
```go
	actuatorRouter := mux.NewRouter()
	// 添加健康检查路由
	actuatorRouter = healthx.Append(actuatorRouter)
	actuatorSrv := actuator.NewServer(16060, actuatorRouter)
```

## 启动App
```go
	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv, actuatorSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
```

## 访问健康检查接口
全部
```
curl -v 'http://localhost:16060/health'
```

```
> GET /health HTTP/1.1
> Host: localhost:16060
> User-Agent: curl/8.4.0
> Accept: */*
> 
< HTTP/1.1 503 Service Unavailable
< Date: Wed, 12 Mar 2025 08:26:32 GMT
< Content-Length: 97
< Content-Type: text/plain; charset=utf-8
< 
{"status":"NOT_SERVING","components":{"grpc":"SERVING","http":"SERVING"}}
```

grpc
```
curl -v --location 'http://localhost:16060/health/grpc'
```
```
> GET /health/grpc HTTP/1.1
> Host: localhost:16060
> User-Agent: curl/8.4.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 12 Mar 2025 08:27:40 GMT
< Content-Length: 21
< Content-Type: text/plain; charset=utf-8
< 
{"status":"SERVING"}
```

http
```
curl -v 'http://localhost:16060/health/http' 
```
```
> GET /health/http HTTP/1.1
> Host: localhost:16060
> User-Agent: curl/8.4.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Wed, 12 Mar 2025 08:28:16 GMT
< Content-Length: 21
< Content-Type: text/plain; charset=utf-8
< 
{"status":"SERVING"}
```

# 自定义健康检查
```go
// 自定义健康检查
func customHealthChecker() {
	// 创建一个自定义的健康检查器，传入一个key
	checker := healthx.NewChecker("custom")
	rand, err := randx.NewChaCha8()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			time.Sleep(time.Duration(rand.Int64N(int64(time.Second))))
			if rand.Int()%2 == 0 {
				// 恢复
				checker.Resume()
			} else {
				// 关闭
				checker.Shutdown()
			}
		}
	}()
	// 注册健康检查器
	healthx.RegisterChecker(checker)
}
```
查询接口
```go
curl -v 'http://localhost:16060/health/custom' 
```
```go
> GET /health/custom HTTP/1.1
> Host: localhost:16060
> User-Agent: curl/8.4.0
> Accept: */*
> 
< HTTP/1.1 503 Service Unavailable
< Date: Wed, 12 Mar 2025 08:30:13 GMT
< Content-Length: 25
< Content-Type: text/plain; charset=utf-8
< 
{"status":"NOT_SERVING"}
```

具体代码见[health](../example/health)