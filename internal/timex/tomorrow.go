package timex

import "time"

func Tomorrow(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}
