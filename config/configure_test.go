package config_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config"
	"context"
	"encoding/json"
	"math/rand"

	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

var content = `{"family":{"last_name":"Hashimoto"},"location":{"city":"San Francisco"},"first_name":"Mitchell","name":"Mitchell","age":91,"emails":["one","two","three"],"extra":{"twitter":"mitchellh"}}`

var content2 = `{"family":{"last_name":"last_name"},"location":{"city":"city"},"first_name":"first_name","name":"name","age":54,"emails":["one","two","three"],"extra":{"qq":"42213445"}}`

type MockResource struct {
}

func (m *MockResource) Load(ctx context.Context) (*config.Source, error) {
	return config.NewSource("mock", []byte(content), "json"), nil
}

func (m *MockResource) Watch(ctx context.Context) (config.Watcher, error) {
	mw := &MockWatcher{}
	mw.init()
	return mw, nil
}

type MockWatcher struct {
	eventC chan<- config.Event
}

func (m *MockWatcher) init() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if m.eventC == nil {
				return
			}
			if rand.Int()%2 == 0 {
				m.eventC <- config.SourceEvent(config.NewSource("mock2", []byte(content2), "json"))
			} else {
				m.eventC <- config.SourceEvent(config.NewSource("mock", []byte(content), "json"))
			}
		}
	}()
}

func (m *MockWatcher) Notify(eventC chan<- config.Event) {
	m.eventC = eventC
}

func (m *MockWatcher) StopNotify(eventC chan<- config.Event) {
	m.eventC = nil
}

func (m *MockWatcher) Close(ctx context.Context) error {
	return nil
}

type MockDecoder struct {
}

func (md MockDecoder) IsSupported(extension string) bool {
	return true
}

func (md MockDecoder) Decode(data []byte, m map[string]any) error {
	return json.Unmarshal(data, &m)
}

func TestConfigure(t *testing.T) {
	configure, err := config.NewConfigure(
		context.Background(),
		config.Resources(&MockResource{}),
		config.Decoders(&MockDecoder{}))
	assert.NoError(t, err)
	familyLastName, err := configure.Get("family.last_name").String()
	assert.NoError(t, err)
	assert.Equal(t, "Hashimoto", familyLastName)

	eo, err := configure.Get("emails.0").String()
	assert.NoError(t, err)
	assert.Equal(t, "one", eo)

	type Location struct {
		City string `mapstructure:"City" json:"city" yaml:"city"`
	}

	var l Location
	err = configure.Get("location").Scan(&l)
	assert.NoError(t, err)
	assert.Equal(t, Location{City: "San Francisco"}, l)

	watcher, err := configure.Watch(context.Background())
	assert.NoError(t, err)

	eventC := make(chan config.Event)
	watcher.Notify(eventC)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range eventC {
			source, err := event.Get()
			if err != nil {
				t.Log(err)
			} else {
				t.Log(string(source.Value()))
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(15 * time.Second)
		watcher.StopNotify(eventC)
		close(eventC)
		_ = watcher.Close(context.Background())
	}()

	wg.Wait()
}
