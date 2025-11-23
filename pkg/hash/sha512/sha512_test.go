package sha512

import (
	"crypto/sha512"
	"encoding/hex"
	"testing"
)

func TestSum(t *testing.T) {
	// 测试空数据
	hash := Sum([]byte{})
	if len(hash) != sha512.Size {
		t.Errorf("Sum() for empty data returned incorrect size: got %d, want %d", len(hash), sha512.Size)
	}

	// 测试单个字节
	hash = Sum([]byte{0x00})
	expected, _ := hex.DecodeString("b8244d028981d693af7b456af8efa4cad63d282e19ff14942c246e50d9351d22704a802a71c3580b6370de4ceb293c324a8423342557d4e5c38438f0e36910ee")
	if !bytesEqual(hash, expected) {
		t.Errorf("Sum() for single byte returned incorrect hash: got %s, want %s", hex.EncodeToString(hash), hex.EncodeToString(expected))
	}

	// 测试简单字符串
	hash = Sum([]byte("hello"))
	expected, _ = hex.DecodeString("9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043")
	if !bytesEqual(hash, expected) {
		t.Errorf("Sum() for 'hello' returned incorrect hash")
	}

	// 测试较长字符串
	hash = Sum([]byte("The quick brown fox jumps over the lazy dog"))
	expected, _ = hex.DecodeString("07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6")
	if !bytesEqual(hash, expected) {
		t.Errorf("Sum() for long string returned incorrect hash")
	}
}

func TestSumHex(t *testing.T) {
	// 测试空数据
	hash := SumHex([]byte{})
	expected := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"
	if hash != expected {
		t.Errorf("SumHex() for empty data returned incorrect hash: got %s, want %s", hash, expected)
	}

	// 测试简单字符串
	hash = SumHex([]byte("hello"))
	expected = "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
	if hash != expected {
		t.Errorf("SumHex() for 'hello' returned incorrect hash: got %s, want %s", hash, expected)
	}

	// 测试较长字符串
	hash = SumHex([]byte("The quick brown fox jumps over the lazy dog"))
	expected = "07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6"
	if hash != expected {
		t.Errorf("SumHex() for long string returned incorrect hash: got %s, want %s", hash, expected)
	}
}

func TestSumString(t *testing.T) {
	// 测试空字符串
	hash := SumString("")
	expected, _ := hex.DecodeString("cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e")
	if !bytesEqual(hash, expected) {
		t.Errorf("SumString() for empty string returned incorrect hash")
	}

	// 测试简单字符串
	hash = SumString("hello")
	expected, _ = hex.DecodeString("9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043")
	if !bytesEqual(hash, expected) {
		t.Errorf("SumString() for 'hello' returned incorrect hash")
	}
}

func TestSumStringHex(t *testing.T) {
	// 测试空字符串
	hash := SumStringHex("")
	expected := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"
	if hash != expected {
		t.Errorf("SumStringHex() for empty string returned incorrect hash: got %s, want %s", hash, expected)
	}

	// 测试简单字符串
	hash = SumStringHex("hello")
	expected = "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"
	if hash != expected {
		t.Errorf("SumStringHex() for 'hello' returned incorrect hash: got %s, want %s", hash, expected)
	}
}

func TestNew(t *testing.T) {
	digest := New()
	if digest == nil {
		t.Errorf("New() returned nil")
	}

	// 测试写入单个字节
	digest.Write([]byte{0x00})
	hash := digest.Sum(nil)
	expected, _ := hex.DecodeString("b8244d028981d693af7b456af8efa4cad63d282e19ff14942c246e50d9351d22704a802a71c3580b6370de4ceb293c324a8423342557d4e5c38438f0e36910ee")
	if !bytesEqual(hash, expected) {
		t.Errorf("New().Write().Sum() for single byte returned incorrect hash: got %s, want %s", hex.EncodeToString(hash), hex.EncodeToString(expected))
	}

	// 测试写入字符串
	digest = New()
	digest.Write([]byte("hello"))
	hash = digest.Sum(nil)
	expected, _ = hex.DecodeString("9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043")
	if !bytesEqual(hash, expected) {
		t.Errorf("New().Write().Sum() for 'hello' returned incorrect hash")
	}
}

func TestDigestReset(t *testing.T) {
	digest := New()
	digest.Write([]byte("hello"))
	digest.Reset()
	digest.Write([]byte("hello"))
	hash := digest.Sum(nil)
	expected, _ := hex.DecodeString("9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043")
	if !bytesEqual(hash, expected) {
		t.Errorf("Digest.Reset() did not work correctly")
	}
}

func TestDigestSizeAndBlockSize(t *testing.T) {
	digest := New()
	if size := digest.Size(); size != sha512.Size {
		t.Errorf("Digest.Size() returned %d, want %d", size, sha512.Size)
	}
	if blockSize := digest.BlockSize(); blockSize != sha512.BlockSize {
		t.Errorf("Digest.BlockSize() returned %d, want %d", blockSize, sha512.BlockSize)
	}
}

func TestSumAppended(t *testing.T) {
	digest := New()
	digest.Write([]byte("hello"))
	prefix := []byte("prefix:")
	hash := digest.Sum(prefix)
	if !bytesEqual(prefix, hash[:len(prefix)]) {
		t.Errorf("Sum() did not append to the provided slice correctly")
	}
	expectedHash, _ := hex.DecodeString("9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043")
	if !bytesEqual(hash[len(prefix):], expectedHash) {
		t.Errorf("Sum() appended incorrect hash")
	}
}

func TestSum512_224(t *testing.T) {
	hash := Sum512_224([]byte("hello"))
	expected, _ := hex.DecodeString("fe8509ed1fb7dcefc27e6ac1a80eddbec4cb3d2c6fe565244374061c")
	if !bytesEqual(hash, expected) {
		t.Errorf("Sum512_224() for 'hello' returned incorrect hash")
	}
}

func TestSum512_224Hex(t *testing.T) {
	hash := Sum512_224Hex([]byte("hello"))
	expected := "fe8509ed1fb7dcefc27e6ac1a80eddbec4cb3d2c6fe565244374061c"
	if hash != expected {
		t.Errorf("Sum512_224Hex() for 'hello' returned incorrect hash: got %s, want %s", hash, expected)
	}
}

func TestNew512_224(t *testing.T) {
	digest := New512_224()
	digest.Write([]byte("hello"))
	hash := digest.Sum(nil)
	expected, _ := hex.DecodeString("fe8509ed1fb7dcefc27e6ac1a80eddbec4cb3d2c6fe565244374061c")
	if !bytesEqual(hash, expected) {
		t.Errorf("New512_224().Write().Sum() for 'hello' returned incorrect hash")
	}
}

func TestSum512_256(t *testing.T) {
	hash := Sum512_256([]byte("hello"))
	expected, _ := hex.DecodeString("e30d87cfa2a75db545eac4d61baf970366a8357c7f72fa95b52d0accb698f13a")
	if !bytesEqual(hash, expected) {
		t.Errorf("Sum512_256() for 'hello' returned incorrect hash")
	}
}

func TestSum512_256Hex(t *testing.T) {
	hash := Sum512_256Hex([]byte("hello"))
	expected := "e30d87cfa2a75db545eac4d61baf970366a8357c7f72fa95b52d0accb698f13a"
	if hash != expected {
		t.Errorf("Sum512_256Hex() for 'hello' returned incorrect hash: got %s, want %s", hash, expected)
	}
}

func TestNew512_256(t *testing.T) {
	digest := New512_256()
	digest.Write([]byte("hello"))
	hash := digest.Sum(nil)
	expected, _ := hex.DecodeString("e30d87cfa2a75db545eac4d61baf970366a8357c7f72fa95b52d0accb698f13a")
	if !bytesEqual(hash, expected) {
		t.Errorf("New512_256().Write().Sum() for 'hello' returned incorrect hash")
	}
}

// 辅助函数，用于比较两个字节切片
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
