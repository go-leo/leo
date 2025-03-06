package cqrs

import (
	"fmt"
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

type Bird interface {
	Fly()
}

type Eagle struct {
}

func (e *Eagle) Fly() {
	fmt.Println("Eagle is flying")
}

func fly[T Bird]() {
	var bird T
	bird.Fly()
}

func TestName(t *testing.T) {
	fly[*Eagle]()
}
