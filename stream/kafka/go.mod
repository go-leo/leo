module codeup.aliyun.com/qimao/leo/leo/stream/kafka

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.12
	github.com/confluentinc/confluent-kafka-go/v2 v2.1.1
	github.com/go-leo/gox v0.0.0-20230616023204-abcd5dbca361
	github.com/stretchr/testify v1.8.4
	github.com/ugorji/go v1.2.11
	github.com/ugorji/go/codec v1.2.11
)

require (
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.12 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	google.golang.org/grpc v1.59.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/utils v0.0.0-20230505201702-9f6742963106 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo => ../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../internal/gox
