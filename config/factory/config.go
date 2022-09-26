package factory

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/config/medium/apollo"
	"github.com/go-leo/leo/config/medium/file"
	"github.com/go-leo/leo/config/medium/nacos"
	"github.com/go-leo/leo/config/parser"
	"github.com/go-leo/leo/config/valuer"

	"github.com/go-leo/stringx"
)

type Property struct {
	Key string
	Val any
}

func NewConfigMgr(configURLs string) (*config.Mgr, error) {
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
			loader, watcher := NewFileLoaderAndWatcher(uri)
			loaders = append(loaders, loader)
			if watcher != nil {
				watchers = append(watchers, watcher)
			}
		case "nacos":
			loader, watcher, err := NewNanosLoaderAndWatcher(uri)
			if err != nil {
				return nil, err
			}
			loaders = append(loaders, loader)
			if watcher != nil {
				watchers = append(watchers, watcher)
			}
		case "apollo":
			loader, watcher, err := NewApolloLoaderAndWatcher(uri)
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

func NewFileLoaderAndWatcher(uri *url.URL) (config.Loader, config.Watcher) {
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

func NewNanosLoaderAndWatcher(uri *url.URL) (config.Loader, config.Watcher, error) {
	query := uri.Query()
	group := query.Get("group")
	dataID := query.Get("dataID")
	contentType := query.Get("contentType")
	loader := nacos.NewLoader("", "", "", group, dataID, contentType)

	isWatch, _ := strconv.ParseBool(query.Get("watch"))
	if !isWatch {
		return loader, nil, nil
	}
	return loader, nacos.NewWatcher("", "", "", group, dataID), nil
}

func NewApolloLoaderAndWatcher(uri *url.URL) (config.Loader, config.Watcher, error) {
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
