package server

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/grpclog"

	"github.com/go-leo/leo/runner/net/http/internal/codec"
	"github.com/go-leo/leo/runner/net/http/internal/util"
)

// HandlerFunc将xxx.leo.pb.go里的xxx_HTTP_Handler包装gin的handler
func HandlerFunc(cli any, desc *ServiceDesc, methodDesc *MethodDesc) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取请求的Content-Type，默认是application/json
		contentCodec := codec.GetCodec(util.GetContentType(c.Request.Header))
		// 获取客户端Accept类型，默认是application/json
		acceptCodec := codec.GetCodec(util.GetAcceptType(c.Request.Header))

		metadata := new(Metadata)
		ctx := NewContextWithMetadata(c.Request.Context(), metadata)

		ctx, cancel, err := _GRPCTimeout(c)
		if err != nil {
			errorHandler(ctx, c, acceptCodec, err)
			return
		}
		defer cancel()

		// 解析http.Header，转成metadata，创建OutgoingContext
		ctx, err = newOutgoingContext(ctx, c)
		if err != nil {
			errorHandler(ctx, c, acceptCodec, err)
			return
		}

		// 读请求的body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			errorHandler(ctx, c, acceptCodec, err)
			return
		}
		// request解码的方法
		in := methodDesc.RequestConstructor()
		if err := contentCodec.Unmarshal(body, in); err != nil {
			errorHandler(ctx, c, acceptCodec, err)
			return
		}

		// 调用handler，此handler就是xxx.leo.pb.go里的xxx_HTTP_Handler
		reply, appErr := methodDesc.Handler(cli, ctx, in)
		if appErr != nil {
			errorHandler(ctx, c, acceptCodec, appErr)
			return
		}
		// 编码reply
		data, err := acceptCodec.Marshal(reply)
		if err != nil {
			grpclog.Infof("Marshal error: %v", err)
			errorHandler(ctx, c, acceptCodec, err)
			return
		}

		// 设置Header
		handleHeaderMetadata(c, metadata)
		if requestAcceptsTrailers(c) {
			handleForwardResponseTrailerHeader(c, metadata)
			c.Writer.Header().Set("Transfer-Encoding", "chunked")
		}

		// 响应
		c.Data(http.StatusOK, acceptCodec.ContentType(), data)

		// 设置Trailer
		if requestAcceptsTrailers(c) {
			handleTrailerMetadata(c, metadata)
		}
	}
}
