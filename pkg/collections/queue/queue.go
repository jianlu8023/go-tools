package queue

// queue 是基础队列的实现
type queue[T any] struct {
	data []T
}

// New 创建一个新的队列
func New[T any]() Queue[T] {
	return &queue[T]{
		data: make([]T, 0),
	}
}

// Enqueue 向队列尾部添加元素
func (q *queue[T]) Enqueue(element T) Queue[T] {
	q.data = append(q.data, element)
	return q
}

// Dequeue 从队列头部移除并返回元素
func (q *queue[T]) Dequeue() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}

	element := q.data[0]
	q.data = q.data[1:]
	return element, true
}

// Front 返回队列头部元素但不移除
func (q *queue[T]) Front() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}

	return q.data[0], true
}

// Len 返回队列长度
func (q *queue[T]) Len() int {
	return len(q.data)
}

// Empty 判断队列是否为空
func (q *queue[T]) Empty() bool {
	return len(q.data) == 0
}

// Clear 清空队列
func (q *queue[T]) Clear() Queue[T] {
	q.data = make([]T, 0)
	return q
}
