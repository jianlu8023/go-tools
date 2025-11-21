package concurrent

// Entry 定义Map的键值对条目
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Iterator 定义迭代器接口
type Iterator[T any] interface {
	// HasNext 移动到下一个元素，如果还有元素则返回true，否则返回false
	HasNext() bool
	// Value 返回当前元素的值
	Value() T
	// Close 关闭迭代器，释放相关资源
	Close() error
}
