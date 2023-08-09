package nacosv2_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/config/resource/nacosv2"
	"context"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

var content = `controller:
    port: 8080
provider:
    port: 9090
actuator:
    port: 16060`

func TestResource(t *testing.T) {
	factory := func() (config_client.IConfigClient, error) {
		sc := []constant.ServerConfig{
			*constant.NewServerConfig(os.Getenv("NACOS_ADDR"), 8848, constant.WithContextPath("/nacos")),
		}

		cc := *constant.NewClientConfig(
			constant.WithNamespaceId(""),
			constant.WithTimeoutMs(5000),
			constant.WithNotLoadCacheAtStart(true),
			constant.WithLogDir("/tmp/nacos/log"),
			constant.WithCacheDir("/tmp/nacos/cache"),
			constant.WithLogLevel("debug"),
		)
		return clients.NewConfigClient(
			vo.NacosClientParam{
				ClientConfig:  &cc,
				ServerConfigs: sc,
			},
		)
	}
	client, err := factory()
	assert.NoError(t, err)

	publishConfig, err := client.PublishConfig(vo.ConfigParam{
		DataId:  "demo",
		Group:   "yaml",
		Content: content,
	})
	assert.NoError(t, err)
	assert.True(t, publishConfig)

	resource, err := nacosv2.NewResource("demo", "yaml", factory, nacosv2.Extension("yaml"))
	assert.NoError(t, err)

	res, err := resource.Load(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, "yaml", res.Extension())
	assert.Equal(t, content, string(res.Value()))

	watcher, err := resource.Watch(context.Background())
	assert.NoError(t, err)

	eventC := make(chan config.Event)
	watcher.Notify(eventC)

	go func() {
		time.Sleep(2 * time.Second)
		publishConfig, err := client.PublishConfig(vo.ConfigParam{
			DataId:  "demo",
			Group:   "yaml",
			Content: content + "\n" + content,
		})
		assert.NoError(t, err)
		assert.True(t, publishConfig)
	}()

	e := <-eventC
	src, err := e.Get()
	assert.NoError(t, err)
	assert.Equal(t, content+"\n"+content, string(src.Value()))
}
