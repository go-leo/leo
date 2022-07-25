package config

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

const (
	stop uint64 = iota
	start
)

// Mgr config manager, support mutil load
type Mgr struct {
	// loaders
	loaders  []Loader
	parsers  []Parser
	watchers []Watcher
	valuer   Valuer
	locker   sync.Mutex
	state    uint64
	stopFunc context.CancelFunc
}

type Option func(mgr *Mgr)

func WithLoader(loader ...Loader) Option {
	return func(mgr *Mgr) {
		mgr.loaders = append(mgr.loaders, loader...)
	}
}

func WithParser(parser ...Parser) Option {
	return func(mgr *Mgr) {
		mgr.parsers = append(mgr.parsers, parser...)
	}
}

func WithWatcher(watcher ...Watcher) Option {
	return func(mgr *Mgr) {
		mgr.watchers = append(mgr.watchers, watcher...)
	}
}

func WithValuer(valuer Valuer) Option {
	return func(mgr *Mgr) {
		mgr.valuer = valuer
	}
}

func NewManager(opts ...Option) *Mgr {
	m := new(Mgr)
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *Mgr) ReadConfig() error {
	loaders := m.loaders
	configs := make([]map[string]any, 0, len(loaders))
	for _, loader := range loaders {
		if err := loader.Load(); err != nil {
			return err
		}
		parser, err := m.getParser(loader.ContentType())
		if err != nil {
			return err
		}
		if err := parser.Parse(loader.RawData()); err != nil {
			return err
		}
		configs = append(configs, parser.ConfigMap())
	}
	m.valuer.AddConfig(configs...)
	return nil
}

func (m *Mgr) StartWatch(ctx context.Context) (<-chan *Event, error) {
	m.locker.Lock()
	defer m.locker.Unlock()
	if m.state == start {
		return nil, errors.New("watcher is started")
	}
	stopCtx, stopFunc := context.WithCancel(ctx)
	var eventCs []<-chan *Event
	for _, watcher := range m.watchers {
		eventC, err := watcher.Start(ctx)
		if err != nil {
			stopFunc()
			return nil, err
		}
		eventCs = append(eventCs, eventC)
	}
	eventC := make(chan *Event)
	var eg sync.WaitGroup
	for _, c := range eventCs {
		eg.Add(1)
		go func(c <-chan *Event) {
			defer eg.Done()
			for {
				select {
				case e := <-c:
					eventC <- e
				case <-stopCtx.Done():
					return
				}
			}
		}(c)
	}
	go func() {
		eg.Wait()
		close(eventC)
	}()
	m.state = start
	m.stopFunc = stopFunc
	return eventC, nil
}

func (m *Mgr) StopWatch(ctx context.Context) error {
	m.locker.Lock()
	defer m.locker.Unlock()
	if m.state != start {
		return errors.New("watcher is stopped")
	}
	var multiErr error
	for _, watcher := range m.watchers {
		err := watcher.Stop(ctx)
		if err != nil {
			multiErr = multierror.Append(multiErr, err)
		}
	}
	m.stopFunc()
	m.state = stop
	return multiErr
}

func (m *Mgr) AsMap() map[string]any {
	return m.valuer.Config()
}

func (m *Mgr) Get(key string) (any, error) {
	return m.valuer.Get(key)
}

// GetString returns the value associated with the key as a string.
func (m *Mgr) GetString(key string) (string, error) {
	str, err := m.Get(key)
	if err != nil {
		return "", err
	}
	return cast.ToStringE(str)
}

// GetInt64 returns the value associated with the key as an int64.
func (m *Mgr) GetInt64(key string) (int64, error) {
	str, err := m.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToInt64E(str)
}

// GetUint64 returns the value associated with the key as an uint64.
func (m *Mgr) GetUint64(key string) (uint64, error) {
	str, err := m.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToUint64E(str)
}

// GetFloat64 returns the value associated with the key as a float64.
func (m *Mgr) GetFloat64(key string) (float64, error) {
	str, err := m.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToFloat64E(str)
}

// GetBool returns the value associated with the key as a bool.
func (m *Mgr) GetBool(key string) (bool, error) {
	str, err := m.Get(key)
	if err != nil {
		return false, err
	}
	return cast.ToBoolE(str)
}

// GetTime returns the value associated with the key as time.
func (m *Mgr) GetTime(key string) (time.Time, error) {
	str, err := m.Get(key)
	if err != nil {
		return time.Time{}, err
	}
	return cast.ToTimeE(str)
}

// GetDuration returns the value associated with the key as a duration.
func (m *Mgr) GetDuration(key string) (time.Duration, error) {
	str, err := m.Get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToDurationE(str)
}

// GetSlice returns the value associated with the key as a slice of interface values.
func (m *Mgr) GetSlice(key string) ([]any, error) {
	str, err := m.Get(key)
	if err != nil {
		return nil, err
	}
	return cast.ToSliceE(str)
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (m *Mgr) GetStringMap(key string) (map[string]any, error) {
	str, err := m.Get(key)
	if err != nil {
		return nil, err
	}
	return cast.ToStringMapE(str)
}

// UnmarshalKey takes a single key and unmarshal it into a struct.
func (m *Mgr) UnmarshalKey(key string, rawVal any) error {
	input, err := m.Get(key)
	if err != nil {
		return err
	}
	c := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           rawVal,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToIPHookFunc(),
			mapstructure.StringToIPNetHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure.TextUnmarshallerHookFunc(),
		),
	}
	decoder, err := mapstructure.NewDecoder(c)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}

// Unmarshal merged config into a struct.
func (m *Mgr) Unmarshal(rawVal any) error {
	return m.UnmarshalKey("", rawVal)
}

func (m *Mgr) getParser(contentType string) (Parser, error) {
	for _, parser := range m.parsers {
		if parser.Support(contentType) {
			return parser, nil
		}
	}
	return nil, fmt.Errorf("not found %s parser", contentType)
}
