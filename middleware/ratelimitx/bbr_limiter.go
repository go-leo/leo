package ratelimitx

import (
	"github.com/go-kratos/aegis/ratelimit/bbr"
)

type bbrLimiter struct {
	limiter *bbr.BBR
}

func (limiter *bbrLimiter) Allow() bool {
	_, err := limiter.limiter.Allow()
	if err != nil {
		return false
	}
	return true
}
