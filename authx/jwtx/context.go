package jwtx

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
)

type jwtClaimsKey struct{}

func NewContentWithClaims(ctx context.Context, claims jwt.Claims) context.Context {
	return context.WithValue(ctx, jwtClaimsKey{}, claims)
}

func ClaimsFromContext(ctx context.Context) (jwt.Claims, bool) {
	v, ok := ctx.Value(jwtClaimsKey{}).(jwt.Claims)
	return v, ok
}

type jwtTokenKey struct{}

func NewContentWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, jwtTokenKey{}, token)
}

func TokenFromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(jwtTokenKey{}).(string)
	return v, ok
}
