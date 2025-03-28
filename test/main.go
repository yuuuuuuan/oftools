package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

// 使用盐（salt）生成固定长度的哈希字符串，仅包含字母和数字
func generateFixedString(salt string, length int) (string, error) {
	hash := sha256.New()
	hash.Write([]byte(salt))
	hashed := hash.Sum(nil)

	// 使用 Base64 标准编码（无填充）
	encoded := base64.RawStdEncoding.EncodeToString(hashed)

	// 只保留字母和数字
	encoded = strings.Map(func(r rune) rune {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		return -1 // 过滤掉非字母数字字符
	}, encoded)

	// 截取前 length 个字符
	if len(encoded) > length {
		encoded = encoded[:length]
	} else if len(encoded) < length {
		return "", nil // 确保总是返回足够长的字符串
	}

	return encoded, nil
}

func main() {
	// 给定盐
	salt := "NF3282"

	// 生成固定的 28 位字符
	result, err := generateFixedString(salt, 28)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result) // 输出固定长度的字符（28 位）
}
