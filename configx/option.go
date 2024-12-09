package configx

// options 结构体用于存储配置项，包括资源、解析器和合并器。
type options struct {
	Resources  []Resource
	Parsers    []Parser
	Merger     Merger
	BufferSize int
}

func newOptions() *options {
	return &options{
		Merger:     &merger{},
		BufferSize: 16,
	}
}

func (o *options) Apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Option func(*options)

// WithResource 函数用于设置 Resource 选项。
func WithResource(resources ...Resource) Option {
	return func(o *options) {
		o.Resources = append(o.Resources, resources...)
	}
}

// WithParser 函数用于设置 Parser 选项。
func WithParser(parsers ...Parser) Option {
	return func(o *options) {
		o.Parsers = append(o.Parsers, parsers...)
	}
}

// WithMerger 函数用于设置 Merger 选项。
func WithMerger(merger Merger) Option {
	return func(o *options) {
		o.Merger = merger
	}
}

// WithBufferSize 函数用于设置 BufferSize 选项。
func WithBufferSize(size int) Option {
	return func(o *options) {
		o.BufferSize = size
	}
}
