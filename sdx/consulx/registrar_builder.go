package consulx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/hashicorp/consul/api"
	"net/url"
)

type RegistrarBuilder struct {
	ConfigParser func(rawURL *url.URL) (*api.Config, error)
}

func (b *RegistrarBuilder) Build(ctx context.Context, service sdx.ServiceInstance) sd.Registrar {

}
