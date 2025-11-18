package check

import (
	"time"
)

// dateIterFunc 定义了一个日期迭代器函数类型
// 用于DateRanger函数的返回值类型
type dateIterFunc = func(yield func(V time.Time) bool)

// DateRanger 创建一个日期范围迭代器，从start到end，步长为step
//
// 参数:
//   - start: 开始时间
//   - end: 结束时间
//   - step: 时间步长
//
// 返回值:
//   - 一个日期迭代器函数，可以用于遍历日期范围
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

// timeIterFunc 定义了一个时间迭代器函数类型
// 用于TimeRanger函数的返回值类型
type timeIterFunc = func(yield func(K, V time.Time) bool)

// TimeRanger 创建一个时间范围迭代器，从start到end，步长为step
//
// 参数:
//   - start: 开始时间
//   - end: 结束时间
//   - step: 时间步长
//
// 返回值:
//   - 一个时间迭代器函数，可以用于遍历时间范围
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
