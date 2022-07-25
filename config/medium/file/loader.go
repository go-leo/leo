package file

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/go-leo/leo/common/stringx"
	"github.com/go-leo/leo/config"
	"github.com/go-leo/leo/log"
)

var _ config.Loader = new(Loader)

type Loader struct {
	filename    string
	contentType string
	rawData     []byte
	log         log.Logger
}

func (loader *Loader) ContentType() string {
	return loader.contentType
}

func (loader *Loader) Load() error {
	loader.log.Info("reading file:", loader.filename)
	rawData, err := ioutil.ReadFile(loader.filename)
	if err != nil {
		return err
	}
	loader.log.Debug("file content:", string(rawData))
	loader.rawData = rawData
	return nil
}

func (loader *Loader) RawData() []byte {
	return loader.rawData
}

type LoaderOption func(loader *Loader)

func ContentType(contentType string) LoaderOption {
	return func(loader *Loader) {
		loader.contentType = contentType
	}
}

func Logger(log log.Logger) LoaderOption {
	return func(loader *Loader) {
		loader.log = log
	}
}

func NewLoader(filename string, opts ...LoaderOption) *Loader {
	loader := &Loader{filename: filename, log: log.Discard{}}
	for _, opt := range opts {
		opt(loader)
	}
	if stringx.IsBlank(loader.contentType) {
		loader.contentType = filepath.Ext(filename)
		loader.contentType = strings.TrimPrefix(loader.contentType, ".")
	}
	return loader
}
