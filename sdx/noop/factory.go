package noop

import (
	"context"
	"github.com/go-kit/kit/sd"
	kitlog "github.com/go-kit/log"

	"github.com/go-leo/leo/v3/sdx"
	"net"
	"net/url"
)

var _ sdx.Builder = (*Builder)(nil)

const schemeName = "noop"

type Builder struct{}

func (Builder) Scheme() string {
	return schemeName
}

func (b Builder) BuildInstancer(ctx context.Context, instance *url.URL, color string, logger kitlog.Logger) (sd.Instancer, error) {
	//TODO implement me
	panic("implement me")
}

func (b Builder) BuildRegistrar(ctx context.Context, instance *url.URL, ip net.IP, port int, color string, logger kitlog.Logger) (sd.Registrar, error) {
	//TODO implement me
	panic("implement me")
}
