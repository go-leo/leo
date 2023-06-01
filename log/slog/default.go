package slog

import "codeup.aliyun.com/qimao/leo/leo/log"

func init() {
	New(LevelAdapt(log.LevelDebug), JSON()).SetDefault()
}
