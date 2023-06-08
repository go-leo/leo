package timex

import "time"

func Date(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func Today(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return Date(time.Now())
	}
	if ts[0].IsZero() {
		return Date(time.Now())
	}
	return Date(ts[0])
}

func Tomorrow(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return Today().AddDate(0, 0, 1)
	}
	if ts[0].IsZero() {
		return Today().AddDate(0, 0, 1)
	}
	return ts[0].AddDate(0, 0, 1)
}

func Yesterday(ts ...time.Time) time.Time {
	if len(ts) <= 0 {
		return Today().AddDate(0, 0, -1)
	}
	if ts[0].IsZero() {
		return Today().AddDate(0, 0, -1)
	}
	return ts[0].AddDate(0, 0, -1)
}
