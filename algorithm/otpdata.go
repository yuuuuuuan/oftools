package algorithm

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// OtpdataGetSingle 读取多个行号内容，提取空格后的第一个字符串，
// 横向写入 CSV 文件；若文件存在则追加新行，并在首列写入文件名。
func OtpdataGetSingle(sourceDir string, nums []string) error {
	file, err := os.Open(sourceDir)
	if err != nil {
		return fmt.Errorf("无法打开文件: %v", err)
	}
	defer file.Close()

	// 读取文件所有行到内存
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("读取文件出错: %v", err)
	}

	// 检查 CSV 文件是否存在
	csvPath := "output.csv"
	fileExists := false
	if _, err := os.Stat(csvPath); err == nil {
		fileExists = true
	}

	// 打开 CSV 文件（存在则追加，不存在则创建）
	csvFile, err := os.OpenFile(csvPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("打开/创建 CSV 文件失败: %v", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// 提取文件名（不含路径）
	fileName := filepath.Base(sourceDir)

	var headers []string
	var values []string

	// 第一列：文件名
	headers = append(headers, "文件")
	values = append(values, fileName)

	// 遍历每个行号
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return fmt.Errorf("行号 %s 不是有效数字", numStr)
		}
		if num < 1 || num > len(lines) {
			return fmt.Errorf("行号 %d 超出文件范围", num)
		}

		line := strings.TrimSpace(lines[num-1])
		parts := strings.Fields(line)

		hexStr := fmt.Sprintf("0x%X", num-1)
		headers = append(headers, fmt.Sprintf("%s", hexStr))

		if len(parts) >= 2 {
			values = append(values, parts[1])
		} else {
			values = append(values, "")
		}
	}

	// 若文件不存在则先写入表头
	if !fileExists {
		writer.Write(headers)
	}

	// 写入一行值（追加模式）
	writer.Write(values)

	return nil
}

// OtpdataGetMuti 扫描指定目录下所有 .ini 文件并依次执行 OtpdataGetSingle
func OtpdataGetMuti(sourceDir string) error {
	// 检查目录是否存在
	info, err := os.Stat(sourceDir)
	if err != nil {
		return fmt.Errorf("无法访问路径 %s: %v", sourceDir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("%s 不是目录", sourceDir)
	}

	// 遍历目录下所有文件
	err = filepath.WalkDir(sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 只处理 .ini 文件
		if !d.IsDir() && filepath.Ext(path) == ".ini" {
			fmt.Printf("🔍 发现 INI 文件: %s\n", path)

			// 示例：指定要提取的行号
			nums := []string{"9", "9819", "9820", "9821", "9822", "9823", "9824", "9838", "9840", "9891", "9892", "9893", "9894", "9895", "9896", "9913", "9914", "9915", "9916", "9917", "9918"}

			// 调用前面定义的单文件处理函数
			err := OtpdataGetSingle(path, nums)
			if err != nil {
				fmt.Printf("⚠️ 处理文件 %s 时出错: %v\n", path, err)
			} else {
				fmt.Printf("✅ 已处理文件: %s\n", filepath.Base(path))
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("遍历目录失败: %v", err)
	}

	fmt.Println("✨ 所有 INI 文件已处理完成。")
	return nil
}
