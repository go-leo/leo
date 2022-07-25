package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-leo/leo/runner/net/http/internal/codec"
)

// 错误处理，把grpc的错误信息包装成http的错误处理形式
func errorHandler(ctx context.Context, c *gin.Context, marshaller codec.Codec, err error) {
	_ = c.Error(err)
	// error转成grpc.status(兼容grpc形式的错误处理)
	s := status.Convert(err)
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
	c.Data(HTTPStatusFromGRPCCode(s.Code()), marshaller.ContentType(), buf)

	// 设置Trailer
	if requestAcceptsTrailers(c) {
		handleTrailerMetadata(c, metadata)
	}
}

// HTTPStatusFromGRPCCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromGRPCCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
