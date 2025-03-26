package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

// 使用盐（salt）和其他输入数据生成固定的哈希值
func generateFixedString(salt string, length int) (string, error) {
	// 使用 SHA-256 哈希算法
	hash := sha256.New()

	// 将盐值作为输入
	hash.Write([]byte(salt))

	// 获取哈希值
	hashed := hash.Sum(nil)

	// 使用 Base64 编码
	encoded := base64.URLEncoding.EncodeToString(hashed)

	// 如果生成的字符串长度大于目标长度，则截取
	if len(encoded) > length {
		encoded = encoded[:length]
	}

	return encoded, nil
}

func main() {
	// 给定盐
	salt := "NF3248"

	// 生成固定的 28 位字符
	result, err := generateFixedString(salt, 28)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result) // 输出固定长度的字符（28 位）
}
