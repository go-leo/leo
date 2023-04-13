package log

import "github.com/gin-gonic/gin"

type Skip func(ctx *gin.Context) bool

type options struct {
	Skips []Skip
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func (o *options) init() {

}

func Skips(skips ...Skip) Option {
	return func(o *options) {
		o.Skips = append(o.Skips, skips...)
	}
}
