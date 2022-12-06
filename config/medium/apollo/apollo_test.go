package apollo_test

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-leo/leo/v2/config/medium/apollo"
)

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

func TestLoader(t *testing.T) {
	loader := apollo.NewLoader(
		"39.107.67.209",
		"8080",
		"test.apollo.config",
		"TEST",
		"mgr.yaml",
		apollo.Secret("8b8de8f58bf64668bbcb665581761b05"),
	)
	contentType := loader.ContentType()
	assert.Equal(t, confType, contentType, "content type not match")

	err := loader.Load()
	assert.Nil(t, err, "failed load config")

	data := loader.RawData()
	assert.Equal(t, confContent, string(data), "config content not equal")
}

func TestExt(t *testing.T) {
	ext := filepath.Ext("mgr.yaml")
	assert.Equal(t, ".yaml", ext)

	ext = filepath.Ext("application")
	assert.Equal(t, "", ext)
}

func TestWatcher(t *testing.T) {
	watcher := apollo.NewWatcher(
		"39.107.67.209",
		"8080",
		"test.apollo.config",
		"TEST",
		"mgr.yaml",
		apollo.WithSecret("8b8de8f58bf64668bbcb665581761b05"),
	)
	c, err := watcher.Start(context.Background())
	assert.Nil(t, err, "failed watch file")
	go func() {
		time.Sleep(time.Minute)
		t.Log(watcher.Stop(context.Background()))
		t.Log(watcher.Stop(context.Background()))
		t.Log(watcher.Stop(context.Background()))
		t.Log(watcher.Stop(context.Background()))
	}()
	for event := range c {
		data := event.Data()
		assert.Equal(t, confContent+"\nfloat_32: 0.3", string(data))
	}

}
