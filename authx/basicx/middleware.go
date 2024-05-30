package basicx

import (
	"bytes"
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"github.com/go-kit/kit/auth/basic"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/transportx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

func Middleware(requiredUser, requiredPassword, realm string) endpoint.Middleware {
	requiredUserBytes := toHashSlice([]byte(requiredUser))
	requiredPasswordBytes := toHashSlice([]byte(requiredPassword))
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			name, ok := transportx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			if name == transportx.HttpServer {
				return basic.AuthMiddleware(requiredUser, requiredPassword, realm)(next)(ctx, request)
			}
			if name == transportx.GrpcServer {
				md, ok := metadata.FromIncomingContext(ctx)
				if !ok {
					return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
				}

				givenUser, givenPassword, ok := parseBasicAuth(md.Get("authorization"))
				if !ok {
					return nil, status.Errorf(codes.Unauthenticated, `invalid token, Basic realm=%q`, realm)
				}

				givenUserBytes := toHashSlice(givenUser)
				givenPasswordBytes := toHashSlice(givenPassword)

				if subtle.ConstantTimeCompare(givenUserBytes, requiredUserBytes) == 0 ||
					subtle.ConstantTimeCompare(givenPasswordBytes, requiredPasswordBytes) == 0 {
					return nil, status.Errorf(codes.Unauthenticated, `invalid token, Basic realm=%q`, realm)
				}

				// Continue execution of handler after ensuring a valid token.
				return next(ctx, request)
			}
			return next(ctx, request)
		}
	}
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ([]byte("Aladdin"), []byte("open sesame"), true).
func parseBasicAuth(authorization []string) (username, password []byte, ok bool) {
	if len(authorization) == 0 {
		return
	}
	auth := authorization[0]
	const prefix = "Basic "
	if !strings.HasPrefix(auth, prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}

	s := bytes.IndexByte(c, ':')
	if s < 0 {
		return
	}
	return c[:s], c[s+1:], true
}

// Returns a hash of a given slice.
func toHashSlice(s []byte) []byte {
	hash := sha256.Sum256(s)
	return hash[:]
}
