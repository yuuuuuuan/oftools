package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

//const baseurl = "http://192.168.124.126/client"

func main() {
	// 创建应用程序
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// 创建主窗口
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("V1.0.0.17")
	window.SetMinimumSize2(400, 300)

	// 创建中央部件
	centralWidget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(centralWidget)

	// 创建布局
	layout := widgets.NewQVBoxLayout()
	centralWidget.SetLayout(layout)

	// go func(ch chan []string) {
	// 	url := baseurl + "/manufactures"
	// 	resp, err := ofhttp.HttpPost(url, "")
	// 	if err != nil {
	// 		println("", err)
	// 	} else {
	// 		data, err := ofhttp.ConvertRespToJson(resp)
	// 		if err != nil {
	// 			println("", err)
	// 		} else {
	// 			var stringSlice []string
	// 			for key, value := range data {
	// 				// 使用类型断言检查值的类型
	// 				switch value.(type) {
	// 				case string:
	// 					stringSlice = append(stringSlice, key+" : "+value.(string))
	// 				case int:
	// 					stringSlice = append(stringSlice, key+" : "+fmt.Sprintf("%d", value.(int)))
	// 				case float64:
	// 					stringSlice = append(stringSlice, key+" : "+fmt.Sprintf("%f", value.(float64)))
	// 				default:
	// 					// 处理其他类型，例如struct等
	// 					stringSlice = append(stringSlice, key+" : "+fmt.Sprintf("%v", value))
	// 				}
	// 			}
	// 			ch <- stringSlice
	// 		}
	// 	}
	// }()
	// ch_yq := make(chan []string)
	domainLabel := widgets.NewQLabel2("园区：", nil, 0)
	domainComboBox := widgets.NewQComboBox(nil)
	domainComboBox.AddItems([]string{"测试", "生产", "其他"})

	domainLayout := widgets.NewQHBoxLayout()
	domainLayout.AddWidget(domainLabel, 0, 0)
	domainLayout.AddWidget(domainComboBox, 0, 0)
	layout.AddLayout(domainLayout, 0)

	// 显示窗口
	window.Show()

	// 运行 Qt 主循环
	app.Exec()
}

//CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" -o DownloadTools.exe interface.go
