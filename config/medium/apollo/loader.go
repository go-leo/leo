package apollo

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-leo/netx/httpx"
	"github.com/go-leo/stringx"

	"github.com/go-leo/filex"

	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/log"
)

var _ config.Loader = new(Loader)

type Loader struct {
	scheme        string
	host          string
	port          string
	appID         string
	cluster       string
	namespaceName string
	secret        string
	timeout       time.Duration
	client        *http.Client
	rawData       []byte
	log           log.Logger
}

func (loader *Loader) ContentType() string {
	contentType := filex.ExtName(loader.namespaceName)
	if stringx.IsBlank(contentType) {
		return "properties"
	}
	return contentType
}

func (loader *Loader) Load() error {
	content, err := loader.getConfigFromApollo()
	if err != nil {
		return err
	}
	loader.log.Debug("config content:", content)
	loader.rawData = []byte(content)
	return nil
}

func (loader *Loader) getConfigFromApollo() (string, error) {
	uri := fmt.Sprintf("%s://%s:%s/configs/%s/%s/%s", loader.scheme, loader.host, loader.port, loader.appID, loader.cluster, loader.namespaceName)
	loader.log.Info("reading apollo config:", uri)
	ctx := context.Background()
	if loader.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, loader.timeout)
		defer cancel()
	}
	return getConfigContent(ctx, uri, loader.appID, loader.secret, loader.ContentType(), loader.client)
}

func (loader *Loader) RawData() []byte {
	return loader.rawData
}

type LoaderOption func(loader *Loader)

func Scheme(scheme string) LoaderOption {
	return func(loader *Loader) {
		loader.scheme = scheme
	}
}

func Secret(secret string) LoaderOption {
	return func(loader *Loader) {
		loader.secret = secret
	}
}

func Timeout(timeout time.Duration) LoaderOption {
	return func(loader *Loader) {
		loader.timeout = timeout
	}
}

func Logger(log log.Logger) LoaderOption {
	return func(loader *Loader) {
		loader.log = log
	}
}

func NewLoader(host string, port string, appID string, cluster string, namespaceName string, opts ...LoaderOption) *Loader {
	loader := &Loader{
		scheme:        "http",
		host:          host,
		port:          port,
		appID:         appID,
		cluster:       cluster,
		namespaceName: namespaceName,
		log:           log.Discard{},
		client:        httpx.PooledClient(),
	}
	for _, opt := range opts {
		opt(loader)
	}
	return loader
}
