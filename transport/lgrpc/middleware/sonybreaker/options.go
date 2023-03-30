package grpcsonybreaker

import (
	"time"

	"github.com/sony/gobreaker"
)

// Option represents the gobreaker.Settings
type Option func(*gobreaker.Settings)

// WithMaxRequests sets MaxRequests.
// MaxRequests is the maximum number of requests allowed to pass through
// when the CircuitBreaker is half-open.
// If MaxRequests is 0, the CircuitBreaker allows only 1 request.
func WithMaxRequests(maxRequests uint32) Option {
	return func(st *gobreaker.Settings) {
		st.MaxRequests = maxRequests
	}
}

// WithInterval sets Interval.
// Interval is the cyclic period of the closed state
// for the CircuitBreaker to clear the internal Counts.
// If Interval is 0, the CircuitBreaker doesn't clear internal Counts during the closed state.
func WithInterval(interval time.Duration) Option {
	return func(st *gobreaker.Settings) {
		st.Interval = interval
	}
}

// WithTimeout sets Timeout.
// Timeout is the period of the open state,
// after which the state of the CircuitBreaker becomes half-open.
// If Timeout is 0, the timeout value of the CircuitBreaker is set to 60 seconds.
func WithTimeout(timeout time.Duration) Option {
	return func(st *gobreaker.Settings) {
		st.Timeout = timeout
	}
}

// WithReadyToTrip set ReadyToTrip.
// ReadyToTrip is called with a copy of Counts whenever a request fails in the closed state.
// If ReadyToTrip returns true, the CircuitBreaker will be placed into the open state.
// If ReadyToTrip is nil, default ReadyToTrip is used.
// Default ReadyToTrip returns true when the number of consecutive failures is more than 5.
func WithReadyToTrip(readyToTrip func(counts gobreaker.Counts) bool) Option {
	return func(st *gobreaker.Settings) {
		st.ReadyToTrip = readyToTrip
	}
}

// WithOnStateChange sets OnStateChange.
// OnStateChange is called whenever the state of the CircuitBreaker changes.
func WithOnStateChange(onStateChange func(name string, from gobreaker.State, to gobreaker.State)) Option {
	return func(st *gobreaker.Settings) {
		st.OnStateChange = onStateChange
	}
}

func defaultSettings(name string) *gobreaker.Settings {
	return &gobreaker.Settings{
		Name:          name,
		MaxRequests:   0,
		Interval:      0,
		Timeout:       0,
		ReadyToTrip:   nil,
		OnStateChange: nil,
	}
}

func apply(st *gobreaker.Settings, opts ...Option) {
	for _, opt := range opts {
		opt(st)
	}
}
