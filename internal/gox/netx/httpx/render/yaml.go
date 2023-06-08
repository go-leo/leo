package render

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/yamlx"
)

// YAML marshals the given interface object and writes data with custom ContentType.
func YAML(w http.ResponseWriter, Data any) error {
	writeContentType(w, []string{"application/x-yaml; charset=utf-8"})
	bytes, err := yamlx.Marshal(Data)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}
