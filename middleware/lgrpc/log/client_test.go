package log

import (
	"codeup.aliyun.com/qimao/leo/leo/log"
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestUnaryClientInterceptor(t *testing.T) {
	ctx := context.Background()
	method := "/test.Service/TestMethod"
	req := "test request"
	reply := "test reply"
	cc := &grpc.ClientConn{}
	opts := []grpc.CallOption{}

	invoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		return nil
	}

	loggerFactory := func(ctx context.Context) log.Logger {
		// slog.New(slog.LevelAdapt(log.LevelDebug))
		return log.L()
	}

	interceptor := UnaryClientInterceptor(loggerFactory)

	// 模拟拦截器的调用
	err := interceptor(ctx, method, req, reply, cc, invoker, opts...)
	t.Log(err)
}
