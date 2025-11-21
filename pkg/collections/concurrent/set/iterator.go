package set

// setIterator 实现Iterator接口，用于遍历Set的元素
type setIterator[T comparable] struct {
	elements []T
	index    int
	closed   bool
}

// HasNext 移动到下一个元素
func (it *setIterator[T]) HasNext() bool {
	if it.closed || it.index >= len(it.elements) {
		return false
	}
	it.index++
	return it.index < len(it.elements)
}

// Value 返回当前元素的值
func (it *setIterator[T]) Value() T {
	if it.closed || it.index >= len(it.elements) {
		var empty T
		return empty
	}
	return it.elements[it.index]
}

// Close 关闭迭代器
func (it *setIterator[T]) Close() error {
	it.closed = true
	it.elements = nil
	return nil
}
