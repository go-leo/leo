package binding

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/yamlx"
	"net/http"
)

func YAML(req *http.Request, obj any) error {
	return yamlx.NewDecoder(req.Body).Decode(obj)
}
