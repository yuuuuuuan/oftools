package algorithm

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"


	"github.com/PuerkitoBio/goquery"
)

func IworkGet() error {
	// Base URL
	targetURL := "https://it.ofilm.com/hr/hr-ks//rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/"

	// 发送 HTTP 请求
	resp, err := http.Get(targetURL)
	if err != nil {
		log.Fatalf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 解析 HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("解析 HTML 失败: %v", err)
	}

	// 获取所有符合 `/域名/*` 规则的链接
	fmt.Println("发现的链接:")
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		link, exists := element.Attr("href")
		if exists {
			parsedURL, err := url.Parse(link)
			if err == nil {
				// 确保链接属于同一域名，并且符合 `/域名/*` 格式
				if parsedURL.IsAbs() {
					// 绝对路径直接匹配
					if parsedURL.Host == "it.ofilm.com/hr/hr-ks//rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/" && parsedURL.Path != "/NF3266" {
						fmt.Println(parsedURL.String())
					}
				} else {
					// 处理相对路径
					fullURL := targetURL + parsedURL.Path
					fmt.Println(fullURL)
				}
			}
		}
	})
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
