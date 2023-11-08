package auth

import (
	"context"
	"encoding/base64"
	"google.golang.org/grpc/credentials"
	"strings"
)

var _ credentials.PerRPCCredentials = BasicPerRPCCredentials{}

type BasicPerRPCCredentials struct {
	Username string
	Password string
	Security bool
}

func (b BasicPerRPCCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(b.Username+":"+b.Password))}, nil
}

func (b BasicPerRPCCredentials) RequireTransportSecurity() bool { return b.Security }

func BasicAuthorizer(username string, password string) Authorizer {
	return DefaultAuthorizer(func(ctx context.Context, scheme, credentials string) bool {
		if !strings.EqualFold(scheme, "Basic") {
			return false
		}
		if credentials != base64.StdEncoding.EncodeToString([]byte(username+":"+password)) {
			return false
		}
		return true
	})
}
