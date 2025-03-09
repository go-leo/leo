package retryx

import (
	"github.com/go-leo/gox/backoff"
)

// BackoffFunc 退避函数
type BackoffFunc = backoff.Factory

// Constant 常数退避
// Constant it waits for a fixed period of time between calls.
func Constant() BackoffFunc {
	return backoff.ConstantFactory()
}

// Linear 线性退避
// Linear it waits for "delta * attempt" time between calls.
func Linear() BackoffFunc {
	return backoff.LinearFactory()
}

// Exponential 指数退避
// Exponential it waits for "delta * e^attempts" time between calls.
func Exponential() BackoffFunc {
	return backoff.ExponentialFactory()
}

// Exponential2 指数退避
// Exponential2 it waits for "delta * 2^attempts" time between calls.
func Exponential2() BackoffFunc {
	return backoff.Exponential2Factory()
}

// Fibonacci 斐波那契退避
// Fibonacci it waits for "delta * fibonacci(attempt)" time between calls.
func Fibonacci() BackoffFunc {
	return backoff.FibonacciFactory()
}

// JitterUp 退避函数支持抖动
// JitterUp adds random jitter to the interval.
//
// This adds or subtracts time from the interval within a given jitter fraction.
// For example for 10s and jitter 0.1, it will return a time within [9s, 11s])
func JitterUp(f BackoffFunc, jitter float64) BackoffFunc {
	return backoff.JitterUpFactory(f, jitter)
}
