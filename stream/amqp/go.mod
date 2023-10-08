module codeup.aliyun.com/qimao/leo/leo/stream/amqp

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.10
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.10
	github.com/rabbitmq/amqp091-go v1.8.1
	github.com/spf13/cast v1.5.1
	github.com/stretchr/testify v1.8.3
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
	golang.org/x/sync v0.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/utils v0.0.0-20230505201702-9f6742963106 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo => ../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../internal/gox
