package main

import (
	"encoding/json"
	"log"
	"oftools/DownloadTools/ofhttp"
	"time"
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

	go ofhttp.SendPostRequset(baseurl+"/manufactures", reqChan1, respChan1)
	go ofhttp.SendPostRequset(baseurl+"/projects", reqChan2, respChan2)
	go ofhttp.SendPostRequset(baseurl+"/stations", reqChan3, respChan3)

	var input map[string]interface{}
	var err error
	input_json := `{}`
	err = json.Unmarshal([]byte(input_json), &input)
	if err != nil {
		log.Println(err)
	}
	reqChan1 <- input
	time.Sleep(500 * time.Millisecond)
	output := <-respChan1

	data, err := ofhttp.ExtractDataAsStringMap(output)
	if err != nil {
		log.Println(err)
	}

	var values []string
	for _, v := range data {
		// 使用类型断言确认值是 string 类型
		if str, ok := v.(string); ok {
			values = append(values, str)
		}
	}

	log.Println(values)
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
