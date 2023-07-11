package stringx_test

import (
	"testing"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
)

type str string

func TestBlank(t *testing.T) {
	t.Log(stringx.IsNotBlank(""))
	t.Log(stringx.IsNotBlank(" "))
	t.Log(stringx.IsNotBlank("	 "))
	t.Log(stringx.IsNotBlank("1"))
	t.Log(stringx.IsNotBlank("2 "))

	t.Log(stringx.IsNotBlank(str("")))
	t.Log(stringx.IsNotBlank(str(" ")))
	t.Log(stringx.IsNotBlank(str("	 ")))
	t.Log(stringx.IsNotBlank(str("1")))
	t.Log(stringx.IsNotBlank(str("2 ")))
}
