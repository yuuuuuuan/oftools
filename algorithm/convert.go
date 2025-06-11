package algorithm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Conversion struct to store values in all numeral systems
type Conversion struct {
	Hex string
	Dec int64
	Bin string
	Oct string
}

// ConvertHexToOthers converts a hexadecimal value to decimal, binary, and octal
func ConvertHexToOthers(hexValue string) (Conversion, error) {
	// Remove the '0x' prefix if it exists
	hexValue = strings.TrimPrefix(hexValue, "0x")

	// Convert hex to decimal
	decimalValue, err := strconv.ParseInt(hexValue, 16, 64)
	if err != nil {
		return Conversion{}, fmt.Errorf("invalid hex value: %v", err)
	}

	// Convert decimal to binary and octal
	binaryValue := fmt.Sprintf("%b", decimalValue)
	octalValue := fmt.Sprintf("%o", decimalValue)
	hexValue = fmt.Sprintf("0x%X", decimalValue)

	// Return the Conversion struct with all the systems
	return Conversion{
		Hex: hexValue,
		Dec: decimalValue,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}

// ConvertDecToOthers converts a decimal value to hex, binary, and octal
func ConvertDecToOthers(decimalValue string) (Conversion, error) {
	value, err := strconv.ParseInt(decimalValue, 10, 64)
	if err != nil {
		return Conversion{}, fmt.Errorf("Error:%v", err)

	}
	// Convert decimal to hex, binary, and octal
	hexValue := fmt.Sprintf("0x%X", value)
	binaryValue := fmt.Sprintf("%b", value)
	octalValue := fmt.Sprintf("%o", value)

	return Conversion{
		Hex: hexValue,
		Dec: value,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}

// ConvertBinToOthers converts a binary value to hex, decimal, and octal
func ConvertBinToOthers(binaryValue string) (Conversion, error) {
	// Convert binary to decimal
	decimalValue, err := strconv.ParseInt(binaryValue, 2, 64)
	if err != nil {
		return Conversion{}, fmt.Errorf("invalid binary value: %v", err)
	}

	// Convert to hex and octal
	hexValue := fmt.Sprintf("0x%X", decimalValue)
	octalValue := fmt.Sprintf("%o", decimalValue)

	return Conversion{
		Hex: hexValue,
		Dec: decimalValue,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}

// ConvertOctToOthers converts an octal value to hex, decimal, and binary
func ConvertOctToOthers(octalValue string) (Conversion, error) {
	// Convert octal to decimal
	decimalValue, err := strconv.ParseInt(octalValue, 8, 64)
	if err != nil {
		return Conversion{}, fmt.Errorf("invalid octal value: %v", err)
	}

	// Convert to hex and binary
	hexValue := fmt.Sprintf("0x%X", decimalValue)
	binaryValue := fmt.Sprintf("%b", decimalValue)

	return Conversion{
		Hex: hexValue,
		Dec: decimalValue,
		Bin: binaryValue,
		Oct: octalValue,
	}, nil
}

// BinaryToHexFile reads a text file containing binary strings line by line,
// converts every 8 bits into a byte, and writes the resulting binary data to the output file.
func BinaryToHexFile(inputPath, outputPath string) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("无法打开输入文件: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("无法创建输出文件: %v", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// 移除所有空白字符
		cleaned := strings.Map(func(r rune) rune {
			if r == '0' || r == '1' {
				return r
			}
			return -1
		}, line)

		for i := 0; i < len(cleaned); i += 8 {
			//end := i + 8
			byteStr := cleaned[i:]
			if len(byteStr) > 8 {
				byteStr = byteStr[:8]
			} else if len(byteStr) < 8 {
				// 补齐末尾不足8位的字节
				byteStr += strings.Repeat("0", 8-len(byteStr))
			}

			// 二进制字符串转为字节
			var b byte
			_, err := fmt.Sscanf(byteStr, "%08b", &b)
			if err != nil {
				return fmt.Errorf("第 %d 行处理失败，二进制无效: %v", lineNumber, err)
			}

			_, err = outputFile.Write([]byte{b})
			if err != nil {
				return fmt.Errorf("写入文件失败: %v", err)
			}
		}

		fmt.Printf("✅ 已写入第 %d 行\n", lineNumber)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}

	return nil
}
