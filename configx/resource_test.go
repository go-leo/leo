package configx

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestResource_Env_Load(t *testing.T) {
	// Mock environment variables
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	r := &Environ{}
	data, err := r.Load(context.Background())
	if err != nil {
		t.Errorf("Load() error = %v", err)
		return
	}
	if !strings.Contains(string(data), "TEST_KEY=test_value") {
		t.Errorf("Load() data = %v, want data to contain 'TEST_KEY=test_value'", string(data))
	}
}

func TestResource_Env_Watch(t *testing.T) {
	// Mock environment variables
	os.Setenv("TEST_KEY", "test_value")
	defer os.Unsetenv("TEST_KEY")

	notifyC := make(chan *Event, 1)
	r := &Environ{}

	// Start watching
	stopFunc, err := r.Watch(context.Background(), notifyC)
	if err != nil {
		t.Errorf("Watch() error = %v", err)
		return
	}

	// Give some time for the watcher to detect the change
	time.Sleep(1 * time.Second)

	// Modify the environment variable to trigger a change
	os.Setenv("TEST_KEY", "new_test_value")

	// Wait for the event
	select {
	case event := <-notifyC:
		if data, ok := event.AsDataEvent(); !ok || data.Data == nil {
			t.Error("Expected DataEvent with non-nil data")
		}
	case <-time.After(1 * time.Second):
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
		if !ok || err.Err == nil || !errors.Is(err.Err, ErrStopWatch) {
			t.Error("Did not expect to receive an event after stopping the watcher")
		}
	case <-time.After(100 * time.Millisecond):
		// Expected behavior
	}
}

func TestResource_File_Load(t *testing.T) {
	// Mock environment variables
	file, _ := os.CreateTemp("", ".env")
	_, _ = fmt.Fprintln(file, "TEST_KEY=test_value")

	defer os.Remove(file.Name())

	r := File{Formatter: Env{}, Filename: file.Name()}
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

	r := File{Formatter: Env{}, Filename: file.Name()}

	notifyC := make(chan *Event, 1)

	// Start watching
	stopFunc, err := r.Watch(context.Background(), notifyC)
	if err != nil {
		t.Errorf("Watch() error = %v", err)
		return
	}

	// Give some time for the watcher to detect the change
	time.Sleep(1 * time.Second)

	_, _ = fmt.Fprintln(file, "TEST_KEY_NEW=test_value_new")
	_ = file.Sync()

	// Wait for the event
	select {
	case event := <-notifyC:
		if data, ok := event.AsDataEvent(); !ok || data.Data == nil {
			t.Error("Expected DataEvent with non-nil data")
		}
	case <-time.After(1 * time.Second):
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
		if !ok || err.Err == nil || !errors.Is(err.Err, ErrStopWatch) {
			t.Error("Did not expect to receive an event after stopping the watcher")
		}
	case <-time.After(100 * time.Millisecond):
		// Expected behavior
	}
}
