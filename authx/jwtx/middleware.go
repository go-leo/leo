package jwtx

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

const (
	prefix  = "Bearer "
	authKey = "authorization"
)

// NewSigner creates a new JWT generating middleware, specifying key ID,
// signing string, signing method and the claims you would like it to contain.
// Tokens are signed with a Key ID header (kid) which is useful for determining
// the key to use for parsing.
func NewSigner(kid string, key []byte, method jwt.SigningMethod, claims jwt.Claims) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			name, ok := transportx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			switch name {
			case grpcx.GrpcClient, httpx.HttpClient:
				token := jwt.NewWithClaims(method, claims)
				token.Header["kid"] = kid

				// Sign and get the complete encoded token as a string using the secret
				tokenString, err := token.SignedString(key)
				if err != nil {
					return nil, err
				}
				ctx = NewContentWithToken(ctx, tokenString)

				metadata := metadatax.Pairs(authKey, fmt.Sprintf("%s%s", prefix, tokenString))
				ctx = metadatax.AppendToOutgoingContext(ctx, metadata)

				return next(ctx, request)
			}
			return next(ctx, request)
		}
	}
}

type claimsFactory interface {
	New() jwt.Claims
}

// ClaimsFactory is a factory for jwt.Claims.
// Useful in NewParser middleware.
type ClaimsFactory struct {
	Factory claimsFactory
}

// MapClaimsFactory is a ClaimsFactory that returns
// an empty jwt.MapClaims.
type MapClaimsFactory struct{}

func (MapClaimsFactory) New() jwt.Claims { return jwt.MapClaims{} }

// StandardClaimsFactory is a ClaimsFactory that returns
// an empty jwt.StandardClaims.
type StandardClaimsFactory struct{}

func (StandardClaimsFactory) New() jwt.Claims {
	return &jwt.StandardClaims{}
}

var (
	// ErrMissMetadata denotes a metadata was not found.
	ErrMissMetadata = statusx.ErrInvalidArgument.WithMessage("missing metadata")

	// ErrInvalidAuthorization denotes a token was not able to be parsed.
	ErrInvalidAuthorization = statusx.ErrUnauthenticated.WithMessage("invalid authorization")

	// ErrTokenInvalid denotes a token was not able to be validated.
	ErrTokenInvalid = statusx.ErrUnauthenticated.WithMessage("JWT was invalid")
)

// NewParser creates a new JWT parsing middleware, specifying a
// jwt.Keyfunc interface, the signing method and the claims type to be used. NewParser
// adds the resulting claims to endpoint context or returns error on invalid token.
// Particularly useful for servers.
func NewParser(keyFunc jwt.Keyfunc, method jwt.SigningMethod, claimsFactory ClaimsFactory) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			name, ok := transportx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			switch name {
			case grpcx.GrpcServer, httpx.HttpServer:
				md, ok := metadatax.FromIncomingContext(ctx)
				if !ok {
					return nil, ErrMissMetadata
				}

				tokenString, ok := parseAuthorization(md.Values(authKey))
				if !ok {
					return nil, ErrInvalidAuthorization
				}

				// Parse takes the token string and a function for looking up the
				// key. The latter is especially useful if you use multiple keys
				// for your application.  The standard is to use 'kid' in the head
				// of the token to identify which key to use, but the parsed token
				// (head and claims) is provided to the callback, providing
				// flexibility.
				token, err := jwt.ParseWithClaims(tokenString, claimsFactory.Factory.New(), func(token *jwt.Token) (any, error) {
					// Don't forget to validate the alg is what you expect:
					if token.Method != method {
						return nil, errors.New("unexpected signing method")
					}
					return keyFunc(token)
				})
				if err != nil {
					return nil, statusx.ErrUnauthenticated.WithMessage(err.Error())
				}

				if !token.Valid {
					return nil, ErrTokenInvalid
				}

				ctx = NewContentWithToken(ctx, tokenString)
				ctx = NewContentWithClaims(ctx, token.Claims)

				return next(ctx, request)
			}
			return next(ctx, request)
		}
	}
}

func parseAuthorization(authorization []string) (string, bool) {
	if len(authorization) == 0 {
		return "", false
	}
	auth := authorization[0]
	if !strings.HasPrefix(auth, prefix) {
		return "", false
	}
	return auth[len(prefix):], true
}
