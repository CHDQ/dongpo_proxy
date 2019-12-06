package main

import (
	"dongpo_proxy/server/biz"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start proxy server at " + time.Now().Format("2006-01-02 15:04:05"))
	biz.StartServer(":8080")
}
