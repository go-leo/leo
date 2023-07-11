package log_test

import (
	"testing"

	"codeup.aliyun.com/qimao/leo/leo/log"
)

func TestParseLevel(t *testing.T) {
	level, err := log.ParseLevel("debug")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(level)

	level, err = log.ParseLevel("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(level)

	level, err = log.ParseLevel("info")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(level)

	level, err = log.ParseLevel("warn")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(level)

	level, err = log.ParseLevel("panic")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(level)

	level, err = log.ParseLevel("fatal")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(level)

	level, err = log.ParseLevel("fatal+")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(level)
}
