package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	"github.com/go-leo/leo/runner/net/http/internal/codec"
	"github.com/go-leo/leo/runner/net/http/internal/status"
)

// 错误处理，把grpc的错误信息包装成http的错误处理形式
func errorHandler(ctx context.Context, c *gin.Context, marshaller codec.Codec, err error) {
	_ = c.Error(err)
	// error转成grpc.status(兼容grpc形式的错误处理)
	s := grpcstatus.Convert(err)
	// grpc.status转成proto
	pb := s.Proto()
	if s.Code() == codes.Unauthenticated {
		c.Header("WWW-Authenticate", s.Message())
	}
	c.Writer.Header().Del("Trailer")
	c.Writer.Header().Del("Transfer-Encoding")

	// 对错误进行编码
	buf, e := marshaller.Marshal(pb)
	if e != nil {
		c.JSON(http.StatusInternalServerError, &spb.Status{Code: int32(codes.Internal), Message: "failed to marshal error message"})
		return
	}

	metadata, _ := MetadataFromContext(ctx)
	// 设置Header
	handleHeaderMetadata(c, metadata)
	if requestAcceptsTrailers(c) {
		handleForwardResponseTrailerHeader(c, metadata)
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
	}

	// 返回错误信息
	c.Data(status.HTTPStatusFromGRPCCode(s.Code()), marshaller.ContentType(), buf)

	// 设置Trailer
	if requestAcceptsTrailers(c) {
		handleTrailerMetadata(c, metadata)
	}
}
