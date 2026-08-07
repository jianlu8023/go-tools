package des

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// CTREncrypt 使用DES CTR模式加密字节切片
// key: 8字节的密钥
// plaintext: 明文字节切片
// 返回加密后的字节切片和错误信息
func CTREncrypt(key, plaintext []byte) ([]byte, error) {
	// 验证密钥长度
	if len(key) != 8 {
		return nil, fmt.Errorf("DES密钥长度必须为8字节，当前长度为%d字节", len(key))
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建DES块错误: %v", err)
	}

	// 生成随机IV
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("生成IV向量错误: %v", err)
	}

	// 加密数据
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// 将IV附加到密文前面
	result := make([]byte, blockSize+len(ciphertext))
	copy(result[:blockSize], iv)
	copy(result[blockSize:], ciphertext)

	return result, nil
}

// CTRDecrypt 使用DES CTR模式解密字节切片
// key: 8字节的密钥
// ciphertext: 密文字节切片（包含IV）
// 返回解密后的字节切片和错误信息
func CTRDecrypt(key, ciphertext []byte) ([]byte, error) {
	// 验证密钥长度
	if len(key) != 8 {
		return nil, fmt.Errorf("DES密钥长度必须为8字节，当前长度为%d字节", len(key))
	}

	// 验证密文长度
	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("密文长度不足，至少需要%d字节", blockSize)
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建DES块错误: %v", err)
	}

	// 提取IV和密文
	iv := ciphertext[:blockSize]
	actualCiphertext := ciphertext[blockSize:]

	// 解密数据
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(actualCiphertext))
	stream.XORKeyStream(plaintext, actualCiphertext)

	return plaintext, nil
}

// CTREncryptText 使用DES CTR模式加密文本
// key: 8字节的密钥
// text: 明文文本
// 返回加密后的字节切片和错误信息
func CTREncryptText(key []byte, text string) ([]byte, error) {
	return CTREncrypt(key, []byte(text))
}

// CTRDecryptText 使用DES CTR模式解密文本
// key: 8字节的密钥
// ciphertext: 密文字节切片（包含IV）
// 返回解密后的文本和错误信息
func CTRDecryptText(key, ciphertext []byte) (string, error) {
	plaintext, err := CTRDecrypt(key, ciphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// CTREncryptFile 使用des ctr模式加密文件
// key: 8字节的密钥
// inputFile: 输入文件
// outputFile: 输出文件
// error: 错误信息
func CTREncryptFile(key, inputFile, outputFile string) (err error) {
	// 验证密钥长度
	if len(key) != 8 {
		return fmt.Errorf("DES密钥长度必须为8字节，当前长度为%d字节", len(key))
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

	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return fmt.Errorf("创建DES块错误: %v", err)
	}

	iv := make([]byte, blockSize)
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

// CTRDecryptFile 使用des ctr模式解密文件
// key: 8字节的密钥
// inputFile: 输入文件
// outputFile: 输出文件
// error: 错误信息
func CTRDecryptFile(key, inputFile string, outputFile string) (err error) {
	// 验证密钥长度
	if len(key) != 8 {
		return fmt.Errorf("DES密钥长度必须为8字节，当前长度为%d字节", len(key))
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

	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return fmt.Errorf("创建DES块错误: %v", err)
	}

	iv := make([]byte, blockSize)
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
