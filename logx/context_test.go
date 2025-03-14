package logx

import (
	"context"
	"testing"
)

func TestWithContext(t *testing.T) {
	ctx := InjectKeyVals(context.Background(), "trace_id", "123456")
	ctx = InjectKeyVals(ctx, "span_id", "987654")
	Debugln(ctx, "arg1", "arg2")
}
