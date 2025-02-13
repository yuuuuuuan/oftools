package ofhttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpGetyuanqu() (map[string]interface{}, error) {
	// 请求的 URL
	url := "http://192.168.124.126/client/manufactures"

	// 创建一个新的 POST 请求，没有请求体
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端并执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error executing request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// 解析 JSON 响应体
	var jsonResponse map[string]interface{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// 打印响应状态码
	fmt.Println("Response Status:", resp.Status)

	return jsonResponse, nil
}
