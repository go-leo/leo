package circuitbreaker

import (
	"time"

	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/streadway/handy/breaker"
)

type options struct {
	breakerCreator func() Breaker
}

type Option func(o *options)

func FailureRatio(failureRatio float64, window time.Duration, cooldown time.Duration, minObservations uint) Option {
	return func(o *options) {
		creator := func(failureRatio float64, window time.Duration, cooldown time.Duration, minObservations uint) func() Breaker {
			return func() Breaker {
				breaker := breaker.NewBreaker(failureRatio)
				if window > 0 {
					breaker.WithWindow(window)
				}
				if cooldown > 0 {
					breaker.WithCooldown(cooldown)
				}
				if minObservations > 0 {
					breaker.WithMinObservation(minObservations)
				}
				return &FailureRatioBreaker{breaker: breaker}
			}
		}
		o.breakerCreator = creator(failureRatio, window, cooldown, minObservations)
	}
}

func GoogleSRE(success float64, request int64, bucket int, window time.Duration) Option {
	return func(o *options) {
		creator := func(success float64, request int64, bucket int, window time.Duration) func() Breaker {
			return func() Breaker {
				var opts []sre.Option
				if success > 0 {
					opts = append(opts, sre.WithSuccess(success))
				}
				if request > 0 {
					opts = append(opts, sre.WithRequest(request))
				}
				if bucket > 0 {
					opts = append(opts, sre.WithBucket(bucket))
				}
				if window > 0 {
					opts = append(opts, sre.WithWindow(window))
				}
				breaker := sre.NewBreaker(opts...)
				return &GoogleSREBreaker{breaker: breaker}
			}
		}
		o.breakerCreator = creator(success, request, bucket, window)
	}
}

func (o *options) init() {
	if o.breakerCreator == nil {
		o.breakerCreator = func() Breaker {
			return &NoopBreaker{}
		}
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}
