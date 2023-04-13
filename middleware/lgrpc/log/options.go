package log

type options struct {
	Payload          bool
	Skips            []string
	PayloadWhenError bool
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func (o *options) init() {

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
