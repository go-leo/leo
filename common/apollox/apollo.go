package apollox

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)

// apollo://ip:port?app_id=x&cluster=x&namespace=x&is_backup_config=true&backup_config_path=x&secret=x&sync_server_timeout=1s

func NewClient(uriStr string) (agollo.Client, error) {
	uri, err := url.Parse(uriStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse uri %s, %w", uriStr, err)
	}
	query := uri.Query()
	isBackupConfig, _ := strconv.ParseBool(query.Get("is_backup_config"))
	syncServerTimeout, _ := time.ParseDuration(query.Get("sync_server_timeout"))

	c := &config.AppConfig{
		AppID:             query.Get("app_id"),
		Cluster:           query.Get("cluster"),
		NamespaceName:     query.Get("namespace"),
		IP:                uri.Host,
		IsBackupConfig:    isBackupConfig,
		BackupConfigPath:  query.Get("backup_config_path"),
		Secret:            query.Get("secret"),
		SyncServerTimeout: int(syncServerTimeout / time.Second),
	}
	return agollo.StartWithConfig(func() (*config.AppConfig, error) { return c, nil })

}
