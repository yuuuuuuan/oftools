package encode

import (
	"testing"
	"unicode"
)

// 测试 Set 生成的字符串是否符合 28 位字母+数字格式
func TestSet(t *testing.T) {
	salt := "NF3260"
	result := Set(salt)

	// 1. 检查长度是否为 28
	if len(result) != 28 {
		t.Errorf("生成的字符串长度错误，期望: 28，实际: %d", len(result))
	}

	// 2. 确保只包含字母和数字
	for _, char := range result {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			t.Errorf("生成的字符串包含非字母数字字符: %c", char)
		}
	}
}

// 测试相同 salt 生成相同的哈希字符串
func TestSetDeterministic(t *testing.T) {
	salt := "NF3260"
	result1 := Set(salt)
	result2 := Set(salt)

	if result1 != result2 {
		t.Errorf("相同 salt 生成的字符串不一致: %s vs %s", result1, result2)
	}
}

// 测试不同 salt 生成不同的哈希字符串
func TestSetUnique(t *testing.T) {
	salt1 := "NF3260"
	salt2 := "NF3282"

	result1 := Set(salt1)
	result2 := Set(salt2)

	if result1 == result2 {
		t.Errorf("不同 salt 生成了相同的字符串: %s", result1)
	}
}
