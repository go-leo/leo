package render

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/jsonx"
	"net/http"
)

type jsonOptions struct {
	Pure bool
}

func (o *jsonOptions) apply(opts ...JSONOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *jsonOptions) init() {
}

type JSONOption func(o *jsonOptions)

func PureJSON() JSONOption {
	return func(o *jsonOptions) {
		o.Pure = true
	}
}

// JSON marshals the given interface object and writes it with custom ContentType.
func JSON(w http.ResponseWriter, data any, opts ...JSONOption) (err error) {
	o := &jsonOptions{}
	o.apply(opts...)
	o.init()
	return sampleJson(w, data, o)
}

func sampleJson(w http.ResponseWriter, data any, o *jsonOptions) error {
	writeContentType(w, []string{"application/json; charset=utf-8"})
	encoder := jsonx.NewEncoder(w)
	if o.Pure {
		encoder.SetEscapeHTML(false)
	}
	return encoder.Encode(data)
}
