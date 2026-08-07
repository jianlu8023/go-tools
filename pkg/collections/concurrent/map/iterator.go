package _map

import (
	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
)

// HasNext 移动到下一个元素
func (it *entryIterator[K, V]) HasNext() bool {
	if it.closed || it.index >= len(it.entries) {
		return false
	}
	it.index++
	return it.index < len(it.entries)
}

// Value 返回当前元素的值
func (it *entryIterator[K, V]) Value() concurrent.Entry[K, V] {
	if it.closed || it.index < 0 || it.index >= len(it.entries) {
		var empty concurrent.Entry[K, V]
		return empty
	}
	return it.entries[it.index]
}

// Close 关闭迭代器
func (it *entryIterator[K, V]) Close() error {
	it.closed = true
	it.entries = nil
	return nil
}

// HasNext 移动到下一个元素
func (it *mapIterator[K, V]) HasNext() bool {
	if it.closed || it.index >= len(it.entries) {
		return false
	}
	it.index++
	return it.index < len(it.entries)
}

// Value 返回当前元素的值
func (it *mapIterator[K, V]) Value() concurrent.Entry[K, V] {
	if it.closed || it.index < 0 || it.index >= len(it.entries) {
		var empty concurrent.Entry[K, V]
		return empty
	}
	return it.entries[it.index]
}

// Close 关闭迭代器
func (it *mapIterator[K, V]) Close() error {
	it.closed = true
	it.entries = nil
	return nil
}
