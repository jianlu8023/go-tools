package check

// filterFunc 定义了一个函数类型，接收类型T的值并返回布尔值
// 用作Filter函数中的过滤函数类型
type filterFunc[T any] func(T) bool

// Filter 将过滤函数应用于值切片，并返回一个新切片
// 该切片仅包含过滤函数返回true的值
//
// 参数:
//   - fn: 过滤函数，接收类型T的值并返回布尔值
//   - arr: 要过滤的类型T的值的可变参数切片
//
// 返回值:
//   - 一个新切片，仅包含过滤函数返回true的值
//
// 示例:
//   numbers := []int{1, 2, 3, 4, 5}
//   evenNumbers := Filter(func(n int) bool { return n%2 == 0 }, numbers...)
//   // evenNumbers 将会是 [2, 4]
func Filter[T any](fn filterFunc[T], arr ...T) []T {
	res := make([]T, 0, len(arr))
	for _, v := range arr {
		if ok := fn(v); ok {
			res = append(res, v)
		}
	}
	return res
}
