package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

// HashFunc 哈希函数类型
type HashFunc func() hash.Hash

// 支持的哈希函数
// 注意：这些变量为函数类型，外部代码不应重新赋值，否则会破坏包内的默认行为。
var (
	SHA256 HashFunc = sha256.New
	SHA384 HashFunc = sha512.New384
	SHA512 HashFunc = sha512.New
)

// Sign 使用HMAC对消息进行签名
// key: 密钥
// message: 消息
// hashFunc: 哈希函数
// 返回签名和错误信息
func Sign(key, message []byte, hashFunc HashFunc) ([]byte, error) {
	// 验证参数
	if len(key) == 0 {
		return nil, fmt.Errorf("密钥不能为空")
	}

	if len(message) == 0 {
		return nil, fmt.Errorf("消息不能为空")
	}

	if hashFunc == nil {
		return nil, fmt.Errorf("哈希函数不能为空")
	}

	// 创建HMAC
	mac := hmac.New(hashFunc, key)

	// 写入消息
	_, err := mac.Write(message)
	if err != nil {
		return nil, fmt.Errorf("写入消息失败: %v", err)
	}

	// 计算签名
	signature := mac.Sum(nil)

	return signature, nil
}

// Verify 使用HMAC验证签名
// key: 密钥
// message: 消息
// signature: 签名
// hashFunc: 哈希函数
// 返回验证结果和错误信息
func Verify(key, message, signature []byte, hashFunc HashFunc) (bool, error) {
	// 验证参数
	if len(key) == 0 {
		return false, fmt.Errorf("密钥不能为空")
	}

	if len(message) == 0 {
		return false, fmt.Errorf("消息不能为空")
	}

	if len(signature) == 0 {
		return false, fmt.Errorf("签名不能为空")
	}

	if hashFunc == nil {
		return false, fmt.Errorf("哈希函数不能为空")
	}

	// 创建HMAC
	mac := hmac.New(hashFunc, key)

	// 写入消息
	_, err := mac.Write(message)
	if err != nil {
		return false, fmt.Errorf("写入消息失败: %v", err)
	}

	// 验证签名
	valid := hmac.Equal(mac.Sum(nil), signature)

	return valid, nil
}

// SignText 使用HMAC对文本进行签名
// key: 密钥
// text: 文本
// hashFunc: 哈希函数
// 返回签名和错误信息
func SignText(key []byte, text string, hashFunc HashFunc) ([]byte, error) {
	return Sign(key, []byte(text), hashFunc)
}

// VerifyText 使用HMAC验证文本签名
// key: 密钥
// text: 文本
// signature: 签名
// hashFunc: 哈希函数
// 返回验证结果和错误信息
func VerifyText(key []byte, text string, signature []byte, hashFunc HashFunc) (bool, error) {
	return Verify(key, []byte(text), signature, hashFunc)
}

// SignSHA256 使用HMAC-SHA256对消息进行签名
// key: 密钥
// message: 消息
// 返回签名和错误信息
func SignSHA256(key, message []byte) ([]byte, error) {
	return Sign(key, message, SHA256)
}

// VerifySHA256 使用HMAC-SHA256验证签名
// key: 密钥
// message: 消息
// signature: 签名
// 返回验证结果和错误信息
func VerifySHA256(key, message, signature []byte) (bool, error) {
	return Verify(key, message, signature, SHA256)
}

// SignTextSHA256 使用HMAC-SHA256对文本进行签名
// key: 密钥
// text: 文本
// 返回签名和错误信息
func SignTextSHA256(key []byte, text string) ([]byte, error) {
	return SignText(key, text, SHA256)
}

// VerifyTextSHA256 使用HMAC-SHA256验证文本签名
// key: 密钥
// text: 文本
// signature: 签名
// 返回验证结果和错误信息
func VerifyTextSHA256(key []byte, text string, signature []byte) (bool, error) {
	return VerifyText(key, text, signature, SHA256)
}
