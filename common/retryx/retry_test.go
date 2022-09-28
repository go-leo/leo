package retryx_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-leo/backoffx"

	"github.com/go-leo/leo/common/retryx"
)

func TestCall(t *testing.T) {
	maxAttempts := 3
	ctx := context.Background()
	method := func(attemptTime int) error {
		fmt.Println(attemptTime)
		if attemptTime < maxAttempts {
			return errors.New("mock error")
		}
		return nil
	}
	backoffFunc := backoffx.Constant(time.Second)
	err := retryx.Call(ctx, uint(maxAttempts), backoffFunc, method)
	assert.Nil(t, err)
}
