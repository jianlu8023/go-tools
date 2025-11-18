package _map

import (
	"testing"

	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
	"github.com/stretchr/testify/assert"
)

func TestRWMap(t *testing.T) {
	// 创建一个新的RWMap
	rwMap := NewRWMap[string, int]()

	// 测试Put和Get
	rwMap.Put("key1", 100)
	rwMap.Put("key2", 200)

	if value, ok := rwMap.Get("key1"); assert.True(t, ok, "应该能够获取到key1") {
		assert.Equal(t, 100, value, "key1的值应该为100")
	}

	if value, ok := rwMap.Get("key2"); assert.True(t, ok, "应该能够获取到key2") {
		assert.Equal(t, 200, value, "key2的值应该为200")
	}

	// 测试HasKey
	assert.True(t, rwMap.HasKey("key1"), "map应该包含key1")
	assert.False(t, rwMap.HasKey("key3"), "map不应该包含key3")

	// 测试Len
	assert.Equal(t, 2, rwMap.Len(), "map长度应该为2")

	// 测试Del
	rwMap.Del("key1")

	assert.False(t, rwMap.HasKey("key1"), "删除后map不应该包含key1")
	assert.Equal(t, 1, rwMap.Len(), "删除后map长度应该为1")

	// 测试Keys
	keys := rwMap.Keys()
	assert.Equal(t, 1, len(keys), "keys长度应该为1")
	assert.Equal(t, "key2", keys[0], "keys[0]应该为key2")

	// 测试Values
	values := rwMap.Values()
	assert.Equal(t, 1, len(values), "values长度应该为1")
	assert.Equal(t, 200, values[0], "values[0]应该为200")

	// 测试Empty
	assert.False(t, rwMap.Empty(), "map不应该为空")

	// 测试Clear
	rwMap.Clear()

	assert.True(t, rwMap.Empty(), "Clear后map应该为空")
	assert.Equal(t, 0, rwMap.Len(), "Clear后map长度应该为0")
}

func TestMap(t *testing.T) {
	// 创建一个新的Map
	m := NewMap[string, int]()

	// 测试Put和Get
	m.Put("key1", 100)
	m.Put("key2", 200)

	if value, ok := m.Get("key1"); assert.True(t, ok, "应该能够获取到key1") {
		assert.Equal(t, 100, value, "key1的值应该为100")
	}

	if value, ok := m.Get("key2"); assert.True(t, ok, "应该能够获取到key2") {
		assert.Equal(t, 200, value, "key2的值应该为200")
	}

	// 测试HasKey
	assert.True(t, m.HasKey("key1"), "map应该包含key1")
	assert.False(t, m.HasKey("key3"), "map不应该包含key3")

	// 测试Len
	assert.Equal(t, 2, m.Len(), "map长度应该为2")

	// 测试Del
	m.Del("key1")

	assert.False(t, m.HasKey("key1"), "删除后map不应该包含key1")
	assert.Equal(t, 1, m.Len(), "删除后map长度应该为1")

	// 测试Keys
	keys := m.Keys()
	assert.Equal(t, 1, len(keys), "keys长度应该为1")
	assert.Equal(t, "key2", keys[0], "keys[0]应该为key2")

	// 测试Values
	values := m.Values()
	assert.Equal(t, 1, len(values), "values长度应该为1")
	assert.Equal(t, 200, values[0], "values[0]应该为200")

	// 测试Empty
	assert.False(t, m.Empty(), "map不应该为空")

	// 测试Clear
	m.Clear()

	assert.True(t, m.Empty(), "Clear后map应该为空")
	assert.Equal(t, 0, m.Len(), "Clear后map长度应该为0")
}

func TestConcurrentMapInterface(t *testing.T) {
	// 测试RWMap是否实现了concurrent.Map接口
	var _ concurrent.Map[string, int] = NewRWMap[string, int]()

	// 测试Map是否实现了concurrent.Map接口
	var _ concurrent.Map[string, int] = NewMap[string, int]()
}
