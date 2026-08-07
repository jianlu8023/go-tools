package ed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

const (
	PrivateKeyPEMType = "PRIVATE KEY"
	PublicKeyPEMType  = "PUBLIC KEY"
)

// GenerateKey 生成Ed25519密钥对
// 返回公钥、私钥和错误信息
func GenerateKey() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(rand.Reader)
}

// Sign 使用Ed25519私钥对消息进行签名
// privateKey: Ed25519私钥
// message: 要签名的消息
// 返回签名和错误信息
func Sign(privateKey ed25519.PrivateKey, message []byte) ([]byte, error) {
	if privateKey == nil {
		return nil, fmt.Errorf("私钥不能为空")
	}

	if len(message) == 0 {
		return nil, fmt.Errorf("消息不能为空")
	}

	signature := ed25519.Sign(privateKey, message)
	return signature, nil
}

// Verify 使用Ed25519公钥验证签名
// publicKey: Ed25519公钥
// message: 要验证的消息
// signature: 签名
// 返回验证结果和错误信息
func Verify(publicKey ed25519.PublicKey, message, signature []byte) (bool, error) {
	if publicKey == nil {
		return false, fmt.Errorf("公钥不能为空")
	}

	if len(message) == 0 {
		return false, fmt.Errorf("消息不能为空")
	}

	if len(signature) == 0 {
		return false, fmt.Errorf("签名不能为空")
	}

	valid := ed25519.Verify(publicKey, message, signature)
	return valid, nil
}

// SignText 使用Ed25519私钥对文本进行签名
// privateKey: Ed25519私钥
// text: 要签名的文本
// 返回签名和错误信息
func SignText(privateKey ed25519.PrivateKey, text string) ([]byte, error) {
	return Sign(privateKey, []byte(text))
}

// VerifyText 使用Ed25519公钥验证文本签名
// publicKey: Ed25519公钥
// text: 要验证的文本
// signature: 签名
// 返回验证结果和错误信息
func VerifyText(publicKey ed25519.PublicKey, text string, signature []byte) (bool, error) {
	return Verify(publicKey, []byte(text), signature)
}

// SavePrivateKey 保存私钥到PEM文件
// privateKey: Ed25519私钥
// filename: 文件名
// 返回错误信息
func SavePrivateKey(privateKey ed25519.PrivateKey, filename string) (err error) {
	// 将私钥编码为PKCS#8格式
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("编码私钥失败: %v", err)
	}

	// 创建PEM块
	privateKeyPEM := &pem.Block{
		Type:  PrivateKeyPEMType,
		Bytes: privateKeyBytes,
	}

	// 写入文件
	privateKeyFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建私钥文件失败: %v", err)
	}
	defer func() {
		if closeErr := privateKeyFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		return fmt.Errorf("写入私钥文件失败: %v", err)
	}

	return nil
}

// SavePublicKey 保存公钥到PEM文件
// publicKey: Ed25519公钥
// filename: 文件名
// 返回错误信息
func SavePublicKey(publicKey ed25519.PublicKey, filename string) (err error) {
	// 将公钥编码为PKIX格式
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("编码公钥失败: %v", err)
	}

	// 创建PEM块
	publicKeyPEM := &pem.Block{
		Type:  PublicKeyPEMType,
		Bytes: publicKeyBytes,
	}

	// 写入文件
	publicKeyFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建公钥文件失败: %v", err)
	}
	defer func() {
		if closeErr := publicKeyFile.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		return fmt.Errorf("写入公钥文件失败: %v", err)
	}

	return nil
}

// LoadPrivateKey 从PEM文件加载私钥
// filename: 文件名
// 返回Ed25519私钥和错误信息
func LoadPrivateKey(filename string) (ed25519.PrivateKey, error) {
	// 读取文件
	privateKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取私钥文件失败: %v", err)
	}

	// 解码PEM块（只处理第一个PEM块，剩余数据将被忽略）
	privateKeyPEM, _ := pem.Decode(privateKeyBytes)
	if privateKeyPEM == nil {
		return nil, fmt.Errorf("解析私钥PEM块失败")
	}

	// 解析PKCS#8私钥
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyPEM.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析私钥失败: %v", err)
	}

	// 类型断言
	ed25519PrivateKey, ok := privateKey.(ed25519.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("私钥类型不正确")
	}

	return ed25519PrivateKey, nil
}

// LoadPublicKey 从PEM文件加载公钥
// filename: 文件名
// 返回Ed25519公钥和错误信息
func LoadPublicKey(filename string) (ed25519.PublicKey, error) {
	// 读取文件
	publicKeyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取公钥文件失败: %v", err)
	}

	// 解码PEM块（只处理第一个PEM块，剩余数据将被忽略）
	publicKeyPEM, _ := pem.Decode(publicKeyBytes)
	if publicKeyPEM == nil {
		return nil, fmt.Errorf("解析公钥PEM块失败")
	}

	// 解析PKIX公钥
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyPEM.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析公钥失败: %v", err)
	}

	// 类型断言
	ed25519PublicKey, ok := publicKey.(ed25519.PublicKey)
	if !ok {
		return nil, fmt.Errorf("公钥类型不正确")
	}

	return ed25519PublicKey, nil
}
