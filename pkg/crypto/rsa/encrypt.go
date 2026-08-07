package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"os"
)

// Encrypt 使用RSA公钥加密字节切片
// plaintext: 明文字节切片
// publicKey: RSA公钥
// 返回加密后的字节切片和错误信息
func Encrypt(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	// 验证公钥
	if publicKey == nil {
		return nil, fmt.Errorf("公钥不能为空")
	}

	// 验证明文长度
	blockSize := publicKey.Size() - 11
	if len(plaintext) > blockSize {
		return nil, fmt.Errorf("明文长度超过最大限制(%d字节)", blockSize)
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err != nil {
		return nil, fmt.Errorf("加密失败: %v", err)
	}

	return encryptedData, nil
}

// EncryptText 使用RSA公钥加密文本
// text: 明文文本
// publicKey: RSA公钥
// 返回加密后的字节切片和错误信息
func EncryptText(text string, publicKey *rsa.PublicKey) ([]byte, error) {
	return Encrypt([]byte(text), publicKey)
}

// EncryptFile 使用RSA公钥加密文件
// inputFile: 待加密文件路径
// outputFile: 加密后文件路径
// publicKey: RSA公钥
// error: 错误信息
func EncryptFile(inputFile, outputFile string, publicKey *rsa.PublicKey) (err error) {
	// 验证参数
	if publicKey == nil {
		return fmt.Errorf("公钥不能为空")
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开待加密文件失败: %v", err)
	}
	defer func() {
		if closeErr := inFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建加密文件失败: %v", err)
	}
	defer func() {
		if closeErr := outFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	blockSize := publicKey.Size() - 11
	buffer := make([]byte, blockSize)

	for {
		n, err := inFile.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("使用buffer读取输入文件错误: %v", err)
		}
		if n == 0 {
			break
		}

		// 检查块大小
		if n > blockSize {
			return fmt.Errorf("读取的数据块大小超过限制")
		}

		encryptedBlock, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, buffer[:n])
		if err != nil {
			return fmt.Errorf("加密文件失败: %v", err)
		}

		_, err = outFile.Write(encryptedBlock)
		if err != nil {
			return fmt.Errorf("写入加密文件失败: %v", err)
		}
	}
	return nil
}
