package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/gui"
    "github.com/therecipe/qt/widgets"
)

// 定义结构体来解析JSON数据
type ResponseData struct {
    Data []string `json:"data"`
}

func main() {
    // 初始化Qt
    core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
    gui.QGuiApplication_SetDefaultColor Qt__white
    widgets.QApplication_SetOrganizationName("YourCompany")
    widgets.QApplication_SetOrganizationDomain("yourcompany.com")
    widgets.QApplication_SetApplicationName("YourApp")

    // 创建窗口
    window := widgets.NewQMainWindow(nil, 0)
    window.SetWindowTitle("QComboBox Demo")
    window.SetMinimumSize2(400, 300)

    // 创建主布局
    centralWidget := widgets.NewQWidget(nil, 0)
    layout := widgets.NewQVBoxLayout()

    // 创建QComboBox
    comboBox := widgets.NewQComboBox(nil)
    layout.addWidget(comboBox)

    // 将布局设置到主窗口
    centralWidget.SetLayout(layout)
    window.SetCentralWidget(centralWidget)

    // 显示窗口
    window.Show()

    // 发送POST请求
    go func() {
        url := "https://your-api-endpoint.com/data"
        method := "POST"
        
        // 创建请求
        req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte("{}")))
        if err != nil {
            fmt.Printf("创建请求失败: %v\n", err)
            return
        }

        // 设置请求头
        req.Header.Set("Content-Type", "application/json")

        // 发送请求
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            fmt.Printf("发送请求失败: %v\n", err)
            return
        }
        defer resp.Body.Close()

        // 检查状态码
        if resp.StatusCode != http.StatusOK {
            fmt.Printf("HTTP错误: %v\n", resp.Status)
            return
        }

        // 读取响应体
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            fmt.Printf("读取响应失败: %v\n", err)
            return
        }

        // 解析JSON
        var data ResponseData
        err = json.Unmarshal(body, &data)
        if err != nil {
            fmt.Printf("解析JSON失败: %v\n", err)
            return
        }

        // 将数据通过channel发送到主线程
        dataChan <- data
    }()

    // 在主线程更新GUI
    dataChan := make(chan ResponseData)
    go func() {
        data := <-dataChan
        // 更新QComboBox中的选项
        comboBox.Clear()
        for _, item := range data.Data {
            comboBox.AddItems([]string{item})
        }
    }()

    // 运行Qt主循环
    core.QCoreApplication_Exec()
}
