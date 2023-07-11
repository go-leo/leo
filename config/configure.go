package config

import (
	"context"
	"errors"
	"sync"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
)

var (
	ErrValueNotFound = errors.New("value not found")

	ErrValueIsNil = errors.New("value is nil")
)

type options struct {
	Resources []Resource
	Decoders  []Decoder
}

type Option func(o *options)

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
}

func Resources(resources ...Resource) Option {
	return func(o *options) {
		o.Resources = append(o.Resources, resources...)
	}
}

func Decoders(decoders ...Decoder) Option {
	return func(o *options) {
		o.Decoders = append(o.Decoders, decoders...)
	}
}

type Configure struct {
	options *options
	parser  *parser
	data    *Data
	mutex   sync.RWMutex
}

func (configure *Configure) Refresh(ctx context.Context) error {
	configure.mutex.Lock()
	defer configure.mutex.Unlock()
	return configure.read(ctx)
}

func (configure *Configure) Get(key string) *Value {
	configure.mutex.RLock()
	defer configure.mutex.RUnlock()
	return configure.get(key)
}

func (configure *Configure) Watch(ctx context.Context) (Watcher, error) {
	configure.mutex.RLock()
	defer configure.mutex.RUnlock()
	return configure.watch(ctx)
}

func (configure *Configure) read(ctx context.Context) error {
	var ds []*Data
	for _, resource := range configure.options.Resources {
		source, err := resource.Load(ctx)
		if err != nil {
			return err
		}
		d, err := configure.parser.Parse(source)
		if err != nil {
			return err
		}
		ds = append(ds, d)
	}
	d, err := mutilData(ds...)
	if err != nil {
		return err
	}
	configure.data = d
	return nil
}

func (configure *Configure) get(key string) *Value {
	if stringx.IsBlank(key) {
		return &Value{val: configure.data.AsMap()}
	}
	node, ok := configure.data.AsTree().Find(key)
	if !ok {
		return &Value{err: ErrValueNotFound}
	}
	meta := node.Meta()
	if meta == nil {
		return &Value{err: ErrValueIsNil}
	}
	return &Value{val: meta}
}

func (configure *Configure) watch(ctx context.Context) (Watcher, error) {
	var watchers []Watcher
	for _, resource := range configure.options.Resources {
		watcher, err := resource.Watch(ctx)
		if err != nil {
			return nil, err
		}
		watchers = append(watchers, watcher)
	}
	return MultiWatcher(watchers...), nil
}

// Config is an interface abstraction for dynamic configuration
type Config interface {
	Refresh(ctx context.Context) error
	Get(key string) *Value
	Watch(ctx context.Context) (Watcher, error)
}

func NewConfigure(ctx context.Context, opts ...Option) (Config, error) {
	o := &options{}
	o.apply(opts...)
	o.init()
	configure := &Configure{
		options: o,
		parser:  &parser{Decoders: o.Decoders},
	}
	defaultConfig = configure
	return configure, configure.read(ctx)
}

// Default Config Manager
var defaultConfig Config

func SetConfig(l Config) {
	defaultConfig = l
}

// Get a value from the config
func Get(path string) *Value {
	return defaultConfig.Get(path)
}

// Watch a value for changes
func Watch(ctx context.Context) (Watcher, error) {
	return defaultConfig.Watch(ctx)
}

func Refresh(ctx context.Context) error {
	return defaultConfig.Refresh(ctx)
}
