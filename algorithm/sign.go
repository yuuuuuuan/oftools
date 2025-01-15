package algorithm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignPingpong() error {
	// requst URL
	url := "https://it.ofilm.com/Bus/admin/Bus/AddByBusRecord"

	openids := []string{
		"olZaE61bCD9kOYCIJTKKfMeQxgsM",
		"olZaE61bCD9kOYCIJTKKfMe7aB5sZl",
		"olZaE61bCD9kOYCIJTKKfMeYjVd3F7",
		"olZaE61bCD9kOYCIJTKKfMehTQw9dZ",
		"olZaE61bCD9kOYCIJTKKfMe3Fm8u5P",
		"olZaE61bCD9kOYCIJTKKfMevQ8wQ4p",
		"olZaE61bCD9kOYCIJTKKfMe6LjgDqU",
		"olZaE61bCD9kOYCIJTKKfMe1Q2lJ8Y",
		"olZaE61bCD9kOYCIJTKKfMeZ4f9R1A",
		"olZaE61bCD9kOYCIJTKKfMeTm5Xq0M",
	}

	// Json Data at req Body
	data := []map[string]string{
		{
			"u_id":    "NF3266",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "wLRJm5Rf",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
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
			fmt.Println("JSON 编码失败:", err)
			return err
		}
		// new req
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("请求创建失败:", err)
			return err
		}

		// set header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Xweb_xhr", "1")
		req.Header.Set("Openid", openid) // 将 openid 设置在请求头
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
			fmt.Println("请求发送失败:", err)
			return err
		}
		defer resp.Body.Close()

		// return statusCode
		fmt.Printf("响应状态码 for openid %s: %d\n", openid, resp.StatusCode)
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
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "NF3272",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "NF3247",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "NF3258",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "N0940570",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "N0940540",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "N0940539",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "N0940572",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "NF3101",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
		},

		{
			"u_id":    "NF3116",
			"u_name":  "",
			"u_dept":  "影像软件开发课",
			"u_group": "影像事业部",
			"u_park":  "南昌未来城",
			"content": "m8E3XeiQ",
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
			fmt.Println("JSON 编码失败:", err)
			return err
		}
		// new req
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("请求创建失败:", err)
			return err
		}

		// set header
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Xweb_xhr", "1")
		req.Header.Set("Openid", openid) // 将 openid 设置在请求头
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
			fmt.Println("请求发送失败:", err)
			return err
		}
		defer resp.Body.Close()

		// return statusCode
		fmt.Printf("响应状态码 for openid %s: %d\n", openid, resp.StatusCode)
		// fmt.Println(string(jsonData))
		// fmt.Println(openid)
	}

	return nil
}
