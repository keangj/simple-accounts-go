package jwt_helper

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user_id int) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
	})

	key, err := generateHMACKey()
	if err != nil {
		return "", err
	}

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(key)
}

func generateHMACKey() ([]byte, error) {
	// 生成 HMAC key
	// key 是一个 64 字节的随机字节序列
	key := make([]byte, 64)
	// 从 rand.Reader 中读取随机字节序列填充 key
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate HMAC key: %w", err)
	}
	return key, nil
}
