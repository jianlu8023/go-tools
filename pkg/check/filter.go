package check

type filterFunc[T any] func(T) bool

func Filter[T any](fn filterFunc[T], arr ...T) []T {
	res := make([]T, 0, len(arr))
	for _, v := range arr {
		if ok := fn(v); ok {
			res = append(res, v)
		}
	}
	return res
}
