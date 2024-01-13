package metric

import (
	"strings"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	kInstrumentationName = "codeup.aliyun.com/qimao/leo/leo/middleware/lgrpc/metric"
)

var (
	vRPCSystemGRPCServer = semconv.RPCSystemKey.String("grpc.server")
	vRPCSystemGRPCClient = semconv.RPCSystemKey.String("grpc.client")
)

func parseFullMethod(fullMethod string) []attribute.KeyValue {
	name := strings.TrimLeft(fullMethod, "/")
	parts := strings.SplitN(name, "/", 2)
	if len(parts) != 2 {
		return []attribute.KeyValue(nil)
	}
	return []attribute.KeyValue{
		attribute.Key("grpc.service").String(parts[0]),
		attribute.Key("grpc.method").String(parts[1]),
	}
}

func gRPCType() attribute.KeyValue {
	return attribute.Key("grpc.type").String("unary")
}

func parseError(err error) attribute.KeyValue {
	var code codes.Code
	if s, ok := status.FromError(err); ok {
		code = s.Code()
	} else {
		code = status.FromContextError(err).Code()
	}
	return attribute.Key("grpc.code").String(code.String())
}
