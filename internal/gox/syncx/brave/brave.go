package brave

func Do(f func(), rs ...func(p any)) {
	defer func() {
		if len(rs) <= 0 {
			return
		}
		r := rs[0]
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			r(p)
		}
	}()
	f()
}

func Go(f func(), rs ...func(p any)) {
	go Do(f, rs...)
}

func DoE(f func() error, rs ...func(p any) error) (err error) {
	defer func() {
		if len(rs) <= 0 {
			return
		}
		r := rs[0]
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			err = r(p)
		}
	}()
	return f()
}

func GoE(f func() error, rs ...func(p any) error) <-chan error {
	errC := make(chan error)
	go func() {
		defer close(errC)
		err := DoE(f, rs...)
		if err != nil {
			errC <- err
		}
	}()
	return errC
}

func DoRE(f func() (any, error), rs ...func(p any) error) (ret any, err error) {
	defer func() {
		if len(rs) <= 0 {
			return
		}
		r := rs[0]
		if r == nil {
			return
		}
		if p := recover(); p != nil {
			err = r(p)
		}
	}()
	return f()
}

func GoRE(f func() (any, error), rs ...func(p any) error) (<-chan any, <-chan error) {
	retC := make(chan any)
	errC := make(chan error)
	go func() {
		defer close(errC)
		defer close(retC)
		ret, err := DoRE(f, rs...)
		if err != nil {
			errC <- err
			return
		}
		retC <- ret
	}()
	return retC, errC
}
