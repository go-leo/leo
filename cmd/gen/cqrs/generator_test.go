package cqrs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDir(t *testing.T) {
	wd, _ := os.Getwd()
	dir := wd
	for {
		t.Log(dir)
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
}
