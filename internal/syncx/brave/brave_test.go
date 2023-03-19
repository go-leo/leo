package brave

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	Do(func() {
		panic("this is a panic")
	}, func(p any) {
		t.Log(p)
	})
}

func TestGo(t *testing.T) {
	Go(func() {
		panic("this is a panic")
	}, func(p any) {
		t.Log(p)
	})
	time.Sleep(time.Second)
}

func TestDoE(t *testing.T) {
	err := DoE(func() error {
		return errors.New("this is do error")
	}, func(p any) error {
		return fmt.Errorf("%v", p)
	})
	t.Log(err)
}

func TestDoEWithPanic(t *testing.T) {
	err := DoE(func() error {
		panic("this is a panic")
		return nil
	}, func(p any) error {
		return fmt.Errorf("%v", p)
	})
	t.Log(err)
}

func TestGoE(t *testing.T) {
	errC := GoE(func() error {
		return errors.New("this is do error")
	}, func(p any) error {
		return fmt.Errorf("%v", p)
	})
	t.Log(<-errC)
}

func TestGoEWithPanic(t *testing.T) {
	errC := GoE(func() error {
		panic("this is a panic")
		return nil
	}, func(p any) error {
		return fmt.Errorf("%v", p)
	})
	t.Log(<-errC)
}
