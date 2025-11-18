package check

import (
	"cmp"
)

// OrUnwish 从数组中返回第一个既不等于unwish值也不等于零值的元素
//
// 参数:
//   - unwish: 不希望得到的值
//   - arr: 要搜索的值的可变参数数组
//
// 返回值:
//   - 返回第一个既不等于unwish值也不等于零值的元素，如果没有找到则返回零值
func OrUnwish[T comparable](unwish T, arr ...T) (res T) {
	var zero = res
	for _, v := range arr {
		if v != unwish && v != zero {
			return v
		}
	}
	return zero
}

// Or 返回数组中第一个非零值元素
// 使用标准库的cmp.Or函数实现
//
// 参数:
//   - arr: 要搜索的值的可变参数数组
//
// 返回值:
//   - 返回数组中第一个非零值元素
func Or[T comparable](arr ...T) T {
	return cmp.Or(arr...)
}
