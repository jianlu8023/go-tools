package _map

import (
	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
)

// entryIterator 实现Iterator接口，用于遍历Map的键值对
type entryIterator[K comparable, V any] struct {
	entries []concurrent.Entry[K, V]
	index   int
	closed  bool
}

// mapIterator 实现Iterator接口，用于遍历sync.Map的键值对
type mapIterator[K comparable, V any] struct {
	entries []concurrent.Entry[K, V]
	index   int
	closed  bool
}
