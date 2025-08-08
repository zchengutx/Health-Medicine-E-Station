package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// AESEncrypt 使用AES算法加密数据
// key: 加密密钥，长度必须是16、24或32字节
// plaintext: 要加密的明文
// 返回: base64编码的密文和错误信息
func AESEncrypt(key []byte, plaintext string) (string, error) {
	// 创建AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 将明文转换为字节数组
	plaintextBytes := []byte(plaintext)

	// 创建GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成随机nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密数据
	ciphertext := gcm.Seal(nonce, nonce, plaintextBytes, nil)

	// 返回base64编码的结果
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AESDecrypt 使用AES算法解密数据
// key: 解密密钥，必须与加密时使用的密钥相同
// ciphertext: base64编码的密文
// 返回: 解密后的明文和错误信息
func AESDecrypt(key []byte, ciphertext string) (string, error) {
	// 创建AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 解码base64
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// 创建GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 检查密文长度
	nonceSize := gcm.NonceSize()
	if len(ciphertextBytes) < nonceSize {
		return "", errors.New("密文长度不足")
	}

	// 提取nonce和实际密文
	nonce, ciphertextBytes := ciphertextBytes[:nonceSize], ciphertextBytes[nonceSize:]

	// 解密数据
	plaintextBytes, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintextBytes), nil
}

// GenerateAESKey 生成指定长度的AES密钥
// keySize: 密钥长度，支持16、24、32字节（对应AES-128、AES-192、AES-256）
// 返回: 生成的密钥和错误信息
func GenerateAESKey(keySize int) ([]byte, error) {
	if keySize != 16 && keySize != 24 && keySize != 32 {
		return nil, errors.New("密钥长度必须是16、24或32字节")
	}

	key := make([]byte, keySize)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

// StringToAESKey 将字符串转换为AES密钥
// keyStr: 密钥字符串
// keySize: 目标密钥长度
// 返回: 处理后的密钥
func StringToAESKey(keyStr string, keySize int) []byte {
	key := []byte(keyStr)
	
	// 如果密钥长度不足，用0填充
	if len(key) < keySize {
		padding := make([]byte, keySize-len(key))
		key = append(key, padding...)
	}
	
	// 如果密钥长度超出，截取前面部分
	if len(key) > keySize {
		key = key[:keySize]
	}
	
	return key
}