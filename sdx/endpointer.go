package sdx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/log"
	"net/url"
)

func NewEndpointer(
	ctx context.Context,
	target string,
	color string,
	builder Builder,
	factory sd.Factory,
	logger log.Logger,
	options ...sd.EndpointerOption,
) (sd.Endpointer, error) {
	instanceUrl, err := url.Parse(target)
	if err == nil && builder.Scheme() == instanceUrl.Scheme {
		return newEndpointer(ctx, instanceUrl, color, builder, factory, logger, options...)
	}
	canonicalTarget := builder.Scheme() + ":///" + target
	instanceUrl, err = url.Parse(canonicalTarget)
	if err != nil {
		return nil, fmt.Errorf("sdx: failed to parse canonical target instance: %q", canonicalTarget)
	}
	return newEndpointer(ctx, instanceUrl, color, builder, factory, logger, options...)
}

func newEndpointer(ctx context.Context, instanceUrl *url.URL, color string, builder Builder, factory sd.Factory, logger log.Logger, options ...sd.EndpointerOption) (sd.Endpointer, error) {
	instancer, err := builder.BuildInstancer(ctx, instanceUrl, color, logger)
	if err != nil {
		return nil, fmt.Errorf("sdx: failed to new instancer, target url: %q, color: %q", instanceUrl.String(), color)
	}
	return sd.NewEndpointer(instancer, factory, logger, options...), nil
}
