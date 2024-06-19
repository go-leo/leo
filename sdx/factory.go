package sdx

import (
	"context"
	"github.com/go-kit/kit/sd"
)

type Factory func(ctx context.Context, args any) sd.Factory
