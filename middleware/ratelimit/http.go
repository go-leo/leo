package ratelimit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllowerGinMiddleware(limit Allower) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limit.Allow() {
			c.String(http.StatusTooManyRequests, ErrLimited.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func WaiterGinMiddleware(limit Waiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := limit.Wait(c.Request.Context()); err != nil {
			c.String(http.StatusTooManyRequests, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
