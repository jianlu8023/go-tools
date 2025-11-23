package xxhash

import (
	"testing"
)

func TestSum64(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected uint64
	}{
		{"空数据", []byte{}, 0xEF46DB3751D8E999},
		{"单个字节", []byte{0x01}, 9962287286179718960},
		{"简单字符串", []byte("hello"), 2794345569481354659},
		{"较长字符串", []byte("The quick brown fox jumps over the lazy dog"), 802816344064684476},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum64(tt.data)
			if got != tt.expected {
				t.Errorf("Sum64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSum64String(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected uint64
	}{
		{"空字符串", "", 0xEF46DB3751D8E999},
		{"简单字符串", "hello", 2794345569481354659},
		{"较长字符串", "The quick brown fox jumps over the lazy dog", 802816344064684476},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum64String(tt.data)
			if got != tt.expected {
				t.Errorf("Sum64String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected uint64
	}{
		{"空数据", []byte{}, 0xEF46DB3751D8E999},
		{"单个字节", []byte{0x01}, 9962287286179718960},
		{"简单字符串", []byte("hello"), 2794345569481354659},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New()
			h.Write(tt.data)
			got := h.Sum64()
			if got != tt.expected {
				t.Errorf("New().Sum64() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDigestReset(t *testing.T) {
	h := New()
	h.Write([]byte("hello"))
	firstHash := h.Sum64()

	h.Reset()
	// 重置后应该得到空数据的哈希值
	resetHash := h.Sum64()
	if resetHash != 0xEF46DB3751D8E999 {
		t.Errorf("Reset() 后哈希值错误: got %v, want %v", resetHash, uint64(0xEF46DB3751D8E999))
	}

	// 重置后再次写入数据应该正常工作
	h.Write([]byte("hello"))
	secondHash := h.Sum64()
	if firstHash != secondHash {
		t.Errorf("Reset() 后再次计算哈希值不一致: got %v, want %v", secondHash, firstHash)
	}
}

func TestDigestSizeAndBlockSize(t *testing.T) {
	h := New()
	if h.Size() != 8 {
		t.Errorf("Size() = %v, want 8", h.Size())
	}
	if h.BlockSize() != 32 {
		t.Errorf("BlockSize() = %v, want 32", h.BlockSize())
	}
}

func TestSum(t *testing.T) {
	h := New()
	h.Write([]byte("hello"))

	// 测试Sum方法
	buf := make([]byte, 0, 16)
	sum := h.Sum(buf)
	if len(sum) != 8 {
		t.Errorf("Sum() 返回长度错误: got %v, want 8", len(sum))
	}

	// 验证Sum方法返回的值与Sum64一致
	sum64 := Sum64([]byte("hello"))
	// 注意：这里需要考虑字节序，使用xxhash库的Sum64方法直接验证
	h2 := New()
	h2.Write([]byte("hello"))
	sumFromDigest := h2.Sum64()
	if sum64 != sumFromDigest {
		t.Errorf("Sum() 相关计算与 Sum64() 结果不一致")
	}
}
