package binding

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/xmlx"
	"net/http"
)

func XML(req *http.Request, obj any) error {
	return xmlx.NewDecoder(req.Body).Decode(obj)
}
