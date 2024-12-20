package sdx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/log"
	"net/url"
)

func NewEndpointer(ctx context.Context, target string, color string, instancerFactory InstancerFactory, factory sd.Factory, logger log.Logger, options ...sd.EndpointerOption) (sd.Endpointer, error) {
	targetUrl, err := url.Parse(target)
	if err != nil {
		canonicalTarget := instancerFactory.Scheme() + ":///" + target
		targetUrl, err = url.Parse(canonicalTarget)
		if err != nil {
			return nil, fmt.Errorf("sdx: failed to parse canonical target: %q", canonicalTarget)
		}
	}
	instancer, err := instancerFactory.New(ctx, targetUrl, color)
	if err != nil {
		return nil, fmt.Errorf("sdx: failed to new instancer, target url: %q, color: %q", targetUrl.String(), color)
	}
	return sd.NewEndpointer(instancer, factory, logger, options...), nil
}
