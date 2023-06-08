package render

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/tomlx"
)

// TOML marshals the given interface object and writes data with custom ContentType.
func TOML(w http.ResponseWriter, data any) error {
	writeContentType(w, []string{"application/toml; charset=utf-8"})
	bytes, err := tomlx.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(bytes)
	return err
}
