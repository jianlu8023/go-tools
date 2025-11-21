package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetOperations(t *testing.T) {
	// 创建一个新的set
	s := New[int]()

	// 测试Add和Len
	s.Add(1, 2, 3, 4, 5)
	assert.Equal(t, 5, s.Len(), "添加5个元素后长度应该为5")

	// 测试Has
	assert.True(t, s.Has(3), "set应该包含元素3")
	assert.False(t, s.Has(6), "set不应该包含元素6")

	// 测试Members
	members := s.Members()
	assert.Equal(t, 5, len(members), "成员数量应该为5")

	// 测试Remove
	s.Remove(3, 4)
	assert.Equal(t, 3, s.Len(), "删除2个元素后长度应该为3")
	assert.False(t, s.Has(3), "删除后set不应该包含元素3")

	// 测试Update
	s.Update(1, 10)
	assert.False(t, s.Has(1), "更新后set不应该包含元素1")
	assert.True(t, s.Has(10), "更新后set应该包含元素10")

	// 测试Pop
	lengthBeforePop := s.Len()
	s.Pop()
	lengthAfterPop := s.Len()

	assert.Equal(t, lengthBeforePop-1, lengthAfterPop, "Pop后长度应该减少1")

	// 测试One
	lengthBeforeOne := s.Len()
	s.One()
	lengthAfterOne := s.Len()

	assert.Equal(t, lengthBeforeOne, lengthAfterOne, "One后长度应该保持不变")

	// 测试Clear
	s = s.Clear()
	assert.Equal(t, 0, s.Len(), "Clear后长度应该为0")
}

func TestSetUnion(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()

	s1.Add(1, 2, 3)
	s2.Add(3, 4, 5)

	// 测试方法形式的Union
	s1.Union(s2)

	assert.Equal(t, 5, s1.Len(), "Union后长度应该为5")

	expected := []int{1, 2, 3, 4, 5}
	for _, v := range expected {
		assert.True(t, s1.Has(v), "Union后的set应该包含元素%d", v)
	}
}

func TestSetDiff(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()

	s1.Add(1, 2, 3, 4, 5)
	s2.Add(3, 4, 5, 6, 7)

	// 测试方法形式的Diff
	s1.Diff(s2)

	assert.Equal(t, 2, s1.Len(), "Diff后长度应该为2")
	assert.True(t, s1.Has(1), "Diff后的set应该包含元素1")
	assert.True(t, s1.Has(2), "Diff后的set应该包含元素2")
	assert.False(t, s1.Has(3), "Diff后的set不应该包含元素3")
	assert.False(t, s1.Has(4), "Diff后的set不应该包含元素4")
	assert.False(t, s1.Has(5), "Diff后的set不应该包含元素5")
}

func TestSetIntersection(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()

	s1.Add(1, 2, 3, 4, 5)
	s2.Add(3, 4, 5, 6, 7)

	// 测试方法形式的Intersection
	s1.Intersection(s2)

	assert.Equal(t, 3, s1.Len(), "Intersection后长度应该为3")
	assert.True(t, s1.Has(3), "Intersection后的set应该包含元素3")
	assert.True(t, s1.Has(4), "Intersection后的set应该包含元素4")
	assert.True(t, s1.Has(5), "Intersection后的set应该包含元素5")
	assert.False(t, s1.Has(1), "Intersection后的set不应该包含元素1")
	assert.False(t, s1.Has(2), "Intersection后的set不应该包含元素2")
}

func TestSetFunctions(t *testing.T) {
	s1 := New[int]()
	s2 := New[int]()

	s1.Add(1, 2, 3, 4, 5)
	s2.Add(3, 4, 5, 6, 7)

	// 测试函数形式的Diff
	diff := Diff(s1, s2)
	assert.Equal(t, 2, diff.Len(), "Diff函数后长度应该为2")

	// 重新创建s1因为Diff修改了原set
	s1 = New[int]()
	s1.Add(1, 2, 3, 4, 5)

	// 测试函数形式的Union
	union := Union(s1, s2)
	assert.Equal(t, 7, union.Len(), "Union函数后长度应该为7")

	// 重新创建s1因为Union修改了原set
	s1 = New[int]()
	s1.Add(1, 2, 3, 4, 5)

	// 测试函数形式的Intersection
	intersection := Intersection(s1, s2)
	assert.Equal(t, 3, intersection.Len(), "Intersection函数后长度应该为3")
}

func TestSetLoop(t *testing.T) {
	s := New[int]()
	s.Add(1, 2, 3, 4, 5)

	sum := 0
	s.Loop(func(element int) {
		sum += element
	})

	assert.Equal(t, 15, sum, "循环累加结果应该为15")
}

func TestSetIterator(t *testing.T) {
	// 创建一个Set并添加一些数据
	s := New[int]()
	s.Add(1, 2, 3, 4, 5)

	// 使用迭代器遍历
	iter := s.Iterator()
	defer iter.Close()

	count := 0
	values := make(map[int]bool)

	// 遍历所有元素
	for iter.HasNext() {

		value := iter.Value()
		values[value] = true
		count++

	}

	// 验证结果
	assert.Equal(t, 5, count)
	assert.True(t, values[1])
	assert.True(t, values[2])
	assert.True(t, values[3])
	assert.True(t, values[4])
	assert.True(t, values[5])
}

func TestEmptySetIterator(t *testing.T) {
	// 测试空Set的迭代器
	s := New[int]()
	iter := s.Iterator()
	defer iter.Close()

	// 对于空Set，HasNext应该返回false
	assert.False(t, iter.HasNext())

	// Next应该返回false

}

func TestSetIteratorClose(t *testing.T) {
	// 测试关闭迭代器
	s := New[string]()
	s.Add("a", "b", "c")

	iter := s.Iterator()

	// 使用一次迭代器
	assert.True(t, iter.HasNext())
	value := iter.Value()
	assert.Equal(t, "a", value)

	// 关闭迭代器
	err := iter.Close()
	assert.NoError(t, err)

	// 关闭后再次调用HasNext应该返回false
	assert.False(t, iter.HasNext())

	// 关闭后再次调用Next应该返回false

}

func TestSetIteratorValue(t *testing.T) {
	// 测试Value方法的边界情况
	s := New[int]()
	s.Add(42)

	iter := s.Iterator()
	defer iter.Close()

	// 在调用Next之前，Value应该返回零值
	var zero int
	assert.Equal(t, zero, iter.Value())

	// 调用Next后，Value应该返回正确的值
	assert.True(t, iter.HasNext())
	assert.Equal(t, 42, iter.Value())

	// 超出范围后，Value应该返回零值
	assert.False(t, iter.HasNext())
	assert.Equal(t, zero, iter.Value())
}
