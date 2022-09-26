package nacos_test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/stretchr/testify/assert"

	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/config/medium/nacos"
)

var namespaceID = "public" // nolint
var group = "test"
var dataID = "demo"
var confType = "yaml"
var confContent = `bool: true
int: 10
int_32: -200
int_64: -3000
u_int: 133
u_int_32: 413
u_int_64: 564
float_64: 1.3
time: 2021-07-07T17:16:12.361234+08:00
duration: 1s`

var client config_client.IConfigClient

func TestMain(m *testing.M) {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}
	cc := constant.NewClientConfig()
	var err error
	client, err = clients.NewConfigClient(vo.NacosClientParam{ClientConfig: cc, ServerConfigs: sc})
	if err != nil {
		log.Fatalln(err)
	}

	ok, err := client.PublishConfig(vo.ConfigParam{Group: group, DataId: dataID, Content: confContent})
	if err != nil {
		log.Fatalln(err)
	}

	if !ok {
		log.Fatalln(ok)
	}

	m.Run()

	ok, err = client.DeleteConfig(vo.ConfigParam{Group: group, DataId: dataID})
	if err != nil {
		log.Fatalln(err)
	}

	if !ok {
		log.Fatalln(ok)
	}

	os.Exit(0)
}

func TestLoader(t *testing.T) {
	loader := nacos.NewLoader("localhost", "8848", "1233", group, dataID, confType)
	contentType := loader.ContentType()
	assert.Equal(t, confType, contentType, "content type not match")

	err := loader.Load()
	assert.Nil(t, err, "failed load file")

	data := loader.RawData()
	assert.Equal(t, confContent, string(data), "config content not equal")
}

func TestWatcher(t *testing.T) {
	watcher := nacos.NewWatcher("localhost", "8848", "1233", group, dataID)
	c, err := watcher.Start(context.Background())
	assert.Nil(t, err, "failed watch file")

	go func() {
		time.Sleep(3 * time.Second)
		client, err := clients.NewConfigClient(vo.NacosClientParam{ClientConfig: constant.NewClientConfig(), ServerConfigs: []constant.ServerConfig{*constant.NewServerConfig("127.0.0.1", 8848)}})
		if err != nil {
			log.Fatalln(err)
		}
		content := confContent + "\nfloat_32: 0.3"
		ok, err := client.PublishConfig(vo.ConfigParam{Group: group, DataId: dataID, Content: content})
		assert.Nil(t, err, "failed open file")
		assert.True(t, ok, "publish config not ok")

		time.Sleep(3 * time.Second)
		_ = watcher.Stop(context.Background())
	}()
	events := c
	var e *config.Event
	for event := range events {
		e = event
	}
	assert.Equal(t, confContent+"\nfloat_32: 0.3", string(e.Data()))
}
