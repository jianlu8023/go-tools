package check

import (
	"time"
)

type dateIterFunc = func(yield func(V time.Time) bool)

func DateRanger(start, end time.Time, step time.Duration) dateIterFunc {
	return func(yield func(V time.Time) bool) {
		for start.Before(end) {
			if !yield(start) {
				return
			}
			start = start.Add(step)
		}
	}
}

type timeIterFunc = func(yield func(K, V time.Time) bool)

func TimeRanger(start, end time.Time, step time.Duration) timeIterFunc {
	return func(yield func(K, V time.Time) bool) {
		for start.Before(end) {
			if !yield(start, start.Add(step)) {
				return
			}
			start = start.Add(step)
		}
	}
}
