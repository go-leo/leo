package filex_test

import (
	"testing"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/osx/filex"
)

func TestHumanReadableSize(t *testing.T) {
	t.Log(filex.HumanReadableSize(1))
	t.Log(filex.HumanReadableSize(10))
	t.Log(filex.HumanReadableSize(100))
	t.Log(filex.HumanReadableSize(1000))
	t.Log(filex.HumanReadableSize(10000))
	t.Log(filex.HumanReadableSize(100000))
	t.Log(filex.HumanReadableSize(1000000))
	t.Log(filex.HumanReadableSize(10000000))
	t.Log(filex.HumanReadableSize(100000000))
	t.Log(filex.HumanReadableSize(1000000000))
	t.Log(filex.HumanReadableSize(10000000000))
	t.Log(filex.HumanReadableSize(100000000000))
	t.Log(filex.HumanReadableSize(1000000000000))
	t.Log(filex.HumanReadableSize(10000000000000))
	t.Log(filex.HumanReadableSize(100000000000000))
}
