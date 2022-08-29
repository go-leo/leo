package concurrentx

func BraveGo(f func(), r func(p any)) {
	go BraveDo(f, r)
}

func BraveDo(f func(), r func(p any)) {
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
}
