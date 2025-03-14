package logx

import (
	"context"
	"testing"
)

func TestWithContext(t *testing.T) {
	ctx := KeyValsExtractorInjector(context.Background(), "trace_id", "123456")
	ctx = KeyValsExtractorInjector(ctx, "span_id", "987654")
	Debugln(ctx, "arg1", "arg2")
}
