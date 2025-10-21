package sm2

import (
	"fmt"
	"io"
	"os"

	"github.com/tjfoc/gmsm/sm2"
)

// DecryptBytes 使用 SM2 算法解密字节切片
// ciphertext: 密文字节切片
// privateKey: SM2 私钥
// c1c2c3: 是否使用 C1C2C3 格式
// 返回解密后的字节切片和错误信息
func DecryptBytes(ciphertext []byte, privateKey *sm2.PrivateKey, c1c2c3 bool) ([]byte, error) {
	// 验证私钥
	if privateKey == nil {
		return nil, fmt.Errorf("私钥不能为空")
	}

	// 验证密文
	if len(ciphertext) == 0 {
		return nil, fmt.Errorf("密文不能为空")
	}

	var mode int
	if c1c2c3 {
		mode = sm2.C1C2C3
	} else {
		mode = sm2.C1C3C2
	}

	decryptedData, err := sm2.Decrypt(privateKey, ciphertext, mode)
	if err != nil {
		return nil, fmt.Errorf("解密失败: %v", err)
	}

	return decryptedData, nil
}

// DecryptText 使用 SM2 算法解密文本
// ciphertext: 密文字节切片
// privateKey: SM2 私钥
// c1c2c3: 是否使用 C1C2C3 格式
// 返回解密后的文本和错误信息
func DecryptText(ciphertext []byte, privateKey *sm2.PrivateKey, c1c2c3 bool) (string, error) {
	plaintext, err := DecryptBytes(ciphertext, privateKey, c1c2c3)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// DecryptFile 使用 SM2 算法解密文件
// inputFile: 待解密的文件路径
// outputFile: 解密后的文件路径
// privateKey: SM2 私钥
// c1c2c3: 是否使用 C1C2C3 格式
// error: 错误信息
func DecryptFile(inputFile, outputFile string, privateKey *sm2.PrivateKey, c1c2c3 bool) error {
	// 验证参数
	if privateKey == nil {
		return fmt.Errorf("私钥不能为空")
	}

	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("打开待解密文件失败: %v", err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("创建解密文件失败: %v", err)
	}
	defer outFile.Close()

	var mode int
	if c1c2c3 {
		mode = sm2.C1C2C3
	} else {
		mode = sm2.C1C3C2
	}

	buffer := make([]byte, decryptBlockSize)

	for {
		n, err := inFile.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("使用buffer读取输入文件错误: %v", err)
		}
		if n == 0 {
			break
		}

		decryptedBlock, err := sm2.Decrypt(privateKey, buffer[:n], mode)
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
