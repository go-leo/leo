package ratelimitx

import (
	"container/list"
	"sync"
	"time"
)

type OverLoadProtection struct {
	sync.RWMutex
	// queue 请求时间队列
	queue *list.List
	// interval 窗口大小
	interval time.Duration
	// rate 窗口内允许的最大请求数
	rate int
}


func NewOverLoadProtection(interval time.Duration, rate int) *OverLoadProtection {
	return &OverLoadProtection{
		queue:    list.New(),
		interval: interval,
		rate:     rate,
	}
}