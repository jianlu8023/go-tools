package queue

import (
	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
)

// queueIterator 是队列的迭代器实现
type queueIterator[T any] struct {
	elements []T
	index    int
	closed   bool
}

var _ concurrent.Iterator[int] = (*queueIterator[int])(nil)

// HasNext 移动到下一个元素
func (qi *queueIterator[T]) HasNext() bool {
	if qi.closed || qi.index >= len(qi.elements)-1 {
		return false
	}
	qi.index++
	return true
}

// Value 返回当前元素的值
func (qi *queueIterator[T]) Value() T {
	if qi.closed || qi.index < 0 || qi.index >= len(qi.elements) {
		var zero T
		return zero
	}
	return qi.elements[qi.index]
}

// Close 关闭迭代器，释放相关资源
func (qi *queueIterator[T]) Close() error {
	qi.closed = true
	qi.elements = nil
	return nil
}
