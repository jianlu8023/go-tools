package queue_test

import (
	"fmt"

	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent/queue"
)

// ExampleRWQueue 示例：如何使用并发安全的队列
func ExampleRWQueue() {
	// 创建一个新的并发安全队列
	q := queue.NewRWQueue[string]()

	// 向队列中添加元素
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")

	// 查看队列长度
	fmt.Printf("队列长度: %d\n", q.Len())

	// 查看队列头部元素
	if front, ok := q.Front(); ok {
		fmt.Printf("队列头部元素: %s\n", front)
	}

	// 从队列中取出元素
	for !q.Empty() {
		if value, ok := q.Dequeue(); ok {
			fmt.Printf("取出元素: %s\n", value)
		}
	}

	// Output:
	// 队列长度: 3
	// 队列头部元素: first
	// 取出元素: first
	// 取出元素: second
	// 取出元素: third
}
