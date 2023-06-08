package jsonx

import "codeup.aliyun.com/qimao/leo/leo/internal/gox/encodingx"

type JSONEncoder interface {
	encodingx.Encoder
	SetIndent(prefix, indent string)
	SetEscapeHTML(escapeHTML bool)
}
