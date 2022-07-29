package lb

import (
	"context"

	"golang.org/x/exp/maps"

	"github.com/go-leo/leo/common/sortx"
	"github.com/go-leo/leo/runner/net/http/client"
)

type PickFirst struct{}

func (p *PickFirst) Pick(_ context.Context, interfaces map[string]client.Interface) (client.Interface, error) {
	if len(interfaces) <= 0 {
		return nil, ErrNoInterface
	}
	keys := maps.Keys(interfaces)
	sortx.Asc(keys)
	return interfaces[keys[0]], nil
}
