package recovery

import (
	"fmt"
)

var _ error = new(PanicError)

type PanicError struct {
	err   error
	stack []byte
}

func (p *PanicError) Error() string {
	return fmt.Sprintln(p.err, "panic", "stack", "...\n"+string(p.stack))
}
