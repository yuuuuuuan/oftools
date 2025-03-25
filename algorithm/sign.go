package algorithm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"oftools/oflog"
)

mapping := map[string]string{
	"pingpong":  "wLRJm5Rf",
	"badminton": "LIgGfIxc",
}



func SignSingle(name string, id string) error {
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
			"u_id":    "NF3101",
			"u_name":  "万永存",
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
