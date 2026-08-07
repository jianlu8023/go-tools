package set

import (
	"sync"

	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
)

// RWSet 是一个基于读写锁的并发安全Set实现
type RWSet[T comparable] struct {
	mutex sync.RWMutex
	m     map[T]struct{}
}

var _ concurrent.Set[int] = (*RWSet[int])(nil)

// NewRWSet 创建一个新的RWSet实例
func NewRWSet[T comparable]() concurrent.Set[T] {
	return &RWSet[T]{
		m: make(map[T]struct{}),
	}
}

// Add 向set中添加元素
func (rw *RWSet[T]) Add(elements ...T) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()

	for _, element := range elements {
		rw.m[element] = struct{}{}
	}
	return rw
}

// Remove 从set中删除元素
func (rw *RWSet[T]) Remove(elements ...T) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	for _, element := range elements {
		delete(rw.m, element)
	}
	return rw
}

// Update 将set中的旧数据替换为新数据
func (rw *RWSet[T]) Update(old, new T) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	delete(rw.m, old)
	rw.m[new] = struct{}{}
	return rw
}

// Has 判断set中是否包含指定元素
func (rw *RWSet[T]) Has(element T) bool {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	_, ok := rw.m[element]
	return ok
}

// Len 返回set的长度
func (rw *RWSet[T]) Len() int {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	return len(rw.m)
}

// Members 返回set中所有的元素
func (rw *RWSet[T]) Members() []T {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	result := make([]T, 0, len(rw.m))
	for elem := range rw.m {
		result = append(result, elem)
	}
	return result
}

// Pop 随机返回并删除一个元素
func (rw *RWSet[T]) Pop() T {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	var elem T
	for k := range rw.m {
		elem = k
		delete(rw.m, k)
		break // 只删除一个元素
	}
	return elem
}

// One 随机返回一个元素（不删除）
func (rw *RWSet[T]) One() T {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	var elem T
	for k := range rw.m {
		elem = k
		break
	}
	return elem
}

// Clear 清空set
func (rw *RWSet[T]) Clear() concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.m = make(map[T]struct{})
	return rw
}

// Union 和其他set做并集
// 注意：为避免在持有自身写锁的同时回调 other 的方法（可能反向获取 other 的锁）造成死锁，
// 这里先通过 other.Members() 在无锁状态下拷贝一份快照，再在自身锁内完成合并。
func (rw *RWSet[T]) Union(other concurrent.Set[T]) concurrent.Set[T] {
	others := other.Members()
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	for _, elem := range others {
		rw.m[elem] = struct{}{}
	}
	return rw
}

// Diff 和其他set做差集
// 同 Union，先拷贝 other 快照，再在自身锁内完成删除，避免死锁。
func (rw *RWSet[T]) Diff(other concurrent.Set[T]) concurrent.Set[T] {
	others := other.Members()
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	for _, elem := range others {
		delete(rw.m, elem)
	}
	return rw
}

// Intersection 和其他set做交集
// 同 Union，先拷贝 other 快照，再在自身锁内完成判断，避免死锁。
func (rw *RWSet[T]) Intersection(other concurrent.Set[T]) concurrent.Set[T] {
	others := other.Members()
	otherSet := make(map[T]struct{}, len(others))
	for _, elem := range others {
		otherSet[elem] = struct{}{}
	}
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	for k := range rw.m {
		if _, ok := otherSet[k]; !ok {
			delete(rw.m, k)
		}
	}
	return rw
}

// Loop 遍历set中的每个元素
func (rw *RWSet[T]) Loop(fn func(element T)) {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	for elem := range rw.m {
		fn(elem)
	}
}

// Iterator 返回set的迭代器
func (rw *RWSet[T]) Iterator() concurrent.Iterator[T] {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()

	elements := make([]T, 0, len(rw.m))
	for elem := range rw.m {
		elements = append(elements, elem)
	}

	return &setIterator[T]{
		elements: elements,
		index:    -1, // 初始化为-1，第一次调用Next时会移动到0
	}
}
