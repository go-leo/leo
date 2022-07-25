package circuitbreaker

import (
	"time"

	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/streadway/handy/breaker"
)

// Breaker is an interface representing the ability to conditionally allow
// requests to pass, and to report on the result of passed requests.
type Breaker interface {
	Allow() bool
	Success()
	Failure()
}

type NoopBreaker struct{}

func (breaker *NoopBreaker) Allow() bool { return true }

func (breaker *NoopBreaker) Success() {}

func (breaker *NoopBreaker) Failure() {}

type FailureRatioBreaker struct {
	breaker breaker.Breaker
}

func (breaker *FailureRatioBreaker) Allow() bool {
	return breaker.breaker.Allow()
}

func (breaker *FailureRatioBreaker) Success() {
	breaker.breaker.Success(time.Duration(0))
}

func (breaker *FailureRatioBreaker) Failure() {
	breaker.breaker.Failure(time.Duration(0))
}

type GoogleSREBreaker struct {
	breaker circuitbreaker.CircuitBreaker
}

func (breaker *GoogleSREBreaker) Allow() bool {
	return breaker.breaker.Allow() == nil
}

func (breaker *GoogleSREBreaker) Success() {
	breaker.breaker.MarkSuccess()
}

func (breaker *GoogleSREBreaker) Failure() {
	breaker.breaker.MarkFailed()
}
