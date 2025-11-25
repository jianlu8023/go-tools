package queue

import (
	"sync"

	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
)

// RWQueue 是一个基于读写锁的并发安全队列实现
type RWQueue[T any] struct {
	mutex sync.RWMutex
	data  []T
}

var _ concurrent.Queue[int] = (*RWQueue[int])(nil)

// NewRWQueue 创建一个新的RWQueue实例
func NewRWQueue[T any]() concurrent.Queue[T] {
	return &RWQueue[T]{
		data: make([]T, 0, 8),
	}
}

// Enqueue 向队列尾部添加元素
func (rq *RWQueue[T]) Enqueue(element T) concurrent.Queue[T] {
	rq.mutex.Lock()
	defer rq.mutex.Unlock()
	rq.data = append(rq.data, element)
	return rq
}

// Dequeue 从队列头部移除并返回元素
func (rq *RWQueue[T]) Dequeue() (T, bool) {
	rq.mutex.Lock()
	defer rq.mutex.Unlock()

	if len(rq.data) == 0 {
		var zero T
		return zero, false
	}

	element := rq.data[0]
	rq.data = rq.data[1:]
	return element, true
}

// Front 返回队列头部元素但不移除
func (rq *RWQueue[T]) Front() (T, bool) {
	rq.mutex.RLock()
	defer rq.mutex.RUnlock()

	if len(rq.data) == 0 {
		var zero T
		return zero, false
	}

	return rq.data[0], true
}

// Len 返回队列长度
func (rq *RWQueue[T]) Len() int {
	rq.mutex.RLock()
	defer rq.mutex.RUnlock()
	return len(rq.data)
}

// Empty 判断队列是否为空
func (rq *RWQueue[T]) Empty() bool {
	rq.mutex.RLock()
	defer rq.mutex.RUnlock()
	return len(rq.data) == 0
}

// Clear 清空队列
func (rq *RWQueue[T]) Clear() concurrent.Queue[T] {
	rq.mutex.Lock()
	defer rq.mutex.Unlock()
	rq.data = make([]T, 0)
	return rq
}

// Iterator 返回队列的迭代器
func (rq *RWQueue[T]) Iterator() concurrent.Iterator[T] {
	rq.mutex.RLock()
	defer rq.mutex.RUnlock()

	elements := make([]T, len(rq.data))
	copy(elements, rq.data)

	return &queueIterator[T]{
		elements: elements,
		index:    -1,
	}
}
