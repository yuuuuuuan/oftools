package algorithm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"oftools/oflog"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

type RequestData struct {
	ExamId string `json:"examId"`
	UserNo string `json:"userNo"`
}

// 定义数据结构
type ListAnswer struct {
	Content string `json:"content"`
	IsRight int    `json:"isRight"`
}

type ListElqu struct {
	ID      int          `json:"id"`
	Content string       `json:"content"`
	Answers []ListAnswer `json:"listanswer"`
}

type ResponseBody struct {
	ElTestpaper struct {
		Listelqu []ListElqu `json:"listelqu"`
	} `json:"elTestpaper"`
}

type Output struct {
	ID      int          `json:"id"`
	Content string       `json:"content"`
	Answers []ListAnswer `json:"answers"`
}

// LoginResponse 定义返回结构体
type LoginResponse struct {
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
	Token string `json:"token"`
}

func OfyxGetquiz() error {
	reader := bufio.NewReader(os.Stdin)

	// 输入账号
	fmt.Print("请输入账号: ")
	usernameRaw, _ := reader.ReadString('\n')
	username := strings.TrimSpace(usernameRaw)

	// 输入密码（遮罩）
	fmt.Print("请输入密码: ")
	passwordBytes, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println() // 换行
	if err != nil {
		fmt.Println("读取密码失败:", err)
		return err
	}
	password := string(passwordBytes)

	// 输入题库
	fmt.Print("请输入题库: ")
	testpaperidRaw, _ := reader.ReadString('\n')
	testpaperid := strings.TrimSpace(testpaperidRaw)

	// 执行登录
	token, err := login(username, password)
	if err != nil {
		fmt.Println("❌ 登录失败:", err)
		return err
	}
	fmt.Println("✅ 登录成功，Token:")
	fmt.Println(token)

	requst(token, username, testpaperid)
	return nil
}

// Login 返回 token 或错误
func login(username, password string) (string, error) {
	url := "https://ofyx.ofilm.com/study-prod-api/login"

	// 构造请求体
	payload := map[string]string{
		"username": username,
		"password": password,
	}
	jsonBody, _ := json.Marshal(payload)

	// 构造请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	// 设置 Headers（模拟浏览器）
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Origin", "https://ofyx.ofilm.com")
	req.Header.Set("Referer", "https://ofyx.ofilm.com/dist/login?redirect=/home")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取并解析响应
	body, _ := io.ReadAll(resp.Body)

	var loginRes LoginResponse
	if err := json.Unmarshal(body, &loginRes); err != nil {
		return "", err
	}

	// 返回错误或token
	if loginRes.Code != 200 {
		return "", errors.New("登录失败: " + loginRes.Msg)
	}
	return loginRes.Token, nil
}

// 定义请求函数，接收token和username
func requst(token, username, testpaperid string) {
	url := "https://ofyx.ofilm.com/study-prod-api/exam/exam/selectExamById"
	headers := map[string]string{
		"Host":            "ofyx.ofilm.com",
		"Content-Length":  "36",
		"Accept":          "application/json, text/plain, */*",
		"Authorization":   "Bearer " + token, // 使用传入的token
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090c2d)XWEB/11581",
		"Content-Type":    "application/json;charset=UTF-8",
		"Origin":          "https://ofyx.ofilm.com",
		"Sec-Fetch-Site":  "same-origin",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Dest":  "empty",
		"Referer":         "https://ofyx.ofilm.com/dist/exam/examform?id=" + testpaperid,
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9",
		"Connection":      "keep-alive",
		"Cookie":          "Admin-Token=" + token, // 使用传入的token
	}

	// 定义请求的数据
	data := RequestData{
		ExamId: testpaperid,
		UserNo: username, // 使用传入的username
	}

	// 打开文件以保存响应数据
	file, err := os.Create("response1.json")
	if err != nil {
		oflog.Print.Errorf("无法创建文件: %v", err)
	}
	defer file.Close()

	// 设置请求频率间隔时间
	requestInterval := 500 * time.Millisecond // 设置请求之间的间隔时间，1秒
	requstTime := 100
	// 发送请求并保存响应数据
	for i := 0; i < requstTime; i++ {
		// 将请求数据编码为JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			oflog.Print.Errorf("无法编码JSON: %v", err)
		}

		// 创建请求
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			oflog.Print.Errorf("创建请求失败: %v", err)
		}

		// 设置请求头
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		// 执行请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			oflog.Print.Errorf("请求失败: %v", err)
		}
		defer resp.Body.Close()

		// 读取响应 body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			oflog.Print.Errorf("读取响应失败: %v", err)
		}

		// 解析JSON响应
		var response ResponseBody
		err = json.Unmarshal(body, &response)
		if err != nil {
			oflog.Print.Errorf("解析JSON失败: %v", err)
		}

		// 提取数据
		var output []Output
		for _, item := range response.ElTestpaper.Listelqu {
			answers := []ListAnswer{}
			for _, ans := range item.Answers {
				answers = append(answers, ListAnswer{
					Content: ans.Content,
					IsRight: ans.IsRight,
				})
			}
			output = append(output, Output{
				ID:      item.ID,
				Content: item.Content,
				Answers: answers,
			})
		}

		// 将 listelqu 内容写入文件，保存为JSON格式并去掉前后字符
		listelquJson, err := json.Marshal(output)
		// 转换为字符串并去掉首尾的中括号
		jsonStr := strings.Trim(string(listelquJson), "[]")
		listelquJson = []byte(jsonStr)
		if err != nil {
			oflog.Print.Errorf("无法编码 listelqu 数据为JSON: %v", err)
		}

		// 写入到文件中
		_, err = file.Write(listelquJson)
		if err != nil {
			oflog.Print.Errorf("写入文件失败: %v", err)
		}

		// 添加逗号
		_, err = file.WriteString(",\n")
		if err != nil {
			oflog.Print.Errorf("写入逗号失败: %v", err)
		}

		// 添加换行符
		_, err = file.WriteString("\n")
		if err != nil {
			oflog.Print.Errorf("写入换行符失败: %v", err)
		}

		// 输出请求完成的消息
		oflog.Print.Infof("请求 %d 完成，listelqu数据已保存", i+1)

		// 添加延迟，控制请求频率
		time.Sleep(requestInterval)
	}

	oflog.Print.Infof("请求已完成，listelqu数据已保存到response.json")

	// 第一步：读取 response1.txt 文件内容
	responseFile, err := os.Open("response1.json")
	if err != nil {
		oflog.Print.Infof("无法打开 response1.json:", err)
		return
	}
	defer responseFile.Close()

	var responseContent string
	scanner := bufio.NewScanner(responseFile)
	for scanner.Scan() {
		responseContent += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		oflog.Print.Errorf("读取 response1.txt 错误:", err)
		return
	}

	// 第二步：读取目标文件
	targetFile, err := os.Open("template.html")
	if err != nil {
		oflog.Print.Errorf("无法打开目标文件:", err)
		return
	}
	defer targetFile.Close()

	// 将目标文件的内容存入一个切片中
	var lines []string
	scanner = bufio.NewScanner(targetFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		oflog.Print.Errorf("读取目标文件错误:", err)
		return
	}

	// 第三步：在指定行列插入内容
	lineToInsert := 44   // 假设我们在第3行插入
	columnToInsert := 32 // 假设我们在第5列插入

	// 获取指定行的内容并插入到该列位置
	if lineToInsert-1 < len(lines) {
		line := lines[lineToInsert-1]
		// 将内容插入指定列位置
		line = line[:columnToInsert-1] + responseContent + line[columnToInsert-1:]
		lines[lineToInsert-1] = line
	}

	// 第四步：将修改后的内容写回到文件
	outputfilename := "index_" + testpaperid + ".html"
	outputFile, err := os.Create(outputfilename)
	if err != nil {
		oflog.Print.Errorf("无法创建输出文件:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	writer.Flush()

	oflog.Print.Infof("内容已成功插入到" + outputfilename + "文件中！")
}
