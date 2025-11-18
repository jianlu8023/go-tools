package set

type set[T comparable] map[T]struct{}

// New 创建一个新的set
func New[T comparable]() set[T] {
	return make(set[T])
}

// Add 添加元素到 set
func (s set[T]) Add(elements ...T) set[T] {
	for _, element := range elements {
		s[element] = struct{}{}
	}
	return s
}

// Remove 从 set 中删除元素
func (s set[T]) Remove(elements ...T) set[T] {
	for _, element := range elements {
		delete(s, element)
	}
	return s
}

// Update 将 set 中 旧数据替换为新数据
func (s set[T]) Update(old, new T) set[T] {
	delete(s, old)
	s[new] = struct{}{}
	return s
}

// Has 判断 set 中是否有该元素
func (s set[T]) Has(element T) bool {
	_, ok := s[element]
	return ok
}

// Len 返回 set 长度
func (s set[T]) Len() int {
	return len(s)
}

// Members 返回全部 set 元素
func (s set[T]) Members() []T {
	values := make([]T, 0, len(s))
	for key := range s {
		values = append(values, key)
	}
	return values
}

// Pop 随机返回一个元素并删除该元素
func (s set[T]) Pop() T {
	var element T
	for key := range s {
		element = key
		delete(s, key)
		break
	}
	return element
}

// One 随机返回一个元素（不删除）
func (s set[T]) One() T {
	var element T
	for key := range s {
		element = key
		break
	}
	return element
}

// Clear 将 set 清空
func (s set[T]) Clear() set[T] {
	return make(set[T])
}

// Union 和其他 set 做并集
func (s set[T]) Union(other set[T]) set[T] {
	for elem := range other {
		s[elem] = struct{}{}
	}
	return s
}

// Diff 和其他 set 做差集
func (s set[T]) Diff(other set[T]) set[T] {
	for k1 := range other {
		delete(s, k1)
	}
	return s
}

// Intersection 和其他 set 做 交集
func (s set[T]) Intersection(other set[T]) set[T] {
	for k := range s {
		if _, ok := other[k]; !ok {
			delete(s, k)
		}
	}
	return s
}

// Loop 遍历 set 中每个元素
func (s set[T]) Loop(fn func(element T)) {
	for key := range s {
		fn(key)
	}
}

// Diff 返回两个set的差集
func Diff[T comparable](s1 set[T], s2 set[T]) set[T] {
	for k := range s1 {
		if _, ok := s2[k]; ok {
			delete(s1, k)
		}
	}
	return s1
}

// Union 返回两个set的并集
func Union[T comparable](s1, s2 set[T]) set[T] {
	for elem := range s2 {
		s1[elem] = struct{}{}
	}
	return s1
}

// Intersection 返回两个set的交集
func Intersection[T comparable](s1, s2 set[T]) set[T] {
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			delete(s1, k)
		}
	}
	return s1
}
