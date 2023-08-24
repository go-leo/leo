package context

import (
	"context"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestStreamClientInterceptor_ContextValue(t *testing.T) {
	// Create a sample context function that adds a value
	contextFunc := func(ctx context.Context) context.Context {
		return context.WithValue(ctx, "key", "value")
	}

	// Create a mock Stream with a context
	mockStream := &mockClientStream{ctx: context.Background()}

	// Create a mock handler function
	mockHandler := func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		// Test if the context value is correctly passed to the handler
		val := ctx.Value("key")
		if val == nil || val != "value" {
			t.Errorf("Context value not passed to the handler")
		}
		return mockStream, nil
	}

	// Test the interceptor
	interceptor := StreamClientInterceptor(contextFunc)
	_, _ = interceptor(context.Background(), nil, nil, "", mockHandler)
}

func TestUnaryClientInterceptor_ContextTimeout(t *testing.T) {
	// Create a sample context function that sets a timeout
	contextFunc := func(ctx context.Context) context.Context {
		ctx, _ = context.WithTimeout(ctx, time.Millisecond*100) // Set a short timeout
		return ctx
	}

	// Create a mock unary invoker function
	mockInvoker := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
		// Wait for a longer time than the context timeout
		time.Sleep(time.Millisecond * 200)

		// Test if the context timeout is effective
		if ctx.Err() == nil {
			t.Errorf("Context timeout not effective")
		} else if ctx.Err() != context.DeadlineExceeded {
			t.Errorf("Expected context timeout error, got %v", ctx.Err())
		}
		return nil
	}

	// Test the interceptor
	interceptor := UnaryClientInterceptor(contextFunc)
	_ = interceptor(context.Background(), "", nil, nil, nil, mockInvoker)
}

// Mock for grpc.ClientStream
type mockClientStream struct {
	grpc.ClientStream
	ctx context.Context
}

func (m *mockClientStream) Context() context.Context {
	return m.ctx
}
