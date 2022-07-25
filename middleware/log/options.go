package log

type options struct {
	// Payload is true, the interceptor logs the payloads of requests and responses.
	Payload bool
	// Skips is a grpc full method array or url path array which metrics are collected.
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
