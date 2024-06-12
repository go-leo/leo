package consulx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/log"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-playground/form/v4"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// schemeName for the urls
// All target URLs like 'consul://.../...' will be resolved by this builder
const schemeName = "consul"

type InstancerBuilder struct {
	log log.Logger
}

func (b *InstancerBuilder) Build(ctx context.Context, target *url.URL, colors []string) (sd.Instancer, error) {
	dsn := strings.Join([]string{schemeName + ":/", target.Host, target.Path + "?" + target.RawQuery}, "/")
	service, config, err := parseURL(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "Wrong consul URL")
	}
	cli, err := api.NewClient(config)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't connect to the Consul API")
	}
	return consul.NewInstancer(consul.NewClient(cli), b.log, service, colors, true), nil
}

func (b *InstancerBuilder) Scheme() string {
	return schemeName
}

func NewInstancerBuilder(log log.Logger) sdx.InstancerBuilder {
	return &InstancerBuilder{log: log}
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

	username := rawURL.User.Username()
	password, _ := rawURL.User.Password()
	addr := rawURL.Host
	service := strings.TrimLeft(rawURL.Path, "/")

	q := queries{}
	decoder := form.NewDecoder()
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) { return time.ParseDuration(vals[0]) }, time.Duration(0))
	err = decoder.Decode(&q, rawURL.Query())
	if err != nil {
		return "", nil, fmt.Errorf("malformed url parameters, %w", err)
	}
	if len(q.Near) == 0 {
		q.Near = "_agent"
	}
	if q.MaxBackoff == 0 {
		q.MaxBackoff = time.Second
	}
	var auth *api.HttpBasicAuth
	if len(username) > 0 && len(password) > 0 {
		auth = &api.HttpBasicAuth{
			Username: username,
			Password: password,
		}
	}
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

type queries struct {
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
