module codeup.aliyun.com/qimao/leo/leo/stream/kafka

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.20
	github.com/confluentinc/confluent-kafka-go/v2 v2.1.1
	github.com/go-leo/gox v0.0.0-20230616023204-abcd5dbca361
	github.com/stretchr/testify v1.8.4
	github.com/ugorji/go v1.2.11
	github.com/ugorji/go/codec v1.2.11
)

require (
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.20 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20240112132812-db7319d0e0e3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/utils v0.0.0-20230505201702-9f6742963106 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo => ../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../internal/gox
