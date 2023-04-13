package log

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

func Middleware(loggerFactory func(ctx context.Context) log.Logger, opts ...Option) gin.HandlerFunc {
	o := new(options)
	o.apply(opts...)
	o.init()
	return func(c *gin.Context) {
		if loggerFactory == nil {
			return
		}
		logger := loggerFactory(c.Request.Context())
		if logger == nil {
			return
		}
		for _, skip := range o.Skips {
			if skip(c) {
				return
			}
		}
		// 开始时间
		startTime := time.Now()
		// 处理中间件和业务逻辑
		c.Next()
		builder := NewFieldBuilder().
			System("http.server").
			StartTime(startTime).
			Deadline(c.Request.Context()).
			Path(c.Request.URL.Path).
			Method(c.Request.Method).
			Status(http.StatusText(c.Writer.Status())).
			Error(c.Errors.ByType(gin.ErrorTypePrivate).String()).
			Latency(time.Since(startTime))
		if len(c.Errors) > 0 {
			logger.ErrorF(builder.Build()...)
		} else {
			logger.InfoF(builder.Build()...)
		}
	}
}
