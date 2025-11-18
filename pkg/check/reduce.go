package check

// reduceFunc 定义了一个归约函数类型，接收两个类型T的值并返回一个类型T的值
// 用作Reduce函数中的归约函数类型
type reduceFunc[T any] func(T, T) T

// Reduce 将归约函数应用于值切片，并返回一个归约后的结果
//
// 参数:
//   - fn: 归约函数，接收两个类型T的值并返回一个类型T的值
//   - arr: 要归约的类型T的值的可变参数切片
//
// 返回值:
//   - 归约后的结果值
//
// 示例:
//   numbers := []int{1, 2, 3, 4, 5}
//   sum := Reduce(func(a, b int) int { return a + b }, numbers...)
//   // sum 将会是 15
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
