package check

// mapFunc 定义了一个函数类型，接收类型T的值并返回类型V的值
// 用作Map函数中的映射函数类型
type mapFunc[T, V any] func(T) V

// Map 将映射函数应用于值切片，并返回一个新切片
// 新切片包含对每个输入值应用映射函数后的结果
//
// 参数:
//   - fn: 映射函数，接收类型T的值并返回类型V的值
//   - arr: 要映射的类型T的值的可变参数切片
//
// 返回值:
//   - 一个新切片，包含对每个输入值应用映射函数后的结果
//
// 示例:
//   numbers := []int{1, 2, 3, 4, 5}
//   squared := Map(func(n int) int { return n * n }, numbers...)
//   // squared 将会是 [1, 4, 9, 16, 25]
func Map[T, V any](fn mapFunc[T, V], arr ...T) []V {
	res := make([]V, 0, len(arr))
	for _, v := range arr {
		res = append(res, fn(v))
	}
	return res
}
