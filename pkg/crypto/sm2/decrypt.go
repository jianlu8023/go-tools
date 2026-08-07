package sm2

import (
	"encoding/binary"
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
//
// 警告: SM2 为非对称加密算法，按块调用私钥解密性能远低于对称加密（如 SM4），
// 不适合解密大文件。建议对大文件采用混合加密方案：
// 先用 SM4 等对称算法解密文件内容，再用 SM2 解密 SM4 密钥。
//
// 文件格式: 与 EncryptFile 配套，循环读取 4 字节大端序长度前缀，
// 再按长度读取对应密文块，逐块解密拼接。
//
// inputFile: 待解密的文件路径
// outputFile: 解密后的文件路径
// privateKey: SM2 私钥
// c1c2c3: 是否使用 C1C2C3 格式
// error: 错误信息
func DecryptFile(inputFile, outputFile string, privateKey *sm2.PrivateKey, c1c2c3 bool) (err error) {
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
	defer func() {
		if closeErr := outFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	var mode int
	if c1c2c3 {
		mode = sm2.C1C2C3
	} else {
		mode = sm2.C1C3C2
	}

	lenBuf := make([]byte, 4)
	for {
		// 读取 4 字节长度前缀
		_, err = io.ReadFull(inFile, lenBuf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("读取密文长度前缀失败: %v", err)
		}
		cipherLen := binary.BigEndian.Uint32(lenBuf)

		// 读取对应长度的密文块
		cipherData := make([]byte, cipherLen)
		if _, err = io.ReadFull(inFile, cipherData); err != nil {
			return fmt.Errorf("读取密文块失败: %v", err)
		}

		// 解密
		plainData, decErr := sm2.Decrypt(privateKey, cipherData, mode)
		if decErr != nil {
			return fmt.Errorf("解密文件失败: %v", decErr)
		}

		// 写入明文
		if _, wErr := outFile.Write(plainData); wErr != nil {
			return fmt.Errorf("写入解密文件失败: %v", wErr)
		}
	}

	return nil
}
