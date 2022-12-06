package file_test

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/go-leo/leo/v2/config/medium/file"
)

var confFilename = "test.yaml"
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
duration: 1s
`

func TestMain(t *testing.M) {
	fp := filepath.Join(os.TempDir(), confFilename)

	if err := os.WriteFile(fp, []byte(confContent), os.ModePerm); err != nil {
		log.Fatalln(err)
	}

	t.Run()

	if err := os.Remove(fp); err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}

func TestLoader(t *testing.T) {
	fp := filepath.Join(os.TempDir(), confFilename)
	loader := file.NewLoader(fp)
	contentType := loader.ContentType()
	assert.Equal(t, confType, contentType, "content type not match")

	err := loader.Load()
	assert.Nil(t, err, "failed load file")

	data := loader.RawData()
	assert.Equal(t, confContent, string(data), "config content not equal")
}

func TestWatcher(t *testing.T) {
	fp := filepath.Join(os.TempDir(), confFilename)
	watcher := file.NewWatcher(fp)
	events, err := watcher.Start(context.Background())
	assert.Nil(t, err, "failed watch file")
	go func() {
		file, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, os.ModePerm)
		defer func() {
			_ = file.Close()
		}()
		assert.Nil(t, err, "failed open file")
		n, err := file.WriteString("float_32: 0.3")
		assert.Nil(t, err, "failed write string")
		assert.Greater(t, n, 0, "failed write string")

		time.Sleep(5 * time.Second)
		_ = watcher.Stop(context.Background())
	}()
	event := <-events
	data := event.Data()
	assert.Equal(t, confContent+"float_32: 0.3", string(data))
	_ = watcher.Stop(context.Background())
}
