package config

import (
	"context"
	"errors"

	"github.com/go-leo/gox/stringx"
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
}

func NewConfigure(opts ...Option) *Configure {
	o := &options{}
	o.apply(opts...)
	o.init()
	configurer := &Configure{
		options: o,
		parser:  &parser{Decoders: o.Decoders},
	}
	return configurer
}

func (configure *Configure) Import(ctx context.Context) error {
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

func (configure *Configure) Get(key string) *Value {
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

func (configure *Configure) Watch(ctx context.Context) (Watcher, error) {
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
