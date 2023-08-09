package env_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config"
	"codeup.aliyun.com/qimao/leo/leo/config/resource/env"
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestResource(t *testing.T) {
	os.Setenv("TEST_PATH", "test")

	resource := env.NewResource(
		env.Prefix("TEST"),
	)
	res, err := resource.Load(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, "env", res.Extension())
	assert.Equal(t, "TEST_PATH=test", string(res.Value()))

	watcher, err := resource.Watch(context.Background())
	assert.NoError(t, err)

	eventC := make(chan config.Event)
	watcher.Notify(eventC)

	go func() {
		time.Sleep(2 * time.Second)
		os.Setenv("TEST_REPO", "leo")
	}()

	e := <-eventC
	src, err := e.Get()
	assert.NoError(t, err)
	assert.Equal(t, "TEST_PATH=test\nTEST_REPO=leo", string(src.Value()))
}
