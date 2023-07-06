package jwt_helper

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(user_id int) (string, error) {
	// 创建一个新的 token 对象，指定签名算法和 claims
	// jwt.SigningMethodHS256 是指 HMAC-SHA256 算法
	// jwt.MapClaims 是指标准的 payload 数据
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
	})
	// 生成 HMAC key
	key, err := getHMACKey()
	if err != nil {
		return "", err
	}

	// 使用 HMAC key 签名 token
	return token.SignedString(key)
}

func getHMACKey() ([]byte, error) {
	keyPath := viper.GetString("jwt.hmac.key_path") // 从配置文件中读取 key_path
	return ioutil.ReadFile(keyPath)                 // 从文件中读取 HMAC key
}
func GenerateHMACKey() ([]byte, error) {
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
