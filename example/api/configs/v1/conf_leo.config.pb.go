// Code generated by protoc-gen-go-leo. DO NOT EDIT.

package configs

import (
	context "context"
	protox "github.com/go-leo/gox/protox"
	configx "github.com/go-leo/leo/v3/configx"
	sync "sync"
)

var (
	_ApplicationConfig      = &Application{}
	_ApplicationConfigMutex sync.RWMutex
)

func GetApplicationConfig() *Application {
	_ApplicationConfigMutex.RLock()
	defer _ApplicationConfigMutex.RUnlock()
	return protox.Clone(_ApplicationConfig)
}

func SetApplicationConfig(conf *Application) {
	_ApplicationConfigMutex.Lock()
	_ApplicationConfig = protox.Clone(conf)
	_ApplicationConfigMutex.Unlock()
}

func LoadApplicationConfig(ctx context.Context, opts ...configx.Option) error {
	conf, err := configx.Load[*Application](ctx, opts...)
	if err != nil {
		return err
	}
	SetApplicationConfig(conf)
	return nil
}

func WatchApplicationConfig(ctx context.Context, opts ...configx.Option) error {
	confC, err := configx.Watch[*Application](ctx, opts...)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case conf := <-confC:
				SetApplicationConfig(conf)
			}
		}
	}()
	return nil
}

func LoadAndWatchApplicationConfig(ctx context.Context, opts ...configx.Option) error {
	if err := LoadApplicationConfig(ctx, opts...); err != nil {
		return err
	}
	return WatchApplicationConfig(ctx, opts...)
}
