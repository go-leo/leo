package consulx

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/configx"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/hashicorp/go-hclog"
	"path/filepath"
	"strings"
)

var _ configx.Resource = (*Resource)(nil)

type Resource struct {
	Formatter configx.Formatter
	Client    *api.Client
	Key       string
}

func (r *Resource) Format() string {
	if r.Formatter == nil {
		return strings.TrimPrefix(filepath.Ext(r.Key), ".")
	}
	return r.Formatter.Format()
}

func (r *Resource) Load(ctx context.Context) ([]byte, error) {
	pair, _, err := r.Client.KV().Get(r.Key, nil)
	if err != nil {
		return nil, err
	}
	return pair.Value, nil
}

func (r *Resource) Watch(ctx context.Context, notifyC chan<- *configx.Event) error {
	params := map[string]any{
		"type": "key",
		"key":  r.Key,
	}
	plan, err := watch.Parse(params)
	if err != nil {
		return err
	}
	plan.Handler = func(idx uint64, raw interface{}) {
		if raw == nil {
			return
		}
		if pair, ok := raw.(*api.KVPair); ok {
			notifyC <- configx.NewDataEvent(pair.Value)
		}
	}
	go func() {
		err = plan.RunWithClientAndHclog(r.Client, &consuleLogger{Logger: hclog.NewNullLogger(), notifyC: notifyC})
		if err != nil {
			notifyC <- configx.NewErrorEvent(err)
		}
		notifyC <- configx.NewErrorEvent(configx.ErrStopWatch)
	}()
	go func() {
		<-ctx.Done()
		plan.Stop()
	}()
	return nil
}

type consuleLogger struct {
	hclog.Logger
	notifyC chan<- *configx.Event
}

func (l *consuleLogger) Error(msg string, args ...interface{}) {
	l.notifyC <- configx.NewErrorEvent(fmt.Errorf(msg, args...))
}
