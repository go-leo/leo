package text_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/config/resource/text"
	"context"
	"github.com/stretchr/testify/assert"
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
	txt := text.NewText(content)
	resource := text.NewResource(txt, "content", "yaml")

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
		txt.Set(content + "\n" + content)
	}()

	e := <-eventC
	src, err := e.Get()
	assert.NoError(t, err)
	assert.Equal(t, content+"\n"+content, string(src.Value()))
}
