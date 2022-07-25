package lb

import "context"

type TargetTaker interface {
	Target() string
}

type LoadBalancer interface {
	Pick(ctx context.Context, takers []TargetTaker) (TargetTaker, error)
}
