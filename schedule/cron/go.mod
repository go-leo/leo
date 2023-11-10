module codeup.aliyun.com/qimao/leo/leo/schedule/cron

go 1.20

require (
	codeup.aliyun.com/qimao/leo/leo v0.0.17
	codeup.aliyun.com/qimao/leo/leo/internal/gox v0.0.17
	github.com/robfig/cron/v3 v3.0.1
	golang.org/x/exp v0.0.0-20231108232855-2478ac86f678
)

replace codeup.aliyun.com/qimao/leo/leo => ../..

replace codeup.aliyun.com/qimao/leo/leo/internal/gox => ../../internal/gox
