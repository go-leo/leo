package contextx_test

import (
	"syscall"
	"testing"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/contextx"
)

func TestSignal(t *testing.T) {
	ctx, cancel := contextx.Signal(syscall.SIGHUP)
	t.Log(ctx)
	t.Log(cancel)
}
