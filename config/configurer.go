package config

import (
	"context"
	"errors"
)

type Configurer struct {
	importers []*Importer
	valuer    *Valuer
	onChanged func(importerName string, data []byte, desc string)
	onError   func(importerName string, err error, desc string)
}

func (configurer *Configurer) apply(opts ...ConfigurerOption) {
	for _, opt := range opts {
		opt(configurer)
	}
}

func (configurer *Configurer) init() {
	if configurer.onChanged == nil {
		configurer.onChanged = func(importerName string, data []byte, desc string) {}
	}
	if configurer.onError == nil {
		configurer.onError = func(string string, err error, desc string) {}
	}
}

type ConfigurerOption func(*Configurer)

func WithImporter(importers ...*Importer) ConfigurerOption {
	return func(configurer *Configurer) {
		configurer.importers = append(configurer.importers, importers...)
	}
}

func WithOnChanged(f func(importerName string, data []byte, desc string)) ConfigurerOption {
	return func(configurer *Configurer) {
		configurer.onChanged = f
	}
}

func WithOnError(f func(importerName string, err error, desc string)) ConfigurerOption {
	return func(configurer *Configurer) {
		configurer.onError = f
	}
}

func NewConfigurer(opts ...ConfigurerOption) *Configurer {
	configurer := &Configurer{valuer: newValuer()}
	configurer.apply(opts...)
	configurer.init()
	return configurer
}

func (configurer *Configurer) Read() error {
	for _, importer := range configurer.importers {
		configMap, err := importer.importConfig()
		if err != nil {
			return err
		}
		err = configurer.valuer.Merge(configMap)
		if err != nil {
			return err
		}
	}
	return nil
}

func (configurer *Configurer) Get(key string) (*Value, error) {
	return configurer.valuer.Value(key)
}

func (configurer *Configurer) StartWatch(ctx context.Context) error {
	var errs []error
	for _, importer := range configurer.importers {
		e := importer.startWatchConfig(ctx, configurer.onChanged, configurer.onError)
		if e != nil {
			errs = append(errs, e)
		}
	}
	return errors.Join(errs...)
}

func (configurer *Configurer) StopWatch(ctx context.Context) error {
	var errs []error
	for _, importer := range configurer.importers {
		e := importer.stopWatchConfig(ctx)
		if e != nil {
			errs = append(errs, e)
		}
	}
	return errors.Join(errs...)
}
