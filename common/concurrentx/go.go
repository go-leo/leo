package concurrentx

func BraveGo(f func(), r func(p any)) {
	go func() {
		defer func() {
			// if r is nil, which means panics are not recovered.
			if r == nil {
				return
			}
			if p := recover(); p != nil {
				r(p)
			}
		}()
		f()
	}()
}
