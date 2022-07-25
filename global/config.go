package global

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/config/medium/apollo"
	"github.com/go-leo/leo/config/medium/file"
	"github.com/go-leo/leo/config/medium/nacos"
	"github.com/go-leo/leo/config/parser"
	"github.com/go-leo/leo/config/valuer"

	"github.com/go-leo/leo/common/nacosx"
	"github.com/go-leo/leo/common/stringx"
)

// --config deployments/bff/local/configs/bff.yaml
// --config nacos://ip1:port1,ip2:port2?contentType=yaml&namespace=ns&group=g&dataID=d
// --config apollo://ip:port?appID=x&cluster=x&namespaceName=x&secret=x

var (
	configuration Config
	configMgr     *config.Mgr
	configLocker  sync.RWMutex
)

func Configuration() Config {
	configLocker.RLock()
	defer configLocker.RUnlock()
	return configuration
}

func SetConfig(conf Config) func() {
	configLocker.Lock()
	defer configLocker.Unlock()
	prev := configuration
	configuration = conf
	return func() { SetConfig(prev) }
}

func ConfigMgr() *config.Mgr {
	configLocker.RLock()
	defer configLocker.RUnlock()
	return configMgr
}

func SetConfigMgr(mgr *config.Mgr) func() {
	configLocker.Lock()
	defer configLocker.Unlock()
	prev := configMgr
	configMgr = mgr
	return func() { SetConfigMgr(prev) }
}

type Property struct {
	Key string
	Val any
}

func initConfig(configURLs string, properties []*Property) error {
	configLocker.Lock()
	defer configLocker.Unlock()
	mgr, err := newConfigMgr(configURLs)
	if err != nil {
		return err
	}
	configMgr = mgr
	if err := configMgr.ReadConfig(); err != nil {
		return err
	}
	if err := configMgr.Unmarshal(&configuration); err != nil {
		return err
	}
	for _, property := range properties {
		if err := configMgr.UnmarshalKey(property.Key, property.Val); err != nil {
			return err
		}
	}
	configuration.Init()
	return nil
}

func newConfigMgr(configURLs string) (*config.Mgr, error) {
	configPaths := strings.Split(configURLs, ",")
	loaders := make([]config.Loader, 0)
	watchers := make([]config.Watcher, 0)
	for _, confPath := range configPaths {
		uri, err := url.Parse(confPath)
		if err != nil {
			return nil, fmt.Errorf("failed parse url %s, %w", configURLs, err)
		}
		switch uri.Scheme {
		case "":
			loader, watcher := newFileLoaderAndWatcher(uri)
			loaders = append(loaders, loader)
			if watcher != nil {
				watchers = append(watchers, watcher)
			}
		case "nacos":
			loader, watcher, err := newNanosLoaderAndWatcher(uri)
			if err != nil {
				return nil, err
			}
			loaders = append(loaders, loader)
			if watcher != nil {
				watchers = append(watchers, watcher)
			}
		case "apollo":
			loader, watcher, err := newApolloLoaderAndWatcher(uri)
			if err != nil {
				return nil, err
			}
			loaders = append(loaders, loader)
			if watcher != nil {
				watchers = append(watchers, watcher)
			}
		}
	}
	manager := config.NewManager(
		config.WithLoader(loaders...),
		config.WithWatcher(watchers...),
		config.WithParser(parser.Parsers()...),
		config.WithValuer(valuer.NewTrieTreeValuer()),
	)
	return manager, nil
}

func newFileLoaderAndWatcher(uri *url.URL) (config.Loader, config.Watcher) {
	// 创建file loader
	loader := file.NewLoader(uri.String())
	query := uri.Query()
	// 是否需要watch
	isWatch, _ := strconv.ParseBool(query.Get("watch"))
	if !isWatch {
		return loader, nil
	}
	// 返回
	return loader, file.NewWatcher(uri.String())
}

func newNanosLoaderAndWatcher(uri *url.URL) (config.Loader, config.Watcher, error) {
	client, err := nacosx.NewNacosConfigClient(uri)
	if err != nil {
		return nil, nil, err
	}
	query := uri.Query()
	group := query.Get("group")
	dataID := query.Get("dataID")
	contentType := query.Get("contentType")
	loader := nacos.NewLoader(client, group, dataID, contentType)

	isWatch, _ := strconv.ParseBool(query.Get("watch"))
	if !isWatch {
		return loader, nil, nil
	}
	return loader, nacos.NewWatcher(client, group, dataID), nil
}

func newApolloLoaderAndWatcher(uri *url.URL) (config.Loader, config.Watcher, error) {
	host := uri.Hostname()
	port := uri.Port()
	query := uri.Query()
	appID := query.Get("appID")
	cluster := query.Get("cluster")
	namespaceName := query.Get("namespaceName")
	secret := query.Get("secret")

	var loaderOpts []apollo.LoaderOption
	var watcherOpts []apollo.WatcherOption
	if stringx.IsNotBlank(secret) {
		loaderOpts = append(loaderOpts, apollo.Secret(secret))
		watcherOpts = append(watcherOpts, apollo.WithSecret(secret))
	}
	loader := apollo.NewLoader(host, port, appID, cluster, namespaceName, loaderOpts...)

	isWatch, _ := strconv.ParseBool(query.Get("watch"))
	if !isWatch {
		return loader, nil, nil
	}
	return loader, apollo.NewWatcher(host, port, appID, cluster, namespaceName, watcherOpts...), nil
}
