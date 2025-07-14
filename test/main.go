package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("0d30Zyll20vDWf4Nezol2BFUvu20ZylU")

	// URL-safe Base64 编码（保留 padding）
	encoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("Encoded (with padding):", encoded)

	// URL-safe Base64 编码（去除 padding）
	encodedRaw := base64.RawURLEncoding.EncodeToString(data)
	fmt.Println("Encoded (no padding):", encodedRaw)

	// 解码（适配 Raw 或有填充的都支持）
	decoded, _ := base64.RawURLEncoding.DecodeString(encodedRaw)
	fmt.Println("Decoded:", string(decoded))
}
