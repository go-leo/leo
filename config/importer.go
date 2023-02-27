package config

import (
	"context"

	"github.com/go-leo/gox/slicex"
)

type Importer struct {
	name    string
	loader  Loader
	watcher Watcher
	decoder Decoder
}

func (importer *Importer) apply(opts ...ImporterOption) {
	for _, opt := range opts {
		opt(importer)
	}
}

func (importer *Importer) init() {
	if importer.loader == nil {
		importer.loader = &noopLoader{}
	}
	if importer.watcher == nil {
		importer.watcher = &noopWatcher{}
	}
	if importer.decoder == nil {
		importer.decoder = &noopDecoder{}
	}
}

type ImporterOption func(*Importer)

func WithLoader(loader Loader) ImporterOption {
	return func(resource *Importer) {
		resource.loader = loader
	}
}

func WithWatcher(watcher Watcher) ImporterOption {
	return func(resource *Importer) {
		resource.watcher = watcher
	}
}

func WithDecoder(decoder Decoder) ImporterOption {
	return func(resource *Importer) {
		resource.decoder = decoder
	}
}

func NewImporter(name string, opts ...ImporterOption) *Importer {
	importer := &Importer{name: name}
	importer.apply(opts...)
	importer.init()
	return importer
}

func (importer *Importer) importConfig() (map[string]any, error) {
	data, err := importer.loader.Load()
	if err != nil {
		return nil, err
	}
	configMap := make(map[string]any)
	err = importer.decoder.Decode(data, configMap)
	if err != nil {
		return nil, err
	}
	return configMap, nil
}

func (importer *Importer) startWatchConfig(
	ctx context.Context,
	onChanged func(string, []byte, string),
	onError func(string, error, string),
) error {
	eventC, err := importer.watcher.StartWatch(ctx)
	if err != nil {
		return err
	}
	go func() {
		for event := range eventC {
			if slicex.IsNotEmpty(event.Data()) && onChanged != nil {
				onChanged(importer.String(), event.Data(), event.Description())
			}
			if event.Err() != nil && onError != nil {
				onError(importer.String(), event.Err(), event.Description())
			}
		}
	}()
	return nil
}

func (importer *Importer) stopWatchConfig(ctx context.Context) error {
	return importer.watcher.StopWatch(ctx)
}

func (importer *Importer) String() string {
	return importer.name
}
