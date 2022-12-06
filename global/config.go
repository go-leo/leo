package global

import (
	"sync"

	leoconfig "github.com/go-leo/leo/v2/config"
)

var (
	config       any
	configMgr    *leoconfig.Mgr
	configLocker sync.RWMutex
)

func Config() any {
	configLocker.RLock()
	defer configLocker.RUnlock()
	return config
}

func SetConfig(conf any) func() {
	configLocker.Lock()
	defer configLocker.Unlock()
	prev := config
	config = conf
	return func() { SetConfig(prev) }
}

func ConfigMgr() *leoconfig.Mgr {
	configLocker.RLock()
	defer configLocker.RUnlock()
	return configMgr
}

func SetConfigMgr(mgr *leoconfig.Mgr) func() {
	configLocker.Lock()
	defer configLocker.Unlock()
	prev := configMgr
	configMgr = mgr
	return func() { SetConfigMgr(prev) }
}
