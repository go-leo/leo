module codeup.aliyun.com/qimao/leo/leo/middleware/lgrpc/metric

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo/middleware/lgrpc/noop v0.0.1
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/metric v1.16.0
	google.golang.org/grpc v1.55.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo/middleware/lgrpc/noop => ../../../middleware/lgrpc/noop
