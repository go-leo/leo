module codeup.aliyun.com/qimao/leo/leo/config/resource/file

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.12
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.12
	github.com/fsnotify/fsnotify v1.6.0
	github.com/stretchr/testify v1.8.4

)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/derekparker/trie v0.0.0-20230829180723-39f4de51ef7d // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace codeup.aliyun.com/qimao/leo/leo => ../../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../../internal/gox
