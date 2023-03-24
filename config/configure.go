package config

import (
	"context"
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
	valuer  *valuer
	parser  *parser
}

func NewConfigure(opts ...Option) *Configure {
	o := &options{}
	o.apply(opts...)
	o.init()
	configurer := &Configure{
		valuer: &valuer{},
		parser: &parser{Decoders: o.Decoders},
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
	configure.valuer.data = d
	return nil
}

func (configure *Configure) Get(key string) (*Value, error) {
	return configure.valuer.Value(key)
}
