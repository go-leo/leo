package render

import (
	"bytes"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/convx"
	"fmt"
	"html/template"
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/jsonx"
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
)

type indentedOptions struct {
	Prefix string
	Indent string
}

type secureOptions struct {
	Prefix string
}

type jsonpOptions struct {
	Callback string
}

type jsonOptions struct {
	IndentedOptions *indentedOptions
	SecureOptions   *secureOptions
	Pure            bool
	Ascii           bool
	JsonpOptions    *jsonpOptions
}

func (o *jsonOptions) apply(opts ...JSONOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *jsonOptions) init() {
	if o.IndentedOptions != nil {
		if stringx.IsBlank(o.IndentedOptions.Prefix) && stringx.IsBlank(o.IndentedOptions.Indent) {
			o.IndentedOptions.Prefix = ""
			o.IndentedOptions.Indent = "  "
		}
	}
}

type JSONOption func(o *jsonOptions)

func IndentedJSON(prefix, indent string) JSONOption {
	return func(o *jsonOptions) {
		o.IndentedOptions = &indentedOptions{
			Prefix: prefix,
			Indent: indent,
		}
	}
}

func PureJSON() JSONOption {
	return func(o *jsonOptions) {
		o.Pure = true
	}
}

func SecureJSON(prefix string) JSONOption {
	return func(o *jsonOptions) {
		o.SecureOptions = &secureOptions{Prefix: prefix}
	}
}

func JsonpJSON(callback string) JSONOption {
	return func(o *jsonOptions) {
		o.JsonpOptions = &jsonpOptions{Callback: callback}
	}
}

func AsciiJSON() JSONOption {
	return func(o *jsonOptions) {
		o.Ascii = true
	}
}

// JSON marshals the given interface object and writes it with custom ContentType.
func JSON(w http.ResponseWriter, data any, opts ...JSONOption) (err error) {
	o := &jsonOptions{}
	o.apply(opts...)
	o.init()
	if o.SecureOptions != nil {
		return secureJSON(w, data, o)
	}
	if o.Ascii {
		return asciiJSON(w, data)
	}
	if o.JsonpOptions != nil {
		return jsonpJSON(w, data, o)
	}
	return sampleJson(w, data, o)
}

func secureJSON(w http.ResponseWriter, data any, o *jsonOptions) error {
	writeContentType(w, []string{"application/json; charset=utf-8"})
	jsonBytes, err := jsonx.Marshal(data)
	if err != nil {
		return err
	}
	// if the jsonBytes is array values
	if bytes.HasPrefix(jsonBytes, convx.StringToBytes("[")) &&
		bytes.HasSuffix(jsonBytes, convx.StringToBytes("]")) {
		if _, err = w.Write(convx.StringToBytes(o.SecureOptions.Prefix)); err != nil {
			return err
		}
	}
	_, err = w.Write(jsonBytes)
	return err
}

func asciiJSON(w http.ResponseWriter, data any) (err error) {
	writeContentType(w, []string{"application/json"})
	ret, err := jsonx.Marshal(data)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	for _, r := range convx.BytesToString(ret) {
		cvt := string(r)
		if r >= 128 {
			cvt = fmt.Sprintf("\\u%04x", int64(r))
		}
		buffer.WriteString(cvt)
	}

	_, err = w.Write(buffer.Bytes())
	return err
}

func jsonpJSON(w http.ResponseWriter, Data any, o *jsonOptions) error {
	writeContentType(w, []string{"application/javascript; charset=utf-8"})
	ret, err := jsonx.Marshal(Data)
	if err != nil {
		return err
	}

	if o.JsonpOptions.Callback == "" {
		_, err = w.Write(ret)
		return err
	}

	callback := template.JSEscapeString(o.JsonpOptions.Callback)
	if _, err = w.Write(convx.StringToBytes(callback)); err != nil {
		return err
	}

	if _, err = w.Write(convx.StringToBytes("(")); err != nil {
		return err
	}

	if _, err = w.Write(ret); err != nil {
		return err
	}

	if _, err = w.Write(convx.StringToBytes(");")); err != nil {
		return err
	}

	return nil
}

func sampleJson(w http.ResponseWriter, data any, o *jsonOptions) error {
	writeContentType(w, []string{"application/json; charset=utf-8"})
	encoder := jsonx.NewEncoder(w)
	if o.IndentedOptions != nil {
		encoder.SetIndent(o.IndentedOptions.Prefix, o.IndentedOptions.Indent)
	}
	if o.Pure {
		encoder.SetEscapeHTML(false)
	}
	return encoder.Encode(data)
}
