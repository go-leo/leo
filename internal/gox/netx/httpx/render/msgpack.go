package render

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/msgpackx"
)

// MsgPack encodes the given interface object and writes data with custom ContentType.
func MsgPack(w http.ResponseWriter, data any) error {
	writeContentType(w, []string{"application/msgpack; charset=utf-8"})
	return msgpackx.NewEncoder(w).Encode(data)
}
