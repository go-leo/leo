package circuitbreakerx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"sync"
	"time"
)

// Counts holds the numbers of requests and their successes/failures.
// CircuitBreaker clears the internal Counts either
// on the change of the state or at the closed-state intervals.
// Counts ignores the results of the requests sent before clearing.
type Counts struct {
	// Requests is the number of requests.
	Requests uint32
	// TotalSuccesses is the number of successes.
	TotalSuccesses uint32
	// TotalFailures is the number of failures.
	TotalFailures uint32
	// ConsecutiveSuccesses is the number of consecutive successes.
	ConsecutiveSuccesses uint32
	// ConsecutiveFailures is the number of consecutive failures.
	ConsecutiveFailures uint32
}

// onRequest increments the number of requests.
func (c *Counts) onRequest() {
	c.Requests++
}

// onSuccess increments the number of successes.
func (c *Counts) onSuccess() {
	c.TotalSuccesses++
	c.ConsecutiveSuccesses++
	c.ConsecutiveFailures = 0
}

// onFailure increments the number of failures.
func (c *Counts) onFailure() {
	c.TotalFailures++
	c.ConsecutiveFailures++
	c.ConsecutiveSuccesses = 0
}

// clear clears all metrics.
func (c *Counts) clear() {
	c.Requests = 0
	c.TotalSuccesses = 0
	c.TotalFailures = 0
	c.ConsecutiveSuccesses = 0
	c.ConsecutiveFailures = 0
}

// CircuitBreaker is a state machine to prevent sending requests that are likely to fail.
type CircuitBreaker struct {

	// maxRequests is the maximum number of requests allowed to pass through when the CircuitBreaker is half-open.
	maxRequests uint32

	// interval is the cyclic period of the closed state
	interval time.Duration

	// timeout is the period of the open state,
	// The purpose of the timeout timer is to give the system time to rectify the problem that caused the failure
	// before allowing the application to attempt to perform the operation again.
	timeout time.Duration

	// readyToTrip is called with a copy of Counts whenever a request fails in the closed state.
	readyToTrip func(counts Counts) bool

	// isSuccessful is called with the error returned from a request.
	// If IsSuccessful returns true, the error is counted as a success.
	// Otherwise, the error is counted as a failure.
	isSuccessful func(err error) bool

	mutex      sync.Mutex
	state      State
	generation uint64
	counts     Counts

	// expiry 记录当前状态的有效期
	expiry time.Time
}

type Option func(cb *CircuitBreaker)

// MaxRequests is the maximum number of requests allowed to pass through
// when the CircuitBreaker is half-open.
// The Half-Open state is useful to prevent a recovering service from suddenly being inundated with requests.
// As a service recovers, it may be able to support a limited volume of requests until the recovery is complete,
// but while recovery is in progress a flood of work may cause the service to time out or fail again.
// If MaxRequests is 0, the CircuitBreaker allows only 1 request.
func MaxRequests(maxRequests uint32) Option {
	return func(cb *CircuitBreaker) {
		cb.maxRequests = maxRequests
	}
}

// Interval is the cyclic period of the closed state
// for the CircuitBreaker to clear the internal Counts.
// If Interval is 0, the CircuitBreaker doesn't clear internal Counts during the closed state.
func Interval(interval time.Duration) Option {
	return func(cb *CircuitBreaker) {
		cb.interval = interval
	}
}

// Timeout is the period of the open state,
// after which the state of the CircuitBreaker becomes half-open.
// If Timeout is 0, the timeout value of the CircuitBreaker is set to 60 seconds.
func Timeout(timeout time.Duration) Option {
	return func(cb *CircuitBreaker) {
		cb.timeout = timeout
	}
}

// ReadyToTrip is called with a copy of Counts whenever a request fails in the closed state.
// If ReadyToTrip returns true, the CircuitBreaker will be placed into the open state.
// If ReadyToTrip is nil, default ReadyToTrip is used.
// Default ReadyToTrip returns true when the number of consecutive failures is more than 5.
func ReadyToTrip(readyToTrip func(counts Counts) bool) Option {
	return func(cb *CircuitBreaker) {
		cb.readyToTrip = readyToTrip
	}
}

func isSuccessful(isSuccessful func(err error) bool) Option {
	return func(cb *CircuitBreaker) {
		cb.isSuccessful = isSuccessful
	}
}

// NewCircuitBreaker returns a new CircuitBreaker configured with the given Options.
func NewCircuitBreaker(opts ...Option) *CircuitBreaker {
	cb := &CircuitBreaker{
		maxRequests: 1,
		interval:    0,
		timeout:     60 * time.Second,
		readyToTrip: func(counts Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
		isSuccessful: func(err error) bool {
			return err == nil
		},
		mutex:      sync.Mutex{},
		state:      StateClosed,
		generation: 0,
		counts:     Counts{},
		expiry:     time.Time{},
	}
	for _, opt := range opts {
		opt(cb)
	}
	// set initial generation
	cb.toNewGeneration(time.Now())
	return cb
}

// State returns the current state of the CircuitBreaker.
func (cb *CircuitBreaker) State() State {
	cb.mutex.Lock()
	state, _ := cb.currentState(time.Now())
	cb.mutex.Unlock()
	return state
}

// Execute runs the given request if the CircuitBreaker accepts it.
// Execute returns an error instantly if the CircuitBreaker rejects the request.
// Otherwise, Execute returns the result of the request.
// If a panic occurs in the request, the CircuitBreaker handles it as an error
// and causes the same panic again.
func (cb *CircuitBreaker) Execute(ctx context.Context, request any, next endpoint.Endpoint) (any, error, bool) {
	generation, err := cb.beforeRequest()
	if err != nil {
		return nil, nil, false
	}

	defer func() {
		if e := recover(); e != nil {
			cb.afterRequest(generation, false)
			panic(e)
		}
	}()

	response, err := next(ctx, request)
	cb.afterRequest(generation, cb.isSuccessful(err))
	return response, err, true
}

// beforeRequest returns the current generation and an error if the CircuitBreaker rejects the request.
func (cb *CircuitBreaker) beforeRequest() (uint64, error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	// The state of the CircuitBreaker is determined by the current generation.
	state, generation := cb.currentState(time.Now())
	switch state {
	case StateClosed:
		// The request from the application is routed through to the operation.
	case StateOpen:
		// The request from the application fails immediately and an exception is returned to the application.
		return generation, errors.New("circuitbreakerx: circuit breaker is open")
	case StateHalfOpen:
		// The Half-Open state is useful to prevent a recovering service from suddenly being inundated with requests.
		// As a service recovers, it may be able to support a limited volume of requests until the recovery is complete,
		// but while recovery is in progress a flood of work may cause the service to time out or fail again.
		if cb.counts.Requests >= cb.maxRequests {
			return generation, errors.New("circuitbreakerx: too many requests")
		}
		// single test
	}
	cb.counts.onRequest()
	return generation, nil
}

func (cb *CircuitBreaker) afterRequest(before uint64, success bool) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	now := time.Now()
	state, generation := cb.currentState(now)
	if generation != before {
		return
	}

	if success {
		cb.onSuccess(state, now)
	} else {
		cb.onFailure(state, now)
	}
}

// onSuccess is called when the request is successful.
func (cb *CircuitBreaker) onSuccess(state State, now time.Time) {
	switch state {
	case StateClosed:
		// If the request is successful, the counter is incremented.
		cb.counts.onSuccess()
		return
	case StateHalfOpen:
		// If the request is successful, the counter is incremented.
		cb.counts.onSuccess()
		// If these requests are successful, it is assumed that the fault that was previously causing the failure has
		// been fixed and the circuit breaker switches to the Closed state (the failure counter is reset).
		if cb.IsRecovered() {
			cb.setState(StateClosed, now)
		}
	case StateOpen:
	}
}

// onFailure is called when the request is failed.
func (cb *CircuitBreaker) onFailure(state State, now time.Time) {
	switch state {
	case StateClosed:
		// If the request is fails, the counter is decremented.
		cb.counts.onFailure()
		// If the number of recent failures exceeds a specified threshold within a given time period, the proxy is
		// placed into the Open state.
		if cb.readyToTrip(cb.counts) {
			cb.setState(StateOpen, now)
		}
	case StateHalfOpen:
		// If any request fails, the circuit breaker assumes that the fault is still present so it reverts back to the
		// Open state and restarts the timeout timer to give the system a further period of time to recover from the failure.
		cb.setState(StateOpen, now)
	case StateOpen:
	}
}

// IsRecovered returns true if the CircuitBreaker is in the half-open state.
func (cb *CircuitBreaker) IsRecovered() bool {
	return cb.counts.ConsecutiveSuccesses >= cb.maxRequests
}

// currentState returns the current state of the CircuitBreaker and the current generation.
func (cb *CircuitBreaker) currentState(now time.Time) (State, uint64) {
	switch cb.state {
	case StateClosed:
		// expiry 时间到达后，内部的统计计数将被重置。
		if !cb.expiry.IsZero() && cb.expiry.Before(now) {
			cb.toNewGeneration(now)
		}
	case StateHalfOpen:
	case StateOpen:
		// 从 StateOpen 转换到 StateHalfOpen
		// 当 CircuitBreaker 处于 StateOpen 状态时，expiry 记录了 StateOpen 状态结束的时间点。
		// 一旦这个时间点过去，CircuitBreaker 将自动切换到 StateHalfOpen，允许有限数量的新请求尝试。
		if cb.expiry.Before(now) {
			cb.setState(StateHalfOpen, now)
		}
	}
	return cb.state, cb.generation
}

// setState sets the state of the CircuitBreaker.
func (cb *CircuitBreaker) setState(state State, now time.Time) {
	if cb.state == state {
		return
	}
	cb.state = state
	cb.toNewGeneration(now)
}

// toNewGeneration 方法的作用是在 CircuitBreaker 类中推进其内部的状态和统计信息到一个新的“世代”或周期。具体来说，它有以下几个主要作用：
// 更新状态有效期：根据 CircuitBreaker 当前的状态，toNewGeneration 方法会设置或重置 expiry 时间。例如，在 StateClosed 状态下，如果设置了 Interval，则 expiry 会被设置为当前时间加上 Interval 的未来某个时刻；而在 StateOpen 状态下，expiry 会被设置为当前时间加上 Timeout 的未来某个时刻。
// 增加世代计数：每次调用 toNewGeneration 方法时，generation 变量都会递增。这有助于检测并防止过时的状态检查结果被用来做出决定，因为每个请求在执行前后都会检查 generation 是否发生了变化。
// 通过这些操作，toNewGeneration 方法帮助维持了断路器的状态机的正确性和时效性，确保了断路器能够在适当的时候基于最新的统计数据做出正确的状态转换决策。这对于实现一个高效且响应迅速的服务调用管理机制至关重要。
func (cb *CircuitBreaker) toNewGeneration(now time.Time) {
	cb.generation++
	// 重置统计信息：每当 CircuitBreaker 进入一个新状态或者在关闭状态下达到了一个新的统计周期时，它会调用 toNewGeneration 方法来清空之前的请求统计信息（如成功次数、失败次数等）。
	// 这有助于避免旧数据对新的状态决策产生影响。
	cb.counts.clear()

	var zero time.Time
	switch cb.state {
	case StateClosed:
		// 当 CircuitBreaker 处于 StateClosed 状态且设置了 Interval 时，expiry 记录了下一次清除统计计数的时间点。
		// 如果在这段时间内没有达到 ReadyToTrip 的条件，那么在 expiry 时间到达后，内部的统计计数将被重置。
		if cb.interval == 0 {
			cb.expiry = zero
		} else {
			cb.expiry = now.Add(cb.interval)
		}
	case StateOpen:
		// 当 CircuitBreaker 处于 StateOpen 状态时，expiry 记录了 StateOpen 状态结束的时间点。
		cb.expiry = now.Add(cb.timeout)
	case StateHalfOpen:
		cb.expiry = zero
	}
}
