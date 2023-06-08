package binding

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/tomlx"
	"net/http"
)

func TOML(req *http.Request, obj any) error {
	return tomlx.NewDecoder(req.Body).Decode(obj)
}
