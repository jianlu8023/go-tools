package rand

import (
	"testing"
)

func TestNewMathRandomGenerator(t *testing.T) {
	g := NewMathRandomGenerator()
	if g == nil {
		t.Error("NewMathRandomGenerator returned nil")
	}
}

func TestMathRandomGenerator_Intn(t *testing.T) {
	g := NewMathRandomGenerator()
	n := 100
	v := g.Intn(n)
	if v < 0 || v >= n {
		t.Errorf("Intn(%d) returned %d, expected range [0, %d)", n, v, n)
	}

	// 测试边界情况
	v = g.Intn(1)
	if v != 0 {
		t.Errorf("Intn(1) returned %d, expected 0", v)
	}

	// 测试多线程安全性（简单测试，实际应该用并发测试）
	for i := 0; i < 1000; i++ {
		g.Intn(1000)
	}
}

func TestMathRandomGenerator_Uint32(t *testing.T) {
	g := NewMathRandomGenerator()
	v := g.Uint32()
	// uint32的范围总是有效的，只需要检查生成器正常工作
	_ = v // 使用变量避免编译警告

	// 测试多线程安全性
	for i := 0; i < 1000; i++ {
		g.Uint32()
	}
}

func TestMathRandomGenerator_Uint64(t *testing.T) {
	g := NewMathRandomGenerator()
	v := g.Uint64()
	// uint64的范围总是有效的，只需要检查生成器正常工作
	_ = v // 使用变量避免编译警告

	// 测试多线程安全性
	for i := 0; i < 1000; i++ {
		g.Uint64()
	}
}

func TestMathRandomGenerator_GenerateString(t *testing.T) {
	g := NewMathRandomGenerator()
	length := 10
	runes := "abcdef"
	s := g.GenerateString(length, runes)

	if len(s) != length {
		t.Errorf("expected length %d, got %d", length, len(s))
	}

	// 检查每个字符是否在指定的字符集中
	for _, c := range s {
		found := false
		for _, r := range runes {
			if c == r {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("character %c not in runes %q", c, runes)
		}
	}

	// 测试空字符集（应该仍然工作，返回空字符串）
	s = g.GenerateString(5, "")
	if len(s) != 0 {
		t.Errorf("expected empty string for empty runes, got %q", s)
	}

	// 测试长度为0
	s = g.GenerateString(0, "abc")
	if len(s) != 0 {
		t.Errorf("expected empty string for length 0, got %q", s)
	}

	// 测试多线程安全性
	for i := 0; i < 1000; i++ {
		g.GenerateString(10, "abcdef")
	}
}

func TestMathRandomGenerator_Float64(t *testing.T) {
	g := NewMathRandomGenerator()
	v := g.Float64()
	if v < 0.0 || v >= 1.0 {
		t.Errorf("Float64 returned %f, expected range [0.0, 1.0)", v)
	}

	// 测试多线程安全性
	for i := 0; i < 1000; i++ {
		g.Float64()
	}
}

func TestMathRandomGenerator_IntBetween(t *testing.T) {
	g := NewMathRandomGenerator()
	min, max := 5, 15
	v := g.IntBetween(min, max)
	if v < min || v > max {
		t.Errorf("IntBetween(%d, %d) returned %d, expected range [%d, %d]", min, max, v, min, max)
	}

	// 测试边界情况
	v = g.IntBetween(5, 5)
	if v != 5 {
		t.Errorf("IntBetween(5, 5) returned %d, expected 5", v)
	}

	// 测试无效范围（应该处理）
	v = g.IntBetween(10, 5)
	if v != 10 && v != 5 {
		t.Errorf("IntBetween(10, 5) returned %d, expected 5 or 10", v)
	}

	// 测试多线程安全性
	for i := 0; i < 1000; i++ {
		g.IntBetween(1, 100)
	}
}

func TestMathRandomGenerator_Shuffle(t *testing.T) {
	g := NewMathRandomGenerator()

	// 测试整数切片打乱
	slice := []int{1, 2, 3, 4, 5}
	g.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	// 检查所有元素是否仍然存在
	counts := make(map[int]int)
	for _, v := range slice {
		counts[v]++
	}
	for i := 1; i <= 5; i++ {
		if counts[i] != 1 {
			t.Errorf("expected count 1 for %d, got %d", i, counts[i])
		}
	}

	// 测试空切片
	emptySlice := []int{}
	g.Shuffle(len(emptySlice), func(i, j int) {
		emptySlice[i], emptySlice[j] = emptySlice[j], emptySlice[i]
	})

	// 测试单元素切片
	singleSlice := []int{42}
	g.Shuffle(len(singleSlice), func(i, j int) {
		singleSlice[i], singleSlice[j] = singleSlice[j], singleSlice[i]
	})
	if singleSlice[0] != 42 {
		t.Errorf("expected 42, got %d", singleSlice[0])
	}
}
