package algorithm

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"oftools/oflog"
	"regexp"
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

type JsonBody struct {
	LoginID string `json:"loginid"`
	ID      string `json:"id"`
}

func OaResults(name string) error {
	url := "https://it.ofilm.com/ofilm-oa/oa/queryoajixiaojd"
	jsonBody := JsonBody{
		LoginID: name,
		ID:      "智能汽车研发二部",
	}
	body, err := json.Marshal(jsonBody)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	// ✅ 保留 gzip 但不设置 br/zstd，Go 支持自动解压 gzip；或者你想保留 br，必须手动解压
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip") // ✅ 建议只保留 gzip
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

	// ✅ 检查是否 gzip 响应
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("gzip 解压失败: %w", err)
		}
		defer gzReader.Close()
		reader = gzReader
	}

	respBody, err := io.ReadAll(reader)
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
	oflog.Print.Infof("姓名：%s（工号：%s）\n", result.Data[0].Name, result.Data[0].ClerkCode)
	for _, r := range result.Data {
		oflog.Print.Infof("✔ %s：%s", r.CYearPeriod, r.C7)
	}

	return nil
}

func OaInfo(name string) error {
	baseURL := "https://oa.ofilm.com/workflow/request/DataInputFrom.jsp"

	// 构造查询参数
	params := url.Values{}
	params.Set("id", "558445")
	params.Set("formid", "-916")
	params.Set("bill", "1")
	params.Set("node", "565381")
	params.Set("detailsum", "0")
	params.Set("trg", "field68618")
	params.Set("trgv", name+",")
	params.Set("rand", "1752162706921")
	params.Set("tempflag", "0.785598242543941")

	// 替换字段中的 NF3447 为 name
	params.Set("635861|field68618", name)
	params.Set("635863|field68663", "")
	params.Set("635865|field451051", "RSQH101")
	params.Set("636360|field68618", name)
	params.Set("635862|field68658", "")
	params.Set("635866|field68618", name)
	params.Set("636358|field68618", name)
	params.Set("635864|field68664", "")
	params.Set("636359|field68618", name)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// 构造请求
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置请求头
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-HK;q=0.8,en;q=0.7,zh-HK;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="137", "Chromium";v="137", "Not/A)Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	// 设置 Cookie
	req.Header.Set("Cookie", `wfcookie=0; yili_userid=NF3266; yili_token=2707dcb6-af45-4d13-a5a1-6069f161c873; loginfileweaver=%2Fwui%2Ftheme%2Fecology8%2Fpage%2Flogin.jsp%3FtemplateId%3D101%26logintype%3D1%26gopage%3D; loginidweaver=253272; languageidweaver=7; JSESSIONID=abcZEBsaemyz_7daqt9Fz; ecology_JSessionid=abcZEBsaemyz_7daqt9Fz`)

	// 发起请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应失败: %w", err)
	}

	js := string(body)

	re := regexp.MustCompile(`var\s+mainjs\s*=\s*"getElementByDocument\(window\.parent\.document,\s*\\?"(field\d{5})\\?"\)\.value\\?=\\"([^\\"]+)\\";`)
	
	matches := re.FindAllStringSubmatch(js, -1)
	for _, m := range matches {
		fmt.Printf("字段: %s, 值: %s\n", m[1], m[2])
	}
	return nil
}
