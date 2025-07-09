package algorithm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OaResponse struct {
	Msg  string     `json:"msg"`
	Code string     `json:"code"`
	Data []OaRecord `json:"data"`
}

type OaRecord struct {
	ClerkCode   string `json:"clerkcode"`
	Name        string `json:"name"`
	CYearPeriod string `json:"cyearperiod"`
	CYear       string `json:"cyear"`
	CPeriod     string `json:"cperiod"`
	C7          string `json:"c_7"`
	Xzfl        string `json:"xzfl"`
}

func OaResults() error {
	url := "https://it.ofilm.com/ofilm-oa/oa/queryoajixiaojd"
	jsonBody := `{"loginid":"NF3266","id":"智能汽车研发二部"}`
	body := []byte(jsonBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// 添加 Headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-HK;q=0.8,en;q=0.7,zh-HK;q=0.6")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Origin", "https://oa.ofilm.com")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://oa.ofilm.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("sec-ch-ua", `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	var result OaResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	if result.Code != "0" {
		return fmt.Errorf("查询失败: %s", result.Msg)
	}

	if len(result.Data) == 0 {
		fmt.Println("无绩效数据")
		return nil
	}

	// 输出格式化结果
	fmt.Printf("姓名：%s（工号：%s）\n\n", result.Data[0].Name, result.Data[0].ClerkCode)
	for _, r := range result.Data {
		fmt.Printf("✔ %s：%s\n", r.CYearPeriod, r.C7)
	}

	return nil
}
