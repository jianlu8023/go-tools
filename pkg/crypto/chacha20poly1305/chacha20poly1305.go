package chacha20poly1305

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	// KeySize ChaCha20-Poly1305密钥长度
	KeySize = chacha20poly1305.KeySize

	// NonceSize ChaCha20-Poly1305 nonce长度
	NonceSize = chacha20poly1305.NonceSize
)

// GenerateKey 生成ChaCha20-Poly1305密钥
// 返回32字节密钥和错误信息
func GenerateKey() ([]byte, error) {
	key := make([]byte, KeySize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, fmt.Errorf("生成密钥失败: %v", err)
	}
	return key, nil
}

// Encrypt 使用ChaCha20-Poly1305加密数据
// key: 32字节密钥
// plaintext: 明文数据
// additionalData: 附加数据（可选）
// 返回密文和错误信息
func Encrypt(key, plaintext, additionalData []byte) ([]byte, error) {
	// 验证密钥长度
	if len(key) != KeySize {
		return nil, fmt.Errorf("密钥长度必须为%d字节，当前长度为%d字节", KeySize, len(key))
	}

	// 验证明文
	if len(plaintext) == 0 {
		return nil, fmt.Errorf("明文不能为空")
	}

	// 创建ChaCha20-Poly1305实例
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, fmt.Errorf("创建ChaCha20-Poly1305实例失败: %v", err)
	}

	// 生成随机nonce
	nonce := make([]byte, NonceSize)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("生成nonce失败: %v", err)
	}

	// 加密数据
	ciphertext := aead.Seal(nil, nonce, plaintext, additionalData)

	// 将nonce附加到密文前面
	result := make([]byte, NonceSize+len(ciphertext))
	copy(result[:NonceSize], nonce)
	copy(result[NonceSize:], ciphertext)

	return result, nil
}

// Decrypt 使用ChaCha20-Poly1305解密数据
// key: 32字节密钥
// ciphertext: 密文数据（包含nonce）
// additionalData: 附加数据（可选）
// 返回明文和错误信息
func Decrypt(key, ciphertext, additionalData []byte) ([]byte, error) {
	// 验证密钥长度
	if len(key) != KeySize {
		return nil, fmt.Errorf("密钥长度必须为%d字节，当前长度为%d字节", KeySize, len(key))
	}

	// 验证密文长度
	if len(ciphertext) < NonceSize {
		return nil, fmt.Errorf("密文长度不足，至少需要%d字节", NonceSize)
	}

	// 创建ChaCha20-Poly1305实例
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, fmt.Errorf("创建ChaCha20-Poly1305实例失败: %v", err)
	}

	// 提取nonce和密文
	nonce := ciphertext[:NonceSize]
	actualCiphertext := ciphertext[NonceSize:]

	// 解密数据
	plaintext, err := aead.Open(nil, nonce, actualCiphertext, additionalData)
	if err != nil {
		return nil, fmt.Errorf("解密失败: %v", err)
	}

	return plaintext, nil
}

// EncryptText 使用ChaCha20-Poly1305加密文本
// key: 32字节密钥
// text: 明文文本
// additionalData: 附加数据（可选）
// 返回密文和错误信息
func EncryptText(key []byte, text string, additionalData []byte) ([]byte, error) {
	return Encrypt(key, []byte(text), additionalData)
}

// DecryptText 使用ChaCha20-Poly1305解密文本
// key: 32字节密钥
// ciphertext: 密文数据（包含nonce）
// additionalData: 附加数据（可选）
// 返回明文文本和错误信息
func DecryptText(key, ciphertext, additionalData []byte) (string, error) {
	plaintext, err := Decrypt(key, ciphertext, additionalData)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// EncryptFile 使用ChaCha20-Poly1305加密文件
// key: 32字节密钥
// inputFile: 输入文件路径
// outputFile: 输出文件路径
// additionalData: 附加数据（可选）
// 返回错误信息
func EncryptFile(key []byte, inputFile, outputFile string, additionalData []byte) error {
	// 验证密钥长度
	if len(key) != KeySize {
		return fmt.Errorf("密钥长度必须为%d字节，当前长度为%d字节", KeySize, len(key))
	}

	// 打开输入文件
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开输入文件失败: %v", err)
	}
	defer inFile.Close()

	// 创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %v", err)
	}
	defer outFile.Close()

	// 读取文件内容
	plaintext, err := io.ReadAll(inFile)
	if err != nil {
		return fmt.Errorf("读取输入文件失败: %v", err)
	}

	// 加密数据
	ciphertext, err := Encrypt(key, plaintext, additionalData)
	if err != nil {
		return fmt.Errorf("加密文件失败: %v", err)
	}

	// 写入加密后的数据
	_, err = outFile.Write(ciphertext)
	if err != nil {
		return fmt.Errorf("写入输出文件失败: %v", err)
	}

	return nil
}

// DecryptFile 使用ChaCha20-Poly1305解密文件
// key: 32字节密钥
// inputFile: 输入文件路径
// outputFile: 输出文件路径
// additionalData: 附加数据（可选）
// 返回错误信息
func DecryptFile(key []byte, inputFile, outputFile string, additionalData []byte) error {
	// 验证密钥长度
	if len(key) != KeySize {
		return fmt.Errorf("密钥长度必须为%d字节，当前长度为%d字节", KeySize, len(key))
	}

	// 打开输入文件
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开输入文件失败: %v", err)
	}
	defer inFile.Close()

	// 创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %v", err)
	}
	defer outFile.Close()

	// 读取文件内容
	ciphertext, err := io.ReadAll(inFile)
	if err != nil {
		return fmt.Errorf("读取输入文件失败: %v", err)
	}

	// 解密数据
	plaintext, err := Decrypt(key, ciphertext, additionalData)
	if err != nil {
		return fmt.Errorf("解密文件失败: %v", err)
	}

	// 写入解密后的数据
	_, err = outFile.Write(plaintext)
	if err != nil {
		return fmt.Errorf("写入输出文件失败: %v", err)
	}

	return nil
}
