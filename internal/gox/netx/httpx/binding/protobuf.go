package binding

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/protobufx"
	"net/http"
)

func ProtoBuf(req *http.Request, obj any) error {
	return protobufx.NewDecoder(req.Body).Decode(obj)
}
