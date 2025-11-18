package check

// In 检查目标值是否在数组中存在
//
// 参数:
//   - target: 要检查的目标值
//   - arr: 要搜索的值的可变参数数组
//
// 返回值:
//   - 如果目标值在数组中存在则返回true，否则返回false
//
// 示例:
//   result := In(3, 1, 2, 3, 4, 5)
//   // result 将会是 true
func In[T comparable](target T, arr ...T) bool {
	for _, obj := range arr {
		if target == obj {
			return true
		}
	}
	return false
}
