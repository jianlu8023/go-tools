package sha256

import (
	"encoding/hex"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected []byte
	}{
		{
			"空数据",
			[]byte{},
			[]byte{0xe3, 0xb0, 0xc4, 0x42, 0x98, 0xfc, 0x1c, 0x14, 0x9a, 0xfb, 0xf4, 0xc8, 0x99, 0x6f, 0xb9, 0x24, 0x27, 0xae, 0x41, 0xe4, 0x64, 0x9b, 0x93, 0x4c, 0xa4, 0x95, 0x99, 0x1b, 0x78, 0x52, 0xb8, 0x55},
		},
		{
			"单个字节",
			[]byte{0x00},
			[]byte{0x6e, 0x34, 0x0b, 0x9c, 0xff, 0xb3, 0x7a, 0x98, 0x9c, 0xa5, 0x44, 0xe6, 0xbb, 0x78, 0x0a, 0x2c, 0x78, 0x90, 0x1d, 0x3f, 0xb3, 0x37, 0x38, 0x76, 0x85, 0x11, 0xa3, 0x06, 0x17, 0xaf, 0xa0, 0x1d},
		},
		{
			"简单字符串",
			[]byte("hello"),
			[]byte{0x2c, 0xf2, 0x4d, 0xba, 0x5f, 0xb0, 0xa3, 0x0e, 0x26, 0xe8, 0x3b, 0x2a, 0xc5, 0xb9, 0xe2, 0x9e, 0x1b, 0x16, 0x1e, 0x5c, 0x1f, 0xa7, 0x42, 0x5e, 0x73, 0x04, 0x33, 0x62, 0x93, 0x8b, 0x98, 0x24},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sum(tt.data)
			if len(got) != len(tt.expected) {
				t.Errorf("Sum() 长度错误: got %v, want %v", len(got), len(tt.expected))
				return
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Errorf("Sum() 第 %d 个字节错误: got %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}

func TestSumHex(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{
			"空数据",
			[]byte{},
			"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			"简单字符串",
			[]byte("hello"),
			"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			"较长字符串",
			[]byte("The quick brown fox jumps over the lazy dog"),
			"d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumHex(tt.data)
			if got != tt.expected {
				t.Errorf("SumHex() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSumString(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected string // 十六进制表示，方便比较
	}{
		{
			"空字符串",
			"",
			"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			"简单字符串",
			"hello",
			"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumString(tt.data)
			expectedBytes, _ := hex.DecodeString(tt.expected)
			if len(got) != len(expectedBytes) {
				t.Errorf("SumString() 长度错误: got %v, want %v", len(got), len(expectedBytes))
				return
			}
			for i := range got {
				if got[i] != expectedBytes[i] {
					t.Errorf("SumString() 第 %d 个字节错误: got %v, want %v", i, got[i], expectedBytes[i])
				}
			}
		})
	}
}

func TestSumStringHex(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected string
	}{
		{
			"空字符串",
			"",
			"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			"简单字符串",
			"hello",
			"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			"较长字符串",
			"The quick brown fox jumps over the lazy dog",
			"d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SumStringHex(tt.data)
			if got != tt.expected {
				t.Errorf("SumStringHex() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string // 十六进制表示，方便比较
	}{
		{
			"空数据",
			[]byte{},
			"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			"简单字符串",
			[]byte("hello"),
			"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New()
			h.Write(tt.data)
			gotBytes := h.Sum(nil)
			expectedBytes, _ := hex.DecodeString(tt.expected)
			if len(gotBytes) != len(expectedBytes) {
				t.Errorf("New().Sum() 长度错误: got %v, want %v", len(gotBytes), len(expectedBytes))
				return
			}
			for i := range gotBytes {
				if gotBytes[i] != expectedBytes[i] {
					t.Errorf("New().Sum() 第 %d 个字节错误: got %v, want %v", i, gotBytes[i], expectedBytes[i])
				}
			}
		})
	}
}

func TestDigestReset(t *testing.T) {
	h := New()
	h.Write([]byte("hello"))
	firstHash := hex.EncodeToString(h.Sum(nil))

	h.Reset()
	// 重置后应该得到空数据的哈希值
	resetHash := hex.EncodeToString(h.Sum(nil))
	if resetHash != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
		t.Errorf("Reset() 后哈希值错误: got %v, want e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", resetHash)
	}

	// 重置后再次写入数据应该正常工作
	h.Write([]byte("hello"))
	secondHash := hex.EncodeToString(h.Sum(nil))
	if firstHash != secondHash {
		t.Errorf("Reset() 后再次计算哈希值不一致: got %v, want %v", secondHash, firstHash)
	}
}

func TestDigestSizeAndBlockSize(t *testing.T) {
	h := New()
	if h.Size() != 32 {
		t.Errorf("Size() = %v, want 32", h.Size())
	}
	if h.BlockSize() != 64 {
		t.Errorf("BlockSize() = %v, want 64", h.BlockSize())
	}
}

func TestSumAppended(t *testing.T) {
	h := New()
	h.Write([]byte("hello"))

	// 测试Sum方法的附加功能
	prefix := []byte{0x01, 0x02, 0x03}
	sum := h.Sum(prefix)
	if len(sum) != len(prefix)+32 {
		t.Errorf("Sum() 附加功能长度错误: got %v, want %v", len(sum), len(prefix)+32)
	}

	// 验证前缀是否正确
	for i, b := range prefix {
		if sum[i] != b {
			t.Errorf("Sum() 前缀错误: got %v, want %v", sum[i], b)
		}
	}

	// 验证哈希部分是否正确
	expectedHash := []byte("2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824")
	expectedBytes, _ := hex.DecodeString(string(expectedHash))
	for i, b := range expectedBytes {
		if sum[i+len(prefix)] != b {
			t.Errorf("Sum() 哈希部分错误: got %v, want %v", sum[i+len(prefix)], b)
		}
	}
}
