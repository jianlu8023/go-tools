package set

import (
	"testing"

	"github.com/jianlu8023/go-tools/v2/pkg/collections/concurrent"
	"github.com/stretchr/testify/assert"
)

func TestRWSet(t *testing.T) {
	// 创建一个新的RWSet
	rwSet := NewRWSet[int]()

	// 测试Add和Len
	rwSet.Add(1, 2, 3, 4, 5)
	assert.Equal(t, 5, rwSet.Len(), "添加5个元素后长度应该为5")

	// 测试Has
	assert.True(t, rwSet.Has(3), "set应该包含元素3")
	assert.False(t, rwSet.Has(6), "set不应该包含元素6")

	// 测试Members
	members := rwSet.Members()
	assert.Equal(t, 5, len(members), "成员数量应该为5")

	// 测试Remove
	rwSet.Remove(3, 4)
	assert.Equal(t, 3, rwSet.Len(), "删除2个元素后长度应该为3")
	assert.False(t, rwSet.Has(3), "删除后set不应该包含元素3")

	// 测试Update
	rwSet.Update(1, 10)
	assert.False(t, rwSet.Has(1), "更新后set不应该包含元素1")
	assert.True(t, rwSet.Has(10), "更新后set应该包含元素10")

	// 测试Pop
	lengthBeforePop := rwSet.Len()
	rwSet.Pop()
	lengthAfterPop := rwSet.Len()

	assert.Equal(t, lengthBeforePop-1, lengthAfterPop, "Pop后长度应该减少1")

	// 测试One
	lengthBeforeOne := rwSet.Len()
	rwSet.One()
	lengthAfterOne := rwSet.Len()

	assert.Equal(t, lengthBeforeOne, lengthAfterOne, "One后长度应该保持不变")

	// 测试Clear
	rwSet.Clear()
	assert.Equal(t, 0, rwSet.Len(), "Clear后长度应该为0")
}

func TestRWSetUnion(t *testing.T) {
	s1 := NewRWSet[int]()
	s2 := NewRWSet[int]()

	s1.Add(1, 2, 3)
	s2.Add(3, 4, 5)

	// 测试Union
	s1.Union(s2)

	assert.Equal(t, 5, s1.Len(), "Union后长度应该为5")

	expected := []int{1, 2, 3, 4, 5}
	for _, v := range expected {
		assert.True(t, s1.Has(v), "Union后的set应该包含元素%d", v)
	}
}

func TestRWSetDiff(t *testing.T) {
	s1 := NewRWSet[int]()
	s2 := NewRWSet[int]()

	s1.Add(1, 2, 3, 4, 5)
	s2.Add(3, 4, 5, 6, 7)

	// 创建另一个set用于Diff操作
	s3 := NewRWSet[int]()
	s3.Add(3, 4, 5, 6, 7)

	// 测试Diff
	s1.Diff(s3)

	assert.Equal(t, 2, s1.Len(), "Diff后长度应该为2")
	assert.True(t, s1.Has(1), "Diff后的set应该包含元素1")
	assert.True(t, s1.Has(2), "Diff后的set应该包含元素2")
	assert.False(t, s1.Has(3), "Diff后的set不应该包含元素3")
	assert.False(t, s1.Has(4), "Diff后的set不应该包含元素4")
	assert.False(t, s1.Has(5), "Diff后的set不应该包含元素5")
}

func TestRWSetIntersection(t *testing.T) {
	s1 := NewRWSet[int]()
	s2 := NewRWSet[int]()

	s1.Add(1, 2, 3, 4, 5)
	s2.Add(3, 4, 5, 6, 7)

	// 创建另一个set用于Intersection操作
	s3 := NewRWSet[int]()
	s3.Add(3, 4, 5, 6, 7)

	// 测试Intersection
	s1.Intersection(s3)

	assert.Equal(t, 3, s1.Len(), "Intersection后长度应该为3")
	assert.True(t, s1.Has(3), "Intersection后的set应该包含元素3")
	assert.True(t, s1.Has(4), "Intersection后的set应该包含元素4")
	assert.True(t, s1.Has(5), "Intersection后的set应该包含元素5")
	assert.False(t, s1.Has(1), "Intersection后的set不应该包含元素1")
	assert.False(t, s1.Has(2), "Intersection后的set不应该包含元素2")
}

func TestRWSetLoop(t *testing.T) {
	s := NewRWSet[int]()
	s.Add(1, 2, 3, 4, 5)

	sum := 0
	s.Loop(func(element int) {
		sum += element
	})

	assert.Equal(t, 15, sum, "循环累加结果应该为15")
}

func TestConcurrentSetInterface(t *testing.T) {
	// 测试RWSet是否实现了concurrent.Set接口
	var _ concurrent.Set[int] = NewRWSet[int]()
}
