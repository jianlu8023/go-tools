package sm4

import (
	"bytes"
	"fmt"

	"github.com/tjfoc/gmsm/sm4"
)

const blockSize = sm4.BlockSize // SM4 块大小为 16 字节

// pkcs7Padding 添加 PKCS#7 填充
func pkcs7Padding(src []byte, blockSize int) []byte {
	if blockSize <= 0 {
		panic("块大小必须大于0")
	}
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	result := make([]byte, len(src), len(src)+padding)
	copy(result, src)
	return append(result, padText...)
}

// pkcs7UnPadding 移除 PKCS#7 填充
func pkcs7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, fmt.Errorf("输入为空")
	}

	// 检查是否是块大小的倍数
	if length%blockSize != 0 {
		return nil, fmt.Errorf("输入长度不是块大小(%d)的倍数", blockSize)
	}

	unPadding := int(src[length-1])

	// 验证填充值的有效性
	if unPadding <= 0 || unPadding > blockSize {
		return nil, fmt.Errorf("无效的填充长度: %d", unPadding)
	}

	// 验证填充内容的一致性
	for i := length - unPadding; i < length; i++ {
		if src[i] != byte(unPadding) {
			return nil, fmt.Errorf("填充内容不一致")
		}
	}

	return src[:length-unPadding], nil
}
