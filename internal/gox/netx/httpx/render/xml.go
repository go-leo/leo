package render

import (
	"net/http"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/xmlx"
)

// XML encodes the given interface object and writes data with custom ContentType.
func XML(w http.ResponseWriter, data any) error {
	writeContentType(w, []string{"application/xml; charset=utf-8"})
	return xmlx.NewEncoder(w).Encode(data)
}
