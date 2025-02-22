package sdx

import (
	"google.golang.org/grpc/resolver"
	"net/url"
	"testing"
)

func TestParseTarget(t *testing.T) {
	target, err := parseTarget("localhost:50051")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(target)
}

func parseTarget(target string) (resolver.Target, error) {
	u, err := url.Parse(target)
	if err != nil {
		return resolver.Target{}, err
	}

	return resolver.Target{URL: *u}, nil
}
