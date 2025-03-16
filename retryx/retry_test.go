package retryx

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"testing"
)

func TestDuration(t *testing.T) {
	duration := durationpb.New(-1)
	t.Log(duration.AsDuration())
}
