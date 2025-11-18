package check

// IF 根据条件返回两个值中的一个
//
// 参数:
//   - check: 布尔条件
//   - rtTrue: 当条件为true时返回的值
//   - rtFalse: 当条件为false时返回的值
//
// 返回值:
//   - 当条件为true时返回rtTrue，否则返回rtFalse
//
// 示例:
//   result := IF(true, "yes", "no")
//   // result 将会是 "yes"
func IF[T any](check bool, rtTrue T, rtFalse T) T {
	if check {
		return rtTrue
	}
	return rtFalse
}
