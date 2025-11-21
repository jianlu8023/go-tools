package _map

import (
	"sync"

	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
)

// RWMap 是一个基于读写锁的并发安全Map实现
type RWMap[K comparable, V any] struct {
	mutex sync.RWMutex
	m     map[K]V
}

var _ concurrent.Map[int, int] = (*RWMap[int, int])(nil)

// NewRWMap 创建一个新的RWMap实例
func NewRWMap[K comparable, V any]() concurrent.Map[K, V] {
	return &RWMap[K, V]{
		m: make(map[K]V),
	}
}

// Put 向map中添加键值对
func (rw *RWMap[K, V]) Put(key K, value V) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.m[key] = value
}

// Get 从map中获取指定键的值
func (rw *RWMap[K, V]) Get(key K) (V, bool) {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	value, ok := rw.m[key]
	return value, ok
}

// Del 从map中删除指定键
func (rw *RWMap[K, V]) Del(key K) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	delete(rw.m, key)
}

// Len 返回map的长度
func (rw *RWMap[K, V]) Len() int {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	return len(rw.m)
}

// Clear 清空map
func (rw *RWMap[K, V]) Clear() {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.m = make(map[K]V)
}

// Empty 判断map是否为空
func (rw *RWMap[K, V]) Empty() bool {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	return len(rw.m) == 0
}

// HasKey 判断map中是否包含指定键
func (rw *RWMap[K, V]) HasKey(key K) bool {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	_, ok := rw.m[key]
	return ok
}

// Keys 返回map中所有的键
func (rw *RWMap[K, V]) Keys() []K {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	keys := make([]K, 0, len(rw.m))
	if len(rw.m) == 0 {
		return keys
	}
	for k := range rw.m {
		keys = append(keys, k)
	}
	return keys
}

// Values 返回map中所有的值
func (rw *RWMap[K, V]) Values() []V {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	values := make([]V, 0, len(rw.m))
	if len(rw.m) == 0 {
		return values
	}
	for _, v := range rw.m {
		values = append(values, v)
	}
	return values
}

// Iterator 返回map的迭代器
func (rw *RWMap[K, V]) Iterator() concurrent.Iterator[concurrent.Entry[K, V]] {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()

	entries := make([]concurrent.Entry[K, V], 0, len(rw.m))
	for k, v := range rw.m {
		entries = append(entries, concurrent.Entry[K, V]{Key: k, Value: v})
	}

	return &entryIterator[K, V]{
		entries: entries,
		index:   -1, // 初始化为-1，第一次调用Next时会移动到0
	}
}

// Map 是一个基于sync.Map的并发安全Map实现
// 适用于只读、只写或读多写少的场景
type Map[K comparable, V any] struct {
	m   sync.Map
	len int
}

var _ concurrent.Map[int, int] = (*Map[int, int])(nil)

// NewMap 创建一个新的Map实例
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:   sync.Map{},
		len: 0,
	}
}

// Put 向map中添加键值对
func (wm *Map[K, V]) Put(key K, val V) {
	_, loaded := wm.m.LoadOrStore(key, val)
	if !loaded {
		wm.len++
	}
}

// Get 从map中获取指定键的值
func (wm *Map[K, V]) Get(key K) (V, bool) {
	result, ok := wm.m.Load(key)
	return result.(V), ok
}

// Del 从map中删除指定键
func (wm *Map[K, V]) Del(key K) {
	// 先检查键是否存在再删除
	_, ok := wm.m.Load(key)
	if ok {
		wm.m.Delete(key)
		wm.len--
	}
}

// Len 返回map的长度
func (wm *Map[K, V]) Len() int {
	return wm.len
}

// Clear 清空map
func (wm *Map[K, V]) Clear() {
	wm.m = sync.Map{}
	wm.len = 0
}

// Empty 判断map是否为空
func (wm *Map[K, V]) Empty() bool {
	return wm.len == 0
}

// HasKey 判断map中是否包含指定键
func (wm *Map[K, V]) HasKey(key K) bool {
	_, ok := wm.m.Load(key)
	return ok
}

// Keys 返回map中所有的键
func (wm *Map[K, V]) Keys() []K {
	result := make([]K, 0, wm.len)
	wm.m.Range(func(key, value any) bool {
		result = append(result, key.(K))
		return true
	})
	return result
}

// Values 返回map中所有的值
func (wm *Map[K, V]) Values() []V {
	result := make([]V, 0, wm.len)
	wm.m.Range(func(key, value any) bool {
		result = append(result, value.(V))
		return true
	})
	return result
}

// Iterator 返回map的迭代器
func (wm *Map[K, V]) Iterator() concurrent.Iterator[concurrent.Entry[K, V]] {
	entries := make([]concurrent.Entry[K, V], 0, wm.len)
	wm.m.Range(func(key, value any) bool {
		entries = append(entries, concurrent.Entry[K, V]{Key: key.(K), Value: value.(V)})
		return true
	})

	return &mapIterator[K, V]{
		entries: entries,
		index:   -1, // 初始化为-1，第一次调用Next时会移动到0
	}
}
