package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// CBCEncrypt 使用DES CBC模式加密字节切片
// key: 8字节的密钥
// plaintext: 明文字节切片
// 返回加密后的字节切片和错误信息
func CBCEncrypt(key, plaintext []byte) ([]byte, error) {
	// 验证密钥长度
	if len(key) != 8 {
		return nil, fmt.Errorf("DES密钥长度必须为8字节，当前长度为%d字节", len(key))
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建DES块错误: %v", err)
	}

	// 处理数据并进行填充
	paddedData := pkcs5Padding(plaintext, blockSize)

	// 生成随机IV
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("生成IV向量错误: %v", err)
	}

	// 加密数据
	mode := cipher.NewCBCEncrypter(block, iv)
	ciphertext := make([]byte, len(paddedData))
	mode.CryptBlocks(ciphertext, paddedData)

	// 将IV附加到密文前面
	result := make([]byte, blockSize+len(ciphertext))
	copy(result[:blockSize], iv)
	copy(result[blockSize:], ciphertext)

	return result, nil
}

// CBCDecrypt 使用DES CBC模式解密字节切片
// key: 8字节的密钥
// ciphertext: 密文字节切片（包含IV）
// 返回解密后的字节切片和错误信息
func CBCDecrypt(key, ciphertext []byte) ([]byte, error) {
	// 验证密钥长度
	if len(key) != 8 {
		return nil, fmt.Errorf("DES密钥长度必须为8字节，当前长度为%d字节", len(key))
	}

	// 验证密文长度
	if len(ciphertext) == 0 || len(ciphertext)%blockSize != 0 {
		return nil, fmt.Errorf("密文长度必须为%d的倍数", blockSize)
	}

	block, err := des.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("创建DES块错误: %v", err)
	}

	// 提取IV和密文
	iv := ciphertext[:blockSize]
	actualCiphertext := ciphertext[blockSize:]

	// 解密数据
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(actualCiphertext))
	mode.CryptBlocks(plaintext, actualCiphertext)

	// 移除填充
	unpaddedData, err := pkcs5UnPadding(plaintext)
	if err != nil {
		return nil, fmt.Errorf("移除填充错误: %v", err)
	}

	return unpaddedData, nil
}

// CBCEncryptText 使用DES CBC模式加密文本
// key: 8字节的密钥
// text: 明文文本
// 返回加密后的字节切片和错误信息
func CBCEncryptText(key []byte, text string) ([]byte, error) {
	return CBCEncrypt(key, []byte(text))
}

// CBCDecryptText 使用DES CBC模式解密文本
// key: 8字节的密钥
// ciphertext: 密文字节切片（包含IV）
// 返回解密后的文本和错误信息
func CBCDecryptText(key, ciphertext []byte) (string, error) {
	plaintext, err := CBCDecrypt(key, ciphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// CBCEncryptFile 使用DES CBC模式加密文件
// key: 8字节的密钥
// inputFile: 输入文件路径
// outputFile: 输出文件路径
// error: 错误信息
func CBCEncryptFile(key, inputFile, outputFile string) error {
	// 验证密钥长度
	if len(key) != 8 {
		return fmt.Errorf("DES密钥长度必须为8字节，当前长度为%d字节", len(key))
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开输入文件错误: %v", err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件错误: %v", err)
	}
	defer outFile.Close()

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

	mode := cipher.NewCBCEncrypter(block, iv)
	buf := make([]byte, blockSize)
	buffer := bytes.NewBuffer(nil)

	for {
		n, err := inFile.Read(buf)
		if n > 0 {
			buffer.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("使用buffer读取输入文件错误: %v", err)
		}
	}

	// 处理最后一部分数据并进行填充
	paddedData := pkcs5Padding(buffer.Bytes(), blockSize)

	// 加密整个数据
	ciphertext := make([]byte, len(paddedData))
	mode.CryptBlocks(ciphertext, paddedData)

	// 写入加密后的数据
	if _, err := outFile.Write(ciphertext); err != nil {
		return fmt.Errorf("将加密结果写入输出文件错误: %v", err)
	}

	return nil
}

// CBCDecryptFile 使用DES CBC模式解密文件
// key: 8字节的密钥
// inputFile: 输入文件路径
// outputFile: 输出文件路径
// error: 错误信息
func CBCDecryptFile(key, inputFile, outputFile string) error {
	// 验证密钥长度
	if len(key) != 8 {
		return fmt.Errorf("DES密钥长度必须为8字节，当前长度为%d字节", len(key))
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开输入文件错误: %v", err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件错误: %v", err)
	}
	defer outFile.Close()

	block, err := des.NewCipher([]byte(key))
	if err != nil {
		return fmt.Errorf("创建DES块错误: %v", err)
	}

	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(inFile, iv); err != nil {
		return fmt.Errorf("读取IV向量错误: %v", err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	buf := make([]byte, blockSize)
	buffer := bytes.NewBuffer(nil)

	for {
		n, err := inFile.Read(buf)
		if n > 0 {
			buffer.Write(buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("使用buffer读取输入文件错误: %v", err)
		}
	}

	// 解密整个数据
	ciphertext := buffer.Bytes()
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 移除填充
	unPaddedBuf, err := pkcs5UnPadding(plaintext)
	if err != nil {
		return fmt.Errorf("移除填充错误: %v", err)
	}

	// 写入解密后的数据
	if _, err := outFile.Write(unPaddedBuf); err != nil {
		return fmt.Errorf("将解密结果写入输出文件错误: %v", err)
	}

	return nil
}
