package consulx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/hashicorp/consul/api"
	"net/url"
	"strings"
)

// schemeName for the urls
// All target URLs like 'consul://.../...' will be resolved by this builder
const schemeName = "consul"

type InstancerBuilder struct {
	ConfigParser func(rawURL *url.URL) (*api.Config, error)
}

func (b *InstancerBuilder) Build(ctx context.Context, target *sdx.Target) (sd.Instancer, error) {
	dsn := strings.Join([]string{schemeName + ":/", target.URL.Host, target.URL.Path + "?" + target.URL.RawQuery}, "/")
	rawURL, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("malformed url, %v", err)
	}
	if rawURL.Scheme != schemeName || len(rawURL.Host) == 0 || len(strings.TrimLeft(rawURL.Path, "/")) == 0 {
		return nil, fmt.Errorf("malformed url('%s'). must be in the next format: 'consul://[username:password]@host/service?param=value'", dsn)
	}
	service := strings.TrimLeft(rawURL.Path, "/")
	if b.ConfigParser == nil {
		b.ConfigParser = DefaultConfigParser
	}
	config, err := b.ConfigParser(rawURL)
	if err != nil {
		return nil, err
	}
	cli, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	color, ok := sdx.ExtractColor(ctx)
	if !ok {
		return consul.NewInstancer(consul.NewClient(cli), logx.FromContext(ctx), service, nil, true), nil
	}
	return consul.NewInstancer(consul.NewClient(cli), logx.FromContext(ctx), service, []string{color}, true), nil
}

func (b *InstancerBuilder) Scheme() string {
	return schemeName
}
