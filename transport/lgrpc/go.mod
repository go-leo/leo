module codeup.aliyun.com/qimao/leo/leo/transport/lgrpc

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.11
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.11
	golang.org/x/sync v0.4.0
	google.golang.org/grpc v1.56.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo => ../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../internal/gox
