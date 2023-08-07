package requestid

import (
	"context"
	"testing"
)

func TestFromContext(t *testing.T) {
	testValue := "test-request-id"
	ctx := NewContext(context.Background(), testValue)

	val, ok := FromContext(ctx)

	if !ok {
		t.Errorf("Expected to find a value in the context, but found none")
	}

	if val != testValue {
		t.Errorf("Expected %s, but got %s", testValue, val)
	}
}

func TestNewContext(t *testing.T) {
	testValue := "test-request-id"
	ctx := NewContext(context.Background(), testValue)

	val, ok := ctx.Value(key).(string)

	if !ok {
		t.Errorf("Expected to find a value in the context, but found none")
	}

	if val != testValue {
		t.Errorf("Expected %s, but got %s", testValue, val)
	}
}
