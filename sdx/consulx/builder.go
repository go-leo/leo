package consulx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-playground/form/v4"
	"github.com/hashicorp/consul/api"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// schemeName for the urls
// All target URLs like 'consul://.../...' will be resolved by this builder
const schemeName = "consul"

type InstancerBuilder struct {
	ConfigParser func(rawURL *url.URL) (*api.Config, error)
}

func (b *InstancerBuilder) Build(ctx context.Context, target *sdx.Target, color *sdx.Color) (sd.Instancer, error) {
	dsn := strings.Join([]string{schemeName + ":/", target.URL.Host, target.URL.Path + "?" + target.URL.RawQuery}, "/")
	rawURL, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("malformed url, %v", err)
	}
	if rawURL.Scheme != schemeName || len(rawURL.Host) == 0 || len(strings.TrimLeft(rawURL.Path, "/")) == 0 {
		return nil, fmt.Errorf("malformed url('%s'). must be in the next format: 'consul://[username:password]@host/service?param=value'", dsn)
	}
	service := strings.TrimLeft(rawURL.Path, "/")
	config, err := DefaultConfigParser(rawURL)
	if err != nil {
		return nil, err
	}
	cli, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	if color == nil {
		return consul.NewInstancer(consul.NewClient(cli), logx.FromContext(ctx), service, nil, true), nil
	}
	return consul.NewInstancer(consul.NewClient(cli), logx.FromContext(ctx), service, color.Color(), true), nil
}

func (b *InstancerBuilder) Scheme() string {
	return schemeName
}

func DefaultConfigParser(rawURL *url.URL) (*api.Config, error) {
	q := args{}
	decoder := form.NewDecoder()
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) { return time.ParseDuration(vals[0]) }, time.Duration(0))
	if err := decoder.Decode(&q, rawURL.Query()); err != nil {
		return nil, fmt.Errorf("malformed url parameters, %v", err)
	}
	if len(q.Near) == 0 {
		q.Near = "_agent"
	}
	if q.MaxBackoff == 0 {
		q.MaxBackoff = time.Second
	}

	var auth *api.HttpBasicAuth
	username := rawURL.User.Username()
	password, _ := rawURL.User.Password()
	if len(username) > 0 && len(password) > 0 {
		auth = &api.HttpBasicAuth{
			Username: username,
			Password: password,
		}
	}

	addr := rawURL.Host

	config := &api.Config{
		Address:    addr,
		HttpAuth:   auth,
		WaitTime:   q.Wait,
		HttpClient: &http.Client{Timeout: q.Timeout},
		TLSConfig: api.TLSConfig{
			InsecureSkipVerify: q.TLSInsecure,
		},
		Token: q.Token,
	}
	return config, nil
}

type args struct {
	Wait              time.Duration `form:"wait"`
	Timeout           time.Duration `form:"timeout"`
	MaxBackoff        time.Duration `form:"max-backoff"`
	Tag               string        `form:"tag"`
	Near              string        `form:"near"`
	Limit             int           `form:"limit"`
	Healthy           bool          `form:"healthy"`
	TLSInsecure       bool          `form:"insecure"`
	Token             string        `form:"token"`
	Dc                string        `form:"dc"`
	AllowStale        bool          `form:"allow-stale"`
	RequireConsistent bool          `form:"require-consistent"`
}
