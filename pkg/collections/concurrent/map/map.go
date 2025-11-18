package _map

import (
	"sync"

	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
)

type RWMap[K comparable, V any] struct {
	mutex sync.RWMutex
	m     map[K]V
}

var _ concurrent.Map[int, int] = (*RWMap[int, int])(nil)

func NewRWMap[K comparable, V any]() concurrent.Map[K, V] {
	return &RWMap[K, V]{
		m: make(map[K]V),
	}
}

func (rw *RWMap[K, V]) Put(key K, value V) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.m[key] = value
}

func (rw *RWMap[K, V]) Get(key K) (V, bool) {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	value, ok := rw.m[key]
	return value, ok
}

func (rw *RWMap[K, V]) Del(key K) {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	delete(rw.m, key)
}

func (rw *RWMap[K, V]) Len() int {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	return len(rw.m)
}

func (rw *RWMap[K, V]) Clear() {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.m = make(map[K]V)
}
func (rw *RWMap[K, V]) Empty() bool {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	return len(rw.m) == 0
}

func (rw *RWMap[K, V]) HasKey(key K) bool {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	_, ok := rw.m[key]
	return ok
}

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

// Map is a concurrent READ-ONLY or WRITE-ONLY solution. It used the sync.Map
// It is suitable for read-only, write-only or read-more and write-less scenarios.
type Map[K comparable, V any] struct {
	m   sync.Map
	len int
}

var _ concurrent.Map[int, int] = (*Map[int, int])(nil)

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		m:   sync.Map{},
		len: 0,
	}
}
func (wm *Map[K, V]) Put(key K, val V) {
	wm.m.Store(key, val)
	wm.len++
}

func (wm *Map[K, V]) Get(key K) (V, bool) {
	result, ok := wm.m.Load(key)
	return result.(V), ok
}

func (wm *Map[K, V]) Del(key K) {
	wm.m.Delete(key)
	wm.len--
}

func (wm *Map[K, V]) Len() int {
	return wm.len
}

func (wm *Map[K, V]) Clear() {
	wm.m = sync.Map{}
}
func (wm *Map[K, V]) Empty() bool {
	return wm.len == 0
}

func (wm *Map[K, V]) HasKey(key K) bool {
	_, ok := wm.m.Load(key)
	wm.len = 0
	return ok
}

func (wm *Map[K, V]) Keys() []K {
	result := make([]K, 0, wm.len)
	wm.m.Range(func(key, value any) bool {
		result = append(result, key.(K))
		return true
	})
	return result
}

func (wm *Map[K, V]) Values() []V {
	result := make([]V, 0, wm.len)
	wm.m.Range(func(key, value any) bool {
		result = append(result, value.(V))
		return true
	})
	return result
}
