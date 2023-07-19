module codeup.aliyun.com/qimao/leo/leo/stream/amqp

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.6
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.6
	github.com/rabbitmq/amqp091-go v1.8.1
	github.com/stretchr/testify v1.8.3
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1 // indirect
	golang.org/x/sync v0.3.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/utils v0.0.0-20230505201702-9f6742963106 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo => ../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../internal/gox
