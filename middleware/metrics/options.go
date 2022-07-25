package metrics

type options struct {
	// Skips is a grpc full method array or url path array which metrics are collected.
	Skips []string
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func (o *options) init() {
}

func WithSkips(skips ...string) Option {
	return func(o *options) {
		o.Skips = append(o.Skips, skips...)
	}
}
