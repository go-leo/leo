package recovery

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-leo/gox/runtimex"
)

func Middleware(handlers ...func(*gin.Context, any) error) gin.HandlerFunc {
	var handle func(*gin.Context, any) error
	if len(handlers) == 0 {
		handle = func(c *gin.Context, p any) error {
			return fmt.Errorf("panic triggered: %+v, stack: %s", p, runtimex.Stack(0))
		}
	} else {
		handle = handlers[0]
	}
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := r.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(handle(c, r)) // nolint: errcheck
					c.Abort()
				} else {
					_ = c.AbortWithError(http.StatusInternalServerError, handle(c, r))
				}
			}
		}()
		c.Next()
	}
}

// 用于panic后自定义返回结构
func HandleRecovery(c *gin.Context, err any) {
	c.Set("panic", true)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "服务器异常:" + errorToString(err),
		"code": 500,
	})
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
