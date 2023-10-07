package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// keyMutex 生成公私钥对时同步块
var (
	keyMutex = sync.Mutex{}
)

// RsaEncrypt
// @Description: RsaEncrypt
// @author ght
// @date 2023-10-07 16:40:08
// @param data: 待加密的数据
// @return []byte: 加密结果
func RsaEncrypt(data []byte) []byte {
	if !checkRsaFile() {
		generateKeyPair()
	}
	wd, _ := os.Getwd()
	publicKeyPath := filepath.Join(wd, "metadata", "rsa", "rsa_public_key.pem")
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println(fmt.Errorf("读取公钥失败:%s", err))
	}
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		fmt.Println(fmt.Errorf("私钥无效:%s", err))
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(fmt.Errorf("解析公钥失败:%s", err))
	}
	key := publicKey.(*rsa.PublicKey)
	v15, err := rsa.EncryptPKCS1v15(rand.Reader, key, data)

	if err != nil {
		fmt.Println(fmt.Errorf("加密失败:%s", err))
	}
	return v15
}

// RsaDecrypt
// @Description: RsaDecrypt
// @author ght
// @date 2023-10-07 16:40:13
// @param data: 待解密数据
// @return []byte: 解密结果
func RsaDecrypt(data []byte) []byte {
	if !checkRsaFile() {
		generateKeyPair()
	}
	wd, _ := os.Getwd()
	privateKeyPath := filepath.Join(wd, "metadata", "rsa", "rsa_private_key.pem")
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		fmt.Println(fmt.Errorf("读取公钥失败:%s", err))
	}
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		fmt.Println(fmt.Errorf("私钥无效:%s", err))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(fmt.Errorf("解析私钥失败:%s", err))
	}

	v15, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
	if err != nil {
		fmt.Println(fmt.Errorf("解密失败:%s", err))
	}
	return v15
}

// generateKeyPair
// @Description: generateKeyPair
// @author ght
// @date 2023-10-07 16:39:06
func generateKeyPair() {
	keyMutex.Lock()
	defer keyMutex.Unlock()
	if checkRsaFile() {
		return
	}
	key, err := rsa.GenerateKey(rand.Reader, 8192)
	if err != nil {
		fmt.Println(fmt.Errorf("生成密钥对失败:%s", err))
		return
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(key)

	privateKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})
	wd, _ := os.Getwd()

	privateKeyPath := filepath.Join(wd, "metadata", "rsa", "rsa_private_key.pem")
	err = saveFile(privateKeyPath, privateKeyPem)
	if err != nil {
		fmt.Println(fmt.Errorf("保存rsa私钥失败:%s", err))
	}
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		fmt.Println(fmt.Errorf("获取rsa公钥失败:%s", err))
	}
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	publicKeyPath := filepath.Join(wd, "metadata", "rsa", "rsa_public_key.pem")
	err = saveFile(publicKeyPath, publicKeyPem)
	if err != nil {
		fmt.Println(fmt.Errorf("保存Rsa公钥失败:%s", err))
	}
}

// saveFile
// @Description: saveFile
// @author ght
// @date 2023-10-07 16:39:22
// @param path: 文件存储路径
// @param content: 文件内容
// @return error: 保存文件可能出现的错误
func saveFile(path string, content []byte) error {
	dir := filepath.Dir(path)
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 	文件夹不存在，创建文件夹
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Println(fmt.Errorf("创建文件夹失败:%s", err))
		}
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Println(fmt.Errorf("创建文件 %s 失败:%s", path, err))
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(fmt.Errorf("将内容写入 %s 失败:%s", path, err))
		return err
	}
	return nil
}

// checkRsaFile
// @Description: checkRsaFile
// @author ght
// @date 2023-10-07 16:39:35
// @return bool: 是否存在公私钥文件
func checkRsaFile() bool {
	wd, _ := os.Getwd()
	privateKeyPath := filepath.Join(wd, "metadata", "rsa", "rsa_private_key.pem")
	publicKeyPath := filepath.Join(wd, "metadata", "rsa", "rsa_public_key.pem")
	_, privateKeyerr := os.Stat(privateKeyPath)
	_, publicKeyerr := os.Stat(publicKeyPath)
	if privateKeyerr == nil && publicKeyerr == nil {
		return true
	} else if os.IsNotExist(privateKeyerr) {
		return false
	} else if os.IsNotExist(publicKeyerr) {
		return false
	} else {
		return false
	}
}
