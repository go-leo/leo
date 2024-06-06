package statusx

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestIs(t *testing.T) {
	tokenErr := grpcstatus.New(codes.Unauthenticated, "token is valid").Err()
	signErr := grpcstatus.New(codes.Unauthenticated, "signature is valid").Err()
	is := errors.Is(tokenErr, signErr)
	t.Log(is)

	fmt.Println(time.Now().Unix())
}
