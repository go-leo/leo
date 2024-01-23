module codeup.aliyun.com/qimao/leo/leo/schedule/cron

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.1.0
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.1.0
	github.com/robfig/cron/v3 v3.0.1
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a
)

replace codeup.aliyun.com/qimao/leo/leo => ../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../internal/gox
