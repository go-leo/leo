package ratelimiter

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Limiter interface {
	Limit(ctx context.Context) bool
}

var alwaysFalseLimiter = alwaysLimiter{v: false}

type alwaysLimiter struct {
	v bool
}

func (l alwaysLimiter) Limit(ctx context.Context) bool {
	return l.v
}

func Middleware(limiter Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Limit(c.Request.Context()) {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
	}
}
