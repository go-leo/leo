package basicx

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/statusx"
	"net/http"
	"strings"
)

const (
	prefix  = "Basic "
	authKey = "authorization"
)

func Client(user, password string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			tokenString := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, password)))
			metadata := metadatax.Pairs(authKey, fmt.Sprintf("%s%s", prefix, tokenString))
			ctx = metadatax.AppendOutgoingContext(ctx, metadata)
			return next(ctx, request)
		}
	}
}

func Server(user, password string) endpoint.Middleware {
	requiredUserBytes := toHashSlice([]byte(user))
	requiredPasswordBytes := toHashSlice([]byte(password))
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			md, ok := metadatax.FromIncomingContext(ctx)
			if !ok {
				return nil, statusx.InvalidArgument(statusx.Message("missing metadata"))
			}

			givenUser, givenPassword, ok := parseAuthorization(md.Values(authKey))
			if !ok {
				header := http.Header{"WWW-Authenticate": []string{fmt.Sprintf(`Basic realm=%q`, "Restricted")}}
				return nil, statusx.Unauthenticated(statusx.Headers(header), statusx.Message("invalid authorization"))
			}

			givenUserBytes := toHashSlice(givenUser)
			givenPasswordBytes := toHashSlice(givenPassword)

			if subtle.ConstantTimeCompare(givenUserBytes, requiredUserBytes) == 0 ||
				subtle.ConstantTimeCompare(givenPasswordBytes, requiredPasswordBytes) == 0 {
				header := http.Header{"WWW-Authenticate": []string{fmt.Sprintf(`Basic realm=%q`, "Restricted")}}
				return nil, statusx.Unauthenticated(statusx.Headers(header), statusx.Message("invalid authorization"))
			}
			// Continue execution of handler after ensuring a valid token.
			return next(ctx, request)
		}
	}
}

// parseAuthorization parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ([]byte("Aladdin"), []byte("open sesame"), true).
func parseAuthorization(authorization []string) ([]byte, []byte, bool) {
	if len(authorization) == 0 {
		return nil, nil, false
	}
	auth := authorization[0]
	if !strings.HasPrefix(auth, prefix) {
		return nil, nil, false
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return nil, nil, false
	}

	s := bytes.IndexByte(c, ':')
	if s < 0 {
		return nil, nil, false
	}
	return c[:s], c[s+1:], true
}

// Returns a hash of a given slice.
func toHashSlice(s []byte) []byte {
	hash := sha256.Sum256(s)
	return hash[:]
}
