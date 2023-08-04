package recovery

import (
	"errors"
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
			return nil
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
					c.Error(r.(error)) // nolint: errcheck
					c.Abort()
				} else {
					handle(c, r)
				}
			}
		}()
		c.Next()
	}
}

// 用于panic后自定义返回结构
func HandleRecovery(c *gin.Context, err any) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg":  "服务器异常:" + errorToString(err),
		"code": 500,
	})
}

// 用于panic后自定义返回结构
func HandleRecoveryWithErr(c *gin.Context, err any) error {
	msg := "服务器异常:" + errorToString(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg":  msg,
		"code": 500,
	})
	return errors.New(msg)
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
