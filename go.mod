module codeup.aliyun.com/qimao/leo/leo

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.9
	github.com/derekparker/trie v0.0.0-20221221181808-1424fce0c981
	github.com/mitchellh/mapstructure v1.5.0
	github.com/spf13/cast v1.5.1
	github.com/stretchr/testify v1.8.3
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	golang.org/x/sync v0.3.0
	k8s.io/utils v0.0.0-20230505201702-9f6742963106
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ./internal/gox
