package set

import (
	"sync"

	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
)

type RWSet[T comparable] struct {
	mutex sync.RWMutex
	m     map[T]struct{}
}

var _ concurrent.Set[int] = (*RWSet[int])(nil)

func NewRWSet[T comparable]() concurrent.Set[T] {
	return &RWSet[T]{
		m: make(map[T]struct{}),
	}
}

func (rw *RWSet[T]) Add(elements ...T) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()

	for _, element := range elements {
		rw.m[element] = struct{}{}
	}
	return rw
}

func (rw *RWSet[T]) Remove(elements ...T) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	for _, element := range elements {
		delete(rw.m, element)
	}
	return rw
}

func (rw *RWSet[T]) Update(old, new T) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	delete(rw.m, old)
	rw.m[new] = struct{}{}
	return rw
}

func (rw *RWSet[T]) Has(element T) bool {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	_, ok := rw.m[element]
	return ok
}

func (rw *RWSet[T]) Len() int {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	return len(rw.m)
}

func (rw *RWSet[T]) Members() []T {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	result := make([]T, 0, len(rw.m))
	for elem := range rw.m {
		result = append(result, elem)
	}
	return result
}

func (rw *RWSet[T]) Pop() T {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	var elem T
	for k := range rw.m {
		elem = k
		delete(rw.m, k)
	}
	return elem
}

func (rw *RWSet[T]) One() T {
	rw.mutex.RLock()
	defer rw.mutex.RUnlock()
	var elem T
	for k := range rw.m {
		elem = k
	}
	return elem
}

func (rw *RWSet[T]) Clear() concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.m = make(map[T]struct{})
	return rw
}

func (rw *RWSet[T]) Union(other concurrent.Set[T]) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	other.Loop(func(elem T) {
		rw.m[elem] = struct{}{}
	})
	return rw
}

func (rw *RWSet[T]) Diff(other concurrent.Set[T]) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	other.Loop(func(k1 T) {
		delete(rw.m, k1)
	})
	return rw
}

func (rw *RWSet[T]) Intersection(other concurrent.Set[T]) concurrent.Set[T] {
	rw.mutex.Lock()
	defer rw.mutex.Unlock()
	rw.Loop(func(k T) {
		if !other.Has(k) {
			delete(rw.m, k)
		}
	})
	return rw
}

func (rw *RWSet[T]) Loop(fn func(element T)) {
	for elem := range rw.m {
		fn(elem)
	}
}
