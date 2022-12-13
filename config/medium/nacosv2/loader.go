package nacosv2

import (
	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/log"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var _ config.Loader = new(Loader)

type Loader struct {
	group       string
	dataId      string
	contentType string
	data        string
	log         log.Logger
	client      config_client.IConfigClient
}

func (loader *Loader) ContentType() string {
	return loader.contentType
}

func (loader *Loader) Load() error {
	loader.log.Infof("get config DataId: %s, Group: %s", loader.dataId, loader.group)
	content, err := loader.client.GetConfig(vo.ConfigParam{
		DataId: loader.dataId,
		Group:  loader.group,
	})
	if err != nil {
		return err
	}
	loader.data = content
	return nil
}

func (loader *Loader) RawData() []byte {
	return []byte(loader.data)
}

type LoaderOption func(loader *Loader)

func Logger(log log.Logger) LoaderOption {
	return func(loader *Loader) {
		loader.log = log
	}
}

func NewLoader(
	client config_client.IConfigClient,
	group string,
	dataId string,
	contentType string,
	opts ...LoaderOption,
) *Loader {
	loader := &Loader{
		group:       group,
		dataId:      dataId,
		contentType: contentType,
		data:        "",
		log:         log.Discard{},
		client:      client,
	}
	for _, opt := range opts {
		opt(loader)
	}
	return loader
}
