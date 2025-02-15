package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

//baseurl := "http://192.168.124.126/client"

type Interface struct {
	Yq []int
	Lx int
	Cs map[string]string
}

func main() {
	// 1. 初始化QApplication
	app := widgets.NewQApplication(len(os.Args), os.Args)
	// 2. 创建主窗口
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Go + Qt桌面客户端示例")
	window.Resize2(400, 200)

	// 3. 创建中央部件和布局
	centralWidget := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQVBoxLayout2(centralWidget)

	// 创建QComboBox
	comboBox := widgets.NewQComboBox(nil)
	//comboBox.SetPlaceholderText("请选择一个选项")
	comboBox.AddItem("选项1", nil)
	comboBox.AddItem("选项2", nil)
	comboBox.AddItem("选项3", nil)

	// 设置默认选项为“选项1”
	comboBox.SetCurrentIndex(0)

	// 将QComboBox添加到布局中
	layout.AddWidget(comboBox, 0, 0)

	// 设置中心部件并显示窗口
	window.SetCentralWidget(centralWidget)
	window.Show()

	// 4. 执行应用
	app.Exec()
}

//CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" -o DownloadTools.exe interface.go
