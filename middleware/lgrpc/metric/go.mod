module codeup.aliyun.com/qimao/leo/leo/middleware/lgrpc/metric

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo/middleware/lgrpc/noop v0.0.1
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/metric v1.16.0
	google.golang.org/grpc v1.59.0
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	go.opentelemetry.io/otel/trace v1.16.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo/middleware/lgrpc/noop => ../../../middleware/lgrpc/noop
