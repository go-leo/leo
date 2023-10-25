package recovery

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware(handlers ...func(*gin.Context, any) error) gin.HandlerFunc {
	var handle func(*gin.Context, any) error
	if len(handlers) == 0 {
		handle = func(c *gin.Context, p any) error {
			return c.Error(fmt.Errorf("panic triggered: %+v", p))
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
					_ = handle(c, r)
					c.Abort()
				} else {
					_ = handle(c, r)
				}
			}
		}()
		c.Next()
	}
}

// 用于panic后自定义返回结构
func HandleRecoveryWithErr(c *gin.Context, p any) error {
	msg := "服务器异常:" + errorToString(p)
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg":  msg,
		"code": 500,
	})
	return c.Error(errors.New(msg))
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return fmt.Sprint(r)
	}
}
