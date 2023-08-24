package context

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"testing"
)

func TestUnaryServerInterceptor_ContextValue(t *testing.T) {
	// Create a sample context function that adds a value
	contextFunc := func(ctx context.Context) context.Context {
		return context.WithValue(ctx, "key", "value")
	}

	// Create a mock unary handler function
	mockHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
		// Test if the context value is correctly passed to the handler
		val := ctx.Value("key")
		if val == nil || val != "value" {
			return nil, errors.New("Context value not passed to the handler")
		}
		return "response", nil
	}

	// Test the interceptor
	interceptor := UnaryServerInterceptor(contextFunc)
	_, err := interceptor(context.Background(), nil, &grpc.UnaryServerInfo{}, mockHandler)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestStreamServerInterceptor(t *testing.T) {
	// Create a sample context function that modifies the context
	contextFunc := func(ctx context.Context) context.Context {
		return context.WithValue(ctx, "key", "modified_value")
	}

	// Create a mock handler function
	mockHandler := func(srv interface{}, stream grpc.ServerStream) error {
		ctx := stream.Context()
		val := ctx.Value("key")
		if val != "modified_value" {
			t.Errorf("Expected modified_value in context, got %v", val)
		}
		return nil
	}

	// Create a mock ServerStream
	mockServerStream := &mockStream{ctx: contextFunc(context.Background())}

	// Test the interceptor
	interceptor := StreamServerInterceptor(contextFunc)
	err := interceptor(nil, mockServerStream, &grpc.StreamServerInfo{}, mockHandler)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

type mockStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (m *mockStream) Context() context.Context {
	return m.ctx
}
