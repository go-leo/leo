package binding

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx/formx"
)

func Uri(m map[string][]string, obj any, tag string) error {
	return formx.Unmarshal(m, obj, tag)
}
