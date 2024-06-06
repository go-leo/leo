package basicx

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/errorx"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"testing"

	httptransport "github.com/go-kit/kit/transport/http"
)

func TestHttpWithBasicAuth(t *testing.T) {
	requiredUser := "test-user"
	requiredPassword := "test-pass"
	realm := "test realm"

	type want struct {
		result interface{}
		err    *statusx.Error
	}
	tests := []struct {
		name       string
		authHeader interface{}
		want       want
	}{
		{"Isn't valid with nil header", nil, want{nil, statusx.ErrUnauthenticated}},
		{"Isn't valid with non-string header", 42, want{nil, statusx.ErrUnauthenticated}},
		{"Isn't valid without authHeader", "", want{nil, statusx.ErrUnauthenticated}},
		{"Isn't valid for wrong user", makeAuthString("wrong-user", requiredPassword), want{nil, statusx.ErrUnauthenticated}},
		{"Isn't valid for wrong password", makeAuthString(requiredUser, "wrong-password"), want{nil, statusx.ErrUnauthenticated}},
		{"Is valid for correct creds", makeAuthString(requiredUser, requiredPassword), want{true, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			ctx = transportx.InjectName(ctx, httpx.HttpServer)
			md := metadatax.New()
			md.Set("Authorization", convx.ToString(tt.authHeader))
			ctx = metadatax.NewIncomingContext(ctx, md)
			ctx = context.WithValue(ctx, httptransport.ContextKeyRequestAuthorization, tt.authHeader)

			result, err := Middleware(requiredUser, requiredPassword, realm)(passedValidation)(ctx, nil)
			if result != tt.want.result || !tt.want.err.Equals(err) {
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
		{"Isn't valid with nil header", nil, want{nil, statusx.ErrUnauthenticated}},
		{"Isn't valid with non-string header", 42, want{nil, statusx.ErrUnauthenticated}},
		{"Isn't valid without authHeader", "", want{nil, statusx.ErrUnauthenticated}},
		{"Isn't valid for wrong user", makeAuthString("wrong-user", requiredPassword), want{nil, statusx.ErrUnauthenticated}},
		{"Isn't valid for wrong password", makeAuthString(requiredUser, "wrong-password"), want{nil, statusx.ErrUnauthenticated}},
		{"Is valid for correct creds", makeAuthString(requiredUser, requiredPassword), want{true, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			ctx = transportx.InjectName(ctx, grpcx.GrpcServer)
			md := metadatax.New()
			md.Set("authorization", convx.ToString(tt.authHeader))
			ctx = metadatax.NewIncomingContext(ctx, md)

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
