package timex

import "time"

func Yesterday(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}
