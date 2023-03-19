package nacos

import "testing"

func TestChannel(t *testing.T) {
	ch := make(chan error)
	nch := rch(ch)
	t.Log(ch == nch)
}

func rch(ch chan error) <-chan error {
	return ch
}
