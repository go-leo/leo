package rewrite

import (
	"bytes"

	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data,omitempty" `
	Msg  string `json:"msg,omitempty"`
	// TraceId string      `json:"trace_id"` // 用于链路追踪 id
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		c.Writer = blw.ResponseWriter
		if !c.GetBool("panic") {
			var responseBody interface{}
			if err := json.Unmarshal(blw.body.Bytes(), &responseBody); err == nil {
				obj := Response{}
				obj.Code = 200
				obj.Data = responseBody
				updatedBody, _ := json.Marshal(obj)
				c.Writer.Write(updatedBody)
			}
		} else {
			c.Writer.Write(blw.body.Bytes())
		}
	}
}
