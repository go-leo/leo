package binding

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/formx"
	"net/http"
)

func Query(req *http.Request, obj any, tag string) error {
	return formx.Unmarshal(req.URL.Query(), obj, tag)
}
