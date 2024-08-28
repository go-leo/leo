package ratelimitx

import (
	"container/list"
	"sync"
	"time"
)

type SlideWindowLimiter struct {
	sync.RWMutex
	// queue 请求时间队列
	queue *list.List
	// interval 窗口大小
	interval time.Duration
	// rate 窗口内允许的最大请求数
	rate int
}

func NewSlideWindowLimiter(interval time.Duration, rate int) *SlideWindowLimiter {
	return &SlideWindowLimiter{
		queue:    list.New(),
		interval: interval,
		rate:     rate,
	}
}

func (limiter *SlideWindowLimiter) Allow() bool {
	now := time.Now()
	boundary := now.Add(-limiter.interval).UnixNano()

	limiter.Lock()
	defer limiter.Unlock()

	// 快速路径：直接判断是否允许请求
	if limiter.allow(now) {
		// 记住了请求的时间戳
		limiter.put(now)
		return true
	}

	// 慢速路径：清理过期数据后再判断
	limiter.clean(boundary)
	if limiter.allow(now) {
		// 记住了请求的时间戳
		limiter.put(now)
		return true
	}

	// 拒绝请求
	return false
}

// clean 清理过期的数据
func (limiter *SlideWindowLimiter) clean(boundary int64) {
	for timestamp := limiter.queue.Front(); timestamp != nil && timestamp.Value.(int64) < boundary; timestamp = limiter.queue.Front() {
		limiter.queue.Remove(timestamp)
	}
}

// allow 具体判断是否允许请求
func (limiter *SlideWindowLimiter) allow(now time.Time) bool {
	// 如果窗口内请求数小于窗口大小，则允许请求
	if limiter.queue.Len() < limiter.rate {
		return true
	}
	return false
}

func (limiter *SlideWindowLimiter) put(now time.Time) *list.Element {
	return limiter.queue.PushBack(now.UnixNano())
}
