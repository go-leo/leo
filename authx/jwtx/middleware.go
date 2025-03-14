package jwtx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

const (
	prefix  = "Bearer "
	authKey = "authorization"
)

func Client(key []byte) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			claims, ok := ClaimsFromContext(ctx)
			if !ok {
				return next(ctx, request)
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
			tokenString, err := token.SignedString(key)
			if err != nil {
				return nil, err
			}
			ctx = NewContentWithToken(ctx, tokenString)
			metadata := metadatax.Pairs(authKey, fmt.Sprintf("%s%s", prefix, tokenString))
			ctx = metadatax.AppendOutgoingContext(ctx, metadata)
			return next(ctx, request)
		}
	}
}

func Server(keyFunc jwt.Keyfunc) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			md, ok := metadatax.FromIncomingContext(ctx)
			if !ok {
				return nil, statusx.InvalidArgument(statusx.Message("missing metadata"))
			}
			tokenString, ok := parseAuthorization(md.Values(authKey))
			if !ok {
				return nil, statusx.Unauthenticated(statusx.Message("invalid authorization"))
			}
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				if token.Method != jwt.SigningMethodHS512 {
					return nil, statusx.Unauthenticated(statusx.Message("unexpected signing method"))
				}
				return keyFunc(token)
			})
			if err != nil {
				return nil, statusx.Unauthenticated(statusx.Message(err.Error()))
			}
			if !token.Valid {
				return nil, statusx.Unauthenticated(statusx.Message("JWT was invalid"))
			}
			ctx = NewContentWithToken(ctx, tokenString)
			ctx = NewContentWithClaims(ctx, token.Claims)
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
