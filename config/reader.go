package config

import (
	"context"
)

type Reader struct {
	Loader
	Watcher
	Decoder
}

func newReader(loader Loader, watcher Watcher, decoder Decoder) *Reader {
	return &Reader{Loader: loader, Watcher: watcher, Decoder: decoder}
}

func (r *Reader) Load() ([]byte, error) {
	return r.Loader.Load()
}

func (r *Reader) StartWatch(ctx context.Context) (<-chan Event, error) {
	return r.Watcher.StartWatch(ctx)
}

func (r *Reader) StopWatch(ctx context.Context) error {
	return r.Watcher.StopWatch(ctx)
}

func (r *Reader) Decode(data []byte, configMap map[string]interface{}) error {
	return r.Decoder.Decode(data, configMap)
}
