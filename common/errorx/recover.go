package errorx

// PanicHandler is called for recovering from panics spawned internally to the library (and thus
// not recoverable by the caller's goroutine).
type PanicHandler func(any)

func WithRecover(handler PanicHandler, fn func()) {
	defer func() {
		// if handler to nil, which means panics are not recovered.
		if handler != nil {
			if err := recover(); err != nil {
				handler(err)
			}
		}
	}()
	fn()
}
