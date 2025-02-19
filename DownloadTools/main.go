package main

import (
	"encoding/json"
	"log"
	"oftools/DownloadTools/ofhttp"
)

const baseurl = "http://192.168.124.126/client"

func main() {
	// // 创建应用程序
	// app := widgets.NewQApplication(len(os.Args), os.Args)

	// // 创建主窗口
	// window := widgets.NewQMainWindow(nil, 0)
	// window.SetWindowTitle("V1.0.0.17")
	// window.SetMinimumSize2(400, 300)

	// // 创建中央部件
	// centralWidget := widgets.NewQWidget(nil, 0)
	// window.SetCentralWidget(centralWidget)

	// // 创建布局
	// layout := widgets.NewQVBoxLayout()
	// centralWidget.SetLayout(layout)

	//var reqChan1, respChan1 chan map[string]interface{}
	//var reqChan2, respChan2 chan map[string]interface{}
	//var reqChan3, respChan3 chan map[string]interface{}
	reqChan1 := make(chan map[string]interface{})
	respChan1 := make(chan map[string]interface{})
	reqChan2 := make(chan map[string]interface{})
	respChan2 := make(chan map[string]interface{})
	reqChan3 := make(chan map[string]interface{})
	respChan3 := make(chan map[string]interface{})

	go ofhttp.SendPostRequset(baseurl + "/manufactures", reqChan1, respChan1)
	go ofhttp.SendPostRequset(baseurl + "/projects", reqChan2, respChan2)
	go ofhttp.SendPostRequset(baseurl + "/stations", reqChan3, respChan3)

	var input map[string]interface{}
	var err error
	input_json := `{}`
	err = json.Unmarshal([]byte(input_json), &input)
	if err != nil {
		log.Fatal(err)
	}
	reqChan1 <- input

	output := <-respChan1
	if output == nil {
		print("err")
		return
	}
	data, _ := ofhttp.ExtractDataAsStringMap(output)
	// 提取所有值并存入 []string
	var values []string
	for _, value := range data {
		values = append(values, value)
	}
	print(values)
	// domainLabel := widgets.NewQLabel2("园区：", nil, 0)
	// domainComboBox := widgets.NewQComboBox(nil)
	// domainComboBox.AddItems(values)

	// domainLayout := widgets.NewQHBoxLayout()
	// domainLayout.AddWidget(domainLabel, 0, 0)
	// domainLayout.AddWidget(domainComboBox, 0, 0)
	// layout.AddLayout(domainLayout, 0)
	// println("2")
	// // 显示窗口
	// window.Show()

	// // 运行 Qt 主循环
	// app.Exec()
}

//CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" -o DownloadTools.exe main.go
