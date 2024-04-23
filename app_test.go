package leo

import (
	"context"
	"os"
	"testing"

	"github.com/go-kit/log"
)

func TestNewApp(t *testing.T) {
	app := NewApp(Logger(log.NewJSONLogger(os.Stdout)))
	_ = app.Run(context.Background())
}
