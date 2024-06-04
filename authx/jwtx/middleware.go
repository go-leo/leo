package jwtx

import (
	"context"
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
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
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

				metadata := metadatax.New()
				metadata.Set(authKey, fmt.Sprintf("%s%s", prefix, tokenString))
				ctx = metadatax.NewOutgoingContext(ctx, metadata)

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
	ErrMissMetadata = statusx.InvalidArgument("missing metadata").Err()

	// ErrInvalidToken denotes a token was not able to be parsed.
	ErrInvalidToken = statusx.Unauthenticated("invalid token").Err()

	// ErrTokenInvalid denotes a token was not able to be validated.
	ErrTokenInvalid = statusx.Unauthenticated("JWT was invalid").Err()

	// ErrTokenExpired denotes a token's expire header (exp) has since passed.
	ErrTokenExpired = statusx.Unauthenticated("JWT is expired").Err()

	// ErrTokenMalformed denotes a token was not formatted as a JWT.
	ErrTokenMalformed = statusx.Unauthenticated("JWT is malformed").Err()

	// ErrTokenNotActive denotes a token's not before header (nbf) is in the
	// future.
	ErrTokenNotActive = statusx.Unauthenticated("token is not valid yet").Err()

	// ErrUnexpectedSigningMethod denotes a token was signed with an unexpected
	// signing method.
	ErrUnexpectedSigningMethod = statusx.Unauthenticated("unexpected signing method").Err()
)

// NewParser creates a new JWT parsing middleware, specifying a
// jwt.Keyfunc interface, the signing method and the claims type to be used. NewParser
// adds the resulting claims to endpoint context or returns error on invalid token.
// Particularly useful for servers.
func NewParser(keyFunc jwt.Keyfunc, method jwt.SigningMethod, claimsFactory ClaimsFactory) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
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
					return nil, ErrInvalidToken
				}

				// Parse takes the token string and a function for looking up the
				// key. The latter is especially useful if you use multiple keys
				// for your application.  The standard is to use 'kid' in the head
				// of the token to identify which key to use, but the parsed token
				// (head and claims) is provided to the callback, providing
				// flexibility.
				token, err := jwt.ParseWithClaims(tokenString, claimsFactory.Factory.New(), func(token *jwt.Token) (interface{}, error) {
					// Don't forget to validate the alg is what you expect:
					if token.Method != method {
						return nil, ErrUnexpectedSigningMethod
					}

					return keyFunc(token)
				})
				if err != nil {
					if e, ok := err.(*jwt.ValidationError); ok {
						switch {
						case e.Errors&jwt.ValidationErrorMalformed != 0:
							// Token is malformed
							return nil, ErrTokenMalformed
						case e.Errors&jwt.ValidationErrorExpired != 0:
							// Token is expired
							return nil, ErrTokenExpired
						case e.Errors&jwt.ValidationErrorNotValidYet != 0:
							// Token is not active yet
							return nil, ErrTokenNotActive
						case e.Inner != nil:
							// report e.Inner
							return nil, e.Inner
						}
						// We have a ValidationError but have no specific Go kit error for it.
						// Fall through to return original error.
					}
					return nil, err
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
