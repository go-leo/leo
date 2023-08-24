package auth

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"testing"
)

type MockAuthorizer struct {
	valid bool
}

func (a *MockAuthorizer) Authorize(ctx context.Context, fullMethodName string) (context.Context, error) {
	if a.valid {
		return ctx, nil
	}
	return ctx, status.Errorf(codes.Unauthenticated, "invalid token")
}

func TestUnaryServerInterceptor(t *testing.T) {
	mockAuthorizer := &MockAuthorizer{valid: true}
	authorizerFunc := AuthorizerFunc(mockAuthorizer.Authorize)
	interceptor := UnaryServerInterceptor(authorizerFunc)

	_, err := interceptor(context.Background(), nil, &grpc.UnaryServerInfo{}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return "response", nil
	})

	assert.Nil(t, err, "Expected no error")
}

func TestStreamServerInterceptor(t *testing.T) {
	mockAuthorizer := &MockAuthorizer{valid: true}
	authorizerFunc := AuthorizerFunc(mockAuthorizer.Authorize)
	interceptor := StreamServerInterceptor(authorizerFunc)

	err := interceptor(nil, &mockServerStream{}, &grpc.StreamServerInfo{}, func(srv interface{}, ss grpc.ServerStream) error {
		return nil
	})

	assert.Nil(t, err, "Expected no error")
}

type mockServerStream struct {
	grpc.ServerStream
}

func (m *mockServerStream) Context() context.Context {
	return context.Background()
}

func TestDefaultAuthorizer_ValidToken(t *testing.T) {
	valid := func(ctx context.Context, scheme, credentials string) bool {
		return scheme == "Bearer" && credentials == "validToken"
	}
	authorizer := DefaultAuthorizer(valid)

	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer validToken"))
	newCtx, err := authorizer.Authorize(ctx, "method")
	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, ctx, newCtx, "Expected contexts to be equal")
}

func TestDefaultAuthorizer_InvalidToken(t *testing.T) {
	valid := func(ctx context.Context, scheme, credentials string) bool {
		return scheme == "Bearer" && credentials == "validToken"
	}
	authorizer := DefaultAuthorizer(valid)

	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer invalidToken"))
	newCtx, err := authorizer.Authorize(ctx, "method")
	assert.Error(t, err, "Expected error")
	assert.Contains(t, err.Error(), "invalid token", "Expected error message to contain 'invalid token'")
	assert.Equal(t, codes.Unauthenticated, status.Code(err), "Expected Unauthenticated error code")
	assert.Equal(t, ctx, newCtx, "Expected contexts to be equal")
}

func TestDefaultAuthorizer_MissingAuthorization(t *testing.T) {
	valid := func(ctx context.Context, scheme, credentials string) bool {
		return scheme == "Bearer" && credentials == "validToken"
	}
	authorizer := DefaultAuthorizer(valid)

	ctx := context.Background()
	newCtx, err := authorizer.Authorize(ctx, "method")
	assert.Error(t, err, "Expected error")
	assert.Contains(t, err.Error(), "missing metadata", "Expected error message to contain 'missing metadata'")
	assert.Equal(t, codes.InvalidArgument, status.Code(err), "Expected InvalidArgument error code")
	assert.Equal(t, ctx, newCtx, "Expected contexts to be equal")
}

func TestUnaryServerInterceptor_AuthorizerSrv(t *testing.T) {
	// Create a mock authorizer
	mockAuthorizer := &MockAuthorizer{valid: true}

	// Create a server with Authorizer interface
	server := struct {
		Authorizer
	}{mockAuthorizer}

	interceptor := UnaryServerInterceptor(nil)

	_, err := interceptor(context.Background(), nil, &grpc.UnaryServerInfo{Server: server}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return "response", nil
	})

	assert.Nil(t, err, "Expected no error")
}

func TestStreamServerInterceptor_AuthorizerSrv(t *testing.T) {
	// Create a mock authorizer
	mockAuthorizer := &MockAuthorizer{valid: true}

	// Create a server with Authorizer interface
	server := struct {
		Authorizer
	}{mockAuthorizer}

	interceptor := StreamServerInterceptor(nil)

	err := interceptor(server, &mockServerStream{}, &grpc.StreamServerInfo{}, func(srv interface{}, ss grpc.ServerStream) error {
		return nil
	})

	assert.Nil(t, err, "Expected no error")
}

func TestUnaryServerInterceptor_NoAuthorizer(t *testing.T) {
	interceptor := UnaryServerInterceptor(nil)

	_, err := interceptor(context.Background(), nil, &grpc.UnaryServerInfo{}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return "response", nil
	})

	assert.Nil(t, err, "Expected no error")
}

func TestStreamServerInterceptor_NoAuthorizer(t *testing.T) {
	interceptor := StreamServerInterceptor(nil)

	err := interceptor(nil, &mockServerStream{}, &grpc.StreamServerInfo{}, func(srv interface{}, ss grpc.ServerStream) error {
		return nil
	})

	assert.Nil(t, err, "Expected no error")
}
