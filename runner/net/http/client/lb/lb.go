package lb

import (
	"context"
	"errors"

	"github.com/go-leo/leo/runner/net/http/client"
)

// Balancer selects on Interface from multiple interfaces based on a certain policy
type Balancer interface {
	Interface(context.Context, map[string]client.Interface) (client.Interface, error)
}

var ErrNoInterface = errors.New("no interface available")
