package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"os"
)

// Decrypt 使用RSA私钥解密字节切片
// ciphertext: 密文字节切片
// privateKey: RSA私钥
// 返回解密后的字节切片和错误信息
func Decrypt(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	// 验证私钥
	if privateKey == nil {
		return nil, fmt.Errorf("私钥不能为空")
	}

	// 验证密文长度
	if len(ciphertext) != privateKey.Size() {
		return nil, fmt.Errorf("密文长度不正确，期望%d字节，实际%d字节", privateKey.Size(), len(ciphertext))
	}

	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("解密失败: %v", err)
	}

	return decryptedData, nil
}

// DecryptText 使用RSA私钥解密文本
// ciphertext: 密文字节切片
// privateKey: RSA私钥
// 返回解密后的文本和错误信息
func DecryptText(ciphertext []byte, privateKey *rsa.PrivateKey) (string, error) {
	plaintext, err := Decrypt(ciphertext, privateKey)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// DecryptFile 使用私钥解密文件
// inputFile: 待解密文件路径
// outputFile: 解密后文件路径
// privateKey: RSA私钥
// error: 错误信息
func DecryptFile(inputFile, outputFile string, privateKey *rsa.PrivateKey) error {
	// 验证参数
	if privateKey == nil {
		return fmt.Errorf("私钥不能为空")
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开输入文件错误: %v", err)
	}
	defer func(inFile *os.File) {
		if err := inFile.Close(); err != nil {
			fmt.Printf("关闭输入文件错误: %v\n", err)
		}
	}(inFile)

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件错误: %v", err)
	}
	defer func(outFile *os.File) {
		if err := outFile.Close(); err != nil {
			fmt.Printf("关闭输出文件错误: %v\n", err)
		}
	}(outFile)

	blockSize := privateKey.Size()
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
		if n != blockSize {
			// 最后一个块可能小于blockSize，这是正常的
		}

		decryptedBlock, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, buffer[:n])
		if err != nil {
			return fmt.Errorf("解密文件失败: %v", err)
		}

		_, err = outFile.Write(decryptedBlock)
		if err != nil {
			return fmt.Errorf("写入解密文件失败: %v", err)
		}
	}
	return nil
}
