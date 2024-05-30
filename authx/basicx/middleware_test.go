package basicx

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-kit/kit/auth/basic"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/leo/v3/transportx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"testing"

	httptransport "github.com/go-kit/kit/transport/http"
)

func TestHttpWithBasicAuth(t *testing.T) {
	requiredUser := "test-user"
	requiredPassword := "test-pass"
	realm := "test realm"

	type want struct {
		result interface{}
		err    error
	}
	tests := []struct {
		name       string
		authHeader interface{}
		want       want
	}{
		{"Isn't valid with nil header", nil, want{nil, basic.AuthError{realm}}},
		{"Isn't valid with non-string header", 42, want{nil, basic.AuthError{realm}}},
		{"Isn't valid without authHeader", "", want{nil, basic.AuthError{realm}}},
		{"Isn't valid for wrong user", makeAuthString("wrong-user", requiredPassword), want{nil, basic.AuthError{realm}}},
		{"Isn't valid for wrong password", makeAuthString(requiredUser, "wrong-password"), want{nil, basic.AuthError{realm}}},
		{"Is valid for correct creds", makeAuthString(requiredUser, requiredPassword), want{true, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			ctx = transportx.InjectName(ctx, transportx.HttpServer)
			ctx = context.WithValue(ctx, httptransport.ContextKeyRequestAuthorization, tt.authHeader)

			result, err := Middleware(requiredUser, requiredPassword, realm)(passedValidation)(ctx, nil)
			if result != tt.want.result || err != tt.want.err {
				t.Errorf("WithBasicAuth() = result: %v, err: %v, want result: %v, want error: %v", result, err, tt.want.result, tt.want.err)
			}
		})
	}
}

func TestGrpcWithBasicAuth(t *testing.T) {
	requiredUser := "test-user"
	requiredPassword := "test-pass"
	realm := "test realm"

	type want struct {
		result interface{}
		err    error
	}
	tests := []struct {
		name       string
		authHeader interface{}
		want       want
	}{
		{"Isn't valid with nil header", nil, want{nil, status.Errorf(codes.Unauthenticated, `invalid token, Basic realm=%q`, realm)}},
		{"Isn't valid with non-string header", 42, want{nil, status.Errorf(codes.Unauthenticated, `invalid token, Basic realm=%q`, realm)}},
		{"Isn't valid without authHeader", "", want{nil, status.Errorf(codes.Unauthenticated, `invalid token, Basic realm=%q`, realm)}},
		{"Isn't valid for wrong user", makeAuthString("wrong-user", requiredPassword), want{nil, status.Errorf(codes.Unauthenticated, `invalid token, Basic realm=%q`, realm)}},
		{"Isn't valid for wrong password", makeAuthString(requiredUser, "wrong-password"), want{nil, status.Errorf(codes.Unauthenticated, `invalid token, Basic realm=%q`, realm)}},
		{"Is valid for correct creds", makeAuthString(requiredUser, requiredPassword), want{true, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			ctx = transportx.InjectName(ctx, transportx.GrpcServer)
			ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("authorization", convx.ToString(tt.authHeader)))

			result, err := Middleware(requiredUser, requiredPassword, realm)(passedValidation)(ctx, nil)
			if result != tt.want.result || err.Error() != tt.want.err.Error() {
				t.Errorf("WithBasicAuth() = result: %v, err: %v, want result: %v, want error: %v", result, err, tt.want.result, tt.want.err)
			}
		})
	}
}

func makeAuthString(user string, password string) string {
	data := []byte(fmt.Sprintf("%s:%s", user, password))
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(data))
}

func passedValidation(ctx context.Context, request interface{}) (response interface{}, err error) {
	return true, nil
}
