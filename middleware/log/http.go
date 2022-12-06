package log

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/v2/log"
)

func GinMiddleware(loggerFactory func(ctx context.Context) log.Logger, opts ...Option) gin.HandlerFunc {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	return func(c *gin.Context) {
		if _, ok := skipMap[c.FullPath()]; ok {
			c.Next()
			return
		}
		if loggerFactory == nil {
			c.Next()
			return
		}
		logger := loggerFactory(c.Request.Context())
		if logger == nil {
			c.Next()
			return
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
