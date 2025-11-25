package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Enqueue(t *testing.T) {
	q := New[int]()
	assert.True(t, q.Empty())
	assert.Equal(t, 0, q.Len())

	q.Enqueue(1)
	assert.False(t, q.Empty())
	assert.Equal(t, 1, q.Len())

	q.Enqueue(2)
	assert.Equal(t, 2, q.Len())
}

func TestQueue_Dequeue(t *testing.T) {
	q := New[int]()

	// 测试从空队列出队
	value, ok := q.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, 0, value)

	// 测试正常出队
	q.Enqueue(1).Enqueue(2).Enqueue(3)

	value, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, value)
	assert.Equal(t, 2, q.Len())

	value, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, value)
	assert.Equal(t, 1, q.Len())

	value, ok = q.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, value)
	assert.Equal(t, 0, q.Len())

	// 再次尝试从空队列出队
	value, ok = q.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestQueue_Front(t *testing.T) {
	q := New[string]()

	// 测试空队列的Front
	value, ok := q.Front()
	assert.False(t, ok)
	assert.Equal(t, "", value)

	// 测试非空队列的Front
	q.Enqueue("first").Enqueue("second")

	value, ok = q.Front()
	assert.True(t, ok)
	assert.Equal(t, "first", value)

	// 确保Front不会移除元素
	assert.Equal(t, 2, q.Len())
}

func TestQueue_Clear(t *testing.T) {
	q := New[int]()
	q.Enqueue(1).Enqueue(2).Enqueue(3)
	assert.Equal(t, 3, q.Len())

	q.Clear()
	assert.True(t, q.Empty())
	assert.Equal(t, 0, q.Len())
}
