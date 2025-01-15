package consulx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/stain"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	stdconsul "github.com/hashicorp/consul/api"
	"net"
	"net/url"
	"strings"
)

var _ sdx.Builder = (*Builder)(nil)

// schemeName for the urls
// All target URLs like 'consul://.../...' will be resolved by this builder
const schemeName = "consul"

type Builder struct {
	ClientCreator func(rawURL *url.URL, color string) (*api.Client, error)
}

func (Builder) Scheme() string {
	return schemeName
}

func (b Builder) BuildInstancer(ctx context.Context, instance *url.URL, color string) (sd.Instancer, error) {
	dsn := strings.Join([]string{schemeName + ":/", instance.Host, instance.Path + "?" + instance.RawQuery}, "/")
	rawURL, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("malformed url, %v", err)
	}
	if rawURL.Scheme != schemeName || len(rawURL.Host) == 0 || len(strings.TrimLeft(rawURL.Path, "/")) == 0 {
		return nil, fmt.Errorf("malformed url('%s'). must be in the next format: 'consul://[username:password]@host/service?param=value'", dsn)
	}
	if b.ClientCreator == nil {
		b.ClientCreator = DefaultClientCreator
	}
	cli, err := b.ClientCreator(rawURL, color)
	if err != nil {
		return nil, err
	}

	service := strings.TrimLeft(rawURL.Path, "/")
	color, ok := stain.ExtractColor(ctx)
	if !ok {
		return consul.NewInstancer(consul.NewClient(cli), logx.FromContext(ctx), service, nil, true), nil
	}
	return consul.NewInstancer(consul.NewClient(cli), logx.FromContext(ctx), service, []string{color}, true), nil
}

func (b Builder) BuildRegistrar(ctx context.Context, instance *url.URL, ip net.IP, port int, color string) (sd.Registrar, error) {
	dsn := strings.Join([]string{schemeName + ":/", instance.Host, instance.Path + "?" + instance.RawQuery}, "/")
	rawURL, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("malformed url, %v", err)
	}
	if rawURL.Scheme != schemeName || len(rawURL.Host) == 0 || len(strings.TrimLeft(rawURL.Path, "/")) == 0 {
		return nil, fmt.Errorf("malformed url('%s'). must be in the next format: 'consul://[username:password]@host/service?param=value'", dsn)
	}
	if b.ClientCreator == nil {
		b.ClientCreator = DefaultClientCreator
	}
	cli, err := b.ClientCreator(rawURL, color)
	if err != nil {
		return nil, err
	}
	service := strings.TrimLeft(rawURL.Path, "/")
	return consul.NewRegistrar(consul.NewClient(cli), &stdconsul.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    service,
		Tags:    []string{color},
		Port:    port,
		Address: ip.String(),
	}, logx.L()), nil
}
