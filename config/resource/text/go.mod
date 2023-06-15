module codeup.aliyun.com/qimao/leo/leo/config/resource/text

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.1
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
)

require codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.1

require github.com/mitchellh/mapstructure v1.5.0 // indirect

replace codeup.aliyun.com/qimao/leo/leo => ../../../

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../../internal/gox
