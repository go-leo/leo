package consulx

import (
	"context"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/hashicorp/consul/api"
	"net/http"
	"net/url"
	"time"
)

func DefaultClientCreator(ctx context.Context, rawURL *url.URL, color string) (*api.Client, error) {
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
	return api.NewClient(config)
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
