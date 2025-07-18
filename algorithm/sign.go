package algorithm

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"oftools/encode"
	"oftools/oflog"
)

type Input struct {
	Id      string `json:"u_id"`
	Name    string `json:"u_name"`
	Dept    string `json:"u_dept"`
	Group   string `json:"u_group"`
	Park    string `json:"u_park"`
	Content string `json:"content"`
}

type Ifsuccess struct {
	Code int `json:"code"`
}

// 定义结构体
type Data struct {
	PsnCode string `json:"psncode"`
	PsnName string `json:"psnname"`
	Bm      string `json:"bm"`
	Syq     string `json:"syq"`
	Yq      string `json:"yq"`
}

type Response struct {
	Data Data `json:"data"`
}

func SignSingle(name string, id string) error {

	mapping := map[string]string{
		"pingpong":  "wLRJm5Rf",
		"badminton": "LIgGfIxc",
		"billiard":  "v0wudzzk",
	}

	// 检查键是否存在
	value, exists := mapping[name]
	if !exists {
		oflog.Print.Errorf("key '%s' not exsit", name)
		return nil
	}

	// Define the request URL
	url := "https://it.ofilm.com/hr/hr-ks//rest/kskinsfolk/kskinsfolk/findUserNoNcHrEK/" + id

	// Define the request body (empty JSON data)
	jsonData := []byte(`{}`)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		oflog.Print.Errorf("Error creating request:%s", err)
		return err
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
		oflog.Print.Errorf("Error sending request:%s", err)
		return err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		oflog.Print.Errorf("Error reading response:%s", err)
		return err
	}

	// // Print the response status and body
	// fmt.Printf("Status Code: %d\n", resp.StatusCode)
	// fmt.Printf("Response Body: %s\n", body)
	var ifsuccess Ifsuccess
	err = json.Unmarshal([]byte(body), &ifsuccess)
	if err != nil {
		oflog.Print.Errorf("JSON 解析错误:%s", err)
		return err
	}
	if ifsuccess.Code == 500 {
		oflog.Print.Errorf("User not exist")
		return nil
	}
	var result Response

	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		oflog.Print.Errorf("JSON 解析错误:%s", err)
		return err
	}

	fmt.Println("psncode:", result.Data.PsnCode)
	fmt.Println("psnname:", result.Data.PsnName)
	fmt.Println("bm:", result.Data.Bm)
	fmt.Println("syq:", result.Data.Syq)
	fmt.Println("yq:", result.Data.Yq)

	input := Input{
		Id:      result.Data.PsnCode,
		Name:    result.Data.PsnName,
		Dept:    result.Data.Bm,
		Group:   result.Data.Syq,
		Park:    result.Data.Yq,
		Content: value, // 可以添加额外的信息
	}

	if err = postsign(input); err != nil {
		oflog.Print.Errorf("%s Error:failed at algorithm.postsign!", getFunctionName())
		return err
	}

	return nil
}

func SignAuto(name string, num int) error {
	
	return nil
	
}

func postsign(input Input) error {

	url := "https://it.ofilm.com/Bus/admin/Bus/AddByBusRecord"
	openid := encode.Set(input.Id)

	jsonData, err := json.Marshal(input)
	if err != nil {
		oflog.Print.Errorf("JSON Encode failed.")
		return err
	}
	// new req
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		oflog.Print.Errorf("New Request failed.")
		return err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Xweb_xhr", "1")
	req.Header.Set("Openid", openid)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090c11)XWEB/11275")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://servicewechat.com/wx7222634c083face7/179/page-frame.html")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")

	// new http client req
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		oflog.Print.Errorf("Send Requst failed.")
		return err
	}
	defer resp.Body.Close()

	oflog.Print.Infof("%s statusCode for openid %s: %d\n", jsonData, openid, resp.StatusCode)
	return nil
}

func SignPingpong() error {
	// requst URL
	url := "https://it.ofilm.com/Bus/admin/Bus/AddByBusRecord"

	openids := []string{
		"olZaE61bCD9kOYCIJTKKfMeQxgsM",
		"olZaE61bCD9kOYCIJTKKfMe7aB5s",
		"olZaE61bCD9kOYCIJTKKfMeYjVd3",
		"olZaE61bCD9kOYCIJTKKfMehTQw9",
		"olZaE61bCD9kOYCIJTKKfMe3Fm8u",
		"olZaE61bCD9kOYCIJTKKfMevQ8wQ",
		"olZaE61bCD9kOYCIJTKKfMe6LjgD",
		"olZaE61bCD9kOYCIJTKKfMe1Q2lJ",
		"olZaE61bCD9kOYCIJTKKfMeZ4f9R",
		"olZaE61bCD9kOYCIJTKKfMeTm5Xq",
	}

	// Json Data at req Body
	data := []map[string]string{
		{
			"u_id":    "NF3266",
			"u_name":  "邹航远",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "朱毅林",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3247",
			"u_name":  "田武超",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3258",
			"u_name":  "马杰健",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "N0940570",
			"u_name":  "李明",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "N0940540",
			"u_name":  "黄安",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "N0940539",
			"u_name":  "陈其敏",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "N0940572",
			"u_name":  "叶俊旺",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3067",
			"u_name":  "辛洪亮",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3116",
			"u_name":  "马博",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},
	}

	// convert to json
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("JSON 编码失败:", err)
	// 	return err
	// }
	//fmt.Println(string(jsonData))
	//index := 0;
	// go through openid
	for v, openid := range openids {

		jsonData, err := json.Marshal(data[v])
		if err != nil {
			oflog.Print.Errorf("JSON Encode failed.")
			return err
		}
		// new req
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			oflog.Print.Errorf("New Request failed.")
			return err
		}

		// set header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Xweb_xhr", "1")
		req.Header.Set("Openid", openid)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090c11)XWEB/11275")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Sec-Fetch-Site", "cross-site")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Referer", "https://servicewechat.com/wx7222634c083face7/179/page-frame.html")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
		req.Header.Set("Connection", "keep-alive")

		// new http client req
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			oflog.Print.Errorf("Send Requst failed.")
			return err
		}
		defer resp.Body.Close()

		// return statusCode
		oflog.Print.Infof("%s statusCode for openid %s: %d\n", jsonData, openid, resp.StatusCode)
		// fmt.Println(string(jsonData))
		// fmt.Println(openid)
	}

	return nil
}

func SignBadminton() error {
	// requst URL
	url := "https://it.ofilm.com/Bus/admin/Bus/AddByBusRecord"

	openids := []string{
		//"olZaE61bCD9kOYCIJTKKfMeQxgsM",
		"olZaE61bCD9kOYCIJTKKfMe7aB5s",
		"olZaE61bCD9kOYCIJTKKfMeYjVd3",
		"olZaE61bCD9kOYCIJTKKfMehTQw9",
	}

	// Json Data at req Body
	data := []map[string]string{
		// {
		// 	"u_id":    "NF3266",
		// 	"u_name":  "邹航远",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "LIgGfIxc",
		// },

		{
			"u_id":    "NF3272",
			"u_name":  "朱毅林",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "LIgGfIxc",
		},

		{
			"u_id":    "NF3247",
			"u_name":  "田武超",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "LIgGfIxc",
		},

		{
			"u_id":    "NF3258",
			"u_name":  "马杰健",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "LIgGfIxc",
		},

		// {
		// 	"u_id":    "N0940570",
		// 	"u_name":  "李明",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "N0940540",
		// 	"u_name":  "黄安",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "N0940539",
		// 	"u_name":  "陈其敏",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "N0940572",
		// 	"u_name":  "叶俊旺",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "NF3101",
		// 	"u_name":  "万永存",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "NF3116",
		// 	"u_name":  "马博",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },
	}

	// convert to json
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("JSON 编码失败:", err)
	// 	return err
	// }
	//fmt.Println(string(jsonData))
	//index := 0;
	// go through openid
	for v, openid := range openids {

		jsonData, err := json.Marshal(data[v])
		if err != nil {
			oflog.Print.Errorf("JSON Encode failed.")
			return err
		}
		// new req
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			oflog.Print.Errorf("New Request failed.")
			return err
		}

		// set header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Xweb_xhr", "1")
		req.Header.Set("Openid", openid)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090c11)XWEB/11275")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Sec-Fetch-Site", "cross-site")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Referer", "https://servicewechat.com/wx7222634c083face7/179/page-frame.html")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
		req.Header.Set("Connection", "keep-alive")

		// new http client req
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			oflog.Print.Errorf("Send Requst failed.")
			return err
		}
		defer resp.Body.Close()

		// return statusCode
		oflog.Print.Infof("%s statusCode for openid %s: %d\n", jsonData, openid, resp.StatusCode)
		// fmt.Println(string(jsonData))
		// fmt.Println(openid)
	}

	return nil
}

func SignTest() error {
	// requst URL
	url := "https://it.ofilm.com/Bus/admin/Bus/AddByBusRecord"

	openids := []string{
		"olZaE61bCD9kOYCIJTKKfMeQxgsM",
		"olZaE61bCD9kOYCIJTKKfMe7aB5s",
		"olZaE61bCD9kOYCIJTKKfMeYjVd3",
		"olZaE61bCD9kOYCIJTKKfMehTQw9",
	}

	// Json Data at req Body
	data := []map[string]string{
		{
			"u_id":    "NF3266",
			"u_name":  "邹航远",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "朱毅林",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3247",
			"u_name":  "田武超",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3258",
			"u_name":  "马杰健",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		// {
		// 	"u_id":    "N0940570",
		// 	"u_name":  "李明",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "N0940540",
		// 	"u_name":  "黄安",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "N0940539",
		// 	"u_name":  "陈其敏",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "N0940572",
		// 	"u_name":  "叶俊旺",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "NF3101",
		// 	"u_name":  "万永存",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },

		// {
		// 	"u_id":    "NF3116",
		// 	"u_name":  "马博",
		// 	"u_dept":  "影像软件开发课",
		// 	"u_group": "影像事业部",
		// 	"u_park":  "南昌未来城",
		// 	"content": "wLRJm5Rf",
		// },
	}

	// convert to json
	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	fmt.Println("JSON 编码失败:", err)
	// 	return err
	// }
	//fmt.Println(string(jsonData))
	//index := 0;
	// go through openid
	for v, openid := range openids {

		jsonData, err := json.Marshal(data[v])
		if err != nil {
			oflog.Print.Errorf("JSON Encode failed.")
			return err
		}
		// new req
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			oflog.Print.Errorf("New Request failed.")
			return err
		}

		// set header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Xweb_xhr", "1")
		req.Header.Set("Openid", openid)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 MicroMessenger/7.0.20.1781(0x6700143B) NetType/WIFI MiniProgramEnv/Windows WindowsWechat/WMPF WindowsWechat(0x63090c11)XWEB/11275")
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Sec-Fetch-Site", "cross-site")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Referer", "https://servicewechat.com/wx7222634c083face7/179/page-frame.html")
		req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
		req.Header.Set("Connection", "keep-alive")

		// new http client req
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			oflog.Print.Errorf("Send Requst failed.")
			return err
		}
		defer resp.Body.Close()

		// return statusCode
		oflog.Print.Infof("%s statusCode for openid %s: %d\n", jsonData, openid, resp.StatusCode)
		// fmt.Println(string(jsonData))
		// fmt.Println(openid)
	}

	return nil
}
