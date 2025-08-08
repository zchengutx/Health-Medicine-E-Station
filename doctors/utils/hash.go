package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 使用bcrypt算法对密码进行哈希加密
// password: 要加密的密码
// cost: 加密强度，建议使用bcrypt.DefaultCost(10)或更高
// 返回: 哈希后的密码和错误信息
func HashPassword(password string, cost int) (string, error) {
	if cost < bcrypt.MinCost {
		cost = bcrypt.DefaultCost
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("密码哈希失败: %w", err)
	}

	return string(hash), nil
}

// VerifyPassword 验证密码是否与哈希值匹配
// password: 要验证的明文密码
// hash: 存储的哈希值
// 返回: 是否匹配和错误信息
func VerifyPassword(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, fmt.Errorf("密码验证失败: %w", err)
	}

	return true, nil
}

// MD5Hash 计算字符串的MD5哈希值
// data: 要哈希的数据
// 返回: MD5哈希值的十六进制字符串
func MD5Hash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SHA1Hash 计算字符串的SHA1哈希值
// data: 要哈希的数据
// 返回: SHA1哈希值的十六进制字符串
func SHA1Hash(data string) string {
	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SHA256Hash 计算字符串的SHA256哈希值
// data: 要哈希的数据
// 返回: SHA256哈希值的十六进制字符串
func SHA256Hash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// SHA512Hash 计算字符串的SHA512哈希值
// data: 要哈希的数据
// 返回: SHA512哈希值的十六进制字符串
func SHA512Hash(data string) string {
	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

// HashWithSalt 使用盐值进行SHA256哈希
// data: 要哈希的数据
// salt: 盐值
// 返回: 加盐后的SHA256哈希值
func HashWithSalt(data, salt string) string {
	combined := data + salt
	return SHA256Hash(combined)
}

// GenerateSalt 生成指定长度的随机盐值
// length: 盐值长度
// 返回: 随机盐值和错误信息
func GenerateSalt(length int) (string, error) {
	if length <= 0 {
		length = 16 // 默认16字节
	}

	saltBytes, err := GenerateAESKey(length)
	if err != nil {
		return "", fmt.Errorf("生成盐值失败: %w", err)
	}

	return hex.EncodeToString(saltBytes), nil
}

// QuickHashPassword 快速密码哈希（使用默认强度）
// password: 要加密的密码
// 返回: 哈希后的密码和错误信息
func QuickHashPassword(password string) (string, error) {
	return HashPassword(password, bcrypt.DefaultCost)
}

// ComparePasswords 比较两个密码的哈希值是否相同
// password1, password2: 要比较的密码
// 返回: 是否相同
func ComparePasswords(password1, password2 string) bool {
	hash1 := SHA256Hash(password1)
	hash2 := SHA256Hash(password2)
	return hash1 == hash2
}
