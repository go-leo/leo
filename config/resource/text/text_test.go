package text_test

import (
	"codeup.aliyun.com/qimao/leo/leo/config/resource/text"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestText(t *testing.T) {
	txt := text.NewText("hello")
	txt.Get()
	assert.Equal(t, "hello", txt.Get())
	observer1 := text.ObserverFunc(func(newText, oldText string) {
		t.Log(newText, oldText)
	})
	txt.AddObserver(observer1)
	observer2 := text.ObserverFunc(func(newText, oldText string) {
		t.Log(newText, oldText)
	})
	txt.AddObserver(observer2)
	txt.Set("world")

	txt.RemoveObserver(observer2)

	txt.Set("hello leo")
}
