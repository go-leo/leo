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
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"strings"
)

const prefix = "Basic "

func Middleware(requiredUser, requiredPassword, realm string) endpoint.Middleware {
	requiredUserBytes := toHashSlice([]byte(requiredUser))
	requiredPasswordBytes := toHashSlice([]byte(requiredPassword))
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			name, ok := transportx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			if name == grpcx.GrpcServer {
				key := "authorization"
				return handleIncoming(ctx, request, next, key, requiredUserBytes, requiredPasswordBytes, realm)
			}
			if name == grpcx.GrpcClient {
				key := "authorization"
				return handleOutgoing(ctx, request, next, key, requiredUser, requiredPassword)
			}
			if name == httpx.HttpServer {
				key := "Authorization"
				return handleIncoming(ctx, request, next, key, requiredUserBytes, requiredPasswordBytes, realm)
			}
			if name == httpx.HttpClient {
				key := "Authorization"
				return handleOutgoing(ctx, request, next, key, requiredUser, requiredPassword)
			}
			return next(ctx, request)
		}
	}
}

func handleIncoming(ctx context.Context, request interface{}, next endpoint.Endpoint, key string, requiredUserBytes []byte, requiredPasswordBytes []byte, realm string) (interface{}, error) {
	md, ok := metadatax.FromIncomingContext(ctx)
	if !ok {
		return nil, statusx.InvalidArgument("missing metadata").Err()
	}

	givenUser, givenPassword, ok := parseBasicAuth(md.Values(key))
	if !ok {
		return nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()
	}

	givenUserBytes := toHashSlice(givenUser)
	givenPasswordBytes := toHashSlice(givenPassword)

	if subtle.ConstantTimeCompare(givenUserBytes, requiredUserBytes) == 0 ||
		subtle.ConstantTimeCompare(givenPasswordBytes, requiredPasswordBytes) == 0 {
		return nil, statusx.Unauthenticated(fmt.Sprintf(`invalid token, Basic realm=%q`, realm)).Err()
	}
	// Continue execution of handler after ensuring a valid token.
	return next(ctx, request)
}

// parseBasicAuth parses an HTTP Basic Authentication string.
// "Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns ([]byte("Aladdin"), []byte("open sesame"), true).
func parseBasicAuth(authorization []string) ([]byte, []byte, bool) {
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

func handleOutgoing(ctx context.Context, request interface{}, next endpoint.Endpoint, key string, requiredUser, requiredPassword string) (interface{}, error) {
	metadata := metadatax.New()
	metadata.Set(key, fmt.Sprintf("%s%s", prefix, base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", requiredUser, requiredPassword)))))
	ctx = metadatax.NewOutgoingContext(ctx, metadata)
	return next(ctx, request)
}
