package consulx

import (
	"fmt"
	"net/url"
	"time"

	"github.com/hashicorp/consul/api"

	"github.com/go-leo/leo/common/stringx"
)

// consul://username:password@ip:port?scheme=http&datacenter=dev&token=12345&wait_time=1s&tls=true

func NewClient(uri *url.URL) (*api.Client, error) {
	query := uri.Query()
	config := &api.Config{
		Address:    uri.Host,
		Scheme:     query.Get("scheme"),
		Datacenter: query.Get("datacenter"),
		Token:      query.Get("token"),
		TokenFile:  query.Get("token_file"),
		TLSConfig: api.TLSConfig{
			Address:            "",
			CAFile:             "",
			CAPath:             "",
			CertFile:           "",
			KeyFile:            "",
			InsecureSkipVerify: false,
		},
	}

	if uri.User != nil && stringx.IsNotBlank(uri.User.Username()) {
		password, _ := uri.User.Password()
		config.HttpAuth = &api.HttpBasicAuth{Username: uri.User.Username(), Password: password}
	}

	waitTimeStr := query.Get("wait_time")
	if stringx.IsNotBlank(waitTimeStr) {
		duration, err := time.ParseDuration(waitTimeStr)
		if err != nil {
			return nil, fmt.Errorf("failed parse wait time %s, %w", waitTimeStr, err)
		}
		config.WaitTime = duration
	}

	if query.Get("tls") == "true" {
		config.TLSConfig = api.TLSConfig{
			Address:  query.Get("address"),
			CAFile:   query.Get("ca_file"),
			CAPath:   query.Get("ca_path"),
			CertFile: query.Get("cert_file"),
			KeyFile:  query.Get("key_file"),
		}
		if query.Get("insecure") == "true" {
			config.TLSConfig.InsecureSkipVerify = true
		}
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed new consul client, %w", err)
	}
	return client, nil
}
