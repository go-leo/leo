package internal

import (
	"strings"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	InstrumentationName = "github.com/go-leo/leo/v2"
)

var (
	RPCSystemGRPCServer = semconv.RPCSystemKey.String("grpc.server")
	RPCSystemGRPCClient = semconv.RPCSystemKey.String("grpc.client")
	RPCSystemHTTPServer = semconv.RPCSystemKey.String("http.server")
	RPCSystemHTTPClient = semconv.RPCSystemKey.String("http.client")
)

func ParseFullMethod(fullMethod string) []attribute.KeyValue {
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

func GRPCType() attribute.KeyValue {
	return attribute.Key("grpc.type").String("unary")
}

func ParseError(err error) attribute.KeyValue {
	if err != nil {
		s, _ := status.FromError(err)
		return attribute.Key("grpc.code").String(s.String())
	}
	return attribute.Key("grpc.code").String(codes.OK.String())
}
