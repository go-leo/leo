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

type InstancerBuilder struct{}

func (b *InstancerBuilder) Build(ctx context.Context, target *sdx.Target, color *sdx.Color) (sd.Instancer, error) {
	dsn := strings.Join([]string{schemeName + ":/", target.URL.Host, target.URL.Path + "?" + target.URL.RawQuery}, "/")
	service, config, err := parseURL(dsn)
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

func NewInstancerBuilder() sdx.InstancerBuilder {
	return &InstancerBuilder{}
}

//	parseURL with parameters
//
// see README.md for the actual format
// URL schema will stay stable in the future for backward compatibility
func parseURL(u string) (string, *api.Config, error) {
	rawURL, err := url.Parse(u)
	if err != nil {
		return "", nil, fmt.Errorf("malformed url, %w", err)
	}

	if rawURL.Scheme != schemeName ||
		len(rawURL.Host) == 0 || len(strings.TrimLeft(rawURL.Path, "/")) == 0 {
		return "", nil, fmt.Errorf("malformed url('%s'). must be in the next format: 'consul://[username:password]@host/service?param=value'", u)
	}

	service := strings.TrimLeft(rawURL.Path, "/")

	q := args{}
	decoder := form.NewDecoder()
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) { return time.ParseDuration(vals[0]) }, time.Duration(0))
	if err := decoder.Decode(&q, rawURL.Query()); err != nil {
		return "", nil, fmt.Errorf("malformed url parameters, %w", err)
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
	return service, config, nil
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
