package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

//baseurl := "http://192.168.124.126/client"

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

	// —— 第一行：域（测试/生产/其他） ——
	domainLabel := widgets.NewQLabel2("域：", nil, 0)
	domainComboBox := widgets.NewQComboBox(nil)
	domainComboBox.AddItems([]string{"测试", "生产", "其他"})

	domainLayout := widgets.NewQHBoxLayout()
	domainLayout.AddWidget(domainLabel, 0, 0)
	domainLayout.AddWidget(domainComboBox, 0, 0)
	layout.AddLayout(domainLayout, 0)

	// —— 第二行：名（示例：O1000-N-NULL-CTS） ——
	fileLabel := widgets.NewQLabel2("名：", nil, 0)
	fileEdit := widgets.NewQLineEdit(nil)
	fileEdit.SetText("O1000-N-NULL-CTS")

	fileLayout := widgets.NewQHBoxLayout()
	fileLayout.AddWidget(fileLabel, 0, 0)
	fileLayout.AddWidget(fileEdit, 0, 0)
	layout.AddLayout(fileLayout, 0)

	// —— 第三行：机（示例：OTP1 / NTS） ——
	choiceLabel := widgets.NewQLabel2("机：", nil, 0)
	choiceComboBox := widgets.NewQComboBox(nil)
	choiceComboBox.AddItems([]string{"OTP1", "NTS"})

	choiceLayout := widgets.NewQHBoxLayout()
	choiceLayout.AddWidget(choiceLabel, 0, 0)
	choiceLayout.AddWidget(choiceComboBox, 0, 0)
	layout.AddLayout(choiceLayout, 0)

	// —— 第四行：下载按钮 ——
	downloadButton := widgets.NewQPushButton2("下载", nil)
	downloadButton.ConnectClicked(func(bool) {
		// 这里添加下载逻辑，如HTTP下载或其他操作
		widgets.QMessageBox_Information(
			nil,
			"提示",
			"开始下载: "+fileEdit.Text(),
			widgets.QMessageBox__Ok,
			widgets.QMessageBox__Ok,
		)
	})
	layout.AddWidget(downloadButton, 0, core.Qt__AlignCenter)

	// 设置中心部件并显示窗口
	window.SetCentralWidget(centralWidget)
	window.Show()

	// 4. 执行应用
	app.Exec()
}

//CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui -s -w" -o DownloadTools.exe interface.go
