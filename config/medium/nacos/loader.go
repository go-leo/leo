package nacos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-leo/netx/httpx"

	"github.com/go-leo/leo/v2/config"
	"github.com/go-leo/leo/v2/log"
)

var _ config.Loader = new(Loader)

type Loader struct {
	scheme      string
	host        string
	port        int
	namespace   string
	group       string
	dataId      string
	contentType string
	data        []byte
	log         log.Logger
	client      *http.Client
}

func (loader *Loader) ContentType() string {
	return loader.contentType
}

func (loader *Loader) Load() error {
	loader.log.Infof("get config DataId: %s, Group: %s", loader.dataId, loader.group)
	uri := fmt.Sprintf("%s://%s:%d/nacos/v1/cs/configs", loader.scheme, loader.host, loader.port)
	helper := httpx.NewRequestBuilder().
		Get().
		URLString(uri).
		Query("tenant", loader.namespace).
		Query("group", loader.group).
		Query("dataId", loader.dataId).
		Execute(context.Background(), loader.client)
	if err := helper.Err(); err != nil {
		return err
	}
	content, err := helper.BytesBody()
	if err != nil {
		return err
	}
	loader.data = content
	return nil
}

func (loader *Loader) RawData() []byte {
	return loader.data
}

type LoaderOption func(loader *Loader)

func Scheme(scheme string) LoaderOption {
	return func(loader *Loader) {
		loader.scheme = scheme
	}
}

func Logger(log log.Logger) LoaderOption {
	return func(loader *Loader) {
		loader.log = log
	}
}

func NewLoader(
	host string,
	port int,
	namespace string,
	group string,
	dataId string,
	contentType string,
	opts ...LoaderOption,
) *Loader {
	loader := &Loader{
		scheme:      "http",
		host:        host,
		port:        port,
		namespace:   namespace,
		group:       group,
		dataId:      dataId,
		contentType: contentType,
		data:        nil,
		log:         log.Discard{},
		client:      httpx.PooledClient(),
	}
	for _, opt := range opts {
		opt(loader)
	}
	return loader
}
