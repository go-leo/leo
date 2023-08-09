package file_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/config/resource/file"
	"context"
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
	f, err := os.CreateTemp("", "config*.yaml")
	assert.NoError(t, err)
	_, err = f.WriteString(content)
	assert.NoError(t, err)
	err = f.Sync()
	assert.NoError(t, err)

	resource := file.NewResource(f.Name())
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
		_, err = f.WriteString("\n")
		assert.NoError(t, err)
		_, err = f.WriteString(content)
		assert.NoError(t, err)
		err = f.Sync()
		assert.NoError(t, err)
	}()

	e := <-eventC
	src, err := e.Get()
	assert.NoError(t, err)
	assert.Equal(t, content+"\n"+content, string(src.Value()))
}
