package leo

import (
	"context"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	_ = app.Run(context.Background())
}
