package binding

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/msgpackx"
	"net/http"
)

func MsgPack(req *http.Request, obj any) error {
	return msgpackx.NewDecoder(req.Body).Decode(obj)
}
