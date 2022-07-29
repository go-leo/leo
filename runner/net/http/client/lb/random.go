package lb

import (
	"context"
	"sync/atomic"

	"golang.org/x/exp/maps"

	"github.com/go-leo/leo/common/sortx"
	"github.com/go-leo/leo/runner/net/http/client"
)

type Random struct {
	number uint64
}

func (r *Random) Interface(_ context.Context, interfaces map[string]client.Interface) (client.Interface, error) {
	if len(interfaces) <= 0 {
		return nil, ErrNoInterface
	}
	keys := maps.Keys(interfaces)
	sortx.Asc(keys)
	number := atomic.AddUint64(&r.number, 1)
	index := number % uint64(len(keys))
	return interfaces[keys[index]], nil
}
