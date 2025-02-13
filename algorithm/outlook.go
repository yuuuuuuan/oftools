package algorithm

import (
    "fmt"
    "github.com/go-ole/go-ole"
    "github.com/go-ole/go-ole/oleutil"
    "time"
)

func OutlookListen() error{
    // 初始化 OLE
    ole.CoInitialize(0)
    defer ole.CoUninitialize()

    // 连接到 Outlook
    outlook, err := oleutil.CreateObject("Outlook.Application")
    if err != nil {
        fmt.Println("无法创建 Outlook 对象:", err)
        return err
    }
    defer outlook.Release()

    app, err := outlook.QueryInterface(ole.IID_IDispatch)
    if err != nil {
        fmt.Println("无法获取 Outlook IDispatch 接口:", err)
        return err
    }
    defer app.Release()

    // 获取 MAPI 命名空间
    namespace, err := oleutil.CallMethod(app, "GetNamespace", "MAPI")
    if err != nil {
        fmt.Println("无法获取 MAPI 命名空间:", err)
        return err
    }
    defer namespace.Clear()

    // 选择收件箱
    inbox, err := oleutil.CallMethod(namespace.ToIDispatch(), "GetDefaultFolder", 6) // 6 = Inbox
    if err != nil {
        fmt.Println("无法获取收件箱:", err)
        return err 
    }
    defer inbox.Clear()

    // 获取邮件列表
    items, err := oleutil.GetProperty(inbox.ToIDispatch(), "Items")
    if err != nil {
        fmt.Println("无法获取邮件:", err)
        return err
    }
    defer items.Clear()

    // 监听邮件
    fmt.Println("开始监听 Outlook 收件箱...")
    for {
        count, _ := oleutil.GetProperty(items.ToIDispatch(), "Count")
        if count.Val > 0 {
            lastMail, _ := oleutil.CallMethod(items.ToIDispatch(), "GetLast")
            subject, _ := oleutil.GetProperty(lastMail.ToIDispatch(), "Subject")
            sender, _ := oleutil.GetProperty(lastMail.ToIDispatch(), "SenderName")
            fmt.Printf("新邮件: 来自 %s, 主题: %s\n", sender.ToString(), subject.ToString())
            lastMail.Clear()
        }
        time.Sleep(5 * time.Second) // 每 5 秒轮询
    }
}
