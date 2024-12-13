package consulx

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/leo/v3/configx"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/hashicorp/go-hclog"
)

var _ configx.Resource = (*Consul)(nil)

type consulParamKey struct{}

type consulParam struct {
	Key          string
	QueryOptions *api.QueryOptions
}

func WithConsulParam(ctx context.Context, key string, options *api.QueryOptions) context.Context {
	return context.WithValue(ctx, consulParamKey{}, consulParam{Key: key, QueryOptions: options})
}

type Consul struct {
	configx.Formatter
	client *api.Client
}

func (r *Consul) Load(ctx context.Context) ([]byte, error) {
	param, ok := ctx.Value(consulParamKey{}).(consulParam)
	if !ok {
		return nil, errors.New("resourcex: Consul param is nil")
	}

	pair, _, err := r.client.KV().Get(param.Key, param.QueryOptions)
	if err != nil {
		return nil, err
	}
	return pair.Value, nil
}

func (r *Consul) Watch(ctx context.Context, notifyC chan<- *configx.Event) (func(), error) {
	param, ok := ctx.Value(consulParamKey{}).(consulParam)
	if !ok {
		return nil, errors.New("resourcex: Consul param is nil")
	}
	params := map[string]any{
		"type": "key",
		"key":  param.Key,
	}
	if options := param.QueryOptions; options != nil {
		if options.Datacenter != "" {
			params["datacenter"] = options.Datacenter
		}
		if options.Token != "" {
			params["token"] = options.Token
		}
	}
	plan, err := watch.Parse(params)
	if err != nil {
		return nil, err
	}
	err = plan.RunWithClientAndHclog(r.client, &consuleLogger{Logger: hclog.NewNullLogger(), notifyC: notifyC})
	if err != nil {
		return nil, err
	}
	return func() {
		plan.Stop()
		notifyC <- configx.NewErrorEvent(configx.ErrStopWatch)
	}, nil
}

func NewConsul(client *api.Client, formatter configx.Formatter) *Consul {
	return &Consul{Formatter: formatter, client: client}
}

type consuleLogger struct {
	hclog.Logger
	notifyC chan<- *configx.Event
}

func (l *consuleLogger) Error(msg string, args ...interface{}) {
	l.notifyC <- configx.NewErrorEvent(fmt.Errorf(msg, args...))
}
