package nacosx

import (
	"context"
	"github.com/go-leo/leo/v3/configx"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"path/filepath"
	"strings"
)

var _ configx.Resource = (*Resource)(nil)

type Resource struct {
	Formatter configx.Formatter
	Client    config_client.IConfigClient
	Group     string
	DataId    string
}

func (r *Resource) Format() string {
	if r.Formatter == nil {
		return strings.TrimPrefix(filepath.Ext(r.DataId), ".")
	}
	return r.Formatter.Format()
}

func (r *Resource) Load(ctx context.Context) ([]byte, error) {
	content, err := r.Client.GetConfig(vo.ConfigParam{
		Group:  r.Group,
		DataId: r.DataId,
	})
	if err != nil {
		return nil, err
	}
	return []byte(content), nil
}

func (r *Resource) Watch(ctx context.Context, notifyC chan<- *configx.Event) error {
	err := r.Client.ListenConfig(vo.ConfigParam{
		Group:  r.Group,
		DataId: r.DataId,
		OnChange: func(_, _, _, data string) {
			notifyC <- configx.NewDataEvent([]byte(data))
		},
	})
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		if err := r.Client.CancelListenConfig(vo.ConfigParam{
			Group:  r.Group,
			DataId: r.DataId,
		}); err != nil {
			notifyC <- configx.NewErrorEvent(err)
		}
		notifyC <- configx.NewErrorEvent(configx.ErrStopWatch)
	}()
	return nil
}
