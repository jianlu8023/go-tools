package check

type reduceFunc[T any] func(T, T) T

func Reduce[T any](fn reduceFunc[T], arr ...T) (res T) {
	if len(arr) == 0 {
		return
	}
	res = arr[0]
	l := len(arr)
	for i := 1; i < l; i++ {
		res = fn(res, arr[i])
	}
	return res
}
