package nacosx

import (
	"context"
	"errors"
	"github.com/go-leo/leo/v3/configx"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"os"
	"strings"
	"testing"
	"time"
)

func nacosFactory() (config_client.IConfigClient, error) {
	sc := []constant.ServerConfig{*constant.NewServerConfig(os.Getenv("Addr"), 8848)}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId("dev"),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("debug"),
		constant.WithLogDir("/tmp/nacos.log"),
		constant.WithAccessKey(os.Getenv("AccessKey")),
		constant.WithSecretKey(os.Getenv("SecretKey")),
	)
	return clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
}

func TestResource_Load_Nacos(t *testing.T) {
	configClient, err := nacosFactory()
	if err != nil {
		t.Errorf("factory() error = %v", err)
		return
	}

	_, err = configClient.PublishConfig(vo.ConfigParam{
		DataId:  "nacos",
		Group:   "test",
		Content: "TEST_KEY=test_value",
	})
	if err != nil {
		t.Errorf("PublishConfig() error = %v", err)
		return
	}

	defer func() {
		_, err = configClient.DeleteConfig(vo.ConfigParam{
			DataId: "nacos",
			Group:  "test",
		})
		if err != nil {
			t.Errorf("DeleteConfig() error = %v", err)
			return
		}
	}()

	time.Sleep(time.Second)

	r := Resource{
		Formatter: configx.Env{},
		Client:    configClient,
		Group:     "test",
		DataId:    "nacos",
	}
	ctx := context.Background()
	content, err := r.Load(ctx)
	if err != nil {
		t.Errorf("Load() error = %v", err)
		return
	}

	if !strings.Contains(string(content), "TEST_KEY=test_value") {
		t.Errorf("Load() data = %v, want data to contain 'TEST_KEY=test_value'", string(content))
	}

	time.Sleep(time.Second)

}

func TestResource_Watch_Nacos(t *testing.T) {
	configClient, err := nacosFactory()
	if err != nil {
		t.Errorf("factory() error = %v", err)
		return
	}

	_, err = configClient.PublishConfig(vo.ConfigParam{
		DataId:  "nacos",
		Group:   "test",
		Content: "TEST_KEY=test_value",
	})
	if err != nil {
		t.Errorf("PublishConfig() error = %v", err)
		return
	}

	defer func() {
		_, err = configClient.DeleteConfig(vo.ConfigParam{
			DataId: "nacos",
			Group:  "test",
		})
		if err != nil {
			t.Errorf("DeleteConfig() error = %v", err)
			return
		}
	}()

	time.Sleep(time.Second)

	r := Resource{
		Formatter: configx.Env{},
		Client:    configClient,
		Group:     "test",
		DataId:    "nacos",
	}
	notifyC := make(chan *configx.Event, 1)
	// Start watching
	ctx := context.Background()
	stopFunc, err := r.Watch(ctx, notifyC)
	if err != nil {
		t.Errorf("Watch() error = %v", err)
		return
	}

	// Give some time for the watcher to detect the change
	go func() {
		for {
			time.Sleep(time.Second)
			ok, err := configClient.PublishConfig(vo.ConfigParam{
				DataId:  "nacos",
				Group:   "test",
				Content: "TEST_KEY_NEW=test_value_new" + time.Now().Format(time.RFC3339),
			})
			if err != nil {
				t.Errorf("PublishConfig() error = %v", err)
				return
			}
			t.Log(ok)
		}
	}()

	// Wait for the event
	select {
	case event := <-notifyC:
		if data, ok := event.AsDataEvent(); !ok || data.Data == nil {
			t.Error("Expected DataEvent with non-nil data")
		}
	case <-time.After(100 * time.Second):
		t.Error("No event received within the timeout")
	}

	stopFunc()

	_, err = configClient.PublishConfig(vo.ConfigParam{
		DataId:  "nacos",
		Group:   "test",
		Content: "TEST_KEY=another_test_value",
	})
	if err != nil {
		t.Errorf("PublishConfig() error = %v", err)
		return
	}

	select {
	case event := <-notifyC:
		err, ok := event.AsErrorEvent()
		if !ok || err.Err == nil || !errors.Is(err.Err, configx.ErrStopWatch) {
			t.Error("Did not expect to receive an event after stopping the watcher")
		}
	case <-time.After(100 * time.Millisecond):
		// Expected behavior
	}
}
