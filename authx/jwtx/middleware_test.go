package jwtx

import (
	"context"
	"errors"
	"fmt"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/metadatax"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"github.com/golang-jwt/jwt/v4"
	"sync"
	"testing"
	"time"
)

var (
	kid            = "kid"
	key            = []byte("test_signing_key")
	myProperty     = "some value"
	method         = jwt.SigningMethodHS256
	invalidMethod  = jwt.SigningMethodRS256
	mapClaims      = jwt.MapClaims{"user": "go-kit"}
	standardClaims = jwt.StandardClaims{Audience: "go-kit"}
	// Signed tokens generated at https://jwt.io/
	signedKey         = "eyJhbGciOiJIUzI1NiIsImtpZCI6ImtpZCIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiZ28ta2l0In0.14M2VmYyApdSlV_LZ88ajjwuaLeIFplB8JpyNy0A19E"
	standardSignedKey = "eyJhbGciOiJIUzI1NiIsImtpZCI6ImtpZCIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJnby1raXQifQ.L5ypIJjCOOv3jJ8G5SelaHvR04UJuxmcBN5QW3m_aoY"
	customSignedKey   = "eyJhbGciOiJIUzI1NiIsImtpZCI6ImtpZCIsInR5cCI6IkpXVCJ9.eyJteV9wcm9wZXJ0eSI6InNvbWUgdmFsdWUiLCJhdWQiOiJnby1raXQifQ.s8F-IDrV4WPJUsqr7qfDi-3GRlcKR0SRnkTeUT_U-i0"
	invalidKey        = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.e30.vKVCKto-Wn6rgz3vBdaZaCBGfCBDTXOENSo_X2Gq7qA"
	malformedKey      = "malformed.jwt.token"
)

func signingValidator(t *testing.T, ctx context.Context, signer endpoint.Endpoint, expectedKey string) {
	resp, err := signer(ctx, struct{}{})
	if err != nil {
		t.Fatalf("Signer returned error: %s", err)
	}

	token, ok := TokenFromContext(resp.(context.Context))
	if !ok {
		t.Fatal("Token did not exist in context")
	}

	if token != expectedKey {
		t.Fatalf("JWTs did not match: expecting %s got %s", expectedKey, token)
	}
}

func TestNewSigner(t *testing.T) {

	ctx := transportx.InjectName(context.TODO(), httpx.HttpClient)
	testNewSigner(t, ctx)

	ctx = transportx.InjectName(context.TODO(), grpcx.GrpcClient)
	testNewSigner(t, ctx)

}

func testNewSigner(t *testing.T, ctx context.Context) {
	e := func(ctx context.Context, i interface{}) (interface{}, error) { return ctx, nil }
	signer := NewSigner(kid, key, method, mapClaims)(e)
	signingValidator(t, ctx, signer, signedKey)

	signer = NewSigner(kid, key, method, standardClaims)(e)
	signingValidator(t, ctx, signer, standardSignedKey)
}

func TestJWTParser(t *testing.T) {
	ctx := transportx.InjectName(context.TODO(), httpx.HttpServer)
	testJWTParser(t, ctx)

	ctx = transportx.InjectName(context.TODO(), grpcx.GrpcServer)
	testJWTParser(t, ctx)
}

func testJWTParser(t *testing.T, ctx context.Context) {
	e := func(ctx context.Context, i interface{}) (interface{}, error) { return ctx, nil }

	keys := func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}

	parser := NewParser(keys, method, ClaimsFactory{Factory: MapClaimsFactory{}})(e)

	// No Token is passed into the parser
	_, err := parser(ctx, struct{}{})
	if err == nil {
		t.Error("Parser should have returned an error")
	}

	if !errors.Is(err, ErrMissMetadata) {
		t.Errorf("unexpected error returned, expected: %s got: %s", ErrMissMetadata, err)
	}

	// Invalid Token is passed into the parser
	ctx = NewContentWithToken(ctx, invalidKey)
	_, err = parser(ctx, struct{}{})
	if err == nil {
		t.Error("Parser should have returned an error")
	}

	// Invalid Method is used in the parser
	badParser := NewParser(keys, invalidMethod, ClaimsFactory{Factory: MapClaimsFactory{}})(e)
	ctx = metadatax.NewIncomingContext(ctx, metadatax.Pairs("authorization", fmt.Sprintf("%s%s", prefix, signedKey)))
	_, err = badParser(ctx, struct{}{})
	if err == nil {
		t.Error("Parser should have returned an error")
	}

	if !errors.Is(err, ErrUnexpectedSigningMethod) {
		t.Errorf("unexpected error returned, expected: %s got: %s", ErrUnexpectedSigningMethod, err)
	}

	// Invalid key is used in the parser
	invalidKeys := func(token *jwt.Token) (interface{}, error) {
		return []byte("bad"), nil
	}

	badParser = NewParser(invalidKeys, method, ClaimsFactory{Factory: MapClaimsFactory{}})(e)
	ctx = metadatax.NewIncomingContext(ctx, metadatax.Pairs("authorization", fmt.Sprintf("%s%s", prefix, signedKey)))
	_, err = badParser(ctx, struct{}{})
	if err == nil {
		t.Error("Parser should have returned an error")
	}

	// Correct token is passed into the parser
	ctx = metadatax.NewIncomingContext(ctx, metadatax.Pairs("authorization", fmt.Sprintf("%s%s", prefix, signedKey)))
	ctx1, err := parser(ctx, struct{}{})
	if err != nil {
		t.Fatalf("Parser returned error: %s", err)
	}

	claims, ok := ClaimsFromContext(ctx1.(context.Context))
	if !ok {
		t.Fatal("Claims were not passed into context correctly")
	}
	cl := claims.(jwt.MapClaims)

	if cl["user"] != mapClaims["user"] {
		t.Fatalf("JWT Claims.user did not match: expecting %s got %s", mapClaims["user"], cl["user"])
	}

	// Test for malformed token error response
	parser = NewParser(keys, method, ClaimsFactory{Factory: StandardClaimsFactory{}})(e)
	ctx = metadatax.NewIncomingContext(ctx, metadatax.Pairs("authorization", fmt.Sprintf("%s%s", prefix, malformedKey)))
	ctx1, err = parser(ctx, struct{}{})
	if !errors.Is(err, ErrTokenMalformed) {
		t.Fatalf("Expected %+v, got %+v", ErrTokenMalformed, err)
	}

	// Test for expired token error response
	parser = NewParser(keys, method, ClaimsFactory{Factory: StandardClaimsFactory{}})(e)
	expired := jwt.NewWithClaims(method, jwt.StandardClaims{ExpiresAt: time.Now().Unix() - 100})
	token, err := expired.SignedString(key)
	if err != nil {
		t.Fatalf("Unable to Sign Token: %+v", err)
	}
	ctx = metadatax.NewIncomingContext(ctx, metadatax.Pairs("authorization", fmt.Sprintf("%s%s", prefix, token)))
	ctx1, err = parser(ctx, struct{}{})
	if !errors.Is(err, ErrTokenExpired) {
		t.Fatalf("Expected %+v, got %+v", ErrTokenExpired, err)
	}

	// Test for not activated token error response
	parser = NewParser(keys, method, ClaimsFactory{Factory: StandardClaimsFactory{}})(e)
	notactive := jwt.NewWithClaims(method, jwt.StandardClaims{NotBefore: time.Now().Unix() + 100})
	token, err = notactive.SignedString(key)
	if err != nil {
		t.Fatalf("Unable to Sign Token: %+v", err)
	}
	ctx = metadatax.NewIncomingContext(ctx, metadatax.Pairs("authorization", fmt.Sprintf("%s%s", prefix, token)))
	ctx1, err = parser(ctx, struct{}{})
	if !errors.Is(err, ErrTokenNotActive) {
		t.Fatalf("Expected %+v, got %+v", ErrTokenNotActive, err)
	}

	// test valid standard claims token
	parser = NewParser(keys, method, ClaimsFactory{Factory: StandardClaimsFactory{}})(e)
	ctx = metadatax.NewIncomingContext(ctx, metadatax.Pairs("authorization", fmt.Sprintf("%s%s", prefix, standardSignedKey)))
	ctx1, err = parser(ctx, struct{}{})
	if err != nil {
		t.Fatalf("Parser returned error: %s", err)
	}
	claims, ok = ClaimsFromContext(ctx1.(context.Context))
	if !ok {
		t.Fatal("Claims were not passed into context correctly")
	}
	stdCl := claims.(*jwt.StandardClaims)
	if !stdCl.VerifyAudience("go-kit", true) {
		t.Fatalf("JWT jwt.StandardClaims.Audience did not match: expecting %s got %s", standardClaims.Audience, stdCl.Audience)
	}

}

func TestIssue(t *testing.T) {
	var (
		kf  = func(token *jwt.Token) (interface{}, error) { return []byte("secret"), nil }
		e   = NewParser(kf, jwt.SigningMethodHS256, ClaimsFactory{Factory: MapClaimsFactory{}})(endpoint.Nop)
		key = kitjwt.JWTContextKey
		val = "eyJhbGciOiJIUzI1NiIsImtpZCI6ImtpZCIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiZ28ta2l0In0.14M2VmYyApdSlV_LZ88ajjwuaLeIFplB8JpyNy0A19E"
		ctx = context.WithValue(context.Background(), key, val)
	)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			e(ctx, struct{}{}) // fatal error: concurrent map read and map write
		}()
	}
	wg.Wait()
}
