module codeup.aliyun.com/qimao/leo/leo

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.1.0
	github.com/derekparker/trie v0.0.0-20230829180723-39f4de51ef7d
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1
	github.com/spf13/cast v1.5.1
	github.com/stretchr/testify v1.8.4
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a
	golang.org/x/sync v0.6.0
	k8s.io/utils v0.0.0-20230505201702-9f6742963106
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ./internal/gox
