package circuitbreakerx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

// Gutter Breaker

// GoogleSreBreaker
// see: https://landing.google.com/sre/sre-book/chapters/handling-overload/
type GoogleSreBreaker struct {
}

func (breaker *GoogleSreBreaker) Execute(ctx context.Context, request any, next endpoint.Endpoint) (any, bool, error) {

}

// rejectionProbability
// requests
//   - The number of requests attempted by the application layer(at the client, on top of the adaptive throttling system).
//
// accepts
//   - The number of requests accepted by the backend.
//
// k
//   - accepts multiplier.
//     Reducing the multiplier will make adaptive throttling behave more aggressively.
//     Increasing the multiplier will make adaptive throttling behave less aggressively.
//
// see: https://sre.google/sre-book/handling-overload/#eq2101
func rejectionProbability(requests int, accepts int, k float64) float64 {
	return max(0, float64(requests)-k*float64(accepts)/float64(requests+1))
}
