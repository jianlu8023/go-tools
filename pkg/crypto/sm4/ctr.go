package sm4

import (
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/tjfoc/gmsm/sm4"
)

// CTREncrypt 使用SM4 CTR模式加密字节切片
// key: 16字节密钥
// plaintext: 明文字节切片
// 返回加密后的字节切片和错误信息
func CTREncrypt(key, plaintext []byte) ([]byte, error) {
	// 验证密钥长度
	if len(key) != 16 {
		return nil, fmt.Errorf("SM4密钥长度必须为16字节，当前长度为%d字节", len(key))
	}

	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建SM4块错误: %v", err)
	}

	// 生成随机IV
	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("生成IV向量错误: %v", err)
	}

	// 加密数据
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// 将IV附加到密文前面
	result := make([]byte, block.BlockSize()+len(ciphertext))
	copy(result[:block.BlockSize()], iv)
	copy(result[block.BlockSize():], ciphertext)

	return result, nil
}

// CTRDecrypt 使用SM4 CTR模式解密字节切片
// key: 16字节密钥
// ciphertext: 密文字节切片（包含IV）
// 返回解密后的字节切片和错误信息
func CTRDecrypt(key, ciphertext []byte) ([]byte, error) {
	// 验证密钥长度
	if len(key) != 16 {
		return nil, fmt.Errorf("SM4密钥长度必须为16字节，当前长度为%d字节", len(key))
	}

	// 验证密文长度
	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("密文长度不足，至少需要%d字节", blockSize)
	}

	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建SM4块错误: %v", err)
	}

	// 提取IV和密文
	iv := ciphertext[:block.BlockSize()]
	actualCiphertext := ciphertext[block.BlockSize():]

	// 解密数据
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(actualCiphertext))
	stream.XORKeyStream(plaintext, actualCiphertext)

	return plaintext, nil
}

// CTREncryptText 使用SM4 CTR模式加密文本
// key: 16字节密钥
// text: 明文文本
// 返回加密后的字节切片和错误信息
func CTREncryptText(key []byte, text string) ([]byte, error) {
	return CTREncrypt(key, []byte(text))
}

// CTRDecryptText 使用SM4 CTR模式解密文本
// key: 16字节密钥
// ciphertext: 密文字节切片（包含IV）
// 返回解密后的文本和错误信息
func CTRDecryptText(key, ciphertext []byte) (string, error) {
	plaintext, err := CTRDecrypt(key, ciphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// CTREncryptFile 使用SM4 CTR模式加密文件
// key: 16字节密钥
// inputFile: 输入文件路径
// outputFile: 输出文件路径
// error: 错误信息
func CTREncryptFile(key, inputFile, outputFile string) (err error) {
	// 验证密钥长度
	if len(key) != 16 {
		return fmt.Errorf("SM4密钥长度必须为16字节，当前长度为%d字节", len(key))
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开输入文件错误: %v", err)
	}
	defer func() {
		if closeErr := inFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件错误: %v", err)
	}
	defer func() {
		if closeErr := outFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	block, err := sm4.NewCipher([]byte(key))
	if err != nil {
		return fmt.Errorf("创建SM4块错误: %v", err)
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return fmt.Errorf("生成IV向量错误: %v", err)
	}

	if _, err := outFile.Write(iv); err != nil {
		return fmt.Errorf("写入IV向量错误: %v", err)
	}

	stream := cipher.NewCTR(block, iv)
	buf := make([]byte, blockSize)
	for {
		n, err := inFile.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf[:n], buf[:n])
			if _, err := outFile.Write(buf[:n]); err != nil {
				return fmt.Errorf("将加密结果写入输出文件错误: %v", err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("使用buffer读取输入文件错误: %v", err)
		}
	}
	return nil
}

// CTRDecryptFile 使用SM4 CTR模式解密文件
// key: 16字节密钥
// inputFile: 输入文件路径
// outputFile: 输出文件路径
// error: 错误信息
func CTRDecryptFile(key, inputFile, outputFile string) (err error) {
	// 验证密钥长度
	if len(key) != 16 {
		return fmt.Errorf("SM4密钥长度必须为16字节，当前长度为%d字节", len(key))
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开输入文件错误: %v", err)
	}
	defer func() {
		if closeErr := inFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件错误: %v", err)
	}
	defer func() {
		if closeErr := outFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	block, err := sm4.NewCipher([]byte(key))
	if err != nil {
		return fmt.Errorf("创建SM4块错误: %v", err)
	}

	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(inFile, iv); err != nil {
		return fmt.Errorf("读取IV向量错误: %v", err)
	}

	stream := cipher.NewCTR(block, iv)
	buf := make([]byte, blockSize)

	for {
		n, err := inFile.Read(buf)
		if n > 0 {
			stream.XORKeyStream(buf[:n], buf[:n])
			if _, err := outFile.Write(buf[:n]); err != nil {
				return fmt.Errorf("将加密结果写入输出文件错误: %v", err)
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("使用buffer读取输入文件错误: %v", err)
		}
	}
	return nil
}
