package main

import (
	"dongpo_proxy/client/windows/view"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	manager := view.UIManager{}
	go manager.CreatePanel()
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	_ = <-c
	view.ShutdownCallback() //关闭时回调钩子
}
