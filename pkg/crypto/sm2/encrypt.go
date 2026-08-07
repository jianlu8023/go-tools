package sm2

import (
	"crypto/rand"
	"encoding/binary"
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
//
// 警告: SM2 为非对称加密算法，按块调用公钥加密性能远低于对称加密（如 SM4），
// 不适合加密大文件。建议对大文件采用混合加密方案：
// 先用 SM4 等对称算法加密文件内容，再用 SM2 加密 SM4 密钥。
//
// 文件格式: 每个密文块前写入 4 字节大端序长度前缀，即 [len(密文块)]+[密文块]，
// 所有块依次拼接，与 DecryptFile 配套使用。
//
// inputFile: 待加密的文件路径
// outputFile: 加密后的文件路径
// publicKey: SM2 公钥
// c1c2c3: 是否使用 C1C2C3 格式
// error: 错误信息
func EncryptFile(inputFile, outputFile string, publicKey *sm2.PublicKey, c1c2c3 bool) (err error) {
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
	defer func() {
		if closeErr := outFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	buffer := make([]byte, encryptBlockSize)

	var mode int
	if c1c2c3 {
		mode = sm2.C1C2C3
	} else {
		mode = sm2.C1C3C2
	}

	lenBytes := make([]byte, 4)
	for {
		n, readErr := io.ReadFull(inFile, buffer)
		if n > 0 {
			encryptedBlock, encErr := sm2.Encrypt(publicKey, buffer[:n], rand.Reader, mode)
			if encErr != nil {
				return fmt.Errorf("加密文件失败: %v", encErr)
			}

			// 写入 4 字节大端序长度前缀 + 密文
			binary.BigEndian.PutUint32(lenBytes, uint32(len(encryptedBlock)))
			if _, wErr := outFile.Write(lenBytes); wErr != nil {
				return fmt.Errorf("写入加密文件长度前缀失败: %v", wErr)
			}
			if _, wErr := outFile.Write(encryptedBlock); wErr != nil {
				return fmt.Errorf("写入加密文件失败: %v", wErr)
			}
		}
		if readErr == io.EOF || readErr == io.ErrUnexpectedEOF {
			break
		}
		if readErr != nil {
			return fmt.Errorf("使用buffer读取输入文件错误: %v", readErr)
		}
	}
	return nil
}
