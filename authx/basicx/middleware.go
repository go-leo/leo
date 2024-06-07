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
	httpstatus "github.com/go-leo/leo/v3/statusx/http"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"strings"
)

const (
	prefix  = "Basic "
	authKey = "authorization"
)

func Middleware(requiredUser, requiredPassword, realm string) endpoint.Middleware {
	requiredUserBytes := toHashSlice([]byte(requiredUser))
	requiredPasswordBytes := toHashSlice([]byte(requiredPassword))
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			name, ok := transportx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			switch name {
			case grpcx.GrpcServer, httpx.HttpServer:
				return handleIncoming(ctx, request, next, requiredUserBytes, requiredPasswordBytes, realm)
			case grpcx.GrpcClient, httpx.HttpClient:
				return handleOutgoing(ctx, request, next, requiredUser, requiredPassword)
			}
			return next(ctx, request)
		}
	}
}

var (
	ErrMissMetadata         = statusx.ErrInvalidArgument.WithMessage("missing metadata")
	ErrInvalidAuthorization = statusx.ErrUnauthenticated.WithMessage("invalid authorization")
)

func handleIncoming(ctx context.Context, request any, next endpoint.Endpoint, requiredUserBytes []byte, requiredPasswordBytes []byte, realm string) (any, error) {
	md, ok := metadatax.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissMetadata
	}

	givenUser, givenPassword, ok := parseAuthorization(md.Values(authKey))
	if !ok {
		return nil, ErrInvalidAuthorization.WithHttpHeader(&httpstatus.Header{Key: "WWW-Authenticate", Value: fmt.Sprintf(`Basic realm=%q`, realm)})
	}

	givenUserBytes := toHashSlice(givenUser)
	givenPasswordBytes := toHashSlice(givenPassword)

	if subtle.ConstantTimeCompare(givenUserBytes, requiredUserBytes) == 0 ||
		subtle.ConstantTimeCompare(givenPasswordBytes, requiredPasswordBytes) == 0 {
		return nil, ErrInvalidAuthorization.WithHttpHeader(&httpstatus.Header{Key: "WWW-Authenticate", Value: fmt.Sprintf(`Basic realm=%q`, realm)})
	}
	// Continue execution of handler after ensuring a valid token.
	return next(ctx, request)
}

func handleOutgoing(ctx context.Context, request any, next endpoint.Endpoint, requiredUser, requiredPassword string) (any, error) {
	metadata := metadatax.Pairs(authKey, fmt.Sprintf("%s%s", prefix, base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", requiredUser, requiredPassword)))))
	ctx = metadatax.AppendToOutgoingContext(ctx, metadata)
	return next(ctx, request)
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
