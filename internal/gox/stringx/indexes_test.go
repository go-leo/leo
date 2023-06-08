package stringx_test

import (
	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
	"testing"
)

func TestIndexes(t *testing.T) {
	indexes := stringx.Indices("codeup.aliyun.com/qimao/leo/leo/log/slog_test.TestLog", "/")
	t.Log(indexes)

	indexes = stringx.Indices("Hello, World! World is beautiful.", "World")
	t.Log(indexes)

	indexes = stringx.Indices("Hello, World! World is beautiful.", "bye")
	t.Log(indexes)

	indexes = stringx.Indices("Hello, World! World is beautiful.", "Hello, World! World is beautiful")
	t.Log(indexes)
	indexes = stringx.Indices("Hello, World! World is beautiful.", "Hello, World! World is beautiful.")
	t.Log(indexes)
}
