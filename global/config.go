package global

import (
	"sync"

	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/config/factory"
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
	mgr, err := factory.NewConfigMgr(configURLs)
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
