package filex

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-leo/leo/v3/configx"
	"os"
	"strings"
	"testing"
	"time"
)

func TestResource_File_Load(t *testing.T) {
	// Mock environment variables
	file, _ := os.CreateTemp("", ".env")
	_, _ = fmt.Fprintln(file, "TEST_KEY=test_value")

	defer os.Remove(file.Name())

	r := Resource{Formatter: configx.Env{}, Filename: file.Name()}
	data, err := r.Load(context.Background())
	if err != nil {
		t.Errorf("Load() error = %v", err)
		return
	}
	if !strings.Contains(string(data), "TEST_KEY=test_value") {
		t.Errorf("Load() data = %v, want data to contain 'TEST_KEY=test_value'", string(data))
	}
}

func TestResource_File_Watch(t *testing.T) {
	// Mock environment variables
	file, _ := os.CreateTemp("", ".env")
	_, _ = fmt.Fprintln(file, "TEST_KEY=test_value")

	defer os.Remove(file.Name())

	r := Resource{Formatter: configx.Env{}, Filename: file.Name()}

	notifyC := make(chan *configx.Event, 1)

	// Start watching
	stopFunc, err := r.Watch(context.Background(), notifyC)
	if err != nil {
		t.Errorf("Watch() error = %v", err)
		return
	}

	go func() {
		time.Sleep(5 * time.Second)
		_, _ = fmt.Fprintln(file, "TEST_KEY_NEW=test_value_new")
		_ = file.Sync()
	}()

	// Wait for the event
	select {
	case event := <-notifyC:
		if data, ok := event.AsDataEvent(); !ok || data.Data == nil {
			t.Error("Expected DataEvent with non-nil data")
		}
	case <-time.After(10 * time.Second):
		t.Error("No event received within the timeout")
	}

	// Stop the watcher
	stopFunc()

	// Ensure that the watcher has been stopped
	// by sending an environment variable change and
	// not receiving an event within a short period.
	os.Setenv("TEST_KEY", "another_test_value")
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
