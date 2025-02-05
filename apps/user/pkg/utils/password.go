package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

const saltLength = 16

// GenerateSalt 生成随机盐值
func GenerateSalt() (string, error) {
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// HashPassword 密码加盐哈希
func HashPassword(password string, salt string) string {
	combined := password + salt
	hasher := sha256.New()
	hasher.Write([]byte(combined))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
} 