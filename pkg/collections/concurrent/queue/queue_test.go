package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRWQueue_Enqueue(t *testing.T) {
	queue := NewRWQueue[int]()
	assert.True(t, queue.Empty())
	assert.Equal(t, 0, queue.Len())

	queue.Enqueue(1)
	assert.False(t, queue.Empty())
	assert.Equal(t, 1, queue.Len())

	queue.Enqueue(2)
	assert.Equal(t, 2, queue.Len())
}

func TestRWQueue_Dequeue(t *testing.T) {
	queue := NewRWQueue[int]()

	// 测试从空队列出队
	value, ok := queue.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, 0, value)

	// 测试正常出队
	queue.Enqueue(1).Enqueue(2).Enqueue(3)

	value, ok = queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, value)
	assert.Equal(t, 2, queue.Len())

	value, ok = queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 2, value)
	assert.Equal(t, 1, queue.Len())

	value, ok = queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 3, value)
	assert.Equal(t, 0, queue.Len())

	// 再次尝试从空队列出队
	value, ok = queue.Dequeue()
	assert.False(t, ok)
	assert.Equal(t, 0, value)
}

func TestRWQueue_Front(t *testing.T) {
	queue := NewRWQueue[string]()

	// 测试空队列的Front
	value, ok := queue.Front()
	assert.False(t, ok)
	assert.Equal(t, "", value)

	// 测试非空队列的Front
	queue.Enqueue("first").Enqueue("second")

	value, ok = queue.Front()
	assert.True(t, ok)
	assert.Equal(t, "first", value)

	// 确保Front不会移除元素
	assert.Equal(t, 2, queue.Len())
}

func TestRWQueue_Clear(t *testing.T) {
	queue := NewRWQueue[int]()
	queue.Enqueue(1).Enqueue(2).Enqueue(3)
	assert.Equal(t, 3, queue.Len())

	queue.Clear()
	assert.True(t, queue.Empty())
	assert.Equal(t, 0, queue.Len())
}

func TestRWQueue_Iterator(t *testing.T) {
	queue := NewRWQueue[int]()
	queue.Enqueue(1).Enqueue(2).Enqueue(3)

	iterator := queue.Iterator()

	// 测试迭代器
	values := make([]int, 0)
	for iterator.HasNext() {
		values = append(values, iterator.Value())
	}

	assert.Equal(t, []int{1, 2, 3}, values)

	// 测试关闭迭代器
	err := iterator.Close()
	assert.NoError(t, err)
}

func TestRWQueue_Concurrent(t *testing.T) {
	queue := NewRWQueue[int]()

	// 并发测试不是单元测试的重点，这里只是简单验证
	// 实际并发安全性需要通过go test -race来验证

	queue.Enqueue(1)
	assert.Equal(t, 1, queue.Len())

	value, ok := queue.Dequeue()
	assert.True(t, ok)
	assert.Equal(t, 1, value)
}
