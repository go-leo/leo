package basicx

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"google.golang.org/grpc/metadata"
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
		{"Isn't valid with nil header", nil, want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Isn't valid with non-string header", 42, want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Isn't valid without authHeader", "", want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Isn't valid for wrong user", makeAuthString("wrong-user", requiredPassword), want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Isn't valid for wrong password", makeAuthString(requiredUser, "wrong-password"), want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Is valid for correct creds", makeAuthString(requiredUser, requiredPassword), want{true, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			ctx = transportx.InjectName(ctx, httpx.HttpServer)
			ctx = context.WithValue(ctx, httptransport.ContextKeyRequestAuthorization, tt.authHeader)

			result, err := Middleware(requiredUser, requiredPassword, realm)(passedValidation)(ctx, nil)
			if result != tt.want.result || !errorx.Equals(err, tt.want.err) {
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
		{"Isn't valid with nil header", nil, want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Isn't valid with non-string header", 42, want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Isn't valid without authHeader", "", want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Isn't valid for wrong user", makeAuthString("wrong-user", requiredPassword), want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Isn't valid for wrong password", makeAuthString(requiredUser, "wrong-password"), want{nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()}},
		{"Is valid for correct creds", makeAuthString(requiredUser, requiredPassword), want{true, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			ctx = transportx.InjectName(ctx, grpcx.GrpcServer)
			ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("authorization", convx.ToString(tt.authHeader)))

			result, err := Middleware(requiredUser, requiredPassword, realm)(passedValidation)(ctx, nil)
			if result != tt.want.result || !errorx.Equals(err, tt.want.err) {
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
