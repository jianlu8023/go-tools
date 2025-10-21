package sm2

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/tjfoc/gmsm/sm2"
)

// EncryptBytes 使用 SM2 算法加密字节切片
// plaintext: 明文字节切片
// publicKey: SM2 公钥
// c1c2c3: 是否使用 C1C2C3 格式
// 返回加密后的字节切片和错误信息
func EncryptBytes(plaintext []byte, publicKey *sm2.PublicKey, c1c2c3 bool) ([]byte, error) {
	// 验证公钥
	if publicKey == nil {
		return nil, fmt.Errorf("公钥不能为空")
	}

	// 验证明文
	if len(plaintext) == 0 {
		return nil, fmt.Errorf("明文不能为空")
	}

	var mode int
	if c1c2c3 {
		mode = sm2.C1C2C3
	} else {
		mode = sm2.C1C3C2
	}

	encryptedData, err := sm2.Encrypt(publicKey, plaintext, rand.Reader, mode)
	if err != nil {
		return nil, fmt.Errorf("加密失败: %v", err)
	}

	return encryptedData, nil
}

// EncryptText 使用 SM2 算法加密文本
// text: 明文文本
// publicKey: SM2 公钥
// c1c2c3: 是否使用 C1C2C3 格式
// 返回加密后的字节切片和错误信息
func EncryptText(text string, publicKey *sm2.PublicKey, c1c2c3 bool) ([]byte, error) {
	return EncryptBytes([]byte(text), publicKey, c1c2c3)
}

// EncryptFile 使用 SM2 算法加密文件
// inputFile: 待加密的文件路径
// outputFile: 加密后的文件路径
// publicKey: SM2 公钥
// c1c2c3: 是否使用 C1C2C3 格式
// error: 错误信息
func EncryptFile(inputFile, outputFile string, publicKey *sm2.PublicKey, c1c2c3 bool) error {
	// 验证参数
	if publicKey == nil {
		return fmt.Errorf("公钥不能为空")
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开待加密文件失败: %v", err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建加密文件失败: %v", err)
	}
	defer outFile.Close()

	buffer := make([]byte, encryptBlockSize)

	var mode int
	if c1c2c3 {
		mode = sm2.C1C2C3
	} else {
		mode = sm2.C1C3C2
	}

	for {
		n, err := inFile.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("使用buffer读取输入文件错误: %v", err)
		}
		if n == 0 {
			break
		}

		encryptedBlock, err := sm2.Encrypt(publicKey, buffer[:n], rand.Reader, mode)
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
