module codeup.aliyun.com/qimao/leo/leo/schedule/cron

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.20
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.20
	github.com/robfig/cron/v3 v3.0.1
	golang.org/x/exp v0.0.0-20240112132812-db7319d0e0e3
)

replace codeup.aliyun.com/qimao/leo/leo => ../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../internal/gox
