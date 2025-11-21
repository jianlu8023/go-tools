package concurrent

// Map 定义并发安全的Map接口
type Map[K comparable, V any] interface {
	// Put 向 map 中添加元素
	Put(K, V)
	// Get 从 map 中取元素
	Get(K) (V, bool)
	// Del 将 map 中元素删除
	Del(K)
	// Len 返回 map 长度
	Len() int
	// Clear 清空 map
	Clear()
	// Empty 判断 map 是否是空
	Empty() bool
	// HasKey 判断 map 中是否有 key
	HasKey(K) bool
	// Keys 返回 map 所有 key
	Keys() []K
	// Values 返回 map 所有 value
	Values() []V
	// Iterator 返回 map 的迭代器
	Iterator() Iterator[Entry[K, V]]
}

// Set 定义并发安全的Set接口
type Set[T comparable] interface {
	// Add 添加元素到 set
	Add(elements ...T) Set[T]
	// Remove 从 set 中删除元素
	Remove(elements ...T) Set[T]
	// Update 将 set 中 旧数据替换为新数据
	Update(old, new T) Set[T]
	// Has 判断 set 中是否有该元素
	Has(element T) bool
	// Len 返回 set 长度
	Len() int
	// Members 返回全部 set 元素
	Members() []T
	// Pop 随机返回一个 删除该元素
	Pop() T
	// One 随机返回一个 删除该元素
	One() T
	// Clear 将 set 清空
	Clear() Set[T]
	// Union 和其他 set 做并集
	Union(other Set[T]) Set[T]
	// Diff 和其他 set 做差集
	Diff(other Set[T]) Set[T]
	// Intersection 和其他 set 做 交集
	Intersection(other Set[T]) Set[T]
	// Loop 遍历 set 中每个元素
	Loop(fn func(element T))
	// Iterator 返回 set 的迭代器
	Iterator() Iterator[T]
}
