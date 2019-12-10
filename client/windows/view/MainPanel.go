package view

import (
	"dongpo_proxy/client/proxy/http"
	"encoding/json"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"
)

type UIManager struct {
	LocalPort *ui.Entry
	RPCIp     *ui.Entry
	RPCPort   *ui.Entry
	Info      *ui.Label
	CommitBtn *ui.Button
	CancelBtn *ui.Button
}

func (uiManager *UIManager) initParam() {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}
	dict := make(map[string]string)
	err = json.Unmarshal(data, &dict)
	if err != nil {
		return
	}
	if _, ok := dict["local.port"]; ok {
		uiManager.LocalPort.SetText(dict["local.port"])
	}
	if _, ok := dict["rpc.ip"]; ok {
		uiManager.RPCIp.SetText(dict["rpc.ip"])
	}
	if _, ok := dict["rpc.port"]; ok {
		uiManager.RPCPort.SetText(dict["rpc.port"])
	}

}

func (uiManager *UIManager) storeConfig(localPort string, rpcIp string, rpcPort string) {
	dict := make(map[string]string)
	dict["local.port"] = localPort
	dict["rpc.ip"] = rpcIp
	dict["rpc.port"] = rpcPort
	data, err := json.Marshal(dict)
	if err == nil {
		ioutil.WriteFile("./config.json", data, os.ModeAppend)
	}
}
func (uiManager *UIManager) CreatePanel() {
	err := ui.Main(func() {
		uiManager.LocalPort = ui.NewEntry()
		uiManager.RPCIp = ui.NewEntry()
		uiManager.RPCPort = ui.NewEntry()
		uiManager.Info = ui.NewLabel(``)
		// 生成：按钮
		uiManager.CommitBtn = ui.NewButton(`启动`)
		uiManager.CancelBtn = ui.NewButton(`停止`)
		// 设置：按钮点击事件
		uiManager.CommitBtn.OnClicked(uiManager.commitButtonClick)
		uiManager.CancelBtn.OnClicked(uiManager.cancelButtonClick)
		uiManager.CancelBtn.Disable()
		// 生成：垂直容器
		box := ui.NewVerticalBox()
		// 往 垂直容器 中添加 控件
		box.Append(ui.NewLabel(`本地监听端口：`), false)
		box.Append(uiManager.LocalPort, false)
		box.Append(ui.NewLabel(`代理服务器ip：`), false)
		box.Append(uiManager.RPCIp, false)
		box.Append(ui.NewLabel(`代理服务器端口：`), false)
		box.Append(uiManager.RPCPort, false)
		box.Append(uiManager.CommitBtn, false)
		box.Append(uiManager.CancelBtn, false)
		box.Append(uiManager.Info, false)
		// 生成：窗口（标题，宽度，高度，是否有 菜单 控件）
		window := ui.NewWindow(`dongpo proxy`, 300, 250, false)
		window.SetMargined(true)
		// 窗口容器绑定
		window.SetChild(box)

		// 设置：窗口关闭时
		window.OnClosing(func(*ui.Window) bool {
			// 窗体关闭
			ui.Quit()
			return true
		})
		uiManager.initParam()
		// 窗体显示
		window.Show()
	})
	if err != nil {
		panic(err)
	}

}
func (uiManager *UIManager) commitButtonClick(button *ui.Button) {
	rpcIp := uiManager.RPCIp.Text()
	localPort := uiManager.LocalPort.Text()
	rpcPort := uiManager.RPCPort.Text()
	address := net.ParseIP(rpcIp)
	if address == nil {
		uiManager.Info.SetText("代理服务器ip格式错误")
		return
	}
	localPortNum, err := strconv.Atoi(localPort)
	if err != nil || localPortNum < 1024 || localPortNum > 65535 {
		uiManager.Info.SetText("监听端口格式错误")
		return
	}
	rpcPortNum, err := strconv.Atoi(rpcPort)
	if err != nil {
		uiManager.Info.SetText("代理服务器端口格式错误")
		return
	}
	if rpcPortNum >= 1024 && rpcPortNum <= 65535 {
		button.Disable()
		uiManager.CancelBtn.Enable()
		uiManager.Info.SetText("start proxy client at " + time.Now().Format("2006-01-02 15:04:05"))
		uiManager.storeConfig(localPort, rpcIp, rpcPort)
		go func() {
			err := http.StartClient(":"+localPort, rpcIp+":"+rpcPort)
			if err != nil {
				uiManager.Info.SetText("[ERROR] " + err.Error())
				button.Enable()
				uiManager.CancelBtn.Disable()
				return
			}
		}()
		return
	}
	uiManager.Info.SetText("代理服务器端口格式错误")
}
func (uiManager *UIManager) cancelButtonClick(button *ui.Button) {
	http.ShutdownClient()
	uiManager.Info.SetText("stop proxy client at " + time.Now().Format("2006-01-02 15:04:05"))
	button.Disable()
	uiManager.CommitBtn.Enable()
}
