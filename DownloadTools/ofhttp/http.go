package ofhttp

import (
	"fmt"
	"net/http"
)

func HttpGetyuanqu() (*http.Response, error) {
	// 请求的 URL
	url := "http://192.168.124.126/client/manufactures"

	// 创建一个新的 POST 请求，没有请求体
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 创建 HTTP 客户端并执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error executing request: %v\n", err)
		return nil, err
	}

	// 注意：通常应在函数内部关闭响应体，以防止资源泄漏
	// 但如果要让调用者处理响应体，应移除以下行
	// defer resp.Body.Close()

	// 打印响应状态码
	fmt.Printf("Response Status: %v\n", resp.Status)

	return resp, nil
}
