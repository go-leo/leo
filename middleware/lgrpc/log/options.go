package log

type options struct {
	Payload          bool
	Skips            []string
	PayloadWhenError bool
	ErrorChecker     func(err error) bool
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func (o *options) init() {
	if o.ErrorChecker == nil {
		o.ErrorChecker = func(err error) bool {
			return err != nil
		}
	}
}

func ErrorCheck(f func(err error) bool) Option {
	return func(o *options) {
		o.ErrorChecker = f
	}
}

func WithSkip(skips ...string) Option {
	return func(o *options) {
		o.Skips = append(o.Skips, skips...)
	}
}

func WithPayload() Option {
	return func(o *options) {
		o.Payload = true
	}
}

func WithPayloadWhenError() Option {
	return func(o *options) {
		o.PayloadWhenError = true
	}
}
