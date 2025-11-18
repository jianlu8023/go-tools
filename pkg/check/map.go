package check

type mapFunc[T, V any] func(T) V

func Map[T, V any](fn mapFunc[T, V], arr ...T) []V {
	res := make([]V, 0, len(arr))
	for _, v := range arr {
		res = append(res, fn(v))
	}
	return res
}
