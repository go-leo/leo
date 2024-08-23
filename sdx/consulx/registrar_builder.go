package consulx

import (
	"github.com/hashicorp/consul/api"
	"net/url"
)

type RegistrarBuilder struct {
	ConfigParser func(rawURL *url.URL) (*api.Config, error)
}
